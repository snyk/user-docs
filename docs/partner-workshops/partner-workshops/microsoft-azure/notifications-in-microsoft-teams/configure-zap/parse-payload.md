# Parse Payload

Next, we will add another action and name it **Parse payload**:

![](../../../../.gitbook/assets/zappier-parse-payload-main.png)

Select **Run Javascript** as the **Action Event**:

![](../../../../.gitbook/assets/zappier-parse-payload-script.png)

Be sure to add **body** under **Input Data** as shown below:

![](../../../../.gitbook/assets/zappier-parse-payload-setup.png)

Copy and paste the following snippet in the **Code** field:

```javascript
try {
    output = JSON.parse(inputData.body);
} catch (err) {
    output = { err: err.message };
}
```

Test the action:

![](../../../../.gitbook/assets/zappier-parse-payload-test.png)

