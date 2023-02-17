# Step 3: Create and scan a Snyk Cloud Environment for Azure (Web UI)

{% hint style="info" %}
**Recap**\
****You have created the Azure app registration, federated identity credential, and service principal for Snyk Cloud. Now you can create and scan a Snyk Cloud Environment.
{% endhint %}

## Process overview

1. In the Snyk Web UI **Add Azure Environment** modal where you downloaded the Terraform template or Bash script, enter your application ID in the **Application ID** field.
2. Optionally, enter an environment name. If one is not provided, Snyk will use your subscription name.\
   ![Enter Application ID section of the Add Azure Environment modal in Snyk Cloud](../../../../.gitbook/assets/snyk-cloud-onboard-azure-step-2.png)
3. Select **Approve and begin scan**.
4. You'll see a confirmation message: "Azure environment successfully added." Select **Add another environment** to return to the **Add Azure Environment** modal and onboard a new subscription, or select **Go to settings** if you are finished.

### What's next?

You can now view misconfiguration issues in the API or Snyk Web UI. See [Snyk Cloud issues](../../snyk-cloud-issues/) for more information.
