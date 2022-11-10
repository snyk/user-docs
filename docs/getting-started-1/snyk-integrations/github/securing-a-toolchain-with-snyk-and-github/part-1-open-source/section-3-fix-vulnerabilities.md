---
description: >-
  With our Security Gate in place and alerting us to risks, we want to get
  started fixing issues in the develop branch so we can downstream the fixes to
  the PROD branch.
---

# Section 3: Fix Vulnerabilities

With the GitHub Integration we configured in Section 1, Snyk is able to open Pull Requests to upgrade dependencies to non-vulnerable versions, helping to accelerate fixes.

## Step 1: Explore a vulnerability in more detail

Log into Snyk, and go into the `gh-actions-academy` project imported earlier. Scroll down to see the list of vulnerabilities present, ordered by [our proprietary Priority Score](https://snyk.io/blog/snyk-priority-score/). For each Vulnerability, Snyk displays:

* The module that introduced it and, in the case of transitive dependencies, its direct dependency,
* Details on the path and proposed fixes, as well as the specific vulnerable functions

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-vuln.png)

## Step 2: Create a Fix Pull Request in Snyk

When using the GitHub integration, and if a fix is available, Snyk can automatically upgrade the vulnerable dependency to a non-vulnerable version through a Pull Request. Click on "Fix this vulnerability" to do so.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-fixvuln.png)

On the next screen, you'll be able to confirm the issue to fix with this PR.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-prconfirm.png)

Looks good! Go ahead and open the PR.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-propen.png)

Once it's ready, you'll be taken to the PR in GitHub, where you can review the changes in the file diff view:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/gh-prdiff.png)

We see that CI checks completed successfully, assuring us we didn't introduce a breaking change.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/gh-prchecks.png)

Now, go ahead and merge the PR! You can also delete the branch. Back in Snyk, we can appreciate that our `package.json` file has 1 less High Severity Vulnerability.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/snyk-postpr.png)

## Step 3: Fix the rest of the vulnerabilities

Repeat steps 1 and 2 to fix all of the vulnerabilities that can be fixed.
