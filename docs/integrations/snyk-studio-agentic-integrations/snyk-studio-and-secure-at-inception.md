# Snyk Studio and Secure at Inception

The [Snyk quickstart guides for MCP](https://docs.snyk.io/integrations/developer-guardrails-for-agentic-workflows/quickstart-guides-for-mcp) enable you to install and configure Snyk Studio in your preferred coding assistant, such as Cursor, in minutes. To scale this process at the company level, you can roll out the [Secure at Inception ruleset](https://docs.snyk.io/integrations/developer-guardrails-for-agentic-workflows/quickstart-guides-for-mcp/cursor-guide#secure-at-inception-rules) broadly with as little friction as possible.

Secure at Inception is a configurable option that prompts the LLM to scan code for security issues during generation in the IDE. Developers can manually run Snyk security scans to add custom rules to ensure secure code at inception, or can enable the Secure at Inception option at the company level.

Below is a request using Windsurf and Claude to generate a login page.

![Request to generate a login page using Claude and Windsurf](<../../.gitbook/assets/unknown (4).png>)

Claude then uses Snyk’s MCP tools to validate the security of the code it just wrote due to deployed environmental rules.

![Validation steps using Snyk MCP tools](<../../.gitbook/assets/unknown (5).png>)

In this case, a vulnerability was found:

![Example of a vulnerabilty found](<../../.gitbook/assets/unknown (6).png>)

A fix is suggested:

![Example of a suggested fix](<../../.gitbook/assets/unknown (7).png>)

## Secure at Inception configuration

Using the VS Code Snyk IDE extension, the Secure at Inception option introduced settings that automatically set up Snyk Studio.

<figure><img src="../../.gitbook/assets/Security at inception setting.jpg" alt=""><figcaption><p>Secure at inception settings</p></figcaption></figure>

The rules of Secure at Inception are as follows:

* Always run the Snyk Code scanning tool for new first-party code generated.
* Always run the Snyk SCA scanning tool for new dependencies or dependency updates.
* Always run the Snyk IaC scanning tool for new Infrastructure as Code updates.
* If security issues arise from new or modified code or dependencies, attempt to fix them using the results context from Snyk.
* Rescan the code after fixing issues to verify they are resolved and no new issues are introduced.
* Repeat this process until no issues are found.

Snyk offers reference scripts for deploying the IDE extension on developer machines with default settings. You can customize these scripts as needed.

{% hint style="info" %}
Snyk recommends leveraging the Snyk IDE extension in order to install and configure Snyk Studio. However, it is not always mandatory. Contact the Snyk account team for other deployment and management options.
{% endhint %}

It is possible that these scripts need your input and customization.

## Prerequisites

Ensure that:

* You have an active Snyk account
* Access to an Endpoint Management tool, such as JAMF
* Snyk testing was performed on Mac with VSCode, Cursor and Windsurf using JAMF as the MDM. Further testing will be conducted on other operating systems (such as Windows), additional MDMs (such as Intune), and additional IDEs.

## Deploy the Secure at Inception option

For broad deployment, use an MDM or Endpoint Management tools like JAMF to target MCP configuration to developer devices.

For JAMF, targeting for Group and configurations must be separated as follows:

* For machines with Windsurf installed, target Snyk’s Windsurf MCP scripts to [deploy](https://learn.jamf.com/en-US/bundle/jamf-pro-documentation-current/page/Scripts.html).
* For machines with Cursor installed, target Snyk’s Cursor MCP scripts to deploy.
* Where neither is installed, deploy the desired [.pkg](https://learn.jamf.com/en-US/bundle/jamf-pro-documentation-current/page/Package_Deployment.html) and configure matching scripts to be run after installation.

You can configure the MDM to re-apply the desired Secure at Inception configuration across all endpoints on a regular basis.

## Sample scripts

### Install the Snyk extension and configure the Snyk MCP server

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

### Enable Secure at Inception by applying the rules in the Project

```sh
#!/bin/bash
set -u -o pipefail

log_info(){ echo "[INFO] $1"; }
log_error(){ echo "[ERROR] $1" >&2; }

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

set_key(){
  local plistbuddy="/usr/libexec/PlistBuddy"
  local file_path="$1"
  local full_key_path="$2"

  run_as_user "$plistbuddy" -c "Delete :$full_key_path" "$file_path" 2>/dev/null || true
  if ! run_as_user "$plistbuddy" -c "Add :$full_key_path bool true" "$file_path"; then
    log_error "Failed to set $full_key_path in $file_path"
    return 1
  fi

  log_info "Successfully set $full_key_path = true in $file_path"
  return 0
}

update_json(){
  local file_path="$1"

  local parent_key="snyk.securityAtInception"
  local all_keys=("autoConfigureMcpServer" "publishSecurityAtInceptionRules")
  local cursor_key="persistRulesInProjects"

  if [[ "$file_path" == *"Cursor"* ]]; then
    all_keys+=("$cursor_key")
  fi

  # Ensure parent dir + file
  run_as_user /bin/mkdir -p "$(dirname "$file_path")"
  if ! run_as_user test -f "$file_path"; then
    log_info "Config file missing; creating"
    run_as_user /bin/sh -lc "printf '{}' > \"${file_path}\""
  fi

  # Convert whatever is there to XML plist
  if ! run_as_user /usr/bin/plutil -convert xml1 "$file_path" 2>/dev/null; then
    # If conversion fails (corrupt file), reset to empty dict and convert again
    run_as_user /bin/sh -lc "printf '{}' > \"${file_path}\""
    run_as_user /usr/bin/plutil -convert xml1 "$file_path"
  fi

  local ok=0
  for key in "${all_keys[@]}"; do
    if ! set_key "$file_path" "$parent_key:$key"; then ok=1; fi
  done

  # Convert back to JSON
  run_as_user /usr/bin/plutil -convert json -r "$file_path" || {
    log_error "Failed to convert $file_path back to JSON"
    return 1
  }

  return $ok
}

is_valid_json_target(){
  local p="$1"
  [[ -n "${p}" ]] || return 1
  [[ "${p}" != "/" ]] || return 1
  [[ "${p}" == *.json ]] || return 1
  if run_as_user test -d "${p}"; then return 1; fi
  return 0
}

process_config_file(){
  local config_file="$1"
  log_info "Processing configuration file: $config_file"
  if ! is_valid_json_target "${config_file}"; then
    log_error "Invalid target path: ${config_file}. Must be a .json file, not '/' or a directory."
    return 1
  fi
  update_json "$config_file"
}

main(){
  local user_home; user_home="$(get_console_home)"

  local config_files=(
    "${user_home}/Library/Application Support/Windsurf/User/settings.json"
    "${user_home}/Library/Application Support/Cursor/User/settings.json"
    )

  local ok=0
  for cf in "${config_files[@]}"; do
    if ! process_config_file "$cf"; then ok=1; fi
  done
  exit $ok
}

main "$@"
```

Below is an example of the Windsurf.pkg and both Snyk scripts ready to be deployed on devices using JAMF:

![Example of the Windsurf.pkg with Snyk scripts](<../../.gitbook/assets/unknown (9).png>)

In order for the IDE and the MCP to be properly configured:

* Ensure that The MCP server is present. In Windsurf, navigate to **Windsurf** > **Settings** > **Advanced settings** > **Cascade** > **MCP Servers** > **Manage MCPs**.

![Manage MCP servers screen in Windsurf](<../../.gitbook/assets/unknown (10).png>)

* Ensure that the user is prompted to trust Snyk. The code scan does not work if trust is not provided.
* If Secure at Inception is configured, ensure that the rules file snyk\_rules.md includes the rules and that **Activation Mode** is set to **Always On**.

![Rules file and rules activation mode option in Windsurf](<../../.gitbook/assets/unknown (13).png>)

To learn more, see [Troubleshooting for the Snyk MCP](https://docs.snyk.io/integrations/developer-guardrails-for-agentic-workflows/troubleshooting-for-the-snyk-mcp-server).
