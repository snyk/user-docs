# SAST scanning results (SAST, Snyk Code)

## SAST scanning results (SAST, Snyk Code)

Starting version 2.0.0 and above, Snyk is introducing a deeper integration within Eclipse with the native flows of Eclipse (inline highlights, problems integrations, information about the issue on hover). The following shows all of these for a high severity security vulnerability found in a `js` file:

1. The security vulnerability is highlighted (the red squiggly line) indicating there is a high severity security vulnerability in your code. You will see the vulnerability ID, what the issue is, on hover. Snyk is updating the panel to include more information about the found Code issues, so you have even more details like data flow and example fixes.
2. You see the integration with the problems tab, which comes in handy if you use the problems tab to show only issues in the current file. Snyk also shows the line where the issue is.
3. You can see the gutter icons on the left and file map highlights (with colors matching the priorities) on the right.

{% hint style="info" %}
Currently the hover information is limited to JavaEditor and GenericEditor, which is the default editor for plugins like Wild Web Developer, etc.
{% endhint %}

![SAST scanning results](<../../.gitbook/assets/Screenshot 2022-05-13 at 12.56.46.png>)

