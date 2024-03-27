# Start scanning

You can use Snyk to scan your code manually and automatically using the [Snyk CLI](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-cli), the [Snyk Web UI](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-web-ui), the [Snyk API](start-scanning-using-the-cli-web-ui-or-api.md#scan-using-the-api), and by running [PR Checks](start-scanning-using-the-cli-web-ui-or-api.md#using-pr-checks).

{% hint style="info" %}
Scans (tests) may be limited on your account, depending on your[ Pricing Plan](../implement-snyk/enterprise-implementation-guide/trial-limitations.md). For more information, see [What counts as a test?](scanning-overview/what-counts-as-a-test.md)
{% endhint %}

<table><thead><tr><th width="220">Features</th><th width="126">Snyk Web UI</th><th width="111">Snyk CLI</th><th width="135">Snyk API</th><th>PR Checks</th></tr></thead><tbody><tr><td>Auto scanning</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>Manual scanning</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>Local scans</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>Incorporate into the CI/CD pipelines</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>Obtain results precisely reflecting the Project vulnerabilities and configurations</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr></tbody></table>

## Scan using the CLI

{% hint style="info" %}
See [Getting started with the CLI](../snyk-cli/getting-started-with-the-snyk-cli.md) for more details.&#x20;
{% endhint %}

Use the following Snyk [CLI commands](../snyk-cli/cli-commands-and-options-summary.md) for specific scanning methods:

<table><thead><tr><th width="190">Command</th><th width="236">Function</th><th>More details</th></tr></thead><tbody><tr><td><a href="../snyk-cli/commands/test.md">snyk test</a></td><td>Scan open-source code</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/">Use Snyk Open Source from the CLI</a></td></tr><tr><td><a href="../snyk-cli/commands/code.md">snyk code test</a></td><td>Scan application code</td><td> <a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/">Use Snyk Code from the CLI</a></td></tr><tr><td><a href="../snyk-cli/commands/container.md">snyk container test</a></td><td>Scan container images</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/">Use Snyk Container from the CLI</a></td></tr><tr><td><a href="../snyk-cli/commands/iac.md">snyk iac test</a></td><td>Scan infrastructure as code (IaC) files</td><td> <a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/">Snyk CLI for IaC</a></td></tr><tr><td><a href="../snyk-cli/commands/monitor.md">snyk monitor</a> and <a href="../snyk-cli/commands/container-monitor.md">snyk container monitor</a></td><td>Constantly monitor a Project for new vulnerabilities.</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md">Monitor your projects at regular intervals</a></td></tr></tbody></table>

## Scan using the Web UI

A scan runs when you import a Snyk Project (see [Import a Project](../getting-started/quickstart/import-a-project.md)) or click the **Retest now** button on a Project. Snyk then automatically runs periodic scans on that imported Project, to see if your code is affected by newly disclosed vulnerabilities.

See [Explore Snyk through the Web UI](../getting-started/snyk-web-ui.md).

The default **scanning frequency** and available frequencies vary depending on the type of Project: for more information, see [Usage settings](../snyk-admin/groups-and-organizations/usage-settings.md). You can also set the frequency in the Project **Settings** (see [View and edit project settings](../snyk-admin/snyk-projects/view-and-edit-project-settings.md)) or use the Snyk REST API; see the  [Updates project by project ID](https://apidocs.snyk.io/?version=2023-02-15#patch-/orgs/-org\_id-/projects/-project\_id-) endpoint.

## Scan using the API

The Snyk API v1 offers a set of endpoints to test your code. Scans are counted when calls are made to the **https://snyk.io/api/v1/test** endpoint.

See the [API documentation](https://snyk.docs.apiary.io/#reference/test) and the [API v1 Test docs](https://snyk.docs.apiary.io/#reference/test) for more information.

## Using PR Checks

Snyk can scan every new Pull Request (PR) submitted on your monitored repositories to help prevent new vulnerabilities from being added to your codebase.

See [Run PR Checks](run-pr-checks/) for details.
