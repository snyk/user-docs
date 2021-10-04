# Format Issues

Next, create another action and name it **Format issues**:

![](https://github.com/snyk/user-docs/tree/0874305e3aea1ea3c57b0398879776ac062b3479/.gitbook/assets/zappier-format-issues-main.png)

Select **Run Javascript** as the **Action Event**:

![](https://github.com/snyk/user-docs/tree/0874305e3aea1ea3c57b0398879776ac062b3479/.gitbook/assets/zappier-format-issues-script.png)

Be sure to add **body** in the **Input Data** section:

![](https://github.com/snyk/user-docs/tree/0874305e3aea1ea3c57b0398879776ac062b3479/.gitbook/assets/zappier-format-issues-setup.png)

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

![](https://github.com/snyk/user-docs/tree/0874305e3aea1ea3c57b0398879776ac062b3479/.gitbook/assets/zappier-format-issues-test.png)

