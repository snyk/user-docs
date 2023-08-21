# Reachable vulnerabilities

A reachable vulnerability has a path from your code to the vulnerable function in a dependency. Snyk reachable vulnerability scanning allows you to gauge risk by identifying whether a vulnerable function is being called by your application.

Scanning for reachable vulnerabilities is a way to prioritize vulnerabilities. By analyzing how your code uses open-source dependencies, and also how the open-source dependencies interact with each other, Snyk can provide insight into which vulnerabilities are reachable, directly from your own code and indirectly through transitive dependencies.

You can use reachability as one signal in a broader strategy for prioritizing issues.

## How finding reachable vulnerabilities works

First, Snyk identifies the vulnerable functions associated with specific vulnerabilities. Snyk does this using a call graph, essentially a map of function calls throughout your application from the top level right down through the transitive dependencies.

Next, Snyk builds a call graph that is the union of the calls within the application and the calls within open-source libraries scanned by Snyk. If a vulnerable function is reachable using this unified call graph from the application code, Snyk identifies the vulnerability as reachable.

Each vulnerability then receives a reachability status:

* **Reachable** - A direct or indirect path from your code to vulnerable functions was found OR
* **No path found** - No path found from your code to vulnerable functions

{% hint style="warning" %}
If a "no path found" status is given, do not assume that the vulnerability is totally unreachable.
{% endhint %}

Producing a precise call graph may require that the analysis sometimes under-approximates the set of reachable functions.

For example, although the Snyk solution does support code using reflection, it is not always possible to give a precise answer if reflection is used. In such a case, there is a tradeoff between returning a large set of false positives, or possibly returning that no path was found even if a call path is theoretically possible.

## Supported languages and integrations for reachable vulnerabilities

Reachable vulnerabilities analysis is available for **Java** (Maven and Gradle) with [supported versions](../../scan-applications/snyk-open-source/snyk-open-source-supported-languages-and-package-managers/snyk-for-java-and-kotlin.md) using the **GitHub** integration.

The public beta version supports **Bitbucket Cloud** and **Bitbucket Server**, **Gitlab**, and **Azure Repos**, including brokered connections.

Reachable vulnerability analysis via the Snyk CLI, other Git integrations, and other languages is not currently supported.

## Enabling reachability for all SCM providers

You can enable reachable vulnerability analysis for other SCM providers. This option is available in public beta, and you can enable it from the **Snyk Preview** interface.&#x20;

1. Open the Snyk Web UI.
2. Go to **Settings** > **Snyk Preview**.
3. Activate the **Enable reachability for all SCM providers** option**.**&#x20;

<figure><img src="../../.gitbook/assets/Enable_Beta_Feature (1).png" alt="Enabling reachability for all SCM providers."><figcaption><p>Enabling reachability for all SCM providers</p></figcaption></figure>



If you are setting up reachable vulnerability analysis for the first time, you must enable it.

1. Open the Snyk Web UI.
2. Go to **Settings** > **Languages**.
3. Enable the **Reachable vulnerabilities** option**.**

<figure><img src="../../.gitbook/assets/Enable_rechable_vullerabilities.png" alt="Enabling the reachable vulnerabilities option"><figcaption><p>Enabling the reachable vulnerabilities option</p></figcaption></figure>

## Brokered connections

If you use a brokered connection to your SCM, configure the Broker to provide access to your source files. See the [Snyk Code - Clone capability with Broker for Docker](../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md), the [Broker rules for Snyk Code](../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-helm-chart-installation/broker-rules-for-snyk-code.md), and the [Snyk Broker - Code Agent](../../enterprise-setup/snyk-broker/snyk-broker-code-agent/) documentation for configuration details when the Broker is used with Snyk Code.&#x20;

## Scanning for reachable vulnerabilities using Git integrations

{% hint style="info" %}
To provide this feature, Snyk creates a temporary copy of your Git repository contents.

For more information, see [How Snyk handles your data](../../more-info/how-snyk-handles-your-data.md).
{% endhint %}

1. Set up your [GitHub integration](../../integrations/git-repository-scm-integrations/github-integration.md).
2. Enable Reachable Vulnerabilities analysis:
   * In the **Organization** settings, go to the **Languages** section.
   * Go to the **Reachable vulnerabilities analysis** section.
   * Activate the **Reachable vulnerabilities analysis** and save your changes.
3. Go to the **import projects** page and choose the repositories to import to Snyk.\
   Selected Projects are imported and analyzed for reachable vulnerabilities issues.

If you have existing Projects to which you want to add reachability analysis, enable it via the Organization settings and retest the Projects manually or wait for the daily recurring test.

## Reachable vulnerabilities on the Project page

After import or testing of a Project via Snyk UI, the Project is monitored by Snyk, and the results of the reachable vulnerabilities analysis appear on the Project page in the following places:

1. **Filters** - Allows you to focus on reachable vulnerabilities first by filtering results based on reachability.
2. **Reachability badge** - Allows you to quickly understand the reachability level of vulnerabilities.
3. **Call path** - Allows you to see the path from your code to the vulnerable function to validate the result.
4. **Priority score** - Allows you to see reachable vulnerabilities first, as Snyk will increase the priority score for reachable vulnerabilities.

<figure><img src="../../.gitbook/assets/image (124) (1) (1) (1) (2) (1) (1) (1) (2).png" alt="Reachability call path"><figcaption><p>Reachability call oath</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (126) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Priority Score"><figcaption><p>Priority score</p></figcaption></figure>

## Reachability status in reports

You can filter by reachability status to quickly show the reachable issues.&#x20;

Reachability is an input into the priority score. If an issue is found to be reachable, it can increase the score.

<figure><img src="../../.gitbook/assets/image (137) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Issue Filters"><figcaption><p>Issue Filters</p></figcaption></figure>
