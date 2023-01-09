# Format Issues

Next, create another action and name it **Format issues**:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-format-issues-main.png)

Select **Run Javascript** as the **Action Event**:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-format-issues-script.png)

Be sure to add **body** in the **Input Data** section:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-format-issues-setup.png)

Copy and paste the snippet into the **Code** field:

```javascript
function formatIssue({ pkgName, pkgVersions, issueData }) {
    return `
    <a href="${issueData.url}">${issueData.title}</a><br/>
    Vulnerability in ${pkgName} (${pkgVersions.join(', ')}). ${issueData.severity} severity.
    `;
}

try {
    const { newIssues } = JSON.parse(inputData.body);
    output = { newIssues: newIssues.map(formatIssue) };
} catch (err) {
    output = { newIssues: [], err: err.message };
}
```

Test the action:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/zappier-format-issues-test.png)

