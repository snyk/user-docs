# Risk Score

{% hint style="info" %}
**Release status**

Risk Score is in Early Access and available for Snyk Open Source and Snyk Container, with Snyk Enterprise and Snyk Free plans. If you want to set it up in your Group, contact your Snyk account team.

Use [Snyk Preview](../../snyk-platform-administration/snyk-preview.md) to replace the Priority Score with the new Risk Score for Snyk Open Source and Snyk Container issues.
{% endhint %}

The Snyk Risk Score is a single value assigned to an issue, applied by automatic risk analysis for all vulnerability-type issues. License issues are not supported. Risk Score is based on the potential impact and likelihood of exploitability. Ranging from 0 to 1,000, the score represents the risk imposed on your environment and enables a risk-based prioritization approach.&#x20;

Risk score remains the same over time if the contributing factors do not change. However, some contributing factors,  such as the Exploit Prediction Scoring System (EPSS), can potentially change frequently. The number of days since the vulnerability was first published is also a factor and causes the score to change once, when the number of days becomes more than one year, and the likelihood subscore decreases.&#x20;

Since real risk is scarce, you should expect a significant drift in the distribution of scores, as can be seen in this example of Project score distributions:&#x20;

<div data-full-width="false"><figure><img src="../../.gitbook/assets/image (18).png" alt="Example Project scores distribution"><figcaption><p>Example Project scores distribution</p></figcaption></figure></div>

Risk Score replaces the Priority Score directly. See the [priority score docs](priority-score.md) for how to interact with the Risk Score in the UI, API, and Reports, where the Risk Score is now introduced when enabled.&#x20;

Risk Score is not available in the CLI.&#x20;

{% hint style="info" %}
The Risk Score replaces the Priority Score only after the Snyk Open Source and Snyk Container Projects are retested.&#x20;

You can wait for the Projects to be automatically retested (daily for Snyk Open Source and weekly for Snyk Container), or you can manually retest the Project.
{% endhint %}

{% hint style="warning" %}
In the API, the relevant fields are still named with `priority`. When Risk Score is enabled, the scores and factors populated in these fields are based on the Risk Score model as part of the early access stage.
{% endhint %}

## Explore the Risk Score by issue&#x20;

When looking at the Issue card information, hover over the score to see the type of score (Priority or Risk Score) that is being displayed. The Risk Score tooltip provides information about the subscore and the Risk Factors contributing to the score.

<div data-full-width="false"><figure><img src="../../.gitbook/assets/image (166).png" alt="Risk Score tooltip" width="563"><figcaption><p>Risk Score tooltip showing Impact and Likelihood</p></figcaption></figure></div>

## About the Risk Score model&#x20;

<figure><img src="../../.gitbook/assets/matrix (2) (2).png" alt="he Snyk Risk Score Model"><figcaption><p>The Snyk Risk Score Model</p></figcaption></figure>

The model that powers the Risk Score applies automatic risk analysis for each security issue based on the potential impact and likelihood of exploitability.

{% hint style="info" %}
The Risk model results from extensive research conducted by the Snyk Security Data Science team and experienced security researchers. The model draws on expertise gained over the years in developing the [Snyk Vulnerability Database](https://security.snyk.io/).
{% endhint %}

### Impact subscore

Objective impact factors are the CVSS impact metrics, Availability, Confidentiality, Integrity, and Scope, calculated based on the CVSS impact subscore. For Container issues, Provider Urgency is also taken into account.&#x20;

The business criticality Project attribute will be taken into account as a contextual impact factor, increasing or decreasing the impact subscore. For more information, see [Project attributes](../../snyk-platform-administration/snyk-projects/project-attributes.md).

### Likelihood subscore&#x20;

Objective likelihood factors are taken into account:

* Exploit Maturity
* Exploit Prediction Scoring System (EPSS)&#x20;
* Age of advisory&#x20;
* CVSS exploitability metrics: Attack vector, Privileges required, User interaction, and Scope
* Social Trends
* Malicious Package
* Provider Urgency (Snyk Container)
* Package popularity (Snyk Open Source)&#x20;

{% hint style="warning" %}
The Exploit Prediction Scoring System (EPSS)  is updated daily.
{% endhint %}

Contextual likelihood factors then increase or decrease the likelihood subscore: &#x20;

* Reachability (Snyk Open Source Java, JavaScript)&#x20;
* Transitive depth

## Risk factors drill down

### Objective impact risk factors

#### Confidentiality

Represents the impact on the customer’s data confidentiality, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-3-1-Confidentiality-C).\
Possible input values: `None`, `Low`, `High`

