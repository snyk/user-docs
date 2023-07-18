# Risk Score



{% hint style="info" %}
Risk Score is currently in Closed Beta for Snyk Open Source and Snyk Container. If you are interested in replacing the Priority Score with the Risk Score, please get in touch with your Snyk account team. See the [Snyk feature release process](../../more-info/snyk-feature-release-process.md) for more details.
{% endhint %}

## Overview

The Snyk Risk Score is a single value assigned to an issue, applied by automatic risk analysis for each security issue and based on the potential impact and likelihood of exploitability. From 0 to 1000, the score represents the risk imposed on your environment and enables a risk-based prioritization approach.&#x20;

Since real risk is scarce, you should expect a significant drift in the distribution of scores, as can be seen in this example Project scores distributions:&#x20;

<figure><img src="../../.gitbook/assets/image (2) (1) (3).png" alt="Example Project scores distribution"><figcaption><p>Example Project scores distribution</p></figcaption></figure>

{% hint style="info" %}
As part of the closed beta, the Risk Score replaces the Priority Score directly. See the [priority score docs](priority-score.md) for how to interact with the Risk Score in the UI, API, and Reports, where it is now introduced when enabled.&#x20;
{% endhint %}

{% hint style="info" %}
The Priority Score will be replaced with the Risk Score upon retesting Snyk Open Source and Snyk Container Projects.&#x20;
{% endhint %}

{% hint style="warning" %}
Note that in the API, the relevant fields are still named with `priority.`When enabled as part of the beta, the scores and factors populated in these fields are based on the Risk Score model.&#x20;
{% endhint %}

## About the Risk Score Model&#x20;

<figure><img src="../../.gitbook/assets/matrix (2).png" alt=""><figcaption><p>The Snyk Risk Score Model</p></figcaption></figure>

The model that powers the Risk Score applies automatic risk analysis for each security issue based on the potential impact and likelihood of exploitability.

