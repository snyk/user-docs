# Group and Organization navigation

Snyk shows your [preferred Organization](https://docs.snyk.io/features/user-and-group-management/managing-groups-and-organizations/manage-snyk-organizations#setting-your-preferred-organization) by default when you log into the Snyk Web UI. Snyk also uses the settings for your preferred Organization when you test a Project locally using the CLI.

## Switch Group

If your company uses multiple Groups, then to change which Organization you're viewing, click the Group navigation switcher (![](<../../.gitbook/assets/image (4) (3) (2).png>) icon):

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-04-25 at 10.06.46.png" alt="Switch Group"><figcaption><p>Switch Group</p></figcaption></figure>

</div>

{% hint style="info" %}
To view your personal Organizations, select **Ungrouped**.
{% endhint %}

## Switch Organization

To change which Organization you're viewing, click the Organization navigation switcher:

<div align="left">

<img src="../../.gitbook/assets/Screenshot 2023-03-13 at 10.31.14.png" alt="Switch Organizations">

</div>

{% hint style="info" %}
If you add Projects through GitHub integration, these Projects are added to the currently chosen Organization.
{% endhint %}

### **Switch Organization In the CLI**

1. If you have only your default Organization, any Projects you add or update by running `snyk monitor` are automatically associated with your default Organization.
2. If you have more than one Organization, you can configure the organization with which newly added Projects should be associated by running `snyk config set org=ORG_ID`.
3. If you would like to override this global configuration for individual runs of `snyk monitor`, run `snyk monitor --org=ORG_ID`.

The default `<ORG_ID>` is the currently preferred Organization in your [Account settings](https://app.snyk.io/account).

See [How to select the organization to use in the CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli.md) for more details.

