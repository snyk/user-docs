# Snyk Vs NPM Audit

You may have noticed some differences in scanning results between Snyk and NPM Audit. Snyk has its own [Vulnerability Database](https://snyk.io/vuln) that is patented by our Security team and is analyzed according to the Testing methodology

Each vulnerability in the npm database has a unique id \(nsp ip\) which was checked against Snyk's own internal vulnerability database to ensure there was a corresponding nsp\_id matched for that vulnerability. This allowed for side by side comparisons to be made between the two databases with Snyk incorporating all vulnerabilities which existed in npm’s database as well as highlighting which Snyk vulnerabilities were missing from npm. Disclosure times were compared using the public disclosure times logged by Snyk and npm, as well as the publication times declared by each database.

Checkout the differences between Snyk and NPM Audit

