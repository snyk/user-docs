# Snyk Configuration

## Obtain your Snyk API token

From the Snyk console, navigate to **Settings** and under the **General** menu `Copy` your **Organization ID**.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-api-token.png)

Once you have copied your token, you will need to go back to the Bitbucket Cloud UI and define the `SNYK_TOKEN` **repository variable**.

## Enable Bitbucket integration

From the Snyk console, navigate to **Integrations** and select **Bitbucket Cloud**.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-integrations-menu.png)

From the **Bitbucket Cloud** integration page, enter your **Bitbucket username** in the **Username** field and the **Bitbucket app password** from the previous step in the **App password** field. Then, click `Save`.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-bitbucket-integration-01.png)

Once you have successfully connected your Snyk and Bitbucket accounts you will see a confirmation message and the ability to **Add your Bitbucket Cloud repository to Snyk**. Click the button to proceed.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-bitbucket-integration-02.png)

Find the repository you forked in the \[**Configure Environment**\]\(\) module. Click the checkbox to select it then click the **Add selected repository** button to import your project.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-bitbucket-add-repo.png)

Let's proceed to the next section.

