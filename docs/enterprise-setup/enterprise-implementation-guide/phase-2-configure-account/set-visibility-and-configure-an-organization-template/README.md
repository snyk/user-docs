# Set visibility and configure an Organization template

Whether you want to create a single Organization or build a template to create multiple Organizations, here are some initial settings you should configure.&#x20;

{% hint style="info" %}
For importing Projects, see the [Import Projects](../../phase-3-gain-visibility/import-projects.md) and [Rollout](../../phase-5-initial-rollout-to-team/) discussions
{% endhint %}

### Create your Organization Structure using a template

When you create a new Organization, you can select an existing Organization to use as the model for settings and integrations. To streamline creating your Organizations, we recommend configuring a template Organization, before creating your full Organization structure.

This template-based approach will save you considerable time, avoiding manually configuring each integration for each Organization.

{% hint style="info" %}
There is no specific ‘template’ functionality in Snyk. Our recommended process is to create an Organization called **Template**, then fully configure this Organization. \
When you then create more Organizations, you can use the option to clone settings from an existing Organization, using **Template** as the basis.
{% endhint %}

#### Templating using API tools

Templating functionality is also available if you are using creating your Organizations using the API, whether you are using the [snyk-api-import](../../../../snyk-api-info/other-tools/tool-snyk-api-import/) tool to mirror Organization from an existing source (such as GitHub Organizations), or using the [API endpoints](https://snyk.docs.apiary.io/#reference/organizations/create-organization/create-a-new-organization) by providing a sourceOrgId.



Configure template Organization settings

In your template Organization, configure a range of settings which you can then choose to copy when creating your full Organization structure:

* Set up all relevant integrations (for example, GitHub Enterprise, Docker Hub)
  * If you have on-premise source code management tools, you must configure and run the [Snyk Broker](../../../snyk-broker/) to enable the integration.
* Integration settings (for example, Configuring whether you want Snyk to run tests on PRs)
  * The default settings for a new Git repository integration include Snyk running tests on newly raised PRs and the option to automatically raise PRs when new vulnerabilities are found. We recommend disabling these initially and turning them on when you are ready to introduce these features in the [Prevention Stage](../../phase-6-rolling-out-the-prevention-stage/).
  * The following [Integrations](configure-integrations.md) section discusses integrations you may want to add to your "template" before copying them.
* Product settings (for example, enabling Snyk Code)
  * By default, Snyk Code is disabled for new Organizations. If you want to enable SAST scanning of repositories imported through a Git repository integration, please ensure you use the toggle to enable this product before importing your Projects.

{% hint style="success" %}
Using Snyk Code and forgot to enable it before importing? \
\
You will need to enable it, and reimport your projects from Git. See [Enabling Snyk Code](enable-snyk-code.md).
{% endhint %}

{% hint style="info" %}
Notification settings (email notifications)

* Snyk suggests that you initially disable all group/organization email notifications so that users do not receive a large number of notifications while projects are imported.
* You would disable Notifications at the ‘Group’ level (for new organizations), and at the ‘Organization’ level for all existing Organizations.&#x20;
{% endhint %}

The following table illustrates what is copied from the "template Organization" to the new organizations you create via web interface or API\


<table><thead><tr><th width="466">All integrations and their settings will be copied across</th><th>The following will not be copied</th></tr></thead><tbody><tr><td>Source control integrations</td><td>Snyk Service accounts</td></tr><tr><td>Container registry integrations</td><td>Members</td></tr><tr><td>Container orchestrators integrations (Kubernetes)</td><td>Projects</td></tr><tr><td>PaaS and Serverless integrations</td><td>Notification preferences</td></tr><tr><td>Notificiation integrations (Slack/Jira)</td><td></td></tr><tr><td>Policies</td><td></td></tr><tr><td>Ignore settings</td><td></td></tr><tr><td>Language settings</td><td></td></tr><tr><td>Infrastructure as code (IaC) settings</td><td></td></tr><tr><td>Snyk Code settings</td><td></td></tr></tbody></table>

###

### Duplicate template organization

After the Organization is created and configured, use it as a template when creating new Organizations, to build out the remaining structure.
