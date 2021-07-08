# Getting started with Snyk Infrastructure as Code \(IaC\)

* [ Create a Snyk account](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360017098237-Create-a-Snyk-account/README.md)
* [ Select a Snyk product / tool](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360014959818-Select-a-Snyk-product-tool/README.md)
* [ Getting started with Snyk Open Source](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360014875297-Getting-started-with-Snyk-Open-Source/README.md)
* [ Getting started with Snyk Code](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360016765157-Getting-started-with-Snyk-Code/README.md)
* [ Getting started with Snyk Container](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360014877957-Getting-started-with-Snyk-Container/README.md)
* [ Getting started with Snyk Infrastructure as Code \(IaC\)](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360014938398-Getting-started-with-Snyk-Infrastructure-as-Code-IaC-/README.md)
* [ Getting Started with Snyk License Compliance Management](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015235618-Getting-Started-with-Snyk-License-Compliance-Management/README.md)
* [ Getting started with Snyk intel vulnerability DB access](https://github.com/snyk/user-docs/tree/53fce7f51125484bfae446936b09a98076f1d418/hc/en-us/articles/360015452178-Getting-started-with-Snyk-intel-vulnerability-DB-access/README.md)

## Getting started with Snyk Infrastructure as Code \(IaC\)

Get started with Snyk IaC to inspect, find and fix issues in configuration files for Terraform or Kubernetes \(including Helm\) environments. For more information, see [Scan your Kubernetes configuration files](https://support.snyk.io/hc/en-us/sections/360001881957-Scan-your-Kubernetes-configuration-files) and [Scan your Terraform files](https://support.snyk.io/hc/en-us/sections/360003156537-Scan-your-Terraform-files).

#### Stage 2: View configuration file issues

View results for configuration files in imported projects.

1. Select **Projects**, then click on the imported project entry, to see information for scanned configuration files, including the number of high, medium and low severity issues found. For example:  \(Issues are sorted into project types: Helm, Kubernetes and Terraform.\)
2. Click on a project to see more information and details of the issues in a configuration file: 

If you encounter any errors during import, see the [Importing projects](https://support.snyk.io/hc/en-us/sections/360000923478-Importing-projects) information.

#### Stage 3: View and fix config files

Act on the recommendations produced by Snyk IaC.

1. IaC results appear as direct issues in the relevant scanned configuration files. For example:
2. Click on an issue to see the details for that issue, and specific recommendations from Snyk IaC. For example:
3. Edit the configuration file to fix the issue identified, based on the recommendations, then commit the change.
4. Snyk automatically rescans the changed file, and you can see the change reflected in the issue display.

### For more information

See [Infrastructure as Code](https://support.snyk.io/hc/en-us/categories/360001342678-Infrastructure-as-code).

