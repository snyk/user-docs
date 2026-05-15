# Troubleshooting Snyk for .NET

### .NET SDK resolution limitations

Snyk uses static analysis instead of the .NET SDK to resolve dependencies in the following scenarios:

* Legacy Project files: Projects use legacy non-SDK style Project files.
* Legacy Broker: Organizations use the Legacy Broker instead of the Universal Broker to connect to brokered private package repositories and SCMs.

To improve accuracy, migrate to the Universal Broker.

### Static analysis limitations

The following functionality is not supported when using static analysis:

* Private dependencies: Snyk cannot access private dependencies, including brokered and non-brokered. Transitive dependencies are not resolved.
* Build configuration: Snyk ignores `Directory.Build.props`, `Directory.Build.targets`, and `global.json` files.
* Runtime precision: Snyk cannot reliably identify the specific runtime, which may increase false positives

### Handle runtime false positives

Static analysis may flag vulnerabilities already patched in your environment because it may not detect your specific runtime patch version.

#### Vulnerabilities in SCM

If your application runs on a system updated with the latest Microsoft patches, a flagged vulnerability may not be relevant. You can ignore these vulnerabilities in the Snyk Web UI.

#### Vulnerabilities in CLI

If your production environment pulls the latest patches or you deploy a self-contained application, use the following configurations in your project file to improve accuracy:

* Latest SDK patch: If your application always runs on the latest SDK patch version in production, set `TargetLatestRuntimePatch` to `true` in the project file. Ensure you upgrade all environments to the latest runtime version.
* Self-contained applications: If you publish a self-contained app that includes the runtime, set `RuntimeFrameworkVersion` to the specific patch version in the project file. You can ignore vulnerabilities that are no longer relevant.
