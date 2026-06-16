# Invalid findings and false positives

In security scans, Snyk API & Web makes a great effort to avoid false positives because they bring a costly overhead of manual analysis and resolution of vulnerabilities.

When Snyk API & Web finds a vulnerability, a finding is created as **Not fixed**. This state can change automatically or manually, as explained in [Finding states](finding-states.md).

If a user manually changes the state to **Invalid**, it means the vulnerability does not exist and is a candidate for a false positive.

This triggers Snyk API & Web to proceed with further investigation to confirm it:

* If it is a false positive, Snyk API & Web improves the identification of vulnerabilities so that it is no longer reported.
* If not, Snyk API & Web contacts the user to explain how the vulnerability can be exploited.

## Accept risk instead of marking as invalid

In some situations, a finding should not be considered a false positive (that is, **Invalid**) and should be accepted in the specific context instead.

For example, if a target uses Cloudflare with Transport Layer Security (TLS) 1.2 with the default configuration, the weak cipher suites enabled vulnerability is reported by the scanner due to the existence of the Cipher Block Chaining (CBC) weak cipher.

However, using TLS 1.2 may be acceptable to support a broader range of users for business reasons.

In this case, consider changing the state of the finding to **Accept risk**.

Visit [Manage findings](manage-findings.md) to learn how to change the state.
