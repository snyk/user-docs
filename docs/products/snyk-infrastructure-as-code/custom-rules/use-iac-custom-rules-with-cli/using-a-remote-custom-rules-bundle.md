# Using a remote custom rules bundle

After you generate your custom rules bundle you can distribute it to one of our supported OCI registries by following the steps in [Pushing a bundle](../getting-started-with-the-sdk/pushing-a-bundle.md).

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
* `SNYK_CFG_OCI_REGISTRY_PASSWORD`

Once you have done that, run a Snyk IaC scan as normal. The CLI will pull the bundle pushed to the configured container registry in the background.

```
snyk iac test <file>
```

The resulting configuration scan issues will include issues from both the default Snyk rules, and your custom rules. Also see [Understanding configuration issues](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/understanding-configuration-scan-issues).

{% hint style="warning" %}
Only one method for defining the bundle's path should be defined at any given time. Make sure to either disable the custom rules settings via the Snyk Settings page or the Snyk API. Alternatively, clear any previously-stored settings using `snyk config unset`.
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

Similarly, configuring them on the group level can be done by navigating to `Settings` > `Infrastructure as Code.`
{% endhint %}

![](<../../../../.gitbook/assets/image (215) (1) (1) (1) (1) (1).png>)

* Enable the usage of remote bundles configuration using the **Enable rules** toggle. Doing so will display the form as shown below:

![](<../../../../.gitbook/assets/image (11) (1) (1).png>)

* Configure the OCI registry URL and tag for your remote bundle of custom rules, and click **Save changes** to save.

![](<../../../../.gitbook/assets/image (184) (1) (1) (1) (1) (1).png>)

Your remote bundle of custom rules is now configured and will be used when testing IaC files.

#### Overriding a group's remote bundle configurations

By default, configuring a remote bundle for a group applies the remote bundle to all the organizations in the group. So if the group configurations are updated, these changes apply to all of its organizations.

However, an organization can still override the group's configurations and define its own bundle and tag. These will not change when the group updates its configurations.

In order to override the group's configurations, go to the organization's `Rules` section on the Infrastructure as Code Settings.

* Initially, the section is populated with the configurations inherited from the organization's group.

![](<../../../../.gitbook/assets/image (79) (1).png>)

* Update the configurations to those customized for your organization, and click **Save changes**.

![](<../../../../.gitbook/assets/image (83) (1) (1) (2).png>)

* Now, configurations on the group level will not override these customized settings for your organization.

You can restore the inheritance of group configurations at any time by using the `Reset to group default` button.

### Snyk API

If manually updating the settings through the Snyk Settings page is too time-consuming, another option is to use the Snyk API. It currently allows to send any variation of the custom rules settings via an API call.

* For example, in order to configure the custom rules bundle at the **group** level, call the [**Group IaC Settings API**](https://snykv3.docs.apiary.io/#reference/group-settings/infrastructure-as-code/update-infrastructure-as-code-settings) endpoint by providing the following body:

```
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
             "oci_registry_tag": "1.3.14",
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
             "oci_registry_tag": "1.3.14"
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

The API will reply with the group's settings, so you can confirm the changes:

```
{
  "type": "iac_settings",
  "id": "<group id>",
  "attributes": {
    "custom_rules": {
      "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
      "oci_registry_tag": "1.3.14",
      "is_enabled": true
    },
   "updated": "2021-11-27T11:34:33.132Z"
  }
```

#### Overriding a group's remote bundle configurations

Similarly to the Settings page, the [**Group IaC Settings API**](https://snykv3.docs.apiary.io/#reference/group-settings/infrastructure-as-code/update-infrastructure-as-code-settings) \*\*\*\* applies the remote bundle to all the organizations in the group. An organization can override the group's configurations and define its own bundle and tag by using an API call.

* To override the group's configurations, call the [**Org IaC Settings API**](https://snykv3.docs.apiary.io/#reference/organization-settings/infrastructure-as-code/update-infrastructure-as-code-settings) endpoint by providing a different custom rules bundle and tag in the request body:

```
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "oci_registry_url": "registry-1.docker.io/org-account/org-bundle-imageage",
             "oci_registry_tag": "1.3.15",
             "is_enabled": true
           }
       }
   }
}
```

* The API replies with the organization's settings, and the group settings under the `parents` section, so you can compare the two:

```
{
  "type": "iac_settings",
  "id": "<org id>",
  "attributes": {
    "custom_rules": {
      "oci_registry_url": "registry-1.docker.io/org-account/org-bundle-image",
      "oci_registry_tag": "1.3.15",
      "is_enabled": true
    },
   "updated": "2021-11-27T11:34:33.132Z",
   "parents": {
      "group": {
        "id": "<group id>",
        "type": "iac_settings",
        "attributes": {
          "custom_rules": {
            "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
            "oci_registry_tag": "1.3.14",
            "is_enabled": true
          },
          "updated": "2021-11-19T10:59:45.259Z"
        }
      }
    }
  }
```

* To revert back to the group settings, call the API by providing the following request body:

```
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "inherit_from_parent": "group"
           }
       }
   }
}
```

* The API replies with the organization's settings, and the group settings under the `parents` section, so you can compare the two:

```
{
  "type": "iac_settings",
  "id": "<org id>",
  "attributes": {
    "custom_rules": {
      "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
      "oci_registry_tag": "1.3.14",
      "is_enabled": true,
      "inherit_from_parent": "group"
    },
   "updated": "2021-11-19T10:59:45.259Z",
   "parents": {
      "group": {
        "id": "<group id>",
        "type": "iac_settings",
        "attributes": {
          "custom_rules": {
            "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
            "oci_registry_tag": "1.3.14",
            "is_enabled": true
          },
          "updated": "2021-11-19T10:59:45.259Z"
        }
      }
    }
  }
```

### Environment variables

You can also configure the location of the custom rules bundle using Snyk config for your organization. In your project’s folder, run the following commands to configure your container registry with the Snyk IaC CLI:

```
snyk config set oci-registry-url=registry-1.docker.io/org-account/org-bundle-image:1.3.14
```

This will set the following Snyk environment variable: `SNYK_CFG_OCI_REGISTRY_URL`

{% hint style="info" %}
Ensure the OCI Registry URL is a valid URL; for example, for DockerHub:

`registry-1.docker.io/org-account/org-bundle-image:1.3.14`

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

If you have found a discrepancy that you cannot explain, raise a support ticket.
