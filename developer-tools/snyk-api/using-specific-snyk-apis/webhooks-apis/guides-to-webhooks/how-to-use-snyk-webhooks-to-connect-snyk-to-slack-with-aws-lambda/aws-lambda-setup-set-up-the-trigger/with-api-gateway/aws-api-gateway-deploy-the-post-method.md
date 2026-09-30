---
description: How to deploy the POST method in AWS API Gateway for the Snyk to Slack integration
nav_context: agnostic
---

# AWS API Gateway: deploy the POST method

Deploy with configured POST method so the AWS Lambda function can start receiving the information.

Follow these steps to deploy the POST method:

1. Go to the **Resources** tab.
2. Click **POST**.
3. On the **Actions** tab, click **Deploy API**.
4. Select the **Deployment stage** to which you want to deploy the new API, in this case, the **default** stage.
5. Navigate back to your Lambda function and In the Lambda trigger configuration, verify you see a new API endpoint.
6. Copy the API endpoint from the API Gateway boxes for use in setting up the Snyk webhook.
7. Now that the API endpoint is saved, set up the Snyk Webhook.
