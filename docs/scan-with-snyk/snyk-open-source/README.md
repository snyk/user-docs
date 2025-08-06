# Snyk Open Source

{% hint style="info" %}
To start scanning, see [Scan open-source libraries and licenses](scan-open-source-libraries-and-licenses/).
{% endhint %}

Snyk Open Source is a developer-first software composition analysis (SCA) solution. Snyk Open Source allows you to find and fix vulnerabilities in the open-source libraries used by your applications. You can also find and address licensing issues in or caused by these open-source libraries.

Developers worldwide use open-source code because it enables fast development. The vast majority of the code making up modern applications is open source. This growing reliance exposes organizations to security vulnerabilities and license issues.

Sometimes, these issues are rooted deep in the code. Open-source packages often reference other packages, and many vulnerabilities are found in these indirect dependencies. Developers may not realize which packages are being called. By using Snyk Open Source, you can reduce the risks introduced by open-source components. Snyk Open Source can help you find, prioritize, and fix security vulnerabilities and license risks in open-source dependencies throughout the SDLC.

Snyk Open Source is available in many common languages and platforms. See [Supported languages and package managers](../../supported-languages-package-managers-and-frameworks/).

{% hint style="info" %}
Only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.&#x20;

In the case of Projects that have a package manager, this means a release of the package manager.&#x20;

In the case of Go and Unmanaged scans (C/C++), this requires an official release or tag on the GitHub repository.
{% endhint %}

## Find and fix vulnerabilities

Use Snyk Open Source to find and fix vulnerabilities in the open-source libraries in your application. Snyk provides actionable fix advice for vulnerabilities and supports workflows to fix vulnerabilities using pull requests. For more information, see [Snyk Pull or Merge Requests](../pull-requests/snyk-pull-or-merge-requests/).

Snyk Open Source also helps prioritize and report on vulnerabilities discovered. For more information, see [Manage risk](../../manage-risk/).

## Find and fix license issues

Snyk Open Source can also scan your Projects for license compliance, checking against licenses known to Snyk. For more information, see [Open-source license compliance](scan-open-source-libraries-and-licenses/open-source-license-compliance.md).

You can also use license policies to define how your company deals with license issues. For more information, see [License policies](../../manage-risk/policies/license-policies/).
