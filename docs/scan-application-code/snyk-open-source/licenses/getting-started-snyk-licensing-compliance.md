# Getting Started with Snyk License Compliance Management

**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.

Get started with Snyk license compliance management, to check compliance for the [open source licenses](https://snyk.io/learn/open-source-licenses/) in your code, as part of your [Snyk Open Source](../open-source-basics/) solution.

The Snyk Default License Policy defines how Snyk identifies potential license issues in the open source packages your projects are using. The default policy applies to all organizations created within your group.

{% hint style="info" %}
This process describes how to use the Snyk Web UI and a [supported source code management](https://support.snyk.io/hc/en-us/sections/360001138098-Git-repository-SCM-integrations) system.\
You can also use an [IDE tool](../../../integrations/ide-tools/) or a [CI/CD integration](../../../integrations/ci-cd-integrations/), or use the [Snyk CLI tool](../../../snyk-cli/) to get started using the command line.
{% endhint %}

## **Prerequisites**

Ensure you have:

* A Snyk [paid plan](https://snyk.io/plans/).
* Integrated and installed your projects, as described for [Snyk Open Source](../getting-started-snyk-open-source.md).

## **Stage 1: Define policies**

To take effective action based on license issues, you need to define policies defining these actions, based on license types. Policies provide a way to capture different requirements within an organization, based on factors such as line of business. Work with your legal team to create policies that are specific to your company.

To open your Snyk Group default license policy, complete the following steps:

1. From your Snyk Organization, **Switch group** and select the relevant Group Overview.
2. Select the **Policies** tab.

<figure><img src="../../../.gitbook/assets/polices4license_mgmt-03oct2022.png" alt="Defining policy."><figcaption><p>Defining policy</p></figcaption></figure>

### Create policy rules

Each policy contains rules, detailing which licenses are acceptable and which are forbidden for use, together with a severity level that indicates how severe the license violation is. For example, severity levels for internal-only license issues may be less severe than for those released externally.

To define your group license policy, complete the following steps:

1. Select **Snyk Default License Policy**.
2. Enter a **Description** for the policy to summarize what it does and the reason it was added.
3. For each license listed on the right, choose a **Severity** from the menu list to define which license issues you want to identify when Snyk tests run.
4. If you select a severity other than **None**, and you want to include additional instructions that will display when that license issue is identified, select the icon to the right of the Severity dropdown and enter the text for the license instruction.
5. Click **Add** or **Update** to confirm your changes.

<figure><img src="../../../.gitbook/assets/license-policy.png" alt="Policy configuration."><figcaption><p>Policy configuration</p></figcaption></figure>

The additional instructions show in the Snyk Web UI and the CLI when the selected licenses are identified.

See [Licenses overview](./) and [Setting a license policy](../../../manage-issues/policies/license-policies/create-a-license-policy-and-rules.md).

## Stage 2: View issues

Snyk’s [Git-based integrations ](../../../integrations/git-repository-scm-integrations/)support license scanning as part of the regular workflow. During scanning, license issues appear as a filterable list in the **Issues** tab:

<figure><img src="../../../.gitbook/assets/image3 (1).png" alt="Issues overview within a dependency project."><figcaption><p>Issues overview within a dependency project</p></figcaption></figure>

This example shows a high-severity issue for a GPL-2.0 license, with accompanying instructions as defined in the policies for that license.

You can also view license issues using the Snyk CLI tool, after running **snyk test**:

<figure><img src="../../../.gitbook/assets/image2-1-.png" alt="License issue overview in Snyk CLI."><figcaption><p>License issue overview in Snyk CLI</p></figcaption></figure>

**View dependencies**

Snyk shows license issues in both your direct and transitive dependencies, in a full dependency tree to show what dependency introduced the license issue.

![Dependencies overview within a dependency project](<../../../.gitbook/assets/image4 (1).png>)

This example includes two high severity license policy violations, caused by:

* a direct dependency on an npm package called **wicket@1.3.5**
* a transitive dependency on a package called **flickity@2.2.1** introduced by **web-project-starter@0.0.3**

**View lists and copyrights**

You can view and share detailed lists of licenses being used, and see a report that lists all the Open Source components and licenses along with copyright information.

<figure><img src="../../../.gitbook/assets/copyright.png" alt="Dependencies overview with focus on licensing and copyright information."><figcaption><p>Dependencies overview with focus on licensing and copyright information</p></figcaption></figure>

## **Stage 3: Process issues**

You can now take action to resolve the license issues identified during the scan, to help you build and deploy your application without outstanding licensing issues.

The actions you take depend on the license conditions and on your policies. For example, if a license violation is surfaced, this issue can be mitigated by either approaching your legal team or by replacing the dependency which added the violation.
