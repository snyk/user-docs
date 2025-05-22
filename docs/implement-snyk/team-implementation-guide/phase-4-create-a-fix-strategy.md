# Phase 4: Create a fix strategy

After setting up your integrations, creating your Organization, and importing your Projects, you now have visibility into your business's current vulnerability backlog.&#x20;

## Decide important focus areas

Before diving into specific vulnerabilities, consider your Organizations and repositories (Targets in Snyk) to consider any areas of specific importance. The discovery of business-critical applications you did in Phase 1 can help inform and prioritize during this phase. For example, if your Organizations match different products, you can initially focus on the one product with the most users or where security matters most.&#x20;

In your Organization, you can then consider the repositories that make up the application's different parts. Areas that handle sensitive data or are public-facing may be more important to secure, so this could be another way to narrow down your initial list of Projects to review.

{% hint style="info" %}
If you have used **Attributes** to add metadata to your Projects, these can be a great way to filter down the number of Projects that you are considering.
{% endhint %}

## Group work by development teams

After you have your reduced set of Projects to prioritize, you may want to split these between different development teams. For example, you may have one Development team fixing issues with your open-source and first-party code and a separate DevOps team responsible for Containers and base image vulnerabilities.

## Prioritization method

Filters are available to help prioritize what issues need to be fixed more urgently. The following search criteria are most commonly used when building a prioritization plan and can be used iteratively or in combination as you analyze results.&#x20;

* Severity (Start with **High** and **Critical**). It is common to filter by critical severity. However, Snyk Code, Snyk code only goes up to High, so if you are using Snyk Code, start there for Code Analysis results.
* [Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) (Issues with **Mature** or **Proof of Concept** are more exploitable). By choosing this filter, you implicitly only filter the results to Open Source.
* Fixable (if there’s a fix available by upgrading a package, it’s much faster to fix).  &#x20;
* CVSS Score for Open Source Vulnerabilities
* [Priority Score](../../manage-risk/prioritize-issues-for-fixing/priority-score.md) (The above values are used to calculate this score). One strategy is to eliminate the vulnerabilities with a score of 900-1000, and then move to vulnerabilities with a score of 800-900, and so on.

Decide which metrics will be used when planning your fix strategy, and get specific with your timeline. If you choose to fix by severity, for example, estimate the time it will take to resolve a vulnerability per severity. It’s recommended to be specific with your fix strategy.&#x20;

**Example**

If there are fifty critical-severity issues and one hundred high-severity issues,  you may plan on two weeks to fix critical vulnerabilities and then four weeks to fix high-severity, based on the size of your team and workload.&#x20;

Alternatively, you can fix by issue type.

## Fix by issue type

It is also common to initially focus on specific issue types (for example, open source vulnerabilities), as then you can more easily compare issues across different Projects.&#x20;

Here are examples of processes based on whether fixes are led by different types of teams.

### Example: Developer-led priorities

Developer-led implementation, with instruction from executives to minimize license risk:

1. Resolve all License issues identified by Snyk Open Source.
2. Address all fixable Critical or High vulnerabilities in Snyk Open Source.
3. Shift focus to Code Analysis Projects (using Snyk Code), starting with High severity issues.
4. Scan the Containers and IaC files they are using to run their applications/environments.

### **Example: DevSecOps-led priorities**

DevSecOps-led implementation, focusing on securing your custom images and environment:

1. Scan and secure any custom base images that your development teams pull from.
2. Integrate with your Container Registry and scan the images you provide to your development teams.
   1. Scan the image you have chosen and built after adding common internal tools and standardizing image parameters.
   2. Ensure developers also scan their containers after adding their custom tools/packages to ensure that the container remains secure before deployment. This scan will also detect application vulnerabilities.
3. After you have secured your Containers, start actively scanning your IaC files and cloud environments for misconfigurations that could lead to security breaches.
4. Introduce Snyk to development teams to scan their Open Source and Code to reduce application vulnerabilities.

## Targeted Vulnerabilities Campaigns

As you operationalize security testing in your development process, another option for your fix strategy is to have campaigns to eliminate vulnerability types, for example, SQL injection. Using CWE in your search filters can be very useful in reporting to identify and log issues.

## Update your timeline

Once you have created your fix strategy, update the timeline for Phase 7.
