# Priority Score vs Risk Score

The Snyk Risk score and Priority score are keys to security management. Both types of score help Organizations handle current threats and prepare for future vulnerabilities, leading to a more robust security framework.

The Priority and Risk Scores rank the issues and the urgency of fixing them. Both scores provide a number between 1 and 1000, where 1 means low severity and 1000 means high severity. Snyk uses these numbers to indicate the urgency of remediating a vulnerability.

{% hint style="info" %}
Risk score assesses the potential impact of vulnerabilities, prioritizing those with severe consequences.

Priority score helps teams quickly identify and address critical security vulnerabilities by ranking them based on urgency.
{% endhint %}

Risk Score and Priority Score are fundamental to vulnerability management. Risk Score predicts the long-term impact of a vulnerability. Priority Score assesses vulnerabilities based on their immediacy and likelihood of exploitation, prioritizing responses to potential threats that could immediately compromise security. In contrast, Risk Score guides strategic resource allocation to prevent long-term damage.

If you want to compare scores, ensure you are looking at the same type of score and for the same product. For example, you can compare a Risk score from Snyk Open Source with another Risk score from Snyk Open Source but cannot compare a Risk score from Snyk Open Source with another Risk score from Snyk Container.

By default, the Priority score is enabled for Snyk Container, Snyk Code, Snyk IaC, and Snyk Open Source issues. If the user enables the Risk score from Snyk Preview, then the scores are applied like this:

* Risk score is applied to Snyk Container and Snyk Open Source issues.
* Priority score is applied to Snyk Code and Snyk IaC issues.

## Priority score and Risk score comparison

You can use the following table to decide which type of score works best for your scenario.

|                      | Priority Score                                                                                         | Risk Score                                                                                             |
| -------------------- | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| **Availability**     | <ul><li>General availability</li></ul>                                                                 | <ul><li>Early access</li></ul>                                                                         |
| **Applicability**    | <ul><li>Vulnerability issues</li><li>License issues</li></ul>                                          | <ul><li>Vulnerability issues</li></ul>                                                                 |
| **Snyk coverage**    | <ul><li>Snyk Open Source</li><li>Snyk Code</li><li>Snyk Container</li></ul>                            | <ul><li>Snyk Open Source</li><li>Snyk Container</li></ul>                                              |
| **Coverage**         | <ul><li>Project view from Snyk Web UI</li><li>Reports view from Snyk Web UI</li><li>Snyk API</li></ul> | <ul><li>Project view from Snyk Web UI</li><li>Reports view from Snyk Web UI</li><li>Snyk API</li></ul> |
| **Integrations**     | <ul><li>Kubernetes</li></ul>                                                                           | NA                                                                                                     |
| **Assessment model** | <ul><li>Impact</li><li>Actionability</li></ul>                                                         | <ul><li>Impact</li><li>Likelihood</li></ul>                                                            |

## When and why to use the Priority Score

The Priority Score tool helps manage vulnerabilities by prioritizing the most urgent issues based on exploitability, ease of mitigation, and potential for exploitation. Use it to first address the critical vulnerabilities and if you want to prioritize the Snyk Code issues (since Risk score is not available for Snyk Code).

* Time-Sensitive Projects - When you are working on Projects with tight deadlines, it is important to address security concerns right away.
* Initial Triage - Priority score quickly identifies and mitigates immediate threats from multiple vulnerabilities.
* Resource Allocation - Teams can allocate security resources to promptly address urgent risks.

Priority Score helps your team quickly prioritize and address urgent Project vulnerabilities to reduce the risk of attacks.

The assessment model used by the Priority score focuses on two factors:

* Impact - Snyk analyses the possibility of a fix to address multiple vulnerabilities. The Priority score increases exponentially with the number of vulnerabilities addressed by a fix.
* Actionability - Snyk analyses how easy it is to remediate a vulnerability.

## When and why to use the Risk score

The Risk Score assesses security threats based on their potential impact and likelihood, allowing for more nuanced vulnerability management.

* Comprehensive Risk Assessment - Use the Risk Score to assess both exploitability and potential damage when evaluating vulnerabilities.
* Long-Term Security Planning - Identifies and prioritizes potential risks, assisting in planning future security infrastructure improvements.
* Stakeholder Communication - The Risk Score is a metric that communicates security risks to non-technical stakeholders, aiding informed decisions on security investments.

The Risk Score helps balance immediate threat mitigation with long-term security posture, creating a comprehensive approach to managing vulnerabilities.

The assessment model used by the Risk score focuses on two factors:

* Impact - the possibility of a fix to address multiple vulnerabilities.
* Likelihood - the probability for a vulnerability to be exploited in a user's code.

## Use cases

Gain a better understanding of when and how to use a specific type of score by going through the following use cases. These scenarios demonstrate the vital utility of both scores in refining the security strategy of your Organization.

These examples use the same set of filters for both Priority and Risk score.

### Priority score use case

Imagine a scenario where your Organization uses a widely popular web application platform, which has recently been identified as having several security vulnerabilities. Due to limited resources, you may not be able to address all issues right away. Prioritization scores are crucial. Your team patches critical vulnerabilities first to mitigate the most significant threats to the security of your Organization.

You can focus on using the Priority score if you also want to scan the source code from Snyk Code, besides Snyk Open Source and Snyk Container.\
The Priority score focuses on the impact and means of remediating a vulnerability.

Scan your source code and apply the following filters to your list of found vulnerabilities:

* Issue type: Vulnerabilities
* Severity: Critical
* Fixed in available: Yes
* Computed fixability: Fixable
* Exploit maturity: Mature

### Risk score use case

Let's assume that you are integrating a new third-party library into an existing application, and after a scan, you discover that the library has several vulnerabilities. Filter the vulnerabilities using the Risk score to determine which vulnerabilities pose the greatest threat.

Remember that Risk score must first be enabled from the [Snyk Preview](../../snyk-platform-administration/snyk-preview.md) screen and can only be applied to Snyk Open Source and Snyk Container.

Scan your source code and apply the following filters to your list of found vulnerabilities to prioritize remediation efforts:

* Issue type: Vulnerabilities
* Severity: Critical
* Fixed availability: Yes
* Computed fixability: Fixable
* Exploit maturity: Mature

After you apply the filters, you have a list of the most critical vulnerabilities that need to be fixed before safely integrating the third-party library with your application. Fixing these critical issues you are preventing a potential security breach and are also safeguarding the integrity of your application.

Given the high severity and the mature exploit, the risk score for this issue would be elevated, indicating an urgent need for action. In this scenario, organizations should prioritize patching this vulnerability immediately due to the high risk of exploitation.

This example demonstrates how risk scores guide decision-makers in prioritizing security efforts effectively, ensuring that the most critical vulnerabilities are addressed promptly.
