# View analysis results from Visual Studio Code extension

Snyk analysis shows a list of security vulnerabilities and code issues found in the application code. For more details and examples of how others fixed the issue, select a security vulnerability or a code security issue. **Snyk suggestion information** for the issue selected appears in a panel on the right side, including the **Description**, **Impact** statement, **Path** by which the issue was introduced, and suggested **Remediation**:

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-16 at 16.07.06.png" alt="Snyk analysis panel and Snyk suggestion information for the issue selected"><figcaption><p>Snyk analysis panel and Snyk suggestion information for the issue selected</p></figcaption></figure>

The **Snyk analysis panel** on the left on the preceding screen shows how much time the analysis took and a list of issues with the suggestions found for them.

The icons have the following meaning:

| ![](<../../../../.gitbook/assets/image (201) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png>) Critical severity                                                                                                    | May allow attackers to access sensitive data and run code on your application.                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| ![](<../../../../.gitbook/assets/image (10) (1) (1) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (5) (3).png>) High severity | May allow attackers to access sensitive data on your application.                                                                            |
| ![](<../../../../.gitbook/assets/image (116) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (5) (6).png>) Medium severity                                          | May allow attackers under some conditions to access sensitive data on your application.                                                      |
| ![](<../../../../.gitbook/assets/image (114) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png>) Low severity                                                                                                         | The application may expose some data allowing vulnerability mapping, which can be used with other vulnerabilities to attack the application. |

You can filter the issues by setting by using the `snyk.severity` setting to show severities you want to see. For example, set `"snyk.severity": { "critical": true, "high": true, "medium": true, "low": false }` to hide low severity issues. You can also apply the setting in the Settings UI.

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-16 at 16.05.27.png" alt="Snyk severity settings"><figcaption><p>Snyk severity settings</p></figcaption></figure>
