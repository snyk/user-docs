# Getting Started with Snyk License Compliance Management

{% hint style="info" %}
**Feature availability**\
This feature is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Get started with Snyk license compliance management, to check compliance for the [open source licenses](https://snyk.io/learn/open-source-licenses/) in your code.

The Snyk Default License Policy defines how Snyk identifies potential license issues in the open-source packages your Projects are using. This policy applies to all Organizations created in your Group.

## **Prerequisites**

Ensure you have:

* A Snyk [paid plan](https://snyk.io/plans/).
* Integrated and imported your Projects, as described for [Snyk Open Source](../getting-started-snyk-open-source.md).

## **Stage 1: Define policies**

To take effective action based on license issues, you need to define policies defining these actions, based on license types. Policies provide a way to capture different requirements within an Organization, based on factors such as line of business. Work with your legal team to create policies that are specific to your company.

To open your Snyk Group default license policy, select the **Policies** menu option in your Group:

<div align="left">

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-15 at 12.34.36.png" alt="Access Group policies"><figcaption><p>Access Group policies</p></figcaption></figure>

</div>

### Create policy rules

Each policy contains rules, detailing which licenses are acceptable and which are forbidden for use, together with a severity level that indicates how severe the license violation is. For example, severity levels for internal-only license issues may be less severe than for those released externally.

You can create and edit multiple license policies for Organizations. See [Create a license policy and rules](../../../manage-issues/policies/license-policies/create-a-license-policy-and-rules.md) for details.

## Stage 2: View issues

Snyk’s [Git-based integrations ](../../../integrations/git-repository-scm-integrations/)support license scanning as part of the regular workflow. During scanning, license issues appear as a filterable list in the **Issues** tab:

<figure><img src="../../../.gitbook/assets/image3 (1) (2).png" alt="Issues overview within a dependency project."><figcaption><p>Issues overview within a dependency project</p></figcaption></figure>

This example shows a high-severity issue for a GPL-2.0 license, with accompanying instructions as defined in the policies for that license.

You can also view license issues using the Snyk CLI tool, after running **snyk test**:

<figure><img src="../../../.gitbook/assets/image2-1-.png" alt="License issue overview in Snyk CLI."><figcaption><p>License issue overview in Snyk CLI</p></figcaption></figure>

**View dependencies**

Snyk shows license issues in both your direct and transitive dependencies, in a full dependency tree to show what dependency introduced the license issue.

![Dependencies overview within a dependency project](<../../../.gitbook/assets/image4 (1).png>)

This example includes two high severity license policy violations, caused by:

* a direct dependency on an npm package called **wicket@1.3.5**
* a transitive dependency on a package called **flickity@2.2.1** introduced by **web-project-starter@0.0.3**

### **View all license information**

You can view and share detailed lists of licenses being used by all Projects in your Organization, and see a report that lists all the Open Source components and licenses along with copyright information.

See [Dependencies and licenses](../../../manage-issues/dependencies-and-licenses/) for more details.

## **Stage 3: Process issues**

You can now take action to resolve the license issues identified during the scan, to help you build and deploy your application without outstanding licensing issues.

The actions you take depend on the license conditions and on your policies. For example, if a license violation is surfaced, this issue can be mitigated by either approaching your legal team or by replacing the dependency which added the violation.
