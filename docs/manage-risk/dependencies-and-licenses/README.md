# Dependencies and licenses

You can [view dependencies](view-dependencies.md) and [license information](view-licenses.md) for all Projects in your Group or Organization using the **Dependencies** option in your Group or Organization menu.

{% hint style="info" %}
When you import or re-test a Project, changes are reflected on the **Dependencies** UI after a ten-second delay.
{% endhint %}

For both dependencies and licenses, you can filter by Project or other filter criteria:

* From the **Projects** dropdown, select specific Projects.
* From the **Filters** dropdown, check the applicable boxes to filter by [Severity level](../prioritize-issues-for-fixing/severity-levels.md) or Project type.

{% hint style="info" %}
Results from the Dockerfile Project type are filtered out by default in the filter criteria as they can result in duplication of results from scans of the images resulting from building the Dockerfiles. To match results from [AP](../../snyk-api/reference/reporting-api-v1.md)I calls, either filter out Dockerfiles from the API results or turn on Dockfiles in the Project type column of the filter.
{% endhint %}
