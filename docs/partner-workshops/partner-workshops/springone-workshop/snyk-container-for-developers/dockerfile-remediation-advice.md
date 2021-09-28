# Dockerfile Remediation Advice

## Add Dockerfile remediation advice to your image

Snyk can provide remediation advice to help developers make good security decisions during the development process. Snyk can do this from the Snyk CLI or the Snyk UI. In our sample application **goof** we will configure the Snyk UI to display this information.

{% hint style="info" %}
For [details](https://support.snyk.io/hc/en-us/articles/360003946917-Test-your-container-images-with-our-CLI-tool) on using the Snyk CLI to review remediation advice.
{% endhint %}

To get started select your container image from the projects page.

![](../../../.gitbook/assets/container_image_snyk_ui.png)

The Snyk UI informs users that it can provide remediation advice and offers a link to configure our container image. Select the settings area link in the purple box or the upper right corner to configure your container image with its Dockerfile.

![](../../../.gitbook/assets/screen-shot-2020-08-21-at-4.38.33-pm.png)

Configure the location of your Dockerfile using the source control repository for the container image. 

{% hint style="info" %}
To use the Dockerfile remediation feature in the Snyk UI you must have your source control repository configured in the Snyk UI. For details see [configuring](https://support.snyk.io/hc/en-us/articles/360004032117-GitHub-scan-monitor-and-remediate#UUID-3a5c7195-5847-36b1-c55b-3801d71a4e8e) and [importing]() your project 
{% endhint %}

![](../../../.gitbook/assets/screen-shot-2020-04-18-at-1.52.23-pm.png)

Select the source control repository

![](../../../.gitbook/assets/screen-shot-2020-04-18-at-1.53.02-pm.png)

Select the location of the Dockerfile in your repository. The file typically lives at the root level and is called Dockerfile.

![](../../../.gitbook/assets/screen-shot-2020-04-18-at-1.53.16-pm.png)

The Snyk UI will show its testing your container image.

![](../../../.gitbook/assets/screen-shot-2020-04-18-at-2.16.19-pm.png)

Once the test is complete the Snyk UI shows the configured Dockerfile for your container image.

![](../../../.gitbook/assets/container_image_goof_dockerfile_set.png)

The Snyk UI shows remediation advice for your container image.  The Snyk UI displays the base image, the number of vulnerabilities, and the severity of the vulnerabilities. It also offers minor, major, and alternative advice to let the developer make a better choice for the base image.

![](../../../.gitbook/assets/image_redmiation_advice_spc.png)

