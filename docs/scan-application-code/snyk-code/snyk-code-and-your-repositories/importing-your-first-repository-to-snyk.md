# Importing your first repository to Snyk

If your Snyk account currently does not include any repositories, you must import to Snyk the repositories you want to be analyzed and tested by Snyk Code.

Follow these steps to import your first repository for Snyk Code testing:

1\. Go to [snyk.io](http://snyk.io) and click the **Log in** button.

2\. On the Snyk Web UI, select the **Projects** option.

{% hint style="info" %}
If your **Project** tab already contains Projects, follow the instructions on one of the following pages: [Importing additional repositories to Snyk](importing-additional-repositories-to-snyk.md) or [Re-importing existing repositories](re-importing-existing-repositories-for-snyk-code-testing.md).
{% endhint %}

3\. On the **Add your first project to scan for vulnerabilities** page in the **Suggested Repositories** section, select the checkboxes of the repositories you want to test with Snyk Code

{% hint style="info" %}
You can click the **Show all repositories** link to display all the repositories in your integrated SCM.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (2) (4) (1).png" alt="Select repositories and click Import and scan"><figcaption><p>Select repositories and click Import and scan</p></figcaption></figure>

4\. Click **Import and scan**.

The selected repositories are imported to Snyk Code, and a progress bar appears:

<figure><img src="../../../.gitbook/assets/image (10) (3) (1) (1).png" alt="Import progress bar"><figcaption><p>Import progress bar</p></figcaption></figure>

When the import is completed, a confirmation message appears on the Projects page. Your imported repositories appear as Target folders, each containing the Code analysis Project that includes the findings of the Snyk Code test:

<figure><img src="../../../.gitbook/assets/image (296) (1).png" alt="Target folder on Projects page"><figcaption><p>Target folder on Projects page</p></figcaption></figure>
