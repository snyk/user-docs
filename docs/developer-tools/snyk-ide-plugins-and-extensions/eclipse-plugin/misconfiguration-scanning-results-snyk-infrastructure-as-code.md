# Misconfiguration scanning results (Snyk Infrastructure as Code)

In the Eclipse plugin version 2.0.0 and later, Snyk has enhanced integrations with the native flows of Eclipse: inline code highlights with displays of information about the issue on hover, and Eclipse Problems integrations. The following illustrates all of these for a high-severity security vulnerability found in a `js` file:

* The misconfiguration is highlighted in your code, with the underline color and icon adjacent to the line number indicating the severity of the issue. You can see detailed information and advice on how to resolve the misconfiguration by hovering over the highlighted code. The hover information is limited to the JavaEditor and GenericEditor, the latter being the default editor for plugins like Wild Web Developer.
* Misconfigurations are displayed in the **Problems** view, which allows for filtering and grouping of issues. The line containing the issue is shown in the **Location** column. You can navigate to the issue in the code by clicking the issue in the list of problems.

<figure><img src="../../../.gitbook/assets/image (95).png" alt=""><figcaption><p>Snyk IaC findings displayed in Eclipse</p></figcaption></figure>

In addition to this, the **Snyk** view offers detailed issue descriptions, including **Remediation** suggestions. In this view, you can start and stop scans, filter issues, and more.&#x20;

To filter issues in the **Snyk** view, click the three vertical dots menu at the top right of the view, and then select any combination of options from the **Severity**, **Snyk Product**, **Issues Status**, and **Fixability** submenus to customize your filter.

<figure><img src="../../../.gitbook/assets/image (325).png" alt=""><figcaption><p>Snyk IaC findings displayed in Snyk View</p></figcaption></figure>
