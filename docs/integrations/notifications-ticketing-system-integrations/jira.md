# Jira

A new integration, Snyk Security for Jira Software, is coming soon. If you are interested in early access, you can [request it](https://earlyaccessprogram.atlassian.net/servicedesk/customer/). This page covers the integration available in Snyk.

{% hint style="info" %}
**Feature availability**\
\- This feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.\
\- For the availability of this feature for Snyk Infrastructure as Code, see [Jira Integration for IaC](../../scan-cloud-deployment/snyk-infrastructure-as-code/jira-integration.md).
{% endhint %}

## **Set up your Jira integration**

Snyk Jira integration allows you to manually raise Jira issues in the Snyk UI for vulnerabilities or license issues. The Jira integration also includes APIs ([see Snyk API docs](https://snyk.docs.apiary.io/#reference/projects/project-jira-issues)).

{% hint style="info" %}
If your Jira instance is private, you will need to use it via [the Snyk Broker deployment method](../snyk-broker/snyk-broker-set-up-examples/setup-broker-with-jira.md).
{% endhint %}

## **Prerequisites**

* Snyk supports Jira version 5 and later versions.
* The following [Jira permissions](https://confluence.atlassian.com/adminjiraserver073/managing-project-permissions-861253293.html) are required: "**Browse Projects**" and "**Create Issues**".

## **How to set up your Jira integration**

The Jira account credentials are configured in the Snyk Web UI - **Organization Settings > Integrations** page.

It is best practice to set up a new user in Jira for this integration, instead of using the credentials of an existing account.

Cloud-hosted Jira implementations require a username and API token authentication. Jira API tokens are generated in [Atlassian API tokens](https://id.atlassian.com/manage/api-tokens). Self-hosted implementations are able to authenticate with a username and password.

## **Create a Jira issue**

Once you set up the Jira integration connection, open one of your Snyk Projects in the Snyk Web UI. A new button, **Create an issue**, now appears at the bottom of each vulnerability and license issue card. This button allows you to create a Jira issue.

![](<../../.gitbook/assets/Jira - new button.png>)

When you click the **Create an issue** button, a Jira issue creation form appear. This form includes the Snyk issue details, which are copied into the relevant fields. You can review and edit this form before creating the issue.

Select the Jira project to which you want to send the issue. The fields in the example below are based on the fields that the specific project has, so switching between projects may show different options.

![](../../.gitbook/assets/uuid-67202f8e-7f70-1e84-6044-f65ec36138b3-en.png)

After you created a Jira issue, the Jira key with a link are displayed on the issue card. If you are using the Jira API, you can generate multiple Jira issues for the same issue in Snyk.

![](<../../.gitbook/assets/Jira - Button with a link.png>)

You can also see which Jira issues have been created on the **Issues** view in your Reports.

![](<../../.gitbook/assets/Jira - Isuues in Report.png>)

### See also:

[Setup Broker with Jira](../snyk-broker/snyk-broker-set-up-examples/setup-broker-with-jira.md)
