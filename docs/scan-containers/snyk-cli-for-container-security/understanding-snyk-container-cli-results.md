# Understanding Snyk Container CLI results

## **Vulnerability information**

When Snyk Container detects vulnerabilities, they are presented in the output:

<figure><img src="../../.gitbook/assets/clivulnerabiilities.png" alt="Vulnerabilities detected with Snyk Container"><figcaption><p>Vulnerabilities detected with Snyk Container</p></figcaption></figure>

Each vulnerability is shown with the following information:

| **Field**              | **Description**                                                                                                                                                              |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Severity**           | The importance of the specific vulnerability. See [Snyk Container severity](../how-snyk-container-works/understanding-linux-vulnerability-severity.md) for more information. |
| **A clear heading**    | Summary of the vulnerability and the affected package                                                                                                                        |
| **Description**        | A brief description of the type of issue or Common Vulnerabilities and Exposure (CVE) reference if a CVE exists                                                              |
| **Info**               | A link to more details about the vulnerability, including links to upstream sources and global vulnerabilities databases                                                     |
| **Introduced through** | The top-level package names affected by the vulnerability                                                                                                                    |
| **From**               | How the affected packages came to be in the image                                                                                                                            |
| **Introduced by**      | Whether the vulnerability is in the base image or which line in the Dockerfile introduced the vulnerability                                                                  |
| **Fixed in**           | If available, the version of the package which has a fix for the vulnerability                                                                                               |

Vulnerabilities appear in reverse severity order to limit scrolling to see the most important issues.

Snyk also reports the total dependencies tested for known vulnerabilities and the total number of vulnerabilities.

<figure><img src="../../.gitbook/assets/clisummary.png" alt="Total dependencies tested and issues fount"><figcaption><p>Total dependencies tested and issues fount</p></figcaption></figure>

{% hint style="info" %}
Snyk groups the same vulnerability discovered in multiple different packages together. This helps you focus on the number of vulnerabilities rather than the instances only.
{% endhint %}

## Base image recommendations

If Snyk determines the base image used, and the image uses an [Official Docker image](https://docs.docker.com/docker-hub/official\_images/), the output includes recommendations for upgrades to resolve some of the discovered vulnerabilities.

<figure><img src="../../.gitbook/assets/clirecommendations.png" alt="Recommendations for base image upgrade"><figcaption><p>Recommendations for base image upgrade</p></figcaption></figure>

This provides a level of situational awareness, showing the vulnerability counts in minor and major upgrades or in alternative base images, which might have fewer vulnerabilities.
