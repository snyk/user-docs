# Snyk runtime monitoring: an overview of the app interface

####  [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

###  [ ]() <a id="category-name"></a>

#### Runtime monitoring

* [ Snyk runtime monitoring: introduction](/hc/en-us/articles/360003737297-Snyk-runtime-monitoring-introduction)
* [ Snyk runtime monitoring: an overview of the app interface](/hc/en-us/articles/360003699038-Snyk-runtime-monitoring-an-overview-of-the-app-interface)
* [ Snyk runtime monitoring: prerequisites for configuration](/hc/en-us/articles/360003737317-Snyk-runtime-monitoring-prerequisites-for-configuration)
* [ Snyk runtime monitoring: install the Snyk agent for Node.js](/hc/en-us/articles/360003699058-Snyk-runtime-monitoring-install-the-Snyk-agent-for-Node-js)
* [ Snyk runtime monitoring: upgrade the Snyk agent for Node.js](/hc/en-us/articles/360003737337-Snyk-runtime-monitoring-upgrade-the-Snyk-agent-for-Node-js)
* [ Snyk runtime monitoring: disable the Snyk agent](/hc/en-us/articles/360003699078-Snyk-runtime-monitoring-disable-the-Snyk-agent)
* [ Snyk runtime monitoring: uninstall the Snyk agent for Node.js](/hc/en-us/articles/360003737357-Snyk-runtime-monitoring-uninstall-the-Snyk-agent-for-Node-js)
* [ Snyk runtime: what data is used when monitoring my Node.js projects?](/hc/en-us/articles/360003699098-Snyk-runtime-what-data-is-used-when-monitoring-my-Node-js-projects-)
* [ Snyk runtime monitoring: install the Snyk agent for Java](/hc/en-us/articles/360003699118-Snyk-runtime-monitoring-install-the-Snyk-agent-for-Java)
* [ Snyk runtime monitoring: upgrade the Snyk agent for Java](/hc/en-us/articles/360003699138-Snyk-runtime-monitoring-upgrade-the-Snyk-agent-for-Java)

 [See more](/hc/en-us/sections/360001051838-Runtime-monitoring)

1.  [Docs Library \| Snyk](/hc/en-us)
2.  [Reports](/hc/en-us/categories/360000598418-Reports)
3.  [Runtime monitoring](/hc/en-us/sections/360001051838-Runtime-monitoring)

##  Snyk runtime monitoring: an overview of the app interface

When Snyk runtime monitoring is successfully monitoring your projects, there are a few positive indications in the app:

* From the **Projects** tab, an animated indicator appears on the rows for all projects monitored at runtime:

From within a project that is monitored at runtime:

* **Monitored at Runtime** appears at the top of the project page.
* **Runtime agent upgrade available**, indicates we've updated the Snyk agent and we recommend you upgrade.
* **Live runtime agents**, displays the number of application instances running and monitored by Snyk at runtime
* **Called At Runtime** indicates that vulnerable functions were recently invoked in a monitored application instance. This tag appears only when vulnerable functions are called for a monitored project.
* Per vulnerability, the vulnerable functions are listed and **Monitored** appears next to those functions being watched at runtime.
* An indicator also displays how long it has been since one of the vulnerable functions was last called.
* 
