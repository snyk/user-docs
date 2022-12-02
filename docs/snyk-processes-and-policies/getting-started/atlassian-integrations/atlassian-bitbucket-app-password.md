---
description: Setup a Bitbucket application password to enable integration with Snyk
---

# Atlassian Bitbucket App Password

The Atlassian App Password is required to enable an integration between Snyk and Atlassian Bitbucket. &#x20;

## Create an app password

You will need to [create an app password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/) in order to authorize Snyk to access your repository and enable Snyk's [Bitbucket Cloud integration](https://support.snyk.io/hc/en-us/articles/360004032097-Bitbucket-Cloud-integration).

To create an app password:

1. From your avatar in the bottom left, click **Personal settings**.
2. Click **App passwords** under **Access management**.
3. Click **Create app password**.
4. Give the app password a name related to the application that will use the password.
5. Select the specific access and permissions you want this application password to have.
6. Copy the generated password and either record or paste it into the application you want to give access. **The password is only displayed this one time**.

You will need the following permissions:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/bitbucket-api-token.png)

* Account: `read`
* Team membership: `read`
* Projects: `read`
* Repositories: `read and write`
* Pull requests: `read and write`
* Webhooks: `read and write`

