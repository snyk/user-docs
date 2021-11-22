# Using a remote custom rules bundle

After you generate your custom rules bundle you can distribute it to one of our supported OCI registries by following the steps in [Pushing a bundle](../getting-started-with-the-sdk/pushing-a-bundle.md).&#x20;

After successfully pushing your custom rules bundle, you can enforce its usage using:

* [Snyk Settings page](using-a-remote-custom-rules-bundle.md#snyk-settings-page)
* [Snyk API](using-a-remote-custom-rules-bundle.md#snyk-api)
* [Environment variables](using-a-remote-custom-rules-bundle.md#environment-variables)

Finally, once you've enforced your custom rules via one of the options above, configure the Snyk Snyk CLI with your username and password in order to allow us to authorize a pull from your OCI registry:

```
snyk config set oci-registry-username=<org registry username>
snyk config set oci-registry-password=<org registry password>
```

This will set the following Snyk environment variables:

* `SNYK_CFG_OCI_REGISTRY_USERNAME`
* `SNYK_CFG_OCI_REGISTRY_PASSWORD`&#x20;

Once you have done that, run a Snyk IaC scan as normal. The CLI will pull the bundle pushed to the configured container registry in the background.

```
snyk iac test <file>
```

The resulting configuration scan issues will include issues from both the default Snyk rules, and your custom rules. Also see [Understanding configuration issues](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/understanding-configuration-scan-issues).&#x20;

{% hint style="warning" %}
Only one method for defining the bundle's path should be defined at any given time. Make sure to either disable the custom rules settings via the Snyk Settings page or the Snyk API. Alternatively, clear any previously-stored settings using `snyk config unset`.&#x20;
{% endhint %}

### Snyk Settings page

We recommend you use the Snyk Settings page to configure custom rules settings. This method gives a simple way to update the custom rules bundle's URL and tag whenever these are modified.

{% hint style="info" %}
Tags are helpful for versioning your custom rules bundles. Configuring your updated bundle can be easily accomplished by setting the new version tag.
{% endhint %}

You can configure these remote bundles on both the organizational level and the group level. Configuring a remote bundle for a group applies the remote bundle to all the organizations in the group.

To configure remote bundles:

* In the Infrastructure as Code Settings, locate the **Rules** section.

{% hint style="info" %}
Configuring remote custom rules bundles on the organizational level can be done by navigating to `Settings` > `Infrastructure as Code`.

Similarly, configuring them on the group level can be done by navigating to` Settings` > `Infrastructure as Code.`
{% endhint %}

![](<../../../../.gitbook/assets/image (78) (1) (1).png>)

* Enable the usage of remote bundles configuration using the **Enable rules **toggle. Doing so will display the form as shown below:

![](<../../../../.gitbook/assets/image (83) (1).png>)

* Configure the OCI registry URL and tag for your remote bundle of custom rules, and click **Save changes** to save.

![](<../../../../.gitbook/assets/image (74) (1) (1).png>)

Your remote bundle of custom rules is now configured and will be used when testing IaC files.

#### Overriding a group's remote bundle configurations

By default, configuring a remote bundle for a group applies the remote bundle to all the organizations in the group. So if the group configurations are updated, these changes apply to all of its organizations.

However, an organization can still override the group's configurations and define its own bundle and tag. These will not change when the group updates its configurations.

In order to override the group's configurations, go to the organization's `Rules` section on the Infrastructure as Code Settings.

* Initially, the section is populated with the configurations inherited from the organization's group.

![](<../../../../.gitbook/assets/image (66) (1) (1).png>)

* Update the configurations to those customized for your organization, and click **Save changes**.

![](<../../../../.gitbook/assets/image (68) (1) (1).png>)

* Now, configurations on the group level will not override these customized settings for your organization.

You can restore the inheritance of group configurations at any time by using the `Reset to group default` button.

### Snyk API

If manually updating the settings through the Snyk Settings page is too time-consuming, another option is to use the Snyk API. The API currently allows to send any variation of the custom rules settings via a `PATCH` API call to the group-level settings.

{% hint style="warning" %}
Currently there is no API to override the custom rules settings at the organization level.
{% endhint %}

* For example, in order to configure the custom rules bundle at the group level, call the Group IaC Settings API as explained in the [documentation](https://snykv3.docs.apiary.io/#reference/group-settings/infrastructure-as-code/update-infrastructure-as-code-settings) by providing the following body:

```
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "oci_registry_url": "https://group-registry/library/image",
             "oci_registry_tag": "latest",
             "is_enabled": true
           }
       }
   }
}
```

* If you want to update the tag only, you can send over a simpler body:

```
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "oci_registry_tag": "v1"
           }
       }
   }
}
```

* And if you want to disable custom rules, you can just send over the `is_enabled` flag:

```
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "is_enabled": false
           }
       }
   }
}
```



### Environment variables

You can also configure the location of the custom rules bundle using Snyk config for your organization. In your project’s folder, run the following commands to configure your container registry with the Snyk IaC CLI:

```
snyk config set oci-registry-url=https://org-registry/library/image
```

This will set the following Snyk environment variable: `SNYK_CFG_OCI_REGISTRY_URL`

{% hint style="info" %}
Ensure the OCI Registry URL is a valid URL; for example, for DockerHub:

`https://registry-1.docker.io/repository-name/bundle-image:1.0.0`



Make sure to clear any previously defined URLs in the Snyk Settings page or disable custom rules, as only one method for defining the bundle's path should be defined at any given time.
{% endhint %}



### Troubleshooting

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
