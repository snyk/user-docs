# Snyk runtime monitoring: install the Snyk agent for Node.js

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

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

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

