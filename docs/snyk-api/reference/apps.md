# Apps

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/self/apps" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/self/apps
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/self/apps/{app_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/self/apps/{app_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/self/apps/{app_id}/sessions" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/self/apps/{app_id}/sessions
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/self/apps/{app_id}/sessions/{session_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/self/apps/{app_id}/sessions/{session_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/self/apps/installs" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/self/apps/installs
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/self/apps/installs/{install_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/self/apps/installs/{install_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/{client_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/{client_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/{client_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/{client_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/{client_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/{client_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/{client_id}/secrets" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/{client_id}/secrets -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/installs" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/installs -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/installs" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/installs
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/installs/{install_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/installs/{install_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/installs/{install_id}/secrets" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/installs/{install_id}/secrets -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/creations" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/creations -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/creations" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/creations
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/creations/{app_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/creations/{app_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/creations/{app_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/creations/{app_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/creations/{app_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/creations/{app_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/apps/creations/{app_id}/secrets" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/apps/creations/{app_id}/secrets -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/app_bots" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/app_bots
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/app_bots/{bot_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/app_bots/{bot_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/apps/installs" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/apps/installs -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/apps/installs" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/apps/installs
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/apps/installs/{install_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/apps/installs/{install_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/apps/installs/{install_id}/secrets" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/apps/installs/{install_id}/secrets -X POST \
  --input body.json
```
