# SCM (Git) and CI/CD integration deployment intro

The most popular integrations for Snyk Code and Open Source are Source Code Management (or Git) and CI/CD pipeline Integration.

![Snyk Integrations](../../.gitbook/assets/scm-ci-cid.png)

Snyk SCM (Git) Integration checks for vulnerabilities and Open Source license issues, prevents pull requests based on policies, and allows you to block newly introduced vulnerabilities. SCM Integration also allows you to rescan automatically on every fix PR; this is enabled by Snyk webhooks.

CI/CD Integration adds Snyk in the build as a step of the pipeline.

Either SCM (Git) or CI/CD integration may achieve the same business objective. Why choose one over the other? The answer is specific to every organization based on workflows and processes. The following considerations explain the benefits of each type of integration.

## SCM Integration considerations

* Easier to setup and maintain
* Ability to block fix PR enabled by webhooks
* Allows scanning earlier in the software development lifecycle
* More friendly experience for developers
* Fix PR from Snyk (SCA and Dockerfiles)
* Does not take resources from your CI/CD pipeline

### CI/CD Integration considerations

* Some package managers require local context and are better run within your environment (notably Scala, Gradle and Go modules)
* More granular options to block
* Strong gatekeeper
* Best practice for container scans
* Allows more custom reports for a complex repo structure such as large monorepo
