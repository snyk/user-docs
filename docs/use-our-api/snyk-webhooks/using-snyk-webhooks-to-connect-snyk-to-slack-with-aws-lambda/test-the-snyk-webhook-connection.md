# Test the Snyk webhook connection

The Snyk Webhook only updates when there is a new vulnerability found, but you can test your implementation using Postman.

If you don’t have Postman you can [install](https://www.postman.com/downloads/) it for free.

To test, you will send a POST request to the AWS API Gateway with a sample payload that is  secured with your API token.

Follow these steps to test the connection:

1. Install Postman.
2.  Create a **Collection**.\


    <figure><img src="https://lh3.googleusercontent.com/j7ab9JGG5IAmqb6xuA7AjwPcF6cmUIhIzrn6p1f7CUQTkwQqHm7P5fVHxDx8I6tysjM93uqu5whBFq_qI1Q5h5y_KK0uR3Hv--uYhcDXJehU5ZCc68Fvdv79S8z7yqCp0CNbLYXOaxwc9cTR0ueQ9lYuydCDhyJmpA5TGNJ08wexCGIpeDMX0fO4Tw" alt="Create new Postman Collection"><figcaption><p>Create new Postman Collection</p></figcaption></figure>
3.  Add your API token (**secret-string**) as an environment variable in the collection .\
    Name the variable `x-hub-signature`so you can use the variable to verify the integrity of the payload.\
    Refer to [Set up the Snyk Webook](set-up-the-snyk-webhook.md) for instructions on finding the API token.\


    <figure><img src="https://lh5.googleusercontent.com/QiPKevkpzyOwSscKxGu9BbzhbfU53bCQKF7y5CaXaImlQFA2VQKuwW5I2TSeKCis1fTDYkyJHaBa8koNDZ1izAHTE1fPWUo2S9bLETght4jPaaKujS8TZKyjOLpk4lUMyeBdSvg5wJvQ553VgK-p_eBJdDyM1St6pXadh9FaVdElZRFh14WBEMLGZA" alt="Add API token (secret-string) as Postman environment variable"><figcaption><p>Create new Postman Collection</p></figcaption></figure>
4.  Within the collection create a new **HTTP Request**.\


    <figure><img src="https://lh3.googleusercontent.com/j7ab9JGG5IAmqb6xuA7AjwPcF6cmUIhIzrn6p1f7CUQTkwQqHm7P5fVHxDx8I6tysjM93uqu5whBFq_qI1Q5h5y_KK0uR3Hv--uYhcDXJehU5ZCc68Fvdv79S8z7yqCp0CNbLYXOaxwc9cTR0ueQ9lYuydCDhyJmpA5TGNJ08wexCGIpeDMX0fO4Tw" alt="Postman create API Request"><figcaption><p>Postman create API Request</p></figcaption></figure>
5.  Change the method to **POST** and add your API Gateway URL. Refer to [AWS API Gateway: add the POST method to connect Snyk to Slack](aws-api-gateway-add-the-post-method-to-connect-snyk-to-slack.md) for instructions on finding the URL.\


    <figure><img src="https://lh4.googleusercontent.com/5QxR-05QtK6FNpoyuPW06L_vyVAl6cCxMnph7euIKafc-YyGIgjaiA74KSNO93uTMGFGxNQnzwyfiZ5Oi3e1y0GA0P2INodvIbamhe6lpwwf1Kc7bCajYUPG0RcfedUOKMqI0l4mmuq1jECRHUiUtnsel7PiBxiIvddcCnplxwVDY9r0FDcYNKZPag" alt="Postman POST method add AWS API Gateway URL"><figcaption><p>Postman POST method add AWS API Gateway URL</p></figcaption></figure>
6.  Configure your pre-request script to look like the following code. This script get your API token and secures it so when Snyk sends it with the payload the Lambda function can decrypt the payload.\
    The code follows; remember to change enter your API token for '`your-secret-string-here'`.\
    \
    `postman.setEnvironmentVariable('x-hub-signature', CryptoJS.HmacSHA256(request.data, 'your-secret-string-here').toString(CryptoJS.digest)); postman.setEnvironmentVariable('x-hub-signature', 'sha256='+ postman.getEnvironmentVariable('x-hub-signature'));`\
    ``&#x20;

    <figure><img src="https://lh4.googleusercontent.com/imlrHdNQOJQVExPXvHiwNSR0zerKrR4qUJKeeXmJsfW-UTarEZtB9S3uW5K0xY4EarI5zft8PqUKEE5AS3TPWIWE5hTMNrLA5iCmv8f9Nv5onoTzPRsS8lXUTOQt4Fl-SFyFMvyTfLs3FBhcu_PCwjfB0zLvFXqGPFjYPw3b6ctorVVZ3YsVMQeVpg" alt="Postman pre-request script"><figcaption><p>Postman pre-request script</p></figcaption></figure>
7.  Add the **Headers**:\
    **Content-Type**: **application/json**\
    **x-hub-signature** `{{x-hub-signature}}`:  (The API token (secret-string) environment variable you created in step 3\


    <figure><img src="https://lh5.googleusercontent.com/SLs1bStNsB5yEBMRpie_PseTXwZuj5qYp_w5CIboLgNcrAJks87wVzoJuwI0TVa71kbXSS-k0zHbrEVSXaKp3j33S3Jn3Fy5dH21Yla8iNqFFSFqHQDf6ArhjbxUheFAaZbPFYoLuyhxoHlsKDNJkdoSk2L7v0vDGUrN4_-Bcf7S91PgvvT6wtZt9w" alt="Postman POST request header"><figcaption><p>Postman POST request headers</p></figcaption></figure>
8.  Add to the **Body** a valid payload, for example the one here and do not beautify it, if it looks readable it will not work:\
    `{"project":{"id":"bc75a806-0893-4ccf-84c5-8fde48a88df8","name":"snyk/juice-shop:frontend/package.json","created":"2022-06-17T06:54:21.326Z","origin":"github","type":"npm","readOnly":false,"testFrequency":"daily","totalDependencies":1216,"issueCountsBySeverity":{"low":2,"high":16,"medium":17,"critical":0},"imageTag":"12.3.0","imagePlatform":"","lastTestedDate":"2022-06-29T05:45:12.569Z","browseUrl":"https://app.snyk.io/org/api-test/project/bc75a806-0893-4ccf-84c5-8fde48a88df7","importingUser":null,"owner":null,"tags":[],"isMonitored":true,"attributes":{"criticality":[],"lifecycle":[],"environment":[]},"branch":"master"},"org":{"id":"570a1e02-8774-4697-80fc-129f5c5195a1","name":"API","slug":"api-quc","url":"https://app.snyk.io/org/api-test","group":null,"created":"2022-05-25T06:29:29.833Z"},"newIssues":[{"id":"SNYK-JS-SCSSTOKENIZER-2339884","issueType":"vuln","pkgName":"scss-tokenizer","pkgVersions":["0.2.3"],"priorityScore":336,"priority":{"score":336,"factors":[{"name":"isFresh","description":"Recently disclosed"},{"name":"cvssScore","description":"CVSS 5.3"}]},"issueData":{"id":"SNYK-JS-SCSSTOKENIZER-2339884","title":"Regular Expression Denial of Service (ReDoS)","severity":"high","url":"https://snyk.io/vuln/SNYK-JS-SCSSTOKENIZER-2339884","description":"Long description","identifiers":{"CWE":["CWE-1333"],"CVE":["CVE-2022-25758"]},"credit":["Paul Bastide"],"exploitMaturity":"no-known-exploit","semver":{"vulnerable":["*"]},"publicationTime":"2022-06-29T10:29:38Z","disclosureTime":"2022-01-13T16:29:34Z","CVSSv3":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L","cvssScore":5.3,"functions":[],"language":"js","patches":[],"nearestFixedInVersion":"","isMaliciousPackage":false},"isPatched":false,"isIgnored":false,"fixInfo":{"isUpgradable":false,"isPinnable":false,"isPatchable":false,"isFixable":false,"isPartiallyFixable":false,"nearestFixedInVersion":"","fixedIn":[]}}],"removedIssues":[]}`\


    <figure><img src="https://lh6.googleusercontent.com/vi_Mt44ag0EzWi9bn9ruwnzBcF-cYxGqajF-F6jQF2nwJEEvNa6KW45ZgszlekP17zLQwRH-z9iar-oTvkOKXdAWEb-ewCJVujrj-pzkHlKftd4Y1GmPyaguELBtbKP-m3RLAN9-R6PxzO1psWDY_KoW7iHwLc3oQax7gcQArwMtf2oxSlmvHUxzWA" alt="Postman POST request body with valid payload"><figcaption><p>Postman POST request body with valid payload</p></figcaption></figure>
9. Click the **Send** button next to the URL
10. Verify that a Success message (`statusCode 200`)appears at the bottom of Postman.\


    <figure><img src="https://lh4.googleusercontent.com/YHelnzIIPgeL7ZkbVOy67hMiaVe6_lz3VvFjhNg8vkeRm4EtevSypMR_PsSRCfzkZcob76KSSgdvrPoqhVwEBL8FwT2LXiIn9u9hv5-bVrF_zh7sK3lB0rJM3lBmqc5w6miUx7hD7ROlLrXROIbAgUWWqCnYpvZ6C8TJcKI_kSTYG5LMaYg2lRm3RA" alt="Postman success message"><figcaption><p>Postman success message</p></figcaption></figure>
11. Verify that you see a new notification in Slack: **New Snyk Vulnerability** with the path, **Package** **name**, **scss-tokenizer,** **Severity level**, **Vulnerability name**, and **Priority Score**.\


    <figure><img src="https://lh5.googleusercontent.com/1nvqWOgUaA6P6kc7MTObqXxfEXrFaP1DKXqHKy7wQhPxpWIA9HyMHV7dwOHd2HGQiJuL9rwn9aVQlhvlg-rBcHTggXh6nhRWB8T7PAtfM4S73bTL1ytUK3ZaKtzbCnofDUg9ER22zcMI84PXv1byQnN9BUToJk49qiOcq6627VLFlDvUBrXpL1Atjg" alt=""><figcaption><p>New Slack notification of vulnerability<br><br></p></figcaption></figure>

    The next time Snyk finds a new vulnerability, you will see a notification from Snyk in Slack.
