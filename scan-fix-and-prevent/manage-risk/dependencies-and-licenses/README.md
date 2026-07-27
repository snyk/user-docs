---
description: How to view dependencies and licenses for your Snyk Projects
nav_context: new
---

# Dependencies and licenses

You can [view dependencies](view-dependencies.md) and [license information](view-licenses.md) for all Projects in your Group or Organization using the **Dependencies** option in your Group or Organization menu.

{% hint style="info" %}
When you import or re-test a Project, changes are reflected on the **Dependencies** UI after a ten-second delay.
{% endhint %}

For both dependencies and licenses, you can filter by Project or other filter criteria:

* From the **Projects** dropdown, select specific Projects.
* From the **Filters** dropdown, check the applicable boxes to filter by [Severity level](../prioritize-issues-for-fixing/severity-levels.md) or Project type.

{% hint style="info" %}
Results from the Dockerfile Project type are filtered out by default in the filter criteria as they can result in duplication of results from scans of the images resulting from building the Dockerfiles. To match results from [AP](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-api/reference/reporting-api-v1)I calls, either filter out Dockerfiles from the API results or turn on Dockfiles in the Project type column of the filter.
{% endhint %}

{% hint style="info" %}
**Snyk 2.0 (Early Access)**

In the Snyk 2.0 UI, information about dependencies and licenses is available under **Analytics** > **Reports**.

Snyk 2.0 introduces UI enhancements to the platform navigation and is available in Early Access. This is being rolled out gradually, so not all users see the new navigation at the same time.

If you are an existing user, you can switch between the new and classic navigation at any time using the toggle in your user profile menu. For more information, visit [Snyk 2.0 platform improvements](https://app.gitbook.com/s/L7HyJj9FsK1W4pNt8Gzl/snyk-2.0-platform-improvements).
{% endhint %}
