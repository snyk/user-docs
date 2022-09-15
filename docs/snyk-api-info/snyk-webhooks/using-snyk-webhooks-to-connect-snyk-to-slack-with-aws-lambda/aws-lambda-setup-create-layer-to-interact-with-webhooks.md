# AWS Lambda setup: create layer to interact with webhooks

The example code for [Create Lambda function to connect Snyk to Slack](aws-lambda-setup-create-lambda-function-to-connect-snyk-to-slack.md) uses [axios](https://github.com/axios/axios) to interact with the Slack webhooks. Therefore you must add a layer to the new AWS Lambda function to be able to use axios.

To add axios as a layer to interact with webhooks, follow these steps:

1. Navigate to the Lambda dashboard where all Lambda functions are displayed.
2. Click **Layers**.
3.  Select **Create Layer** (far right).\


    <figure><img src="https://lh3.googleusercontent.com/TkKmaVQrbpr2G701sNpYnEyLOvjLQuzY9o47XuvzDER9WkbuTcDuTMWKg5EiCPbUf1swUhRAO0Dc5jZkY9sSsyszDHlYuk_4xS7W2MyDFkYZGmnSMcbWkE7h1R-u5YTGsCHraM9iELvEmKx8osR8Ir90HEv7iwSVMYwd5R_WMzN3BJERhglXZp2F5A" alt="AWS Lambda Layers configuration display"><figcaption><p>AWS Lambda Layers configuration display</p></figcaption></figure>
4. Configure the layer to support **Node.js 16.x** and **x86\_64** architecture in **Compatible architectures** and **Compatible runtimes**.
5. Upload axios as a zip using the Upload button.\
   **Note:** To create a zip version of axios, use the following commands in your terminal:\
   \- `mkdir axios` (to create an axios folder)\
   \- `cd axios` (to navigate inside the folder)\
   \- `npm install` (to set up a package. json file)\
   \- `npm install axios –prefix` (to install axios)\
   \- `rm -rf package-lock.json` (to remove the unnecessary file for the purpose of example code)\
   \- `cd ..` (to navigate back to previous folder)\
   \- `zip -r axios.zip axios` (to zip axios into a file which can be uploaded into the layer)
6.  Verify the resulting layer reflects the preceding entries and click **Create**.\


    <figure><img src="https://lh6.googleusercontent.com/xYSsbBBi1R1oI529hloeTqRYx1t9aWFBhddb9igqONiEgo5mRmRkjWtJaBZ7_qx9IIDZLMc5YpIDVXziNvj37kNGzNG9oX_ZgqggIF68RZsFov9gbuFdhM-KzZBNdflSDc4wIDmFOYPAMtFmtQZo0Ge9BQmbIw0xteUB30Ow1hcTovLpccEzEdihOw" alt="Create AWS Lambda layer interface"><figcaption><p>Create AWS Lambda layer interface</p></figcaption></figure>
7. Once the layer is set up, navigate to the Lambda function to add the layer to the Lambda function.
8. Go to the bottom of the page and select **Add Layer**.
9. You should see the axios layer as a possible layer; click **Add**.
10. Verify the layer has been added to the function.



    <figure><img src="https://lh5.googleusercontent.com/JD2vlLpNDne6JpZ90TE0OM3oCiOn-oqbbaPSqSwgzz2cY8FEbuqRKRyTAeYe2m9wAd8EqKvRADRSfzESliC1YWCwO_O24n56nlM5yPO7NSHOxhNe8owaUHAcUYBsOs1-K91My2wSpw8UeFf5LVA4W_dOwhhGUAbZmTLGsH64ica0az7pZKG4mop3Mw" alt="AWS Lambda function add layer interface"><figcaption><p>AWS Lambda function add layer interface</p></figcaption></figure>
