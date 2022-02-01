# Getting started with Snyk intel vulnerability DB access

## Introduction

Our [vulnerability database](https://snyk.io/product/vulnerability-database/) contains the key security information used by our Snyk products to find and fix code vulnerabilities.

For customers who already have their own products, you can still benefit from Snyk’s expertise and accumulated knowledge, with access to this database, giving your development teams access to trusted intelligence, allowing them to rapidly secure open source and container code.

## Process overview

1. Snyk helps you to set an integration up for your company.
2. Snyk provides documentation with instructions for access.
3. Snyk sends you DB information, typically as a JSON file, containing the DB information \(see [sample code](https://snyk.io/partners/api/v4/vulndb/sample.json)\)  **Note**: we recommend that you save the file in a database.
4. You write code to use the DB information in your systems.

## About the DB

A team of security experts and analysts manages Snyk's security database to ensure the database maintains high accuracy and eliminates false positives.

* All items in the database are analyzed and verified.
* The team also invests in proprietary research to discover new vulnerabilities. See our [disclosed vulnerability list](https://app.snyk.io/disclosed-vulnerabilities).  

## Database feeds

Snyk’s security database includes two feeds:

* Application security vulnerabilities: supporting Snyk Open Source, with manually-curated content and summaries, including code snippets where applicable.
* Linux OS vulnerabilities, supporting Snyk Container.

Both feed options can be licensed directly.

