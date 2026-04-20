# Build your first dashboard

This guide provides example queries to build out your first AppSec dashboard, based on key performance indicators (KPIs) and relevant use cases. These are organized by use case and explained in terms of business value and implementation considerations. While the provided queries offer a starting point, Snyk encourages you to customize them to suit your specific requirements.

See queries for the following use cases:

* [Open issues backlog](build-your-first-dashboard.md#open-issues-backlog) - reveal current AppSec risk that requires attention.
* [Aging](build-your-first-dashboard.md#aging) - track the exposure window of open issues.
* [Mean Time To Resolve (MTTR)](build-your-first-dashboard.md#mttr) - analyze the remediation velocity of engineering teams.
* [Service Level Agreement (SLA) ](build-your-first-dashboard.md#sla)- verify that issue remediation meets with your compliance requirements.
* [IDE & CLI test rates](build-your-first-dashboard.md#developers-ide-and-cli-test-usage-and-adoption) - measure the developer adoption of AppSec testing in the development stage.
* [CI/CD Pipelines test rates](build-your-first-dashboard.md#ci-cd-pipelines-test-usage-and-adoption) - measure the adoption of AppSec testing in CI/CD pipelines.
* [Repositories with highest rate of PRs with failed PR checks](build-your-first-dashboard.md#repositories-with-highest-rate-of-prs-with-failed-pr-checks) - analyze how often Snyk is surfacing new issues in PRs and how internal teams are interacting with these results.
* [PR checks by status over time](build-your-first-dashboard.md#pr-checks-by-status-over-time) - analyze trends in PR check outcomes over time.
* [Repository PR check adoption](build-your-first-dashboard.md#repository-pr-check-adoption) - track your repository PR check coverage and surface gaps.

{% hint style="warning" %}
You must update the database and schema names in the example queries provided before execution.
{% endhint %}

## Open issues backlog

### Business value

AppSec teams need to understand the current exposure to risk. To do so, various aspects of the existing issues backlog are examined:

* The number of open issues.
* The number of high or critical-severity open issues.
* If there are available fixes for open issues.

For greater context, those figures are broken down into engineering teams, applications or any meaningful business structure that will make the results more concise and actionable. The following example queries allow this examination.

### Example query - SCA

This query returns open SCA issues backlog counters, distributed by fixability and grouped by Snyk Organization.

The results are based on:

* Open high or critical issues that were found by Snyk Open Source (SCA)
* Noise cancelling:
  * Only issues of monitored projects
  * No deleted issues

```sql
SELECT  o.DISPLAY_NAME AS organization_display_name, -- Update based on the desired aggregation
        COUNT_IF(ISSUE_SEVERITY='Critical' AND COMPUTED_FIXABILITY='Fixable') AS fixable_critical_issues,
        COUNT_IF(ISSUE_SEVERITY='High'AND COMPUTED_FIXABILITY='Fixable') AS fixable_high_issues,
        COUNT_IF(ISSUE_SEVERITY='Critical' AND COMPUTED_FIXABILITY='Partially Fixable') AS partially_fixable_critical_issues,
        COUNT_IF(ISSUE_SEVERITY='High'AND COMPUTED_FIXABILITY='Partially Fixable') AS partially_fixable_high_issues,
        COUNT_IF(ISSUE_SEVERITY='Critical' AND COMPUTED_FIXABILITY='No Fix Supported') AS unfixable_critical_issues,
        COUNT_IF(ISSUE_SEVERITY='High'AND COMPUTED_FIXABILITY='No Fix Supported') AS unfixable_high_issues
FROM SNYK.SNYK.ISSUES__V_1_0 i
     INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.PROJECT_PUBLIC_ID = p.PUBLIC_ID
     INNER JOIN SNYK.SNYK.ORGS__V_1_0 o ON i.ORG_PUBLIC_ID = o.PUBLIC_ID
WHERE p.IS_MONITORED = TRUE                      -- include only monitored projects
     AND i.DELETED_AT IS NULL                    -- remove deleted issues
     AND ISSUE_STATUS = 'Open'                   -- include only open issues
     AND i.PRODUCT_NAME = 'Snyk Open Source'     -- include only Snyk Open Source
GROUP BY o.DISPLAY_NAME                          -- Update based on the desired aggregation
ORDER BY fixable_critical_issues DESC, fixable_high_issues DESC, 
    partially_fixable_critical_issues DESC, partially_fixable_high_issues DESC; -- Update based on the desired order
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (190).png" alt="Output of SQL query for SCA issues backlog counters"><figcaption><p>Output of SQL query for SCA issues backlog counters</p></figcaption></figure>

### Example query - Code

This query returns open Snyk Code issues backlog counters, distributed by severity and grouped by Snyk Organization.

The results are based on:

* Open issues that were found by Snyk Code
* Noise cancelling:
  * Only issues of monitored projects
  * No deleted issues

```sql
SELECT o.DISPLAY_NAME AS organization_display_name, -- Update based on the desired aggregation
        COUNT_IF(ISSUE_SEVERITY='High') AS high_issues,
        COUNT_IF(ISSUE_SEVERITY='Medium') AS medium_issues,
        COUNT_IF(ISSUE_SEVERITY='Low') AS low_issues
FROM SNYK.SNYK.ISSUES__V_1_0 i
     INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.PROJECT_PUBLIC_ID = p.PUBLIC_ID
     INNER JOIN SNYK.SNYK.ORGS__V_1_0 o ON i.ORG_PUBLIC_ID = o.PUBLIC_ID
WHERE p.IS_MONITORED = TRUE                     -- include only monitored projects
     AND i.DELETED_AT IS NULL                   -- remove deleted issues
     AND ISSUE_STATUS = 'Open'                  -- include only open issues
     AND i.PRODUCT_NAME = 'Snyk Code'           -- include only Snyk Open Source
GROUP BY o.DISPLAY_NAME                         -- Update based on the desired aggregation
ORDER BY high_issues DESC, 
         medium_issues DESC, 
         low_issues DESC;                     -- Update based on the desired order
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (191).png" alt="Output format of SQL query for open Snyk Code issues backlog counters"><figcaption><p>Output format of SQL query for open Snyk Code issues backlog counters</p></figcaption></figure>

## Aging

### Business value

Issue aging refers to the time elapsed between an issue’s introduction and the current date. Organizations are concerned about this metric as the exploitation likelihood increases as the exposure window extends.

To mitigate this risk, AppSec teams monitor a predefined SLA criteria, which specifies when an issue has exceeded the expected remediation timeframe.

{% hint style="info" %}
When an issue was reintroduced, the aging is counted based on the last introduction date.
{% endhint %}

### Example query

The query below returns the average aging (in days) of critical issues per Snyk organization.\
The results are based on:

* Open critical issues
* Noise cancelling:
  * Only issues of monitored projects
  * No deleted issues

```sql
SELECT o.DISPLAY_NAME AS organization_display_name, 
    ROUND(AVG(
    CASE
        WHEN LAST_INTRODUCED IS NULL THEN DATEDIFF('DAY', TO_DATE(FIRST_INTRODUCED), CURRENT_DATE)
        WHEN TO_DATE(FIRST_INTRODUCED) <= TO_DATE(LAST_INTRODUCED) THEN DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED), CURRENT_DATE)
    END),0) AS open_issues_aging
FROM SNYK.SNYK.ISSUES__V_1_0 i
     INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.PROJECT_PUBLIC_ID = p.PUBLIC_ID
     INNER JOIN SNYK.SNYK.ORGS__V_1_0 o ON i.ORG_PUBLIC_ID = o.PUBLIC_ID 
WHERE p.IS_MONITORED = TRUE               -- include only monitored projects
     AND i.DELETED_AT IS NULL             -- remove deleted issues
     AND ISSUE_STATUS = 'Open'            -- include only open issues
     AND ISSUE_SEVERITY IN ('Critical')   -- include only critical issues
GROUP BY o.DISPLAY_NAME                   -- Update based on the desired aggregation
ORDER BY open_issues_aging DESC;          -- Update based on the desired order
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (192).png" alt="Output format of SQL query for average aging of critical issues"><figcaption><p>Output format of SQL query for average aging of critical issues</p></figcaption></figure>

## MTTR

### Business value

The MTTR (Mean Time to Resolve) metric tracks the average time it takes to resolve a security issue. It is calculated based on issues that have already been resolved and is measured over a predefined period (typically monthly, quarterly, or annually) according to their last resolution date.

Analyzing the MTTR results provides insight into the the remediation velocity of engineering teams. However, it is important to always measure both MTTR and Aging, as issues that remain open for long periods won’t show up in the MTTR results until they are remediated.

### Example query

The query below returns last month’s MTTR results per issue severity per Snyk organization.\
The results are based on:

* Issues that were resolved in the last month
* Noise cancelling:
  * Only issues of monitored projects
  * No deleted issues

```sql
SELECT 
    o.DISPLAY_NAME AS organization_display_name,
    ROUND(AVG(CASE WHEN ISSUE_SEVERITY = 'Critical' THEN 
        DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED),TO_DATE(LAST_RESOLVED)) ELSE NULL END),2) AS critical_mttr,
    ROUND(AVG(CASE WHEN ISSUE_SEVERITY = 'High' THEN 
        DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED),TO_DATE(LAST_RESOLVED)) ELSE NULL END),2) AS high_mttr,
    ROUND(AVG(CASE WHEN ISSUE_SEVERITY = 'Medium' THEN 
        DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED),TO_DATE(LAST_RESOLVED)) ELSE NULL END),2) AS medium_mttr,
    ROUND(AVG(CASE WHEN ISSUE_SEVERITY = 'Low' THEN 
        DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED),TO_DATE(LAST_RESOLVED)) ELSE NULL END),2) AS low_mttr
FROM SNYK.SNYK.ISSUES__V_1_0 i
    INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.PROJECT_PUBLIC_ID = p.PUBLIC_ID
    INNER JOIN SNYK.SNYK.ORGS__V_1_0 o ON i.ORG_PUBLIC_ID = o.PUBLIC_ID 
WHERE IS_MONITORED = TRUE                       -- include only monitored projects
     AND i.DELETED_AT IS NULL                   -- remove deleted issues
     AND ISSUE_STATUS = 'Resolved'              -- include only resolved issues
     -- issues that were resolved in the last month
     AND TO_DATE(LAST_RESOLVED) >= DATE_TRUNC('MONTH', DATEADD('MONTH', -12, CURRENT_DATE))
     AND TO_DATE(LAST_RESOLVED) <= DATEADD('DAY', -1, DATE_TRUNC('MONTH', CURRENT_DATE))
GROUP BY organization_display_name
ORDER BY organization_display_name ASC;         -- Update based on the desired order
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (193).png" alt="Output format of SQL query for MTTR per issue severity"><figcaption><p>Output format of SQL query for MTTR per issue severity</p></figcaption></figure>

## SLA

### Business value

Remediating vulnerabilities is a crucial practice, however it slows down product development that drives companies' business. Due to this simple fact, engineering teams may neglect open vulnerabilities in favor of product development tasks.

Establishing a service-level agreement (SLA) for vulnerability remediation helps maintaining that fine balance and ensure that while moving forward with product development, evolving security risks are being addressed according to a clear and transparent policy.

SLA targets define the acceptable exposure window for a vulnerability based on factors such as severity, business criticality of the asset, code ownership, or other risk factors.

Snyk issues data enables AppSec teams to track issue aging and identify which vulnerabilities have exceeded SLA targets.

### Example query

The query below returns counters of open issues per SLA status (within SLA, at risk, breached) broken down into issue severities.

The results are based on:

* Open issues
* Noise cancelling:
  * Only issues of monitored projects
  * No deleted issues

```sql
WITH 
    base AS (
    SELECT 
        CASE
            WHEN LAST_INTRODUCED IS NULL THEN DATEDIFF('DAY', TO_DATE(FIRST_INTRODUCED), CURRENT_DATE)
            WHEN TO_DATE(FIRST_INTRODUCED) <= TO_DATE(LAST_INTRODUCED) THEN DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED), CURRENT_DATE)
        END AS ISSUE_AGE,
        ISSUE_SEVERITY,     
        CASE 
           WHEN ISSUE_SEVERITY = 'Critical' AND ISSUE_AGE > c.CRITICAL THEN 'Breached' 
           WHEN ISSUE_SEVERITY = 'Critical' AND ISSUE_AGE >= (c.CRITICAL-c.CRITICAL_AT_RISK) THEN 'At Risk' 
           WHEN ISSUE_SEVERITY = 'Critical' AND ISSUE_AGE < (c.CRITICAL-c.CRITICAL_AT_RISK) THEN 'Within SLA' 
           WHEN ISSUE_SEVERITY = 'High' AND ISSUE_AGE > h.HIGH THEN 'Breached' 
           WHEN ISSUE_SEVERITY = 'High' AND ISSUE_AGE >= (h.HIGH-h.HIGH_AT_RISK) THEN 'At Risk' 
           WHEN ISSUE_SEVERITY = 'High' AND ISSUE_AGE < (h.HIGH-h.HIGH_AT_RISK) THEN 'Within SLA' 
           WHEN ISSUE_SEVERITY = 'Medium' AND ISSUE_AGE > m.MEDIUM THEN 'Breached' 
           WHEN ISSUE_SEVERITY = 'Medium' AND ISSUE_AGE >= (m.MEDIUM-m.MEDIUM_AT_RISK) THEN 'At Risk'
           WHEN ISSUE_SEVERITY = 'Medium' AND ISSUE_AGE < (m.MEDIUM-m.MEDIUM_AT_RISK) THEN 'Within SLA'  
           WHEN ISSUE_SEVERITY = 'Low' AND ISSUE_AGE > l.LOW THEN 'Breached' 
           WHEN ISSUE_SEVERITY = 'Low' AND ISSUE_AGE >= (l.LOW-l.LOW_AT_RISK) THEN 'At Risk' 
           WHEN ISSUE_SEVERITY = 'Low' AND ISSUE_AGE < (l.LOW-l.LOW_AT_RISK) THEN 'Within SLA' 
       END AS SLA_STATUS       
       FROM SNYK.SNYK.ISSUES__V_1_0 i 
            INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.project_public_id = p.public_id
            -- set the SLA TARGETS and AT RISK threshold inside the select clause of each table below
            CROSS JOIN (SELECT 15 AS CRITICAL, 3 AS CRITICAL_AT_RISK) AS c
            CROSS JOIN (SELECT 30 AS HIGH, 10 AS HIGH_AT_RISK) AS h 
            CROSS JOIN (SELECT 90 AS MEDIUM, 20 AS MEDIUM_AT_RISK) AS m 
            CROSS JOIN (SELECT 180 AS LOW, 30 AS LOW_AT_RISK) AS l 
       WHERE IS_MONITORED = true    -- include only monitored projects
        AND ISSUE_STATUS = 'Open'   -- include only open issues
        AND i.DELETED_AT IS NULL    -- remove deleted issues
    )
SELECT
    SLA_STATUS,
    SUM(CASE WHEN ISSUE_SEVERITY = 'Critical' THEN 1 ELSE 0 END) AS critical,
    SUM(CASE WHEN ISSUE_SEVERITY = 'High' THEN 1 ELSE 0 END) AS high,
    SUM(CASE WHEN ISSUE_SEVERITY = 'Medium' THEN 1 ELSE 0 END) AS medium,
    SUM(CASE WHEN ISSUE_SEVERITY = 'Low' THEN 1 ELSE 0 END) AS low
FROM base
GROUP BY SLA_STATUS
ORDER BY SLA_STATUS
```

{% hint style="info" %}
The example query can be extended to support various SLA use-cases, such as defining different SLA targets per Snyk orgs or groups, drilling-down into the at-risk or breached issues and prioritize their remediation or analyzing the SLA status for different business units.
{% endhint %}

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (194).png" alt="Output format of SQL query for open issues counter per SLA status"><figcaption><p>Output format of SQL query for open issues counter per SLA status</p></figcaption></figure>

## Developers IDE & CLI test usage and adoption

### Business value

This section demonstrate how you can discover the adoption of Snyk IDE & CLI tests by your developers. Implementing AppSec testing during the development phase is regarded as one of the most cost-effective methods for preventing new security risks from reaching production. It is more efficient because developers are already in the right context to address issues before the code progresses further in the SDLC. Detecting issues in later stages requires developers to switch context and revisit the problem, which can be less efficient and more time-consuming.

### Example query

The query below returns the unique developers and total number of scans per environment and Snyk Product.

The results are based on:

* Tests executed
* Excluding tests that were performed during CI/CD stage

```sql
SELECT
    ENVIRONMENT_DISPLAY_NAME AS IDE,
    PRODUCT_DISPLAY_NAME AS PRODUCT,
    COUNT(DISTINCT USER_EMAIL) AS UNIQUE_DEVELOPERS,
    COUNT(1) AS TOTAL_SCANS
from SNYK.SNYK.USAGE_EVENTS__V_1_0
WHERE (RUNTIME_APPLICATION_DATA_SCHEMA_VERSION = 'v2' 
AND ARRAY_CONTAINS('test'::VARIANT, INTERACTION_CATEGORIES) 
AND INTERACTION_STAGE IN ('dev')
		OR RUNTIME_APPLICATION_DATA_SCHEMA_VERSION = 'v1'
	 )
GROUP BY IDE, PRODUCT
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (195).png" alt="Output format of SQL query for number of scans per Snyk environment"><figcaption><p>Output format of SQL query for number of scans per Snyk environment</p></figcaption></figure>

## CI/CD pipelines test usage and adoption

### Business value

Preventing vulnerabilities from reaching production involves placing security gates throughout the software development lifecycle (SDLC). One of the most common gates is within the CI/CD pipeline, ensuring that any vulnerabilities missed in earlier stages are caught and blocked during the build process.

Leveraging Snyk Data Share enables you to assess the current adoption of tests and security gates within your CI/CD pipelines.

### Example query

The query below returns the number of tested repositories, total tests, and the test % success rate per Snyk Product.

The results are based on tests executed in the CI/CD stage in the last 3 months.

```sql
SELECT
    PRODUCT_DISPLAY_NAME AS PRODUCT,
    COUNT(DISTINCT INTERACTION_TARGET_ID) AS "TESTED REPOS",
    COUNT(1) AS "TOTAL SCANS",
    ROUND(((SUM(CASE WHEN INTERACTION_EXIT_CODE=0 THEN 1 ELSE 0 END))/
    (NULLIF(SUM(CASE WHEN INTERACTION_EXIT_CODE IN (0,1) THEN 1 ELSE 0 END),0))
    *100),0) AS "SUCCESS RATE"
FROM SNYK.SNYK.USAGE_EVENTS__V_1_0
WHERE INTERACTION_STAGE != 'cicd'
    AND ARRAY_CONTAINS('test'::VARIANT, INTERACTION_CATEGORIES) 
    AND INTERACTION_TIMESTAMP >= DATE_TRUNC('MONTH', DATEADD('MONTH', -3, CURRENT_DATE))
GROUP BY PRODUCT
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (196).png" alt="Output format of SQL query for number of tested repositories, total tests, and the test % success rate per Snyk Product"><figcaption><p>Output format of SQL query for number of tested repositories, total tests, and the test % success rate per Snyk Product</p></figcaption></figure>

## Repositories with highest rate of PRs with failed PR checks

### Business value

Understanding how internal teams are interacting with these PR checks helps AppSec teams assess the effectiveness of any shift-left strategy. Surfacing these patterns at the repository level enables targeted conversations with engineering teams to improve overall security posture.\
\
A high rate of failed PR checks may indicate emerging risk areas or configurations that need tuning.

A high rate of overridden checks may signal that developers are bypassing security gates, which warrants investigation into whether the overrides are justified or represent a policy gap.&#x20;

### Example query

The query below returns the total number of pull requests, the count of PRs with at least one failed check, the count of PRs with at least one overridden check, the % of PRs with failed checks, and the % of PRs with overridden checks per target repository.

The results are based on PR checks executed in the last 4 months and filtered to only include monitored and non-deleted projects.&#x20;

You can also adapt this query to prioritize surfacing repositories that have a high rate of overridden checks.

```sql
SELECT
    PRJ.TARGET_DISPLAY_NAME AS repository_name,
    COUNT(DISTINCT PRCHK.PULL_DATA_REF) AS total_prs,
    COUNT(DISTINCT IFF(PRCHK.PR_CHECK_GROUP_STATE IN ('failure'), PRCHK.PULL_DATA_REF, NULL)) AS total_prs_with_failures,
    COUNT(DISTINCT IFF(PRCHK.MARKED_AS_SUCCESS = TRUE, PRCHK.PULL_DATA_REF, NULL)) AS total_prs_marked_as_success,
    ROUND(DIV0(total_prs_with_failures, total_prs) * 100, 2) AS pct_prs_with_failed_checks,
    ROUND(DIV0(total_prs_marked_as_success, total_prs) * 100, 2) AS pct_prs_with_overridden_checks
FROM SNYK.SNYK.PR_CHECKS__V_1_0 AS PRCHK
    INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 AS PRJ
        ON PRCHK.PROJECT_PUBLIC_ID = PRJ.PUBLIC_ID
WHERE
    PRJ.DELETED IS NULL                                                     -- exclude deleted projects
    AND PRJ.IS_MONITORED = TRUE                                                 -- include only monitored projects
    AND PRCHK.CHECK_GROUP_CREATED_AT::DATE >= (CURRENT_TIMESTAMP::DATE - 120)   -- filter to pr checks within the last 4 months
GROUP BY 1
ORDER BY pct_prs_with_failed_checks DESC, pct_prs_with_overridden_checks DESC
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (304).png" alt=""><figcaption><p>Output format of SQL query for % of PRs with failed PR checks and % of PRs with overridden PR checks by repository</p></figcaption></figure>

## PR checks by status over time

### Business value

Tracking the volume of PR checks by their outcome over time can provide visibility into the overall health and trajectory of any shift-left strategy. By monitoring weekly trends in successful, failed, and errored checks, AppSec teams can detect and investigate spikes in failures and identify product configurations that need attention from error trends.\
\
An increasing success rate over time can demonstrate that developers are producing more secure code earlier in your software development life cycle.

### Example query

The query below returns weekly counts of PR check groups by if the status was success, failure, or error. The aggregation is done on the pr check group id and using the pr check group state.

The results are based on PR check groups created in the last 6 months and filtered to only include monitored and non-deleted projects.&#x20;

```sql
SELECT
    DATE_TRUNC('WEEK', PRCHK.CHECK_GROUP_CREATED_AT::DATE) AS week_starting,
    COUNT(DISTINCT IFF(PRCHK.PR_CHECK_GROUP_STATE = 'success', PRCHK.PR_CHECK_GROUP_ID, NULL)) AS total_checks_success,
    COUNT(DISTINCT IFF(PRCHK.PR_CHECK_GROUP_STATE = 'failure', PRCHK.PR_CHECK_GROUP_ID, NULL)) AS total_checks_failure,
    COUNT(DISTINCT IFF(PRCHK.PR_CHECK_GROUP_STATE = 'error', PRCHK.PR_CHECK_GROUP_ID, NULL)) AS total_checks_error
FROM SNYK.SNYK.PR_CHECKS__V_1_0 AS PRCHK
    INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 AS PRJ
        ON PRCHK.PROJECT_PUBLIC_ID = PRJ.PUBLIC_ID
WHERE
    PRJ.DELETED IS NULL                                                     -- exclude deleted projects
    AND PRJ.IS_MONITORED = TRUE                                                 -- include only monitored projects
    AND PRCHK.CHECK_GROUP_CREATED_AT::DATE >= DATEADD('MONTH', -6, CURRENT_DATE) -- filter to the last 6 months
GROUP BY 1
ORDER BY week_starting ASC
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (308).png" alt=""><figcaption></figcaption></figure>

<figure><img src="../../../../.gitbook/assets/image (306).png" alt=""><figcaption></figcaption></figure>

## Repository PR check adoption

### Business value

Ensuring PR checks are enabled across your repositories is critical to preventing new security risks from reaching your production environments. Tracking PR check adoption at the repository level helps AppSec teams identify and tackle coverage gaps. Comparing current adoption against a previous period surfaces whether coverage is improving or regressing over time.

### Example query

The query below returns PR check enablement status per repository for the current 30-day period vs. the preceding 30-day period for both Snyk Code and Snyk Open Source.

PR check enablement is resolved by combining project-level and integration-level settings. Project-level settings overrides take priority when they exist. Otherwise the integration settings apply. A repository is considered covered if at least one organization importing it has PR Checks enabled for all projects associated with that product type for at least one integration.

Results will be N/A when a repository is new to Snyk or if the specific Snyk product does not apply.&#x20;

```sql
WITH
    -- query latest integration data for the current period
    integration_current AS (
        SELECT ORG_SOURCE_PUBLIC_ID,
               IS_PULL_REQUEST_TEST_OPEN_SOURCE_ENABLED,
               IS_PULL_REQUEST_TEST_CODE_ENABLED
        FROM SNYK.SNYK.PR_CHECKS_INTEGRATION_ADOPTION__V_1_0
        WHERE
          EFFECTIVE_AT::DATE <= CURRENT_DATE
          AND (ENDS_AT IS NULL OR ENDS_AT::DATE >= DATEADD('DAY', -30, CURRENT_DATE))
          AND ORG_SOURCE_DELETED IS NULL                            -- remove deleted integrations

        QUALIFY ROW_NUMBER() OVER (
            PARTITION BY ORG_SOURCE_PUBLIC_ID ORDER BY EFFECTIVE_AT DESC, ENDS_AT DESC
        ) = 1
    ),

    -- query latest integration data for the previous period
    integration_previous AS (
        SELECT ORG_SOURCE_PUBLIC_ID,
               IS_PULL_REQUEST_TEST_OPEN_SOURCE_ENABLED,
               IS_PULL_REQUEST_TEST_CODE_ENABLED
        FROM SNYK.SNYK.PR_CHECKS_INTEGRATION_ADOPTION__V_1_0
        WHERE
          EFFECTIVE_AT::DATE <= DATEADD('DAY', -31, CURRENT_DATE)
          AND (ENDS_AT IS NULL OR ENDS_AT::DATE >= DATEADD('DAY', -61, CURRENT_DATE))
          AND ORG_SOURCE_DELETED IS NULL                            -- remove deleted integrations

        QUALIFY ROW_NUMBER() OVER (
            PARTITION BY ORG_SOURCE_PUBLIC_ID ORDER BY EFFECTIVE_AT DESC, ENDS_AT DESC
        ) = 1
    ),

    -- query latest project pr check setting data for the current period
    project_adoption_current AS (
        SELECT p.TARGET_DISPLAY_NAME AS repository_name,
               prj.ORG_SOURCE_PUBLIC_ID,
               prj.PRODUCT_NAME,
               prj.TEST_PULL_REQUESTS
        FROM SNYK.SNYK.PR_CHECKS_PROJECT_ADOPTION__V_1_0 AS prj
            INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 AS p
                ON prj.PROJECT_PUBLIC_ID = p.PUBLIC_ID

        WHERE
          prj.EFFECTIVE_AT::DATE <= CURRENT_DATE
          AND (prj.ENDS_AT IS NULL OR prj.ENDS_AT::DATE >= DATEADD('DAY', -30, CURRENT_DATE))
          AND prj.PROJECT_DELETED IS NULL                                                       -- exclude deleted projects
          AND p.IS_MONITORED = TRUE                                                             -- include only monitored projects

        QUALIFY ROW_NUMBER() OVER (
            PARTITION BY prj.PROJECT_PUBLIC_ID ORDER BY prj.EFFECTIVE_AT DESC, ENDS_AT DESC
        ) = 1
    ),

    -- query latest project pr check setting data for the previous period
    project_adoption_previous AS (
        SELECT p.TARGET_DISPLAY_NAME AS repository_name,
               prj.ORG_SOURCE_PUBLIC_ID,
               prj.PRODUCT_NAME,
               prj.TEST_PULL_REQUESTS
        FROM SNYK.SNYK.PR_CHECKS_PROJECT_ADOPTION__V_1_0 AS prj
            INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 AS p
                ON prj.PROJECT_PUBLIC_ID = p.PUBLIC_ID

        WHERE
          prj.EFFECTIVE_AT::DATE <= DATEADD('DAY', -31, CURRENT_DATE)
          AND (prj.ENDS_AT IS NULL OR prj.ENDS_AT::DATE >= DATEADD('DAY', -61, CURRENT_DATE))
          AND prj.PROJECT_DELETED IS NULL                                                       -- exclude deleted projects
          AND p.IS_MONITORED = TRUE                                                             -- include only monitored projects

        QUALIFY ROW_NUMBER() OVER (
            PARTITION BY prj.PROJECT_PUBLIC_ID ORDER BY prj.EFFECTIVE_AT DESC, ENDS_AT DESC
        ) = 1
    ),

    -- combine project and integration adoption data for current period
    project_integration_pr_check_adoption_current AS (
        SELECT repository_name,
               prj.ORG_SOURCE_PUBLIC_ID,
               prj.product_name,
               -- project pr check settings override the integration settings if they exist
               -- only consider enabled for a specific product if all projects for that product are enabled
               MIN(
               COALESCE(
                   prj.test_pull_requests,
                   CASE
                       WHEN prj.PRODUCT_NAME = 'Snyk Open Source'
                           THEN NVL(integ.IS_PULL_REQUEST_TEST_OPEN_SOURCE_ENABLED, TRUE)
                       WHEN prj.PRODUCT_NAME = 'Snyk Code'
                           THEN NVL(integ.IS_PULL_REQUEST_TEST_CODE_ENABLED, FALSE)
                   END
               ))::BOOLEAN AS pr_checks_enabled
        FROM project_adoption_current AS prj
            INNER JOIN integration_current AS integ
                ON prj.ORG_SOURCE_PUBLIC_ID = integ.ORG_SOURCE_PUBLIC_ID

        GROUP BY 1, 2, 3
    ),

    -- combine project and integration adoption data for previous period
    project_integration_pr_check_adoption_previous AS (
        SELECT repository_name,
               prj.product_name,
               prj.org_source_public_id,
               -- project pr check settings override the integration settings if they exist
               -- only consider enabled for a specific product if all projects for that product are enabled
               MIN(
               COALESCE(
                   prj.test_pull_requests,
                   CASE
                       WHEN prj.PRODUCT_NAME = 'Snyk Open Source'
                           THEN NVL(integ.IS_PULL_REQUEST_TEST_OPEN_SOURCE_ENABLED, TRUE)
                       WHEN prj.PRODUCT_NAME = 'Snyk Code'
                           THEN NVL(integ.IS_PULL_REQUEST_TEST_CODE_ENABLED, FALSE)
                   END
               ))::BOOLEAN AS pr_checks_enabled
        FROM project_adoption_previous AS prj
            INNER JOIN integration_previous AS integ
                ON prj.ORG_SOURCE_PUBLIC_ID = integ.ORG_SOURCE_PUBLIC_ID

        GROUP BY 1, 2, 3
    ),

    repository_pr_check_adoption_current AS (
        SELECT repository_name,
               -- consider enabled if the repository is enabled for any organizations or integrations
               MAX(IFF(PRODUCT_NAME = 'Snyk Code', pr_checks_enabled::INT, NULL))::BOOLEAN
                   AS snyk_code_enabled,
               MAX(IFF(PRODUCT_NAME = 'Snyk Open Source', pr_checks_enabled::INT, NULL))::BOOLEAN
                   AS snyk_open_source_enabled
        FROM project_integration_pr_check_adoption_current
        GROUP BY 1
    ),

    repository_pr_check_adoption_previous AS (
        SELECT repository_name,
               -- consider enabled if the repository is enabled for any organizations or integrations
               MAX(IFF(PRODUCT_NAME = 'Snyk Code', pr_checks_enabled::INT, NULL))::BOOLEAN
                    AS snyk_code_enabled,
               MAX(IFF(PRODUCT_NAME = 'Snyk Open Source', pr_checks_enabled::INT, NULL))::BOOLEAN
                   AS snyk_open_source_enabled
        FROM project_integration_pr_check_adoption_previous
        GROUP BY 1
    )

SELECT
    c.repository_name,
    CASE
        WHEN c.snyk_code_enabled = true THEN 'Enabled'
        WHEN c.snyk_code_enabled = false THEN 'Not Enabled'
        ELSE 'N/A'
    END AS snyk_code_pr_check_enabled_current,
    CASE
        WHEN p.snyk_code_enabled = true THEN 'Enabled'
        WHEN p.snyk_code_enabled = false THEN 'Not Enabled'
        ELSE 'N/A'
    END AS snyk_code_pr_check_enabled_previous,
    CASE
        WHEN c.snyk_open_source_enabled = true THEN 'Enabled'
        WHEN c.snyk_open_source_enabled = false THEN 'Not Enabled'
        ELSE 'N/A'
    END AS snyk_open_source_pr_check_enabled_current,
    CASE
        WHEN p.snyk_open_source_enabled = true THEN 'Enabled'
        WHEN p.snyk_open_source_enabled = false THEN 'Not Enabled'
        ELSE 'N/A'
    END AS snyk_open_source_pr_check_enabled_previous
FROM repository_pr_check_adoption_current AS c
    LEFT JOIN repository_pr_check_adoption_previous AS p
        ON c.repository_name = p.repository_name
ORDER BY repository_name
```

#### **Output format:**

<figure><img src="../../../../.gitbook/assets/image (278).png" alt=""><figcaption></figcaption></figure>
