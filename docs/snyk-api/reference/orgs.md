# Orgs

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/memberships" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/memberships" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/memberships/{membership_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/memberships/{membership_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/orgs" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}
