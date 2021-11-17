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

Additionally, once you have generated your custom rules bundle you can distribute it to one of our supported OCI registries by following the steps in [Pushing a bundle](getting-started-with-the-sdk/pushing-a-bundle.md).&#x20;

After successfully pushing your custom rules bundle, you can configure it to be applied in either:

1. The settings page.
2. Environment variables.

#### Configure remote custom rules bundles using the settings page

We recommend you use the Snyk settings page to configure custom rules settings. This method gives a simple way to update the custom rules bundle's URL and tag whenever these are modified.

{% hint style="info" %}
Tags are helpful for versioning your custom rules bundles. Configuring your updated bundle can be easily accomplished by setting the new version tag.
{% endhint %}

These remote bundles can be configured on both the organizational level and the group level. Configuring a remote bundle for a group applies the remote bundle to all the organizations in the group.

{% hint style="info" %}
Configuring remote custom rules bundles on the organizational level can be done by navigating to `Settings` > `Infrastructure as Code`.

Similarly, configuring them on the group level can be done by navigating to` Settings` > `Infrastructure as Code.`
{% endhint %}

To configure remote bundles:

* In the Infrastructure as Code Settings, locate the **Rules** section.

![](<../../../.gitbook/assets/image (78) (1) (1).png>)

* Enable the usage of remote bundles configuration using the **Enable rules **toggle. Doing so will display the form as shown below:

![](<../../../.gitbook/assets/image (83) (1).png>)

* Configure the OCI registry URL and tag for your remote bundle of custom rules, and click  Save changes to save.

![](<../../../.gitbook/assets/image (74) (1) (1) (1).png>)

Your remote bundle of custom rules has now been configured and will be used when testing IaC files.

#### Overriding a group's remote bundle configurations

By default, configuring a remote bundle for a group applies the remote bundle to all the organizations in the group. As such, whenever the group configurations are updated, these changes apply to all of its organizations.

However, an organization can still override the group's configurations and define its own bundle and tag. These will not change when the group updates its configurations.

In order to override the group's configurations, head to the organization's `Rules` section on the Infrastructure as Code Settings.

![](<../../../.gitbook/assets/image (66) (1) (1).png>)

Initially, you will see the section is populated with the configurations inherited from the organization's group.

Update the configurations to those customized for your organization, and click to `Save changes`.

![](<../../../.gitbook/assets/image (68) (1) (1) (1).png>)

* Now, configurations on the group level will not override these customized settings for your organization.

You can restore the inheritance of group configurations at any time by using the `Reset to group default` button.

#### Configure remote custom rules bundles using environment variables

Optionally, if you don't want to use the Snyk UI, you can configure the location of the custom rules bundle using Snyk config. In your project’s folder, run the following commands to configure your container registry with the Snyk IaC CLI:

```
snyk config set oci-registry-url=<container registry url>
snyk config set oci-registry-username=<container registry username>
snyk config set oci-registry-password=<container registry password>
```

This will set the following Snyk environment variables:

* `SNYK_CFG_OCI_REGISTRY_URL`
* `SNYK_CFG_OCI_REGISTRY_USERNAME`
* `SNYK_CFG_OCI_REGISTRY_PASSWORD`

{% hint style="info" %}
Ensure the OCI Registry URL is a valid URL, e.g. for DockerHub:

`https://registry-1.docker.io/repository-name/bundle-image:1.0.0`



Make sure to clear any previously defined URLs in the settings page, as only one method for defining the bundle's path should be defined at any given time.
{% endhint %}

Then, run a Snyk IaC scan as normal.  The CLI will pull the bundle pushed to the configured container registry in the background.

```
snyk iac test <file>
```

The resulting configuration scan issues will include issues from both the default Snyk rules, and your custom rules. Also see [Understanding configuration issues](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/understanding-configuration-scan-issues).&#x20;

{% hint style="warning" %}
Make sure to either disable the custom rules settings in the Snyk UI or clear any previously-stored settings using `snyk config unset`. Only one method for defining the bundle's path should be defined at any given time.
{% endhint %}

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

