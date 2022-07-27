# Use the Snyk plugin to secure your Eclipse projects

Once the language server is downloaded and the authentication is completed, the plugin starts the workspace scan. You might notice a confirmation that a workspace scan is starting. Snyk shows such a notification when there is no workspace scan available.

![Starting workspace scan](<../../.gitbook/assets/Screenshot 2022-05-13 at 11.55.41.png>)

All of the issues found by Snyk are now integrated natively with Eclipse flows. Issues are shown in the **Problems** tab as shown in the following screenshot. There is a squiggly line indicating the issue while you code plus gutter icons to indicate where the issue is.

![Eclipse Problems tab](<../../.gitbook/assets/Screenshot 2022-05-13 at 12.20.26.png>)

Continue by following the instructions on the page for the type of scan you are doing:

* [SAST scanning results (SAST, Snyk Code)](https://docs.snyk.io/ide-tools/eclipse-plugin/sast-scanning-results-sast-snyk-code)
* [Misconfiguration scanning results (Snyk Infrastructure as Code)](https://docs.snyk.io/ide-tools/eclipse-plugin/misconfiguration-scanning-results-snyk-infrastructure-as-code)
* [Third party dependency scanning (SCA, Snyk Open Source)](https://docs.snyk.io/ide-tools/eclipse-plugin/third-party-dependency-scanning-sca-snyk-open-source)
