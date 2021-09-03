# Snyk runtime monitoring: uninstall the Snyk agent for Node.js

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

1. From the project CLI, remove the agent as a dependency to your project by running the following command

   `npm uninstall @snyk/nodejs-runtime-agent`

2. Remove the `Require` code snippet that you inserted at installation \(and upgrade\).
3. Commit and push the changes to your manifest file \(for example `package.json`\).

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}