# Example: setting up custom mapping for Google Workspace

The following shows how to use [custom mapping](../) to map roles for a Google Workspace custom SAML connection.

For additional details and guidance, see the [Google documentation, Manage Custom User Fields](https://developers.google.com/admin-sdk/directory/v1/guides/manage-schemas).

To use the API, either log in to your Google Workspace admin area and from there execute the commands that follow , or for flexibility and automation, use the API with your preferred API client or script. To generate API credentials, refer to the Google documentation, [Create access credentials](https://developers.google.com/workspace/guides/create-credentials).

## Add a user schema

Use the [schema endpoint](https://developers.google.com/admin-sdk/directory/reference/rest/v1/schemas/insert) to add a schema that can be tied to the users. An example schema follows. This schema makes it possible to expose the desired custom role mapping in the SAML payload for the user.

```json
{
   "fields":
   [
     {
       "fieldName": "roles",
       "fieldType": "STRING",
       "readAccessType": "ADMINS_AND_SELF",
       "multiValued": true,
       "displayName": "roles"
     }
   ],
   "schemaName": "Snyk-SSO"
 }
```

## Modify user profiles

Attach the desired roles to the user profile with the user [API endpoint](https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/patch). An example payload follows for reference.

```json
{
 "customSchemas": {
   "Snyk-SSO": {
     "roles": [
       {
         "value": "snyk:org:org1:org_admin"
       },
       {
         "value": "snyk:org:org2:org_admin"
       }
     ]
   }
 }
}
```

## Modify SAML attributes

To expose these roles in the SAML payload, you must modify the attributes in the SAML attribute mapping:

1.  Log in to your Google Workspace Admin area and go to **Apps** and then **Web and mobile apps** and open your application.

    <figure><img src="../../../../../.gitbook/assets/x1.png" alt="Open Google SAML app"><figcaption><p>Open Google SAML app</p></figcaption></figure>
2. Click on **SAML attribute mapping** and then **ADD MAPPING**.
3. Click **Select field** and scroll to the bottom until you find **Snyk-SSO - roles** and select it.
4.  In the **App attributes** value field, enter **roles** and click **Save**.

    <figure><img src="../../../../../.gitbook/assets/x2.png" alt="Adding custom mapping app attribute"><figcaption><p>Adding custom mapping app attribute</p></figcaption></figure>

After this, log in as a user and have your Snyk contact validate the SAML payload and finalize the setup in the Snyk backend.
