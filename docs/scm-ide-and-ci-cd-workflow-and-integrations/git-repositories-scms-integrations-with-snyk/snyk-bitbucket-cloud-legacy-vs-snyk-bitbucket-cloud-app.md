# Snyk Bitbucket Cloud (legacy) vs Snyk Bitbucket Cloud App

In general, Snyk recommends using the new Bitbucket app integration. However, the new integration does not fit all cases. The information in this section is intended to help you decide which integration is right for you.

See [Migrate to the Snyk Bitbucket Cloud App](snyk-bitbucket-cloud-integration.md#migrate-to-the-snyk-bitbucket-cloud-app) for detailed migration instructions.

### Main capabilities unlocked by the new app integration

* Allows using Snyk with Bitbucket's [allowlisting IP addresses](https://support.atlassian.com/bitbucket-cloud/docs/control-access-to-your-private-content/) premium tier feature.
* Helps handle rate-limiting issues for companies who spread their repos across multiple workspaces in Bitbucket Cloud.
* Supports the first-party interface in Bitbucket Cloud (Snyk's Security tab) out-of-the-box, meaning you need not install and maintain the first-party extension app to consume Snyk's security insights from Bitbucket Cloud.

### Limitations of the new app integration

* In the new app integration, every Snyk Organization can connect to only one workspace in Bitbucket Cloud. If you want to import Projects from various workspaces in Bitbucket into the same single Organization in Snyk, use the PAT integration.
* The new app integration does not yet support Snyk Multi-Tenant EU, Snyk Multi-Tenant AUS, and Snyk Single-Tenant cloud deployments.
* For customers who are part of the custom branch closed beta, the new app integration's first-party interface in Bitbucket Cloud does not allow importing Projects from non-default branches. It is possible to import a non-default branch; you must do it from the Snyk.io import modal.

### Are there any plans for end-of-life for the Personal Access Token (PAT) integration?

No, the Personal Access Token Bitbucket Cloud integration is fully supported, and there are no plans to stop supporting it.

However, there is a first-party interface _extension_ app that serves as an extension layer to the PAT integration, allowing developers to consume Snyk's findings from within the Bitbucket interface. This extension app was developed and supported by an external contractor company. As this functionality is now an integral part of the new app integration, the extension app has now moved to no-support mode, meaning that customers who use the PAT integration alongside the first-party extension app must migrate to the new app integration to get support for the first-party interface functionality.
