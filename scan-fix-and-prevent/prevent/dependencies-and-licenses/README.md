---
nav_context: new
---

# Dependencies and licenses

{% hint style="info" %}
**Note that this option** (viewing dependencies and licenses as a report across multiple projects in an organization/group) **is not available for Free and Team plan users.** Users of the Free plan can view a list of dependencies **per** [**project**](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects/view-project-issues-fixes-and-dependencies#view-dependency-details) only, and Team plan users will also see license information in that same view.
{% endhint %}

You can [view dependencies](view-dependencies.md) and [license information](view-licenses.md) for all Projects in your Group or Organization using the **Dependencies** option in your Group or Organization menu. <br>

{% hint style="info" %}
When you import or re-test a Project, changes are reflected on the **Dependencies** UI after a ten-second delay.
{% endhint %}

For both dependencies and licenses, you can filter by Project or other filter criteria:

* From the **Projects** dropdown, select specific Projects.
* From the **Filters** dropdown, check the applicable boxes to filter by [Severity level](../../fix/prioritize-issues-for-fixing/severity-levels.md) or Project type.

{% hint style="info" %}
Results from the Dockerfile Project type are filtered out by default in the filter criteria as they can result in duplication of results from scans of the images resulting from building the Dockerfiles. To match results from [Reporting API](https://docs.snyk.io/developer-tools/snyk-api/reference/reporting-api-v1) calls, either filter out Dockerfiles from the API results or turn on Dockerfiles in the Project type column of the filter.
{% endhint %}
