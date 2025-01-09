# SAST scanning results (SAST, Snyk Code)

In Eclipse plugin version 2.0.0 and later, Snyk is introducing a deeper integration with the native flows of Eclipse: inline highlights, problems integrations, and information about the issue on hover. The following shows all of these for a high-severity security vulnerability found in a `js` file:

1. The security vulnerability is highlighted indicating there is a high severity security vulnerability in your code. You see the vulnerability ID and what the issue is on hover.
2. You see the integration with the **Problems** view (bottom of screen) which is useful if you use the **Problems** view to filter and group issues. Snyk also indicates the line where the issue is. Click the issue in the problem view to navigate to the issue.
3. You can see the gutter icons on the left and file map highlights (with colors matching the priorities) on the right.
4.  In addition to this, the **Snyk** view offers detailed issue descriptions, including the dataflow and fix examples, together with possibilities to start and stop scans, filter issues using the toolbar of the view. and more.&#x20;

    <figure><img src="../../../.gitbook/assets/image (643).png" alt=""><figcaption><p>The filter menu opens when clicking on the 3 dots</p></figcaption></figure>

{% hint style="info" %}
The hover information is limited to JavaEditor and GenericEditor, the latter being the default editor for plugins like Wild Web Developer.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (122) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1).png" alt=""><figcaption><p>Snyk Code findings displayed in Eclipse</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (645).png" alt=""><figcaption><p>Snyk Code findings displayed in the Snyk View</p></figcaption></figure>
