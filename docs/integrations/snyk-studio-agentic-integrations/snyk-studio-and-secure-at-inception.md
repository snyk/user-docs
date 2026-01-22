# Secure at Inception with Snyk Studio

## Prerequisites

* An active Snyk user account
* Access to an Endpoint Management tool, such as JAMF

## Custom rules at scale

Snyk provides the following directives which enable issue prevention and remediation at scale:

* Run the Snyk Code scanning tool for new first-party code generated.
* Run the Snyk SCA scanning tool for new dependencies or dependency updates.
* Run the Snyk IaC scanning tool for new Infrastructure as Code updates.
* Fix security issues from new or modified code or dependencies using the results context from Snyk.
* Rescan the code after fixing issues to verify they are resolved and no new issues are introduced.
* Repeat this process until no issues are found.

Snyk offers [reference scripts](snyk-studio-and-secure-at-inception.md#sample-scripts) for deploying the IDE extension on developer machines with default settings. You can customize these scripts as needed.

{% hint style="info" %}
Snyk recommends leveraging the Snyk IDE extension in order to install and configure Snyk Studio. This is not mandatory. Contact the Snyk account team for other deployment and management options.
{% endhint %}

## Execution frequency setting

Set the frequency of scanning and initiation of fixes for AI-generated code in the VS Code Snyk IDE extension.

<figure><img src="../../.gitbook/assets/Security at inception setting.jpg" alt=""><figcaption><p>Scan and fix initiation frequency settings in the IDE</p></figcaption></figure>

## Configuration at scale

{% hint style="info" %}
For broad deployment, use an MDM or Endpoint Management tool like JAMF to target [MCP configuration](snyk-studio-and-secure-at-inception.md#install-the-extension-and-configure-the-snyk-mcp-server) to developer devices.
{% endhint %}

To configure using JAMF, separate targeting for groups and configurations:

* For machines with Windsurf installed, target the Snyk Windsurf MCP [scripts](https://learn.jamf.com/en-US/bundle/jamf-pro-documentation-current/page/Scripts.html).
* For machines with Cursor installed, target the Snyk Cursor MCP scripts.
* For machines where neither is installed, deploy the desired [.pkg](https://learn.jamf.com/en-US/bundle/jamf-pro-documentation-current/page/Package_Deployment.html) and configure matching scripts to run after installation.

Configure the MDM to reapply the desired [custom rules](snyk-studio-and-secure-at-inception.md#custom-rules-at-scale) across all endpoints.

## Sample scripts

To configure the IDE and MCP:

* Ensure that the MCP server shows in the application. In Windsurf, navigate to **Windsurf** > **Settings** > **Advanced settings** > **Cascade** > **MCP Servers** > **Manage MCPs**.
* Ensure that the user is prompted to trust Snyk. The code scan does not work if trust is not provided.
* If the Snyk directives are configured, ensure that the rules file snyk\_rules.md includes the rules and that **Activation Mode** is set to **Always On**.

{% hint style="info" %}
If you encounter any issues, visit [Troubleshooting](troubleshooting.md).
{% endhint %}

### Install the extension and configure the Snyk MCP server

```sh
#!/bin/bash

set -euo pipefail

readonly INSTALL_PLUGIN_ID="snyk-security.snyk-vulnerability-scanner"
readonly REMOVE_PLUGIN_ID="snyk-security.snyk-vulnerability-scanner-preview"

# Set by find_editor_cmd
EDITOR_CMD=""
# Set by fix_zscaler_cert_path
CERT_FILE=""

# --- Logging helpers ---
log_info() { echo "[INFO] $1"; }
log_error() { echo "[ERROR] $1" >&2; }
# -----------------------

# --- Console user helpers ---
get_console_user() { stat -f%Su /dev/console; }

get_console_home() {
    local user
    user="$(get_console_user)"
    launchctl asuser "$(id -u "$user")" sudo -u "$user" /bin/sh -lc 'printf %s "$HOME"'
}

run_as_user() {
    local user uid
    user="$(get_console_user)"
    uid="$(id -u "$user")"
    launchctl asuser "$uid" sudo -u "$user" "$@"
}
# ----------------------------

# --- PATH setup ---
build_path() {
    local editor_paths="$1"
    local user_home="$(get_console_home)"
    echo "${editor_paths}:/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"
}

# Find a usable editor (absolute paths first, then PATH)
find_editor_cmd() {
    local name="$1"
    local paths="$2"

    IFS=':' read -ra path_array <<< "$paths"
    for path in "${path_array[@]}"; do
        local candidate="${path}/$name"
        if run_as_user test -x "$candidate"; then
            EDITOR_CMD="$candidate"
            return 0
        fi
    done

    local user_path resolved
    user_path="$(build_path "$paths")"

    if resolved=$(run_as_user /bin/sh -lc "PATH='${user_path}':\"\\$PATH\"; command -v $name 2>/dev/null || true"); then
        if [[ -n "$resolved" ]]; then
            EDITOR_CMD="$resolved"
            return 0
        fi
    fi

    return 1
}
# ------------------

# --- Zscaler certificate fix ---
fix_zscaler_cert_path() {
    # Check if a Zscaler root certificate is present in the system keychain
    local zscaler_certs
    if zscaler_certs=$(security find-certificate -c "Zscaler Root CA" -p /Library/Keychains/System.keychain 2>/dev/null); then
        log_info "Zscaler certificate found in keychain."

        # Create a temporary file to hold the certificate
        CERT_FILE=$(run_as_user mktemp --tmpdir=/tmp zscaler-cert-XXXXXXXXXX.pem)

        if ! run_as_user printf %s "${zscaler_certs}" > "$CERT_FILE"; then
            log_error "Failed to export Zscaler certificate to a temporary file."
            exit 1
        fi

        log_info "Exported certificate to a temporary file for use during installation."

        # Set a trap to ensure the temporary file is removed on exit
        trap 'run_as_user rm -f "$CERT_FILE"' EXIT
    fi
}
# -------------------------------

# --- Plugin functions ---
editor_installed() {
    local name="$1"
    local paths="$2"

    if find_editor_cmd "$name" "$paths"; then
        if run_as_user "${EDITOR_CMD}" --version >/dev/null 2>&1; then
            log_info "${name} installed (using: ${EDITOR_CMD})"
            return 0
        fi
    fi

    log_info "${name} not installed, skipping"
    return 1
}

plugin_installed() {
    local name="$1" extensions_output

    if ! extensions_output=$(run_as_user "${EDITOR_CMD}" --list-extensions 2>/dev/null); then
        log_error "Failed to list extensions via ${EDITOR_CMD}"
        exit 1
    fi

    if echo "${extensions_output}" | grep -qx "${INSTALL_PLUGIN_ID}"; then
        log_info "Snyk plugin already installed in ${name}, skipping"
        return 0
    fi

    log_info "Snyk plugin not yet installed in ${name}"
    return 1
}

install_plugin() {
    local name="$1" paths="$2"
    local user_home extdir output rc user_path

    user_home="$(get_console_home)"
    user_path="$(build_path "$paths")"
    extdir="${user_home}/.${name}/extensions"
    run_as_user /bin/mkdir -p "${extdir}"

    log_info "Installing Snyk plugin for ${name}"

    log_info "Removing stable extension, if present"
    run_as_user /bin/sh -lc \
        "HOME='${user_home}' PATH='${user_path}' \
        '${EDITOR_CMD}' --uninstall-extension '${REMOVE_PLUGIN_ID}'" 2>&1 || true

    # Fix the SSL certificate issue if Zscaler is present
    fix_zscaler_cert_path

    local install_cmd="HOME='${user_home}' PATH='${user_path}' '${EDITOR_CMD}' --extensions-dir '${extdir}' --install-extension '${INSTALL_PLUGIN_ID}' --force"

    # Add the Zscaler certificate to the environment if it was found
    if [[ -n "$CERT_FILE" ]]; then
        install_cmd="NODE_EXTRA_CA_CERTS='$CERT_FILE' $install_cmd"
        log_info "Using temporary Zscaler certificate for installation."
    else
        log_info "Zscaler certificate not found, proceeding without certificate override."
    fi

    # Run the installation command
    output=$(run_as_user /bin/sh -lc "$install_cmd" 2>&1) || rc=$?

    if [[ -z "${rc:-}" || "${rc}" -eq 0 ]]; then
        log_info "Plugin installed for ${name}"
        return 0
    fi

    log_error "Install failed (exit ${rc}) for ${name}:"
    printf '%s\n' "${output}" >&2
    exit 1
}
# --------------------------

# --- Main execution ---
main() {
    log_info "Starting Snyk plugin installation script"

    ide_paths=(
        "windsurf|$(get_console_home)/.codeium/windsurf/bin:/Applications/Windsurf.app/Contents/Resources/app/bin"
        "cursor|/Applications/Cursor.app/Contents/Resources/app/bin"
    )

    for ide in "${ide_paths[@]}"; do
        IFS='|' read -r name paths <<< "$ide"

        log_info "Processing ${name}..."
        if ! editor_installed "$name" "$paths"; then
            continue
        fi

        if plugin_installed "$name"; then
            continue
        fi

        install_plugin "$name" "$paths"

        log_info "${name} processing complete"
    done

    log_info "Snyk plugin installation script completed successfully"
}

main "$@"
```

### Enable Snyk directives by applying the rules in the Project

```sh
#!/bin/bash
set -u -o pipefail

log_info(){ echo "[INFO] $1"; }
log_error(){ echo "[ERROR] $1" >&2; }
log_warning(){ echo "[WARNING] $1" >&2; }

# --- console user helpers ---
get_console_user(){ stat -f%Su /dev/console; }
get_console_home(){
  local user; user="$(get_console_user)"
  launchctl asuser "$(id -u "$user")" sudo -u "$user" /bin/sh -lc 'printf %s "$HOME"'
}
run_as_user(){
  local user uid; user="$(get_console_user)"; uid="$(id -u "$user")"
  launchctl asuser "$uid" sudo -u "$user" "$@"
}

# --- jq binary management ---
get_system_architecture(){
  local arch
  arch="$(uname -m)"
  case "$arch" in
    arm64|aarch64) echo "arm64" ;;
    x86_64|amd64) echo "amd64" ;;
    *)
      log_error "Unsupported architecture: $arch"
      return 1
      ;;
  esac
}

JQ_VERSION="1.8.1"
JQ_HASH_MACOS_ARM64="d1e04871ef93ffd2807a00d7edfa305b"
JQ_HASH_MACOS_AMD64="d91812b3fbcc20ae2e1f28a9b8141c67"

JQ_BINARY_PATH="/tmp/snyk-ide-mdm-jq-${JQ_VERSION}"
MARKER_FILE_DIR="/tmp/snyk-ide-mdm"

verify_jq_hash(){
  local jq_path="$1"
  local arch="$2"

  local expected_hash
  case "$arch" in
    arm64) expected_hash="$JQ_HASH_MACOS_ARM64" ;;
    amd64) expected_hash="$JQ_HASH_MACOS_AMD64" ;;
    *)
      log_error "Unknown architecture for hash verification: $arch"
      return 1
      ;;
  esac

  local actual_hash
  actual_hash="$(run_as_user /sbin/md5 -q "$jq_path" 2>/dev/null)"

  if [[ "$actual_hash" != "$expected_hash" ]]; then
    log_error "MD5 hash mismatch for jq binary. Expected: $expected_hash, Got: $actual_hash"
    return 1
  fi
 
  log_info "jq binary hash verified successfully with hash: $actual_hash"
  return 0
}

download_jq_if_needed(){
  if [[ -x "$JQ_BINARY_PATH" ]]; then
    log_info "jq binary already exists at $JQ_BINARY_PATH"
    return 0
  fi

  log_info "Downloading jq ${JQ_VERSION} binary..."

  local arch
  arch="$(get_system_architecture)" || return 1

  local download_url

  case "$arch" in
    arm64|amd64) download_url="https://github.com/jqlang/jq/releases/download/jq-${JQ_VERSION}/jq-macos-${arch}" ;;
    *)
      log_error "Unknown architecture: $arch"
      return 1
      ;;
  esac

  log_info "Downloading jq binary from: $download_url"

  if ! run_as_user /usr/bin/curl -L -f -s -o "$JQ_BINARY_PATH" "$download_url"; then
    log_error "Failed to download jq binary from $download_url"
    return 1
  fi

  if ! run_as_user /bin/chmod +x "$JQ_BINARY_PATH"; then
    log_error "Failed to make jq binary executable"
    return 1
  fi

  verify_jq_hash "$JQ_BINARY_PATH" "$arch" || {
    log_error "Hash verification failed, removing binary"
    run_as_user /bin/rm -f "$JQ_BINARY_PATH"
    return 1
  }

  log_info "Successfully downloaded and configured jq binary"
  return 0
}

update_json(){
  local file_path="$1"

  run_as_user /bin/mkdir -p "$(dirname "$file_path")"
  if ! run_as_user test -f "$file_path"; then
    log_warning "Config file missing; creating"
    run_as_user /bin/sh -lc "printf '{}' > \"${file_path}\""
  fi

  # Validate JSON. NB: jq doesn't support jsonc like vscode + derivatives do, so will fail on comments or trailing commas.
  local jq_error
  if ! jq_error=$(run_as_user "$JQ_BINARY_PATH" '.' "$file_path" 2>&1 >/dev/null); then
    log_error "Invalid JSON detected in $file_path: $jq_error"
    return 1
  fi

  # Read original content for comparison
  local original_content
  original_content=$(run_as_user "$JQ_BINARY_PATH" '.' "$file_path" 2>/dev/null)

  # See options: https://github.com/snyk/vscode-extension/blob/929244b7756f28dfb54e879ea616940a7c6d5cf3/package.json#L278-L282
  local jq_filter='."snyk.securityAtInception.autoConfigureSnykMcpServer" = true | ."snyk.securityAtInception.executionFrequency" = "On Code Generation"'

  local output
  if ! output=$(run_as_user "$JQ_BINARY_PATH" "$jq_filter" "$file_path"); then
    log_error "Failed to process settings in $file_path"
    return 1
  fi

  # Compare original and updated content
  if [[ "$original_content" == "$output" ]]; then
    log_info "No updates written in setting file $file_path (settings already configured)"
    return 0
  fi

  if ! run_as_user /bin/sh -c "printf '%s\n' \"\$0\" > \"$file_path\"" "$output"; then
    log_error "Failed to write updated settings to $file_path"
    return 1
  fi

  log_info "Successfully updated settings in $file_path"
  return 0
}

get_marker_file_path(){
  local ide="$1"
  echo "${MARKER_FILE_DIR}/${ide}"
}

is_already_updated(){
  local marker_file="$1"

  if run_as_user test -f "$marker_file"; then
    return 0  # Already updated
  fi

  return 1  # Not updated yet
}

create_marker_file(){
  local ide="$1"
  local marker_file
  marker_file="$(get_marker_file_path "$ide")"

  run_as_user /bin/mkdir -p "$(dirname "$marker_file")"

  # Create marker file with timestamp
  if ! run_as_user /bin/sh -c "printf '%s\n' \"$(date -u +%Y-%m-%dT%H:%M:%SZ)\" > \"$marker_file\""; then
    log_error "Failed to create marker file: $marker_file"
    return 1
  fi

  log_info "Created marker file: $marker_file"
  return 0
}

process_config_file(){
  local ide="$1"
  local config_file="$2"
  log_info "Processing configuration file for ${ide}"

  local marker_file; marker_file="$(get_marker_file_path "$ide")"
  if is_already_updated "$marker_file"; then
    log_info "Skipping update for ${ide} - already configured (marker file exists at $marker_file)"
    return 0
  fi

  if ! update_json "$config_file"; then
    return 1
  fi

  if ! create_marker_file "$ide"; then
    log_warning "Failed to create marker file for ${ide}, but config was updated"
    # Don't fail the whole operation, but log warning
  fi

  return 0
}

main(){
  if ! download_jq_if_needed; then
    log_error "Failed to download jq binary. Cannot proceed."
    exit 1
  fi

  local user_home; user_home="$(get_console_home)"

  # Array of IDE names and config file paths as pairs
  # Format: "IDE_NAME|CONFIG_FILE_PATH"
  local config_entries=(
    "windsurf|${user_home}/Library/Application Support/Windsurf/User/settings.json"
    "cursor|${user_home}/Library/Application Support/Cursor/User/settings.json"
  )

  local ok=0
  for entry in "${config_entries[@]}"; do
    IFS='|' read -r ide config_file <<< "$entry"
    if ! process_config_file "$ide" "$config_file"; then ok=1; fi
  done
  exit $ok
}

main "$@"
```

Here is an example of the Windsurf.pkg and both Snyk scripts ready to be deployed on devices using JAMF:

![Example of the Windsurf.pkg with Snyk scripts](<../../.gitbook/assets/unknown (9).png>)

In order for the IDE and the MCP to be properly configured:

* Check that the MCP server is present. In Windsurf, navigate to **Windsurf** > **Settings** > **Advanced settings** > **Cascade** > **MCP Servers** > **Manage MCPs**.

![Manage MCP servers screen in Windsurf](<../../.gitbook/assets/unknown (10).png>)

* Check that the user is prompted to trust Snyk. The code scan does not work if trust is not provided.
* If Secure at Inception is configured, ensure that the rules file `snyk_rules.md` includes the rules and that **Activation Mode** is set to **Always On**.

![Rules file and rules activation mode option in Windsurf](<../../.gitbook/assets/unknown (13).png>)

To learn more, visit [Troubleshooting for the Snyk MCP](troubleshooting.md).
