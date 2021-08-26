# Snyk runtime monitoring: an overview of the app interface

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

When Snyk runtime monitoring is successfully monitoring your projects, there are a few positive indications in the app:

* From the **Projects** tab, an animated indicator appears on the rows for all projects monitored at runtime:

 

![ProjectsPage\_Filter\_GitHub.gif](../../.gitbook/assets/uuid-520448e8-52c7-250a-be13-df4e4518560c-en.gif)

From within a project that is monitored at runtime:

* **Monitored at Runtime** appears at the top of the project page.
* **Runtime agent upgrade available**, indicates we've updated the Snyk agent and we recommend you upgrade.
* **Live runtime agents**, displays the number of application instances running and monitored by Snyk at runtime
* **Called At Runtime** indicates that vulnerable functions were recently invoked in a monitored application instance. This tag appears only when vulnerable functions are called for a monitored project.
* Per vulnerability, the vulnerable functions are listed and **Monitored** appears next to those functions being watched at runtime.
* An indicator also displays how long it has been since one of the vulnerable functions was last called.

 

![Runtime\_ProjectDetailsExample.png](../../.gitbook/assets/uuid-6e84ffa7-04bd-413e-4bbb-3bd5de1d9092-en.png)

