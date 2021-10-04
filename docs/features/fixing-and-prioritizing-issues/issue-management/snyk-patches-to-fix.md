# Snyk patches to fix

## **Snyk patches to fix vulnerabilities**

### **Getting started with Snyk Open Source patches**

Sometimes there is no direct upgrade that can address the vulnerability or an upgrade is not possible due to functional reasons \(e.g. it’s a major breaking change\).

For such cases Snyk can help you [protect your code with patches](https://docs.snyk.io/snyk-cli/secure-your-projects-in-the-long-term/protect-your-code-with-patches). This option will make the minimal modifications to your locally installed nodemodules files to fix the vulnerability. It will also update the policy to patch this issue when running snyk protect_._

{% hint style="info" %}
**Caution**  
Patching is currently supported for **Node.js** projects only.
{% endhint %}

Patches are applicable in the following scenarios:

1. When there is no upgrade available for the direct dependency
2. When there is no way of upgrading a direct dependency to get to a vulnerability free version of a transitive dependency.
3. When an upgrade would render the package incompatible with the current codebase.

Patches are available via the source code integrations and the Snyk CLI.

### **How do patches work when using the Snyk CLI?**

If you use the Snyk Protect to fix your vulnerable Node.js project by applying a patch then following things will happen:

1. [`@snyk/protect`](https://www.npmjs.com/package/@snyk/protect) will be added as a production dependency of the project
2. A postinstall hook will be added to run `snyk-protect` when npm install or yarn install completes.

This means that whenever the project dependencies are installed with npm install or yarn install then the hook can trigger `snyk-protect` to run and patch the necessary dependencies, on completion you will see a success message in the output.

### **Difference between the `@snyk/protect` and `snyk protect`**

Previously, Snyk was adding the whole [Snyk CLI `snyk` package](https://www.npmjs.com/package/snyk) to your project dependencies in order to run `snyk protect` command. We’ve created a new standalone [package `@snyk/protect`](https://github.com/snyk/snyk/tree/master/packages/snyk-protect#snykprotect), that’s much lighter and faster for applying patches.

If you are still using a `snyk` package to apply patches, we recommend to migrate to `@snyk/protect` by either [following its README](https://github.com/snyk/snyk/tree/master/packages/snyk-protect#snykprotect) or running a [migration script](https://www.npmjs.com/package/@snyk/cli-protect-upgrade) with: `npx @snyk/cli-protect-upgrade`.

### **How do patches work when using the source code integrations?**

When you choose to use a patch to fix a vulnerability,`@snyk/protect` is added as a dependency and a `.snyk` file is created which contains the list of patches to apply.

The `.snyk` file contains the details of the patch per individual path to the dependency as it may appear in multiple locations in the `node_modules`, for example:

```text
# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.
version: v1.22.1
ignore: {}
# patches apply the minimum changes required to fix a vulnerability
patch:
  SNYK-JS-LODASH-567746:
    - tap > nyc > istanbul-lib-instrument > babel-types > lodash:
        patched: '2021-02-17T13:43:51.857Z'
```

However, only a single path is required, as `@snyk/protect` will attempt to patch all applicable instances of the vulnerable dependency.

