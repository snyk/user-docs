# SAST scanning results (SAST, Snyk Code)

In Eclipse plugin version 2.0.0 and later, Snyk is introducing a deeper integration with the native flows of Eclipse: inline highlights, problems integrations, and information about the issue on hover) The following shows all of these for a high severity security vulnerability found in a `js` file:

1. The security vulnerability is highlighted indicating there is a high severity security vulnerability in your code. You see the vulnerability ID and what the issue is on hover.
2. You see the integration with the **Problems** view, (bottom of screen) which is useful if you use the **Problems** view to filter and group issues. Snyk also indicates the line where the issue is, and clicking the issue in the problem view navigates to it.
3. You can see the gutter icons on the left and file map highlights (with colors matching the priorities) on the right.

{% hint style="info" %}
Currently the hover information is limited to JavaEditor and GenericEditor, the latter being the default editor for plugins like Wild Web Developer.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (122) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1).png" alt=""><figcaption><p>Snyk Code findings displayed in Eclipse</p></figcaption></figure>
