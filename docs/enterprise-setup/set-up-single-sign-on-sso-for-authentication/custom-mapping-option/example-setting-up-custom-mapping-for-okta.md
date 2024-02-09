# Example: Setting up custom mapping for Okta

The following shows two different options for [Custom Mapping Options](../../using-single-sign-on-sso-for-authentication/custom-mapping-option/) of Okta roles.

{% hint style="info" %}
For both of these options to work, the Snyk SSO application **must** be assigned at the group level, not the user level.
{% endhint %}

## Option 1: Assign Custom Mapping with user roles

The following describes Custom Mapping with user roles.

* Okta groups are mapped to Snyk Organizations.
* User profiles are tied to each user, not the Okta group.

The steps follow.

{% hint style="info" %}
Note: The Snyk SSO application **must** be assigned at the group level, not the user level.

When you look at a user's Application assignment it should look similar to the image that follows, with set by **Snyk-Roles group** following **Add Another** grayed out.
{% endhint %}

<figure><img src="../../../.gitbook/assets/snyk-roles.png" alt="Roles set on group level"><figcaption><p>Roles set on group level</p></figcaption></figure>

### Create an app attribute that contains the names of your relevant Snyk Organizations

When your Snyk Groups and users have been set up, follow these steps:

1. On the main page of Okta select **Directory -> Profile Editor -> your Snyk SSO app**.
2. Select **+Add Attribute**.
3. In the corresponding fields, add the following details for the first Attribute:\
   **Data type**: string array\
   **Display name**: Snyk Orgs\
   **Variable name**: snyk\_orgs\
   **Group Priority**: Combine values across groups
4. When you are finished, select **Save and Add Another**.

### Create a second app attribute that contains roles

1. In the corresponding fields, add the following details for the second Attribute:\
   **Data type**: String\
   **Display name**: Snyk User Role\
   **Variable name**: user\_role\
   **Enum**: Check **Define enumerated list of values**\
   **Attribute members Collaborator:** collaborator or collab\
   **Attribute members Admin**: administrator or admin\
   **Attribute members GroupAdmin**: groupadmin\
   **Attribute required**: Yes\
   **Scope**: User personal
2. When you are finished, select **Save**.

### Assign the first attribute to your Okta groups

1. On the main page of Okta select **Directory -> Groups**.
2. Select a **Group**, navigate to the **Applications** tab, click **Assign** **application i**f not already assigned, and choose your Snyk SSO app. Then click on the **pencil** next to the displayed Snyk app.
3. In **Edit App Assignment** dialog, add the Snyk Organization name to associate with your Okta group (no spaces or capital letter(s)).
4. Repeat the preceding steps to assign the Snyk app to all your applicable Okta groups, modifying the Snyk Organization name as needed.

### Assign the second attribute to your users

1. On the main page of Okta select **Directory -> People.**
2. Select a **User,** navigate to the **Applications** tab, and click on the **pencil** next to the application.
3. Select the right user type in group (user role): **Collaborator**, **Admin**, or **Group Admin**.

### Construct a value expression that concatenates these two attributes into string values in a roles array to be sent to Snyk

1. Navigate to **Applications -> Applications** and click on the **Snyk app** you configured.
2. Select **General Tab -> SAML Settings -> Edit** and click **next** to go to the Configure SAML step.
3. Select **Attribute Statements** and add an attribute named **roles** with **Name format** **Unspecified** and the **Value** in the following expression:\
   `appuser.user_role == "groupadmin" ? "snyk-groupadmin" : Arrays.flatten(String.replace(String.replace(String.append("snyk-",String.append(Arrays.toCsvString(appuser.snyk_orgs),"-"+appuser.user_role)),",",",snyk-"),",","-"+appuser.user_role+","))`
4. Click **Next -> Finish.**
5. Reach out to your Snyk point of contact so they can complete the configuration. This process may take four to five days.

The following explains the roles expression:

* If the role is groupadmin, the expression ignores everything else and passes `snyk-groupadmin`.
* If the role is not groupadmin, then for each Snyk Org name listed across all groups the expression automatically concatenates the prefix “`snyk-`” with the Snyk Org name and appends `user_role` at the end of each org name. An example result follows.

Example roles are:`[ "snyk-groupadmin", "snyk-org1-admin", "snyk-org2-admin" ]`

## Option 2: Assign Custom Mapping with Groups

In this configuration:

* Okta groups are mapped to Snyk Organizations.
* Okta groups are mapped to Snyk Organization membership roles.
* The user role in Snyk is pre-set in each Okta group for all members of that group.

The steps follow.

{% hint style="info" %}
Note: The Snyk SSO application must be assigned at the group level, not the user level.

When you look at a user's Application assignment it should look similar to the image that follows, with set by **Snyk-Roles group** following **Add Another** grayed out.
{% endhint %}

<figure><img src="../../../.gitbook/assets/snyk-roles.png" alt="Roles set on group level"><figcaption><p>Roles set on group level</p></figcaption></figure>

### Create a single app attribute containing both the Snyk Organization name and role

1. On the main page of Okta, select **Directory -> Profile Editor -> your Snyk SSO app**.
2. Select **+Add Attribute**.
3. In the corresponding fields, add the following details for this Attribute:\
   **Data type**: string array\
   **Display name**: Snyk Orgs\
   **Variable name:** snyk\_orgs\
   **Group Priority**: Combine values across groups
4. When you are finished select **Save.**

### Assign the attribute to the relevant Okta groups

1. On the main page of Okta select **Directory -> Groups**.
2. Select a **Group**, navigate to the **Applications** tab, click **Assign** **application** if not already assigned, and choose your Snyk SSO app,. Then click on the **pencil** next to the displayed Snyk SSO app.
3. In the **Edit App Assignment** dialog, add the Snyk org name + role associated with your Okta group (no spaces or capital letter(s)), following the syntax explained in [Example roles array mapping](../../using-single-sign-on-sso-for-authentication/custom-mapping-option/#example-roles-array-mapping). Example, `snyk-org-role`.
4. Repeat the preceding steps for all your applicable Okta groups to assign the org name and role combination to each user within each configured group.

### Construct a value expression that creates a roles array to be sent to Snyk

1. Navigate to **Applications -> Applications** and click on the **Snyk app** you configured.
2. Select **General Tab -> Edit SAML Settings** and click **next** to go to the **Configure SAML** step.
3. Add an **Attribute Statement** named “roles” of an unspecified type.
4.  Select **Attribute Statements** and set **roles** as the **Name** field with **Name format** **Unspecified** and the **Value** in the following expression:

    `Arrays.flatten(appuser.snyk_orgs)`
5. Reach out to your Snyk point of contact so they can complete the configuration.
