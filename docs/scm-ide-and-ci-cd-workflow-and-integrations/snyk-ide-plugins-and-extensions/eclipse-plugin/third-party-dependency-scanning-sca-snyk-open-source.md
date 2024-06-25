# Third-party dependency scanning (SCA, Snyk Open Source)

In Eclipse plugin version 2.0.0 and later, Snyk is introducing a deeper integration with the native flows of Eclipse, inline highlights, problems integrations, and information about the issue on hover. The following shows all of these for a security vulnerability found in a third party dependency:

1. The vulnerable package is highlighted (the red squiggly line) indicating there is a high severity security vulnerability in this package. You have all the information on hover; you can scroll, read, or click the links for even more information. Advice on what action to take and how is presented right where the vulnerability is.
2. You see the integration with the **Problems** view, which is useful if you use the **Problems** view to filter and group issues. Snyk also indicates the line where the issue is, and clicking the issue in the problem view navigates to it.
3. You can see the gutter icons on the left, as well as the file map highlights (with colors matching the priorities) on the right.

{% hint style="info" %}
The hover information is limited to JavaEditor and GenericEditor, which is the default editor for plugins like Wild Web Developer.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (267) (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

## **Context menu**

Right click menu options include:

**Ignore issue**—Hover over the specific issue that you want to ignore for the next 30 days and then access the context menu.

**Snyk test**—Run the Snyk test for the entire workspace.

**Preferences**—Access and update Snyk Vuln Scanner preferences directly from the right click menu.

## **Snyk View when collapsed**

**Title:** The name of the project.

**Dependency:** A summary of vulnerabilities and the number of affected paths found for each project.

## Snyk View when expanded

**Title:** The full name of the vulnerability affecting your project, linked to a description and complete details of the vulnerability in the Snyk database, to assist you in resolving the issue.

**Dependency:** The name of the direct dependency package in your project (the package you explicitly installed) that is affected by the vulnerability, either directly or indirectly.

All details appear on a single row and the Dependency (the name of the package explicitly used in the code) and Package (the name of the package that actually contains the vulnerability) columns both display the name of the same package:

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-4cdd086d6be47b598fc1a9a52c63023d59cff825%2Fuuid-e7accdc1-7495-e7a5-7a64-2403b066cb03-en.png?alt=media&token=e3bf024a-ba92-4b76-87be-b728d7edf092" %}
Eclipse results details
{% endembed %}

An arrow appears on the row, grouping together all relevant details, similar to the following examples:

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-85e429be9a965c2dc534817a648773176a724531%2Fuuid-c71f67d1-80a3-7485-b33b-e602a1a5050e-en.png?alt=media&token=99e95293-bb37-4fed-8388-d9cb56a73092" %}
Eclipse results arrow on row grouping details
{% endembed %}

## Examples in collapsed and expanded mode

The following shows a **dependency in collapsed mode**, when your project is affected by an indirect vulnerability:

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-85e429be9a965c2dc534817a648773176a724531%2Fuuid-c71f67d1-80a3-7485-b33b-e602a1a5050e-en.png?alt=media&token=99e95293-bb37-4fed-8388-d9cb56a73092" %}
Collapsed mode, indirect vulnerability
{% endembed %}

In this example:

Package X uses Package Y, which in turn uses Package Z.

Package Z contains a Cross-Site Scripting (XSS) vulnerability, indirectly affecting your project.

The Dependency (the name of the package explicitly used in the code) is Package X; the Package field displays Package Z (the name of the package that actually contains the vulnerability).

The following shows a **dependency in expanded mode**, when your project is affected by an indirect vulnerability:

Click the arrow on the row **to expand and view** the full path from the direct dependency to the vulnerable package.

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-992b169b89e7f3c45782fdeb47b205e3c0a95af8%2Fuuid-35658aaf-3359-80c2-c094-41a34c7863cc-en.png?alt=media&token=53c91ccc-f9bc-4ba7-a55f-8def3aa50d86" %}
Expanded mode, indierct vulnerability
{% endembed %}

On the preceding screen the full path would appear as:

\[Name of Package X]-->\[Name of Package Y]-->\[Name of Package Z]

**Package:** The name of the package in your project that is directly affected by the vulnerability. On the preceding screen:

* The Dependency is indicated as Package X—this is the package the developer explicitly uses in the code
* the Package field displays Package Z, which is the package that contains the vulnerability.

**Fix:** The name of the package if any and the version that it can be upgraded to in order to resolve the issue.

## My Snyk Results panel is not visible

If you close the Snyk Results panel by accident, or for some reason you do not see it, you can enable it as follows:

Navigate to **Windows -> Show View -> Other...**

![Show View, Other](<../../../.gitbook/assets/Screenshot 2022-05-13 at 12.04.07.png>)

Search for Snyk in the **Show View** dialog window.

![Show View dialog window](<../../../.gitbook/assets/Screenshot 2022-05-13 at 12.02.06 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (4).png>)

You should now be able to see the Snyk Results panel:

![Snyk Results panel](<../../../.gitbook/assets/Screenshot 2022-05-13 at 12.02.18.png>)

##
