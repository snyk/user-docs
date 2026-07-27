---
nav_context: classic
---

# Unified IDE Configuration Dialog

You can use only one IDE configuration dialog to configure all your IDEs.

<figure><img src="../../.gitbook/assets/Screenshot 2026-07-08 at 10.43.36.png" alt=""><figcaption></figcaption></figure>

## Tabs

The dialog has these tabs:

* **Setup**: authentication, CLI, and trust settings that apply across the IDE.
* **Project defaults**: the scan, filter, and advanced settings applied to every open project.
* A tab for each open project, where you can override the project defaults for that project.

## Setup

### Authentication

| Setting                          | Description                                                                                                                                                              |
| -------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Authentication method            | The method used to authenticate with Snyk: **OAuth2 (recommended)**, **Personal access token**, or **API token (legacy)**. Snyk recommends OAuth2 for stronger security. |
| Token                            | The token used to authenticate with the Snyk API. This field only appears when you select the **Personal access token** or **API token (legacy)** method.                |
| Endpoint                         | The endpoint used to communicate with the Snyk API. When you use SSO with OAuth2, Snyk populates this automatically.                                                     |
| Insecure (skip SSL verification) | Ignore SSL certificate errors. Use this only when connecting to self-signed or internal certificate authority servers.                                                   |
| Authenticate                     | Log in to Snyk. This button only appears when you select the **OAuth2** method.                                                                                          |
| Get Token                        | Open the Snyk page to generate a token. This link only appears when you select the **Personal access token** or **API token (legacy)** method.                           |
| Log out                          | Delete the stored authentication credentials. This button only appears when a credential is stored.                                                                      |

### CLI configuration

| Setting                       | Description                                                                                                                                                                     |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CLI path                      | The file path to the Snyk CLI executable. The plugin uses this path to run the CLI and to store downloads.                                                                      |
| Manage binaries automatically | Download and manage the Snyk CLI binary automatically, saving it to the location in **CLI path**.                                                                               |
| CLI Release                   | The release channel for automatic CLI downloads: **Stable**, **Release Candidate**, or **Preview**. Select **Specify version…** to pin a specific version, such as `v1.1292.0`. |
| CLI download base URL         | The base URL for CLI downloads. The default is `https://downloads.snyk.io`. For FIPS-enabled CLIs on Windows or Linux, use `https://downloads.snyk.io/fips`.                    |
| Path _(Eclipse only)_         | Eclipse has a **Path** field to prepend additional directories to `PATH` before the CLI runs. Restart the IDE to apply a change.                                                |

### Trust settings

| Setting              | Description                                                                 |
| -------------------- | --------------------------------------------------------------------------- |
| Trusted folder paths | The folders trusted for scanning. Each path includes all of its subfolders. |
| Add folder           | Add a trusted folder path.                                                  |

## Project defaults

The **Project defaults** tab sets the values applied to every project unless a project overrides them. Each open project has its own tab with the same settings.

### Scan configuration

| Setting                     | Description                                                                                                                                                                           |
| --------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Snyk Open Source            | Find and fix open source dependency issues.                                                                                                                                           |
| Snyk Code                   | Find and fix security issues in your application code. Snyk Code must be enabled for your Organization.                                                                               |
| Snyk Infrastructure as Code | Find and fix your IaC misconfigurations.                                                                                                                                              |
| Snyk Secrets                | Detect and prevent secrets from being exposed in your code. This scanner appears when the feature is available for your Organization.                                                 |
| Scanning mode               | Select **Auto** to run scans in the background at startup and on save, or **Manual** to run scans only when you run the Scan command.                                                 |
| Organization                | The Organization ID or name for Snyk to run scans against. Snyk uses the project-specific setting when configured, then this setting, then your web account's preferred Organization. |

### Filtering and views

| Setting                              | Description                                                                                                                                                                                                                                                                                                                 |
| ------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Filter by severity                   | Select which severity levels to display: **Critical**, **High**, **Medium**, and **Low**.                                                                                                                                                                                                                                   |
| Issue View Options                   | Select **Show open issues**, **Show ignored issues**, or both. These filters depend on Code Consistent Ignores for the Organization.                                                                                                                                                                                        |
| Filter by risk score _(Closed Beta)_ | The minimum risk score for an issue to display. The valid range is 0 to 1000, where 0 shows all issues.                                                                                                                                                                                                                     |
| Delta findings                       | Select **All issues** to show every issue, or **Net new issues** to show only issues newly introduced by your code changes. See the [delta findings feature](../../snyk-ide-plugins-and-extensions/visual-studio-code-extension/view-analysis-results-from-visual-studio-code-extension/#net-new-issues-versus-all-issues). |

### Advanced

Expand **Advanced** to set defaults for every project. These values combine with any per-project values rather than replacing them.

| Setting                          | Description                                                                                                                                                                          |
| -------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Additional CLI parameters        | Space-separated Snyk CLI parameters for all project scans. For example, `--severity-threshold=high --debug`. CLI parameters from this level and the project level are used together. |
| Additional environment variables | Environment variables passed to the Snyk CLI, in the format `VAR1=value1;VAR2=value2`. A per-project value overrides a project default when both set the same key.                   |

## Project settings

Each open project has its own tab with the same **Scan configuration**, **Filters and views**, and **Advanced** settings described previously. Values you set in a project tab override the project's defaults.

The project tab also includes:

| Setting                  | Description                                                                                                                                                   |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Auto select organization | Snyk selects the Organization for the project automatically, using context found in the repository. Clear this option to set the Organization yourself.       |
| Organization             | The Organization ID or name for this project. When **Auto select organization** is enabled, this field is read-only and shows the Organization Snyk selected. |

### Resetting settings

On a project tab, **Reset overrides** clears all overrides for that project and falls back to the project defaults.
