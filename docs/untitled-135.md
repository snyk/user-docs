# Snyk patches to fix

* [ Evaluating and prioritizing vulnerabilities](/hc/en-us/articles/360006113978-Evaluating-and-prioritizing-vulnerabilities)
* [ Remediate your vulnerabilities](/hc/en-us/articles/360006113798-Remediate-your-vulnerabilities)
* [ Upgrading package versions to fix](/hc/en-us/articles/360005993658-Upgrading-package-versions-to-fix)
* [ Snyk patches to fix](/hc/en-us/articles/360004032437-Snyk-patches-to-fix)
* [ Ignoring issues not prioritized for your project](/hc/en-us/articles/360004002718-Ignoring-issues-not-prioritized-for-your-project)

##  Snyk patches to fix

###  How do patches work when using the Snyk CLI?

If you use the Snyk CLI to fix your vulnerable node project by running `snyk wizard` and choose to apply a patch then two things happen:

1. Adds Snyk as a dependency of the project \(so that the CLI is fetched with `npm install` \)
2. Adds a `postinstall` hook to run `snyk protect` when `npm install` runs

This means that whenever the project is built with `npm install`, all dependencies are downloaded from their source and placed in the `node_modules` folder, and then `snyk protect` runs to patch the necessary dependencies. If the Heroku buildpack invokes `npm install`, the relevant dependencies are patched. This is easily inspected in the buildpack output; look for Successfully applied Snyk patches.

Since running protect is the way to repeatedly apply patches, Snyk protect needs to be run every time you reinstall your modules. Common integration points would be in your CI/build system, or deployment system, and adding it as a post-installation step in the `package.json` file \(which is necessary if you consume this module via npm or yarn\).

###  How do patches work when using the source code integrations?

When you choose to use a patch to fix a vulnerability, Snyk is added as a dependency, a .snyk file is created which contains the list of patches to apply and the Snyk protect command post-install hook is added, and it is this command that is responsible for applying the patches.

The .snyk file contains the details of the patch for example

```text
'npm:negotiator:20160616':
- errorhandler &gt; accepts &gt; negotiator:
patched: '2017-05-05T12:39:16.961Z'
```

The patch is essentially an instruction stating which bits of code in the dependency need to be replaced and the code that should be used to replace it.

The Snyk protect command is what replaces the vulnerable code with the patch.

