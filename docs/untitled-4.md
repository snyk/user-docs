# When I can choose, how should I decide whether to upgrade or patch?

An upgrade is usually the best way to fix a vulnerability. If both an upgrade and a patch are available, Snyk will usually recommend the upgrade.

* Our GitHub integration will suggest the best fix available. If there is both an upgrade and a patch, the fix pull request will recommend to upgrade.
* If you’re using Snyk’s CLI for Node.js, Snyk [wizard](https://snyk.io/docs/using-snyk/#wizard) lets you choose to patch, even if an upgrade is available. You might want to patch if an upgrade would be a potentially breaking change \(we highlight if this is the case\), or if you have other reasons to not upgrade for now.

If you’re unsure and would like to assess the impact before applying a fix, check Snyk’s advisory for the vulnerability.

