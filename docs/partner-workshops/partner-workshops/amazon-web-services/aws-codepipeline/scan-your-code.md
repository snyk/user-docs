# Scan your code

Snyk will scan your code each time the pipeline is triggered. This will take place for any commits on your configured repository, but you can also manually trigger it with a release.

![](../../../.gitbook/assets/snyk-codepipeline-15.png)

Click on the **Release change** button to trigger the pipeline.

![](../../../.gitbook/assets/snyk-codepipeline-16.png)

You will see the **Scan** stage run where Snyk consumes your artifacts to find potential vulnerabilities.

![](../../../.gitbook/assets/snyk-codepipeline-17.png)

Within the **Scan** stage, you will notice two clickable links. These correspond to the integration configuration and the findings report. You will have the ability to reconfigure your integrations for each configured pipeline without having to go through the entire process to enable this integration by clicking the link above. 

Similarly, once the stage completes, you can review your findings report by clicking on the **Details** link.

![](../../../.gitbook/assets/snyk-codepipeline-18.png)

A detailed findings report is available to you within the AWS CodePipeline console. The report provides the total count of known vulnerabilities, dependency paths, and detailed information on how these were introduced.

That's it! You've enabled Snyk's integration to AWS CodePipeline and are ready to take the first step to improve your security posture!

To learn more about Snyk and our [partnership with AWS](https://snyk.io/partners/aws/), please visit our partner page.

