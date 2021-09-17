# Protect your code with patches

The `protect` command applies the patches specified in your `.snyk file` to the local file system. This is currently supported for **Node.js projects only.**

Run `snyk protect` after you’ve created a `.snyk file` and installed your local dependencies \(by running `npm install` or `yarn install`\/). `snyk wizard` does this as a last step.

Since running `protect` is the way to repeatedly apply patches, you should run it every time you reinstall your modules. Common integration points would be your CI/build system, your deployment system, and adding it as a post installation step in your package.json file necessary if you consume this module via npm or yarn\).

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

