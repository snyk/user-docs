# Assets and risk factors for Snyk AppRisk

The capabilities of the SnykWeb UI Issues menu rely on understanding your application context to help you better prioritize your security issues. It does that by understanding how your application is configured and relying on that information to provide you with triage and prioritization of your assets and issues for the Snyk Essentials plan, and it also adds specific [risk factors](./#risk-factors) and [evidence graph](../using-the-issues-ui-with-snyk-apprisk/evidence-graph.md) information for the Snyk AppRisk plan.&#x20;

* [Assets](./#assets) are analyzed using Snyk Issues, focusing on images, Kubernetes resources, and packages to understand how they all interact with each other.
* [Risk factors](./#risk-factors) are analyzed using Snyk Issues and grouped into four main categories:&#x20;
  * [Deployed](risk-factor-deployed.md)&#x20;
  * [Loaded package ](risk-factor-loaded-package.md)
  * [OS conditions](risk-factor-os-condition.md)&#x20;
  * [Public facing](risk-factor-public-facing.md)

## Assets

Snyk AppRisk issues analyzes the assets described on this page.

### Images

Images are assets that represent a Docker image. Snyk Container performs security scans on the Docker images. Images can be mapped one to many with the Snyk Projects created by the scans performed by Snyk Container. Docker images have natural IDs, which are represented by SHAs. Snyk uses this natural ID to correlate the same images even if they are mapped to different Snyk Projects.

### Kubernetes resources

Kubernetes resources are assets that represent a Kubernetes object. The Kubernetes Connector collects resource information from the Kubernetes clusters.&#x20;

Kubernetes resources do not map to the Snyk Projects. These are internal entities used to compute certain risk factors, further detailed on the rest of this page. These risk factors can be related to the packages and images.

### Packages

Packages are assets that represent a software package. Snyk Open Source and Snyk Code products perform security scans on files. These files represent the package manager declaration and the source code of a software package, respectively. A package is a representation of that software package.

Packages can be mapped one to one with the Snyk Projects created by the scans performed by Snyk Open Source and Snyk Code. All the issues identified by these products and attributed to these Projects will be mapped to the package entity.&#x20;

The term Package is a very coarse abstraction. It does not have versions. It is a representation of the current state of the software package at a point in time. The point in time is determined by the time when the Snyk processing pipeline is completed and the state of Snyk Projects at that time.&#x20;

Snyk Open Source uses the word package to refer to the third-party dependencies declared in the package dependency manifest. Snyk does not currently expose the granularity of the third-party dependencies. However, from the prioritization data model point of view, there is no distinction between third-party and first-party packages. These would be represented as a package object at that point in time.

## Risk factors

By understanding your images, packages, and Kubernetes resources as "application context", Snyk can compute the following risk factors:

* [Deployed](risk-factor-deployed.md)
* [Loaded package](risk-factor-loaded-package.md)
* [OS condition](risk-factor-os-condition.md)
* [Public facing](risk-factor-public-facing.md)

You can enable and disable all of these "application context" risk factors through the Group **Settings**, on the **Issues** tab. If you choose to disable a risk factor, a provider selection, or the Kubernetes cluster mapping, Snyk will no longer compute them.&#x20;

{% hint style="info" %}
The Risk Factors are available only for Snyk AppRisk.
{% endhint %}

Depending on the integration options enabled for your application, risk factors are applied differently. You can [prioritize your integrations](../set-up-insights-for-snyk-apprisk/#prioritize-your-integrations) by customizing the available Insights options from the Group settings.

Risk factors are supported for stateful entities such as the following Kubernetes components: [StatefulSet](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/), [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/), and [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/).

<figure><img src="../../../.gitbook/assets/image (230).png" alt=""><figcaption><p>Group settings - Insights tab in the Group settings</p></figcaption></figure>

{% hint style="info" %}
Risk factor settings may take up to four hours to take effect.
{% endhint %}
