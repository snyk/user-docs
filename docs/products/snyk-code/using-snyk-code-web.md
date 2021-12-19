# Using Snyk Code (web)

## Configure Snyk Code

To use Snyk Code, you have to make sure it's enabled for your organization and then import/re-import your repositories via the UI or API.

{% hint style="info" %}
See [Getting started with Snyk Code](https://docs.snyk.io/getting-started/getting-started-snyk-products/getting-started-with-snyk-code) to enable Code and start scanning your projects.
{% endhint %}

Once you have Code enabled, the following will happen:

1. Every time a new repository is imported (or re-imported), if it contains any source code files, a "Code Analysis" project is created for it
2. You may now run Snyk Code tests against this organization via the CLI
3. You may now run Snyk Code tests against this organization via one of the IDE plugins

![](../../.gitbook/assets/screenshot\_2021-06-17\_at\_13.23.19.png)

## View project vulnerabilities

Use Snyk Code with the standard Snyk web interface to find and fix vulnerabilities in your code.

1. In your **Projects** area, select the project to open
2. Snyk Code displays information and vulnerability cards for that project:&#x20;

![](../../.gitbook/assets/snykcofe\_priority\_score.png)

Information available shows standard Snyk project information (see [Snyk projects](https://support.snyk.io/hc/en-us/sections/360004724958-Snyk-projects)), including:

* Snapshot information showing when the project was last tested.
* **Overview**, **History** and **Settings** information. For example, use the **History** section to view previous snapshots of projects.
* Filters on the left of the screen.

## Vulnerability card details

Each vulnerability card shows specific details about that vulnerability:

![](../../.gitbook/assets/snykcode\_issue\_card.png)

Card details include:

* The severity level (for example, **H** for high), the name (for example, **Cross-site Scripting**), and the [Priority Score](https://docs.snyk.io/fixing-and-prioritizing-issues/starting-to-fix-vulnerabilities/snyk-priority-score) value.
* The [CWE type:](https://cwe.mitre.org/data/index.html) click the link to view more information about that type of vulnerability.
* A snippet of your code showing the exact area that is vulnerable.
* A clear and helpful text description of the vulnerability.

See [Issue card information](https://docs.snyk.io/getting-started/introduction-to-snyk-projects/issue-card-information) for more details.

* Click **Ignore** to ignore this vulnerability (see [Ignore a vulnerability](using-snyk-code-web.md))
* Click **Full details** to view more information (see [View full details](using-snyk-code-web.md)).

## Ignore a vulnerability

{% hint style="info" %}
Ignore functionality is not offered for Open Source repositories.
{% endhint %}

Use the **Ignore** button to suppress specific warnings, ignoring suggestions for fixes for that issue. For example, you may have deliberately used hard coded passwords to test your routines in test code, or you are aware of an issue but have decided not to fix it.

After you click **Ignore**, you are prompted to provide more details:

![](../../.gitbook/assets/snykcode-ignore-pic2.png)

Select:

* Whether this is to be marked **Not vulnerable**, **ignore temporarily** (to suppress the message for now), or **ignore permanently** (won’t fix ever)
* Add any notes to yourself or your colleagues in the comment box; we recommend writing a quick explanation for your decision.
* Set a timer for how long to ignore the issue (14, 30, 60, or 90 days, or click **Ignore forever**).

Click **Save** to ignore this issue with the parameters selected. Afer you select to ignore an issue, it does not appear in scan results.

{% hint style="info" %}
There is a status selector for ignored issues on the left side, and you can include those (and remove or edit the ignore flag) if you want to review ignored results.
{% endhint %}

See [Ignoring issues not prioritized for your project](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/ignore-issues) for full details of using the Ignore function.

See [Ignoring issues in Snyk Code](https://docs.snyk.io/fixing-and-prioritizing-issues/issue-management/ignore-issues#Ignore-Snyk-Code) for specific details of how Snyk Code processes the Ignore function.

## View full details

Click **Full details** from a vulnerability card to view more information:

![Data flow page preview.](../../.gitbook/assets/data-flow.png)

![Fix analysis page preview.](../../.gitbook/assets/fix-analysis.png)

Full details include all the information in the vulnerability card, plus:

* **Data Flow**: This area helps you understand where the issue is and how it flows throughout your application by showing you the full taint flow of the issue in the code, from the source (the user input) to the sink (the operation that needs to receive clean input and can be exploited otherwise). In the above example, the developer has not sanitized the input, allowing an attacker to do a pass traversal to potentially access any file on the file system, including sensitive data such as password files.
* **Fix Analysis:** This area provides you with more information about the issue and how it could be fixed. Developers are able to see fix-related information, vulnerability overview information (understanding and approach), and fix examples for this vulnerability type.
* A link to the source file, which you can open to make changes directly (see [Open the source code file](using-snyk-code-web.md)).
* A full window showing the code affected, with specific lines highlighted to accompany the **Data flow** information.

## Open the source code file

1.  Click the code link to open the source code file itself (for example, in GitHub)

    ![](../../.gitbook/assets/link.png)
2. The file opens, showing you exactly where to fix the vulnerability (in this example, by adding the sanitation required to the input).
3. You can now make the fix as required, to address the vulnerability in your code.

![](../../.gitbook/assets/open-code2.png)

## Example: Cross-site Scripting (XSS)

This shows an example of a common vulnerability, Cross-site Scripting (XSS). XSS vulnerabilities allow attackers to compromise the interactions users have with your application, including gaining control over the application's functionality and data.

The vulnerability card shows key information about this vulnerability:

![](../../.gitbook/assets/snykcode\_issue\_card.png)

Click **Full details** to see more information about this vulnerability:

![](../../.gitbook/assets/xss-2.png)

(This example shows a non-sanitized, HTTP input flows into a **write** response returned by the server, so could be running malicious code.)

Cick the code link to open the source code file directly, then make changes to fix this vulnerability.
