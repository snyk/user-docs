# Custom Base Images

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/custom_base_images" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/custom_base_images -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/custom_base_images" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/custom_base_images
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/custom_base_images/{custombaseimage_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/custom_base_images/{custombaseimage_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/custom_base_images/{custombaseimage_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/custom_base_images/{custombaseimage_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/custom_base_images/{custombaseimage_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/custom_base_images/{custombaseimage_id} -X DELETE
```
