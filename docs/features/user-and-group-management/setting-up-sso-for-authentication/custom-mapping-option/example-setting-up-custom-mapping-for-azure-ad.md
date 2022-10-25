# Example: Setting up custom mapping for Azure AD

The following information shows how to configure custom mapping of roles for Azure AD using the [Custom Mapping Option](./).

There are two ways of implementing custom mapping with Azure AD, either using [custom claims](example-setting-up-custom-mapping-for-azure-ad.md#configuration) or [App Roles](example-setting-up-custom-mapping-for-azure-ad.md#configuration-app-roles). Choose the one that suits your situation best.&#x20;

If you need guidance setting up the initial Enterprise application, refer the Azure AD Enterprise Application [example](../self-serve-single-sign-on-sso/example-azure-ad-enterprise-application.md). **Note**: any step on the Snyk side in setting up the Enterprise application must be performed by your Snyk contact as self-serve SSO does not accommodate custom mapping.

## Configuration - custom claims

In this configuration:

* Azure AD Security groups are mapped to Snyk organizations.
* Azure AD Security is mapped to Snyk Organization membership roles.
* The user role in Snyk is pre-set in each Azure AD Security Group for all members of that group.

Once you have set up groups and users, follow these steps:

1.  In your Snyk App in Azure AD, open the Single Sign On Settings: **Dashboard -> Enterprise Applications -> Snyk -> Single Sign On.**

    <img src="../../../../.gitbook/assets/Screen Shot 2022-06-08 at 8.22.43 AM.png" alt="Snyk App Azure AD single sign on settings" data-size="original">
2.  Edit Attributes and Claims.

    <img src="../../../../.gitbook/assets/Screen Shot 2022-06-08 at 8.26.20 AM.png" alt="Attributes and Claims" data-size="line">
3.  Add new Claim

    <img src="../../../../.gitbook/assets/Screen Shot 2022-06-08 at 8.28.37 AM.png" alt="Add new claim" data-size="original">
4.  Configure Attribute.

    <img src="../../../../.gitbook/assets/Screen Shot 2022-06-08 at 8.32.50 AM.png" alt="Configure attribute" data-size="original">

    * **Name**: roles
    * Expand the Claim conditions section:
      * For each unique value, that is, unique combinations of security groups, you will need to create a new condition. As a reminder, each security group reflects a unique combination of organization membership and user role. **Order is important**. If you have more than one condition with the same group(s) included in 'scoped groups' the conditions will be evaluated top to bottom and the last value that includes the group(s) will be used. For this reason, It is suggested to have the conditions in increasing order of scoped groups.
      * Set User type to Members\
        ![Members user type](<../../../../.gitbook/assets/Screen Shot 2022-06-08 at 9.19.38 AM.png>)
      * The scoped groups should be the security groups to which you are assigning one or more Org membership and user role combinations.\
        ![Selected groups](<../../../../.gitbook/assets/select groups.png>)
      * Select Attribute as the Source
      *   Set the Value to the Snyk Org and user role slugs in the following format:\
          snyk-orgslug-role

          * For more than one, separate by comma.
          * There should be no spaces or capital letter(s) in the Org and user role slugs.
          * Do not include double quotes as Azure AD automatically adds them.

          <img src="../../../../.gitbook/assets/Screen Shot 2022-06-08 at 9.20.22 AM.png" alt="snyk-orgslug-role" data-size="original">
      * Repeat for each claim condition.

## Configuration - App roles

Prerequisites for this configuration type:

* Snyk support must configure your Snyk SSO configuration as a Microsoft Azure AD (WAAD, not SAML)&#x20;
* You must have an existing Azure Enterprise application and app registration connected to that SSO configuration

1.  In your app registration menu, select the name of your Enterprise Application.

    <figure><img src="../../../../.gitbook/assets/image.png" alt="App registration, name of Enterprise Application"><figcaption><p>App registration, name of Enterprise Application</p></figcaption></figure>
2.  Select App roles, then Create app role.

    <figure><img src="../../../../.gitbook/assets/image (1).png" alt="App roles, Create app role"><figcaption><p>App roles, Create app role</p></figcaption></figure>
3.  Create an app role with details as needed.

    <figure><img src="../../../../.gitbook/assets/image (2).png" alt="Create app role with details"><figcaption><p>Create app role with details</p></figcaption></figure>
4.  Go into Azure AD and select your Enterprise Application.

    <figure><img src="../../../../.gitbook/assets/image (3).png" alt="Enterprise Application"><figcaption><p>Enterprise Application</p></figcaption></figure>
5.  Select Users and groups; then Add user/group.

    <figure><img src="../../../../.gitbook/assets/image (4).png" alt="Users and groups, Add user/group"><figcaption><p>Users and groups, Add user/group</p></figcaption></figure>
6.  Select Users and groups, then Select a role and finally Assign.

    <figure><img src="../../../../.gitbook/assets/image (5).png" alt="Add assignment"><figcaption><p>Add assignment</p></figcaption></figure>
7.  Repeat for all required groups and roles that should be assigned. Then verify that the list looks similar to this.

    <figure><img src="../../../../.gitbook/assets/image (6).png" alt="Users and group list"><figcaption><p>Users and group list</p></figcaption></figure>
