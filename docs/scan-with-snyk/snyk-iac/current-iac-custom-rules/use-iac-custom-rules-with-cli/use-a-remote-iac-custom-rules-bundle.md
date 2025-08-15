# Use a remote IaC custom rules bundle

After you generate your custom rules bundle, you can distribute it to one of the supported OCI registries by following the steps in [Pushing a bundle](../writing-rules-using-the-sdk/pushing-a-bundle.md).

After successfully pushing your custom rules bundle, you can enforce the use of the bundle using any of the following:

* [Snyk settings](use-a-remote-iac-custom-rules-bundle.md#snyk-settings-and-remote-custom-rules-bundle)
* [Snyk API](use-a-remote-iac-custom-rules-bundle.md#snyk-api-and-remote-custom-rules-bundle)
* [Environment variables](use-a-remote-iac-custom-rules-bundle.md#environment-variables-and-remote-custom-rules-bundle)

Finally, after you have enforced your custom rules using one of these options, configure the Snyk Snyk CLI with your username and password to allow Snyk to authorize a pull from your OCI registry:

```
snyk config set oci-registry-username=<org registry username>
snyk config set oci-registry-password=<org registry password>
```

This sets the following Snyk environment variables:

* `SNYK_CFG_OCI_REGISTRY_USERNAME`
* `SNYK_CFG_OCI_REGISTRY_PASSWORD`

After you have completed this configuration, you can run a Snyk IaC scan. The CLI will pull the bundle pushed to the configured container registry in the background.

```
snyk iac test <file>
```

The resulting configuration scan issues will include issues from both the default Snyk rules and your custom rules. See also [Understanding the IaC CLI test results](../../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/understand-the-iac-cli-test-results/).

{% hint style="info" %}
Only one method for defining the bundle's path should be defined at any given time. Make sure to disable the custom rules settings using the Snyk settings page or the [Snyk API](../../../../snyk-api/reference/iacsettings.md). Alternatively, clear any previously stored settings using [`snyk config unset`](../../../../developer-tools/snyk-cli/commands/config.md#unset-less-than-key-greater-than).
{% endhint %}

## Snyk settings and remote custom rules bundle

Snyk recommends you use the Snyk settings page to configure custom rules settings. This is a simple way to update the custom rules bundle's URL and tag whenever these are modified.

{% hint style="info" %}
Tags are helpful for versioning your custom rules bundles. Configuring your updated bundle can be easily accomplished by setting the new version tag.
{% endhint %}

You can configure these remote bundles on both the Organization and Group levels. Configuring a remote bundle for a Group applies the remote bundle to all the Organizations in the Group.

To configure remote bundles:

* In the Infrastructure as Code Settings, locate the **Rules** section.

{% hint style="info" %}
You can configure remote custom rules bundles on the Organization level by navigating to `Settings` > `Infrastructure as Code`.

You can configure remote custom rules bundles on the Group level by navigating to `Settings` > `Infrastructure as Code.`
{% endhint %}

<figure><img src="../../../../.gitbook/assets/image (161) (1) (1) (1) (1) (1) (2).png" alt="Enable remote custom rules bundles"><figcaption><p>Enable remote custom rules bundles</p></figcaption></figure>

* Enable configuration of remote bundles by using the **Enable rules** toggle. Doing so loads the form to specify the Registry URL and tag as shown in this example:

<figure><img src="../../../../.gitbook/assets/image (102) (2).png" alt="Specify Registry URL and tag form"><figcaption><p>Specify Registry URL and tag form</p></figcaption></figure>

* Configure the OCI registry URL and tag for your remote bundle of custom rules and click **Save changes** to save.

<figure><img src="../../../../.gitbook/assets/image (19).png" alt="Registry URL and tag configured"><figcaption><p>Registry URL and tag configured</p></figcaption></figure>

Your remote bundle of custom rules is now configured and will be used when testing IaC files.

You can override remote bundle configurations for a Group using Snyk Settings.

By default, configuring a remote bundle for a Group applies the remote bundle to all the Organizations in the Group. Thus if the Group configurations are updated, these changes apply to all of its Organizations.

However, an Organization can still override the Group configurations and define its own bundle and tag. These will not change when the Group updates its configurations.

To override the Group configurations, go to the Organization's `Rules` section in the Infrastructure as Code Settings.

* Initially, the section is populated with the configurations inherited from the Organization's Group.

<figure><img src="../../../../.gitbook/assets/image (72).png" alt="Organization Rules section in the IaC Settings"><figcaption><p>Organization Rules section in the IaC Settings</p></figcaption></figure>

* Update the configurations to those customized for your Organization and click **Save changes**.

<figure><img src="../../../../.gitbook/assets/image (152).png" alt="Organization rules configuration updated"><figcaption><p>Organization rules configuration updated</p></figcaption></figure>

* Now, configurations on the Group level will not override these customized settings for your Organization.

You can restore the inheritance of Group configurations at any time by using the **Reset to group default** button.

## Snyk API and remote custom rules bundle

If manually updating the settings through the Snyk Settings page is too time-consuming, you can use the Snyk API, which allows you to send any variation of the custom rules settings using an API call.

* For example, in order to configure the custom rules bundle at the Group level, use the endpoint [Update the Infrastructure as Code settings for a group](../../../../snyk-api/reference/iacsettings.md#groups-group_id-settings-iac) by providing the following body:

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

* If you want to disable custom rules, you can send over the `is_enabled` flag:

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

The API replies with the Group settings so you can confirm the changes:

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

You can override remote bundle configurations using the Snyk API.

Similarly to the Settings page, the endpoint [Update the Infrastructure as Code settings for a group](../../../../snyk-api/reference/iacsettings.md#groups-group_id-settings-iac) allows you to apply the remote bundle to all the Organizations in the Group. An Organization can override the configurations for a Group and define its own bundle and tag by using an API call.

* To override the Group configurations, call the endpoint [Update the Infrastructure as Code settings for an org](../../../../snyk-api/reference/iacsettings.md#orgs-org_id-settings-iac) by providing a different custom rules bundle and tag in the request body:

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

* The API replies with the Organization settings and the Group settings under the `parents` section so that you can compare the two:

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

* To revert to the Group settings, call the API by providing the following request body:

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

* The API replies with the Organization settings and the Group settings under the `parents` section so that you can compare the two:

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
   "parents"  {
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

## Environment variables and remote custom rules bundle

You can also configure the location of the custom rules bundle using Snyk config for your Organization. In your Project folder, use the following command to configure your container registry with the Snyk IaC CLI:

```
snyk config set oci-registry-url=registry-1.docker.io/org-account/org-bundle-image:1.3.14
```

This sets the Snyk environment variable `SNYK_CFG_OCI_REGISTRY_URL.`

{% hint style="info" %}
Ensure the OCI Registry URL is a valid URL, for example, for DockerHub:

`registry-1.docker.io/org-account/org-bundle-image:1.3.14`

Be sure to clear any previously defined URLs in the Snyk Settings page or disable custom rules, as only one method for defining the bundle's path should be defined at any given time.
{% endhint %}

## Troubleshooting for remote custom rules bundle

Enable debug logs by running the command with the `-d` option:

```
snyk iac test <file> -d
```

Some possible problems include:

* Providing an invalid container registry URL. See the note above if you are using Docker Hub. The error is

```
We were unable to download the custom bundle to the disk. 
Please ensure access to the remote Registry and validate you have provided all the right parameters.
```

* Providing invalid credentials. The error is:

```
There was an authentication error. Incorrect credentials provided.
    We were unable to download the custom bundle to the disk.
    Please ensure access to the remote Registry and validate you have provided all the right parameters.
```

If you have found a discrepancy that you cannot explain, contact [Snyk Support](https://support.snyk.io).
