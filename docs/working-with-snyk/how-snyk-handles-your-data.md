# How Snyk handles your data

Snyk is a developer security platform placing the utmost importance on data security. Fully understanding your privacy and security needs, Snyk provides this document with the goal of providing you with transparency as to how and what data is accessed, transferred, and stored by Snyk.

The data that Snyk handles varies depending on the product you are using, how you are integrating with Snyk and your Snyk deployment. Be aware, Snyk is subject to fast-moving changes, therefore, the types of data that Snyk accesses and stores might change with the introduction of a new capability or changes to an existing capability.

## Flexible deployment options

Snyk leverages the latest software development practices and technologies to provide customers with the flexibility to use Snyk’s developer security platform in the manner that best suits the needs of their business.

Snyk’s cloud-first deployment options offer ease of use and scalability while also providing the required level of data protection with multi- and single-tenancy options supported in the US, EU, and AUS.

In the future, Snyk will support additional regions.

* **Multi-Tenant SaaS**: The simplest, most common, and most cost-effective way to use Snyk’s developer security platform
* **Single-Tenant SaaS:** Private Cloud - An isolated and fully managed instance of the Snyk developer security platform on AWS. For details, see [Snyk Deployment Options](https://snyk.io/platform/deployment-options/).
* **Snyk Broker**: A client service installed on your private infrastructure, acting as a proxy between the Snyk developer security platform (multi- or single-tenant) and your on-premise codebase. [Snyk Broker](../enterprise-setup/snyk-broker/) securely handles inbound and outbound connections, encrypting data during transit and deliberately controlling the access Snyk has to your data. Sensitive credentials stay behind your firewall.

## Customer data flows across Snyk

Snyk provides a wide range of development tools and integration points, requiring different types of data and involving different data interactions. The following sections provide an overview of both the common types of data Snyk accesses and stores and product and integration-specific types. The information is reviewed at least twice per year or when a significant change occurs within the product operations.

## Common data types

* **Vulnerability data** - Snyk stores information on the vulnerabilities identified in customer applications and related fix context.
* **Vulnerability source** - Snyk stores information on where the vulnerability was identified. Examples: source code repository or registry, file name and location, dependency tree, vulnerability path.
* **Integration-related data** - Snyk stores information required to set up an integration with Snyk. Examples: tokens and API configurations.
* **User data** - Snyk stores user information required to access and use the platform. Examples: user name, IDs (for example, GitHub user ID), email address, IP address.
* **User list** - For purposes of accurate contributor counting, Snyk accesses commits from the last 90 days for repositories monitored. Upon request, an unhashed version of user emails is produced.
* **Billing data** - Snyk stores information required for billing your Snyk account.
* **User behavior analytics** - Snyk stores various types of information pertaining to usage patterns. Examples: platform navigation and executed CLI commands.

{% hint style="info" %}
All data is handled by Snyk in accordance with the SOC 2 standard. For more details, see [Snyk certifications](how-snyk-handles-your-data.md#snyk-certifications).
{% endhint %}

### Cache retention period related to vulnerability source data

Snyk uses cloud products from Amazon Web Services (AWS) and Google Cloud Platform (GCP) to process and store the vulnerability source data in a cache. Such data is cached according to the shortest possible duration afforded by the cloud service provider's storage lifecycle policy, which is subject to change by these cloud service providers and set out approximately in the periods below:

* For AWS tenants - EU/AU/Private Tenant: 24 to 48 hours under the[ S3 policy](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html)
* For GCP tenants - US (Default): 24 hours under the[ Google Cloud Storage Policy](https://cloud.google.com/storage/docs/lifecycle)

## Product-specific data types

Snyk knows how important it is to protect your data. Therefore we will only access and store the information needed to provide you with Snyk services and to ensure accurate analysis as described for each Snyk product: Snyk Open Source, Snyk Code, Snyk Container, and Snyk IaC.

### **Snyk Open Source**

<figure><img src="../.gitbook/assets/SnykOSS.svg" alt="Snyk Open Source logo"><figcaption><p>Snyk Open Source</p></figcaption></figure>

* Snyk accesses manifest files, lock files, and related configuration files in order to identify your open-source dependencies.
* By default, Snyk does not access your source code with this exception: for CLI scans using the `--unmanaged` option, Snyk accesses your source code files to convert them to file signatures (hashes) and store the file signatures and file names.
* Snyk accesses and stores the names and version numbers of your dependencies.
* Snyk stores the names of associated licenses, including copyright and attribution information.
* Snyk accesses and stores repository-specific information.
* Snyk accesses and stores Git provider push and pull-specific information. Examples: contributor name, filenames, timestamps.

**Optional Snyk Open Source add-ons (opt-in)**

Your account is subject to contract terms which might restrict your ability to enable these features. By enabling these features, you agree on behalf of your company to changes to your contract terms to allow these features, and you will be responsible for the use of these features based on your own circumstances.

* For the Go Modules full source code analysis feature, Snyk will access and store the contents of your Git repository to facilitate the building of an accurate dependency graph. After the Snyk analysis is complete, your code is deleted from the Snyk system.
* For the Reachable Vulnerabilities feature, Snyk will access and store the contents of your git repository to facilitate the building of a call graph. When the analysis completes, your code is deleted from the Snyk system. Only the call graph and function names are maintained.

### **Snyk Code**

<figure><img src="../.gitbook/assets/SnykCode.svg" alt="Snyk Code logo"><figcaption><p>Snyk Code</p></figcaption></figure>

* Snyk accesses your repository code for a one-time analysis and [caches the source code](how-snyk-handles-your-data.md#how-snyk-processes-this-data) according to the Cloud provider's storage minimum policy. After this period, only the location (file path, line, and column) of the issues found; the issue id, and explanations are maintained. Your code is removed and is not stored in the Snyk network or logs.
* Results are stored in a database and used for analytic and monitoring purposes by Snyk.
* Snyk Code does not use any customer code (1) for engine training purposes or (2) to extract examples to show possible fixes.
* The AI model for Snyk Code Fix Suggestions is trained on public repositories with permissive licenses, where any data from repositories with changing licenses are immediately removed. Static analysis, automated assessment, and partial human labeling are used during the data collection.
* The scan results do not contain original source code but rather pointers to positions (for example, files, line, and column numbers), plus identification meta-information so that results are displayed using the correct source code version.
* Snyk stores repository-specific information. Examples: Name of the Git repository and file names.
* The server infrastructure ensures separation between customers by using authentication and authorization. Snyk Code uses software controls to ensure customer data segregation. All communication is encrypted using high-grade industry-standard protocols.

### **Snyk Container**

<figure><img src="../.gitbook/assets/image (201) (1).png" alt="Snyk Container logo"><figcaption><p>Snyk Container</p></figcaption></figure>

* Snyk accesses and stores package versions, executable hashes and versions, operating system, container image metadata (for example, rootfs hashes, history), and image ID.
* Snyk accesses and stores information pertaining to the parent image: name, version, tag.
* Snyk accesses and stores RUN instructions from Dockerfile.
* Kubernetes configurations: Snyk accesses workload security settings (for example, ‘run as root’). This is only accessed if you use Snyk’s Kubernetes integration.
* Container registry integrations: Snyk accesses and then stores a short-term copy of the container image (unless a Broker is used). This copy is removed from the Snyk network after analysis.

### **Snyk IaC**

<figure><img src="../.gitbook/assets/SnykIaC.svg" alt="Snyk IaC logo"><figcaption><p>Snyk Infrastructure as Code</p></figcaption></figure>

#### Current IaC

* CLI tests are performed locally. When results are shared with the Snyk platform using the `--report` option, resource configurations are also included.
* SCM tests require access to your infrastructure as code files. Snyk stores them for the duration of the analysis and subsequently deletes them from the Snyk system. Snyk retains parsed resource configurations to provide context for issues and resources.
* Terraform Cloud and Terraform Enterprise tests analyze plan files. Snyk removes secrets and sensitive values and retains resource configurations to provide context for issues and resources.
* For drift detection using `snyk iac describe`, Snyk relies on the principle of least privilege and requires only read-only access to [AWS](../scan-with-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/configure-cloud-providers/configure-aws-provider.md#least-privileged-policy), [Azure](../scan-with-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/configure-cloud-providers/configure-azure-provider.md#least-privileged-policy), [Google](../scan-with-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/configure-cloud-providers/configure-google-provider.md#least-privileged-policy), or [GitHub](../scan-with-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/configure-cloud-providers/configure-github-provider.md#least-privileged-policy). Provider credentials are not sent to or stored by Snyk.
* Snyk relies on local read-only Terraform State file access and extracts and sends relevant resource configuration data to the platform.

#### IaC+

* Snyk Cloud scans cloud platform APIs to gather information on configured infrastructure deployed in AWS Accounts and Google Cloud Subscriptions.
* To perform scans, Snyk relies on the principle of least privilege, leveraging different authentication mechanisms that are supported by each Cloud platform.
  * For Amazon Web Services (AWS), a read-only AWS IAM role must be provisioned in your AWS Account(s) to provide secure access to required AWS APIs.
  * For Google Cloud, a read-only Google Cloud Service Account must be provisioned to enable secure access to required Google Cloud APIs.
* During scans, Snyk gathers and stores a resource configuration state to perform analysis, and stores the results of that analysis, including the details of misconfigurations that result in issues.
* Snyk Cloud retains resource configuration states found in scans to provide context for Issues and resources but does not store secrets or sensitive values.

### Snyk AppRisk

<figure><img src="../.gitbook/assets/AppRisk_Color_64px.png" alt="Snyk App Risk logo"><figcaption><p>Snyk AppRisk</p></figcaption></figure>

* Snyk AppRisk provides visibility into data at the Snyk Group level and therefore covers the scope of all Snyk Organizations within a Snyk Group.
* From Snyk Targets and Projects within associated Snyk organizations, Snyk AppRisk accesses and stores asset metadata to generate code repository assets, package (first-party) assets, and container image assets. Asset metadata includes the Git remote URL, languages represented in repositories, and commit history metadata (not the underlying source code).
* From SCM integrations configured using the Snyk AppRisk Integration Hub, Snyk AppRisk accesses and stores the following data:
  * Commit history metadata from the last 50 commits for repositories monitored, including profile information on the code committer, such as GitHub user IDs, display name, or email.
  * Metadata on the languages (for example, Python, HTML) used for a given code repository.
  * Repository “topics”, which are represented as asset tags.

## Snyk integrations: Git repository cloning

{% hint style="info" %}
For more information on Git repository cloning enablement and use, see [Git repository cloning for SCM integrations](../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/introduction-to-git-repository-integrations/git-repository-cloning-for-scm-integrations.md).
{% endhint %}

### Safeguards Snyk puts in place to ensure data is secure

* Snyk will perform a clone only when an SCM integration flow requires it, for example, PR check, import, and test.
* Communication between the cloning service and cache is encrypted with TLS 1.2.
* Cloned assets are deleted from the file system immediately after populating the cache.
* All data is handled by Snyk in accordance with the SOC 2 standard. For more details, see [Snyk certifications](how-snyk-handles-your-data.md#snyk-certifications).
* Your data is analyzed only for the purposes of improving code security and code quality and in accordance with your Data Processing Addendum.

### Git cloning applicable contract terms

By enabling this feature, you agree that your Git repository is a Protected Asset as defined in the contract between your company and Snyk.

## Snyk certifications

<figure><img src="../.gitbook/assets/Soc2.png" alt="Soc 2 Type 2 AICPA Soc"><figcaption><p>Soc 2 Type 2 AICPA Soc</p></figcaption></figure>

Snyk is certified to ISO 27001:2013 with the additional objective controls of ISO 27017:2015.

<figure><img src="../.gitbook/assets/Coalfire.png" alt="Coalfire ISO 27001 certified ISO"><figcaption><p>Coalfire ISO 27001 certified ISO</p></figcaption></figure>

## Snyk policies

For additional information, see the following pages on the Snyk website:

* [Privacy](https://snyk.io/policies/privacy/)
* [Snyk Sub-processing](https://snyk.io/policies/subprocessors/)
* [Data Processing Addendum](https://snyk.io/policies/dpa/)
* [Tracking and analytics](https://snyk.io/policies/tracking-and-analytics/)
