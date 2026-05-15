# Third-party dependency scanning (SCA, Snyk Open Source)

In the Eclipse plugin version 2.0.0 and later, Snyk has enhanced integrations with the native flows of Eclipse: inline code highlights with displays of information about the issue on hover, and Eclipse Problems integrations. The following shows all of these for a high-severity security vulnerability found in a `js` file:

* The security vulnerability is highlighted in your code, with the underline color and icon adjacent to the line number indicating the severity of the issue. You can see the vulnerability ID and detailed information by hovering over the highlighted code. The hover information is limited to the JavaEditor and GenericEditor, the latter being the default editor for plugins like Wild Web Developer.
* Vulnerabilities are displayed in the **Problems** view, which allows for filtering and grouping issues. The line containing the issue is shown in the **Location** column. You can navigate to the issue in the code by clicking the issue in the list of problems.

<figure><img src="../../../.gitbook/assets/image (18).png" alt=""><figcaption><p>Snyk Open Source issue displayed in Eclipse</p></figcaption></figure>

In addition to this, the **Snyk Results** view offers detailed issue descriptions, including any available fixes. In this view, you can start and stop scans, filter issues, and more.&#x20;

To filter issues in the **Snyk** view, click the three vertical dots menu at the top right of the view, then select any combination of options from the **Severity**, **Snyk Product**, **Issues Status**, and **Fixability** submenus to customize your filter.

To ignore an issue for 30 days, right-click the issue in the left panel of the **Snyk** view and select **Ignore issue**.

You can enable monitoring for a Project to help spot new issues as they are reported. To do this, right-click a folder in the left panel of the **Snyk** view and select **Monitor project**.&#x20;

<figure><img src="../../../.gitbook/assets/image (235).png" alt=""><figcaption><p>Details for a Snyk Open Source issue dispilayed in the Snyk view</p></figcaption></figure>
