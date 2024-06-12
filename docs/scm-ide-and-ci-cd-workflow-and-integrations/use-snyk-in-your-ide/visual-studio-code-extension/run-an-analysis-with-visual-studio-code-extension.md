# Run an analysis with Visual Studio Code extension

In the IDE, observe that the extension is already picking up the files and uploading them for analysis.

Snyk Open Source requires the Snyk CLI, so it is downloaded in the background.

Snyk Code analysis runs quickly without the CLI, so results may already be available. Otherwise, you see the following screen while Snyk scans your workspace for vulnerabilities and issues:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-16 at 14.53.38.png" alt="Snyk scan in progress"><figcaption><p>Snyk scan in progress</p></figcaption></figure>

Snyk analysis runs automatically when you open a folder or workspace.

* Snyk Code performs scans automatically on file saves (can be disabled by setting **Scanning Mode** to **manual**).
* Snyk Infrastructure as Code (Configuration) scans automatically on file saves (can be disabled by setting **Scanning Mode** to **manual**).

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-16 at 14.56.29.png" alt="Set Snyk scanning mode"><figcaption><p>Set Snyk scanning mode</p></figcaption></figure>

* Snyk Open Source does not automatically run on save by default, but you can enable auto scan in settings:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-16 at 14.58.23.png" alt="Enable auto scan for Snyk Open Source"><figcaption><p>Enable auto scan for Snyk Open Source</p></figcaption></figure>

{% hint style="info" %}
if you do not like to manually save while working, enable [AutoSave](https://code.visualstudio.com/docs/editor/codebasics#\_save-auto-save).
{% endhint %}

## Rescan

To trigger a scan manually, either Save or manually rescan using the rescan icon:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-16 at 15.00.13.png" alt="Rescan icon"><figcaption><p>Rescan icon</p></figcaption></figure>



If you need only ,the Code Quality, Code Security, Open Source Security or Configuration portion of the findings, you can disable the feature with the results you do not want to see or collapse the view:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-16 at 14.28.48.png" alt="Open Source results open"><figcaption><p>Open Source results open</p></figcaption></figure>
