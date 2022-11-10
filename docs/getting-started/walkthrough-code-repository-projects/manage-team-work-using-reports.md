# Manage team work using Reports

{% hint style="info" %}
**Recap**\
You have seen how to view your Snyk Projects, understand Snyk scan results, fix vulnerabilities, and integrate fix work into your development workflow. We'll now look at how to monitor this work in more detail, using Snyk Reports.
{% endhint %}

## View reports

{% hint style="info" %}
**Feature availability**\
Reports are available with Business and Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

Click **Reports** to access the vulnerability information for all Snyk Projects in your Organization:

![An overview of Snyk Project vulnerabilities in Reports](../../.gitbook/assets/reports\_02oct2022.png)

{% hint style="info" %}
By default, results are shown for the last 90 days: use the **Show report for:** list in the top right to change the duration.
{% endhint %}

See [Reports](../../features/snyk-reports/) for more details.

### Security issues

The **Security issues** section shows the number of vulnerabilities, their type, and how many of them Snyk identifies as auto-fixable (with an upgrade, as we saw in [Fix your first vulnerability](fix-your-first-vulnerability.md) topic).

![](<../../.gitbook/assets/image (38).png>)

### Show issues over time

The **Issues over time** graph shows the history of vulnerabilities in your Organization:

![](<../../.gitbook/assets/image (103) (2) (1) (1) (1).png>)

This shows the number of overall vulnerabilities, and the number of projects being scanned (the dotted line in the graph).

#### Why do numbers increase?

Seeing more vulnerabilities over time may not reflect your team's work to improve security.

For example, the total number of vulnerabilities is likely to rise when you add more projects. The graph may display that linkage - when you see jumps in the total numbers of vulnerabilities, you may see corresponding jumps in the total numbers of Projects:

![More Projects, more issues](<../../.gitbook/assets/image (22) (2).png>)

#### Metric: ratio of vulnerabilities to Projects

Instead of measuring at the total numbers of issues, you can compare the number of vulnerabilities to the number of projects, and use this ratio as a measure of security. For example, if you double the number of projects you scan, but only add 10% to the total number of vulnerabilities, your general security is likely to improve.

Alternatively, if you see an increase in the number of vulnerabilities without an increase in the number of projects, this may be because a new vulnerability is discovered in an existing open-source library. Again, this has happened independently of your team's security work.

### Viewing snapshot summary

You can hover over a date to see the summary of “to this date” information at that point:

![](<../../.gitbook/assets/image (320) (1) (1).png>)

This is especially useful when you see sudden changes in issue numbers on a specific day.

### View activity

The **Activity** section shows the activity over the report period:

![](<../../.gitbook/assets/image (261).png>)

For the reporting period (90 days by default), this activity shows:

* **Tests Run**: the number of tests run. By default, Snyk scans each open source Project daily, so an Organization with 100 projects would generate 9,000 scans over 90 days.
* **New issues**: new issues detected.
* **Fixed issues**: the vulnerabilities fixed by your team.
* **Tests preventing issues**: occasions when the team attempted to merge code changes, but Snyk scans informed the team that these changes would have created new issues, so helping prevent new security issues.
* **Ignored issues**: a team member decided to ignore that issue.

### Filtering search results

If you have lots of Projects files to manage and organize in your Organization, you can use filters to focus on specific Projects or specific types of vulnerability:

![](<../../.gitbook/assets/image (71).png>)

For example, if your Organization represents your development team, and you want to focus on (say) front-end work in the next Sprint, click the **Projects** dropdown and select a subset of the front-end Projects to scan, allowing you to focus on fixing vulnerabilities in that area.

{% hint style="info" %}
**Why so many Projects?**\
Remember, a Snyk Project represents a single item that Snyk scans, such as a manifest file. So your application may contain hundreds of Snyk Projects to scan.
{% endhint %}

For Open Source vulnerabilities, we also have tagging, which lets you add your own tags, including custom values for metadata. See [Project tags](../../snyk-web-ui/introduction-to-snyk-projects/project-tags.md) and [Project attributes](../../snyk-web-ui/introduction-to-snyk-projects/project-attributes.md).

#### Dashboard results and report results

Filtering results in Reports means that your Dashboard shows different numbers to your Reports.

For example, by default, Snyk does not scan the Dockerfile in the filter, it scans the docker images themselves

![](<../../.gitbook/assets/image (111) (1) (1) (1) (1) (1) (1).png>)

By default, the Dockerfile option is not selected in the filter: when Snyk scans the Dockerfile, there will be vulnerabilities in the base OS in the container that you are building. Because these vulnerabilities cannot generally be fixed by the development team, this filtering lets your team focus on the issues that they _can_ fix.

{% hint style="info" %}
To see a report across all your Organizations, navigate to the Snyk Group level and look at reports there.
{% endhint %}

### Reports: Issues

Click the **Issues** tab to see a full list of all issues for your Organization:

![](<../../.gitbook/assets/image (108) (1).png>)

Issues are ranked by their Snyk Priority Score; you can also filter based on columns. For example, you may want to look at the highest scores with the most maturity (which are likely to have more exploits). You can also track the list of issues that got fixed.

Click **Export** to export or print these results.

{% hint style="warning" %}
You cannot currently export charts and data in the reports summary page.
{% endhint %}

### What's next?

This concludes this walkthrough.

Please refer to the Snyk documentation in general for more information.
