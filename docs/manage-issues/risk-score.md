# Risk Score

{% hint style="info" %}
Risk Score is currently in Open Beta for Snyk Open Source and Snyk Container. Use [Snyk Preview](https://docs.snyk.io/snyk-admin/manage-settings/snyk-preview) to replace the Priority Score with the new Risk Score. See [Snyk feature release process](../more-info/snyk-feature-release-process.md) for more details.
{% endhint %}

The Snyk Risk Score is a single value assigned to an issue, applied by automatic risk analysis for each security issue. Risk Score is based on the potential impact and likelihood of exploitability. Ranging from 0 to 1,000, the score represents the risk imposed on your environment and enables a risk-based prioritization approach.&#x20;

Since real risk is scarce, you should expect a significant drift in the distribution of scores, as can be seen in this example of Project score distributions:&#x20;

<figure><img src="../.gitbook/assets/image (2) (1) (1) (1).png" alt="Example Project scores distribution"><figcaption><p>Example Project scores distribution</p></figcaption></figure>

{% hint style="info" %}
As part of the Open Beta, the Risk Score replaces the Priority Score directly. See the [priority score docs](priorities-for-fixing-issues/priority-score.md) for how to interact with the Risk Score in the UI, API, and Reports, where the Risk Score is now introduced when enabled. Risk Score is not available in the CLI.&#x20;
{% endhint %}

{% hint style="info" %}
The Priority Score will be replaced with the Risk Score when Snyk Open Source and Snyk Container Projects are re-tested
{% endhint %}

{% hint style="warning" %}
Note that in the API, the relevant fields are still named with `priority.`When Risk Score is enabled, the scores and factors populated in these fields are based on the Risk Score model as part of the beta,&#x20;
{% endhint %}

## Explore the Risk Score by issue&#x20;

When looking at Issue card information, hover over the Risk Score to see the subscore and Risk Factors contributing to the score.

<div data-full-width="false">

<figure><img src="../.gitbook/assets/image (118) (2).png" alt="Risk Score tooltip" width="563"><figcaption><p>Risk Score tooltip showing Impact and Likelihood</p></figcaption></figure>

</div>

## About the Risk Score Model&#x20;

<figure><img src="../.gitbook/assets/matrix (2).png" alt="he Snyk Risk Score Model"><figcaption><p>The Snyk Risk Score Model</p></figcaption></figure>

The model that powers the Risk Score applies automatic risk analysis for each security issue based on the potential impact and likelihood of exploitability.

{% hint style="info" %}
The Risk Model results from extensive research conducted by the Snyk Security Data Science team and experienced security researchers. The model draws on expertise gained over the years in developing the [Snyk Vulnerability Database](https://security.snyk.io/).
{% endhint %}

### Impact subscore

Objective impact factors are the CVSS impact metrics, Availability, Confidentiality, Integrity, and Scope, calculated based on the CVSS impact subscore. For Container issues, Provider Urgency is also taken into account.&#x20;

The business criticality Project attribute will be taken into account as a contextual impact factor, increasing or decreasing the impact subscore. For more information, see [Project attributes](../snyk-admin/introduction-to-snyk-projects/project-attributes.md).

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
* (Forthcoming) Disputed vulnerability

Contextual likelihood factors then increase or decrease the likelihood subscore: &#x20;

* Reachability (Snyk Open Source Java only, JavaScript to be supported)&#x20;
* Transitive depth
* (Forthcoming) Insights such as `Deployed` , `OS condition` and `Public Facing`

{% hint style="info" %}
Fixability is no longer considered part of the Score Calculation, as the effort needed to mitigate a security issue does not affect the risk it imposes. To focus on actionable issues, use Fixability filters and then use the Risk Score to start with the riskiest issues.&#x20;
{% endhint %}

## Risk factors drill down

### Objective impact risk factors

#### Confidentiality

Represents the impact on customer’s data confidentiality, based on [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-3-1-Confidentiality-C).\
**Possible input values:** `None`, `Low`, `High`

#### Integrity

Represents the impact on customer’s data integrity, based on [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-3-2-Integrity-I).\
**Possible input values:** `None`, `Low`, `High`

#### Availability

Represents the impact of customer’s application availability based on [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-3-3-Availability-A).\
**Possible input values:** `None`, `Low`, `High`

#### Scope

Indicates whether the vulnerability can affect components outside of the target’s security scope, based on [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-2-Scope-S).\
**Possible input values:** `Unchanged`, `Changed`.&#x20;

{% hint style="info" %}
**How would these affect the score?** \
The objective impact subscore is calculated based on the CVSS impact subscore. For more information, see the references on CVSS definitions above and the [subscore equations](https://www.first.org/cvss/v3.1/specification-document#7-1-Base-Metrics-Equations).
{% endhint %}

#### Provider urgency (Snyk Container)&#x20;

Urgency rating as provided by the relevant operating system distribution security team. For more information, see [External information sources for relative importance](../scan-using-snyk/snyk-container/how-snyk-container-works/severity-levels-of-detected-linux-vulnerabilities.md#external-information-sources-for-relative-importance) in severity levels of detected Linux vulnerabilities.\
**Possible input values:** `Critical`, `High`, `Medium`, and `Low`_._ When neither CVSS nor Importance Rating is provided, Provider Urgency is set to `Low` by default.

{% hint style="info" %}
**How would this affect the score?** \
`Critical` - Impact subscore will increase significantly\
`High` - Impact subscore will increase\
`Medium` _-_ Impact subscore will decrease significantly\
`Low` _-_ Impact subscore will decrease significantly\
\
Note that Provider Urgency will also affect the Likelihood subscore.&#x20;
{% endhint %}

### Contextual impact risk factor

**Business criticality**&#x20;

User-defined Project attribute representing the subjective business impact of the respective application. For more information, see [Project attributes](../snyk-admin/introduction-to-snyk-projects/project-attributes.md).\
**Possible input values:** `Critical`, `High`, `Medium`, `Low`.

{% hint style="info" %}
**How would this affect the score?** \
`Critical` - Impact subscore will increase\
`High` - Impact subscore will not be affected \
`Medium` _-_ Impact subscore will decrease\
`Low` _-_ Impact subscore will decrease significantly\
\
When no Business Criticality is assigned, the Impact subscore will not be affected.&#x20;
{% endhint %}

{% hint style="info" %}
When you apply a business criticality attribute to a Project, a retest is required for the Risk Scores to incorporate the new data.&#x20;
{% endhint %}

### Objective likelihood risk factors

#### Exploit maturity&#x20;

Represents the existence and maturity of any public exploit retrieved and validated by Snyk. For more information, see [View exploits, How exploits are determined](priorities-for-fixing-issues/view-exploits.md#how-exploits-are-determined).\
**Possible input values:** `No Known Exploit`, `Proof of Concept`, `Functional`, `High`.

{% hint style="info" %}
**How would this affect the score?** \
`High` - Impact subscore will increase significantly\
`Functional` - Impact subscore will increase\
`Proof of Concept` _-_ Impact subscore will decrease slightly\
`No Known Exploit` _-_ Impact subscore will decrease significantly\
\
The likelihood subscore will increase significantly according to the level of exploit maturity
{% endhint %}

#### EPSS score&#x20;

Exploit Prediction Scoring System (EPSS), predicting whether a CVE would be exploited in the wild, based on an elaborated model created and owned by the FIRST Organization. \
The probability is the direct output of the EPSS model and conveys an overall sense of the threat of exploitation in the wild. This data is updated daily, relying on the latest available EPSS model version. See the EPSS [documentation](https://www.first.org/epss/articles/prob\_percentile\_bins) for more details.\
**Possible input values:** `EPSS score [0.00-1.00]`

{% hint style="info" %}
**How would this affect the score?** \
The likelihood subscore will increase significantly according to the EPSS score
{% endhint %}

#### Attack vector&#x20;

Represents the context by which vulnerability exploitation is possible, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-1-Attack-Vector-AV).\
**Possible input values:** `Network, Adjacent, Local, Physical`

{% hint style="info" %}
**How would this affect the score?**&#x20;

`Network` - Likelihood subscore will increase

_`Adjacent, Local, Physical`_ - Likelihood subscore will decrease according to the level of remote access needed to exploit the vulnerability
{% endhint %}

#### Attack complexity&#x20;

Represents the level of complexity defined by the conditions that must exist to exploit the vulnerability, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-2-Attack-Complexity-AC).\
**Possible input values:** `Low`, `High`

{% hint style="info" %}
**How would this affect the score?**&#x20;

`Low` - Likelihood subscore will increase

`High` - Likelihood subscore will decrease
{% endhint %}

#### Privileges required&#x20;

Represents the level of privileges an attacker must possess before successfully exploiting the vulnerability, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-3-Privileges-Required-PR).\
**Possible input values:** `None`, `Low`, `High`

{% hint style="info" %}
**How would this affect the score?**&#x20;

_`None`_- Likelihood subscore will increase&#x20;

`Low, High`- Likelihood subscore will decrease according to the level of privileges required&#x20;
{% endhint %}

#### User interaction&#x20;

Represents the need for action from a user as part of the exploitation process, based on the [CVSS definition](https://www.first.org/cvss/v3.1/specification-document#2-1-4-User-Interaction-UI).\
**Possible input values:** `None`, `Required`&#x20;

{% hint style="info" %}
**How would this affect the score?**&#x20;

`None`-  Likelihood subscore will increase&#x20;

`Required`- Likelihood subscore will decrease&#x20;
{% endhint %}

#### Social trends&#x20;

Represents the social media traffic regarding this vulnerability. Snyk research has shown that greater social media interaction can predict future exploitation or point to existing exploitation. For more information, see [Vulnerabilities with social trends](priorities-for-fixing-issues/vulnerabilities-with-social-trends.md).\
**Possible input values:**  `Trending`, `Not trending`

{% hint style="info" %}
**How would this affect the score?**&#x20;

`Trending` - Likelihood subscore will increase&#x20;

_`Not trending`_ - Likelihood subscore will not change
{% endhint %}

#### Malicious package&#x20;

Malicious code deployed as a supply chain dependency is considered highly exploitable.\
**Possible input values:** `True`, `False`

{% hint style="info" %}
**How would this affect the score?**&#x20;

The Likelihood subscore will increase significantly for Malicious Packages&#x20;
{% endhint %}

#### Age of vulnerability&#x20;

A new vulnerability (up to one year) is more likely to be exploited than an old vulnerability (more than one year since publication)&#x20;

**Possible input values:**  _Days since the vulnerability was first published._

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Less than one year old_ - Likelihood subscore will increase&#x20;

_Over one year old_ -  Likelihood subscore will decrease&#x20;
{% endhint %}

#### Package popularity (Snyk Open Source)&#x20;

If a package is relatively more popular for its ecosystem, it is more likely to be exploited as hackers benefit from a wider pool of potential targets.\
**Possible input values:** `High`, `Medium`, `Low` &#x20;

{% hint style="info" %}
**How would this affect the score?**&#x20;

`High` - Likelihood subscore will increase&#x20;

`Medium` - Likelihood subscore will not change &#x20;

`Low` -  Likelihood subscore will decrease
{% endhint %}

#### CVE disputed (forthcoming)&#x20;

These are CVEs that have been acknowledged as being disputed by their Project maintainer or the community at large. Snyk research shows that none of the disputed CVEs in the Snyk Vulnerability DB have been exploited in the wild.\
**Possible input values:** `True`, `False`

{% hint style="info" %}
**How would this affect the score?**&#x20;

`True` - Likelihood subscore will decrease significantly&#x20;

_`False`_ - Likelihood subscore will not change
{% endhint %}

#### Provider urgency (Snyk Container)&#x20;

Importance rating as provided by the relevant operating system distribution security team. For more information, see [External information sources for relative importance](../scan-using-snyk/snyk-container/how-snyk-container-works/severity-levels-of-detected-linux-vulnerabilities.md#external-information-sources-for-relative-importance) in severity levels of detected Linux vulnerabilities.\
**Possible input values:** `Critical`, `High`, `Medium`, and `Low`_._ When neither CVSS nor Importance rating is provided, provider urgency is set to `Low` by default.

{% hint style="info" %}
**How would this affect the score?** \
`Critical` - Impact subscore will increase significantly\
`High` - Impact subscore will increase\
`Medium` _-_ Impact subscore will decrease\
`Low` _-_ Impact subscore will decrease significantly\
\
Note that Provider Urgency will also affect the Impact subscore.&#x20;
{% endhint %}

### Contextual likelihood risk factors

#### Transitive depth&#x20;

Building on [past studies](https://arxiv.org/pdf/2301.07972.pdf), Snyk research has shown that if a vulnerability is introduced to a Project transitively rather than directly, it is less likely that an exploitable function path will exist.\
**Possible input values:** `Direct dependency`, `Indirect dependency`, `Great transitive depth` _(forthcoming)_

{% hint style="info" %}
**How would this affect the score?**&#x20;

_Direct Dependency_ - Likelihood subscore will not change &#x20;

_Indirect Dependency_ - Likelihood subscore will decrease

_Great transitive depth -_ Likelihood subscore will decrease significantly (coming soon)
{% endhint %}

#### Reachability&#x20;

Snyk static code analysis determines whether the vulnerable method is being called. This is currently supported only in Java; JavaScript support is coming soon. For more information, see [Reachable vulnerabilities](priorities-for-fixing-issues/reachable-vulnerabilities.md).\
**Possible input values:** `Reachable`, `No path found`\
When Reachability is not enabled, the Likelihood subscore will not change, and the factor will not show up.

{% hint style="info" %}
**How would this affect the score?**&#x20;

`Reachable` - Likelihood subscore will increase, and transitive depth will not be considered.&#x20;

_`No path found`_ - Likelihood subscore will not change
{% endhint %}

#### OS condition (forthcoming)&#x20;

Whether or not the operating system and architecture of a given resource are relevant to this specific issue  \
**Possible input values:** `Condition not met`, `Condition met`, `No data`

{% hint style="info" %}
**How would this affect the score?**&#x20;

`Condition not met` - Likelihood subscore will decrease significantly

`Condition met, No data` - Likelihood subscore will not change.&#x20;
{% endhint %}

#### Deployed (forthcoming)&#x20;

Whether or not the container image introducing this issue is deployed. \
**Possible input values:** `Deployed`, `Not Deployed`, `No data`

{% hint style="info" %}
**How would this affect the score?**&#x20;

`Deployed` - Likelihood subscore will increase

`Not Deployed` - Likelihood subscore will decrease

`No data` - Likelihood subscore will not change
{% endhint %}

#### Public facing (forthcoming)&#x20;

Whether or not the asset introducing this issue is exposed to the Internet.

**Possible input values:** `Public Facing`, `Not Public Facing`, `No data`

{% hint style="info" %}
**How would this affect the score?**&#x20;

`Public Facing` - Likelihood subscore will increase

`Not Public Facing` - Likelihood subscore will decrease

`No data` - Likelihood subscore will not change
{% endhint %}

{% hint style="warning" %}
All factor names and their effect on the score are subject to change during the beta period.&#x20;
{% endhint %}

\
\
