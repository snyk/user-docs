---
nav_context: agnostic
description: >-
  How to test the Snyk webhook connection for the Slack and AWS Lambda
  integration
---

# Test the Snyk webhook connection

The Snyk Webhook only updates when there is a new vulnerability found, but you can test your implementation using Postman.

If you don’t have Postman you can [install](https://www.postman.com/downloads/) it for free.

To test, you will send a POST request to the AWS API Gateway with a sample payload that is secured with your API token.

Follow these steps to test the connection:

1. Install Postman.
2. Create a **Collection**.
3. Add your API token (**secret-string**) as an environment variable in the collection .\
    Name the variable `x-hub-signature`so you can use the variable to verify the integrity of the payload.\
    Refer to [Set up the Snyk Webook](set-up-the-snyk-webhook.md) for instructions on finding the API token.
4. Within the collection create a new **HTTP Request**.
5. Change the method to **POST** and add your API Gateway URL or Function URL. Refer to [AWS API Gateway: add the POST method to connect Snyk to Slack](aws-lambda-setup-set-up-the-trigger/with-api-gateway/aws-api-gateway-add-the-post-method-to-connect-snyk-to-slack.md) for instructions on finding the URL.
6. Configure your pre-request script to look like the following code. This script get your API token and secures it so when Snyk sends it with the payload the Lambda function can decrypt the payload.\
    The code follows; remember to change enter your API token for '`your-secret-string-here'`.\
    \
    `postman.setEnvironmentVariable('x-hub-signature', CryptoJS.HmacSHA256(request.data, 'your-secret-string-here').toString(CryptoJS.digest)); postman.setEnvironmentVariable('x-hub-signature', 'sha256='+ postman.getEnvironmentVariable('x-hub-signature'));`
7. Add the **Headers**:\
    **Content-Type**: **application/json**\
    **x-hub-signature** `{{x-hub-signature}}`: (The API token (secret-string) environment variable you created in step 3.
8. Add to the **Body** a valid payload, for example the one here and do not beautify it, if it looks readable it will not work:\
    `{"project":{"id":"bc75a806-0893-4ccf-84c5-8fde48a88df8","name":"snyk/juice-shop:frontend/package.json","created":"2022-06-17T06:54:21.326Z","origin":"github","type":"npm","readOnly":false,"testFrequency":"daily","totalDependencies":1216,"issueCountsBySeverity":{"low":2,"high":16,"medium":17,"critical":0},"imageTag":"12.3.0","imagePlatform":"","lastTestedDate":"2022-06-29T05:45:12.569Z","browseUrl":"https://app.snyk.io/org/api-test/project/bc75a806-0893-4ccf-84c5-8fde48a88df7","importingUser":null,"owner":null,"tags":[],"isMonitored":true,"attributes":{"criticality":[],"lifecycle":[],"environment":[]},"branch":"master"},"org":{"id":"570a1e02-8774-4697-80fc-129f5c5195a1","name":"API","slug":"api-quc","url":"https://app.snyk.io/org/api-test","group":null,"created":"2022-05-25T06:29:29.833Z"},"newIssues":[{"id":"SNYK-JS-SCSSTOKENIZER-2339884","issueType":"vuln","pkgName":"scss-tokenizer","pkgVersions":["0.2.3"],"priorityScore":336,"priority":{"score":336,"factors":[{"name":"isFresh","description":"Recently disclosed"},{"name":"cvssScore","description":"CVSS 5.3"}]},"issueData":{"id":"SNYK-JS-SCSSTOKENIZER-2339884","title":"Regular Expression Denial of Service (ReDoS)","severity":"high","url":"https://snyk.io/vuln/SNYK-JS-SCSSTOKENIZER-2339884","description":"Long description","identifiers":{"CWE":["CWE-1333"],"CVE":["CVE-2022-25758"]},"credit":["Paul Bastide"],"exploitMaturity":"no-known-exploit","semver":{"vulnerable":["*"]},"publicationTime":"2022-06-29T10:29:38Z","disclosureTime":"2022-01-13T16:29:34Z","CVSSv3":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L","cvssScore":5.3,"functions":[],"language":"js","patches":[],"nearestFixedInVersion":"","isMaliciousPackage":false},"isPatched":false,"isIgnored":false,"fixInfo":{"isUpgradable":false,"isPinnable":false,"isPatchable":false,"isFixable":false,"isPartiallyFixable":false,"nearestFixedInVersion":"","fixedIn":[]}}],"removedIssues":[]}`
9. Click the **Send** button next to the URL
10. Verify that a Success message (`statusCode 200`)appears at the bottom of Postman.
11. Verify that you see a new notification in Slack: **New Snyk Vulnerability** with the path, **Package** **name**, **scss-tokenizer,** **Severity level**, **Vulnerability name**, and **Priority Score**.

    The next time Snyk finds a new vulnerability, you will see a notification from Snyk in Slack.