{% hint style="info" %}
The Risk Model results from extensive research conducted by the Snyk Security Data Science team and experienced security researchers. It draws upon years of expertise in developing the [Snyk Vulnerability Database](https://security.snyk.io/).
{% endhint %}

### Impact subscore

* Objective impact factors are the CVSS impact metrics (Availability, Confidentiality, Integrity, and Scope) calculated based on the CVSS impact subscore.
* Coming soon - Business criticality Project attribute ([learn more](https://docs.snyk.io/manage-issues/introduction-to-snyk-projects/project-attributes#available-attributes-and-their-values)) will be considered a contextual impact factor, increasing or decreasing the impact subscore.

### Likelihood subscore&#x20;

* Objective likelihood factors are taken into account:
  * Exploit Maturity
  * Exploit Prediction Scoring System (EPSS)&#x20;
  * Age of advisory&#x20;
  * CVSS exploitability metrics (Attack vector, Privileges required, User interaction, and Scope)&#x20;
  * Social Trends
  * Malicious Package
  * Coming soon -  Disputed vulnerability and Package popularity&#x20;
* Contextual likelihood factors then increase or decrease the likelihood subscore: &#x20;
  * Reachability (Java only, JavaScript coming soon)&#x20;
  * Transitive depth&#x20;
  * Coming soon - Insights such as public exposure and vulnerability condition applicability&#x20;

\
Impact and Likelihood scores are then multiplied into a final Risk Score.&#x20;

{% hint style="info" %}
"Fixability" is no longer considered part of the Score Calculation, as the effort needed to mitigate a security issue does not affect the Risk it imposes. To focus on actionable issues, use Fixability filters and then the Risk Score to start with the riskiest issues.&#x20;
{% endhint %}

## Risk factors drill down

### Objective impact risk factors

#### Confidentiality

Represents the impact on customer’s data confidentiality, based on CVSS definition.

**Possible input values:** _None, Low, High_

#### Integrity

Represents the impact on customer’s data integrity, based on CVSS definition.

**Possible input values:** _None, Low,  High_

#### Availability

Represents the impact of customer’s application availability based on CVSS definition.

**Possible input values:** _None, Low, High_

#### Scope

Represents whether the vulnerability can affect components outside of the target’s security scope, based on CVSS definition.

**Possible input values:** _Unchanged, Changed_

{% hint style="info" %}
**How would these affect the score?** \
The objective impact subscore is calculated based on the CVSS impact subscore. Learn more on CVSS [impact definitions](https://www.first.org/cvss/v3.1/specification-document#2-3-Impact-Metrics) and [subscore equations](https://www.first.org/cvss/v3.1/specification-document#7-1-Base-Metrics-Equations)
{% endhint %}

### Contextual impact risk factors (Coming Soon)

#### Business Criticality&#x20;

User-defined Project attribute representing the subjective business impact of the respective application ([learn more](https://docs.snyk.io/manage-issues/introduction-to-snyk-projects/project-attributes))

**Possible input values:** _Critical, High, Medium, Low_&#x20;

{% hint style="info" %}
**How would this affect the score?** \
_Critical_ - Impact subscore will increase\
_High_ - Impact subscore will not be affected \
_Medium -_ Impact subscore will decrease\
_Low -_ Impact subscore will decrease significantly
{% endhint %}

### Objective likelihood risk factors

#### EPSS score&#x20;

Exploit Prediction Scoring System (EPSS), predicting whether a CVE would be exploited in the wild, based on an elaborated model created and owned by the FIRST Organization. \
The probability is the direct output of the EPSS model, and conveys an overall sense of the threat of exploitation in the wild. This data is updated daily, relying on the latest available EPSS model version. Check out the EPSS [documentation](https://www.first.org/epss/articles/prob\_percentile\_bins) for more details.

**Possible input values:** _EPSS score \[0.00-1.00]_

{% hint style="info" %}
**How would this affect the score?** \
The likelihood subscore will increase significantly according to the EPSS score
{% endhint %}

#### Exploit Maturity&#x20;

Represents the existence and maturity of any public exploit retrieved and validated by Snyk ([learn more](https://docs.snyk.io/manage-issues/issue-management/view-exploits#how-it-works-how-exploits-are-determined))

**Possible input values:** _None, Proof of Concept, Functional, High_

{% hint style="info" %}
**How would this affect the score?** \
The likelihood subscore will increase significantly according to the level of Exploit Maturity
{% endhint %}

#### Malicious Package&#x20;

Malicious code deployed as a supply chain dependency is considered highly exploitable

**Possible input values:** _True, False_

{% hint style="info" %}
**How would this affect the score?**&#x20;

The Likelihood subscore will increase significantly for Malicious Packages&#x20;
{% endhint %}

#### Attack Complexity&#x20;

Represents the level of complexity defined by the conditions that must exist to exploit the vulnerability, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-Exploitability-Metrics).

**Possible input values:** _Low, High_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Low_ - Likelihood subscore will increase

_High_ - Likelihood subscore will decrease
{% endhint %}

#### Attack Vector&#x20;

Represents the context by which vulnerability exploitation is possible, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-Exploitability-Metrics).

**Possible input values:** _Network, Adjacent, Local, Physical_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Network_ - Likelihood subscore will increase

_Adjacent, Local, Physical_ - Likelihood subscore will decrease according to the level of remote access needed for exploit
{% endhint %}

#### Privileges Required&#x20;

Represents the level of privileges an attacker must possess before successfully exploiting the vulnerability, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-Exploitability-Metrics).

**Possible input values:** _None, Low, High_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_None_ - Likelihood subscore will increase&#x20;

_Low, High_ - Likelihood subscore will decrease according to the level of privileges required&#x20;
{% endhint %}

#### User Interaction&#x20;

Represents the need for action from a user as part of the exploitation process, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-Exploitability-Metrics).

**Possible input values:** _None, Required_&#x20;

{% hint style="info" %}
**How would this affect the score?**&#x20;

_None_ -  Likelihood subscore will increase&#x20;

_Required_ - Likelihood subscore will decrease&#x20;
{% endhint %}

#### Social Trends&#x20;

Represents the social media traffic regarding this vulnerability. Our research has shown that greater social media interaction can predict future exploitation or point to existing exploitation ([learn more](https://docs.snyk.io/manage-issues/issue-management/prioritize-by-social-trends)).&#x20;

**Possible input values:**  _Currently Trending, Not trending_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Currently Trending_ - Likelihood subscore will increase&#x20;

_Not trending_ - Likelihood subscore will not change&#x20;
{% endhint %}

#### Age of vulnerability&#x20;

A new vulnerability (up to 1 year) is more likely to be exploited than an old one (more than 1 year since publication)&#x20;

**Possible input values:**  _Less than 1 year old, Over 1 year old_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Less than 1 year old_ - Likelihood subscore will increase&#x20;

_Over 1 year old_ -  Likelihood subscore will decrease&#x20;
{% endhint %}

#### Package Popularity (Coming Soon)&#x20;

If a package is more popular (relative to its ecosystem), it is more likely to be exploited as hackers benefit from a wider pool of potential targets.&#x20;

**Possible input values:** Popularity percentile rank \[1-100]

{% hint style="info" %}
**How would this affect the score?**&#x20;

The likelihood subscore will increase based on the popularity rank.&#x20;
{% endhint %}

#### CVE disputed (Coming Soon)&#x20;

These are CVEs that have been acknowledged as being disputed by their Project maintainer or the community at large. Our research shows that none of the disputed CVEs in the Snyk Vulnerability DB have been exploited in the wild.

**Possible input values:** _True, False_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_True_ - Likelihood subscore will decrease significantly&#x20;

_False_ - Likelihood subscore will not change
{% endhint %}

### Contextual likelihood risk factors

#### Transitive Depth&#x20;

Building on [past studies](https://arxiv.org/pdf/2301.07972.pdf), Snyk’s research has shown that if a vulnerability is introduced to a Project transitively rather than directly, it’s less likely for an exploitable function path to exist

**Possible input values:** _Direct dependency, Indirect dependency, Great transitive depth (coming soon)_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Direct Dependency_ - Likelihood subscore will not change &#x20;

_Indirect Dependency_ - Likelihood subscore will decrease

_Great transitive depth -_ Likelihood subscore will decrease significantly
{% endhint %}

#### Reachability&#x20;

Snyk static code analysis determines whether the vulnerable method is being called. (Currently only supported in Java, JS coming soon. [Learn more](https://docs.snyk.io/manage-issues/issue-management/reachable-vulnerabilities).)&#x20;

**Possible input values:** _Reachable, No path found, Undefined_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Reachable_ - Likelihood subscore will increase, and Transitive Depth will not be considered.&#x20;

_Not reachable_ - Likelihood subscore will not change

_Undefined_ - Likelihood subscore will not change&#x20;
{% endhint %}

#### Platform Condition (Coming Soon)&#x20;

Whether the operating system and architecture of a given resource are relevant to this specific issue or not&#x20;

**Possible input values:** _Condition not met, Condition met, Undefined_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Condition not met_ - Likelihood subscore will decrease significantly

_Condition met, Undefined_ - Likelihood subscore will not change.&#x20;
{% endhint %}

#### Deployed (Coming Soon)&#x20;

Whether the container image introducing this issue is deployed or not&#x20;

**Possible input values:** Deployed, Not Deployed, Undefined

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Deployed_ - Likelihood subscore will increase

_Not Deployed_ - Likelihood subscore will decrease

_Undefined_ - Likelihood subscore will not change
{% endhint %}

#### Public Facing (Coming Soon)&#x20;

Whether the asset introducing this issue is exposed to the internet or not Public Facing Likelihood subscore will increase

**Possible input values:** _Public Facing, Not Public Facing, Undefined_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Public Facing_ - Likelihood subscore will increase

_Not Public Facing_ - Likelihood subscore will decrease

_Undefined_ - Likelihood subscore will not change
{% endhint %}

{% hint style="warning" %}
All factor names and their effect on the score are subject to change during the beta period
{% endhint %}

\
\
