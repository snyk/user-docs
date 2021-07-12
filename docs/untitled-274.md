# Using Snyk Code \(web\)

* [ Introduction to Snyk Code](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360017059758-Introduction-to-Snyk-Code/README.md)
* [ Snyk Code language and framework support](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360016973477-Snyk-Code-language-and-framework-support/README.md)
* [ Using Snyk Code \(web\)](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360017147558-Using-Snyk-Code-web-/README.md)

## Using Snyk Code \(web\)

Use Snyk Code with the standard Snyk web interface to find and fix vulnerabilities in your code.

### View project vulnerabilities

1. In your **Projects** area, select the project to open: 
2. Snyk Code displays information and vulnerability cards for that project:

Information available shows standard Snyk project information \(see [Snyk projects](https://support.snyk.io/hc/en-us/sections/360004724958-Snyk-projects)\), including:

* Snapshot information showing when the project was last tested.
* **Overview**, **History** and **Settings** information. For example, use the **History** section to view previous snapshots of projects.
* Severity and status filters on the left.

#### Vulnerability card details

Each vulnerability card shows specific details about that vulnerability:

Vulnerability card details include:

* The level \(for example, **H** for high\) and type \(for example, **Path Traversal**\).
* The [CWE type:](https://cwe.mitre.org/data/index.html) click the link to view more information about that type of vulnerability.
* A snippet of your code showing the exact area that is vulnerable.
* A clear and helpful text description of the vulnerability.

Click **Full details** from a vulnerability card to view more information. See [View full details](untitled-274.md).

Click **Full details** from a vulnerability card to view more information:

Full details include all the information in the vulnerability card, plus:

* **Data flow**: this area on the left shows the full taint flow of the issue in the code, from the source \(the user input\) to the sink \(the operation that needs to receive clean input and can be exploited otherwise\). In the above example, the developer has not sanitized the input, allowing an attacker to do a pass traversal to potentially access any file on the file system, including sensitive data such as password files.
* **Fix analysis:** insight into the remediation and background of the issue itself. Developers are able to see fix-related information, vulnerability overview information \(understanding and approach\), and fix examples for this vulnerability type.
* A link to the source file, which you can open to make changes directly \(see [Open the source code file](untitled-274.md)\).
* A full window showing the code affected, with specific lines highlighted to accompany the **Data flow** information.

### Example: Cross-site Scripting \(XSS\)

This shows an example of a common vulnerability, Cross-site Scripting \(XSS\). XSS vulnerabilities allow attackers to compromise the interactions users have with your application, including gaining control over the application's functionality and data.

The vulnerability card shows key information about this vulnerability:

Click **Full details** to see more information about this vulnerability:

\(This example shows that an unsanitized HTTP input flows into a **write** response returned by the server, so could be running malicious code.\)

Cick the code link to open the source code file directly, then make changes to fix this vulnerability.

