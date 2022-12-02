# Snyk Vulnerability Database

## Introduction

The [Snyk Vulnerability Database](https://security.snyk.io/) contains the key security information used by Snyk products to find and fix code vulnerabilities.

### Access to the Database

You can inspect the database at [https://security.snyk.io/](https://security.snyk.io/), or you can incorporate database information into your own systems.

Incorporating information into your own systems may be useful for customers who already have their own security products; you can benefit from Snyk’s expertise and accumulated knowledge with access to this database. This gives your development teams access to trusted intelligence, allowing them to rapidly secure open source and container code.

#### Feeds

The Snyk Vulnerability Database includes two feeds:

* [Application security](https://snyk.io/learn/application-security/) vulnerabilities: supporting Snyk Open Source, with manually-curated content and summaries, including code snippets where applicable.
* Linux OS vulnerabilities, supporting Snyk Container.

Both feed options can be licensed directly.

#### Incorporating the Vulnerability Database into your systems

You can incorporate Snyk Vulnerability Database information in your system. Typically:

1. Snyk helps you to set an integration up for your company, providing documentation with instructions for access.
2. Snyk sends you database information, typically as a JSON file (see [sample code](https://snyk.io/partners/api/v4/vulndb/sample.json)) **Note**: It is recommends that you save the file in a database.
3. You can now write code to use the database information in your systems.

Also see [Using the Snyk Vulnerability Database](../../../scan-application-code/starting-to-fix-vulnerabilities/using-the-snyk-vulnerability-database.md).
