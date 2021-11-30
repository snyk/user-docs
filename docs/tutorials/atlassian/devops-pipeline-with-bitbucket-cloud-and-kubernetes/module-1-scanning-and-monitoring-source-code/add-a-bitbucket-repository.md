---
description: Get the code into your Bitbucket
---

# Add a Bitbucket Repository

### Importing the repository

We'll start by importing a repository into your Bitbucket environment.  This establishes a common starting location for all workshop users.  The reference repository contains an intentionally vulnerable application modeled after the popular line of Snyk sample applications named **Goof.**  In our workshop, the implementation is in Java and the repository is aptly named _java-goof._  The _java-goof_ repository is also parameterized to permit you to more easily add it to your running environment.

The source repository is publicly available at this location:

{% embed url="https://github.com/snyk-labs/java-goof" %}

In  Bitbucket, navigate to your workspace and import the repository so you can freely make changes by clicking on the "Create Repository" button in your Repositories view.

![](<../../../../.gitbook/assets/image (93).png>)

Next, you'll be presented with a screen to select a workspace and other details.  In this screen, click on the **Import repository** link.

![](<../../../../.gitbook/assets/image (91) (1).png>)

The new screen asks for the URL to your Old Repository.  Enter [https://github.com/snyk-labs/java-goof](https://github.com/snyk-labs/java-goof) for that field, and select the Workspace, and Project.  The external repository does not require authentication, and your access level as Private or Public is your choice.  If desired, rename the repository.

![](<../../../../.gitbook/assets/image (67).png>)

Click **Import repository** to complete the process.

The next steps are to enable the integration between Snyk and Bitbucket.
