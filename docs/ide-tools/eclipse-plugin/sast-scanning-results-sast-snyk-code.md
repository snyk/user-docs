# SAST scanning results (SAST, Snyk Code)

In Eciplse plugin version 2.0.0 and later, Snyk is introducing a deeper integration with the native flows of Eclipse: inline highlights, problems integrations, and information about the issue on hover) The following shows all of these for a high severity security vulnerability found in a `js` file:

1. The security vulnerability is highlighted (the red squiggly line) indicating there is a high severity security vulnerability in your code. You see the vulnerability ID and what the issue is on hover. Snyk is updating the panel to include more information about the found Code issues, to provide more details like data flow and example fixes.
2. You see the integration with the **Problems** tab, which is useful if you use the **Problems** tab to show only issues in the current file. Snyk also shows the line where the issue is.
3. You can see the gutter icons on the left and file map highlights (with colors matching the priorities) on the right.

{% hint style="info" %}
Currently the hover information is limited to JavaEditor and GenericEditor, which is the default editor for plugins like Wild Web Developer.
{% endhint %}

![SAST scanning results](<../../.gitbook/assets/Screenshot 2022-05-13 at 12.56.46.png>)

