# Snyk runtime monitoring: install the Snyk agent for Node.js

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

##  Snyk runtime monitoring: install the Snyk agent for Node.js

Snyk supports Node.js v8 and above.

1. Add [@snyk/nodejs-runtime-agent](https://www.npmjs.com/package/@snyk/nodejs-runtime-agent) as a dependency to your project in order to start using Snyk Runtime Monitoring for your Node.js applications.
2. Ensure the Require statement for the agent is entered prior to all other require statements that you may add to the code.

   ```text
   //Example of Node.js code to run Snyk runtime agent
   require('@snyk/nodejs-runtime-agent')
   ({projectId: '0462e42b-c92f-4b48-bac8-81eb3ff7f43e',
   });
   ```

3. Commit and push the changes to your manifest file \(for example package.json\).

* 
