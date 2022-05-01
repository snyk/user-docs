# Revoking and regenerating Snyk API tokens

{% hint style="info" %}
When an API token is revoked, all integrations using that key will immediately stop working. Proceed with caution!
{% endhint %}

If you suspect an API token has been leaked, it's good practice to revoke it and generate a new one to use in its place.

To revoke your Snyk user API token, navigate to your User Account Settings in the Snyk Web UI, at [app.snyk.io/account](https://app.snyk.io/account).

![api token screen; revoke; regenerate; click to show](<../../.gitbook/assets/uuid-8d94edf8-b42b-e5b3-ada1-e157d18ff884-en (1) (2) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (9).png>)

Click the **Revoke & Regenerate** button to revoke your API token. A new one will be generated in its place. You can now grab the newly generated API token and update integrations that used the old key.
