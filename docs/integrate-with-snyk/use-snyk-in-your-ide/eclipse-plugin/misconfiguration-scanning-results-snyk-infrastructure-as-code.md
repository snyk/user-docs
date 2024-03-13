# Misconfiguration scanning results (Snyk Infrastructure as Code)

In Eclipse plugin version 2.0.0 and later, Snyk is introducing a deeper integration with the native flows of Eclipse: inline highlights, problems integrations, and information about the issue on hover. The following shows all of these for a high severity misconfiguration found in a Terraform file:

1. The misconfiguration is highlighted (the red squiggly line) indicating there is a high severity security vulnerability in this file and the line number. You have all the information on hover; you can scroll, read, or click the links (when available) for even more information. Advice on how to resolve the misconfiguration is right there where the misconfiguration is.
2. You see the integration with the **Problems** view, which is useful if you use the **Problems** view to filter and group issues. Snyk also indicates the line where the issue is, and clicking the issue in the problem view navigates to it.
3. You can see the gutter icons on the left and file map highlights (with colors matching the priorities) on the right.

{% hint style="info" %}
Currently the hover information is limited to JavaEditor and GenericEditor, the latter being the default editor for plugins like Wild Web Developer.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1).png" alt=""><figcaption><p>Snyk IaC findings displayed in Eclipse</p></figcaption></figure>
