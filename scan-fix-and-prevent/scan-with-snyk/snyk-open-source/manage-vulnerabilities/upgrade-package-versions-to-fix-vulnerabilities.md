# Upgrade package versions to fix vulnerabilities

Snyk always recommends the smallest upgrade of a dependency to resolve a vulnerability.

To resolve a vulnerability in a transitive dependency, Snyk calculates the dependency tree for your Project and determines the minimum upgrade to the direct dependency that results in a vulnerability-free version of the indirect dependency.

Some fixes may require a major upgrade of a dependency. In this situation, if Snyk suspects a major change that would cause breakage, the Fix PR screen indicates this.

See [Upgrading dependencies with automatic PRs](../../pull-requests/snyk-pull-or-merge-requests/enable-automatic-upgrade-prs-for-new-dependency-upgrades.md) for more details.
