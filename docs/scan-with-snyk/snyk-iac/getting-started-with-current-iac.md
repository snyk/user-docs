# Getting started with Snyk IaC

You can use Snyk IaC (Infrastructure as Code) in the Snyk Web UI to find, view, and fix issues in configuration files. You can also use Snyk IaC in the Snyk CLI. For details, see [Snyk CLI for Infrastructure as Code](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/).

On this page, you will find steps to find, view, and fix issues in configuration files for the supported environments: [Terraform](scan-your-iac-source-code/scan-terraform-files/), [AWS CloudFormation](scan-your-iac-source-code/scan-cloudformation-files/), [Kubernetes](scan-your-iac-source-code/scan-kubernetes-configuration-files/), including Helm, and [Azure Resource Manager (ARM)](scan-your-iac-source-code/scan-arm-configuration-files.md). These steps are specific to the current IaC.&#x20;

## **Prerequisites for Snyk IaC**

Before using Snyk IaC, be sure you have the prerequisites as follows:

* A Snyk account. For details, see [Getting started](../../discover-snyk/getting-started/).
* An existing Terraform, CloudFormation, Kubernetes, or ARM environment to work in.
* A Git repository you have integrated with Snyk in the same way as for other Snyk products. For details, see [Git repository (SCM)](../../developer-tools/scm-integrations/organization-level-integrations/).

For more information about IaC and supported environments, see the following pages:

* [Configure your integration to find security issues in your Terraform files](scan-your-iac-source-code/scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-files-current-iac.md)
* [Configure your integration to find security issues in your CloudFormation files](scan-your-iac-source-code/scan-cloudformation-files/configure-your-integration-to-find-security-issues-in-your-cloudformation-files-current-iac.md)
* [Configure your integration to find security issues in your Kubernetes configuration files](scan-your-iac-source-code/scan-kubernetes-configuration-files/configure-integration-to-find-security-issues-in-kubernetes-configuration-files-current-iac.md)

{% hint style="info" %}
You must use the Snyk CLI to scan ARM configuration files. See [Scan ARM configuration files](scan-your-iac-source-code/scan-arm-configuration-files.md).
{% endhint %}

## Import IaC Projects

You will start by importing [Projects](../../snyk-platform-administration/snyk-projects/) you want to scan with Snyk. In these steps, you choose repositories for Snyk to test and re-test:

1. Log in to Snyk and on your dashboard, select **Projects** from the navigation.
2. On the Projects page, from the **Add projects** dropdown, select the SCM where the repositories and projects that you want to scan are; for example, select GitHub.
3. From the list of **Personal and Organization repositories**, select the Git repositories and projects you want to import for scanning.\
   You can select one or more repositories or projects in a repository.
4. Click **Add selected repositories** to import the selected SCM projects and repositories into Snyk.
5. Select **View import Log** to see the results on the import log.\
   You can scan multiple types of configuration files simultaneously.\
   The import completes and the Projects page displays the Snyk Project imported.

{% hint style="info" %}
After you have imported an IaC Project, Snyk re-tests your Project once a week by default. You can de-activate recurring tests on the **Settings** tab of the Projects page; Set **Test & Automated Pull Request Frequency** to **Test never**.
{% endhint %}

## View configuration file issues in IaC

On the Projects page, you can view the results for configuration files in the imported Projects.

* If **Group by targets** is selected, a list of [Targets](../../snyk-platform-administration/snyk-projects/#target) is displayed. These are the repositories with the Projects you imported. Select a Target to expand its list of Projects.
* If **Group by none** is selected: A list of all [Projects](../../snyk-platform-administration/snyk-projects/#project) is displayed.

In your **Projects** listing, select the Project to open to display detailed information about that Project.

<figure><img src="../../.gitbook/assets/snyk-iac-getting-started-list-of-projects.png" alt="A list of Snyk IaC Projects"><figcaption><p>List of Snyk Projects</p></figcaption></figure>

Each Project detail page has a snapshot showing when the Project was last tested, the name of the user who imported the Project, and, on the **Issues** tab, the number of critical, high, medium, and low-severity issues found and issue cards for each scanned configuration file. You can also select the **Overview**, **History,** and **Settings** options. Choose **History** to see previous snapshots of the Project.

<figure><img src="../../.gitbook/assets/image (106).png" alt="Snyk Project issue card"><figcaption><p>Snyk Project issue card</p></figcaption></figure>

## Issue card details for Snyk IaC

Each issue card shows information about the resource and the path by which it was introduced.

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 14.24.14.png" alt="Issue card details"><figcaption><p>Issue card details</p></figcaption></figure>

The information on the issue cards includes the following:

* The severity level, for example, **H** for high, and the name of the issue, for example, **Non-encrypted S3 Bucket**
* The **ID** of the security rule, for example, [SNYK-CC-TF-99](https://security.snyk.io/rules/cloud/SNYK-CC-TF-99).\
  Click the link to view more information on the [Snyk Security Rules](https://security.snyk.io/rules/cloud/).
* A snippet of your code showing the exact area that is vulnerable
* The exact path of the issue
* More details, such as:
  * brief description of the issue
  * impact of the issue
  * remediation advice to resolve the issue

Click **Full details** to see a preview of the full code:

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 14.24.20.png" alt="Preview of the full code"><figcaption><p>Preview of the full code</p></figcaption></figure>

Click **Ignore** to ignore this vulnerability. For details, see [Ignore Issues](../../manage-risk/prioritize-issues-for-fixing/ignore-issues/).

## Fix configuration files in IaC

The steps to act on recommendations produced by Snyk IaC follow.

1. On a Project detail page, select an issue to see the details for that issue and specific recommendations from Snyk IaC.
2. Based on the recommendations, edit the configuration file to fix the issue identified and then commit the change.\
   Snyk automatically rescans the changed file.
3. View the change reflected in the issue display.

<figure><img src="../../.gitbook/assets/snyk-iac-getting-started-issue-card.png" alt="Example of an IaC issues that has been fixed"><figcaption><p>Example of an IaC issues that has been fixed</p></figcaption></figure>

## Examples of IaC results

Examples follow of results displayed for current IaC.

### Terraform Cloud and Helm examples

Terraform Cloud and Helm do not show a code snippet, only the path details. There is no **Full details** button to show the preview of the full code.

<figure><img src="../../.gitbook/assets/image (114) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1).png" alt="Details for Helm"><figcaption><p>Details for Helm</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (100) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (3) (2).png" alt="Details for Terraform Cloud"><figcaption><p>Details for Terraform Cloud</p></figcaption></figure>

### Example showing the code preview is not available

If Snyk can not identify the exact line of the vulnerable path in the file, Snyk does not show a code snippet, only a message and the path details. If possible, Snyk shows the **Full details** button so you can see a preview of the full code.

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 14.28.07.png" alt="Issue card without code snippet"><figcaption><p>Issue card without code snippet</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 14.28.17.png" alt="Full code display"><figcaption><p>Full code display</p></figcaption></figure>
