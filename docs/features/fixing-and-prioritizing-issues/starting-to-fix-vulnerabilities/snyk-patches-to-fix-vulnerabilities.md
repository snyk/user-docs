# Snyk patches to fix vulnerabilities

### Getting started with Snyk Open Source: patches

Sometimes there is no direct upgrade that can address the vulnerability or an upgrade is not possible due to functional reasons (e.g. it’s a major breaking change).

For such cases Snyk can help you [protect your code with patches](https://docs.snyk.io/snyk-cli/secure-your-projects-in-the-long-term/protect-your-code-with-patches). This option will make the minimal modifications to your locally installed`node_modules` files to fix the vulnerability. It will also update the policy to patch this issue when running [snyk protect](https://snyk.io/docs/using-snyk#protect).

{% hint style="info" %}
**Caution**\
Patching is currently supported for **Node.js** projects only.
{% endhint %}

Patches are applicable in the following scenarios:

1. When there is no upgrade available for the direct dependency
2. When there is no way of upgrading a direct dependency to get to a vulnerability free version of a transitive dependency.
3. When an upgrade would render the package incompatible with the current codebase.

Patches are available via the source code integrations and the [Snyk CLI](https://support.snyk.io/hc/en-us/articles/360003812578-Our-full-CLI-reference).

## Getting started with Snyk Open Source: the process for creating patches

Patches are created and maintained by Snyk. If the package owner has made code changes to fix the issues, our patch is based on this official fix, and we remove any cosmetic or unrelated changes. If a package owner has not addressed the vulnerability yet, we write a patch from scratch.

Before releasing it, we verify the patch, backport it to older versions, and test that the patch hasn’t broken functionality.

The patches are a part of [Snyk’s open source vulnerability database](https://github.com/Snyk/vulndb/), so you can check them out before applying them. For example, the patches for the [ms ReDoS vulnerability](https://github.com/Snyk/vulndb/tree/master/data/npm/ms/20151024). We don’t have patches for every case - if you need one that’s missing, let us know.

## How and when are Snyk patches created?

Snyk creates patches for high impact vulnerabilities, a vulnerability is determined as high impact if it is a serious vulnerability in a popular package that affects many users.

The Snyk Security team creates the patch usually by backporting a fix that has been added to the dependency. Backporting is the action of taking a fix that was built for a particular version of a piece of software, and applying it to a previous version of that software, by updating it to be functionally identical but with the fix for the vulnerability applied. For more information take a look at Redhat’s description of [Backporting Security Fixes](https://access.redhat.com/security/updates/backporting).

Once the patch is created by a Snyk Security Engineer it is reviewed by 2 other members of the team. The patch is also tested in the following ways

1. The package is built and tested using the packages automated tests
2. Packages or applications that use that patched package are tested using their automated tests.
3. Custom tests are created and run by the Snyk Security team

All patches are available for download and review by the community and we welcome any feedback.

For unmaintained packages, we will create a patch and open a pull request to the project for it to be merged.

## How do patches work when using the Snyk CLI?

The `protect` command has been replaced by `@snyk/protect`: https://github.com/snyk/snyk/tree/master/packages/snyk-protect; [npm package for `snyk-protect` command](https://www.npmjs.com/package/@snyk/protect). This page has instructions for using the package and migrating from `snyk protect`.

## How do patches work when using the source code integrations?

When you choose to use a patch to fix a vulnerability, `snyk` is added as a dependency and a `.snyk` file is created which contains the list of patches to apply.

The `.snyk` file contains the details of the patch per individual path to the dependency as it may appear in multiple locations in the `node_modules`, for example:

```
'npm:negotiator:20160616':
    - errorhandler > accepts > negotiator:
       patched: '2017-05-05T12:39:16.961Z'
    - negotiator: 
       patched: '2017-05-05T12:39:16.961Z'
```