#### Integrity

Represents the impact on the customer’s data integrity, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-3-2-Integrity-I).\
Possible input values: `None`, `Low`, `High`

#### Availability

Represents the impact of the customer’s application availability based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-3-3-Availability-A).\
Possible input values are `None`, `Low`, and `High.`

#### Scope

Indicates whether the vulnerability can affect components outside of the target’s security scope, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-2-Scope-S).\


The objective impact subscore is calculated based on the CVSS impact subscore. For more information, see the references to CVSS definitions above and the [subscore equations](https://www.first.org/cvss/v3.1/specification-document#7-1-Base-Metrics-Equations).

| Possible input value | Score impact                      |
| -------------------- | --------------------------------- |
| `Unchanged`          | Impact subscore is not affected.  |
| `Changed`            | Impact subscore is affected.      |

#### Provider urgency (Snyk Container)&#x20;

Urgency rating as provided by the relevant operating system distribution security team. For more information, see [External information sources for relative importance](../../scan-with-snyk/snyk-container/how-snyk-container-works/severity-levels-of-detected-linux-vulnerabilities.md#external-information-sources-for-relative-importance) in severity levels of detected Linux vulnerabilities.

| Possible input value | Score impact                             |
| -------------------- | ---------------------------------------- |
| `Critical`           | Impact subscore increases significantly. |
| `High`               | Impact subscore increases.               |
| `Medium`             | Impact subscore decreases significantly. |
| `Low`                | Impact subscore decreases significantly. |

{% hint style="info" %}
Provider Urgency affects the Likelihood subscore.&#x20;
{% endhint %}

### Contextual impact risk factor

**Business criticality**&#x20;

User-defined Project attribute representing the subjective business impact of the respective application. For more information, see [Project attributes](../../snyk-platform-administration/snyk-projects/project-attributes.md).

| Possible input value | Score impact                             |
| -------------------- | ---------------------------------------- |
| `Critical`           | Impact subscore increases.               |
| `High`               | Impact subscore is not affected.         |
| `Medium`             | Impact subscore decreases.               |
| `Low`                | Impact subscore decreases significantly. |

{% hint style="info" %}
When you apply a business criticality attribute to a Project, a retest is required for the Risk Scores to incorporate the new data. When no Business Criticality is assigned, the Impact subscore will not be affected.&#x20;
{% endhint %}

When the business criticality for a Project is not configured, the `high` default value is used so that the subscore remains unaffected.

### Objective likelihood risk factors

#### Exploit maturity&#x20;

Represents the existence and maturity of any public exploit retrieved and validated by Snyk. For more information, see [View exploits, How exploits are determined](view-exploits.md#how-exploits-are-determined).

| Possible input value | Score impact                             |
| -------------------- | ---------------------------------------- |
| `No Known Exploit`   | Impact subscore decreases significantly. |
| `Proof of Concept`   | Impact subscore decreases slightly.      |
| `Functional`         | Impact subscore increases.               |
| `High`               | Impact subscore increases significantly. |

#### EPSS score&#x20;

The Exploit Prediction Scoring System (EPSS) predicts whether a CVE will be exploited in the wild based on an elaborate model created and owned by the FIRST Organization.

The probability is the direct output of the EPSS model and conveys an overall sense of the threat of exploitation in the wild. This data is updated daily, relying on the latest available EPSS model version. See the EPSS [documentation](https://www.first.org/epss/articles/prob_percentile_bins) for more details.\
The possible input values are the `EPSS score [0.00-1.00]`

{% hint style="info" %}
The likelihood subscore increases significantly according to the EPSS score.
{% endhint %}

If the vulnerability is found to be malicious, the EPSS value is set to `1`. If no information regarding the EPSS is available, then the default value is `0.01055`.

#### Attack vector&#x20;

Represents the context by which vulnerability exploitation is possible, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-1-Attack-Vector-AV).

| Possible input values | Score impact                                                                                               |
| --------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Network`             | Likelihood subscore increases.                                                                             |
| `Adjacent`            | Likelihood subscore decreases according to the level of remote access needed to exploit the vulnerability. |
| `Local`               | Likelihood subscore decreases according to the level of remote access needed to exploit the vulnerability. |
| `Physical`            | Likelihood subscore decreases according to the level of remote access needed to exploit the vulnerability. |

#### Attack complexity&#x20;

Represents the level of complexity defined by the conditions that must exist to exploit the vulnerability, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-2-Attack-Complexity-AC).

| Possible input values | Score impact                   |
| --------------------- | ------------------------------ |
| `High`                | Likelihood subscore decreases. |
| `Low`                 | Likelihood subscore increases. |

#### Privileges required&#x20;

Represents the level of privileges an attacker must possess before successfully exploiting the vulnerability, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-3-Privileges-Required-PR).

| Possible input values | Score impact                                                                 |
| --------------------- | ---------------------------------------------------------------------------- |
| `High`                | Likelihood subscore decreases according to the level of privileges required. |
| `Low`                 | Likelihood subscore decreases according to the level of privileges required. |
| `None`                | Likelihood subscore increases.                                               |

#### User interaction&#x20;

Represents the need for action from a user as part of the exploitation process, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-4-User-Interaction-UI).

| Possible input values | Score impact                   |
| --------------------- | ------------------------------ |
| `Required`            | Likelihood subscore decreases. |
| `None`                | Likelihood subscore increases. |

#### Social trends&#x20;

Represents the social media traffic regarding this vulnerability. Snyk research has shown that greater social media interaction can predict future exploitation or point to existing exploitation. For more information, see [Vulnerabilities with social trends](vulnerabilities-with-social-trends.md).

| Possible input values | Score impact                         |
| --------------------- | ------------------------------------ |
| `Trending`            | Likelihood subscore increases.       |
| `Not trending`        | Likelihood subscore does not change. |

#### Malicious package&#x20;

Malicious code deployed as a supply chain dependency is considered highly exploitable.

| Possible input values | Score impact                                                        |
| --------------------- | ------------------------------------------------------------------- |
| `True`                | Likelihood subscore increases significantly for Malicious Packages. |
| `False`               | Likelihood subscore remains unchanged.                              |

#### Age of vulnerability&#x20;

A new vulnerability (up to one year) is more likely to be exploited than an old vulnerability (more than one year since publication)&#x20;

| Possible input values                             | Score impact                                                                                                               |
| ------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| Days since the vulnerability was first published. | <p>Less than one year old - Likelihood subscore increases. </p><p>Over one year old -  Likelihood subscore decreases. </p> |

#### Package popularity (Snyk Open Source)&#x20;

If a package is relatively more popular for its ecosystem, it is more likely to be exploited as hackers benefit from a wider pool of potential targets.

| Possible input values | Score impact                           |
| --------------------- | -------------------------------------- |
| `High`                | Likelihood subscore increases.         |
| `Medium`              | Likelihood subscore does not change.   |
| `Low`                 | Likelihood subscore decreases.         |

#### Provider urgency (Snyk Container)&#x20;

Importance rating as provided by the relevant operating system distribution security team. For more information, see [External information sources for relative importance](../../scan-with-snyk/snyk-container/how-snyk-container-works/severity-levels-of-detected-linux-vulnerabilities.md#external-information-sources-for-relative-importance) in severity levels of detected Linux vulnerabilities.

| Possible input values | Score impact                             |
| --------------------- | ---------------------------------------- |
| `Critical`            | Impact subscore increases significantly. |
| `High`                | Impact subscore increases.               |
| `Medium`              | Impact subscore decreases.               |
| `Low`                 | Impact subscore decreases significantly. |

{% hint style="info" %}
When neither CVSS nor Importance rating is provided, Provider Urgency is set to`Low`by default. Provider Urgency also affects the Impact subscore.&#x20;
{% endhint %}

### Contextual likelihood risk factors

#### Transitive depth&#x20;

Building on [past studies](https://arxiv.org/pdf/2301.07972.pdf), Snyk research has shown that if a vulnerability is introduced to a Project transitively rather than directly, it is less likely that an exploitable function path will exist.

| Possible input values | Score impact                         |
| --------------------- | ------------------------------------ |
| `Direct dependency`   | Likelihood subscore does not change. |
| `Indirect dependency` | Likelihood subscore decreases.       |

#### Reachability&#x20;

Snyk static code analysis determines whether the vulnerable method is being called. This is supported for Java and JavaScript. For more details, navigate to the [Reachability analysis](reachability-analysis.md) page.&#x20;

When Reachability is not enabled, the Likelihood subscore will not change, and the factor will not show up.

| Possible input values | Score impact                                                            |
| --------------------- | ----------------------------------------------------------------------- |
| `Reachable`           | Likelihood subscore increases, and transitive depth is not considered.  |
| `No path found`       | Likelihood subscore does not change.                                    |
