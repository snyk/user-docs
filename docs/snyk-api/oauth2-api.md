# OAuth2 API

Snyk provides an OAuth2 API, primarily for use with [snyk-apps](../snyk-api-info/snyk-apps/ "mention"). It is compliment with RFC 6749.

Most endpoints are served from our API subdomain (eg https://api.snyk.io) with the one exception being `/oauth2/authorize` which is served on the main app subdomain (eg https://app.snyk.io.)

{% swagger src="../.gitbook/assets/oauth-app-spec.yaml" path="/oauth2/authorize" method="get" %}
[oauth-app-spec.yaml](../.gitbook/assets/oauth-app-spec.yaml)
{% endswagger %}

{% swagger src="../.gitbook/assets/oauth-api-spec.yaml" path="/token" method="post" %}
[oauth-api-spec.yaml](../.gitbook/assets/oauth-api-spec.yaml)
{% endswagger %}

{% swagger src="../.gitbook/assets/oauth-api-spec.yaml" path="/revoke" method="post" %}
[oauth-api-spec.yaml](../.gitbook/assets/oauth-api-spec.yaml)
{% endswagger %}
