# Glossary

## A

### Advisor

See [Snyk Advisor](https://snyk.io/advisor/).

### **Asset (Snyk Essentials)**

A Snyk Essentials asset is an identifiable entity that is part of an application, and relevant for security and developers. Snyk is generally focused on the development stages of application software, secures repository assets containing software package assets, and builds artifacts like container image assets.

### Application (Snyk **Essentials**)

An application is software that serves a business purpose and consists of assets that form the app. Organizations often define the scope of an application differently.

### Application graph

Represents the mapping of security issues, application assets, relationships between assets, and all relevant contextual information.

## B

### Base image

The parent image used to construct a container image, usually defined in the `FROM` directive in a Dockerfile. Base images themselves can be constructed from other base images.

### Broker

See [Snyk Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/).

### Build system

A system that takes the source code and builds the deployable application (such as a container).

### Business context

Information related to the organization's objectives, priorities, and regulatory requirements, such as criticality of the application to the business, compliance standards, data sensitivity, and potential impact on revenue or reputation.

## C

### CI/CD

Continuous integration (CI), continuous delivery (CD), and continuous deployment (CD) together comprise a Software Development Lifecycle (SDLC) model, guiding developers to automate the development and delivery of small, frequent changes. This ensures all team members have access to the latest codebase and can ensure the compatibility of committed code during development. See [Snyk CI/CD](../../developer-tools/snyk-ci-cd-integrations/) for details of Snyk CI/CD integrations.

### **Class (Snyk Essentials)**

A way to assign business context to assets and categorize an asset based on the business criticality. Assets can be assigned Classes A, B, C, or D, where Class A (assets that are business critical, deal with sensitive data, are subject to compliance, and so on) is the most important, and Class D (test apps, sandbox environments, and so on) the least important. Assets are assigned Class C by default. A class can be used in policies as well as defined in a policy.

### CLI

Command Line Interface. See [Snyk CLI](glossary.md#snyk-cli).

### Cloud Native Application Security

Implementing security throughout the CI/CD pipeline, automating security embedding in microservices, and maximizing repetition to reduce the introduction of vulnerabilities. Snyk provides a comprehensive [CNAS platform](https://snyk.io/product/cloud-native-application-security/). See [Cloud-native security guide for building secure applications](https://snyk.io/learn/cloud-native-security-for-cloud-native-applications/).

### Code assets (Snyk **Essentials**)

A hierarchical list of all assets retrieved from the scanned repositories.

### Container

Containers allow you to package applications and their dependencies together to be deployed as a single runnable unit. A container is an abstraction provided by the operating system kernel that allows a process to be isolated from other processes running on the system. See also [Snyk Container.](glossary.md#snyk-container)

### Container engine

For users, an application that takes a container image and turns it into a running container. Container engines typically interface with container registries and run containers. Examples of container engines include Docker, CRI-O, and LXC.

### Container image

One or more files that, when instantiated by a container engine or runtime, provide a running container. Images are the packaging and distribution format for containers.

### Container registry

A server that provides a mechanism to store and retrieve container images.

### **Controls (Snyk Essentials)**

The security controls associated with the asset. Navigate to the Snyk Essentials Controls section to see all available statuses for security controls.

### Coverage (Snyk Essentials)

An assessment of whether applicable assets are scanned and tested by security tools (like Snyk Open Source, for instance), as it relates to an application security program. A type of policy that allows you to specify what controls should be applied and, optionally, how often it needs to be run.

### Coverage gap (Snyk **Essentials**)

An assessment of all assets that fall "out of policy" and do not satisfy the coverage criteria you have specified, due to infrequent scanning or no scanning at all.

### CVE

Common Vulnerabilities and Exposures. A widely-used identifier for a well-known vulnerability.

### CVSS

Common Vulnerability Scoring System. An industry standard to assess the severity of vulnerabilities, using a score of 0 (lowest) to 10 (highest). Snyk uses CVSS.

### CWE

Common Weakness Enumeration. An online glossary that categorizes software and hardware weaknesses into different types, for example, CWE-20: Input Validation.

## D

### DAST

Dynamic Application Security Testing. A security analysis technique that tests a running application from the outside to find security issues. See also[ IAST](glossary.md#iast) and [SAST](glossary.md#sast).

### Dependency

When your application uses another package, this other package becomes a dependency in your own software.

* A direct dependency is a package you include in your own Project.
* An indirect dependency (also known as a deep, chained, or transitive dependency), is a package that is used by one of your direct dependencies.

### Dependency tree

Also known as Dependency path. A hierarchical graph showing the dependencies of a software application. This includes both direct and indirect dependencies and thus may be many levels deep.

### Development context

Information and requirements surrounding the development of applications within an organization, such as ownership, development tools, environments, teams, workflows, and processes.

### DevOps

A set of cultural philosophies, practices, and tools that combines software development and IT operations to shorten the systems development lifecycle.

### DevSecOps

The integration of security into emerging agile IT and DevOps development as seamlessly and as transparently as possible.

### Dockerfile

A text file format used to build container images using Docker. The Dockerfile contains all the commands needed to construct the final image, including specifying the parent base image.

## E

### Environment

Can refer to a cloud environment, a [Project attribute](../../snyk-platform-administration/snyk-projects/project-attributes.md), or an interface for working with Snyk, such as the Snyk [CLI](glossary.md#cli), [Web UI](glossary.md#snyk-web-ui), or an [IDE](glossary.md#ide).

### Exploit

A demonstration of how a vulnerability can be taken advantage of. When an exploit is widely published, it is commonly referred to as an exploit "in the wild". See [View exploits](../../manage-risk/prioritize-issues-for-fixing/view-exploits.md).

### Exploit Maturity

A measure of how practical an exploit for a vulnerability is, based on whether the exploit is in the wild, and how "helpful" the exploit is to attackers.

## F

### Fixable / Partially fixable

A measure of whether a vulnerability can be fixed by Sny by applying a patch, upgrade, or pin. See [Vulnerability fix types](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/vulnerability-fix-types.md).

### Fix PR

A pull request with an automatic fix for discovered vulnerabilities that Snyk can offer the user. See [Automated fix PRs](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-backlog-issues-and-known-vulnerabilities-backlog-prs.md).

## G

### Git

A distributed version control system for tracking changes in source code during software development.

## I

### IaC

Infrastructure as Code. See [Snyk Infrastructure as Code.](glossary.md#snyk-infrastructure-as-code)

### IAST

Interactive Application Security Testing. A runtime analysis tool that focuses on code behavior during execution to determine behaviors of interest to you, for example, what packages are loaded, where data flows in the running application, or how end users interact with the application. Capabilities vary between vendors and products. IAST works with a running application to analyze it internally by using sensors or agents placed within the application's runtime environment. IAST offers more detailed insights than [DAST](glossary.md#dast), for example, by tracing the source of vulnerabilities in the code. See also [SAST](glossary.md#sast).

### IDE

Integrated Development Environment. An application that has facilities for software development, typically with a source code editor, build automation tools, and a debugger.

### Image

The stored instance of a container that holds a set of software needed to run an application.

### Image layer

Container images typically consist of several different file system layers, which are combined together at runtime into a single file system.

### Integrations

Third-party products, applications, and platforms that Snyk works with, for example, SCM systems such as GitHub. See [Integrate with Snyk](../../integrations/integrate-with-snyk.md).

### Issue

A license problem, vulnerability, or misconfiguration identified and listed by Snyk. See [Find and manage priority issues](../../manage-risk/prioritize-issues-for-fixing/).

### Issue (Snyk **Essentials**)

An issue is a security problem identified by a Snyk security product when testing an asset, that AppSec teams need to remediate.

### Issues prioritization (Snyk **Essentials**)

Provides a centralized view of all the issues identified by Snyk with additional asset context. This empowers AppSec teams to better triage and remediate issues in Snyk.

### Issue context (Snyk **Essentials**)

Information surrounding a particular security issue that serves as objective risk factors such as issue severity level, availability of a fix, exploit maturity.

## L

### Library

A specific type of package.

### License policy

A set of criteria for evaluating open-source license issues. License policies enable you to set the severity level and define legal instructions for each license. See [License policies](../../manage-risk/policies/license-policies/).

## M

### Manifest

A file containing metadata about other files in a package.

### Monitor

The `snyk monitor` command tests a Project and uploads the results to Snyk. See the CLI help for [Monitor](../../developer-tools/snyk-cli/commands/monitor.md).

## O

### OCI

Open Container Initiative. An independent body set up to facilitate collaboration on standards for containers, to ensure they are interoperable between vendor solutions.

### Organization

An Organization in Snyk is a way to collect and organize your Projects. Members of Organizations have access to these Projects. See [Manage Groups and Organizations](../../snyk-platform-administration/groups-and-organizations/).

### Origin or source

The identifier for the ecosystem that a Target exists in. Snyk can scan Projects from multiple integrations, including CLI, API, GitHub, Kubernetes, and others. See [Snyk Projects](../../snyk-platform-administration/snyk-projects/).

## P

### Package

A group of files and additional metadata about those files, used by package managers.

### Package assets (Snyk **Essentials**)

Package assets are created when you scan the dependencies of a Project through package management systems or by using the Snyk CLI. This enables Snyk Essentials to identify and analyze the security vulnerabilities of the packages used within a Project, offering insights into possible risk exposures and providing recommendations for mitigation.

### Package manager

A set of tools that automate and manage packages of bundled files, and are usually specific to a language. For example, npm.

### Package registry

A software package hosting service that allows customers to host packages and code in one place.

### Pinnable

A fix type. Define and "pin" a specific version of an indirect dependency, to avoid a direct dependency pulling in a vulnerable version.

### Policy

See [license policy](glossary.md#license-policy), [security policy](glossary.md#security-policy), and [`.snyk` policy](glossary.md#snyk-policy).

### **Policy (Snyk Essentials)**

A way to automate actions in certain conditions, like classifying and tagging assets with business context. You can also use a policy to configure actions like sending a message or setting the coverage gap control using a Policy builder UI.

### PR

Pull Request. Allows a user to exchange changes made to source code and collaborate with others on the same branch.

### PR Checks

Use Snyk PR Checks to prevent new security issues from entering your codebase by automatically scanning code changes in real-time as soon as you submit a pull request (PR) in your source code manager (SCM). See [Run PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/).

### Priority Score

Snyk scores issues, including vulnerabilities and licenses for Open Source, to help prioritize the treatment of each one. Scores are based on multiple factors, including the CVSS score, and range from 0 (low) to 1000 (high). See [Priority Score](../../manage-risk/prioritize-issues-for-fixing/priority-score.md).

### Project

An external item scanned by Snyk with configuration to define how to run that scan. Projects appear on the **Projects** menu on the Snyk dashboard. See also [Target](glossary.md#target). For details, see [Snyk Projects](../../snyk-platform-administration/snyk-projects/).

## R

### Reachability

Whether an application contains code that will hit a vulnerable code path during execution. See [Reachable vulnerabilities](../../manage-risk/prioritize-issues-for-fixing/reachability-analysis.md).

### Registry

See [Container registry](glossary.md#container-registry) or [Package registry](glossary.md#package-registry).

### Repository

A storage area that contains all elements necessary for the distribution of an application.

### Repository assets (Snyk **Essentials**)

A repository asset is created by discovering the repositories directly in the SCM, when such integration is configured. Alternatively, a repository asset can be created by scanning a repository, (by Snyk or third-party tools) as long as the scanned code is identified with a specific repository.

### Resource

A cloud infrastructure entity such as an AWS S3 bucket, Identity and Access Management (IAM) role, or Virtual Private Cloud (VPC) flow log.

### Risk-based prioritization **(Snyk AppRisk)**

Assess the risk for each app based on the application context and conduct best-in-class security analysis. Provide fix guidance to direct developer remediation efforts towards the most critical business issues.

### Risk score

A value assigned to an issue, ranging from 0 to 1,000, representing the risk imposed on your environment.

### Rule

A security policy that checks cloud infrastructure and infrastructure as code (IaC) for misconfigurations that can lead to security problems, or a security rule used by Snyk Code when scanning your source code for vulnerabilities. For more information, see [Snyk Code security rules](../../scan-with-snyk/snyk-code/snyk-code-security-rules/) and [IaC custom rules](../../scan-with-snyk/snyk-iac/current-iac-custom-rules/).

### Runtime context (Snyk AppRisk)

Information on where and how an application is running.

## S

### SARIF

Static Analysis Results Interchange Format. A standard, JSON-based format for the output of static analysis tools.

### SAST

Static Application Security Testing. A security analysis technique that examines static source code to identify potential vulnerabilities without running the application. See also [IAST](glossary.md#iast), [DAST](glossary.md#dast), [Snyk Code](glossary.md#snyk-code), and [Snyk Infrastructure as Code](glossary.md#snyk-infrastructure-as-code).

### SBOM

Software Bill Of Materials. A list of components in a piece of software.

### SCA

Software Composition Analysis. A security analysis technique that is used to identify open-source and third-party components in use in an application, their known security vulnerabilities, and typically also adversarial license restrictions. Not to be confused with [Static Code Analysis](glossary.md#static-code-analysis). See also [Snyk Open Source](glossary.md#snyk-open-source).

### Scanned artifacts (Snyk **Essentials**)

A scanned artifact in Snyk Essentials is an entity detected by Snyk that cannot be identified as a repository asset because it does not include identifying information, such as a Git remote URL.

### SCM

Source Code Management. Also known as a code repository (repo) or version control system. The method used by developers to store their source code and track changes to code. SCM helps resolve conflicts when merging updates from multiple contributors. GitHub is an example of a common SCM system. See [Git repositories (SCMs)](../../developer-tools/scm-integrations/organization-level-integrations/).

### SCM Repository freshness (Snyk **Essentials**)

The SCM Repository freshness provides an immediate understanding of the current status of your repositories, including the date of the last commit. This assists you in quickly identifying active and dormant Projects and helps you with the decision-making regarding maintenance, security patching, and resource allocation. Reflects the status of the repository and the date of the last commit.

### SDLC

Software Development Lifecycle. A process followed by a development team, describing how to develop and maintain software.

### Secure at inception

Secure at inception is a strategic imperative of Snyk Studio which ensures that code generated by AI coding assistants (for example, Cursor, Devin, Claude Code) is immediately scanned, verified, and fixed for security issues before you accept the code. This prevents the introduction of new vulnerabilities into the codebase and prevents building on top of insecure foundations which could require more significant refactoring and necessary human intervention. To learn more about this concept, see[ Secure at Inception for AI with Snyk.](https://snyk.io/solutions/secure-at-inception-for-ai/) To learn more about Secure at inception in the context of agentic workflows, see [Secure at inception rules](../../integrations/snyk-studio-agentic-integrations/quickstart-guides-for-snyk-studio/windsurf-guide.md#secure-at-inception-rules).

### Security policy

A set of criteria for evaluating open-source vulnerabilities. Security policies enable you to set custom rules to automatically prioritize or de-prioritize specific vulnerabilities. See [Security policies](../../manage-risk/policies/security-policies/).

### Severity

A severity level is applied to a vulnerability or a license issue, to indicate the risk for that item in an application. See [Severity levels](../../manage-risk/prioritize-issues-for-fixing/severity-levels.md).

### Snapshot

An individual report within the test history of a Project. Includes a tree of dependencies and a list of vulnerabilities that was accurate at the time the test was conducted.

### `.snyk` policy

A policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins. See [The .snyk file](../../manage-risk/policies/the-.snyk-file.md).

### Snyk

A platform providing Cloud Native Application Security (CNAS) solutions, allowing developers to own and build security for the whole application, from code and open source to containers and cloud infrastructure. Snyk is also the company providing the Snyk platform. See [Getting started](./).

### Snyk Advisor

A free web application that allows you to compare software packages across open-source ecosystems. It provides insights into the overall health of a particular package by combining community and security data into a single unified view. See [Snyk Advisor](https://snyk.io/advisor/).

### Snyk API

A Snyk tool that enables developers to integrate programmatically with Snyk. See [Snyk API](../../snyk-api/snyk-api.md).

### Snyk Apps

Snyk Apps are the modern and preferred way to build integrations with Snyk, exposing fine-grained scopes for accessing resources over the Snyk APIs, powered by OAuth 2.0 for a developer-friendly experience. See [Snyk Apps](../../snyk-api/using-specific-snyk-apis/snyk-apps-apis/).

### Snyk Broker

A client/server system that serves as an agent or proxy, allowing Snyk to scan private customer environments: Jira, code repositories, or container registries. Snyk Broker relays messages and allows users to filter which messages are allowed through, for example, allowing users to expose only some GitHub APIs to Snyk. See [Snyk Broker](../../implementation-and-setup/enterprise-setup/snyk-broker/).

### Snyk CLI

A Snyk platform tool that enables developers to find and fix known vulnerabilities in dependencies, using a command line interface. See [Snyk CLI](../../developer-tools/snyk-cli/).

### Snyk Code

A Snyk product. A SAST product enabling developers to find and fix vulnerabilities in your proprietary application code. See [Snyk Code](../../scan-with-snyk/snyk-code/).

### Snyk Container

A Snyk product. Enables developers to find and fix vulnerabilities in container images and Kubernetes applications. See [Snyk Container](../../scan-with-snyk/snyk-container/).

### Snyk Infrastructure as Code

A Snyk product. Enables developers to find and fix vulnerabilities in Kubernetes, Helm, and Terraform configuration files. See [Snyk IaC](../../scan-with-snyk/snyk-iac/).

### Snyk Open Source

A Snyk product. Enables developers to find and fix open-source vulnerabilities. See [Snyk Open Source](../../scan-with-snyk/snyk-open-source/).

### Snyk plugin

A library used by the Snyk CLI to scan a certain language or build system.

### Snyk Studio

Snyk Studio embeds Snyk's AI security platform capabilities into any AI-native workflow. Snyk Studio is built on two core use cases: '[Secure at Inception](glossary.md#secure-at-inception),' which proactively prevents new, AI-generated vulnerabilities using configurable directives, and 'Intelligent Remediation,' which clears existing security backlogs at scale.

### Snyk Security Intelligence

A component powering the Snyk cloud-native application security platform.\
Incorporates the Snyk Intel Vulnerability DB: the Snyk database of vulnerabilities, providing detailed information and fix advice for known vulnerabilities. See [Vulnerability DB](https://snyk.io/vuln).

### Snyk web UI

The browser-based environment that provides users access to Snyk functions.

### Social Trends

Snyk shows a Trending banner on issues that are being actively discussed on X (formerly known as Twitter). See [Vulnerabilities with Social Trends](../../manage-risk/prioritize-issues-for-fixing/vulnerabilities-with-social-trends.md).

### Source

See [Origin](glossary.md#origin-or-source).

### SPDX

Software Package Data Exchange. A file format used to document information on the software licenses under which a piece of computer software is distributed. See [SPDX](https://spdx.dev/).

### Static Code Analysis

A technique for examining source code to identify issues related to code quality, structure, or performance, such as determining code reachability or spotting potential inefficiencies. While this technique may touch on security concerns, its primary focus is often broader, covering various aspects of code health. In contrast, Static Application Security Testing ([SAST](glossary.md#sast)) specifically targets the identification of security vulnerabilities within the code, such as coding flaws that could lead to security risks.

## T

### Target

Representation of an external resource Snyk has scanned. All [Snyk Projects](glossary.md#project) are associated with a parent Target. One Target may relate to many Projects. The structure of the Target depends on the [origin](glossary.md#origin-or-source).

### **Tags (Snyk Essentials)**

A way to categorize assets. Helps you recognize or handle assets differently according to mutual properties. Assets can be filtered by their tags in the inventory or when creating policy rules. A tag can be automatically assigned to an asset, or the asset can be tagged by a policy you created. GitHub and GitLab topics are treated as asset tags and you can use them for creating policies.

### Tenant

The top level of the Snyk hierarchy. It encompasses all your Groups and Organizations and all their corresponding Snyk work items. For more information, see [Tenants, Groups, and Organizations](../../snyk-platform-administration/groups-and-organizations/).

## U

### Upgradable / Patchable

A fix type: a problem can be fixed by upgrading a version of a package or by applying a patch.

## V

### Vulnerability

A security vulnerability that was identified by Snyk. See [Manage vulnerabilities](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/).

## W

### Webhook

A way for an app to provide other applications with real-time information. Snyk uses webhooks to check changes in code. See [Snyk Webhooks](../../snyk-api/using-specific-snyk-apis/webhooks-apis/).

### Web UI

See [Snyk Web UI](glossary.md#snyk-web-ui).

### Workspaces (SCM integrations)

A Snyk feature. This enables Snyk to ingest shallow copies of your Git repositories for scanning, resulting in precise and reliable vulnerability scans.

See [Workspaces for SCM integrations](../../developer-tools/scm-integrations/workspaces.md).
