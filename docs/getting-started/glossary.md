# Glossary

## A

### Advisor

See [Snyk Advisor](https://snyk.io/advisor/).

### **Asset (Snyk AppRisk)**

A Snyk AppRisk asset is an identifiable entity that is part of an application, and relevant for security and developers.

## B

### Base image

The parent image used to construct a container image, usually defined in the `FROM` directive in a Dockerfile. Base images themselves can be constructed from other base images.

### Broker

See [Snyk Broker](../enterprise-configuration/snyk-broker/).

### Build System

A system that takes the source code and builds the deployable application (such as a container).

## C

### CI/CD

Continuous integration (CI), continuous delivery (CD), and continuous deployment (CD) together comprise a Software Development Lifecycle (SDLC) model, guiding developers to automate the development and delivery of small, frequent changes. This ensures all team members have access to the latest codebase and can ensure the compatibility of committed code during development. See [Snyk CI/CD](../integrate-with-snyk/snyk-ci-cd-integrations/) for details of Snyk CI/CD integrations.

### **Class (Snyk AppRisk)**

A way to assign business context to assets and categorize an asset based on the business criticality. Assets can be assigned Classes A, B, C, or D, where Class A (assets that are business critical, deal with sensitive data, are subject to compliance, and so on) is the most important, and Class D (test apps, sandbox environments, and so on) the least important. Assets are assigned Class C by default. A class can be used in policies as well as defined in a policy.

### CLI

Command Line Interface. See [Snyk CLI](glossary.md#snyk-cli).

### Cloud Native Application Security

Implementing security throughout the CI/CD pipeline, automating security embedding in microservices, and maximizing repetition to reduce the introduction of vulnerabilities. Snyk provides a comprehensive [CNAS platform](https://snyk.io/product/cloud-native-application-security/).\
See the article [Cloud-native security guide for building secure applications](https://snyk.io/learn/cloud-native-security-for-cloud-native-applications/).

### Container

Containers allow you to package applications and their dependencies together to be deployed as a single runnable unit. A container is an abstraction provided by the operating system kernel that allows a process to be isolated from other processes running on the system. See also [Snyk Container.](glossary.md#snyk-container)

### Container engine

For users, an application that takes a container image and turns it into a running container. Container engines typically interface with container registries and run containers. Examples of container engines include Docker, CRI-O, and LXC.

### Container image

One or more files that, when instantiated by a container engine or runtime, provide a running container. Images are the packaging and distribution format for containers.

### Container registry

A server that provides a mechanism to store and retrieve container images.

### **Controls (Snyk AppRisk)**&#x20;

The security controls associated with the asset. Navigate to the Snyk AppRisk Controls section to see all available statuses for security controls.

### **Coverage (Snyk AppRisk)**

An assessment of whether applicable assets are scanned and tested by security tools (like Snyk Open Source, for instance), as it relates to an application security program. A type of policy that allows you to specify what controls should be applied and, optionally, how often it needs to be run.

### CVE

Common Vulnerabilities and Exposures. A widely-used identifier for a well-known vulnerability.

### CVSS

Common Vulnerability Scoring System. An industry standard to assess the severity of vulnerabilities, using a score of 0 (lowest) to 10 (highest). Snyk uses CVSS.

### CWE

Common Weakness Enumeration. An online glossary that categorizes software and hardware weaknesses into different types, for example, CWE-20: Input Validation.

## D

### DAST

Dynamic Application Security Testing. An application that you can point at a site or service; it then typically profiles the site or service, then examines the output and behavior to uncover security vulnerabilities. See also [SAST](glossary.md#sast).

### Dependency

When your application uses another package, this other package becomes a dependency in your own software.

* A direct dependency is a package you include in your own Project.
* An indirect dependency (also known as a deep, chained, or transitive dependency), is a package that is used by one of your direct dependencies.

### Dependency tree

Also known as Dependency path. A hierarchical graph showing the dependencies of a software application. This includes both direct and indirect dependencies and thus may be many levels deep.

### DevOps

A set of cultural philosophies, practices, and tools that combines software development and IT operations to shorten the systems development lifecycle.

### DevSecOps

The integration of security into emerging agile IT and DevOps development as seamlessly and as transparently as possible.

### Dockerfile

A text file format used to build container images using Docker. The Dockerfile contains all the commands needed to construct the final image, including specifying the parent base image.

## E

### Environment

Can refer to a cloud environment, a [Project attribute](../snyk-admin/snyk-projects/project-attributes.md), or an interface for working with Snyk, such as the Snyk [CLI](glossary.md#cli), [Web UI](glossary.md#snyk-web-ui), or an [IDE](glossary.md#ide).

### Exploit

A demonstration of how a vulnerability can be taken advantage of. When an exploit is widely published, it is commonly referred to as an exploit "in the wild". See [View exploits](../manage-risk/prioritize-your-issues/view-exploits.md).

### Exploit Maturity

A measure of how practical an exploit for a vulnerability is, based on whether the exploit is in the wild, and how "helpful" the exploit is to attackers.

## F

### Fixable / Partially fixable

A measure of whether a vulnerability can be fixed by Sny by applying a patch, upgrade, or pin. See [Vulnerability fix types](../scan-with-snyk/snyk-open-source/manage-vulnerabilities/vulnerability-fix-types.md).

### Fix PR

A pull request with an automatic fix for discovered vulnerabilities that Snyk can offer the user. See [Automated fix PRs](../scan-with-snyk/snyk-open-source/automatic-and-manual-prs-with-snyk-open-source/automated-fix-pull-requests-for-backlog-issues-and-known-vulnerabilities.md).

## G

### Git

A distributed version control system for tracking changes in source code during software development.

## I

### IaC

Infrastructure as Code. See [Snyk Infrastructure as Code.](glossary.md#snyk-infrastructure-as-code)

### IAST

Interactive Application Security Testing. This approach tests for vulnerabilities while running the application. See [DAST](glossary.md#dast) and [SAST](glossary.md#sast).

### IDE

Integrated Development Environment. An application that has facilities for software development, typically with a source code editor, build automation tools, and a debugger.

### Image

The stored instance of a container that holds a set of software needed to run an application.

### Image layer

Container images typically consist of several different file system layers, which are combined together at runtime into a single file system.

### Integrations

Third-party products, applications, and platforms that Snyk works with, for example, SCM systems such as GitHub. See [Integrate with Snyk](../integrate-with-snyk/).

### Issue

A license problem, vulnerability, or misconfiguration identified and listed by Snyk. See [Find and manage priority issues](../manage-risk/prioritize-your-issues/).

## L

### Library

A specific type of package.

### License policy

A set of criteria for evaluating open-source license issues. License policies enable you to set the severity level and define legal instructions for each license. See [License policies](../scan-with-snyk/policies/license-policies/).

## M

### Manifest

A file containing metadata about other files in a package.

### Monitor

The `snyk monitor` command tests a Project and uploads the results to Snyk. See the CLI help for [Monitor](../snyk-cli/commands/monitor.md).

## O

### OCI

Open Container Initiative. An independent body set up to facilitate collaboration on standards for containers, to ensure they are interoperable between vendor solutions.

### Organization

An Organization in Snyk is a way to collect and organize your Projects. Members of Organizations have access to these Projects. See [Manage Groups and Organizations](../snyk-admin/manage-groups-and-organizations/).

### Origin or source

The identifier for the ecosystem that a Target exists in. Snyk can scan Projects from multiple integrations, including CLI, API, GitHub, Kubernetes, and others. See [Snyk Projects](../snyk-admin/snyk-projects/).

## P

### Package

A group of files and additional metadata about those files, used by package managers.

### Package manager

A set of tools that automate and manage packages of bundled files, and are usually specific to a language. For example, npm.

### Package registry

A software package hosting service that allows customers to host packages and code in one place.

### Pinnable

A fix type: define and "pin" a specific version of an indirect dependency, to avoid a direct dependency pulling in a vulnerable version.

### Policy

See [license policy](glossary.md#license-policy), [security policy](glossary.md#security-policy), and [`.snyk` policy](glossary.md#snyk-policy).

### **Policy (Snyk AppRisk)**

&#x20;A way to automate actions in certain conditions, like classifying and tagging assets with business context. You can also use a policy to configure actions like sending a message or setting the coverage gap control using a Policy builder UI.

### PR

Pull Request. Allows a user to exchange changes made to source code and collaborate with others on the same branch.

### PR Checks

Use Snyk PR Checks to prevent new security issues from entering your codebase by automatically scanning code changes in real-time as soon as you submit a pull request (PR) in your source code manager (SCM). See [Run PR Checks](../scan-with-snyk/run-pr-checks/).

### Priority Score

Snyk scores issues, including vulnerabilities and licenses for Open Source, to help prioritize the treatment of each one. Scores are based on multiple factors including the CVSS score and range from 0 (low) to 1000 (high). See [Priority Score](../manage-risk/prioritize-your-issues/priority-score.md).

### Project

An external item scanned by Snyk with configuration to define how to run that scan. Projects appear on the **Projects** menu on the Snyk dashboard. See also [Target](glossary.md#target). For details, see [Snyk Projects](../snyk-admin/snyk-projects/).

## R

### Reachability

Whether an application contains code that will hit a vulnerable code path during execution. See [Reachable vulnerabilities](../manage-risk/prioritize-your-issues/reachable-vulnerabilities.md).

### Registry

See [Container registry](glossary.md#container-registry) or [Package registry](glossary.md#package-registry).

### Repository

A storage area that contains all elements necessary for the distribution of an application.

### Resource

A cloud infrastructure entity such as an AWS S3 bucket, Identity and Access Management (IAM) role, or Virtual Private Cloud (VPC) flow log.

### Risk score

A value assigned to an issue, ranging from 0 to 1,000, representing the risk imposed on your environment.

### Rule

A security policy that checks cloud infrastructure and infrastructure as code (IaC) for misconfigurations that can lead to security problems, or a security rule used by Snyk Code when scanning your source code for vulnerabilities. For more information, see [Snyk Code security rules](../scan-with-snyk/snyk-code/snyk-code-security-rules/) and [IaC custom rules](../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/).

## S

### SARIF

Static Analysis Results Interchange Format. A standard, JSON-based format for the output of static analysis tools.

### SAST

Static Application Security Testing. A method to secure software by reviewing the source code of your proprietary software and identifying sources of vulnerabilities. See also [DAST](glossary.md#dast).

### SBOM

Software Bill Of Materials. A list of components in a piece of software.

### SCA

Software Composition Analysis. A technology that is used to identify open-source and third-party components in use in an application, including their known security vulnerabilities, and typically adversarial license restrictions. See also [Static Code Analysis](glossary.md#static-code-analysis).

### SCM

Source Code Management. Also known as a code repository (repo) or version control system. The method used by developers to store their source code and track changes to code. SCM helps resolve conflicts when merging updates from multiple contributors. GitHub is an example of a common SCM system. See [Git repositories (SCMs)](../integrate-with-snyk/git-repositories-scms-integrations-with-snyk/).

### SDLC

Software Development Lifecycle. A process followed by a development team, describing how to develop and maintain software.

### Security policy

A set of criteria for evaluating open-source vulnerabilities. Security policies enable you to set custom rules to automatically prioritize or de-prioritize specific vulnerabilities. See [Security policies](../scan-with-snyk/policies/security-policies/).

### Severity

A severity level is applied to a vulnerability or a license issue, to indicate the risk for that item in an application. See [Severity levels](../manage-risk/prioritize-your-issues/severity-levels.md).

### Snapshot

An individual report within the test history of a Project. Includes a tree of dependencies and a list of vulnerabilities that was accurate at the time the test was conducted.

### `.snyk` policy

A policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins. See [The .snyk file](../manage-risk/prioritize-your-issues/the-.snyk-file.md).

### Snyk

A platform providing Cloud Native Application Security (CNAS) solutions, allowing developers to own and build security for the whole application, from code and open source to containers and cloud infrastructure. Snyk is also the company providing the Snyk platform. See [Getting started](./).

### Snyk Advisor

A free web application that allows you to compare software packages across open-source ecosystems. It provides insights into the overall health of a particular package by combining community and security data into a single unified view. See [Snyk Advisor](https://snyk.io/advisor/).

### Snyk API

A Snyk tool that enables developers to integrate programmatically with Snyk. See [Snyk API](../snyk-api/).

### Snyk Apps

Snyk Apps are the modern and preferred way to build integrations with Snyk, exposing fine-grained scopes for accessing resources over the Snyk APIs, powered by OAuth 2.0 for a developer-friendly experience. See [Snyk Apps](../snyk-api-info/snyk-apps/).

### Snyk Broker

A client/server system that serves as an agent or proxy, allowing Snyk to scan private customer environments: Jira, code repositories, or container registries. Snyk Broker relays messages and allows users to filter which messages are allowed through,  for example, allowing users to expose only some GitHub APIs to Snyk. See [Snyk Broker](../enterprise-configuration/snyk-broker/).

### Snyk CLI

A Snyk platform tool that enables developers to find and fix known vulnerabilities in dependencies, using a command line interface. See [Snyk CLI](../snyk-cli/).

### Snyk Code

A Snyk product. A SAST product enabling developers to find and fix vulnerabilities in your proprietary application code. See [Snyk Code](../scan-with-snyk/snyk-code/).

### Snyk Container

A Snyk product. Enables developers to find and fix vulnerabilities in container images and Kubernetes applications. See [Snyk Container](../scan-with-snyk/snyk-container/).

### Snyk Infrastructure as Code

A Snyk product. Enables developers to find and fix vulnerabilities in Kubernetes, Helm, and Terraform configuration files. See [Snyk Infrastructure as Code](../scan-with-snyk/snyk-iac/scan-your-iac-source-code/).

### Snyk Open Source

A Snyk product. Enables developers to find and fix open-source vulnerabilities. See [Snyk Open Source](../scan-with-snyk/snyk-open-source/).

### Snyk plugin

A library used by the Snyk CLI to scan a certain language or build system.

### Snyk Security Intelligence

A component powering the Snk cloud-native application security platform.\
Incorporates the Snyk Intel Vulnerability DB: the Snyk database of vulnerabilities, providing detailed information and fix advice for known vulnerabilities. See [Vulnerability DB](https://snyk.io/vuln).

### Snyk Web UI

The browser-based environment that provides users access to Snyk functions. See [Explore the Snyk Web UI](explore-snyk-through-the-web-ui.md).

### Social Trends

Snyk shows a Trending banner on issues that are being actively discussed on Twitter. See  [Vulnerabilities with Social Trends](../manage-risk/prioritize-your-issues/vulnerabilities-with-social-trends.md).

### Source

See [Origin](glossary.md#origin-or-source).

### SPDX

Software Package Data Exchange. A file format used to document information on the software licenses under which a piece of computer software is distributed. See [SPDX](https://spdx.dev/).

### Static Code Analysis

A method of debugging by examining source code before a program is run. See also [SCA, Software Composition Analysis](glossary.md#sca).

## T

### Target

Representation of an external resource Snyk has scanned. All [Snyk Projects](glossary.md#project) are associated with a parent Target. One Target may relate to many Projects. The structure of the Target depends on the [origin](glossary.md#origin-or-source).

### **Tags (Snyk AppRisk)**

A way to categorize assets. Helps you recognize or handle assets differently according to mutual properties. Assets can be filtered by their tags in the inventory or when creating policy rules. A tag can be automatically assigned to an asset, or the asset can be tagged by a policy you created. GitHub and GitLab topics are treated as asset tags and you can use them for creating policies.

## U

### Upgradable / Patchable

A fix type: a problem can be fixed by upgrading a version of a package or by applying a patch.

## V

### Vulnerability

A security vulnerability that was identified by Snyk. See [Manage vulnerabilities](../scan-with-snyk/snyk-open-source/manage-vulnerabilities/).

## W

### Webhook

A way for an app to provide other applications with real-time information. Snyk uses webhooks to check changes in code. See [Snyk Webhooks](../snyk-api-info/snyk-webhooks/).
