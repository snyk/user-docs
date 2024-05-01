# Level 3: Prioritization setup

Learn what Risk-Based prioritization is and how to set up and deploy the Kubernetes connector for Snyk AppRisk.

Risk-based prioritization is the capability of Snyk AppRisk to understand the context of your application and help you better prioritize your security issues.

The [Snyk Risk-Based prioritization](../risk-based-prioritization-for-snyk-apprisk/prioritization-setup/) product focuses on several risk factors for your vulnerabilities:

* **OS condition**: Does this vulnerability apply to my operating system?
* **Deployed**: Is my code in a container image that is deployed?
* **Public facing**: Does my container image have a configured path to the internet?

The goal of Snyk Risk-based prioritization is to provide application context to your Open Source, Code, and Container issues by understanding how your application is deployed and configured. This enables you to prioritize your issues based on the risk they are posing to your application.

See the [How risk-based prioritization works](../risk-based-prioritization-for-snyk-apprisk/how-risk-based-prioritization-works/) page, with focus on the [Assets](../risk-based-prioritization-for-snyk-apprisk/how-risk-based-prioritization-works/assets.md) and [Risk factors](../risk-based-prioritization-for-snyk-apprisk/how-risk-based-prioritization-works/risk-factors/) pages, for more details and a better understanding of the core concepts.

The purpose of this level is to provide context to your Snyk Container issues or vulnerabilities. After you deploy the kubernetes connector to your cluster, you will be able to identify if a Container is Deployed and Public Facing, allowing you to prioritize your container vulnerabilities.

See the [Prioritization setup: Kubernetes connector](../risk-based-prioritization-for-snyk-apprisk/prioritization-setup/prioritization-setup-kubernetes-connector.md) page for more details about how to set up the Kubernetes connector.
