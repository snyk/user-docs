# Create a Jira Ticket

When your team uses Jira tickets to plan and manage your workloads, the Snyk integration into Atlassian Jira is useful. The Snyk integration simplifies your workload by providing most of the information you need to create issues for your team. This is especially helpful when you need to work through processes to setup and configure issues before planning and scoping the work.

Be sure to follow the instructions outlined in [atlassian-jira.md](../../../atlassian-jira.md "mention") to setup and configure Jira.

Navigate to one of the vulnerabilities in your project and examine its contents. The Snyk database is frequently updated, and your results may reflect a different count or severity from the time this workshop was created. We'll pick the first vulnerability for review and click on the Create a Jira Issue button.

![](<../../../../../../.gitbook/assets/image (159) (1) (1).png>)

Clicking on the button takes you to a screen with most content already filled in. This makes it easier for you to quickly add an issue. The integration uses your permissions to look up Jira projects and users to assign tickets. This lookup depends on your permissions and Snyk will only present projects you are entitled to see. The next screenshot shows some of of the fields selected for the **Project**, **Issue Type**, **Reporter**, and **Assignee**.

![](<../../../../../../.gitbook/assets/image (107) (2) (1) (1) (1) (2).png>)

Once you create your ticket, the Jira Integration adds a link to the issue in-line to help your team more easily see the status of the vulnerability as being tracked in Jira.

![](<../../../../../../.gitbook/assets/image (79) (1) (1) (1).png>)

Now that you've created your Jira issue, your team can plan and prioritize the work per your own processes. The ticket preserves details from the original ticket, including important links to external databases and details about the code.
