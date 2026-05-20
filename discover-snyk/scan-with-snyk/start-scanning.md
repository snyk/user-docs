# Start scanning

You can use Snyk to scan your code manually and automatically using the [Snyk CLI](start-scanning.md#scan-using-the-cli), the [Snyk web UI](start-scanning.md#scan-using-the-web-ui), the [Snyk API](start-scanning.md#scan-using-the-api), and by running [PR Checks](start-scanning.md#using-pr-checks).

{% hint style="info" %}
Scans (tests) may be limited on your account, depending on your [pricing plan](../implementation-and-setup/enterprise-implementation-guide/trial-limitations.md). For more information, see [What counts as a test?](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/ELvljsaLKPkSpffOkmsQ/what-counts-as-a-test)
{% endhint %}

<table><thead><tr><th width="220">Features</th><th width="126">Snyk Web UI</th><th width="111">Snyk CLI</th><th width="135">Snyk API</th><th>PR Checks</th></tr></thead><tbody><tr><td>Auto scanning</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>Manual scanning</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>Local scans</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>Incorporate into the CI/CD pipelines</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>Obtain results precisely reflecting the Project vulnerabilities and configurations</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr></tbody></table>

## Scan using the CLI

{% hint style="info" %}
See [Getting started with the CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/developer-tools/snyk-cli/getting-started-with-the-snyk-cli) for more details.
{% endhint %}

Use the following Snyk [CLI commands](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/developer-tools/snyk-cli/cli-commands-and-options-summary) for specific scanning methods:

<table><thead><tr><th width="190">Command</th><th width="236">Function</th><th>More details</th></tr></thead><tbody><tr><td><a href="../developer-tools/snyk-cli/commands/test.md">snyk test</a></td><td>Scan open-source code</td><td><a href="../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/">Use Snyk Open Source from the CLI</a></td></tr><tr><td><a href="../developer-tools/snyk-cli/commands/code.md">snyk code test</a></td><td>Scan application code</td><td><a href="../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/">Use Snyk Code from the CLI</a></td></tr><tr><td><a href="../developer-tools/snyk-cli/commands/container.md">snyk container test</a></td><td>Scan container images</td><td><a href="../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/">Use Snyk Container from the CLI</a></td></tr><tr><td><a href="../developer-tools/snyk-cli/commands/iac.md">snyk iac test</a></td><td>Scan infrastructure as code (IaC) files</td><td><a href="../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/">Snyk CLI for IaC</a></td></tr><tr><td><a href="../developer-tools/snyk-cli/commands/monitor.md">snyk monitor</a> and <a href="../developer-tools/snyk-cli/commands/container-monitor.md">snyk container monitor</a></td><td>Continually monitor a Project for new vulnerabilities.</td><td><a href="../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md">Monitor your projects at regular intervals</a></td></tr></tbody></table>

## Scan using the Web UI

A scan runs when you import a Snyk Project (see [Import a Project to scan and identify issues ](../discover-snyk/getting-started/#import-a-project-to-scan-and-identify-issues)or click **Retest now** on a Project. Snyk then automatically runs periodic scans on that imported Project, to see if your code is affected by newly disclosed vulnerabilities.

The default scanning frequency and available frequencies vary depending on the type of Project. For more information, see [Usage settings](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IgtgtomLQ2TUgSKOMSAm/snyk-platform-administration/groups-and-organizations/usage-settings).

You can also set the frequency in the Project **Settings** (see [View and edit Project settings](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/BJO0IZx7zB6bOkotxQP2/scan-with-snyk/snyk-projects/view-and-edit-project-settings)) or use the API Endpont [Updates project by project ID](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/projects#orgs-org_id-projects-project_id).

## Scan using the API

The Snyk API offers a set of endpoints to test your code. Scans are counted when calls are made to the test endpoint.

For more information, see the API [Test](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/test-v1) endpoint documentation.

## Using PR Checks

Snyk can scan every new Pull Request (PR) submitted on your monitored repositories to help prevent new vulnerabilities from being added to your codebase.

For more information, see [Pull Request Checks](pull-requests/pull-request-checks/).
