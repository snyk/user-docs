---
description: More easily fix issues with PR feature
---

# Create a Pull Request

When you use Snyk to open a fix PR, you automating the task of setting up and configuring the details to more quickly address the issue. The simplifications aim to minimize typos and help teams focus on fixing issues more efficiently. General details are available in the [Snyk Docs](https://docs.snyk.io/products/snyk-open-source/open-source-basics/fixing-vulnerabilities), and we describe details specific to to this workshop here.

First, you have to ensure you have setup and configured [atlassian-jira.md](../../../getting-started/atlassian-integrations/atlassian-jira.md "mention") with Snyk. We covered this requirement in a previous step within this module.

Let's work through the issues for a Dockerfile, because those are typically about updating the base image. In the Snyk projects view, expand you repository view to see the Docker and click on the details.

![](<../../../../.gitbook/assets/image (63).png>)

When you see the Dockerfile project, you will be presented with upgrade recommendations for the base image as shown below.

![](<../../../../.gitbook/assets/image (87) (2).png>)

Expand the list by clicking on the "Show more upgrade types" to see your available options.

![](<../../../../.gitbook/assets/image (65).png>)

Your options may vary from this screen because the Snyk database is frequently updated, and this includes the latest options for container images. You will see a few options presented.

* Minor upgrades are typically seen as low-impact as they contain fewer chances of breaking changes. Your team may be motivated to make a minor change because they feel it is necessary to stay within a version range for their application stack to maintain compatibility. In the example above, this means the minor change from 8.5.21 to 8.5.71.
* Major upgrades may be advantageous for your team if they address large numbers of vulnerabilities. In the example above, the difference between the current image and minor/major is approximately the same. In other examples, the gaps may be different and your team will use this information to make more informed choices.

For the purpose of this workshop, we'll pick the minor change for the **Open a fix PR** button. At this point, your development team will have an intuitive understanding of what they need to fix and how to fix it. They already know the change requires a few lines in a Dockerfile. What we'll show next is how the automation takes this routine operation of, "I know what to do" and automates the sequence. While simple in nature, this automation eliminates the common typo mistakes that too often come up in this process.

Click on that button now to start the process. The automated process starts by showing you a screen to confirm your choice:

![](<../../../../.gitbook/assets/image (62).png>)

Click on the Open a Fix PR button to confirm the change.

Next what happens is you are taken to Bitbucket Cloud. This is helpful because it is the shared environment for you and your team to manage your open PRs. The screen snippet below shows some of the PR details as populated by Snyk for your team to review as part of your PR process.

![](<../../../../.gitbook/assets/image (80).png>)

Below this information is the code diff your team should easily recognize. This is the fix your developers _knew_ they could make, and Snyk did it for them.

![](<../../../../.gitbook/assets/image (84).png>)

This example is purposefully simple to focus on a single file. The next iteration of this example is to create a pull request on multiple files. While not covered in this workshop, the idea is the same. You select multiple issues you want Snyk to automatically fix across one or more files. Developers frequently "know" what to fix and quickly see the power of being able to let the automated PR process do the right search-and-replace for them. This way, your development team spends more time on the review of the content than the creation of the change.
