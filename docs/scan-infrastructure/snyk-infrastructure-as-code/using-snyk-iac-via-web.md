# Using Snyk IaC with the Web UI

You can use Snyk IaC with the [Snyk Web UI](../../getting-started/quickstart/create-a-snyk-account/logging-in-to-an-existing-account.md) to find and fix issues in configuration files.

1. In your **Projects** listing, select the Project to open.
2. Examine the information and issue cards for that Project:

<figure><img src="../../.gitbook/assets/image (2) (3) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Snyk Project issue card"><figcaption><p>Snyk Project issue card</p></figcaption></figure>

The information available shows [Snyk Project ](../../manage-issues/snyk-projects/)information including:

* Snapshot information showing when the Project was last tested
* **Overview**, **History,** and **Settings** information.\
  You can use the **History** section to view previous snapshots of Project.
* Filters on the left of the screen

## Issue card details

Each issue card shows specific details about that issue:

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 14.24.14.png" alt="Issue card details"><figcaption><p>Issue card details</p></figcaption></figure>

Details on the issue card include:

* The severity level, for example, **H** for high, and the name of the issue, for example, **Non-encrypted S3 Bucket**
* The **ID** of the security rule, for example, [SNYK-CC-00172](https://security.snyk.io/rules/cloud/SNYK-CC-00172).\
  Click the link to view more information on the [Snyk Security Rules](https://security.snyk.io/rules/cloud/).
* A **snippet** of your code showing the exact area that is vulnerable
* The exact **path** of the issue
* More details, such as:
  * a short **description** of the issue
  * the **impact** of the issue
  * the **remediation** advice to resolve the issue

Click **Full details** to see a preview of the full code:

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 14.24.20.png" alt="Preview of the full code"><figcaption><p>Preview of the full code</p></figcaption></figure>

Click **Ignore** to ignore this vulnerability. For details, see [Ignore Issues](../../manage-issues/priorities-for-fixing-issues/ignore-issues.md).

## Examples of results displayed

### Terraform Cloud and Helm examples

Terraform Cloud and Helm do not show a code snippet, only the card details. There is no **Full details** button to show the preview of the full code.

<figure><img src="../../.gitbook/assets/image (114) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1).png" alt="Details for Helm"><figcaption><p>Details for Helm</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (100) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (3) (2).png" alt="Details for Terraform Cloud"><figcaption><p>Details for Terraform Cloud</p></figcaption></figure>

### Examples without the line with the vulnerable path

If Snyk can not identify the exact line of the vulnerable path in the file, Snyk does not show a code snippet, only an info message and the card details. If possible, Snyk shows the **Full details** button so you can see a preview of the full code.

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 14.28.07 (1).png" alt="Card detils with info message"><figcaption><p>Card detils with info message</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/Screenshot 2022-05-23 at 14.28.17 (1).png" alt="Full details display"><figcaption><p>Full details display</p></figcaption></figure>
