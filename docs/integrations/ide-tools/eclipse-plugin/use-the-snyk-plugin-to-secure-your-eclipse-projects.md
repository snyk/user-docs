# Use the Snyk plugin to secure your Eclipse projects

Once the language server is downloaded and the authentication is completed, the plugin starts the workspace scan. You might notice a confirmation that a workspace scan is starting. Alternatively, you can trigger a workspace scan from your project's context menu, or from the Snyk View.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-10-19 at 09.02.25 (1).png" alt="Starting workspace scan"><figcaption><p>Starting workspace scan</p></figcaption></figure>

All of the issues found by Snyk are now integrated natively with Eclipse flows. Issues are shown in the **Problems** tab as shown in the following screenshot. There is a squiggly line indicating the issue while you code plus gutter icons to indicate where the issue is.

<figure><img src="../../../.gitbook/assets/Screenshot 2022-05-13 at 12.20.26.png" alt="Eclipse Problems tab"><figcaption><p>Eclipse Problems tab</p></figcaption></figure>

Continue by following the instructions on the page for the type of scan you are doing:

* [SAST scanning results (SAST, Snyk Code)](https://docs.snyk.io/ide-tools/eclipse-plugin/sast-scanning-results-sast-snyk-code)
* [Misconfiguration scanning results (Snyk Infrastructure as Code)](https://docs.snyk.io/ide-tools/eclipse-plugin/misconfiguration-scanning-results-snyk-infrastructure-as-code)
* [Third party dependency scanning (SCA, Snyk Open Source)](https://docs.snyk.io/ide-tools/eclipse-plugin/third-party-dependency-scanning-sca-snyk-open-source)
