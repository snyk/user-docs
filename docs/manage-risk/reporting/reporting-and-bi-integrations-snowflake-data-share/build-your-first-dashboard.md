# Build your first dashboard

This guide provides example queries to build out your first AppSec dashboard, based on key performance indicators (KPIs) and relevant use cases. These are organized by use case and explained in terms of business value and implementation considerations. While the provided queries offer a starting point, Snyk encourages you to customize them to suit your specific requirements.

See queries for the following use cases:

* [Open issues backlog](build-your-first-dashboard.md#open-issues-backlog) - reveal current AppSec risk that requires attention.  &#x20;
* [Aging](build-your-first-dashboard.md#aging) - track the exposure window of open issues.
* [Mean Time To Resolve (MTTR)](build-your-first-dashboard.md#mttr) - analyze the remediation velocity of engineering teams.
* [Service Level Agreement (SLA) ](build-your-first-dashboard.md#sla) - verify that issue remediation meets with your compliance requirements.
* [IDE & CLI test rates](build-your-first-dashboard.md#developers-ide-and-cli-test-usage-and-adoption) - measure the developer adoption of AppSec testing in the development stage.
* [CI/CD Pipelines test rates](build-your-first-dashboard.md#ci-cd-pipelines-test-usage-and-adoption) - measure the adoption of AppSec testing in CI/CD pipelines.

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

<figure><img src="../../../.gitbook/assets/image (2) (1).png" alt="Output of SQL query for SCA issues backlog counters"><figcaption><p>Output of SQL query for SCA issues backlog counters</p></figcaption></figure>

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

<figure><img src="../../../.gitbook/assets/image (1) (18).png" alt="Output format of SQL query for open Snyk Code issues backlog counters"><figcaption><p>Output format of SQL query for open Snyk Code issues backlog counters</p></figcaption></figure>

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

<figure><img src="../../../.gitbook/assets/image (2) (19).png" alt="Output format of SQL query for average aging of critical issues"><figcaption><p>Output format of SQL query for average aging of critical issues</p></figcaption></figure>

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

<figure><img src="../../../.gitbook/assets/image (3) (1).png" alt="Output format of SQL query for MTTR per issue severity"><figcaption><p>Output format of SQL query for MTTR per issue severity</p></figcaption></figure>

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

<figure><img src="../../../.gitbook/assets/image (4) (2).png" alt="Output format of SQL query for open issues counter per SLA status"><figcaption><p>Output format of SQL query for open issues counter per SLA status</p></figcaption></figure>

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

<figure><img src="../../../.gitbook/assets/image (5).png" alt="Output format of SQL query for number of scans per Snyk environment"><figcaption><p>Output format of SQL query for number of scans per Snyk environment</p></figcaption></figure>

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

<figure><img src="../../../.gitbook/assets/image (6).png" alt="Output format of SQL query for number of tested repositories, total tests, and the test % success rate per Snyk Product"><figcaption><p>Output format of SQL query for number of tested repositories, total tests, and the test % success rate per Snyk Product</p></figcaption></figure>
