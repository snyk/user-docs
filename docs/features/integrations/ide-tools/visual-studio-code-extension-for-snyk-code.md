# Visual Studio Code extension

{% hint style="warning" %}
Snyk's Visual Studio Code extension is available for install on the marketplace: [https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner).

Visual Studio Code requires Snyk CLI, see [install-the-snyk-cli](../../snyk-cli/install-the-snyk-cli/ "mention").
{% endhint %}

### Supported languages, package managers and frameworks

* For Snyk Open Source: the VS Code extension support all the languages and package managers supported by Snyk Open Source and the CLI. See the full list [here](https://docs.snyk.io/products/snyk-open-source/language-and-package-manager-support).
* For Snyk Code: the VS Code extension support all the [languages and frameworks supported by Snyk Code](https://docs.snyk.io/products/snyk-code/snyk-code-language-and-framework-support#language-support-with-snyk-code-ai-engine) today.

### Install the extension

You can find the Snyk Extension in the Visual Studio Code Marketplace. To install, either:

* Navigate to the [Snyk Extension on the Visual Studio Code Marketplace](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner) and follow the instructions for the Snyk extension. The docs from VS Code help you trigger the installation process from Visual Studio Code and guide you through the installation steps.
* Browse for the extension as advised [here](https://code.visualstudio.com/docs/editor/extension-gallery#\_browse-for-extensions) and search for Snyk, then install (as described [here](https://code.visualstudio.com/docs/editor/extension-gallery#\_install-an-extension)).

Once installed you can find a Snyk icon in the sidebar ![](<../../../.gitbook/assets/Screen Shot 2021-12-03 at 8.02.07 AM.png>).

Snyk’s extension provides all the suggestions in a concise and clean view containing all information you need to decide how to fix or act upon:

![](<../../../.gitbook/assets/image (82) (1).png>)

### Configuration

#### Environment

To analyze projects, the plugin uses the Snyk CLI which needs some environment variables:&#x20;

* `PATH`: the path to needed binaries, (for example, to maven).
* `JAVA_HOME`: the path to the JDK you want to use to analyze Java dependencies&#x20;

Setting these variables only in a shell environment (for example,using `~/.bashrc`) is not sufficient, if you do not start the IDE from the command line or create a script file that starts the IDE using a shell environment.&#x20;

* On `Windows`, you can set the variables, using the GUI or on the command line using the `setx` tool.&#x20;
* On `macOS`, the process `launchd` must know the environment variables if you want to launch the IDE from Finder directly. You can set environment variables for applications launched via Finder using the `launchctl setenv` command e.g. on start-up or via a script you launch at user login. \
  **Note:** The provision of environment variables to the macOS UI can change between operating system releases, so it might be easier to create a small shell script that launches the IDE to leverage the shell environment, that can be defined via `~/.bashrc`.&#x20;
* On `Linux`, updating the file /etc/environment can be used to propagate the environment variables to the windows manager and UI.

#### Proxy

If you are behind a proxy, please configure it in the Visual Studio Code settings to be able to use the plugin.&#x20;

### Authentication

To authenticate follow the steps:

1.  Once the extension is installed, click on the Snyk Icon in the left navigation bar:

    ![](<../../../.gitbook/assets/image (67) (1) (1).png>)
2.  Click **Connect VS Code with Snyk**. The extension relies on the Snyk authentication API and it will ask you to authenticate you against Snyk’s web application:

    ![](<../../../.gitbook/assets/image (71) (1) (1).png>)
3. Click **Authenticate**.&#x20;
4.  After successful authentication, you will see a confirmation message:&#x20;

    ![](<../../../.gitbook/assets/image (85) (1).png>)
5. Close the browser window and return to VS Code. VS Code is now reading and saving the authentication on your local machine.

### Run analysis

In the IDE you will notice that the extension is already picking up the files and uploading them for analysis.

Snyk Open Source requires Snyk CLI, so it will proceed with the download in the background.

Snyk Code analysis runs quickly without it, so results may even already be available. Otherwise, you will see the following screen while Snyk scans your workspace for vulnerabilities:

![](<../../../.gitbook/assets/image (80) (1) (1) (1).png>)

Snyk's analysis runs automatically when you open a folder or workspace.

* Snyk Code performs scans automatically on file saves.
* Snyk Open Source does not automatically run on save by default, but you can enable it in settings:

![](<../../../.gitbook/assets/image (73) (1) (1) (1).png>)

**Tip**: if you don't like to manually save while working, enable [AutoSave](https://code.visualstudio.com/docs/editor/codebasics#\_save-auto-save).

### Rescan

To manually trigger a scan, either Save or manually rescan using the rescan icon:

![](<../../../.gitbook/assets/image (61) (1) (1).png>)

If you only need the Code Quality, Code Security or Open Source Security portion of the findings, you can easily disable the feature with the results you don't want to see or simply collapse the view:

![](<../../../.gitbook/assets/image (83) (1) (1) (1) (2).png>)

#### Snyk Code advanced mode

Snyk Code has "Advanced" mode that allows you to control how scan is performed.

To manually perform the analysis, in the configuration of the extension, you can enable Advanced Mode which enables you to control the scanning process:

![](<../../../.gitbook/assets/image (78) (1) (1) (1) (1).png>)

### View analysis results

Snyk analysis shows a list of security vulnerabilities and code issues found in the application code. For more details and examples of how others fixed the issue, select a security vulnerability or a code security issue. Once selected you will see the Snyk suggestion information in a panel on the right side:

![](<../../../.gitbook/assets/image (79) (2).png>)

![](<../../../.gitbook/assets/image (72) (1) (1) (2).png>)

#### Snyk panel

The Snyk analysis panel (on the left of the code screen in the above screenshot) shows how much time the analysis took plus a list of issues with the suggestions found for them.

The icons here mean:

| ![](<../../../.gitbook/assets/image (75) (1) (1).png>) Critical severity        | May allow attackers to access sensitive data and run code on your application.                                                               |
| ------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| ![](<../../../.gitbook/assets/image (77) (1) (1) (1) (1).png>) High severity    | May allow attackers to access sensitive data on your application.                                                                            |
| ![](<../../../.gitbook/assets/image (70) (1) (1) (1).png>) Medium severity      | May allow attackers under some conditions to access sensitive data on your application.                                                      |
| ![](<../../../.gitbook/assets/image (81) (1) (1) (1) (1) (1).png>) Low severity | The application may expose some data allowing vulnerability mapping, which can be used with other vulnerabilities to attack the application. |

You can filter the issues by setting the severities you want to see using the `snyk.severity` setting. E.g. set `"snyk.severity": { "critical": true, "high": true, "medium": true, "low": false }` to hide low severity issues. You can also apply the setting via Settings UI.

![](<../../../.gitbook/assets/image (65) (1) (1).png>)

#### Snyk Code editor window

The editor window (in the middle of the results screen) shows the code that is inspected. This ensures that when you are inspecting a Snyk issue, you always have the code context close to the issue.

#### Snyk Code vulnerability window

![](<../../../.gitbook/assets/image (76) (1) (1) (1) (1) (1).png>)

The Snyk Suggestion panel (on the right of the results screen) shows the argumentation of the Snyk engine using for example variable names of your code and the line numbers in red. You can also see:

* Links to external resources to explain the bug pattern in more detail (the **More info** link).&#x20;
* Tags that were assigned by Snyk, such as **Security** (the issue found is a security issue), **Database** (it is related to database interaction), or **In Test** (the issue is within the test code).&#x20;
* Code from open source repositories that might be of help to see how others fixed the issue.&#x20;
* You can add ignore comments that would make Snyk ignore this particular suggestion, or all of these suggestions for the whole file, by using the two buttons on the lower end of the panel.

We also include a feedback mechanism to report false positives so you others do not see the same issue.

#### Snyk Open Source editor window

Editor window shows security vulnerabilities in open source modules while you code in JavaScript, TypeScript and HTML. Receive feedback in-line with your code, such as how many vulnerabilities a module contains that you are importing. Editor surfaces only top-level dependency vulnerabilities, for the full list of vulnerabilities refer to the side panel.

* Find security vulnerabilities in the npm packages you import: see the number of known vulnerabilities in your imported npm packages as soon as you require them:

![](../../../.gitbook/assets/oss-editor-vulnerability-count.png)

* Code inline vulnerability counts are also shown in your `package.json` file:&#x20;

![](../../../.gitbook/assets/oss-editor-pjson.png)

* Find security vulnerabilities in your JavaScript packages from well-known CDNs: the extension scans any HTML files in your projects and displays vulnerability information about the modules you include from your favorite CDN.
  * Currently supported CDN's:
    * unpkg.com
    * ajax.googleapis.com
    * cdn.jsdelivr.net
    * cdnjs.cloudflare.com
    * code.jquery.com
    * maxcdn.bootstrapcdn.com
    * yastatic.net
    * ajax.aspnetcdn.com

![](<../../../.gitbook/assets/oss-editor-html (1).png>)

You can navigate to the most severe vulnerability by triggering the provided code actions. This opens a vulnerability window to show more details:

![](<../../../.gitbook/assets/oss-editor-show-vulnerability (2).png>)

#### Snyk Open Source vulnerability window

![](<../../../.gitbook/assets/image (68) (1) (1) (1).png>)

OSS vulnerability tab shows information about the vulnerable module.

* Links to external resources (CVE, CWE, Snyk Vulnerability DB) to explain the vulnerability in more detail.
* Displays CVSS score and exploit maturity.
* Provides detailed path on how vulnerability is introduced to the system.
* Shows summary of the vulnerability together with the remediation advice to fix it.

### Extension configuration

After the extension is installed, you can set the following configurations for the extension:

* **Token**: the token the extension uses to connect to Snyk. You can manually replace it, if you need to switch to another account.
* **Features**
  * **Code Security**: configures if code security analysis should run over your code.
  * **Code Quality**: configures if code quality analysis should run over your code.
  * **Open Source Security**: configures if security analysis should run over your open source dependencies.
* **Severity**: sets severity level to display in the analysis result tree.
* **Advanced**
  * **Advanced mode**: toggles a panel to allow the user to manually control when the analysis should be run.
  * **Auto Scan Open Source Security**: sets severity level to display in the analysis result tree.
  * **Additional Parameters**: sets parameters to be passed to Snyk CLI for Open Source Security tests. For the full list you can consult [this reference](https://docs.snyk.io/features/snyk-cli/guides-for-our-cli/cli-reference).

#### Create a .dcignore file

To ignore certain files and directories (for example, **node\_modules**), create a **.dcignore** file. You can create it in any directory on any level starting from the directory where your project resides. The file syntax is identical to .gitignore.

* We recommend adding the file when there is no **.gitignore** file. This will significantly reduce the files that need to be uploaded and speed up the analysis.
* To quickly add the default **.dcignore** file use the command provided by VS Code and the Snyk extension: Snyk create .dcignore file and save the newly created .dcignore file.

### Support / Contact

{% hint style="info" %}
Need more help? [Contact our Support team](https://support.snyk.io/hc/en-us/requests/new).
{% endhint %}

#### Share your experience

We continuously strive to improve our plugins experience. Would you like to share with us your feedback about Snyk's Visual Studio Code extension: [schedule a meeting](https://calendly.com/snyk-georgi/45min?month=2022-01).
