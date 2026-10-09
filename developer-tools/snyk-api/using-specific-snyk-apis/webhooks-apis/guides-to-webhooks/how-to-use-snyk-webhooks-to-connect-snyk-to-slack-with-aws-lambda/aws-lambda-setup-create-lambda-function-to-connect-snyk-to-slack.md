---
nav_context: agnostic
description: How to create an AWS Lambda function to connect Snyk to Slack
---

# AWS Lambda setup: create Lambda function to connect Snyk to Slack

AWS Lambda functions are used to connect Snyk to Slack because these functions are an inexpensive and efficient way of running code triggered by events, for example when there is a new Snyk vulnerability.

**Note:** If publishing the Lambda function through API Gateway, both must be configured in the same region. You can check this on the top right of the AWS Console.

Start by creating a zip file containing the code for the function and the necessary dependencies.

1. Save these two code blocks as package.json and index.js
   1.  package.json (modify to fit any other dependencies needed for your code, the dependencies needed are **axios** and **crypto**)

       ```json
       {
         "name": "snyk-webhook-handler",
         "version": "1.0.0",
         "description": "Snyk to Slack Webhook Integration",
         "main": "index.js",
         "scripts": {
           "test": "echo \"Error: no test specified\" && exit 1"
         },
         "author": "",
         "license": "ISC",
         "dependencies": {
           "axios": "^1.1.3",
           "crypto": "^1.0.1"
         }
       }
       ```
   2.  index.js

       ```javascript
       const crypto = require('crypto')
       const axios = require('axios')

       let slackWebhookUrl = '<your_slackWebhookUrl_here>' // adjust

       //customised messaging to Slack with issue information, modify as needed
       async function messageSlack(
         message,
         snykProjectUrl,
         snykProjectName,
         snykIssuePackage,
         snykIssueUrl,
         snykIssueId,
         severity,
         snykIssuePriority
       ) {
         //strings modified to avoid Axios/Slack errors
         snykProjectUrl = snykProjectUrl.replace(/['"]+/g, '')
         snykProjectName = snykProjectName.replace(/['"]+/g, '')
         snykIssueUrl = snykIssueUrl.replace(/['"]+/g, '')
         snykIssueId = snykIssueId.replace(/['"]+/g, '')
         snykIssuePackage = snykIssuePackage.replace(/['"]+/g, '')
         severity = severity.replace(/['"]+/g, '')

         //construct message
         let payload = {
           text: `${message}`,
           blocks: [
             {
               type: 'header',
               text: {
                 type: 'plain_text',
                 text: `${message}`,
               },
             },
             {
               type: 'section',
               text: {
                 type: 'mrkdwn',
                 text: `Snyk has found a new vulnerability in the project:\n*<${snykProjectUrl}|${snykProjectName}>*`,
               },
             },
             {
               type: 'divider',
             },
             {
               type: 'section',
               fields: [
                 {
                   type: 'mrkdwn',
                   text: `*Package name:*\n${snykIssuePackage}`,
                 },
                 {
                   type: 'mrkdwn',
                   text: `*Vulnerability:*\n<${snykIssueUrl}|${snykIssueId}>`,
                 },
                 {
                   type: 'mrkdwn',
                   text: `*Severity:*\n${severity}`,
                 },
                 {
                   type: 'mrkdwn',
                   text: `*Priority Score:*\n${snykIssuePriority}`,
                 },
               ],
             },
             {
               type: 'actions',
               elements: [
                 {
                   type: 'button',
                   text: {
                     type: 'plain_text',
                     text: 'View in Snyk',
                   },
                   style: 'primary',
                   url: snykProjectUrl,
                   value: 'browseUrl',
                 },
               ],
             },
           ],
         }

         //send message
         const res = await axios.post(slackWebhookUrl, payload)
         const data = res.data
       }

       exports.handler = async (event) => {
         // Securing integrity of payload, this can be moved to another Lambda function and called seperately for authentication
         let response

         const {hmac_verification, severity_threshold} = process.env
         const hmac = crypto.createHmac('sha256', hmac_verification)
         const buffer = JSON.stringify(event.body)
         hmac.update(buffer, 'utf8')
         const stored_signature = `sha256=${hmac.digest('hex')}`

         let sent_signature = event.headers['x-hub-signature']

         if (stored_signature !== sent_signature) {
           console.log('Integrity of request compromised, aborting')
           response = {
             statusCode: 403,
             body: JSON.stringify('Bad request'),
           }
           return response
         }

         // If integrity is ok, verify that the webhook actually contains the project object, iterate and filter
         if (buffer.indexOf('project') !== -1 && buffer.indexOf('newIssues') !== -1) {
           // Iterate through new issues
           var len = buffer['newIssues'] ? buffer['newIssues'].length : 0

           for (let x = 0; x < len; x++) {
             // Get Severity
             let severity = JSON.stringify(buffer['newIssues'][x]['issueData']['severity'])
             // Filter
             if (severity.includes('high') || severity.includes('critical')) {
               let snykProjectName = JSON.stringify(buffer['project'].name)
               let snykProjectUrl = JSON.stringify(buffer['project'].browseUrl)
               let snykIssueUrl = JSON.stringify(buffer['newIssues'][x]['issueData'].url)
               let snykIssueId = JSON.stringify(buffer['newIssues'][x].id)
               let snykIssuePackage = JSON.stringify(buffer['newIssues'][x].pkgName)
               let snykIssuePriority = JSON.stringify(buffer['newIssues'][x]['priority'].score)
               let message = 'New Snyk Vulnerability'

               // Send the result to Slack
               await messageSlack(
                 message,
                 snykProjectUrl,
                 snykProjectName,
                 snykIssuePackage,
                 snykIssueUrl,
                 snykIssueId,
                 severity,
                 snykIssuePriority
               )
             }
           }
         }
         //do nothing, or modify for any preferable action
         else {
           console.log('Valid webhook, but project missing or empty')
         }

         //respond to Snyk
         response = {
           statusCode: 200,
           body: JSON.stringify('Success'),
         }

         return response
       }

       ```
2. Use the following commands in your terminal:\
   \- `cd /path/to/snyk/folder` (to navigate inside the folder where you stored the two files)\
   \- `npm install` (to set up a package. json file)\
   \- `cd ..` (to navigate back to previous folder)\
   \- `zip -r snyk.zip /path/to/snyk/folder` (to zip snyk into a file which can be uploaded to AWS Lambda)

To create an AWS Lambda function, follow these steps:

1. Go to the AWS Console.
2. Navigate to Lambda.
3. Click **Create function**.
4. Choose **Node.js 16.x** for the **Runtime**.
5. Choose **X86\_64** for the **Architecture**.
6. If publishing the Lambda function through API Gateway, add or create a role with the default policy **AmazonAPIGatewayInvokeFullAccess** to interact with the AWS API Gateway.
7. Click **Create Function** and when the function is ready click **Upload from .zip file** and
8. Verify that the code you entered is displayed in the Code Source.
9. In the code, modify the `slackWebhookUrl` to match your Slack webhook URL.
10. For more information on the script you have pasted, go to [Configure the AWS Lambda Script](configure-the-aws-lambda-script.md).
