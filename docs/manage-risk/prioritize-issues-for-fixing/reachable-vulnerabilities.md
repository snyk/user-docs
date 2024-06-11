# Reachable vulnerabilities

Snyk reachable vulnerability scanning allows you to gauge risk by identifying whether a function related to the vulnerability is being called by your application, raising the chances of that vulnerability being exploitable in the context of your application.

Reachable vulnerabilities can be used as a single signal to make decisions or as part of a broader risk-based prioritization approach, like the Risk Score.&#x20;

The following instructions explain how to set up and use reachable vulnerabilities, as well as provide more information on how reachability analysis works at Snyk.&#x20;

## Set up reachable vulnerability analysis

To set up reachable vulnerabilities analysis:&#x20;

* In the Organization settings, navigate to the Languages section.
* Navigate to the **Reachable vulnerabilities** section.
* Activate **Reachable vulnerabilities analysis** and save your changes.

<figure><img src="../../.gitbook/assets/image (2) (9).png" alt="Enabling reachability setting"><figcaption><p>Enabling reachability setting</p></figcaption></figure>

After it is enabled, reachable vulnerabilities analysis is done as part of scanning Projects.&#x20;

{% hint style="info" %}
To update existing Projects with the reachability analysis immediately, trigger a [manual test](../../scan-with-snyk/pull-requests/snyk-fix-pull-or-merge-requests/#manual-pull-and-merge-requests-for-project-code).
{% endhint %}

## Supported languages and integrations for Reachable Vulnerabilities

Reachable vulnerabilities analysis is available for Java (Maven and Gradle) Projects.

The following integrations are supported for Reachable Vulnerabilities analysis:

* [GitHub](../../scm.-ide-and-ci-cd-workflow/git-repositories-scms-integrations-with-snyk/snyk-github-integration.md)&#x20;
* [Bitbucket Cloud](../../scm.-ide-and-ci-cd-workflow/git-repositories-scms-integrations-with-snyk/snyk-bitbucket-cloud-app-integration.md) (Currently in Early Access)
* [Bitbucket Server](../../scm.-ide-and-ci-cd-workflow/git-repositories-scms-integrations-with-snyk/snyk-bitbucket-data-center-server-integration.md) (Currently in Early Access)
* [GitLab](../../scm.-ide-and-ci-cd-workflow/git-repositories-scms-integrations-with-snyk/snyk-gitlab-integration.md) (Currently in Early Access)
* [Azure Repos](../../scm.-ide-and-ci-cd-workflow/git-repositories-scms-integrations-with-snyk/snyk-azure-repositories-tfs-integration.md) (Currently in Early Access)
* [Brokered connections](../../enterprise-configuration/snyk-broker/connections-with-snyk-broker.md) (Currently in Early Access)

{% hint style="warning" %}
**Release status**&#x20;

Reachability for all SCM providers is in [Early Access](../../getting-started/snyk-release-process.md#early-access) and available only for Enterprise plans.

To enable the feature, see [Snyk Preview](https://docs.snyk.io/snyk-admin/manage-settings/snyk-preview).
{% endhint %}

{% hint style="info" %}
Reachable vulnerability analysis using the Snyk CLI, other Git integrations, and other languages is not currently supported.
{% endhint %}

## **Enabling Reachability for brokered connections**

If you use a brokered connection to your SCM, configure the Broker to provide access to your source files. See the [Snyk Code - Clone capability with Broker for Docker](../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md), the [Broker rules for Snyk Code](../../enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-helm-chart-installation/broker-rules-for-snyk-code.md), and the [Snyk Broker - Code Agent](../../enterprise-configuration/snyk-broker/snyk-broker-code-agent/) documentation for configuration details when the Broker is used with Snyk Code.&#x20;

## Using Reachability

### Reachability status&#x20;

After a vulnerability is identified, it has a reachability status

* Reachable - A direct or indirect path was found from your application to the vulnerable code.
* No path found - No path found from your application to the vulnerable code.

{% hint style="info" %}
If a `no path found` status is given, do not assume that the vulnerability is totally unreachable or unexploitable.
{% endhint %}

Reachability status is available [on the Project page](reachable-vulnerabilities.md#on-the-project-page), [as part of the Risk Score](reachable-vulnerabilities.md#as-part-of-the-risk-score), in the [Issues Detail report](../../manage-issues/reporting/available-snyk-reports.md#issues-detail-report),  and through the [Get issues by group ID](https://apidocs.snyk.io/?#get-/groups/-group\_id-/issues) API.

### Reachability as shown on the Project page

After you import  or test a Project using the Snyk UI, the Project is monitored by Snyk, and the results of the reachable vulnerabilities analysis appear on the Project page in the following places:

1. Filters - Allows you to focus on reachable vulnerabilities first by filtering results based on reachability.
2. Reachability badge - Allows you to quickly understand the reachability level of vulnerabilities.
3. Call path - Allows you to see the path from your code to the vulnerable function to validate the result.

<figure><img src="../../.gitbook/assets/image (124) (1) (1) (1) (2) (1) (1) (1) (2) (2).png" alt="Reachability call path"><figcaption><p>Reachability filters, badge and call path on the Projects UI</p></figcaption></figure>

### Reachability as part of the Risk Score

[Risk Score](risk-score.md) helps you apply holistic Risk-Based prioritization that combines multiple factors,  associated with either the vulnerability or the context of your application. Reachability is such a contextual factor that will significantly increase the overall score.&#x20;

Risk Score is available on the Projects page and through the API and Reports.&#x20;

<div data-full-width="false">

<figure><img src="../../.gitbook/assets/image (1) (7).png" alt=""><figcaption><p>Reachability as part of the Risk Score</p></figcaption></figure>

</div>

{% hint style="info" %}
[Priority score](priority-score.md), the legacy model preceding the Risk Score, also takes reachable vulnerabilities into account.&#x20;
{% endhint %}

## Reachable vulnerability analysis&#x20;

Snyk uses a combination of security expert analysis, program analysis, and various AI techniques to determine the reachability of a vulnerability, including these steps of analysis:&#x20;

1. **Enriching vulnerabilities with the patches applied to fix them** - as part of the Snyk vulnerability curation process, Snyk references the fix commit that the maintainer applied.&#x20;
2. **Related elements analysis**- Based on the commit fix, Snyk uses DeepCode AI program analysis to analyze the functions and other elements related to the vulnerability.&#x20;
3. **Root Cause analysis** - Snyk uses DeepCode AI and NLP techniques to automatically rank the related code elements by their chances of being the root cause of the vulnerability.  &#x20;
4. **Reachability analysis** -  As issues are found in your application by a Snyk scan, the DeepCode program analysis engine is used to analyze the call graph of your application in relation to the call graph between the open-source dependencies used. A path between your application and a code element ranked as a root cause will yield a “Reachable” vulnerability.&#x20;
5. **Security experts supervision** - Snyk security experts will manually verify and mark elements as root causes in order to make the entire analysis more accurate over time

The following considerations related to **false positives and false negatives** apply to Reachble vulnerability analysis.&#x20;

Program analysis requires a trade-off between accurate results (minimizing false positives) and recall rates (not missing potentially exploitable vulnerabilities).&#x20;

To facilitate this trade-off, Snyk DeepCode analysis applies real-time decision-making to determine whether to under-approximate the set of reachable elements based on analysis of the likelihood  a reachable path will be found in a specific environment. &#x20;

For example, it is not always possible to give a precise answer when reflection programming is used. In such a case, neither returning a large set of false positives nor returning “Not reachable” will suffice. Snyk Deep Code analysis optimizes in order to retrieve the most accurate and complete result possible for a given code structure. \
