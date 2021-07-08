# Snyk Glossary

[A]() \| [B]() \| [C]() \| [D]() \| [F]() \| [G]() \| [I]() \| [L]() \| [M]() \| [O]() \| [P]() \| [R]() \| [S]() \| [V]() \| [W]()

## API

Application Programming Interface. See [Snyk API]().

## Broker

See [Snyk Broker]().

## Build System

A system that takes the source code and builds the deployable application \(such as a container\).

## CD

Continuous Deployment or Continuous Delivery. Practices that focus on packaging and delivering software, with a fully automated deployment process.

* **Continuous Deployment**: a practice to automate the release of code to a production environment, using automated testing to validate if changes to a codebase are correct and stable.
* **Continuous Delivery**: the frequent shipping of code to a given environment via manual release.

## CI

Continuous Integration. A practice to ensure new code changes are regularly built, tested, and merged to a shared repository.

## CI / CD

Continuous integration \(CI\), continuous delivery \(CD\) and continuous deployment \(CD\) together comprise an SDLC model, guiding developers to automate development and delivery of small, frequent changes. This ensures all team members have access to the latest code base and can ensure the compatibility of committed code during development.

## CLI

Command Line Interface. See [Snyk CLI]().

## Cloud Native Application Security

Implementing security throughout the CI/CD pipeline, automating security embedding in microservices and maximizing repetition to reduce the introduction of vulnerabilities. Snyk provides a comprehensive [CNAS platform](https://snyk.io/product/cloud-native-application-security/).

## Container

A standard unit of software, in runtime, that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another. Containers isolate software from its environment, so the software works uniformly despite differences, for instance between development and staging. Also see [Snyk Container.]()

## Container engine

The platform technology through which containers are built and run. Examples include Docker Engine and Kubernetes.

## Container image

A lightweight, standalone, executable package of software that includes everything need to run an application: code, runtime library, system tools, system libraries, and settings. Container images become containers at runtime.

## Container registry

A repository storing container images created during the application development process, which can then be used throughout the development lifecycle.

## CVE

Common Vulnerabilities and Exposures. A widely-used identifier for a well-known vulnerability.

## CVSS

Common Vulnerability Scoring System. An industry standard to assess the severity of vulnerabilities, using a score of 0 \(lowest\) to 10 \(highest\). Snyk uses CVSS.

## CWE

Common Weakness Enumeration, an online glossary that categorizes software and hardware weaknesses into different types. For example: **CWE-20: Input Validation**.

## DAST

Dynamic Application Security Testing. An application that you can point at a site or service; it then typically profiles the site or service, then examines the output and behaviour to uncover security vulnerabilities.

## Dependency

When your application uses another package, this other package becomes a dependency to your own software.

* A direct dependency is a package you include in your own project.
* An indirect dependency \(also known as a deep, chained, or transitive dependency\), is a package that is used by one of your direct dependencies.

## DevOps

A set of cultural philosophies, practices, and tools that combines software development and IT operations, to shorten the systems development life cycle.

## DevSecOps

The integration of security into emerging agile IT and DevOps development as seamlessly and as transparently as possible.

## Fix PR

A pull request with an automatic fix for vulnerabilities found that Snyk can offer the user.

## Git

A distributed version-control system for tracking changes in source code during software development.

## GitHub

A web-based version control platform for Git.

## IAC

Infrastructure as Code. See [Snyk Infrastructure as Code.]()

## IDE

Integrated Development Environment. An application giving facilities for software development, typically with a source code editor, build automation tools and a debugger.

## Image

The stored instance of a container that holds a set of software needed to run an application.

## Integrations

Third-party products, applications and platforms that Snyk works with, for example SCM systems such as GitHub.

## Issue

A license problem or vulnerability identified and lists by Snyk.

## Library

A specific type of a package.

## Manifest

A file containing metadata about other files in a package.

## Monitor

A run of the **snyk monitor** command that tests the project and uploads results to Snyk.

## OCI

Open Container Initiative. An independent body set up to facilitate collaboration around standards for containers, to ensure they are interoperable between vendor solutions.

## Organization

An organization in Snyk is a way to collect and organize your projects. Members of organizations can then access these projects.

## Open source

Software with a license that complies with OSI \(Open Source Initiative\) criteria.

## PaaS

Platform as a Service. A platform allowing customers to develop, run, and manage applications without building and maintaining the infrastructure themselves. Can also be referred to as \*\*\*\*application platform as a service \(aPaaS\).

## Package

A group of files and additional metadata about those files, used by package managers.

## Package manager

A set of tools that automates and manages packages of bundled files, and are usually specific to a language. For example, npm.

## Package registry

A software package hosting service that allows customers to host packages and code in one place.

## PR

Pull request. Allows a user to exchange changes made to source code and collaborate with others on the same branch.

## Project

An external item that Snyk scans, with configuration to define how to run that scan. Projects appear on the **Projects** menu on the Snyk dashboard. See [Introduction to projects](https://support.snyk.io/hc/en-us/articles/360019058297-Introduction-to-projects).

## Registry

See [Container registry]() or [Package registry]().

## Reports

Snyk produces a wide range of reports to allow you to review and fix vulnerabilities.

## Repository

A storage area that contains all elements necessary for the distribution of an application.

## SAST

Static Application Security Testing. A method to secure software by reviewing the source code of your proprietary software, and identifying sources of vulnerabilities.

## SCA

Software Composition Analysis. Technology used to identify open-source and third-party components in use in an application, including their known security vulnerabilities, and typically adversarial license restrictions.

## SCM

Source Code Management. Also known as a code repo / repository / version control system. The method used by developers to store their source code, and track changes to code. SCM helps resolve conflicts when merging updates from multiple contributors. GitHub is an example of a common SCM system.

## Score

Snyk scores issues \(vulnerabilities and licenses\), to help prioritze treatment of each one. Scores are based on multiple factors including as the CVSS score, and range from 0 to 1000.

## SDLC

Software Development Life Cycle. A process followed by a development team, describing how to develop and, maintain software.

## Serverless

A situation when a cloud provider dynamically manages the allocation of machine resources, rather than the developer doing it.

## Snapshot

An individual report within a project’s test history. Includes a tree of dependancies, and a list of vulnerabilities that was accurate at the time the test was conducted.

## Snyk

A platform providing Cloud Native Application Security \(CNAS\) solutions, allowing developers to own and build security for the whole application, from code and open source to containers and cloud infrastructure.

Also, the company providing the Snyk platform.

## Snyk API

A Snyk tool, Enables developers to programatically integrate with Snyk. See [Snyk API documentation](https://support.snyk.io/hc/en-us/categories/360000665657-Snyk-API).

## Snyk Broker

A client/server system that serves as an agent / proxy, allowing Snyk to scan private customer environments \(Jira, code repositories or container registries\). It relays messages and allows users to filter which messages are allowed through; for example, allowing users to expose only some Github APIs to Snyk. See [Snyk Broker documentation](https://support.snyk.io/hc/en-us/sections/360001138138-Snyk-Broker).

## Snyk CLI

A Snyk platform tool. Snyk CLI enables developers to find and fix known vulnerabilities in dependencies, using a command line interface. See [Snyk CLI documentation](https://support.snyk.io/hc/en-us/categories/360000456217-Snyk-CLI).

## Snyk Container

A Snyk product. Enables developers to find and fix vulnerabilities in container images and Kubernetes applications. See [Snyk Container documentation](https://support.snyk.io/hc/en-us/categories/360000583498-Snyk-Container).

## Snyk Infrastructure as Code

A Snyk product. Enables developers to find and fix vulnerabilities in your Kubernetes, Helm and Terraform configuration files. See [Snyk IaC documentation](https://support.snyk.io/hc/en-us/categories/360001342678-Infrastructure-as-code).

## Snyk License Compliance Management

Part of Snyk Open Source, used to identify, monitor and manage open source license usage across your projects.

## Snyk Open Source

A Snyk product. Enables developers to find and fix open source vulnerabilities. See [Snyk Open Source documentation](https://support.snyk.io/hc/en-us/categories/360003049458-Snyk-Open-Source).

## Snyk plugin

A library used by the Snyk CLI to scan a certain language/build system.

## Webhook

A way for an app to provide other applications with real-time information. Snyk uses webhooks to check changes in code.

