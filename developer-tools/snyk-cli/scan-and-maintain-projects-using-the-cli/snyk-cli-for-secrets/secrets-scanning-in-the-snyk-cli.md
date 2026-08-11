# Secrets scanning in the Snyk CLI

Use Snyk Secrets for the CLI to identify and manage sensitive information (API keys, passwords, tokens) in your source code. You can:

1. [Scan your current directory for hard coded secrets.](secrets-scanning-in-the-snyk-cli.md#run-a-secrets-scan)
2. [Ignore findings.](secrets-scanning-in-the-snyk-cli.md#ignore-findings)
3. [Scan and review ignored secrets.](secrets-scanning-in-the-snyk-cli.md#review-ignored-secrets)
4. [Exclude files and directories from a scan.](secrets-scanning-in-the-snyk-cli.md#exclude-files-and-directories-from-a-scan)
5. [Scan with a pre-commit hook.](secrets-scanning-in-the-snyk-cli.md#scan-with-a-pre-commit-hook)

## Prerequisites

To use this feature, you must:

1. Install the latest version of the Snyk CLI. For installation instructions, visit [Install the Snyk CLI](https://docs.snyk.io/developer-tools/snyk-cli/install-or-update-the-snyk-cli).
2. Authenticate your machine with the following command:

```bash
snyk auth
```

{% hint style="info" %}
If you receive a `SNYK-CLI-0016` error, contact your Snyk account manager to enable the Secrets feature for your Organization.
{% endhint %}

## Run a secrets scan

To scan your current directory for hard-coded secrets, run:

```bash
snyk secrets test
```

To map results to your Organization, include the `--org` option.

```bash
snyk secrets test --org=<your-org-id-or-name>
```

The output includes the following information:

* **Severity**: Snyk assigns a severity level to each finding based on the type and risk of the exposed secret.
* **Finding ID**: A unique identifier for the specific leak. You need this ID to ignore findings.
* **Path and line number**: The locations of the discovered secret in your codebase.

### Ignore findings

If a secret is a known placeholder, a revoked key, or a won't fix scenario, use the `ignore` command to suppress the finding.

To ignore a finding:

1. Copy the Finding ID from the `snyk secrets test` output.
2.  Run the `ignore` command:

    ```bash
    snyk ignore create
    ```
3. Follow the prompts:
   * **Finding ID:** Paste the ID you copied.
   * **Ignore Type**: Select **wont-fix**, **not-vulnerable**, or **temporary-ignore**.
   * **Expiration**: Set an optional expiry date (YYYY-MM-DD).
   * **Reason**: Provide context (for example, `Key has been rotated and is no longer active`).

### Review ignored secrets

Use the `--include-ignores` option to run a scan and view suppressed items. This helps you audit your codebase and ensures you do not hide critical leaks.

```bash
snyk secrets test --include-ignores
```

## Exclude files and directories from a scan

Exclude the paths that you do not want Snyk Secrets to scan, such as test fixtures, vendored dependencies, or sample credentials. Snyk applies the exclusions before the scan, so excluded files are not uploaded to Snyk and produce no findings.

{% hint style="info" %}
Excluding a path is not the same as ignoring a finding. Snyk does not scan excluded paths at all, so they produce no findings and do not appear in the Snyk Web UI as ignored issues. To suppress a specific finding but continue to scan the file, use ignores instead. Visit [Ignore findings](secrets-scanning-in-the-snyk-cli.md#ignore-findings).
{% endhint %}

You can exclude paths in the following ways:

* Commit a `.snyk` file to your repository. Snyk recommends this approach, because everyone who scans the repository shares the exclusions, and Snyk also applies them when the repository is scanned through an SCM integration. Visit [Exclude paths using the `.snyk` file](secrets-scanning-in-the-snyk-cli.md#exclude-paths-using-the-.snyk-file).
* Use the `--exclude` option for a single scan. Visit [Exclude paths for a single scan](secrets-scanning-in-the-snyk-cli.md#exclude-paths-for-a-single-scan).

Snyk Secrets also skips the paths listed in your `.gitignore` files, as well as file types that cannot contain readable secrets, such as binaries, archives, media files, fonts, and dependency lockfiles.

### Exclude paths using the `.snyk` file

The `.snyk` file is a YAML policy file that you commit to your repository. The `snyk secrets test` command applies the patterns from the following `exclude` sections of every `.snyk` file it finds in the scanned directory tree:

* `global`: Applies to Snyk Secrets and to the other Snyk products that support the `global` section.
* `secrets`: Applies to Snyk Secrets only.

The following example excludes the `vendor` directory from all the supported Snyk products, and excludes the PEM test fixtures and the `examples` directory from Snyk Secrets only:

```yaml
# Snyk (https://snyk.io) policy file
exclude:
  global:
    - vendor/**
  secrets:
    - "fixtures/**/*.pem"
    - examples/**
```

{% hint style="info" %}
* Snyk Secrets applies the `global` and `secrets` sections only. It does not apply the `code` and `iac-drift` sections, which apply to Snyk Code and Snyk IaC.
* Patterns follow the `.gitignore` pattern syntax and are relative to the directory that contains the `.snyk` file. A pattern such as `fixtures/` matches at any depth below that directory; start the pattern with a slash, for example `/fixtures/`, to match only at the level of the `.snyk` file. Absolute paths are not supported and are skipped.
* Wrap any pattern that begins with a special character, such as an asterisk (`*`), in double quotation marks.
* If Snyk cannot parse a section of the `.snyk` file, it skips that section and applies the remaining sections. If it cannot parse the file at all, the scan continues without the exclusions from that file.

For the full exclusion pattern syntax and formatting rules, see [Exclusion syntax of the `.snyk` file](https://app.gitbook.com/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/import-project-repository/exclude-directories-and-files-from-project-import#exclusion-syntax-of-the-.snyk-file).
{% endhint %}

### Exclude paths for a single scan

Use the `--exclude` option to exclude directory names and file names for one scan, without committing anything to the repository. Provide a comma-separated list of names, without paths:

```bash
snyk secrets test --exclude=fixtures,secrets.example.yaml
```

This excludes every directory or file with a matching name, at any depth, for example `./fixtures` and `./src/fixtures`.

## Scan with a pre-commit hook

Run Snyk Secrets as a pre-commit hook to catch secrets before they reach a commit. The hook keeps secret detection in the developer workflow and prevents new secrets from entering the repository.

### Authenticate the hook

The `secrets test` command must run as an authenticated CLI command. Each developer authenticates their own CLI locally.

Authenticate in one of two ways:

* Run `snyk auth` to use your local Snyk credentials.
* Export `SNYK_TOKEN` in your shell environment.

{% hint style="warning" %}
**Do not embed a token in the hook command**

Do not embed a token in the hook command, for example `SNYK_TOKEN=... snyk secrets test`. The repository stores and shares hook configuration, so anyone with repository access can copy an embedded token. The token can leak into logs and get flagged by future scans. Keep credentials out of the repository and have the CLI read them from local auth state or `SNYK_TOKEN`.
{% endhint %}

### Pass staged file paths

Pass the staged files to Snyk instead of scanning the whole project. Scanning the full repository on every commit is slower, can block unrelated commits on pre-existing findings, and makes it harder to tell whether a failure came from your change or from older files.

{% hint style="info" %}
The `secrets test` command accepts only one file per invocation. For this reason, the following hooks loop over the staged files and run the command once per path, tracking the worst exit code, instead of passing all files at once.
{% endhint %}

### Set up with the pre-commit framework

Add a local hook to `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: local
    hooks:
      - id: snyk-secrets
        name: snyk secrets test
        entry: sh -c 'status=0; for path do snyk secrets test "$path"; rc=$?; if [ "$rc" -ne 0 ]; then status="$rc"; fi; done; exit "$status"' --
        language: system
        pass_filenames: true
        stages: [pre-commit]
```

With `pass_filenames: true`, pre-commit appends the staged files to the command. The shell wrapper preserves the blocking behavior by exiting non-zero when any scan fails.

### Use other hook managers

The same pattern works with other hook managers: collect the staged files and pass them to `secrets test`.

<details>

<summary>Raw Git hook</summary>

Add the following to `.git/hooks/pre-commit`:

```sh
#!/usr/bin/env sh
files="$(mktemp)"
status=0
git diff --cached --name-only --diff-filter=ACMR > "$files"
while IFS= read -r path; do
  snyk secrets test "$path"
  rc=$?
  if [ "$rc" -ne 0 ]; then
    status="$rc"
  fi
done < "$files"
rm -f "$files"
exit "$status"
```

</details>

<details>

<summary>Husky</summary>

Add the following to `.husky/pre-commit`:

```sh
#!/usr/bin/env sh
. "$(dirname "$0")/_/husky.sh"
files="$(mktemp)"
status=0
git diff --cached --name-only --diff-filter=ACMR > "$files"
while IFS= read -r path; do
  snyk secrets test "$path"
  rc=$?
  if [ "$rc" -ne 0 ]; then
    status="$rc"
  fi
done < "$files"
rm -f "$files"
exit "$status"
```

</details>

<details>

<summary>Lefthook</summary>

Add the following to `lefthook.yml`:

```yaml
pre-commit:
  commands:
    snyk-secrets:
      glob: "*"
      run: sh -c 'status=0; for path do snyk secrets test "$path"; rc=$?; if [ "$rc" -ne 0 ]; then status="$rc"; fi; done; exit "$status"' -- {staged_files}
```

</details>

{% hint style="info" %}
Hook managers vary in how they handle placeholder behavior, such as spaces in filenames and empty file lists. Check the documentation for your specific hook manager.
{% endhint %}
