# Dependencies and licenses

You can [view dependencies](view-dependencies.md) and [license information](view-licenses.md) for all Projects in your Group or Organization using the **Dependencies** option in your Group or Organization menu:

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-11 at 12.45.48.png" alt="Dependencies tab for an Organization"><figcaption><p>Dependencies tab for an Organization</p></figcaption></figure>

{% hint style="info" %}
When you import or re-test a Project, changes will be reflected on the **Dependencies** UI after a ten-second delay.
{% endhint %}

For both dependencies and licenses, you can filter by Project or other filter criteria:

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-05-11 at 13.11.22.png" alt="Select Projects and filters"><figcaption><p>Select Projects and filters</p></figcaption></figure></div>

* From the **Projects** dropdown, select specific Projects.
* From the **Filters** dropdown, check the applicable boxes to filter by [Severity level](../prioritize-issues-for-fixing/severity-levels.md) or Project type.

{% hint style="info" %}
Results from the Dockerfile Project type are filtered out by default in the filter criteria as they can result in duplication of results from scans of the images resulting from building the Dockerfiles. To match results from [AP](../../snyk-api/reference/reporting-api-v1.md)I calls, either filter out Dockerfiles from the API results or turn on Dockfiles in the Project type column of the filter.
{% endhint %}
