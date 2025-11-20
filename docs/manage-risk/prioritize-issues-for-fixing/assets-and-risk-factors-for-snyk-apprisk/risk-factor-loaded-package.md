# Risk factor: loaded package

{% hint style="info" %}
**Release status**

The Loaded package risk factor is in Early Access and available only with Snyk Enterprise plans with Snyk AppRisk. If you want to set it up in your Group, contact your Snyk account team.
{% endhint %}

Vulnerabilities in a package are only exploitable if relevant code is executed. For that to occur, the package must be loaded into memory. Snyk can use data from integrations to match the packages loaded in a Kubernetes container against the issues found when Snyk scans the container, helping to prioritize issues that are more likely to be exploitable.

Snyk collects information on packages loaded into memory from the runtime environment. This information can be collected either from a runtime integration or from the Snyk Runtime Sensor. The Loaded package risk factor helps to identify packages that impose the highest level of risk. Unlike theoretical risks, the likelihood of exploiting a vulnerability increases when a package is loaded into memory.

{% hint style="info" %}
The Loaded package risk factor is an enhancement of the Deployed risk factor.

Loaded packages are associated with a Kubernetes container deployed in a cluster that is known to Snyk. Therefore, an issue cannot have a loaded package risk factor without also having a deployed risk factor.
{% endhint %}

## Types of integration

The Loaded package risk factor works with the Snyk Runtime Sensor and third-party integrations.

### Snyk Runtime Sensor

The Snyk Runtime Sensor improves the identification and management of loaded package risks by providing real-time data on packages loaded into memory within Kubernetes containers. It closely integrates with your runtime environment to continuously monitor active packages and cross-reference this information with known vulnerabilities. This dynamic approach ensures accurate application of the loaded package risk factor, using up-to-date data to prioritize and mitigate risks based on actual runtime conditions.

{% hint style="info" %}
Navigate to the [Snyk Runtime Sensor documentation](broken-reference) to learn more about the supported ecosystems from which the Loaded package risk factor is reported.
{% endhint %}

### Third-party integrations

Third-party integrations make the Loaded package risk factor even stronger by providing extra data sources and analysis capabilities. Snyk integrates with various security tools and platforms to gather comprehensive information about loaded packages from different environments. These integrations allow for cross-referencing loaded packages against extensive vulnerability databases, resulting in a more robust identification and prioritization of risks. Third-party integrations help ensure a comprehensive and up-to-date vulnerability management process by leveraging the strengths and specialized knowledge of various security ecosystems.

## Technical details for the Loaded package risk factor

The [runtime integrations ](../../../integrations/connect-a-third-party-integration.md)must provide a list of packages loaded in a container running within Kubernetes. The container is identified by:

* the Kubernetes cluster name
* the Kubernetes namespace
* the Kubernetes container name
* the Kubernetes image name

When you run the [risk-based prioritization](../prioritization-for-snyk-essentials.md) for Snyk AppRisk pipelines, a snapshot of this information is used to attach lists of loaded packages to the deployed images and facts Snyk knows about.

The packages are identified by:

* package manager
* name
* version

The Snyk issues associated with deployed images are evaluated to label any that have a matching loaded package.

The image name, package name, and other details are provided by the runtime data provider and must match the details Snyk identified in order for the loaded package risk factor to be attached to an issue.
