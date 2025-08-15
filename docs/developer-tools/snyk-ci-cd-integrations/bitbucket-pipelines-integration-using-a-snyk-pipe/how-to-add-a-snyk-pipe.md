# How to add a Snyk pipe

Follow these steps to add a Snyk pipe:

1. Add the Snyk pipe while creating your pipeline or while editing an existing pipeline. See the Bitbucket documentation for more information: [pipelines](https://confluence.atlassian.com/bitbucket/configure-bitbucket-pipelines-yml-792298910.html) and [pipes](https://support.atlassian.com/bitbucket-cloud/docs/pipes/). When adding the Snyk pipe, follow the guidelines in the remaining steps.
2. Use the Bitbucket pipeline editor to update the `.yml` file configuration, select the correct language, and use the Bitbucket Pipes build directory when adding the Snyk pipe.
3. Paste the Snyk pipe into the Bitbucket editor interface after all build steps. Build steps are commands such as these: `npm install / composer install / bundle install / dotnet restore / docker build`
4. Ensure you paste the pipe **before** a deployment step, such as `npm publish` or `docker push`.
5. Configure the mandatory variables **SNYK\_TOKEN** and **LANGUAGE.**
6. (Optional) Choose whether to fail the pipeline on vulnerabilities found with **DONT\_BREAK\_BUILD** and **SEVERITY\_THRESHOLD** if it is used, and consider enabling **MONITOR** . For more information, see [Snyk pipe parameters and values](snyk-pipe-parameters-and-values-bitbucket-cloud.md).
7. After Snyk is included in your pipeline commands, it looks for the manifest files in that repository, for example,`package.json`, `package-lock.json` , and performs the scan.

Results appear in the Bitbucket Pipelines output interface, similar to the following:

![Bitbucket Pipelines output interface - here the pipeline fails due to Snyk finding vulnerabilities](<../../../.gitbook/assets/Screenshot 2023-10-03 at 13.08.45.png>)

{% hint style="info" %}
If the build fails, even if **MONITOR** is set to **True**, Snyk does not continue to the monitor stage because no Projects are deployed until the build succeeds. To enable monitoring on snyk.io of Projects with vulnerabilities, set **DONT\_BREAK\_BUILD** to **True**. You can use **SEVERITY\_THRESHOLD** to tell the pipe the severity threshold for failing the pipe at the scanning stage. For more information, see [Snyk pipe parameters and values](snyk-pipe-parameters-and-values-bitbucket-cloud.md).
{% endhint %}
