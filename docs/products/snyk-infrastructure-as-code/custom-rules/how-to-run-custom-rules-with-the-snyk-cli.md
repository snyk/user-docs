# Running IaC custom rules

With Snyk IaC, you can test your configuration files from the CLI using your custom rules.&#x20;

There are two ways to use your built bundle:

1. Via a local custom rules bundle stored in your disk.
2. Via configuring a remote OCI Registry, in which your custom rules bundle is already stored.&#x20;

{% hint style="info" %}
Ensure you followed the steps in [Getting Started with the SDK](getting-started-with-the-sdk/) to learn how to generate a bundle of your rules.
{% endhint %}

### Scan for issues using a local custom rules bundle

{% hint style="info" %}
Where the examples show `bundle.tar.gz` you can replace this with your bundle name. For example,`bundle-v1.0.0.tar.gz` or `./bundles/team-bundle.tar.gz`
{% endhint %}

In your project’s folder, run the following command:

```
snyk iac test --rules=bundle.tar.gz
```

The resulting configuration scan issues will include issues from both the default Snyk rules, and your custom rules. Also see [Understanding configuration issues](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/understanding-configuration-scan-issues).&#x20;

#### Troubleshooting

Enable debug logs by running the command with a **-d** flag:

```
snyk iac test --rules=bundle.tar.gz -d
```

Some possible problems:

* Providing an incorrect path to the bundle or a path to a non-existent bundle. Make sure the path passed to the `--rules` flag can be accessed from the current location.

```
We were unable to extract the rules provided at: ./invalid/location/bundle.tar.gz
```

* Providing a corrupted or invalid bundle. Make sure you have generated your bundle by following the instructions at [Getting Started with the SDK](getting-started-with-the-sdk/).

```
We were unable run the test. Please run the command again with the `-d` flag and contact support@snyk.io with the contents of the output.
```

If you have found a discrepancy that you cannot explain, please raise a support ticket.

### Scan for issues using a remote custom rules bundle stored in a Container Registry

{% hint style="info" %}
Ensure you followed the steps in [Getting Started with the SDK](getting-started-with-the-sdk/) to learn how to generate a bundle of your rules and distribute the bundle.
{% endhint %}

In your project’s folder, run the following commands to configure your container registry with the Snyk IaC CLI:

```
export OCI_REGISTRY_URL=<container registry url>
export OCI_REGISTRY_USERNAME=<container registry username>
export OCI_REGISTRY_PASSWORD=<container registry password>
```

{% hint style="info" %}
Ensure the `OCI_REGISTRY_URL` is a valid URL, e.g. for DockerHub:

`https://registry-1.docker.io/repository-name/bundle-image:1.0.0`
{% endhint %}

Then, run a Snyk IaC scan as normal.  The CLI will pull the bundle pushed to the configured container registry in the background.

```
snyk iac test <file>
```

The resulting configuration scan issues will include issues from both the default Snyk rules, and your custom rules. Also see [Understanding configuration issues](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/understanding-configuration-scan-issues).&#x20;

#### Troubleshooting

Enable debug logs by running the command with a **-d** flag:

```
snyk iac test <file> -d
```

Some possible problems:

* Providing an invalid container registry URL. See the note above if you're using Docker Hub.

```
We were unable to download the custom bundle to the disk. 
Please ensure access to the remote Registry and validate you have provided all the right parameters.
```

* Providing invalid credentials.

```
There was an authentication error. Incorrect credentials provided.
    We were unable to download the custom bundle to the disk.
    Please ensure access to the remote Registry and validate you have provided all the right parameters.
```

If you have found a discrepancy that you cannot explain, please raise a support ticket.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up\&loc=footer\&page=support\_docs\_page)
{% endhint %}

