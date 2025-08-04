# Use a local IaC custom rules bundle

{% hint style="info" %}
Where the examples show `bundle.tar.gz`, you can replace this with your bundle name. For example,`bundle-v1.0.0.tar.gz` or `./bundles/team-bundle.tar.gz`
{% endhint %}

In your Project folder, run the following command:

```
snyk iac test --rules=bundle.tar.gz
```

The configuration scan results now include issues from both the default Snyk rules and your custom rules. See [Understanding the IaC CLI test results](../../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/understand-the-iac-cli-test-results/).

To troubleshoot a local custom rules bundle, enable debug logs by running the command with the `--d` options:

```
snyk iac test --rules=bundle.tar.gz -d
```

Some possible problems include:

* Providing an incorrect path to the bundle or a path to a non-existent bundle. Ensure the path passed to the `--rules` option can be accessed from the current location. The error is

```
We were unable to extract the rules provided at: ./invalid/location/bundle.tar.gz
```

* Providing a corrupted or invalid bundle. Ensure you have generated your bundle by following the instructions in [Getting Started with the SDK](../writing-rules-using-the-sdk/). The error is

```
We were unable run the test. Please run the command again with the `-d` flag and contact support@snyk.io with the contents of the output.
```

If you have found a discrepancy that you cannot explain, [contact Snyk Support](https://support.snyk.io).
