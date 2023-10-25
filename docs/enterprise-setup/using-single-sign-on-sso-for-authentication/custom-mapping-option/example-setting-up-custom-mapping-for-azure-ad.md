# Example: Setting up custom mapping for Azure AD

The following information shows how to configure the custom mapping of roles for Azure AD using the [Custom Mapping Option](./).

{% hint style="info" %}
See the [Azure AD Enterprise Application example](../self-serve-single-sign-on-sso/example-azure-ad-enterprise-application.md) for guidance setting up the initial Enterprise application.
{% endhint %}

{% hint style="warning" %}
Any step on the Snyk side in setting up the Enterprise application must be performed by your Snyk contact, as self-serve SSO does not accommodate custom mapping.
{% endhint %}

The following are the **prerequisites** for **configuring App roles**:

* Snyk support must configure your Snyk SSO as Microsoft Azure AD (WAAD or SAML).
* If you select SAML, there is a requirement to add a custom claim; the step to do that is in these instructions.
* You must have an existing Azure Enterprise application and app registration connected to that SSO configuration.

The **steps** in **configuring App role**s follow.

1.  In your App registration menu, select the name of your Enterprise Application.

    <figure><img src="../../../.gitbook/assets/image (113) (1).png" alt="App registration, select name of Enterprise Application"><figcaption><p>App registration, select name of Enterprise Application</p></figcaption></figure>
2.  Select **App roles**, then **Create app role**.

    <figure><img src="../../../.gitbook/assets/image (1) (1) (2) (1) (1).png" alt="Select App roles, Create app role"><figcaption><p>Select App roles, Create app role</p></figcaption></figure>
3.  Create an app role with details as needed.\
    Select the **Allowed member types**: **Users/Groups**, **Applications**, or **Both**.\
    Enter the **Value** and **Description** for the selected type.\
    Enable the app role.\
    When you are finished, select **Apply**.

    <figure><img src="../../../.gitbook/assets/image (2) (2) (1).png" alt="Create app role with details"><figcaption><p>Create app role with details</p></figcaption></figure>
4.  In Azure AD, select your Enterprise Application.

    <figure><img src="../../../.gitbook/assets/image (3) (2) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Select Enterprise Application in Azure AD"><figcaption><p>Select Enterprise Application in Azure AD</p></figcaption></figure>
5.  Select **Users and groups**; then **Add user/group**.\
    Search and select the users and groups to add.

    <figure><img src="../../../.gitbook/assets/image (4) (5).png" alt="Select Users and groups, Add user/group"><figcaption><p>Select Users and groups, Add user/group</p></figcaption></figure>
6.  Select **Users and groups**; from the dropdown, select a role; and select **Assign**.

    <figure><img src="../../../.gitbook/assets/image (5) (1) (1) (1).png" alt="Add assignment"><figcaption><p>Add assignment</p></figcaption></figure>
7.  Repeat for all required groups and roles that should be assigned. Then verify that the list looks similar to this.

    <figure><img src="../../../.gitbook/assets/image (6) (1) (1) (2) (1) (1) (1) (1) (1).png" alt="Users and group list"><figcaption><p>Users and group list</p></figcaption></figure>

    Note that it is also possible add multiple Snyk roles to one App role, as the payload can be interpreted as a comma-separated string. However, this can not be used in conjunction with multiple App roles, as only one syntax will be respected (string or array).
8. If you have configured a SAML connection, add a custom claim to pass the roles array in the SAML payload to Snyk. Select **Single sign-on** in the left hand menu.
9.  Select **Edit** next to **Attributes and Claims.**

    <figure><img src="../../../.gitbook/assets/Screenshot 2023-03-10 at 3.19.31 PM.png" alt="Edit attributes and claims"><figcaption><p>Edit attributes and claims</p></figcaption></figure>
10. Select **Add new claim** add the following details, and **Save.**\
    **Name**: roles\
    **Source**: Attribute\
    **Source attribute**: user.assignedroles

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-10 at 2.55.05 PM.png" alt="Custom claim"><figcaption><p>Custom claim</p></figcaption></figure>

When you have completed these steps, reach out to your Snyk point of contact to have the configuration completedd.
