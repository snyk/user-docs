# Use the Snyk plugin to secure your Eclipse projects

Once the language server is downloaded and the authentication is completed, the plugin starts the workspace scan. You might notice a confirmation that a workspace scan is starting. Alternatively, you can trigger a workspace scan from your project's context menu, or from the Snyk View.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-10-19 at 09.02.25 (1).png" alt="Starting workspace scan"><figcaption><p>Starting workspace scan</p></figcaption></figure>

All of the issues found by Snyk are now integrated natively with Eclipse flows. Issues are shown in the **Problems** tab as shown in the following screenshot. There is a squiggly line indicating the issue while you code plus gutter icons to indicate where the issue is.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-05-13 at 12.20.26.png" alt="Eclipse Problems tab"><figcaption><p>Eclipse Problems tab</p></figcaption></figure>

Additionally, starting with version 3+, we provide a custom UI in the Snyk Tab, that displays issue details:

<figure><img src="../../../.gitbook/assets/image (641).png" alt=""><figcaption></figcaption></figure>

### Net new issue scanning (delta scanning)

#### General



<figure><img src="../../../.gitbook/assets/image (647).png" alt=""><figcaption><p>Net new issues can be activated in the dot menu of the Snyk View</p></figcaption></figure>

To filter the displayed issues to only display the issues introduced in the working directory, net new issues scan can be activated. This requires a git repository to work:

1. First, we scan the git `master` or `main` branch, or any reference branch selected in the Snyk View.
2. Second, we scan the working directory
3. We calculate the difference between both and only display the difference

#### Choosing the reference branch

1. Click on the project node, to open the branch chooser dialog.
2. Choose the branch that is the reference against which new issues shall be calculated

<figure><img src="../../../.gitbook/assets/image (653).png" alt=""><figcaption><p>After clicking on the project branch, the reference branch can be chosen</p></figcaption></figure>



Continue by following the instructions on the page for the type of scan you are doing:

* [SAST scanning results (SAST, Snyk Code)](https://docs.snyk.io/ide-tools/eclipse-plugin/sast-scanning-results-sast-snyk-code)
* [Misconfiguration scanning results (Snyk Infrastructure as Code)](https://docs.snyk.io/ide-tools/eclipse-plugin/misconfiguration-scanning-results-snyk-infrastructure-as-code)
* [Third party dependency scanning (SCA, Snyk Open Source)](https://docs.snyk.io/ide-tools/eclipse-plugin/third-party-dependency-scanning-sca-snyk-open-source)
