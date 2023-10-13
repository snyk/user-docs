# Introduction to automated security scans with PR Checks

The Snyk PR Checks feature allows developers to auto-scan their PRs for issues before merging their code. As Snyk acts as an automated pseudo-team member (an “expert security reviewer”), it finds potential issues, leaving review notes on your PR before the code is committed.

While PRs are the points in the development process where code reviews happen, PR Checks allow security scanning to be integrated with developer workflows, empowering dev teams and helping to prevent security issues from occurring in deployed code.

## **Test the change**

The Snyk PR Checks feature allows you to test a change to the current codebase to see if that change introduces a problem. This change testing makes it easier to maintain the security of your codebase on an ongoing basis.

For developers, change-related flaws are relevant and easy to fix, and fixing change-related flaws rather than accumulated flaws makes rolling out secure code easier. You can detect security issues early in the development process, see the test results immediately after you write new code, and find and fix issues as they emerge, all in your native workflow.

## **Testing “before” and “after”**

The Snyk PR Checks feature runs live tests of the “before and after” branch with the PR and fails only if the new branch has more issues. This allows you to address problems that have been introduced since the last scan, for example, new vulnerabilities introduced externally; Snyk PR Checks identifies problems only with your specific code change.
