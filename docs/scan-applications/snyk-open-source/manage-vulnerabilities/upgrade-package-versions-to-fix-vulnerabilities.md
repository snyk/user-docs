# Upgrade package versions to fix vulnerabilities

Snyk will always recommend the smallest upgrade of a dependency to resolve a vulnerability.

To resolve a vulnerability in a transitive dependency, Snyk will calculate the dependency tree for your Project and determine the minimum upgrade to the direct dependency which will result in a vulnerability-free version of the indirect dependency.

Some fixes may require a major upgrade of a dependency. In this situation, if Snyk suspects a major change causing breakage, that is indicated on the Fix PR screen.

See [Upgrading dependencies with automatic PRs](../../../scan-application-code/snyk-open-source/open-source-basics/upgrading-dependencies-with-automatic-prs.md) for more details.
