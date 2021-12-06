---
description: https://github.com/snyk-tech-services/snyk-delta/tree/docs/testing-fix
---

# Snyk-Delta

Snyk Delta Prevent provides the ability to get the delta between 2 Snyk snapshots. Comparing snapshots can provide the following:&#x20;

* New vulnerabilities not found in the baseline snapshot
* New license issues not found in the baseline snapshot
* Dependency delta between the 2 snapshots
  * Direct Dependencies added and removed
  * Indirect Dependencies added and removed
  * Flag path(s) carrying new vulnerabilities\


### Use Cases:

Snyk Delta is useful when running CLI-based scans, like in your local environment, git hooks, etc.

1. Clone the Snyk-Delta project locally
2. Run `npm i -g snyk-delta`&#x20;
3. Run `snyk test --json --print-deps` as this will create a dependency tree in your local folder&#x20;
4. ``

###

\
