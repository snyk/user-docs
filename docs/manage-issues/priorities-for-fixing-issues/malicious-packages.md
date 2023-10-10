# Malicious packages

Malicious packages are a popular and growing method of carrying out software supply chain attacks. This page explains what malicious packages are, how Snyk identifies them, and what customers should do if they have malicious packages in their Projects.

{% hint style="warning" %}
Currently, Snyk does not consider the provenance or origin of a scanned package. In some cases, Snyk may detect false positives when the ecosystem, name, and version of a scanned package match a malicious public package.

Ensure you [verify the provenance of your packages](malicious-packages.md#verifying-the-provenance-of-packages).
{% endhint %}

## Introduction to malicious packages

### Open-source malicious packages

Open-source malicious packages contain malicious code deliberately inserted to perform attacks. This includes infecting the target network, stealing and exfiltrating sensitive information such as passwords and credit card information, and engaging in additional malicious activity done by downloaded malware components.

Malicious packages are published as legitimate packages on popular package registries including npm, PyPI, NuGet, RubyGems, and more.

### Malicious package curation process

The Snyk security research team looks into dozens of data sources daily and triages and verifies vulnerabilities, including malicious packages.&#x20;

This process is continually improved, adding new tools and sources to identify potential threats in a timely manner.

### No CVE ID for malicious packages

Usually, open-source vulnerabilities are assigned a CVE ID, for example, [CVE-2021-44228](https://security.snyk.io/vuln/?search=CVE-2021-44228). The CVE ID helps security vendors like Snyk identify, define, and catalog publicly disclosed vulnerabilities. In almost all cases, malicious packages are not assigned a CVE ID. The absence of this identifier makes sharing new threats in the security community difficult. Snyk uses the name of the package, the package registry, for example, npm, and the affected version range to publicly catalog and share the malicious packages.&#x20;

### Malicious packages - CWE-506&#x20;

All malicious packages in the [Snyk Vulnerability Database](https://security.snyk.io/) are labeled with [CWE-506 - Embedded Malicious Code](https://cwe.mitre.org/data/definitions/506.html).

By filtering on CWE-506 customers can quickly see if they have malicious packages among their Projects.

### Malicious packages severity - Critical

Malicious packages can allow attackers to run remote code on the target machine. By their nature, these packages are untrusted and can be updated to add malicious functionality at any time after their discovery. Therefore, Snyk evaluates malicious code and assigns it (in almost all cases) a Critical severity level.

The remediation advice for malicious packages is to not install them or to remove them if already been installed.

### Malicious package publication timeline

Malicious packages may be publicly available for only a few minutes until identified and taken down. In other cases, a malicious package can be included in a targeted supply chain infection campaign and go unnoticed for months or years.&#x20;

While research and defense capabilities improve, attackers use novel techniques and technologies to overcome their findings. Users may notice that a Snyk advisory was published recently, despite the malicious package's being in the wild for a long time.

## Types of malicious packages

There are many types of malicious packages, and sophisticated hackers always invent new ones. Looking at the recent supply chain campaigns, we can observe the following cluster of common attack vectors:

* [Typosquatting attack](malicious-packages.md#typosquatting-attack)
* [Dependency confusion attack](malicious-packages.md#dependency-confusion-attack)
* [Dependency hijacking](malicious-packages.md#dependency-hijacking)
* [Compromised accounts](malicious-packages.md#compromised-accounts)

### Typosquatting attack

Bad actors publish malicious packages to a registry, hoping to trick users into installing them. An example of a [typosquatting attack](https://snyk.io/blog/typosquatting-attacks/) is [crossenv](https://security.snyk.io/package/npm/crossenv). The attacker used a name similar to the popular package, [cross-env](https://security.snyk.io/package/npm/cross-env), and even wrapped the exact same functionality as the original module to convey that the module is indeed working as expected. However, the module captured environment variables and sent them to an attacker-controlled remote server.

### Dependency confusion attack

A [dependency confusion attack](https://snyk.io/blog/detect-prevent-dependency-confusion-attacks-npm-supply-chain-security/) is a tactic used by attackers who upload malicious packages with private package names to public package registries. The result of this is confusion of automatic processes that import packages, for example, IDE and CIs, so these processes download the public package instead of the intended private one.

For example, ACME Corporation developed an internal UI package and named it `acme-ui`, and hosted it in a private package registry.&#x20;

A malicious user published a malicious package with the same name to npm. A developer at ACME did not configure the environment to pull packages from the private registry. While installing packages, the developer mistakenly downloaded the `acme-ui` package from the public registry instead of the private one, allowing the malicious user to run code on the developer's machine, a successful instance of dependency confusion.&#x20;

### Dependency hijacking

In this attack, the package used by the developer is a legitimate public package, which downstream is using the malicious package as a dependency, for example, due to dependency confusion.&#x20;

More details about this type of attack are described in this article about [malicious code found in the npm package event-stream](https://snyk.io/blog/a-post-mortem-of-the-malicious-event-stream-backdoor/).&#x20;

### Compromised accounts

A bad actor can compromise the account of a public package maintainer and insert malicious code into a package.&#x20;

More details about this type of attack are described in the [malicious versions of ESLint packages](https://eslint.org/blog/2018/07/postmortem-for-malicious-package-publishes/) article.

## Verifying the provenance of packages

At present, Snyk Open Source and Snyk Container scanners are unable to differentiate between internal and external packages, meaning that they cannot determine if a package has been imported from a public or a private registry. This capability will be added in the future. Meanwhile, you might see inaccurate alerts related to packages imported from a private registry.

A Snyk alert is a warning that an attack has targeted the organization, not an implication that the organization published the malware.

Several ecosystems have started implementing support for verifying attestations. The public npm registry, for instance, provides support for package provenance. The more that maintainers adopt this feature and publish packages with provenance details, the more secure and better the ways to validate package provenance become. Snyk highly recommends you start planning readiness for this in your open-source package consumption strategies.

An internal package may be flagged as malicious. If this happens, it means that a package with a similar name was published on a public registry, most likely with malicious code. This publication was intended to make developers in your company import the malicious public package, either by automatic or manual processes, instead of importing the legitimate private package.

Avoid malicious packages by always [verifying the npm registry source](malicious-packages.md#verifying-the-npm-registry-source).

## Understanding malicious packages in npm

Due to the popularity of JavaScript and npm, most malicious packages target this ecosystem.&#x20;

The following nuances are important to understand:

* ["Security holding" in npm](malicious-packages.md#security-holding-in-npm)
* [Malicious package is not available on npm and without "security holding"](malicious-packages.md#malicious-package-is-not-available-on-npm-and-without-security-holding)

### **“Security holding” in npm**&#x20;

A “security holding” on a package means that the package contained malicious code and was removed from the registry by the npm Security Team. A security placeholder was published to ensure users are not affected in the future. Although the package is currently under a “security holding” placeholder, it is important to verify that it was not imported from npm before it was marked in this placeholder. After you confirm that all instances of this package were downloaded from a private registry, you can ignore the issue.

The [Snyk Vulnerability Database](../../scan-applications/snyk-open-source/starting-to-fix-vulnerabilities/using-the-snyk-vulnerability-database.md) also shows packages in a “Security Holding” state, for example, the [flatmap-stream](https://security.snyk.io/package/npm/flatmap-stream) package.

### **Malicious package is not available on npm and without “security holding”**

The package was available publicly in npm, possibly only briefly, before being removed, [unpublished](https://docs.npmjs.com/policies/unpublish), by the owner. Under certain conditions, this can result in the package page being unavailable, though it is possible to see that it was once in use. The name is still available to be used again, with no guarantee that its contents in the new incarnation will be safe.

In these cases, it is important to be extra cautious about where the package was downloaded. Ignoring the issue altogether may result in a future problem in case a malicious actor republishes the package to npm.

## Verifying the npm registry source

To check if you are using a public registry or a private one, you can use the following options:

* [Run CLI commands](malicious-packages.md#run-cli-commands)
* [Check the package-lock.json file](malicious-packages.md#check-the-package-lock.json-file)
* [Check .npmrc file](malicious-packages.md#check-the-.npmrc-file)
* [Use private packages with npm](malicious-packages.md#using-private-packages-with-npm)

### Run CLI commands

Run the following command:  `npm config get registry`

If the result is `registry.npmjs.org`, you are using a public registry.

### Check the package-lock.json file

The `package-lock.json` file contains detailed information about the packages and their sources in your Project. You can open this file and inspect the `resolved` field for each dependency to see the origin of each dependency. Packages from the public npm registry have URLs starting with `https://registry.npmjs.org/`.

### Check the .npmrc file

The `.npmrc` file in your Project's root directory can specify the registry where npm should fetch packages. Check if your `.npmrc` file is configured to use your private npm registry.&#x20;

This is how the file path should look if you are using a private registry:

`registry=https://your-private-registry-url/`&#x20;

If you do not have an `.npmrc` file or it does not specify a registry, npm uses the default public registry, `https://registry.npmjs.org/`.

### Using private packages with npm

If you use npm, you can host your private packages on the npm registry.

If your company is using this [service](https://docs.npmjs.com/about-private-packages), then the public registry will be https://registry.npmjs.org.&#x20;

In this case, you will need to verify that the packages used are private, by logging in to the npm website.&#x20;

## Remediation of malicious package findings

If you find evidence that a malicious package was installed in your environment, you should do the following:

* &#x20;Immediately remove it from the local drive, both the local folder `node_modules` and global package manager cache.&#x20;
* Remove it from the package registry proxy cache and database if it exists.
* Remove it from the package lockfile, `package-lock.json` for npm and `yarn.lock` for Yarn.&#x20;

You can remediate specific cases of malicious packages by implementing the following tactics:

* **Typosquatting**:  Remove the malicious package and switch to the correct safe package.
* **Dependency confusion:** If you imported the public package, either by accident or by default in the CI, before it was placed in "security holding," be sure to remove it. Ensure your development environment and the CI pipeline are configured to use the private registry and install the same-name internal package instead.&#x20;
* **Dependency hijacking** and **Compromised account**: A new safe version is typically released after identifying the malicious package. To resolve this type of attack, update the package to a new version.

It is important to assume that your environment has been infected and to conduct internal security drills. After removing the malicious package, be sure to check for any remnants of the malicious code.
