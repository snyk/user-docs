# Validate Payload

Next, we will add another **Action** and name it **Validate payload**:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-validate-payload-main.png)

Let's define our **Action Event** as **Run Javascript**:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-validate-payload-script.png)

Configure your action to include **body** and **signature** as shown below:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-validate-payload-setup.png)

In the **Code** section, copy and paste the following snippet:

```javascript
const crypto = require('crypto');
const secret = "your_secret";

function makeSignature(body, secret) {
    const hmac = crypto.createHmac('sha256', secret);
    hmac.update(body, 'utf8');
    return `sha256=${hmac.digest('hex')}`;
}

try {
    const body = JSON.parse(inputData.body);
    const { project, org, group, newIssues } = body;
    output = { 
        isValid: inputData.signature === makeSignature(inputData.body, secret)
    };
} catch (err) {
    output = { isValid: false, err: err.message };
}
```

Test your action:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-validate-payload-test.png)

