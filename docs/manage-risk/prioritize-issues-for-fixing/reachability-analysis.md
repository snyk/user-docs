# Reachability analysis

{% hint style="info" %}
**Release status**

Reachability analysis is in Early Access for some integrations and languages.

For information on how to enable the feature for supported integrations and languages, see [Snyk Preview](../../snyk-admin/snyk-preview.md).
{% endhint %}

Snyk reachability analysis allows you to analyze risk by identifying whether your application is calling a code element (e.g. functions, classes, modules, annotations, etc.) related to the vulnerability, thus raising the chances of that vulnerability being exploitable in the context of your application.

Reachability analysis can be used as a standalone signal to make decisions, or as part of a broader risk-based prioritization approach using the Snyk Risk Score.&#x20;

Snyk uses a combination of program analysis and various AI techniques to determine the reachability of a given vulnerability, with validation conducted by security research experts.

The following instructions explain how to set up and use reachability analysis and provide more information on how the underlying analysis works at Snyk.&#x20;

## Set up Reachability analysis

Follow these steps to enable Reachability analysis and begin analyzing Projects for reachable vulnerabilities:&#x20;

* In the Organization **Settings**, navigate to the **Snyk Open Source** section.
* In the **General** section, find **Reachable vulnerabilities**.
* Activate **Enable reachable vulnerabilities analysis**.

<figure><img src="../../.gitbook/assets/image (573).png" alt=""><figcaption><p>Enabling Reachability setting</p></figcaption></figure>

After Reachability analysis is enabled, the analysis is done as part of scanning Projects.&#x20;

