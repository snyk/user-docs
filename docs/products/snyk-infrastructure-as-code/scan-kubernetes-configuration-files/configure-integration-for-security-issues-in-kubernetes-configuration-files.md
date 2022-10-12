# Configure integration for security issues in Kubernetes configuration files

Snyk tests and monitors Kubernetes configurations stored in your source code repositories and provides information, tips, and tricks to better [secure a Kubernetes environment](https://snyk.io/learn/kubernetes-security/). This helps to catch misconfigurations before they are pushed to production, as well as providing fixes for vulnerabilities.

### Supported Git repositories and file formats

Snyk currently scans your Kubernetes configuration files in JSON and YAML format when imported from your integrated Git repository.

### Configure Snyk to scan Kubernetes configuration files

#### **Prerequisites**

* Administrator access for the organization
* Check Git repository access and authorization--more info at [Git repository (SCM) integrations](https://support.snyk.io/hc/en-us/sections/360001138098-Git-repository-SCM-integrations).

#### **Configure Snyk**

Log in to the Snyk Web UI ([app.snyk.io](https://app.snyk.io)), and navigate to the relevant group and organization that you want to manage

{% hint style="info" %}
**Note**\
Integrations are managed per organization.
{% endhint %}

From here:

1. Enable Snyk to detect Kubernetes configuration files by enabling the flag in the settings ![cog\_icon.png](../../../.gitbook/assets/cog\_icon.png) > **Infrastructure as code** page:
2. If needed, review and adjust settings in the **Infrastructure as code** settings:

![](<../../../.gitbook/assets/image (112) (1) (1) (1) (1).png>)

The number of tests per product are based on account plans. For information on plans and test limits, [view our plans](https://snyk.io/plans/).
