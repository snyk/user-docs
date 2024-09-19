# Upgrade package versions to fix vulnerabilities

Snyk will always recommend the smallest upgrade of a dependency to resolve a vulnerability.

To resolve a vulnerability in a transitive dependency, Snyk will calculate the dependency tree for your Project and determine the minimum upgrade to the direct dependency that will result in a vulnerability-free version of the indirect dependency.

Some fixes may require a major upgrade of a dependency. In this situation, if Snyk suspects a major change that would cause breakage, the Fix PR screen indicates this.

See [Upgrading dependencies with automatic PRs](../../pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/) for more details.
