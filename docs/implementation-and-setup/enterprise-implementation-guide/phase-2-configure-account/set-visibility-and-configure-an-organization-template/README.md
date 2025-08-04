# Set visibility and configure an Organization template

Whether you want to create a single Organization or build a template to create multiple Organizations, there are some initial settings you should configure.&#x20;

{% hint style="info" %}
For importing Projects, see the [Import Projects](../../phase-3-gain-visibility/import-projects.md) and [Rollout](../../phase-5-initial-rollout-to-team/) discussions
{% endhint %}

## Create your Organization structure using a template

When you create a new Organization, you can select an existing Organization to use as the model for settings and integrations. To streamline creating your Organizations, Snyk recommends configuring a template Organization before creating your full Organization structure.

This template-based approach will save you considerable time because you will aovid manually configuring each integration for each Organization.

There is no specific template functionality in Snyk. The recommended process is to create an Organization called **Template and** then fully configure this Organization. Afterwards, when you create more Organizations, you can use the option to clone settings from an existing Organization with **Template** as the existing Organization.

{% hint style="info" %}
**Creating a template using the API**\
Templating functionality is also available if you are creating your Organizations using the API, whether you are using the [snyk-api-import](../../../../scan-with-snyk/snyk-tools/tool-snyk-api-import/) tool to mirror an Organization from an existing source, such as GitHub Organizations, or using the endpoint [Create a new organization](../../../../snyk-api/reference/organizations-v1.md#org) and providing a `sourceOrgId`.
{% endhint %}

## Configure template Organization settings

In your template Organization, configure a range of settings that you can choose to copy when creating your full Organization structure:

* All relevant integrations, for example, GitHub Enterprise, Docker Hub.\
  Note: If you have on-premise source code management tools, you must configure and run [Snyk Broker](../../../enterprise-setup/snyk-broker/) to enable the integration.
* Integration settings, for example, configuring whether you want Snyk to run tests on PRs.
  * The default settings for a new Git repository integration include Snyk running tests on newly raised PRs and the option to automatically raise PRs when new vulnerabilities are found. Snyk recommends disabling these settings initially and turning them on when you are ready to introduce these features in the [Prevention Stage](../../phase-6-rolling-out-the-prevention-stage/).
  * The following [Integrations](configure-integrations.md) section discusses integrations you may want to add to your templates before copying them.
* Product settings, for example, enabling Snyk Code.
  * By default, Snyk Code is disabled for new Organizations.
  * If you want to enable SAST scanning of repositories imported through a Git repository integration, ensure you use the toggle to enable Snyk Code before importing your Projects.
  * If you use Snyk Code and **forget to enable it before importing your Project, you must enable Snyk Code and reimport your code** from Git. See [Enabling Snyk Code](enable-snyk-code.md).
* Notification settings (email notifications)
  * Snyk suggests that you initially disable all Group and Organization email notifications so users do not receive a large number of notifications while Projects are imported.
  * To do this, disable Notifications at the Group level for new organizations, and at the Organization level for all existing Organizations.

The following table shows what is copied from the template Organization to the new Organizations you create using the web interface or API\


<table><thead><tr><th width="466">All integrations and their settings WILL be copied</th><th>The following WILL NOT be copied</th></tr></thead><tbody><tr><td>Source control integrations</td><td>Snyk Service accounts</td></tr><tr><td>Container registry integrations</td><td>Members</td></tr><tr><td>Container orchestrators integrations (Kubernetes)</td><td>Projects</td></tr><tr><td>PaaS and Serverless integrations</td><td>Notification preferences</td></tr><tr><td>Notificiation integrations (Slack and Jira)</td><td></td></tr><tr><td>Policies</td><td></td></tr><tr><td>Ignore settings</td><td></td></tr><tr><td>Language settings</td><td></td></tr><tr><td>Infrastructure as code (IaC) settings</td><td></td></tr><tr><td>Snyk Code settings</td><td></td></tr></tbody></table>

## Duplicate the template Organization

After the Organization is created and configured, use it as a template when creating new Organizations, to build out the remaining structure.
