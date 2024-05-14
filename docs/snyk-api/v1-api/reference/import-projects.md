# Import Projects

Import projects into Snyk. Projects can be Git repositories, Docker images, containers, configuration files and much more. See the [snyk-projects](../../../snyk-admin/snyk-projects/ "mention") for more information. A typical import would start with requesting a target to be processed and then polling the Import Job API for further details on completion and resulting Snyk projects.

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/org/{orgId}/integrations/{integrationId}/import" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/org/{orgId}/integrations/{integrationId}/import/{jobId}" method="get" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}
