# Data Model

This section provides references to the Snyk data model and the most popular fields. For a complete reference to the Snyk API data fields please see the [Snyk API documentation](https://snyk.docs.apiary.io/#)

## Snyk Data Model

There are a few core entities and cardinality you should be familiar with in the Snyk API data model. These entities are listed here including a description and URL link to the API documentation.

* [Groups](https://snyk.docs.apiary.io/#reference/groups) - A collection of organizations in Snyk. Groups contain a Group Id, Members, Service Accounts, etc. Groups aggregate organization information and provide global policies. Customers may have zero, one, or more groups. Groups apply to Snyk Pro and Enterprise customers only.
* [Organizations](https://snyk.docs.apiary.io/#reference/organizations) - A collection of Projects in Snyk, Organizations provide configuration for Access Rights, Policies, Service Accounts, etc. Customers may have one or more organizations per group. Snyk customers using Standard or free plans only have organizations.
* [Project](https://snyk.docs.apiary.io/#reference/projects) - A project is a package that is actively tracked in Snyk and created as part of an import action. It represents the package manager manifest file that lists the set of dependencies for the application. Customers may have zero, one, or more projects per organization.
* [Issues](https://snyk.docs.apiary.io/#introduction/overview-and-entities/issue) - A collection of issues in Snyk. Issues include both licensing violations and vulnerabilities and are attached to a specific project. A project may have zero, one, or more issues per project. The issues are the actual scan results that contain the license issues and vulnerabilities.
* [Users](https://snyk.docs.apiary.io/#reference/users) - A user represents the person accessing, configuring, or being assigned to Groups, Orgs, or Projects.
* [Integrations](https://snyk.docs.apiary.io/#reference/integrations) - A integration represents the configurations of a specific DevOps tool like SCM or Container Registry. Integrations can also be used to import or remove projects from Snyk.

## Snyk Data Mapping

Snyk scan results contain a rich set of metadata about the vulnerabilities discovered. Use table below for a quick reference of the **most popular** fields and their potential usage inside your platform:

| Field name | Format | Description | Example |
| :--- | :--- | :--- | :--- |
| credit | Array of strings | Discoverer / Reporter of the vuln | "Matt Scott" |
| cves | Array of strings \(CVE format\) | CVE\(s\) if exists. Not all vulns have a CVE as that takes a while to be assigned. | CVE-2019-13990 |
| cvssScore | Number, 0-10 | Cvss V3.1 score, computed based on the base score of the cvss vector | 5.6 |
| cvssV3 | String \(CVSS format\) | Cvss V3.1 vector, including base score. Might include partial temporal score where applicable | "CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L" |
| cwes | Array of strings \(CWE format\) | The common weakness enumeration \(CWE\) of the vuln | CWE-611description |
| description | String | Full description of vulnerability. This field is meant for human consumption, and repeats a few machine-readable fields such as references and remediation. The format is markdown, making this easy to display to users. | "\#\# Overview\n\n[org.quartz-scheduler.internal:quartz-core](https://mvnrepository.com/artifact/org.quartz-scheduler.internal/quartz-core) is a job scheduling library.\n\n\nAffected versions of this package are vulnerable to XML External Entity \(XXE\) Injection\nvia the `initDocumentParser` method in a job description.\n\n\#\# Details\nXXE Injection is a type of attack against an application that parses XML input.\r\nXML is a markup language that defines a set of rules for encoding documents in a format that is both human-readable and machine-readable. By default, many XML processors allow specification of an external entity, a URI that is dereferenced and evaluated during XML processing. When an XML document is being parsed, the parser can make a request and include the content at the specified URI inside of the XML document.\r\n\r\nAttacks can include disclosing local files, which may contain sensitive data such as passwords or private user data, using file: schemes or relative paths in the system identifier.\r\n\r\nFor example, below is a sample XML document, containing an XML element- username.\r\n\r\n`xml\r\n<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\r\n <username>John</username>\r\n</xml>\r\n`\r\n\r\nAn external XML entity - `xxe`, is defined using a system identifier and present within a DOCTYPE header. These entities can access local or remote content. For example the below code contains an external XML entity that would fetch the content of `/etc/passwd` and display it to the user rendered by `username`.\r\n\r\n`xml\r\n<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\r\n<!DOCTYPE foo [\r\n <!ENTITY xxe SYSTEM \"file:///etc/passwd\" >]>\r\n <username>&xxe;</username>\r\n</xml>\r\n`\r\n\r\nOther XXE Injection attacks can access local resources that may not stop returning data, possibly impacting application availability and leading to Denial of Service.\n\n\#\# Remediation\n\nThere is no fixed version for `org.quartz-scheduler.internal:quartz-core`.\n\n\n\#\# References\n\n- [GitHub Issue](https://github.com/quartz-scheduler/quartz/issues/467)\n" |
| disclosureTime | String \(date-time format\) | Timestamp of when the vulnerability was first made publicly available \(either known to us or as appears the vulnerability source\) | 2019-07-26T20:20:03Zexploit |
| exploit | String, supported exploit maturity values as appear in cvss vector specification | Snyk looks for exploits in the wild and evaluates their maturity. Snyk also writes it’s own POC to evaluate vulnerability exploitability. This knowledge goes into assessing the maturity of the exploit as appears in this field. Values are from section 3.1 in [https://www.first.org/cvss/v3.1/specification-document](https://www.first.org/cvss/v3.1/specification-document) | Functional |
| fixable | Boolean | Is there a fixed version- i.e. a newer version without this specific vulnerability | FALSE |
| id | String | Snyk’s Vulnearbility’s ID | SNYK-JAVA-ORGQUARTZSCHEDULERINTERNAL-455598 |
| initiallyFixedIn | Array of strings | This indicated the earliest version that is vuln-free. As this might be a backport fix, this does not mean that newer versions aren’t vulnerable to it. We highly recommend using the field vulnerableVersions | \[ "2.2", "2.6.14", "2.7.11"\]language |
| language | String | Specifies the ecosystem of the vulnerability \(java,python, etc\) | Java |
| malicious | Boolean | Indication if the vuln is known to mark a malicious package | FALSE |
| package | String | Package name | org.quartz-scheduler.internal:quartz-core |



