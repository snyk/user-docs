---
nav_context: agnostic
description: How to connect Snyk to Slack using webhooks and AWS Lambda
---

# How to use Snyk Webhooks to connect Snyk to Slack with AWS Lambda

You can use Snyk Webhooks alongside a Lambda function to receive and filter new vulnerabilities discovered by Snyk in your Slack.

Exposing the Lambda function to Snyk Webhooks is covered through two options, API Gateway and a Lambda Function URL. Choose the variant that suits your needs and environment the best. If you are not using API Gateway, you may disregard any of those instructions in this guide.

The **prerequisites** are as follows:

* An AWS account with access to:
  * Create new roles (or use an existing one)
  * Modify Lambda functions
  * Modify API Gateway (if you are publishing the Lambda function through API Gateway)
* Snyk account with Organization Admin access
* Slack account that can create a new Slack App with a webhook added to an existing channel
* Ability to run npm commands in a terminal

The Snyk Webhook, AWS Lambda function, and Slack notification **work as follows**:

* Each time Snyk discovers a new vulnerability, it triggers the Snyk API Webhook.
* That triggers the Lambda function to send notifications to Slack. The goal of the Lambda function is to filter these notifications and customize the Slack payload to show the information that is interesting to you.
* When you receive the Slack notification you can act on the newly discovered vulnerability.

This guide **explains** how to use an AWS Lambda function to filter the payload from Snyk Webhooks into Slack.

The following describes the **data and traffic flow**:

The Snyk Project Snapshot Webhook triggers an AWS Lambda function by forwarding headers and the POST body through the API Gateway. The Lambda function then sends a filtered payload (custom message) to the Slack webhook, if a hash header signature validation succeeds and if the payload contains valid data. The Lambda function then filters the POST body to construct the custom message.

If you have **problems** using the Snyk Webhook, **contact** your Solutions Engineer or Technical Success Manager for help.