{% hint style="info" %}
You can apply the reachability analysis to existing Projects by triggering a [manual test](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/#manual-pull-and-merge-requests-for-project-code).
{% endhint %}

## Supported languages and integrations

Reachability analysis is supported for the following languages and package managers:

| Language                                                                                                                                                                   | Package manager | Release status       |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------- | -------------------- |
| [Java](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/)                                                                                         | Maven, Gradle   | General Availability |
| [JavaScript](../../supported-languages-package-managers-and-frameworks/javascript/), [TypeScript](../../supported-languages-package-managers-and-frameworks/typescript.md) | npm, Yarn       | Early Access         |

Reachability analysis is supported in the following integrations:

| Integration                                                                                                                                                              | Release status       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------- |
| [GitHub](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md)                                                                                           | General Availability |
| [GitHub Enterprise](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-scm-contributors-count/scripts-for-scm-contributors-count/github-enterprise/) | General Availability |
| [Bitbucket Cloud](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-cloud-app.md)                                                                     | General Availability |
| [Bitbucket Server](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-data-center-server.md)                                                           | General Availability |
| [GitLab](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab.md)                                                                                           | General Availability |
| [Azure Repos](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/azure-repositories-tfs.md)                                                                      | General Availability |
| [Brokered connections](../../enterprise-setup/snyk-broker/connections-with-snyk-broker.md)                                                                               | General Availability |

{% hint style="info" %}
Reachability analysis using the Snyk CLI, IDE, or other integrations is not supported.
{% endhint %}

## **Enabling Reachability** analysis **for brokered connections**

If you use a brokered connection to your SCM, configure the Broker to provide access to your source files. See the [Snyk Code - Clone capability with Broker for Docker](../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker.md), the [Broker rules for Snyk Code](../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-helm-chart-installation/broker-rules-for-snyk-code.md), and the [Snyk Broker - Code Agent](../../enterprise-setup/snyk-broker/snyk-broker-code-agent/) documentation for configuration details when Broker is used with Snyk Code.&#x20;

## Using Reachability analysis

### Reachability status&#x20;

After a vulnerability is identified, it has one of the following reachability statuses:

* `REACHABLE` - A direct or indirect path was found from your application to the vulnerable code.
* `NO PATH FOUND` - No path found from your application to the vulnerable code.

{% hint style="info" %}
If a `NO PATH FOUND` status is given, do not assume that the vulnerability is totally unreachable or unexploitable.
{% endhint %}

Reachability analysis status is available [on the Project page](reachability-analysis.md#on-the-project-page), [as part of the Risk Score](reachability-analysis.md#as-part-of-the-risk-score), in the [Issues Detail report](../../manage-issues/reporting/available-snyk-reports.md#issues-detail-report), and through the API endpoint  [Get issues by Group ID](../../snyk-api/reference/issues.md#groups-group_id-issues).&#x20;

### Reachability analysis as shown on the Project page

After you import  or test a Project using the Snyk UI, the Project is monitored by Snyk, and the results of the reachable vulnerabilities analysis appear on the Project page in the following places:

* Filters - Allows you to focus on reachable vulnerabilities first by filtering results based on reachability.
* Reachability badge - Allows you to understand the reachability level of vulnerabilities.
* Call path - Allows you to see the path from your code to the vulnerable code element to validate the result.

<figure><img src="../../.gitbook/assets/image (124) (1) (1) (1) (2) (1) (1) (1) (2) (2).png" alt="Reachability filters, badge and call path on the Projects UI"><figcaption><p>Reachability filters, badge and call path on the Projects UI</p></figcaption></figure>

### Reachability analysis as part of the Risk Score

[Risk Score](risk-score.md) helps you apply holistic risk-based prioritization that combines multiple factors,  associated with either the vulnerability or the context of your application. Reachability analysis is such a contextual factor that will significantly increase the overall score.&#x20;

Risk Score is available on the Projects page and through the API and Reports.&#x20;

<div data-full-width="false"><figure><img src="../../.gitbook/assets/image (1) (7).png" alt="Reachability as part of the Risk Score"><figcaption><p>Reachability as part of the Risk Score</p></figcaption></figure></div>

{% hint style="info" %}
[Priority score](priority-score.md), the legacy model preceding the Risk Score, also takes reachable vulnerabilities into account.&#x20;
{% endhint %}

## How Reachable vulnerability analysis works

Snyk uses a combination of security expert analysis, program analysis, and various AI techniques to determine the reachability of a vulnerability, including these steps of analysis:&#x20;

1. **Enriching vulnerabilities with the patches applied to fix them** - as part of the Snyk vulnerability curation process, Snyk references the fix commit that the maintainer applied.&#x20;
2. **Related elements analysis**- Based on the commit fix, Snyk uses DeepCode AI program analysis to analyze the code elements and other parameters related to the vulnerability.&#x20;
3. **Root Cause analysis** - Snyk uses DeepCode AI and NLP techniques to automatically rank the related code elements by their chances of being the root cause of the vulnerability.  &#x20;
4. **Reachability analysis** -  As issues are found in your application by a Snyk scan, the DeepCode program analysis engine is used to analyze the call graph of your application in relation to the call graph between the open-source dependencies used. A path between your application and a code element ranked as a root cause will yield a “Reachable” vulnerability.&#x20;
5. **Security experts supervision** - Snyk security experts will manually verify and mark elements as root causes in order to make the entire analysis more accurate over time

The following considerations related to **false positives and false negatives** apply to Reachable vulnerability analysis.&#x20;

Program analysis requires a trade-off between accurate results, minimizing false positives, and recall rates, by avoiding potentially exploitable vulnerabilities.&#x20;

To facilitate this trade-off, Snyk DeepCode analysis applies real-time decision-making to determine whether to under-approximate the set of reachable elements based on analysis of the likelihood that a reachable path will be found in a specific environment. &#x20;

For example, it is not always possible to give a precise answer when reflection programming is used. In such a case, neither returning a large set of false positives nor returning “Not reachable” will suffice. Snyk Deep Code analysis optimizes in order to retrieve the most accurate and complete result possible for a given code structure.&#x20;
