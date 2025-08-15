# Analysis results: Snyk Open Source

Snyk Open Source analysis shows vulnerabilities in your code with every scan. The scan runs in the background and is enabled by default.

In the **Problems** tab of the Open Source editor window, you can see all vulnerabilities found in your Project.

## Snyk Open Source editor window

The editor window shows security vulnerabilities in open-source modules while you code in JavaScript, TypeScript, and HTML. Receive feedback in line with your code, such as how many vulnerabilities a module that you are importing contains. The editor exposes only top-level dependency vulnerabilities; for the full list of vulnerabilities, refer to the side panel.

You can find security vulnerabilities in the npm packages you import and see the known vulnerabilities in your imported npm packages as soon as you require the information:

<figure><img src="../../../../.gitbook/assets/image (171).png" alt="Vulnerabilities in npm package"><figcaption><p>Vulnerabilities in npm package</p></figcaption></figure>

Code inline vulnerability counts are also shown in your `package.json` file:

<figure><img src="../../../../.gitbook/assets/image (170).png" alt="Results screen showing the vulnerability count"><figcaption><p>Results screen showing the vulnerability count</p></figcaption></figure>

You can find security vulnerabilities in your JavaScript packages from well-known Content Delivery Networks (CDNs). The extension scans any HTML files in your Projects and displays vulnerability information about the modules you include from your favorite CDN.

The following CDNs are supported:

* unpkg.com
* ajax.googleapis.com
* cdn.jsdelivr.net
* cdnjs.cloudflare.com
* code.jquery.com
* maxcdn.bootstrapcdn.com
* yastatic.net
* ajax.aspnetcdn.com

<figure><img src="../../../../.gitbook/assets/oss-editor-html (1).png" alt="Vulnerability from a CDN"><figcaption><p>Vulnerability from a CDN</p></figcaption></figure>

You can navigate to the most severe vulnerability by triggering the provided code actions. This opens a vulnerability window to show more details:

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-17 at 14.04.13.png" alt="Code actions" width="217"><figcaption><p>Code actions</p></figcaption></figure>

## Snyk Open Source vulnerability window

The Open Source Security (OSS) vulnerability window shows information about the vulnerable module.

* Links to external resources (CVE, CWE, Snyk Vulnerability DB) to explain the vulnerability in more detail
* CVSS score and exploit maturity
* Detailed path of how vulnerability is introduced to the system
* Summary of the vulnerability, with the remediation advice to fix it

<figure><img src="../../../../.gitbook/assets/image (172).png" alt="Snyk Open Source vulnerability window"><figcaption><p>Snyk Open Source vulnerability window</p></figcaption></figure>
