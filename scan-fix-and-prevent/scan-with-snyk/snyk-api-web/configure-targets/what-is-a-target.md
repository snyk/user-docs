---
description: What a target is in Snyk API and Web scanning
nav_context: agnostic
---

# What is a target?

A target is the URL of a web application, website, or API that defines the scope of a security scan.

## Understanding targets and scope

Targets define what Snyk API & Web scans. Examples of targets include:

- `https://app1.example.com`
- `https://example.com`
- `https://api.example.com`

The target URL sets the boundaries for scanning. The scanner only tests URLs that begin with the target base URL and never scans outside this scope.

For example, if the target is `https://example.com`, the scanner does not test `https://www.example.com` or any other hosts. The scan includes only URLs prefixed by `example.com`.

## Organize applications with targets

Use targets to organize security testing based on your application structure. For a large application at `https://example.com` with different sections or modules built by different teams, create separate targets such as:

- `https://example.com/sectionA`
- `https://example.com/sectionB`

This approach helps manage workflows and assign security testing responsibilities to specific teams.

## Target types

Snyk API & Web supports two types of targets:

### Web targets

Web targets provide full security scans of web applications, including single-page applications and web applications that interact with APIs. Choose this option to assess the security of user-facing web interfaces.

### API targets

API targets provide detailed security assessments of standalone APIs without a supporting web application. Choose this option to test REST APIs, GraphQL endpoints, or other API implementations.

To scan an API, Snyk API & Web needs the API specification (schema). Define the schema with a URL pointing to the schema file, upload it directly, or use schema introspection for GraphQL. When you select URL or introspection, Snyk API & Web fetches the schema before every scan, ensuring the most up-to-date version.

## Adding and configuring targets

{% hint style="info" %}
Navigation sequences, seeds list, and extra hosts apply only to Web targets.
{% endhint %}

Add targets individually or import multiple targets in bulk. After adding a target, configure settings to enhance scan coverage and accuracy:

- **Authentication**: Configure login forms, login sequences, two-factor authentication (2FA), or basic authentication to scan protected areas
- **Navigation sequences**: Guide the scanner to specific application states
- **Scan profiles**: Choose scan depth and coverage based on your needs
- **Seeds and reject lists**: Control which URLs the scanner includes or excludes
- **Extra hosts**: Include additional domains in the scan scope
- **Scanning agent**: Scan internal applications not accessible from the public internet