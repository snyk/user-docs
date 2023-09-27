# View your first Snyk Projects

## **Introduction:** see what Snyk scanned

We can now see your scan results.

{% hint style="info" %}
**Reminder: where am I?**\
In the Snyk Web UI, you see information specific for your **Organization** (such as your team), which is under a **Group** (such as your company). This allows your company to organize and collect data for the work your teams are doing. See [Managing users & permissions](../../snyk-admin/manage-users-in-organizations-and-groups/).
{% endhint %}

### View imports

In the Snyk Web UI, navigate to your **Projects** page, and see your imported repositories (or **targets** if importing non-code information). For example:

<figure><img src="../../.gitbook/assets/Target-list.png" alt=""><figcaption><p>List of imported targets</p></figcaption></figure>

For each entry, the left icon shows the number of Snyk Projects in each entry, plus the Git-based repository the projects are imported from. For example <img src="../../.gitbook/assets/image (187) (1).png" alt="" data-size="line">.

### Private / public repositories: the lock symbol

When [setting up your GitHub integration](../../integrations/git-repository-scm-integrations/snyk-github-integration.md), you can choose whether Snyk can access public and private repositories, or public repositories only:

<figure><img src="../../.gitbook/assets/image (405) (1).png" alt="Set whether Snyk can access private repos"><figcaption><p>Set whether Snyk can access private repos</p></figcaption></figure>

When you then import a Project, private repositories are identified with a “lock” symbol (<img src="../../.gitbook/assets/image (101) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="" data-size="line">) in the imported scan details:

<figure><img src="../../.gitbook/assets/image (125) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Private repos with lock symbol"><figcaption><p>Private repos with lock symbol</p></figcaption></figure>

For customers on free plans, private repository scans count towards your test count limit.

{% hint style="info" %}
Typically, team leads do the original integration setup and Project import, rather than individual developers.
{% endhint %}

### View lists of Projects

When you open an entry, you see the different Snyk Projects scanned in that entry.

{% hint style="info" %}
**Reminder: what is a Project?**\
A Snyk Project is an item scanned by Snyk; for example, a manifest file listing all your open-source libraries as dependencies. See [Snyk Projects](../../snyk-admin/introduction-to-snyk-projects/).
{% endhint %}

For example:

<figure><img src="../../.gitbook/assets/image (180) (1) (1) (1) (1) (1).png" alt="List of scanned Projects"><figcaption><p>List of scanned Projects</p></figcaption></figure>

### Understand Project information

#### Why are there several items here? What do they mean? Which should I use?

When we import Snyk Projects for the first time, there’s a lot of information - don’t worry!

When you write your application, you may write your own code, import Open Source libraries with dependencies, and build all of that all into a container for deployment.

Snyk scans different parts of this lifecycle, with different icons and entries showing the results for each of these parts of your work, including:

| Example                                                                                         | Description                                                                                                                                                                                                                  |
| ----------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <img src="../../.gitbook/assets/image (297) (1).png" alt="" data-size="line">                   | Your own code analysis results, scanned by[ Snyk Code](../../scan-application-code/snyk-code/).                                                                                                                              |
| <img src="../../.gitbook/assets/Screenshot 2022-07-20 at 11.14.02.png" alt="" data-size="line"> | Your open source libraries, scanned by [Snyk Open Source](../../scan-application-code/snyk-open-source/), displaying each detected manifest, such as **pom.xml**, **package.json**, and other manifests for these libraries. |
| <img src="../../.gitbook/assets/image (307).png" alt="" data-size="line">                       | Container results, scanned by [Snyk Container](../../scan-applications/snyk-container/), for items built into a container, such as a Docker file.                                                                            |
| <img src="../../.gitbook/assets/image (206) (1).png" alt="" data-size="original">               | Kubernetes deployment files, terraform and other IaC files, scanned by [Snyk Infrastructure as Code (IaC)](../../scan-infrastructure/snyk-infrastructure-as-code/).                                                          |

{% hint style="info" %}
Other files and types can be displayed; see [View project information](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects/view-project-information) for more details.
{% endhint %}

### Viewing Project settings

Snyk treats each item in this list as a separate **Project**.

This allows you to control settings for that Project, by clicking on the cog icon (![](<../../.gitbook/assets/image (144).png>)) to define how that Project is scanned:

<figure><img src="../../.gitbook/assets/image (208) (1) (1) (1) (1) (1) (1) (1).png" alt="Click cog icon to edit settings"><figcaption><p>Click cog icon to edit Project settings</p></figcaption></figure>

For example, you can change scan frequency, setting how often scans are run by default. See [View Project Settings](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects/view-project-settings) for more information.

### Scan results

Let’s go back to the results for your scan:

<figure><img src="../../.gitbook/assets/image (167) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Project scan results"><figcaption><p>Project scan results</p></figcaption></figure>

The scan shows you all vulnerabilities in all aspects of an application. Of course, it's unlikely that you are responsible for every entry in this list, but it's important to be aware of the full picture.

So if your Snyk Open Source scan shows no vulnerabilities in your open source libraries - great :tada:! But there may still be a lot of issues identified by other scans, such your container. And even if the developers in your team did not create or manage these issues, you should still know about them.

#### More information

Refer to the Snyk Training course, [Introduction to the Snyk UI](https://training.snyk.io/courses/introduction-to-the-snyk-ui), to learn more about reviewing results from open source, code, container, and infrastructure file scans.&#x20;



### What's next?

Now you understand what results you're seeing, you need to [understand the vulnerabilities](understand-your-vulnerabilities.md) themselves.
