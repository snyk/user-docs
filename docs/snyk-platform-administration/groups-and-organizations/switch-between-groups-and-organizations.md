# Switch between Groups and Organizations

Snyk shows your preferred Organization by default when you log into the Snyk Web UI. Snyk also uses the settings for your preferred Organization when you test a Project using the CLI. For more information, see [Manage Organizations](organizations/create-and-delete-organizations.md).

## Switch Group

If your company has multiple Groups, you must be aware of the Group you are viewing. Each Group contains different Organizations and has different settings.

To navigate to a different Group, click the Group switcher icon and select a Group:

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-04-25 at 10.06.46.png" alt="Switch Group"><figcaption><p>Switch Group</p></figcaption></figure></div>

To view or add to your personal Organizations, select **Ungrouped**.

## Switch Organization in the Web UI

{% hint style="warning" %}
If your SSO email is provided by Google, then logging in with that email through the Google social provider is not the same as logging in using SSO. Using the Google social option creates a separate free account with the standard free user privileges. You can not switch between Organizations that are in different accounts.
{% endhint %}

You must also be aware of the Organization you are viewing. Organizations contain different Projects.

If you add Projects through GitHub integration, these Projects are added to the Organization you are viewing, that is, the Organization you have selected.

To navigate to a different Organization, click the Organization switcher and select an Organization:

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-03-13 at 10.31.14.png" alt="Switch Organization"><figcaption><p>Switch Organization</p></figcaption></figure></div>

## **Switch Organization in the CLI**

1. If you have only your default Organization, any Projects you add or update by running `snyk test` or `snyk monitor` are automatically associated with your default Organization.
2. If you have more than one Organization, you can set the Organization with which new Projects should be associated by running `snyk config set org=ORG_ID`.
3. If you want to override this global configuration for individual runs of `snyk monitor`, run `snyk test --org=ORG_ID` or `snyk monitor --org=ORG_ID`.

The default `<ORG_ID>` is the currently preferred Organization in your [Account settings](https://app.snyk.io/account).

See [How to select the Organization to use in the CLI](../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli.md) for more information.
