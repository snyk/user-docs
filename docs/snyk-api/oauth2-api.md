# OAuth2 API

Snyk provides an OAuth2 API, primarily for use with [Snyk Apps](how-to-use-snyk-apps-apis/). It complies with RFC 6749.

Most endpoints are served from the Snyk API subdomain (for example, https://api.snyk.io), with the one exception being `/oauth2/authorize` which is served on the main app subdomain (for example, https://app.snyk.io).

{% swagger src="../.gitbook/assets/oauth-app-spec.yaml" path="/oauth2/authorize" method="get" %}
[oauth-app-spec.yaml](../.gitbook/assets/oauth-app-spec.yaml)
{% endswagger %}

{% swagger src="../.gitbook/assets/oauth-api-spec.yaml" path="/token" method="post" %}
[oauth-api-spec.yaml](../.gitbook/assets/oauth-api-spec.yaml)
{% endswagger %}

{% swagger src="../.gitbook/assets/oauth-api-spec.yaml" path="/revoke" method="post" %}
[oauth-api-spec.yaml](../.gitbook/assets/oauth-api-spec.yaml)
{% endswagger %}
