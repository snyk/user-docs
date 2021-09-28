# Bitbucket Configuration

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

![](../../../../.gitbook/assets/bitbucket-api-token.png)

* Account: `read`
* Team membership: `read`
* Projects: `read`
* Repositories: `read and write`
* Pull requests: `read and write`
* Webhooks: `read and write`

## Repository variables

You will need to define [repository variables](https://support.atlassian.com/bitbucket-cloud/docs/variables-in-pipelines/#Repository-variables) at the repository level which will later be [referenced in your pipeline](https://support.atlassian.com/bitbucket-cloud/docs/variables-in-pipelines/).

These will consist of the following:

![](../../../../.gitbook/assets/bitubucket-repo-vars.png)

1. Amazon EKS name of your cluster: `AWS_EKS_CLUSTER`
2. Snyk API token for authenticating with your Snyk account: `SNYK_TOKEN`
3. AWS Identity & Access Management User [key and secret](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html) for secure authenticated interactions with the AWS API: `AWS_ACCESS_KEY_ID` & `AWS_SECRET_ACCESS_KEY`
4. AWS region you will be deploying to: `AWS_DEFAULT_REGION`
5. Amazon ECR [URL](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html) for your repository: `AWS_ECR_URI`
6. Container image name: `IMAGE`

{% hint style="info" %}
It is recommended that you use [Snyk Service accounts](https://support.snyk.io/hc/en-us/articles/360004037597-Service-accounts) and [AWS IAM best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html) when creating accounts.
{% endhint %}

