# Getting Started with Snyk License Compliance Management

**Feature availability**  
Basic license policy configuration on a single default license policy is available with Business plans. Full policy creation and management is available with Enterprise plans. See [Pricing plans](https://snyk.io/plans/) for more details.

Get started with Snyk license compliance management, to check compliance for the [open source licenses](https://snyk.io/learn/open-source-licenses/) in your code, as part of your [Snyk Open Source](https://support.snyk.io/hc/en-us/sections/360003454998-Open-source) solution.

To take effective action based on license issues, you need to define policies defining these actions, based on license types. Policies provide a way to capture different requirements within an organization, based on factors such as line of business. Work with your legal team to create policies which are specific to your company.

#### Create policy rules

Each policy contains rules, detailing which licenses are acceptable and which are forbidden for use, together with a severity level which indicates how severe the license violation is. For example, severity levels for internal-only license issues may be less severe than for those released externally.

![license-policy.png](https://support.snyk.io/hc/article_attachments/360012847498/license-policy.png)

See [Licenses overview](https://support.snyk.io/hc/en-us/articles/360003557837-Licenses-overview) and [Setting a license policy](https://support.snyk.io/hc/en-us/articles/360007590258-Setting-a-license-policy).

Snyk’s [Git-based integrations](https://support.snyk.io/hc/en-us/sections/360001138098-Git-repository-SCM-integrations) support license scanning as part of the regular workflow. During scanning, license issues appear as a filterable list in the **Issues** tab:

![image3.png](https://support.snyk.io/hc/article_attachments/360012774437/image3.png)

This example shows a high-severity issue for a GPL-2.0 license, with accompanying instructions as defined in policies for that license.

You can also view license issues using the Snyk CLI tool, after running **snyk test**:

![image2.png](https://support.snyk.io/hc/article_attachments/360012774837/image2.png)

**View dependencies**

Snyk shows license issues in both your direct and transitive dependencies, in a full dependency tree to show what dependency introduced the license issue.

![image4.png](https://support.snyk.io/hc/article_attachments/360012775137/image4.png)

This example includes two high severity license policy violations, caused by:

* a direct dependency on an npm package called **wicket@1.3.5**
* a transitive dependency on a package called **flickity@2.2.1** introduced by **web-project-starter@0.0.3**

**View lists and copyrights**

You can view and share detailed lists of licenses being used, and see a report that lists all the open source components and licenses along with copyright information.

![copyright.png](https://support.snyk.io/hc/article_attachments/360013041917/copyright.png)

You can now take actions to resolve the license issues identified during the scan, to help you build and deploy your application without outstanding licensing issues.

The actions you take depend on the license conditions and on your policies. For example, if a license violation is surfaced, this issue can be mitigated by either approaching your legal team, or by replacing the dependency which added the violation.

### For more information

See [Licenses](https://support.snyk.io/hc/en-us/sections/360001010457-Licenses).

