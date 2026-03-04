# Gain visibility by importing repositories

Gaining visibility over your Organization security begins with importing Projects. This process allows Snyk to monitor your code, dependencies, containers, and infrastructure.

There are several ways you can import Projects, depending on your tech stack and package managers:

* **SCM integration:** Recommended for automatic scanning and developer workflows.
* **Snyk CLI:** Recommended for granular control in CI/CD pipelines.
* **Snyk API:** Recommended for large-scale, programmatic automation.

## Import Projects using an SCM integration

Use this method to connect repositories for automatic scanning. This is the preferred way for teams who prioritize ease of use.

{% hint style="success" %}
**Key decision**: Determine the scale of your import. Use the web UI for under 100 repositories; use the `snyk-api-import` tool for hundreds or thousands of repositories to mirror your source control structure.
{% endhint %}

To do this:

1. In the Snyk web UI, navigate to **Settings** > **Integrations**.
2. Connect to your SCM code repositories using the specific tile.
3. Configure the integration settings:
   * Disable automatic fixes and PR/Merge checks during initial onboarding.
   * Enable these features after reaching a steady state.
4. Add Projects from the Projects listing in the web UI.
5. Monitor results directly in your SCM repositories.

{% hint style="info" %}
Snyk recommends that any monorepos are imported into separate individual Organizations, rather than being split or imported between multiple Orgs.
{% endhint %}

## Import Projects using Snyk CLI

The Snyk CLI provides granular scanning and is typically implemented in build scripts before deployment.

{% hint style="success" %}
**Key decision**: Decide whether to use `test` to break builds based on severity thresholds or `monitor` to report vulnerabilities passively.
{% endhint %}

1. Install the CLI as part of your build script.
2. Navigate to the Project directory.
3. Run the scan command for your Project type:
   * Code: `snyk code test --org=org-id`
   * Open Source: `snyk test --all-projects --org=org-id`
4. Review results locally or in the Snyk web UI.

## Import Projects using Snyk API

Use the API to trigger scans and handle results programmatically across a large portfolio.

{% hint style="success" %}
**Key decision**: Identify which pipelines require real-time issue identification at scale.
{% endhint %}

1. In the Snyk web UI, navigate to **Settings** > **Service Accounts** and generate an API token.
2. Call the Snyk API in your pipelines.
3. Handle the results programmatically to trigger downstream actions.

## Add Projects tags and attributes

After importing Projects, use Project attributes and tags to categorize your data. This metadata allows you to filter and report on specific subsets of your Organization.&#x20;

For example, apply the `frontend` attribute and a `Team:Unicorn` tag to your Projects. You can then generate a report specifically for critical frontend vulnerabilities in production owned by Team Unicorn.&#x20;

Attributes and tags also allow you to:

* Group related Projects for easier management.&#x20;
* Customize how you view and access Project data.
