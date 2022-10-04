# Using Snyk Webhooks to connect Snyk to Slack with AWS Lambda

You can use Snyk Webhooks alongside a Lambda function to receive and filter new vulnerabilities discovered by Snyk in your Slack.

The **prerequisites** are as follows:

* An AWS account with access to:
  * Create new roles (or use an existing one)&#x20;
  * Modify Lambda functions
  * Modify API Gateway
* Snyk account with Organization Admin access
* Slack account that can create a new Slack App with a webhook added to an existing channel

The Snyk Webhook, AWS Lambda function, and Slack notification **work as follows**:

* Each time Snyk discovers a new vulnerability, it triggers the Snyk API Webhook.
* That triggers the Lambda function to send notifications to Slack. The goal of the Lambda function is to filter these notifications and customize the Slack payload to show the information that is interesting to you.
* When you receive the Slack notification you can act on the newly discovered vulnerability.

This guide **explains** how to use an AWS Lambda function to filter the payload from Snyk Webhooks into Slack.

The following describes the **data and traffic flow**:

The Snyk Project Snapshot Webhook triggers an AWS Lambda function by forwarding headers and the POST body through the API Gateway. The Lambda function then sends a filtered payload (custom message) to the Slack webhook, if a hash header signature validation succeeds and if the payload contains valid data. The Lambda function then filters the POST body to construct the custom message.

<figure><img src="https://lh6.googleusercontent.com/VROtTsX240dfLMERpOkm-5epOnvZxQUxjM-qKJYNEOtD_1flwBrpBTiJedo2Uy0RZz6kKplKNQQcINzOW3H30Lf7R9U0teZ4WvivBt1u7TdN_4J3ha_ZmY9wdn3xvXCNxl9036JdYeEzaBMtU53lo6e-do3Bhbmi4Y9tcWDO5y00NT_XRvmt5Z9ipg" alt="Data and traffic flow for using Snyk Webhooks to connect Snyk to Slack with AWS Lambda"><figcaption><p>Data and traffic flow for using Snyk Webhooks to connect Snyk to Slack with AWS Lambda</p></figcaption></figure>

If you have **problems** using the Snyk Webhook, **contact** your Solution Engineer or Technical Success Manager for help.
