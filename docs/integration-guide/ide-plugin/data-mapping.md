---
description: Data fields and mapping
---

# Data Mapping

| Field name | Example | Description | Format |
| :--- | :--- | :--- | :--- |
| creationTime | "2019-07-28T12:26:32.609164Z" | Internal timestamp, will be removed in future versions | String \(date-time format\) |
| credit | "Matt Scott" | Discoverer / Reporter of the vuln | Array of strings |
| cves | CVE-2019-13990 | CVE\(s\) if exists. Not all vulns have a CVE as that takes a while to be assigned. | Array of strings \(CVE format\) |
| cvssScore | 5.6 | Cvss V3.1 score, computed based on the base score of the cvss vector | Number, 0-10 |
| cvssV3 | "CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L" | Cvss V3.1 vector, including base score. Might include partial temporal score where applicable | String \(CVSS format\) |
| description | "\#\# Overview\n\n[org.quartz-scheduler.internal:quartz-core](https://mvnrepository.com/artifact/org.quartz-scheduler.internal/quartz-core) is a job scheduling library.\n\n\nAffected versions of this package are vulnerable to XML External Entity \(XXE\) Injection\nvia the `initDocumentParser` method in a job description.\n\n\#\# Details\nXXE Injection is a type of attack against an application that parses XML input.\r\nXML is a markup language that defines a set of rules for encoding documents in a format that is both human-readable and machine-readable. By default, many XML processors allow specification of an external entity, a URI that is dereferenced and evaluated during XML processing. When an XML document is being parsed, the parser can make a request and include the content at the specified URI inside of the XML document.\r\n\r\nAttacks can include disclosing local files, which may contain sensitive data such as passwords or private user data, using file: schemes or relative paths in the system identifier.\r\n\r\nFor example, below is a sample XML document, containing an XML element- username.\r\n\r\n`xml\r\n<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\r\n <username>John</username>\r\n</xml>\r\n`\r\n\r\nAn external XML entity - `xxe`, is defined using a system identifier and present within a DOCTYPE header. These entities can access local or remote content. For example the below code contains an external XML entity that would fetch the content of `/etc/passwd` and display it to the user rendered by `username`.\r\n\r\n`xml\r\n<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\r\n<!DOCTYPE foo [\r\n <!ENTITY xxe SYSTEM \"file:///etc/passwd\" >]>\r\n <username>&xxe;</username>\r\n</xml>\r\n`\r\n\r\nOther XXE Injection attacks can access local resources that may not stop returning data, possibly impacting application availability and leading to Denial of Service.\n\n\#\# Remediation\n\nThere is no fixed version for `org.quartz-scheduler.internal:quartz-core`.\n\n\n\#\# References\n\n- [GitHub Issue](https://github.com/quartz-scheduler/quartz/issues/467)\n" | Full description of vulnerability. This field is meant for human consumption, and repeats a few machine-readable fields such as references and remediation. The format is markdown, making this easy to display to users. | String |
| disclosureTime | 2019-07-26T20:20:03Z | Timestamp of when the vulnerability was first made publicly available \(either known to us or as appears the vulnerability source\) | String \(date-time format\) |
| exploit | Functional | Snyk looks for exploits in the wild and evaluates their maturity. Snyk also writes it’s own POC to evaluate vulnerability exploitability. This knowledge goes into assessing the maturity of the exploit as appears in this field. Values are from section 3.1 in [https://www.first.org/cvss/v3.1/specification-document](https://www.first.org/cvss/v3.1/specification-document) | String, supported exploit maturity values as appear in cvss vector specification |
| isUpgradable | FALSE | Is there an upgrade a user can take to fix the vulnerability | Boolean value |
| id | SNYK-JAVA-ORGQUARTZSCHEDULERINTERNAL-455598 | Snyk’s Vulnearbility’s ID. | String |
| fixedIn | \[ "2.2", "2.6.14", "2.7.11"\] | This indicated the earliest version that is vuln-free. As this might be a backport fix, this does not mean that newer versions aren’t vulnerable to it. | Array of strings |
| language | Java | Specifies the ecosystem of the vulnerability \(java,python, etc\) | String |
| package | org.quartz-scheduler.internal:quartz-core | Package name | String |



