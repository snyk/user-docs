# Set up Snyk AppRisk Insights

## Prioritize issues with risk factors

Learn about risk-based prioritization and how to set up and deploy the Kubernetes connector for Snyk AppRisk, using the [**Issues**](../../../discover-snyk/getting-started/snyk-web-ui.md#view-and-prioritize-issues) page from the Snyk AppRisk menu.

Risk-based prioritization, or insights, is the capability of Snyk AppRisk to understand the context of your application and help you better prioritize your security issues.

The [Snyk risk-based prioritization](../../../manage-risk/prioritize-issues-for-fixing/#prioritization-based-on-risk) product focuses on several risk factors for your vulnerabilities:

* **Deployed**: Is my code in a container image that is deployed?
* **Public facing**: Does my container image have a configured path to the internet?
* **Loaded package**:  Has a third-party package that is the dependency of an image been loaded?

The goal of Snyk risk-based prioritization is to provide application context to your Open Source, Code, and Container issues by understanding how your application is deployed and configured. This enables you to prioritize your issues based on the risk they pose to your application.

See the [How risk-based prioritization works](../../../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/) page, with focus on the [Assets](../../../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/#assets) and [Risk factors](../../../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/#risk-factors) pages, for more details and a better understanding of the core concepts.

The purpose of this level is to provide context to your Snyk Container issues or vulnerabilities. After you deploy the Kubernetes connector to your cluster, you will be able to identify if a Container is Deployed and Public Facing, allowing you to prioritize your container vulnerabilities.

See the [Prioritization setup: Kubernetes connector](../../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-kubernetes-connector.md) page for more details about how to set up the Kubernetes connector.

## Associate the source code with the container images

Leverage the complete capabilities of risk-based prioritization by associating your Snyk Code and Snyk Open Source Projects with your Container Project. This allows you to prioritize your Snyk Open Source and Snyk Code issues using contextual risk factors.

See the [Prioritization setup: Associating Snyk Open Source, Code, and Container Projects](../../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-associating-snyk-open-source-code-and-container-projects.md) page for more details on how to set up the required linking for the chosen application.
