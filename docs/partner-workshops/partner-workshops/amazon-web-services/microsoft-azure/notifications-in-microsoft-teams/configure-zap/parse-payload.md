# Parse Payload

Next, we will add another action and name it **Parse payload**:

![](https://github.com/snyk/user-docs/tree/0874305e3aea1ea3c57b0398879776ac062b3479/.gitbook/assets/zappier-parse-payload-main.png)

Select **Run Javascript** as the **Action Event**:

![](https://github.com/snyk/user-docs/tree/0874305e3aea1ea3c57b0398879776ac062b3479/.gitbook/assets/zappier-parse-payload-script.png)

Be sure to add **body** under **Input Data** as shown below:

![](https://github.com/snyk/user-docs/tree/0874305e3aea1ea3c57b0398879776ac062b3479/.gitbook/assets/zappier-parse-payload-setup.png)

Copy and paste the following snippet in the **Code** field:

```javascript
try {
    output = JSON.parse(inputData.body);
} catch (err) {
    output = { err: err.message };
}
```

Test the action:

![](https://github.com/snyk/user-docs/tree/0874305e3aea1ea3c57b0398879776ac062b3479/.gitbook/assets/zappier-parse-payload-test.png)

