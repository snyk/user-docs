# Switch between Groups and Organizations

Snyk shows your preferred Organization by default when you log into the Snyk Web UI. Snyk also uses the settings for your preferred Organization when you test a Project using the CLI. For more information, see [Manage Organizations](organizations/create-and-delete-organizations.md).

## Switch Group

If your company has multiple Groups, you must be aware of the Group you are viewing. Each Group contains different Organizations and has different settings.

To navigate to a different Group, click the Group switcher icon and select a Group.

To view or add to your personal Organizations, select **Ungrouped**.

## Switch Organization in the Web UI

{% hint style="warning" %}
If your SSO email is provided by Google, then logging in with that email through the Google social provider is not the same as logging in using SSO. Using the Google social option creates a separate free account with the standard free user privileges. You can not switch between Organizations that are in different accounts.
{% endhint %}

You must also be aware of the Organization you are viewing. Organizations contain different Projects.

If you add Projects through GitHub integration, these Projects are added to the Organization you are viewing, that is, the Organization you have selected.

To navigate to a different Organization, click the Organization switcher and select an Organization.

{% hint style="info" %}
**Snyk 2.0 (Early Access)**

In the Snyk 2.0 UI, you can navigate between different levels of your account (Tenant, Group, and Organization) using the scope selector at the top of the page. When you select a scope, the side menu automatically displays the relevant tools and data for that area.

Snyk 2.0 introduces UI enhancements to the platform navigation and is available in Early Access. This is being rolled out gradually, so not all users see the new navigation at the same time.

If you are an existing user, you can switch between the new and classic navigation at any time using the toggle in your user profile menu. For more information, visit [Snyk 2.0 platform improvements](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/L7HyJj9FsK1W4pNt8Gzl/snyk-2.0-platform-improvements).
{% endhint %}

## **Switch Organization in the CLI**

1. If you have only your default Organization, any Projects you add or update by running `snyk test` or `snyk monitor` are automatically associated with your default Organization.
2. If you have more than one Organization, you can set the Organization with which new Projects should be associated by running `snyk config set org=ORG_ID`.
3. If you want to override this global configuration for individual runs of `snyk monitor`, run `snyk test --org=ORG_ID` or `snyk monitor --org=ORG_ID`.

The default `<ORG_ID>` is the currently preferred Organization in your [Account settings](https://app.snyk.io/account).

See [How to select the Organization to use in the CLI](https://app.gitbook.com/o/-M4tdxG8qotLgGZnLpFR/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli) for more information.
