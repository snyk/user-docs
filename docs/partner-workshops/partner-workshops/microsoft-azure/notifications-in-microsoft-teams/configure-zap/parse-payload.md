# Parse Payload

Next, we will add another action and name it **Parse payload**:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-parse-payload-main.png)

Select **Run Javascript** as the **Action Event**:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-parse-payload-script.png)

Be sure to add **body** under **Input Data** as shown below:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-parse-payload-setup.png)

Copy and paste the following snippet in the **Code** field:

```javascript
try {
    output = JSON.parse(inputData.body);
} catch (err) {
    output = { err: err.message };
}
```

Test the action:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-parse-payload-test.png)

