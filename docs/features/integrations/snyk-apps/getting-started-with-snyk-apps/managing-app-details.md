# Managing App details

## List Apps

You can view a list of Snyk Apps owned by your Snyk Organization by sending a **GET** request to the `apps` endpoint (more details found in the [API documentation](https://snykv3.docs.apiary.io/#reference/apps/app-management/list-existing-apps)):

```
https://api.snyk.io/v3/orgs/{orgId}/apps?version={version}
```

You will not be able to view the **clientSecret** after the App is created. If you have misplaced it, you can [rotate your **clientSecret**](managing-app-details.md#rotate-app-client-secret) and receive a new one.

## Rotate App Client Secret

If your App **clientSecret** is leaked, you can generate a new one using your **clientId**. The clientId can be found using the [List Apps endpoint](managing-app-details.md#view-app-details).

{% hint style="warning" %}
When you rotate your **clientSecret**, your current secret is immediately invalid. Your App will not be able to connect to Snyk until you update the App's configuration with the new secret.
{% endhint %}

To rotate your **clientSecret**, send a **POST** request to the `apps/{clientId}/secrets` endpoint (more details found in the[ API documentation](https://snykv3.docs.apiary.io/#reference/apps/manage-app-secrets/rotate-secret)):

```
https://api.snyk.io/v3/orgs/{orgId}/apps/{clientId}/secrets?version={version}
```

## Update App name

To update your App's name, you can send a **PATCH** request to the `apps/{clientId}` endpoint (more details found in the [API documentation](https://snykv3.docs.apiary.io/#reference/apps/single-app-management/delete-app)). The clientId can be found using the [List Apps endpoint](managing-app-details.md#view-app-details).

```
https://api.snyk.io/v3/orgs/{orgId}/apps/{clientId}?version={version}
```

## Delete an App

If you want to delete an App from your Snyk Organization, you can send a **DELETE** request to the `apps` endpoint (more details found in the [API documentation](https://snykv3.docs.apiary.io/#reference/apps/single-app-management/delete-app)):

```
https://api.snyk.io/v3/orgs/{orgId}/apps?version={version}
```

{% hint style="danger" %}
Deleting an App will revoke your App credentials and remove all of your App's installations. If you have active users, they will no longer be able to connect to Snyk through the App.
{% endhint %}
