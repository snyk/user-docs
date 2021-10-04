# Running IaC custom rules

With Snyk IaC, you can test your configuration files from the CLI using your custom rules.

{% hint style="info" %}
Where the examples show `bundle.tar.gz` you can replace this with your bundle name. For example,`bundle-v1.0.0.tar.gz` or `./bundles/team-bundle.tar.gz`
{% endhint %}

Also see [Snyk CLI for Infrastructure as Code](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code)

### To test for a custom issue

{% hint style="info" %}
Ensure you followed the steps in [Getting Started with the SDK](getting-started-with-the-sdk/) to learn how to generate a bundle of your rules.
{% endhint %}

In your project’s folder, run the following command:

```
snyk iac test --rules=bundle.tar.gz
```

The resulting configuration scan issues include issues from both the default Snyk rules, and your  custom rules. Also see [Understanding configuration issues](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/understanding-configuration-scan-issues). 

### Troubleshooting

Enable debug logs by running the command with a **-d** flag:

```text
snyk iac test --rules=bundle.tar.gz -d
```

Some possible problems:

* Providing an incorrect path to the bundle or a path to a non-existent bundle. Make sure the path passed to the `--rules` flag can be accessed from the current location.

```text
We were unable to extract the rules provided at: ./invalid/location/bundle.tar.gz
```

* Providing a corrupted or invalid bundle. Make sure you have generated your bundle by following the instructions at [Getting Started with the SDK](getting-started-with-the-sdk/).

```text
We were unable run the test. Please run the command again with the `-d` flag and contact support@snyk.io with the contents of the output.
```

If you have found a discrepancy that you cannot explain, please raise a support ticket.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}



