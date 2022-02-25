# Fixed in version vs. fixable attribute in vulnerabilities

### A fixed vulnerability

Will not show up in a project vulnerability list as it is no longer considered a vulnerability.

### Fixed in version

This shows the version of the package that no longer has the vulnerability.

![](../../../.gitbook/assets/fix-desc-1.png)

Compare the fixed in vulnerability card above to one where no fix is available.

![](<../../../.gitbook/assets/fix-desc-2.png (1).png>)

### Fixable

A fixable vulnerability means there is a route within the project that would bring in the secure version rather than the vulnerable version.

This means that a vulnerability can be both fixable and have a fixed in option. If it was fixed, it would not appear in the project list of vulnerabilities as it would then be considered secure.

The easiest way to tell if a vulnerability is fixable in the Snyk app is to look for the "fix this vulnerability" call to action button on the vulnerability card.

![](../../../.gitbook/assets/fix-desc-3.png.png)

### Fixed in showing when issue is not fixable

The difference here is whether it's looking at direct or transitive dependencies. For direct dependencies, this would mean that fixable is true if a fixed (or secure) version of the package exists anywhere in the system. However, this is not the case for transitive dependencies as they require a direct dependency that can be updated to the fixed (or secure) version of the transitive dependency.

![](../../../.gitbook/assets/fix-desc-4.png.png)

The above is an example of a transitive dependency. The detailed paths section (blue outline in image above) shows that no fix path is available; however, it does show that the vulnerability is fixed in the more recent version unlike the no fix available status seen above. This means that Snyk doesn't have the ability to reach to the level that the vulnerability actually exists in this specific project.

To fix a transitive dependency such as this, click on the Vulnerability DB link (screenshot below) and look at the section giving fix advice for more information.

![](../../../.gitbook/assets/fix-desc-5.png)

Vulnerability DB: [https://app.snyk.io/vuln/](https://app.snyk.io/vuln/)

![](../../../.gitbook/assets/fix-desc-6.png)

### Glossary

Be sure to take a look at [Snyk's Glossary of terms](https://support.snyk.io/hc/en-us/articles/360017682058-Snyk-Glossary) to get definitions of Snyk-specific terminology.
