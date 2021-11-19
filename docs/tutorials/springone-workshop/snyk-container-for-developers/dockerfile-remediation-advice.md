# Dockerfile fix advice

## Add Dockerfile fix advice to your image

Snyk can provide fix advice to help developers make good security decisions during the development process. Snyk can do this from the Snyk CLI or the Snyk UI. In our sample application **goof** we will configure the Snyk UI to display this information.

To get started select your container image from the projects page.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/container\_image\_snyk\_ui.png)

The Snyk UI informs users that it can provide fix advice and offers a link to configure our container image. Select the settings area link in the purple box or the upper right corner to configure your container image with its Dockerfile.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/screen-shot-2020-08-21-at-4.38.33-pm.png)

Configure the location of your Dockerfile using the source control repository for the container image.

{% hint style="info" %}
To use the Dockerfile fix feature in the Snyk UI you must have your source control repository configured in the Snyk UI. For details see [GitHub integration](../../../features/integrations/git-repository-scm-integrations/github-integration.md).
{% endhint %}

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/screen-shot-2020-04-18-at-1.52.23-pm.png)

Select the source control repository

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/screen-shot-2020-04-18-at-1.53.02-pm.png)

Select the location of the Dockerfile in your repository. The file typically lives at the root level and is called Dockerfile.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/screen-shot-2020-04-18-at-1.53.16-pm.png)

The Snyk UI will show its testing your container image.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/screen-shot-2020-04-18-at-2.16.19-pm.png)

Once the test is complete the Snyk UI shows the configured Dockerfile for your container image.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/container\_image\_goof\_dockerfile\_set.png)

The Snyk UI shows fix advice for your container image. The Snyk UI displays the base image, the number of vulnerabilities, and the severity of the vulnerabilities. It also offers minor, major, and alternative advice to let the developer make a better choice for the base image.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/image\_redmiation\_advice\_spc.png)
