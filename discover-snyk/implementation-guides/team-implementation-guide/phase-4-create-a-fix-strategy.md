---
description: 'Phase 4 of the Snyk team implementation: prioritize your vulnerability backlog and define focus areas for fixes'
nav_context: agnostic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

# Phase 4: Create a fix strategy

After you set up your integrations, create your Organization, and import your Projects, you have visibility into the current vulnerability backlog of your business.

## Decide important focus areas

Before you investigate specific vulnerabilities, review your Organizations and repositories (Targets in Snyk) to identify any areas of specific importance. The discovery of business-critical applications you did in Phase 1 can help inform and prioritize during this phase. For example, if your Organizations match different products, you can initially focus on the one product with the most users or where security matters most.

In your Organization, you can then consider the repositories that make up the application's different parts. Areas that handle sensitive data or are public-facing are often more important to secure, which gives you another way to narrow down your initial list of Projects to review.

{% hint style="info" %}
If you used Project attributes to add metadata to your Projects, use them to filter down the number of Projects that you are considering.
{% endhint %}

## Group work by development teams

After you have a reduced set of Projects to prioritize, you can split them between development teams. For example, one development team can fix issues in your open-source and first-party code, and a separate DevOps team can handle container and base image vulnerabilities.

## Choose a prioritization method

Use filters to prioritize which issues to fix most urgently. The following search criteria are the most common when building a prioritization plan. You can use them iteratively or in combination as you analyze results.

* Severity (start with **High** and **Critical**). It is common to filter by critical severity. However, Snyk Code severities go up to High only, so if you use Snyk Code, start with High for Code Analysis results.
* [Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) (issues with **Mature** or **Proof of Concept** are more exploitable). Choosing this filter implicitly limits the results to Open Source.
* Fixable (if a fix is available by upgrading a package, the issue is much faster to fix).
* CVSS score for Open Source vulnerabilities
* [Priority Score](https://docs.snyk.io/scan-fix-and-prevent/fix/prioritize-issues-for-fixing/priority-score) (Snyk uses the preceding values to calculate this score). One strategy is to eliminate the vulnerabilities with a score of 900–1000, and then move to vulnerabilities with a score of 800–900, and so on.

Decide which metrics to use when planning your fix strategy, and get specific with your timeline. For example, if you choose to fix by severity, estimate the time it takes to resolve a vulnerability for each severity. Snyk recommends being specific with your fix strategy.

For example, if there are 50 critical-severity issues and 100 high-severity issues, you can plan two weeks to fix critical vulnerabilities and then four weeks to fix high-severity vulnerabilities, based on the size of your team and workload.

Alternatively, you can fix by issue type.

## Fix by issue type

It is also common to focus initially on specific issue types (for example, Open Source vulnerabilities), because you can then more easily compare issues across different Projects.

The following examples show processes for fixes led by different types of teams.

### Example: Developer-led priorities

Developer-led implementation, with instruction from executives to minimize license risk:

1. Resolve all license issues identified by Snyk Open Source.
2. Address all fixable Critical or High vulnerabilities in Snyk Open Source.
3. Shift focus to Code Analysis Projects (using Snyk Code), starting with High severity issues.
4. Scan the containers and IaC files that developers use to run their applications and environments.

### Example: DevSecOps-led priorities

DevSecOps-led implementation, focusing on securing your custom images and environment:

1. Scan and secure any custom base images that your development teams pull from.
2. Integrate with your container registry and scan the images you provide to your development teams.
   1. Scan the image you have chosen and built after adding common internal tools and standardizing image parameters.
   2. Ensure developers also scan their containers after adding their custom tools and packages to ensure that the container remains secure before deployment. This scan also detects application vulnerabilities.
3. After you have secured your containers, start actively scanning your IaC files and cloud environments for misconfigurations that could lead to security breaches.
4. Introduce Snyk to development teams to scan their Open Source and Code to reduce application vulnerabilities.

## Run targeted vulnerability campaigns

As you operationalize security testing in your development process, another option for your fix strategy is to run campaigns to eliminate vulnerability types, for example, SQL injection. Using Common Weakness Enumeration (CWE) in your search filters is useful in reporting to identify and log issues.

## Update your timeline

After you create your fix strategy, update the timeline for [Phase 5: Rolling out the prevention stage](phase-5-rolling-out-the-prevention-stage/).
