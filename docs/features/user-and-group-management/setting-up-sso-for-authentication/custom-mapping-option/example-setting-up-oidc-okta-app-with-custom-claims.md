# Example: Setting up OIDC Okta app with custom claims

The steps follow to configure an integration for OIDC Okta.

## Create an Okta OIDC app

In Okta, select **Applications -> Applications -> OICD OpenID Connect**.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_6_30_22__5_01_PM.png" alt="Create a new app integration in Okta"><figcaption><p>Create a new app integration in Okta</p></figcaption></figure>

## Enter the sign-in redirect URI

<figure><img src="../../../../.gitbook/assets/Enter-redirect-URI-OIDC-Okta.png" alt="Enter the Sign-in redirect URL for the new web app integration"><figcaption><p>Enter the Sign-in redirect URL for the new web app integration</p></figcaption></figure>

Find the [OIDC information to provide to Snyk](https://docs.snyk.io/features/user-and-group-management/setting-up-sso-for-authentication/set-up-snyk-single-sign-on-sso#oidc-information-to-provide-to-snyk): Issuer URL, Client ID, Client Secret, Email domains and subdomains. The information follows \[\[where?]].

## Add Roles in Okta

### Add OIDC claim at the user level

On the Okta main page, select **Directory -> People**.

From the **Person & username list**, select a user.

<figure><img src="../../../../.gitbook/assets/OIDC-claim-steps-1-2.png" alt="Select an Okat user"><figcaption><p>Select an Okat user</p></figcaption></figure>

For the selected user, open the **Profile tab** and select **Edit**.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_14_22__12_26_PM.png" alt="Edit the profile of the selected user"><figcaption><p>Edit the profile of the selected user</p></figcaption></figure>

Select **Add Another Role** and enter the name of that role and select it.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_14_22__12_27_PM.png" alt="Specify the role"><figcaption><p>Specify the role</p></figcaption></figure>

### Add OIDC claim at the group level

On the Okta main page, select **Security -> API**.

On the API tab, enter the Authentication Server, api://snyk.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_11_22__6_12_PM.png" alt="Enter the authentication server"><figcaption><p>Enter the authentication server</p></figcaption></figure>

On the Okta main page, select **Directory -> Groups -> Profile Editor**; in the Profile Editor, select the **Groups** tab.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_23_PM.png" alt="Navigate to the Groups tab in the Profile Editor"><figcaption><p>Navigate to the Groups tab in the Profile Editor</p></figcaption></figure>

On the **Profile Editor Groups** tab, enter the name of the Okta group.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_24_PM.png" alt="Enter the name of the Okta group"><figcaption><p>Enter the name of the Okta group</p></figcaption></figure>

Open the **Okta group** details and select **Add attribute**.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_26_PM.png" alt="Select Add attribute for Okta group"><figcaption><p>Select Add attribute for Okta group</p></figcaption></figure>

Enter the values for the group attributes, here **Data type**, string array; **Display name**, snyk-orgs; **Variable name**, snyk\_orgs; **Description**, List of the Snyk orgs and permissions for user.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_28_PM.png" alt="Enter the values for the Group attribute"><figcaption><p>Enter the values for the Group attribute</p></figcaption></figure>

From the Okta group details, select **Directory -> Groups**.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_31_PM.png" alt="Select Directory -> Groups"><figcaption><p>Select Directory -> Groups</p></figcaption></figure>

Select a group or **Add Group**.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_32_PM.png" alt="Select a group or Add Group"><figcaption><p>Select a group or Add Group</p></figcaption></figure>

On the Snyk Admin screen, select the **Profile** tab.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_35_PM.png" alt="Select Profile tab on Snyk Admin screen"><figcaption><p>Select Profile tab on Snyk Admin screen</p></figcaption></figure>

On the **Profile** tab, select **Edit**.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_36_PM.png" alt="Select Edit on the Profile tab"><figcaption><p>Select Edit on the Profile tab</p></figcaption></figure>

On the Profile tab select **Add Another** and for the **snyk\_orgs Attribute**, Enter a string.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_15_22__5_36_PM (1).png" alt="Add snyk_orgs Attribute"><figcaption><p>Add snyk_orgs Attribute</p></figcaption></figure>

## Create a new claim

On the Okta main page, select **Security -> API**.&#x20;

On the **default** page, select the **Claims** tab and on the tab, select **Add Claim**.

<figure><img src="../../../../.gitbook/assets/Pasted_Image_7_11_22__6_19_PM.png" alt="Create a new claim"><figcaption><p>Create a new claim</p></figcaption></figure>

## Set attributes for the new claim

On the **Edit Claim** page, enter the values for the attributes:\
**Name**: roles\
**Include in token type**: ID Token, Always\
**Value type**: Expression\
Value: Add your attribute, appsuser.roles in the example that follows\
**Include in**: Select the scope that will be passed to Snyk. Any scope in the example that follows. You can also enter scopes.

When you are finished, select **Save**.

<figure><img src="../../../../.gitbook/assets/Untitled (2).png" alt="Example attributes for a claim"><figcaption><p>Example attributes for a claim</p></figcaption></figure>



