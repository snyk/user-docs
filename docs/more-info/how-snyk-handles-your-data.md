# How Snyk handles your data

## Snyk security and compliance: How does Snyk handle your data?

Snyk is a developer security platform and as such, we place the utmost importance on data security. **Fully understanding your privacy and security needs, the goal of this document is to provide you with transparency as to how and what data is accessed, transferred, and stored by Snyk.**

The data handled by Snyk will vary depending on the product you are using, how you are integrating with Snyk, and your Snyk deployment. Because Snyk is a fast moving product, the types of data accessed and stored might change with the introduction of a new capability or changes to an existing capability.

### Flexible deployment options

Snyk leverages the latest software development practices and technologies to provide customers with the flexibility to use Snyk’s developer security platform in the manner that best suits the needs of their business.

Snyk’s cloud-first deployment options offer ease of use and scalability while also providing the required level of data protection with multi and single tenancy options supported in the US, EU and AUS (additional regions will be supported in the future).

* **Multi-Tenant SaaS**: The simplest, most common and most cost-effective way to use Snyk’s developer security platform
* **Single-Tenant SaaS:** Private Cloud - An isolated and fully managed instance of the Snyk developer security platform on AWS.
* **Snyk Broker**: A client service that is installed on your private infrastructure, acting as a proxy between the Snyk developer security platform (multi or single tenant) and your on-premise codebase. The Snyk Broker securely handles inbound and outbound connections, encrypting data during transit, and deliberately controlling the access Snyk has to your data. Sensitive credentials stay behind your firewall.

**Be sure to reach out to your Snyk contact to find out more details on how these different options can be leveraged to meet your needs.**

### Customer dataflows across Snyk

Snyk provides a wide range of development tools and integration points, requiring different types of data, and entailing different data interactions. The sections below provide an overview of both the common types of data Snyk accesses and stores as well as product and integration-specific types. The information is reviewed as a minimum of twice per annum or when a significant change occurs within the product operations.

### Common data types

* **Vulnerability data** - Snyk stores information on the vulnerabilities identified in customers applications and related fix context.
* **Vulnerability source** - Snyk stores information on where the vulnerability was identified. Examples: source code repository/registry, file name and location, dependency tree, vulnerability path.
* **Integration-related data** - Snyk stores information required to set up an integration with Snyk. Examples: tokens and configurations.
* **User data** - Snyk stores user information required to access and use the platform. Examples: user name, IDs (e.g. GitHub user ID), email address, IP address.
* **User list** - For the purposes of an accurate contributor counting, Snyk accesses commits from the last 90 days for repositories monitored. Upon request, an unhashed version of user emails is produced.
* **Billing data** - Snyk stores information required for billing your Snyk account.
* **User behavior analytics** - Snyk stores various types of information pertaining to usage patterns. Examples: platform navigation, executed CLI commands.

### Product-specific data types

We know how important it is for you to protect your data. Our products only access and store the information needed to provide you with Snyk services, and to ensure accurate analysis.

{% tabs %}
{% tab title="Snyk Open Source" %}
**Snyk Open Source**

****![](../.gitbook/assets/SnykOSS.svg)****

* Snyk accesses manifest and build configuration files in order to identify your open source dependencies.
* For tests using the `--unmanaged` flag accesses your source code files to convert them to file signatures (hashes), and store the file signatures and file names. Other Snyk Open Source scans do not access your source code.
* For SCA Scans Snyk does not access your source code
* Snyk accesses and stores the names and version numbers of your dependencies.&#x20;
* Snyk stores the names of associated licenses, including copyright and attribution information.&#x20;
* Snyk accesses and stores repository-specific information.
* Snyk accesses and stores Git provider push & pull specific information. Examples: contributor name, filenames, timestamps.

**Optional ADD-ONS (opt-in)**

* For Reachable Vulnerabilities computation - Snyk stores your source code to facilitate the building of a call graph. Once the analysis completes, your code is removed from the Snyk system. Only the call graph and function names are maintained.
* For Lambda integration only - Snyk pulls a short term copy then destroys it as part of analysis.
{% endtab %}

{% tab title="Snyk Code" %}
**Snyk Code**

****![](../.gitbook/assets/SnykCode.svg)****

* Snyk accesses your repository code for a one-time analysis, caching it for a period of up to 24 hours. After this period, only the location (file path, line, and column) to the issues found, the issue id and explanations are maintained. Your code is removed and is not stored in the Snyk network or logs.
* Results are stored in a database and used for analytic and monitoring purposes by Snyk.
* When viewing Snyk Code issue details on the Snyk App, the relevant files are loaded and cached up to 12 hours.
* Snyk Code does not use any customer code (1) for engine training purposes, or (2) to extract examples to show possible fixes.
* The scan results do not contain original source code but rather pointers to positions (e.g., files, line, and column numbers), plus identification meta-information so that results are displayed using the correct version of the source code.
* Snyk stores repository-specific information. Examples: Names of the Git repository, file names.
* The server infrastructure ensures separation between customers by using authentication and authorization. Snyk Code uses software controls to ensure customer data segregation. All communication is encrypted using high-grade industry-standard protocols.
{% endtab %}

{% tab title="Snyk Container" %}
**Snyk Container**

****![](../.gitbook/assets/SnykContainer.svg)****

* Snyk accesses and stores package versions, executable hashes/versions, operating system, container image metadata (e.g. rootfs hashes, history), image ID.
* Snyk accesses and stores information pertaining to the parent image - name/version/tag.
* Snyk accesses and stores RUN instructions from Dockerfile.
* Kubernetes configurations - Snyk accesses workload security settings (e.g. ‘run as root’). This is only accessed if you use Snyk’s Kubernetes integration.
* Container registry integrations - Snyk accesses and then stores a short term copy of the container image (unless a broker is used). This copy is removed from the Snyk network after analysis.
{% endtab %}

{% tab title="Snyk IaC" %}
**Snyk Infrastructure as Code**

****![](../.gitbook/assets/SnykIaC.svg)****

* CLI scans are performed locally. The contents of your files are not shared or stored by Snyk.
* Git-based scans require access to your Infrastructure as Code files. Snyk stores them for the duration of the analysis and subsequently deletes them from our system. Your IaC files are not stored by Snyk.
* Highlighted code snippets in the Snyk UI are generated dynamically, leveraging git-based scans from the provided source code repository integration. Your code snippets and IaC files are not stored by Snyk.
* Drift detection scans rely on the principle of least privilege.
  * Snyk does not require authorized access to your applications. Additionally, cloud credentials are never collected or stored by Snyk.
  * Snyk requires local read-only Terraform states access. The states are never transmitted to or analyzed by Snyk.
{% endtab %}
{% endtabs %}

### Snyk certifications

![](../.gitbook/assets/Soc2.png)

Snyk is certified to ISO 27001:2013 with the additional objective controls of ISO 27017:2015

![](../.gitbook/assets/Coalfire.png)

### Snyk policies:

* [Privacy](https://snyk.io/policies/privacy/)
* [Snyk Sub-processing](https://snyk.io/policies/sub-processors/)
* [Data Processing Addendum](https://snyk.io/policies/dpa/)
* [Tracking & analytics](https://snyk.io/policies/tracking-and-analytics/)
