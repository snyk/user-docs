# Education reports

{% hint style="info" %}
The Learn Engagement report and the Learn Impact and Opportunities report are available only at the Group level.
{% endhint %}

The Education reports section includes the following reports:

* [Learn Engagement](education-reports.md#learn-engagement) report
* [Learning Impact & Opportunities](education-reports.md#learning-impact-and-opportunities) report

## Learn Engagement

{% hint style="info" %}
Learn Engagement report is available only in the Learning Management add-on offering. For more information, contact your Snyk account team.
{% endhint %}

The goal of the engagement report is to provide insights into your security education and training programs overall progress, and give you insights into which parts of your Organization are engaging with Snyk Learn content. You can use the data and insights to better optimise your program, find security champions, generate reports for compliance and show progress to your executive sponsors.

### Access the report

The Learn Engagement report can be accessed at the Group level from the **Change Report** dropdown in the Reports menu.

### Report features

The report allows you to track:

* Learn engagement snapshot analytics
* Assignment Progress
* Adoption rankings
* Content usage breakdown
* Filtering: custom time periods, users, organizations, organization role, and Lesson titles.

### Learn engagement snapshot and assignment progress

The first section of the report focuses on showing key engagement statistics and the progress of any assignments. Tool tips provide more details on the definitions of the metrics.

<figure><img src="../../../.gitbook/assets/Screenshot 2025-09-29 at 19.30.57 (1).png" alt=""><figcaption></figcaption></figure>

### Adoption rankings

The adoption ranking section shows your organization and individual user engagement with Snyk Learn. This is ranked by "Lessons complete" and also has the estimated duration the org/user has spent on Snyk Learn lessons. Estimated duration calculated using the estimated duration presented at the start of each lesson, and includes estimated time from any progress on "in-progress" lessons in the selected period.

{% hint style="info" %}
The user level adoption ranking is a great way to identify potential security champions who are proactively engaging in security education and training.
{% endhint %}

<figure><img src="../../../.gitbook/assets/Screenshot 2025-09-29 at 19.34.12.png" alt=""><figcaption></figcaption></figure>

### Learning breakdown

The breakdown shows the different types of Learn content the users are engaging with, using lesson completions as the measure. You can see if users are engaging with product training or security education, along with the most popular lessons and insights into which CWE categories users are studying the most.

<figure><img src="../../../.gitbook/assets/Screenshot 2025-09-29 at 19.34.22.png" alt=""><figcaption></figcaption></figure>

## Learning Impact & Opportunities

{% hint style="info" %}
Learning Impact & Opportunities report is in Early Access and available only in the Learning Management add-on offering. For more information, contact your Snyk account team.
{% endhint %}

The goal of the impact and opportunities report is to provide insights into the impact your security education and training programs are having on code issue remediation and code issue prevention. In addition, the report gives recommendations for future training based on your code issue backlog, and issues that were introduced during the selected time period of the report.

### Access the report

The Learning Impact & Opportunities report can be accessed at the Group level from the **Change Report** dropdown in the Reports menu.

### Report features

The report allows you to track:

* Impact of education and training on code issue remediation
* Impact of education and training on code issue prevention
* Recommendations for further training opportunities
* Coverage rates of users trained in identified training opportunities.
* Filtering: custom time periods, users, organizations, lesson title, CWE, issue severity.

### Learning impact snapshot

The first section of the report focuses on the impact education is having on your security program, focusing on code issue resolution and code issue prevention.

The "Learning Impact on Issue Resolution" chart measures the relationship between lesson completion and the resolution of detected code security issues. Resolved issues are counted when a related lesson was completed before the issue was fixed within the selected period. Lesson completions are counted when a related issue was fixed after the lesson was completed within the selected period. Use the filters to drill into specific lessons or CWE categories.

The "Learning Impact on Issue Prevention" chart measures the relationship between lesson completion and the prevention code security issues. Introduced issues are counted when a related lesson was completed within the selected period. Issues introduced on the day a Project was imported are not counted. Use the filters to drill into specific lessons or CWE categories.

<figure><img src="../../../.gitbook/assets/Screenshot 2025-10-23 at 15.00.57 (1).png" alt=""><figcaption></figcaption></figure>

### Top 10 CWEs - open issues / issues introduced in the period

This section of the report shows recommendation for training for your top open code issues, and most frequently introduced issues, by volume. Note issues are only included when Snyk Learn has a related lesson for the CWE category.

You will see coverage for all users within organisation scope of the report filters. This shows you how many people have ever completed a related Snyk Learn lesson on the topic.

<div><figure><img src="../../../.gitbook/assets/Screenshot 2025-10-23 at 14.12.24 (1).png" alt=""><figcaption></figcaption></figure> <figure><img src="../../../.gitbook/assets/Screenshot 2025-10-23 at 14.12.18 (1).png" alt=""><figcaption></figcaption></figure></div>

{% hint style="info" %}
The recommendations in this section allow you to focus on the most impactful training opportunities. Use the filters to further customise the recommendations based on issue severity or for specific Organizations.
{% endhint %}
