# Import Projects (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

You can use this API to import projects into Snyk. Projects imported can be Git repositories, Docker images, containers, configuration files, and much more. See the [Projects and Targets documentation](../../snyk-admin/snyk-projects/#target) for more information. A typical import would start with requesting a target to be processed and then polling the Import Job API for more details on completion of an import and the resulting Snyk Projects.

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}/import" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/integrations/{integrationId}/import/{jobId}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}
