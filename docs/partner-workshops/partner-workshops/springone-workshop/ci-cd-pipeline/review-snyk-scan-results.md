# Review CI/CD results

## View CI/CD results

Now that we have added our Secrets and configured the environment variables in the pipeline let us review the results. Select the last run of the action, and let us examine the Snyk scan results.

![](../../../.gitbook/assets/screen-shot-2020-08-22-at-11.28.15-am.png)

## Snyk scan results

Snyk performed a Synk monitor command based on the configuration of the action. In Snyk, you can choose to test or monitor your application for vulnerabilities and license compliance. Monitoring saves the results of your scan and performs continuous monitoring for new vulnerabilities and license issues. In the output of the Snyk action, you can see the URL for the saved result. You can also use the Snyk UI to review the results, which we will do later in the Synk CI/CD section.

In this GitHub pipeline, we added a GitHub Snyk action to perform the monitor in parallel with the Maven build process. The results were uploaded to your Snyk organization. We will review the results in Snyk later in this workshop.

![](../../../.gitbook/assets/screen-shot-2020-08-22-at-11.52.35-am.png)

