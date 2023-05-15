# Jira integration

A new integration, Snyk Security for Jira Software, is coming soon. If you are interested in early access, you can [request it](https://earlyaccessprogram.atlassian.net/servicedesk/customer/portal/12/group/13/create/56). This page covers the integration available in Snyk.

{% hint style="info" %}
**Feature availability**\
This feature is available with all paid plans. See [pricing plans](https://snyk.io/plans/) for more details.\


For the availability of this feature for Snyk Infrastructure as Code, see [Jira Integration for IaC](../../scan-cloud-deployment/snyk-infrastructure-as-code/jira-integration.md).
{% endhint %}

## **Set up your Jira integration**

Snyk Jira integration allows you to manually raise Jira issues in the Snyk UI for vulnerabilities or license issues. The Jira integration also includes APIs ([see Snyk API docs](https://snyk.docs.apiary.io/#reference/projects/project-jira-issues)).

{% hint style="info" %}
If your Jira instance is private, use [the Snyk Broker deployment method](../../snyk-admin/snyk-broker/install-and-configure-broker-using-docker/setup-broker-with-jira.md).
{% endhint %}

## **Prerequisites for Jira integration with Snyk**

* Snyk supports Jira version 5 and later versions.
* The following [Jira permissions](https://confluence.atlassian.com/adminjiraserver073/managing-project-permissions-861253293.html) are required: **Browse Projects** and **Create Issues.**

## **How to set up your Jira integration**

The Jira account credentials are configured in the Snyk Web UI: **Organization Settings > Integrations** page.

It is best practice to set up a new user in Jira for this integration, instead of using the credentials of an existing account.

Cloud-hosted Jira implementations require a username and API token authentication. Jira API tokens are generated in [Atlassian API tokens](https://id.atlassian.com/manage/api-tokens). Self-hosted implementations can also authenticate with a username and password.

## **Create a Jira issue**

After you set up the Jira integration connection, open one of your Snyk Projects in the Snyk Web UI. A new button, **Create an issue**, appears at the bottom of each vulnerability and license issue card. This button allows you to create a Jira issue.

<figure><img src="../../.gitbook/assets/Jira - new button.png" alt="Create an issue button"><figcaption><p>Create an issue button</p></figcaption></figure>

When you select **Create an issue**, a Jira issue creation form appears. This form includes the Snyk issue details, which are copied into the associated fields. You can review and edit this form before creating the issue.

Select the Jira project to which you want to send the issue. The fields in the example that follows are based on the fields that the specific Project has, so switching between Projects may show different options.

<figure><img src="../../.gitbook/assets/uuid-67202f8e-7f70-1e84-6044-f65ec36138b3-en.png" alt="Crate a Jira issue"><figcaption><p>Crate a Jira issue</p></figcaption></figure>

After you create a Jira issue, the Jira key with a link is displayed on the issue card. If you are using the Jira API, you can generate multiple Jira issues for the same issue in Snyk.

<figure><img src="../../.gitbook/assets/Jira - Button with a link.png" alt="Jira key on issue card"><figcaption><p>Jira key on issue card</p></figcaption></figure>

You can also see which Jira issues have been created in the **Issues** view in your Reports.

<figure><img src="../../.gitbook/assets/Jira - Isuues in Report.png" alt="Jira issues create shown in Reports"><figcaption><p>Jira issues create shown in Reports</p></figcaption></figure>

## Integrate with Jira using Snyk Broker

See [Set up Snyk Broker with Jira](../../snyk-admin/snyk-broker/install-and-configure-broker-using-docker/setup-broker-with-jira.md).
