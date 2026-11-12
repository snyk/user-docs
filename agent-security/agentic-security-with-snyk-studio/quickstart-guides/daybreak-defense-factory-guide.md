---
nav_context: classic
last_reviewed: 2026-11-11T00:00:00.000Z
last_reviewed_source: focused-edit
description: How to import Snyk findings into OpenAI Daybreak Defense Factory
hidden: true
---

# Connect Daybreak Defense Factory to Snyk

Import Snyk findings into OpenAI Daybreak Defense Factory in two steps: produce a Snyk report, then preview and publish it with the Codex Security CLI. A local preview runs offline before anything leaves your machine.

Snyk findings enter Defense Factory at the Discovery stage, where they join other vulnerability reports as candidate vulnerabilities. Results return to the workflows your teams already use, such as pull requests and tickets. To learn how Defense Factory works, visit the [OpenAI Defense Factory playbook](https://cdn.openai.com/defense-factory/downloads/defense-factory-playbook.pdf).

{% hint style="info" %}
The import supports Snyk Open Source findings through all four routes and Snyk Code findings through the monitored-Projects route only. Defense Factory does not import Snyk Container, Snyk Infrastructure as Code (IaC), or software bill of materials (SBOM) results. For details, visit [Supported findings by scan type](daybreak-defense-factory-guide.md#supported-findings-by-scan-type).
{% endhint %}

## Prerequisites

* A Snyk account and the Snyk CLI, authenticated with `snyk auth`
* The Codex Security CLI (`codex-security`), installed from the quick start in the [openai/codex-security repository](https://github.com/openai/codex-security). The CLI requires Node.js 22.13.0 or later in 22.x, or Node.js 24.x or 26.x, and Python 3.10 or later. Python 3.10 also requires `tomli`. Run `codex-security publish findings --help` to confirm that your installed version supports the `snyk` provider.
* Access to Defense Factory. OpenAI recommends that you request cyber model access through [Trusted Access for Cyber](https://chatgpt.com/cyber) and contact your OpenAI account team about architecture.
* For Route B: Node.js and Codex, with the Snyk plugin for Codex installed and authenticated as described in [Set up the Snyk plugin for Codex](daybreak-defense-factory-guide.md#set-up-the-snyk-plugin-for-codex)
* For Route D: a `SNYK_TOKEN` with read access for your region, plus your Organization ID and Project ID
* For Cloud imports: a repository connected in Codex Security Cloud. The Snyk Project must match the selected Cloud repository.

## Choose your route

Choose one of four routes to produce a Snyk report that the Codex Security CLI can import.

| Route                                                   | Use when                                                                      | Findings                                                                                                                                                        | Source key suffix                          |
| ------------------------------------------------------- | ----------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| Route A: native CLI JSON (recommended for dependencies) | You scan dependencies with the Snyk CLI.                                      | Snyk Open Source dependency vulnerabilities, with individual dependency paths and remediation details                                                           | `native-sca`                               |
| Route B: Snyk plugin export                             | You already use the Snyk plugin for Codex.                                    | Snyk Open Source vulnerabilities from a software composition analysis (SCA) scan only. The mapper can shorten descriptions and omit alternate dependency paths. | `plugin-sca`                               |
| Route C: SARIF                                          | Your workflow uses Static Analysis Results Interchange Format (SARIF) output. | Snyk Open Source vulnerabilities only. The SARIF file must be self-contained with direct URIs. Indexed or base-relative manifest locations are not supported.   | `sarif-sca`                                |
| Route D: monitored Snyk Projects                        | Your findings are already in Snyk, including Snyk Code findings.              | Open and resolved findings, including ignored findings, from Snyk Open Source or Snyk Code Projects. Requires a read-authorized token.                          | _`ORG_ID`_`/`_`PROJECT_ID`_ (the full key) |

A source key identifies a set of imported records. Changing the repository or source key creates a separate set, so use a separate source key for each file format. For repeat imports, reuse the exact repository and source key from your earlier local preview.

## Set up the Snyk plugin for Codex

Complete this section only for Route B.

### Install the plugin

```bash
codex plugin marketplace add snyk/studio-recipes
codex plugin add snyk@snyk-studio-recipes
```

Start a new Codex session after you install the plugin. The first session, and the first session after each Snyk CLI release, downloads the Snyk CLI (about 275MB), so the server takes longer to start. The plugin allows 120 seconds for startup.

### Authenticate

```bash
snyk auth
```

The command opens a browser OAuth flow, and the Snyk CLI stores and refreshes the token locally, outside the plugin and your prompts. For Organization, region, and token lifecycle details, visit [Authentication details](daybreak-defense-factory-guide.md#authentication-details).

### Use SNYK\_TOKEN in CI or with a service account

Codex does not pass arbitrary environment variables to plugin MCP servers, so the plugin cannot read an exported `SNYK_TOKEN`. Define the server in your Codex configuration:

{% code title="~/.codex/config.toml" %}
```toml
[mcp_servers.Snyk]
command = "npx"
args = ["-y", "snyk@latest", "mcp", "-t", "stdio"]
env_vars = ["SNYK_TOKEN"]
startup_timeout_sec = 120
```
{% endcode %}

A user-defined server with the same name replaces the plugin server.

## Produce the Snyk report

From the checkout root, set the destination repository and a source key that identifies the repository, branch, scan target, and format. Run all commands on this page in the same shell session so the variables persist.

```bash
snyk_repository=$(pwd -P)
snyk_source_key=github.com/example/project/main/root/native-sca
```

If you scanned a subdirectory, replace `root` with your scan target, such as `services/api`. Change the suffix to match your route, as listed in [Choose your route](daybreak-defense-factory-guide.md#choose-your-route).

{% tabs %}
{% tab title="Route A: native CLI JSON" %}
1.  From the directory that contains your manifest file, run the Snyk CLI. Save the report in a private directory outside the checkout. This example scans `package.json`:

    <pre class="language-bash" data-overflow="wrap"><code class="lang-bash">snyk test --file=package.json --json-file-output=PRIVATE_DIR/snyk-open-source.json
    </code></pre>
2.  Check the exit code immediately after the scan:

    ```bash
    echo $?
    ```

    Exit code 1 means Snyk found vulnerabilities, which is expected. An exit code higher than 1 means the scan failed. Fix the error and scan again before you import the report.
3.  Set the report path:

    ```bash
    snyk_report=PRIVATE_DIR/snyk-open-source.json
    ```

To scan several manifest files under one scan root, replace `--file=package.json` with `--all-projects --fail-fast`. With `--fail-fast`, an error stops the scan with exit code 2 and Snyk reports no vulnerabilities, so fix the error and scan again. Without `--fail-fast`, Snyk scans all manifests but does not report vulnerabilities for manifests it could not scan, so the report is incomplete. An error in a saved result also stops the import. Use separate reports and source keys for independent scan roots.
{% endtab %}

{% tab title="Route B: Snyk plugin export" %}
1. Complete [Set up the Snyk plugin for Codex](daybreak-defense-factory-guide.md#set-up-the-snyk-plugin-for-codex).
2.  From the Git root, ask Codex to use the `snyk-findings-export` skill to run an SCA-only scan and save the bundle outside the checkout. The skill runs SCA and Snyk Code scans by default, so name the scan type in your request. For example:

    <pre data-title="prompt" data-overflow="wrap"><code>Use the snyk-findings-export skill to run a Snyk Open Source (SCA) scan only. Save the bundle in PRIVATE_DIR, outside this checkout.
    </code></pre>
3.  Set the report path and the source key:

    ```bash
    snyk_report=PRIVATE_DIR/BUNDLE_FILE.json
    snyk_source_key=github.com/example/project/main/root/plugin-sca
    ```

The skill scans the current workspace. To import an earlier scan, use its saved report. The bundle records a commit only for a clean workspace with a known revision.
{% endtab %}

{% tab title="Route C: SARIF" %}
1.  Run the Route A command with `--sarif-file-output` in place of `--json-file-output`:

    <pre class="language-bash" data-overflow="wrap"><code class="lang-bash">snyk test --file=package.json --sarif-file-output=PRIVATE_DIR/snyk-open-source.sarif
    </code></pre>

    Check the exit code as in Route A. An exit code higher than 1 means the scan failed.
2.  Set the report path and the source key:

    ```bash
    snyk_report=PRIVATE_DIR/snyk-open-source.sarif
    snyk_source_key=github.com/example/project/main/root/sarif-sca
    ```

Use self-contained SARIF with direct URIs for manifest locations.
{% endtab %}

{% tab title="Route D: monitored Snyk Projects" %}
1. Set `SNYK_TOKEN` to a read-authorized token for your region.
2. Find the Organization ID, Project ID, and source repository URL for the Project.
3.  Navigate to a trusted directory outside the Project checkout, then collect the findings:

    <pre class="language-bash" data-overflow="wrap"><code class="lang-bash">snyk_report=$(mktemp ./snyk-issues.XXXXXX)
    codex-security import snyk PROJECT_ID --snyk-org ORG_ID --source-repository-url https://github.com/example/project --format json > "$snyk_report"
    snyk_source_key=ORG_ID/PROJECT_ID
    </code></pre>

    The command collects Snyk Open Source findings by default. For a Snyk Code Project, add `--snyk-product code`. For a regional deployment, set `--snyk-api-origin` to the API origin for your region: `https://api.us.snyk.io` for SNYK-US-02, `https://api.eu.snyk.io` for SNYK-EU-01, or `https://api.au.snyk.io` for SNYK-AU-01. The default is `https://api.snyk.io` for SNYK-US-01. The command rejects other origins.

Collection reads open and resolved findings across all pages, including ignored findings. For Snyk Code findings from an SCM integration, locations import as repository-relative paths, and the v2 report also imports the start line of each coordinate when available. For resolved findings, closed coordinates, and Projects monitored through the CLI or API, the original location stays in the saved finding evidence.
{% endtab %}
{% endtabs %}

## Preview, publish, and verify

Use the same `snyk_report`, `snyk_repository`, and `snyk_source_key` values for every command in this section.

{% stepper %}
{% step %}
**Preview locally**

Run a dry run. The local preview works offline, with no Cloud login or model call.

{% code overflow="wrap" %}
```bash
codex-security publish findings "$snyk_report" --to local --repository "$snyk_repository" --provider snyk --source-key "$snyk_source_key" --dry-run --format json
```
{% endcode %}

Check the expected advisories, package versions, repository, and exclusions, especially license issues. The `--format json` output lists every finding and exclusion. If the report has no Git URL, confirm the destination checkout yourself.
{% endstep %}

{% step %}
**Save locally**

Run the same command without `--dry-run`, and confirm when prompted. In automation, add `--yes` after you review the preview.

{% code overflow="wrap" %}
```bash
codex-security publish findings "$snyk_report" --to local --repository "$snyk_repository" --provider snyk --source-key "$snyk_source_key"
```
{% endcode %}

To list the saved findings, run:

```bash
codex-security findings imported "$snyk_repository" --format json
```
{% endstep %}

{% step %}
**Send to Cloud**

Keep the file, provider, source key, and options from your local preview. Change the destination to `--to cloud`, and pass the connected repository URL or Cloud ID to `--repository`. The Snyk Project must match the selected Cloud repository.

Preview the Cloud import:

{% code overflow="wrap" %}
```bash
codex-security publish findings "$snyk_report" --to cloud --provider snyk --source-key "$snyk_source_key" --repository https://github.com/example/project --dry-run --format json
```
{% endcode %}

Check the account, endpoint, repository, selected findings, and exclusions. Then run the command again without `--dry-run`, and confirm.
{% endstep %}

{% step %}
**Verify the import**

For a private inbox import (`privateImport` in the preview), check that `privateJob.status` is `completed`, that its counts account for the selected findings, and that the job reports zero errors and zero unprocessed rows.

For a repository publisher import (`environments` in `requests`), check each receipt, the verified and unacknowledged counts, failures, and exclusions.
{% endstep %}
{% endstepper %}

If Cloud accepts a request but the response is lost, run the same command again to resume the saved request. If `publish findings` stops and returns `savedSubmission`, keep that checkpoint and preview the original export again with `--dry-run --format json`. Use the same build, state directory, Cloud account, endpoint, and import settings. If the preview shows `resumed` as `true`, confirmation continues the saved submission. If `resumed` is `false`, review the new submission before you confirm. Finish any pending request before you send updated evidence for the same findings.

If the repository publisher requires a reported severity, the preview excludes findings that have an unknown severity in Snyk. To keep these findings, import them locally or to a private inbox, when available.

New local findings start as `not_assessed`. To check whether a finding affects your code, give its evidence and the matching checkout to the Codex Security triage workflow. Replaying the same evidence leaves records unchanged, and importing a different capture of a finding updates the same record.

## Troubleshooting

| Symptom                                                 | Fix                                                                                                                                                                                                                    |
| ------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Snyk Container, Snyk IaC, or SBOM results do not import | Defense Factory does not import these results. Use any route for Snyk Open Source findings, or Route D for Snyk Code findings.                                                                                         |
| HTTP 429 during collection                              | Wait for the interval in the `Retry-After` header, run the collection again, and preview the complete export before you import it.                                                                                     |
| Duplicate or missing records after a re-import          | Reuse the exact repository and source key from the earlier preview.                                                                                                                                                    |
| Findings missing from the preview                       | Check the exclusions in the `--format json` output, and review license issues.                                                                                                                                         |
| The Snyk server is slow to start (Route B)              | The first launch downloads the Snyk CLI (about 275MB). Wait up to 120 seconds.                                                                                                                                         |
| Authentication error (Route B)                          | Run `snyk auth` again. In CI, define the server with `SNYK_TOKEN` as described in [Use SNYK\_TOKEN in CI or with a service account](daybreak-defense-factory-guide.md#use-snyk_token-in-ci-or-with-a-service-account). |
| The export reports the secrets scan as unavailable      | Older Snyk CLI versions do not expose the secrets tool, so the export records the scan as unavailable. Update the Snyk CLI.                                                                                            |

For other Snyk Studio issues, visit [Troubleshooting](../troubleshooting.md).

## Reference

Use the following tables and lists to look up supported scan types, authentication behavior, data handling, and plugin details.

### Supported findings by scan type

| Finding type                                                     | Snyk product                        | Import support                                                                                                           |
| ---------------------------------------------------------------- | ----------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| Dependencies and packages (SCA)                                  | Snyk Open Source                    | Supported through all four routes                                                                                        |
| Source code findings, static application security testing (SAST) | Snyk Code                           | Supported through monitored Projects only (Route D with `--snyk-product code`). Not supported through the plugin bundle. |
| IaC analysis                                                     | Snyk IaC                            | Not supported. Available through the Snyk MCP server only.                                                               |
| Container images                                                 | Snyk Container                      | Not supported. Available through the Snyk MCP server only.                                                               |
| SBOM files                                                       | Snyk Open Source (`snyk sbom test`) | Not supported. Available through the Snyk MCP server only.                                                               |
| License issues                                                   | Snyk Open Source                    | Not supported. The preview lists them as exclusions.                                                                     |
| Secrets                                                          | Snyk Secrets                        | Not supported. Available through the Snyk MCP server only.                                                               |

The `query_issues` tool in the [Evo MCP server](../../evo-by-snyk/platform-surfaces/evo-mcp-server/) returns policy violations, with the violating asset and the location it was found in. It is not a route for importing scan findings.

### Authentication details

| Topic                   | Details                                                                                                                                                                                                      |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Organization scope      | Scans run under the Organization set in the Snyk CLI (`snyk config set org=`_`ORG_ID`_) or the Organization passed to a tool.                                                                                |
| Region                  | If your account is outside `SNYK-US-01`, set your environment before you run `snyk auth`, for example `snyk config environment SNYK-EU-01`.                                                                  |
| Expiry and refresh      | OAuth tokens refresh automatically. If a scan reports an authentication error, run `snyk auth` again.                                                                                                        |
| Rotation and revocation | Run `snyk logout` to clear local credentials. Revoke tokens or service accounts in the Snyk Web UI. A revoked credential causes scans to fail with an authentication error. Scans never return partial data. |

### Data handling

* Snyk Code scans send source files to Snyk for analysis under your Snyk agreement.
* Snyk Open Source, Snyk Container, Snyk IaC, and SBOM scans send dependency and configuration metadata to Snyk.
* Plugin scans do not create or update Snyk Projects. The export never uses the IaC report option and never runs `snyk monitor`.
* The JSON file stays on your machine until you send it. If the bundle includes Snyk Code results, it contains source snippets from Snyk Code dataflow evidence, so treat it like source code.
* The export never writes secret values or credentials in repository URLs to the bundle. It replaces matched values with `[REDACTED]`.
* The export does not edit, stage, or commit any file other than the bundle. The bundle stays untracked in Git.
* The export skill treats scanner output, such as titles, messages, and snippets, as evidence and never acts on it as instructions.
* The separate `snyk-fix` skill reports counts of fixed and prevented vulnerabilities to Snyk.

### Plugin skills

| Skill                            | What it does                                                                                   |
| -------------------------------- | ---------------------------------------------------------------------------------------------- |
| `snyk-fix`                       | Scans, fixes, validates, and optionally opens a pull request for one vulnerability or a batch. |
| `secure-dependency-health-check` | Checks a package for vulnerabilities and maintenance health before you add it.                 |
| `snyk-findings-export`           | Exports findings and evidence to JSON. Read-only.                                              |

The plugin has no lifecycle hooks. For automatic scanning while Codex writes code, use the [Snyk Studio installer](codex-cli-guide.md), which configures Codex hooks. If the installer already configured Codex, the `[mcp_servers.Snyk]` entry in your Codex configuration overrides the plugin server. The installer also copies `snyk-fix` and `secure-dependency-health-check` into `~/.agents/skills/snyk/`, so Codex sees two copies of each skill after you enable the plugin. Keep one copy: disable the plugin or remove the installer copies of the skills.

### Versions

* Snyk plugin for Codex: 0.1.0
* Snyk CLI: the latest release when the server starts. Every export records the exact version.
* Export format: `snyk-findings-export/v1`

## Related resources

* [OpenAI scanner-import guide for Snyk](https://github.com/openai/codex-security/blob/main/docs/scanner-imports/snyk.md)
* [Snyk plugin for Codex documentation in snyk/studio-recipes](https://github.com/snyk/studio-recipes/tree/main/plugins/codex)
