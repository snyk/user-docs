# Glossary

## A

### Advisor

See [Snyk Advisor](https://snyk.io/advisor/).

## B

### Base image

The parent image used to construct a container image, usually defined in the FROM directive in a Dockerfile. Base images themselves can be constructed from other base images.

### Broker

See [Snyk Broker](https://docs.snyk.io/integrations/snyk-broker).

### Build System

A system that takes the source code and builds the deployable application (such as a container).

## C

### CI / CD

Continuous integration (CI), continuous delivery (CD) and continuous deployment (CD) together comprise an SDLC model, guiding developers to automate development and delivery of small, frequent changes. This ensures all team members have access to the latest code base and can ensure the compatibility of committed code during development.

### CLI

Command Line Interface. See [Snyk CLI](glossary.md#snyk-cli).

### Cloud Native Application Security

Implementing security throughout the CI/CD pipeline, automating security embedding in microservices and maximizing repetition to reduce the introduction of vulnerabilities. Snyk provides a comprehensive [CNAS platform](https://snyk.io/product/cloud-native-application-security/).\
See our article about [Cloud native security guide for building secure applications](https://snyk.io/learn/cloud-native-security-for-cloud-native-applications/).

### Container

Containers allow you to package applications and their dependencies together, to be deployed as a single runnable unit. A container is an abstraction provided by the operating system kernel, that allows a process to be isolated from other processes running on the system. Also see [Snyk Container.](glossary.md#snyk-container)

### Container engine

For users, an application which takes a container image and turns it into a running container. Container engines typically interface with container registries, and run containers. Examples of container engines include Docker, CRI-O or LXC.

### Container image

One or more files which, when instantiated by a container engine or runtime, provide a running container. Images are the packaging and distribution format for containers.

### Container registry

A server which provides a mechanism to store and retrieve container images.

### CVE

Common Vulnerabilities and Exposures. A widely-used identifier for a well-known vulnerability.

### CVSS

Common Vulnerability Scoring System. An industry standard to assess the severity of vulnerabilities, using a score of 0 (lowest) to 10 (highest). Snyk uses CVSS.

### CWE

Common Weakness Enumeration, an online glossary that categorizes software and hardware weaknesses into different types. For example: **CWE-20: Input Validation**.

## D

### DAST

Dynamic Application Security Testing. An application that you can point at a site or service; it then typically profiles the site or service, then examines the output and behaviour to uncover security vulnerabilities. Also see [SAST](glossary.md#sast).

### Dependency

When your application uses another package, this other package becomes a dependency to your own software.

* A direct dependency is a package you include in your own project.
* An indirect dependency (also known as a deep, chained, or transitive dependency), is a package that is used by one of your direct dependencies.

### Dependency tree

(Also known as Dependency path) A hierarchical graph showing the dependencies of a software application. This includes both direct and indirect dependencies, and so may be many levels deep.

### DevOps

A set of cultural philosophies, practices, and tools that combines software development and IT operations, to shorten the systems development life cycle.

### DevSecOps

The integration of security into emerging agile IT and DevOps development as seamlessly and as transparently as possible.

### Dockerfile

A text file format used to build container images using Docker. The Dockerfile contains all the commands needed to construct the final image, including specifying the parent base image.

## E

### Environment

Can refer to a [Snyk Cloud Environment](glossary.md#snyk-cloud-environment), a [project attribute](../../snyk-web-ui/introduction-to-snyk-projects/project-attributes.md), or an interface for working with Snyk, such as the Snyk [CLI](glossary.md#cli), [Web UI](glossary.md#snyk-web-ui), or an [IDE](glossary.md#ide).

### Exploit

A demonstration of how a vulnerability can be taken advantage of. When an exploit is widely published, it is commonly referred to as an exploit in the wild.

### Exploit Maturity

A measure of how practical an exploit for a vulnerability is, based on whether the exploit is in the wild, and how "helpful" the exploit is to attackers. See [Evaluating and prioritizing vulnerabilities](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/evaluating-and-prioritizing-vulnerabilities).

## F

### Fixable / Partially fixable

A measure of whether a vulnerability can be fixed by Snyk, by applying a patch, upgrade, or pin. See [Fixed in version vs. fixable attributes in vulnerabilities](https://support.snyk.io/hc/en-us/articles/4405034808209).

### Fix PR

A pull request with an automatic fix for vulnerabilities found that Snyk can offer the user.

## G

### Git

A distributed version-control system for tracking changes in source code during software development.

## I

### IAC

Infrastructure as Code. See [Snyk Infrastructure as Code.](glossary.md#snyk-infrastructure-as-code)

### IAST

Interactive Application Security Testing. This approach tests for vulnerabilities while running the application. See [DAST](glossary.md#dast) and [SAST](glossary.md#sast).

### IDE

Integrated Development Environment. An application giving facilities for software development, typically with a source code editor, build automation tools and a debugger.

### Image

The stored instance of a container that holds a set of software needed to run an application.

### Image layer

Container images typically consist of several different filesystem layers, which are combined together at runtime into a single filesystem.

### Integrations

Third-party products, applications and platforms that Snyk works with, for example SCM systems such as GitHub.

### Issue

A license problem, vulnerability, or misconfiguration identified and listed by Snyk.

## L

### Library

A specific type of a package.

### License policy

A set of criteria for evaluating open source license issues. License policies enable you to set the severity level and define legal instructions for each license. The Snyk Default License Policy is enabled by default, or you can [create your own license policy](https://docs.snyk.io/products/snyk-open-source/license-policies/setting-a-license-policy). See [License policies](https://docs.snyk.io/products/snyk-open-source/license-policies).

## M

### Manifest

A file containing metadata about other files in a package.

### Monitor

A run of the **snyk monitor** command that tests the project and uploads results to Snyk.

## O

### OCI

Open Container Initiative. An independent body set up to facilitate collaboration around standards for containers, to ensure they are interoperable between vendor solutions.

### Organization

An organization in Snyk is a way to collect and organize your projects. Members of organizations can then access these projects.

### Origin

The identifier for the ecosystem that a Target exists in. Snyk can scan projects from multiple integrations, including CLI, API, GitHub, Kubernetes and others. See [Introduction to projects](https://docs.snyk.io/getting-started/introduction-to-snyk-projects).

## P

### Package

A group of files and additional metadata about those files, used by package managers.

### Package manager

A set of tools that automates and manages packages of bundled files, and are usually specific to a language. For example, npm.

### Package registry

A software package hosting service that allows customers to host packages and code in one place.

### Pinnable

A fix type: define and "pin" a specific version of an indirect dependency, to avoid a direct dependency pulling in a vulnerable version.

### Policy

See [license policy](glossary.md#license-policy), [security policy](glossary.md#security-policy), and [`.snyk` policy](glossary.md#.snyk-policy).

### PR

Pull Request. Allows a user to exchange changes made to source code and collaborate with others on the same branch.

### Priority Score

Snyk scores issues (vulnerabilities and licenses), to help prioritze treatment of each one. Scores are based on multiple factors including as the CVSS score, and range from 0 (low) to 1000 (high). See [Snyk Priority Score](https://docs.snyk.io/fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/snyk-priority-score).

### Project

An external item that Snyk scans, with configuration to define how to run that scan. Projects appear on the **Projects** menu on the Snyk dashboard. See [Introduction to projects](https://docs.snyk.io/getting-started/introduction-to-snyk-projects).

## R

### Reachability

Whether an application contains code which will hit a vulnerable code path during execution. See [Reachable vulnerabilities](https://support.snyk.io/hc/en-us/articles/360010554837-Reachable-Vulnerabilities-).

### Registry

See [Container registry](https://support.snyk.io/hc/en-us/articles/360017682058-Snyk-Glossary#ContainerRegistry) or [Package registry](https://support.snyk.io/hc/en-us/articles/360017682058-Snyk-Glossary#PackageRegistry).

### Repository

A storage area that contains all elements necessary for the distribution of an application.

### Resource

A cloud infrastructure entity such as an AWS S3 bucket, Identity & Access Management (IAM) role, or Virtual Private Cloud (VPC) flow log.

### Rule

A security policy that checks cloud infrastructure and infrastructure as code (IaC) for misconfigurations that can lead to security problems.

## S

### SARIF

Static Analysis Results Interchange Format. A standard, JSON-based format for the output of static analysis tools.

### SAST

Static Application Security Testing. A method to secure software by reviewing the source code of your proprietary software, and identifying sources of vulnerabilities. Also see [DAST](glossary.md#dast).

### SBOM

Software Bill Of Materials. A list of components in a piece of software.

### SCA

Software Composition Analysis. Technology used to identify open-source and third-party components in use in an application, including their known security vulnerabilities, and typically adversarial license restrictions.

{% hint style="info" %}
Not to be confused with static code analysis (a method of debugging by examining source code before a program is run).
{% endhint %}

### SCM

Source Code Management. Also known as a code repo / repository / version control system. The method used by developers to store their source code, and track changes to code. SCM helps resolve conflicts when merging updates from multiple contributors. GitHub is an example of a common SCM system.

### SDLC

Software Development Life Cycle. A process followed by a development team, describing how to develop and, maintain software.

### Security policy

A set of criteria for evaluating open source vulnerabilities. Security policies enable you to set custom rules to automatically prioritize or de-prioritize specific vulnerabilities. The Snyk Default Security Policy is enabled by default, or you can [create your own security policy](https://docs.snyk.io/features/fixing-and-prioritizing-issues/security-policies/how-to-create-a-security-policy-and-set-rules). See [Security policies](https://docs.snyk.io/features/fixing-and-prioritizing-issues/security-policies).

### Serverless

A method to provide pay-as-you-use backend computing services, allowing applications to be constructed entirely from functions provided by the supplier. Examples of serverless providers include AWS Lambda and Azure Functions.

### Severity

A severity level is applied to a vulnerability or a license issue, to indicate the risk for that item in an application. See [Severity levels](https://docs.snyk.io/introducing-snyk/snyks-core-concepts/severity-levels).

### Snapshot

An individual report within a project’s test history. Includes a tree of dependancies, and a list of vulnerabilities that was accurate at the time the test was conducted.

### `.snyk` policy

A policy file that Snyk uses to define certain analysis behaviors and to specify patches for the CLI and CI/CD plugins. See [The .snyk file](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/the-.snyk-file).

### Snyk

A platform providing Cloud Native Application Security (CNAS) solutions, allowing developers to own and build security for the whole application, from code and open source to containers and cloud infrastructure. See [What is Snyk?](https://snyk.io/what-is-snyk/).

Snyk is also the company providing the Snyk platform.

### Snyk Advisor

A free web application which allows you to compare software packages across open source ecosystems. It provides insights into the overall health of a particular package by combining community and security data into a single unified view. See [Snyk Advisor](https://snyk.io/advisor/).

### Snyk API

A Snyk tool that enables developers to programmatically integrate with Snyk. See [Snyk API](../../snyk-api-info/).

### Snyk Apps

Snyk Apps are the modern and preferred way to build integrations with Snyk, exposing fine-grained scopes for accessing resources over the Snyk APIs, powered by OAuth 2.0 for a developer-friendly experience. For more information see [Snyk Apps](https://docs.snyk.io/integrations/snyk-apps).

### Snyk Broker

A client/server system that serves as an agent / proxy, allowing Snyk to scan private customer environments (Jira, code repositories or container registries). It relays messages and allows users to filter which messages are allowed through; for example, allowing users to expose only some Github APIs to Snyk. See [Snyk Broker documentation](https://docs.snyk.io/integrations/snyk-broker).

### Snyk CLI

A Snyk platform tool. Snyk CLI enables developers to find and fix known vulnerabilities in dependencies, using a command line interface. See [Snyk CLI documentation](https://docs.snyk.io/snyk-cli).

### Snyk Cloud

A Snyk product. Enables developers to find and fix cloud infrastructure misconfigurations. See [Snyk Cloud documentation](../../products/snyk-cloud/).

### Snyk Cloud Environment

An organizing concept that equates to an Amazon Web Services (AWS) account, Azure subscription, or Google Cloud project. See [Snyk Cloud concepts](../../products/snyk-cloud/snyk-cloud-concepts.md).

### Snyk Code

A Snyk product. A SAST product enabling developers to find and fix vulnerabilities in your proprietary application code. See [Snyk Code documentation](https://docs.snyk.io/snyk-code).

### Snyk Container

A Snyk product. Enables developers to find and fix vulnerabilities in container images and Kubernetes applications. See [Snyk Container documentation](https://docs.snyk.io/snyk-container).

### Snyk Infrastructure as Code

A Snyk product. Enables developers to find and fix vulnerabilities in your Kubernetes, Helm and Terraform configuration files. See [Snyk IaC documentation](https://docs.snyk.io/snyk-infrastructure-as-code).

### Snyk Open Source

A Snyk product. Enables developers to find and fix open source vulnerabilities. See [Snyk Open Source documentation](https://docs.snyk.io/snyk-open-source).

### Snyk plugin

A library used by the Snyk CLI to scan a certain language/build system.

### Snyk Security Intelligence

A component powering Snyk’s cloud native application security platform.\
Incorporates **Snyk Intel Vulnerability DB**: Snyk’s database of vulnerabilities, providing detailed information and fix advice for known vulnerabilities. See [Vulnerability DB](https://snyk.io/vuln).

### Snyk Web UI

The environment allowing users to access Snyk functions on the web. See [Snyk Web UI](../../snyk-web-ui/).

### Social Trends

Snyk shows a Trending banner on issues that are being actively discussed in Twitter. See the [Prioritizing Social Trends](https://docs.snyk.io/fixing-and-prioritizing-issues/prioritizing-issues/prioritize-by-social-trends) article.

### SPDX

Software Package Data Exchange. A file format used to document information on the software licenses under which a piece of computer software is distributed.

## T

### Target

Representation of an external resource Snyk has scanned through an integration, the CLI, UI or API. Targets may represent a SCM repository, a Kubernetes workload, or other scannable resources external to Snyk. All [projects](glossary.md#project) that Snyk create are associated to a parent target. One target may relate to many projects. The structure of the target depends on the [origin](glossary.md#origin). See [Introduction to projects](https://docs.snyk.io/getting-started/introduction-to-snyk-projects).

## U

### Upgradable / Patchable

A fix type: a problem can be fixed by upgrading a version of a package, or by applying a patch.

## V

### Vulnerability

A security vulnerability, identified by Snyk. See [Fixing vulnerabilities](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/remediate-your-vulnerabilities).

## W

### Webhook

A way for an app to provide other applications with real-time information. Snyk uses webhooks to check changes in code.
