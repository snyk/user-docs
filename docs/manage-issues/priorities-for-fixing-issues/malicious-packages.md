# Malicious packages

Malicious packages are a popular and growing method of carrying out software supply chain attacks. This page explains what malicious packages are, how Snyk identifies them, and what customers should do if they have malicious packages in their Projects.

{% hint style="warning" %}
Currently, Snyk does not consider the provenance or origin of a scanned package. In some cases, Snyk may detect false positives when the ecosystem, name, and version of a scanned package match a malicious public package.

[Learn more](malicious-packages.md#verifying-the-provenance-of-packages) how to verify the source of the package.
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

### Malicious package publication timeline

Malicious packages may be publicly available in the wild for only a few minutes until identified and taken down, and in other cases, they can be part of sophisticated company-specific targeted supply chain infection campaigns and stay under the radar for months or years.&#x20;

While research and defense capabilities improve, attackers use novel techniques and technologies to overcome their findings. Thus, users may see that a Snyk advisory has been published only recently, while the malicious package was initially in the wild for a long time.&#x20;

## Types of malicious packages

There are many types of malicious packages, and sophisticated hackers always invent new ones. Looking at the recent supply chain campaigns, we can observe the following cluster of common attack vectors:

### Typosquatting attack

Bad actors publish malicious packages to a registry, hoping to trick users into installing them. An example of a typosquatting attack is [crossenv](https://security.snyk.io/package/npm/crossenv). The attacker used a name similar to the popular package, [cross-env](https://security.snyk.io/package/npm/cross-env), and had even wrapped the exact same functionality as the original module to convey that the module is indeed working as expected. However, in practice, the module also captured environment variables and sent them to an attacker-controller remote server.

### Dependency confusion attack

Dependency confusion is a tactic used by attackers who upload malicious packages with private package names to public package registries. The result of this is confusion of automatic processes that import packages, for example, IDE and CIs, so these processes download the public package instead of the intended private one.

For example, ACME Corporation developed an internal UI package and named it `acme-ui`, and hosted it in a private package registry.&#x20;

Malicious Bob published a malicious package with the same name to npm. Developer Alice at ACME did not configure the environment to pull packages from the private registry. While installing packages, the developer mistakenly downloaded the `acme-ui` package from the public registry instead of the private one, allowing Malicious Bob to run code on the developer's machine, a successful instance of dependency confusion.&#x20;

### Dependency hijacking

In this attack, the package the developer is using is a legitimate public package, which downstream is using the malicious package as a dependency, for example, due to dependency confusion.

### Compromised accounts

In this attack, a bad actor compromised the account of a public package maintainer and inserted malicious code into a package.

## Verifying the provenance of packages

Today, Snyk Open Source and Snyk Container scanners cannot distinguish between internal and external packages, that is, whether a package was imported from a public or private registry. This capability will be added in the future, and until then, you might see inaccurate alerts related to packages imported from a private registry.

A Snyk alert does not imply that the organizations associated with the packages published the malware; rather, the alert is a warning that such an attack has targeted the organizations.

An internal package may be flagged as malicious. If this happens, it means that a package with a similar name was published on a public registry, most likely with malicious code. This publication was intended to make developers in your company import the malicious public package, either by automatic or manual processes, instead of importing the legitimate private package.

[Learn more](malicious-packages.md#verifying-the-npm-registry-source) how to check the registry source of an npm package.

## Understanding malicious packages in npm

Due to the popularity of JavaScript and npm, most malicious packages target this ecosystem.&#x20;

The following nuances are important to understand.

### **“Security holding” in npm**&#x20;

A “security holding” on a package means that the package contained malicious code and was removed from the registry by the npm Security Team. A security placeholder was published to ensure users are not affected in the future. Although the package is currently under a “security holding” placeholder, it is important to verify that it was not imported from npm before it was marked in this placeholder. After you confirm that all instances of this package were downloaded from a private registry, you can ignore the issue.

The Snyk Vulnerability Database also shows packages in a “Security Holding” state, for example, the [flatmap-stream](https://security.snyk.io/package/npm/flatmap-stream) package.

### **Malicious package is not available on npm and without “security holding”**

The package was available publicly in npm, possibly only briefly, before being removed, that is,  [unpublished](https://docs.npmjs.com/policies/unpublish), by the owner. Under certain conditions, this can result in the package page being unavailable, though it is possible to see that it was once in use. The name is still available to be used again, with no guarantee that its contents in the new incarnation will be safe.

In these cases, it is important to be extra cautious about where the package was downloaded. Ignoring the issue altogether may result in a future problem in case a malicious actor republishes the package to npm.

## Verifying the npm registry source

To check if you are using a public registry or a private one, you can use the following options:

### Run CLI commands

Run the following command:  `npm config get registry`

If the result is `registry.npmjs.org`, you are using a public registry.

### Check package-lock.json file:

The `package-lock.json` file contains detailed information about the packages and their sources in your project. You can open this file and inspect the "`resolved`" field for each dependency to see where it is fetched from. Packages from the public npm registry will have URLs starting with https://registry.npmjs.org/.

### Check .npmrc file:

The `.npmrc` file in your project's root directory can specify the registry where npm should fetch packages. Check if your `.npmrc` file is configured to use your private npm registry. Here's an example of what the file might look like if you're using a private registry:

`registry=https://your-private-registry-url/`&#x20;

If you don't have an `.npmrc` file or it doesn't specify a registry, npm will use the default public registry (https://registry.npmjs.org/).

### Using private packages with npm

Npm allows users to host their private packages on the npm registry.

If your company is using this [service](https://docs.npmjs.com/about-private-packages), then the public registry will be https://registry.npmjs.org.&#x20;

In this case, you will need to verify that the packages used are private, by logging in to the npm website.&#x20;

## Remediation of malicious package findings

If you find that a malicious package has been running in your environment, you should remove it from the local drive, both node\_modules and cache, and remove it from the proxy cache and database if it exists, and from the package-lock file. For npm, this is `package-lock.json`.

**Typosquatting**:  Remove the malicious package and switch to the correct safe package.

**Dependency confusion:** If you imported the public package, either by accident or by default in the CI, before it was placed in "security holding," be sure to remove it. Ensure your development environment and the CI pipeline are configured to use the private registry and install the same-name internal package instead.&#x20;

**Dependency hijacking** and **Compromised account**: A new safe version will usually be released after the malicious package has been identified. To resolve the issue of this kind of attack, update the package to a new version.

In addition, consider your environment to be infected, and ensure you run your internal security drills. It is recommended that you verify there are no vestiges or remainders of the malicious code, even after you have removed the malicious package.
