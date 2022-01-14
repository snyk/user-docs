---
description: Time to Snyk!
---

# Review Application Scan Results

In the previous step, adding your java-goof repository to Snyk automatically starts a repository scan operation. This process normally takes a minute or two and the results are available with Bitbucket and Snyk. We'll review those results in this section.

## Background

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-opensource-01.png)

In your Bitbucket repository, clicking on your **Snyk** link takes you to a page where you can review results from Snyk within Bitbucket as shown below:

![](<../../../../.gitbook/assets/image (73).png>)

Snyk automatically scans your application for vulnerabilities and presents results within Bitbucket. This summary starts with a count of all vulnerabilities for the major elements, or projects, in your repository. For the java-goof application, you will see multiple maven and Dockerfile projects in a logical structure.

If you click into the top-level Dockerfile link, you will a see a screen similar to what is shown below:

![](<../../../../.gitbook/assets/image (81) (1).png>)

The summary line is replicated for the Dockerfile to maintain context.

![](<../../../../.gitbook/assets/image (83) (1) (1) (1).png>)

This summary maintains context from the main page and you'll notice details such as the filename, total counts, and breakdown according to severity (Critical, High, Medium, Low). The reference project has hundreds of vulnerabilities, and Snyk presents the results in an order based on several factors that include the severity, if there are known exploits, and if there is a fix available. This ordered list helps your team focus on the vulnerabilities with the highest scores first.

Each vulnerability also contains links to public databases for users that wish to learn additional background and context.

![](<../../../../.gitbook/assets/image (82).png>)

Another section of the page focuses on the version of your software, available upgrades, and whether there is a known exploit. This level of detail helps teams more quickly assess a vulnerability as they can see the recommended fix when it is available.

![](<../../../../.gitbook/assets/image (85).png>)

You are encouraged to look at the results of other projects and take note of the remediation guidance provided. Some vulnerabilities have fixes, while others do not. Some vulnerabilities do not have known exploits, and that could influence how your team addresses them as issues. The varied nature of vulnerabilities affect how you and your team address them, and reviewing the results are an important part of the process.

Next, let's click into the top-level "visit Snyk" link visible on any project page to take you directly into Snyk. You visit Snyk to find even greater levels of details, filtering, and options beyond the results presented in Bitbucket. You use these pages to get fine-grain visibility and filtering of your vulnerabilities, plus several options to help you mitigate them.

Some of the details you see are shown in the next two images, and they include:

* **Open a fix PR** to help you initiate a Pull Request with the fixes automatically applied by Snyk. This includes guidance so your team can more quickly work through your PR process.
*   [Priority Score](https://snyk.io/blog/snyks-developer-first-prioritization-capabilities/) to help you effectively prioritize fixes.

    The score, ranging from 1-1000, is powered by a proprietary algorithm that processes a wide array of factors, such as [CVSS](https://www.first.org/cvss/) score, the availability of a

    fix known exploits, how new the vulnerability is, and whether it is reachable or not.
* When enabled, the [Jira integration](https://snyk.io/blog/jira-integration/) lets you add a Jira ticket.

![](<../../../../.gitbook/assets/image (86) (1).png>)

![](<../../../../.gitbook/assets/image (66).png>)

In the next two sections, we'll create a Jira ticket and a pull request.
