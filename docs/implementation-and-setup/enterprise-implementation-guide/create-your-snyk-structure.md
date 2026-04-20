# Create your Snyk structure

To create your ideal Snyk structure, reflecting the way your business is structured, you need to clone the template Organization created in the previous phase of this guide. This enables you to create multiple Organizations that cover your critical business units.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1775661958/5._Creating_your_Organization_Structure_pheaur.mp4" %}
Create your Organization structure video guide
{% endembed %}

## Clone the template Organization

{% hint style="success" %}
**Key decision**: Choose between using the Snyk web UI for manual creation or the API for automated, bulk Organization provisioning.
{% endhint %}

After your template is ready, use it to build your structure:

* **Through the web UI**: Select **Template** in the Copy settings from dropdown when you create a new Organization.
* **Through the** [**API**](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/): Use the `sourceOrgId` parameter in the `Create a new organization` endpoint.

{% hint style="info" %}
Snyk recommends using multiple template Orgs for different regions or business units that use the same SCM. This is because different teams or business units may have different risk profiles, legal requirements, and ways of working with the chosen SCM.
{% endhint %}

Use the following table to understand which configurations transfer to the new Organization.

| Settings cloned                                        | Settings not cloned                                 |
| ------------------------------------------------------ | --------------------------------------------------- |
| All integrations and settings                          | Members and Service Accounts                        |
| SCM and Container settings                             | Existing Projects                                   |
| Notification integrations (for example Slack, or Jira) | Custom policies and ignore settings                 |
| PaaS and serverless integrations                       | Infrastructure as Code (IaC) and Snyk Code settings |

## Configure scan and notification behavior

{% hint style="success" %}
**Key decision**: Decide at what stage of the rollout developers should start receiving automated alerts and PR feedback.
{% endhint %}

Snyk recommends a "quiet" start to avoid notification fatigue during the initial import:

* **Disable PR checks**: In your SCM integration settings, turn off automatic PR tests and fix PRs. Enable these only when you reach the Prevention stage of your rollout.
* **Silence notifications**: Disable email notifications at both the Group and Organization levels during the initial setup.
