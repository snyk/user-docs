# Third-party dependency scanning (SCA, Snyk Open Source)

In Eclipse plugin version 2.0.0 and later, Snyk is introducing a deeper integration with the native flows of Eclipse, inline highlights, problems integrations, and information about the issue on hover. The following shows all of these for a security vulnerability found in a third party dependency:

1. The vulnerable package is highlighted (the red squiggly line) indicating there is a high severity security vulnerability in this package. You have all the information on hover; you can scroll, read, or click the links for even more information. Advice on what action to take and how is presented right where the vulnerability is.
2. You see the integration in the **Problems** view, where you can filter and group issues. Snyk also indicates the line where the issue is, and you can click the issue to navigate to it.
3. You can see the gutter icons on the left, as well as the file map highlights (with colors matching the priorities) on the right.
4. In addition to this, the **Snyk view** offers detailed issue descriptions, including the dataflow and fix examples together with possibilities to filter issues using the toolbar of the view.

{% hint style="info" %}
The hover information is limited to JavaEditor and GenericEditor, which is the default editor for plugins like Wild Web Developer.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (267) (1) (1) (1).png" alt=""><figcaption><p>Snyk Eclipse displaying a Snyk Open Source issue</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (649).png" alt=""><figcaption><p>Snyk View displaying issue details for a Snyk Open Source issue</p></figcaption></figure>

In the Snyk View, in the panel on the left, right-click the issue to use the context menu options, I**gnore issue in Snyk** and **Monitor project**.

<figure><img src="../../../.gitbook/assets/image (648).png" alt=""><figcaption><p>Context menu for Snyk Open Source issue</p></figcaption></figure>

**Ignore issue**—Hover over the specific issue that you want to ignore for the next 30 days and then select the context options you want to use.

**Snyk monitor** —Run `snyk test` for the entire workspace.
