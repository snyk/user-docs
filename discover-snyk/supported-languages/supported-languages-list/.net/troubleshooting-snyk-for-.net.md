---
description: How to resolve common Snyk problems for .NET, including static dependency resolution for legacy Project files and Legacy Broker setups
nav_context: agnostic
last_reviewed: 2026-05-15
last_reviewed_source: bulk-sync
---

# Troubleshooting Snyk for .NET

## .NET SDK resolution limitations

Snyk uses static analysis instead of the .NET SDK to resolve dependencies in the following scenarios:

* Legacy Project files: Projects use legacy non-SDK style Project files.
* Legacy Broker: Organizations use the Legacy Broker instead of the Universal Broker to connect to brokered private package repositories and SCMs.

To improve accuracy, migrate to the Universal Broker.

## Static analysis limitations

Static analysis does not support the following:

* Private dependencies: Snyk cannot access private dependencies, brokered or non-brokered. Snyk does not resolve transitive dependencies.
* Build configuration: Snyk ignores `Directory.Build.props`, `Directory.Build.targets`, and `global.json` files.
* Runtime precision: Snyk cannot reliably identify the specific runtime, which can increase false positives.

## Handle runtime false positives

Static analysis can flag vulnerabilities that are already patched in your environment, because it cannot always detect your specific runtime patch version.

### Vulnerabilities in SCM

If your application runs on a system updated with the latest Microsoft patches, a flagged vulnerability can be irrelevant. You can ignore these vulnerabilities in the Snyk Web UI.

### Vulnerabilities in CLI

If your production environment pulls the latest patches or you deploy a self-contained application, use the following configurations in your project file to improve accuracy:

* Latest SDK patch: If your application always runs on the latest SDK patch version in production, set `TargetLatestRuntimePatch` to `true` in the project file. Ensure you upgrade all environments to the latest runtime version.
* Self-contained applications: If you publish a self-contained app that includes the runtime, set `RuntimeFrameworkVersion` to the specific patch version in the project file. You can ignore vulnerabilities that are no longer relevant.
