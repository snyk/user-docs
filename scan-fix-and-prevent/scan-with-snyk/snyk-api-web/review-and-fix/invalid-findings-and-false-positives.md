# Invalid findings and false positives

In security scans, Snyk API & Web avoids false positives because they bring a costly overhead of manual analysis and resolution of vulnerabilities.

When Snyk finds a vulnerability, it creates a finding with a **Not fixed** state. This state can change automatically or manually, as explained in [Finding states](finding-states.md).

When a user manually changes the state to **Invalid**, the vulnerability does not exist and is a candidate for a false positive.

This triggers Snyk to investigate further to confirm it:

* If it is a false positive, Snyk improves the identification of vulnerabilities so that it no longer reports the vulnerability.
* If not, Snyk contacts the user to explain how the vulnerability can be exploited.

## Accept risk instead of marking as invalid

In some situations, a finding is not a false positive (that is, **Invalid**) and is accepted in the specific context instead.

For example, if a target uses Cloudflare with Transport Layer Security (TLS) 1.2 with the default configuration, the scanner reports the weak cipher suites enabled vulnerability because of the Cipher Block Chaining (CBC) weak cipher.

However, using TLS 1.2 is acceptable to support a broader range of users for business reasons.

In this case, change the state of the finding to **Accept risk**.

Visit [Manage findings](manage-findings.md) to learn how to change the state.
