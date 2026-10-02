---
description: How to interpret coverage results in Snyk API & Web
---

# Interpret coverage results

The **Coverage** tab shows which URLs Snyk tested and why it skipped others. Use it to answer two questions:

* Was this URL tested?&#x20;
* If not, why not?

Anyone who can view the scans of a target can open the **Coverage** tab. No additional permission is needed.

## Open the Coverage tab

1. In the **Targets** section, click the target name.
2. Click the **Scan Activity** tab.
3. Click the scan row you want to review.
4. In the scan panel, click **Coverage**.

The **Overview** tab of the scan also links to it with **Review the coverage here**. While a scan is running, the same line reads **Review the preliminary coverage here**.

### Crawl progress

Check the crawl progress and the number of URLs crawled. During the crawl, you can also see the total number of URLs to crawl. After the crawl is complete, only the crawled count remains.

### What the crawler found

Five counts describe everything the crawler discovered: **Discovered** = **Accepted** + **Limits hit** + **Deduplicated** + **Out of scope**.

| Count            | Meaning                                                                                                    |
| ---------------- | ---------------------------------------------------------------------------------------------------------- |
| **Discovered**   | Every URL the crawler reported or observed.                                                                |
| **Accepted**     | URLs sent to the scanner.                                                                                  |
| **Limits hit**   | URLs that reached a crawl limit, such as the number of visits to the same path, query string, or fragment. |
| **Deduplicated** | URLs with the same content as a URL already crawled.                                                       |
| **Out of scope** | URLs the crawler did not crawl. The row shows the reason.                                                  |

Click a count to filter the endpoints table. Click it again to clear the filter. A count that is not available shows **N/A**.

The **Response codes** section breaks down the responses of the scan by HTTP status class, as a share of requests. Click a class to filter the table the same way.

### Crawl health

This section highlights warnings raised during the crawl that affected what the crawler could reach, for example, the target rate-limiting the scan. Where a warning applies, you can jump to the endpoints it affected. When there is nothing to report, the section reads **No issues to report**.

### Read the endpoints table

The endpoints table lists the individual URLs the crawler reported. The table shows these columns by default, in this order:

| Column                               | Description                                                                                                |
| ------------------------------------ | ---------------------------------------------------------------------------------------------------------- |
| **Method**                           | The HTTP method used, such as `GET` or `POST`.                                                             |
| **Host**                             | The host the URL belongs to. Useful on targets with extra hosts, where one scan spans several hosts.       |
| **URL**                              | The URL itself.                                                                                            |
| **Status**                           | The HTTP status code the target returned, as a number. Empty if the URL was never requested.               |
| **Decision**                         | **Accepted** means Snyk sent the endpoint to the scanner. **Rejected** means it did not.                   |
| **Reason**                           | Why Snyk did not send a rejected endpoint to the scanner. Accepted endpoints read **sent to the scanner**. |
| **Authenticated**                    | Whether the endpoint was discovered during the authenticated phase of the crawl.                           |
| **Parameters**                       | How many parameters were tested on this endpoint.                                                          |
| **Request size** / **Response size** | The size of what was sent and received, in human-readable units.                                           |
| **Reported**                         | When this endpoint was last recorded during the scan.                                                      |

{% hint style="info" %}
**Authenticated** records the phase of the crawl in which the endpoint was discovered. **Yes** means the crawler found it after login. **No** means the crawler found it before login. On its own, it does not confirm that the request used a live, valid session.
{% endhint %}

On GraphQL targets, the table shows two additional columns by default: **Operation name** and **Operation type**. Because every endpoint on a GraphQL target shares one URL, these columns distinguish the rows.

Use **Manage columns** to hide any removable column or to add **ID**, the endpoint identifier, which you can use with the API.

Click a column header to sort by it. Sorting applies to the whole result set, not only the page on screen. The page address keeps the sort order, so you can reload or share a sorted view.

### Find out why an endpoint was not tested

Every rejected endpoint carries a reason, shown exactly as the scan engine recorded it. The full list of reasons and what each one means is on the [coverage report](overview-reports/#coverage-reports) page.

Filter by **Reason** to group the omissions of a scan by cause, for example, to separate URLs dropped because a crawl limit was reached from URLs excluded by your reject list.

{% hint style="info" %}
**Rejected endpoints are recorded only on request.** By default, a scan stores only accepted endpoints, so the **Coverage** tab lists only accepted endpoints, but the counts in the **What the crawler found** section can exceed the number of rows.
{% endhint %}

To record rejected endpoints, navigate to **Target Settings** > **Reports** and change **Coverage Detail** to **All URLs**. This applies from the next scan. It cannot recover rows for a scan that has already run.

### Inspect the request and response for an endpoint

Click any row to expand it. The expanded row shows the **Status**, **Decision**, **Reason**, **Authenticated**, and **Parameters** of the endpoint and, for a fetched endpoint, the recorded **Request** and **Response**. Use **Jump to** to move to the response.

You can collapse each pane individually and copy it to the clipboard with one click.

The scan engine stores no response for a rejected endpoint and builds only a request line from the URL. The expanded row states this instead of showing a synthetic exchange.

#### What is and is not stored

Three limits apply to what you see on an accepted endpoint:

* Snyk obfuscates sensitive values before storage. Snyk masks credentials, cookies, and tokens when it writes the exchange, not when it displays it. Snyk never stores the original values.
* Snyk truncates response bodies. Snyk keeps only the first 4KB of the response body and keeps headers in full. When a body is truncated, the pane shows the full size.
* Snyk keeps responses for 90 days. After that, Snyk deletes the stored response body, and the pane reads **This response is no longer stored.** Snyk does not delete the endpoint. Its URL, method, status, decision, reason, sizes, and full request remain. The 90 days start on the date of the scan, so all responses of a scan expire together.

### Filter, search, and group

You can filter the table by **Method**, **Host**, **Status**, **Decision**, **Reason**, and **Authenticated**.&#x20;

The search box matches only the URL. It does not search inside stored requests or responses.

**Group by** organizes the rows under headings by **Decision**, **Status**, **Method**, or **Reason**, without changing which rows are shown. Use it to see the shape of a scan, for example, how many endpoints fall under each reason or how the status codes are distributed. Grouping by **Status** collects rows that were never requested under **Not fetched**.

### Export to CSV

Click **Download CSV** to export the coverage data.

* The export includes every row that your filters and search match, not only the page on screen.
* The export uses the fields you select in the dialog, not the columns visible in the table. The fields include **Info** and raw-byte versions of the two size fields, which are not shown in the table. The **Host** field is not included.

The file uses the same format as the downloadable [coverage report](overview-reports/coverage-report.md), so anything you already parse against that report keeps working.

### Exclude a URL from future scans

If the **Coverage** tab shows that a URL was tested and you do not want it tested again, you can add it to the target's reject list without leaving the page.

1. Find the accepted row and click the reject icon at the end of it.
2. Check the URL in the confirmation dialog.
3. Click **Add to reject list**.

Snyk adds the URL to the reject list of the target, under **Target Settings** > **Scanner**, and skips it from the next scan. Results already recorded do not change.

These conditions apply:

* The action appears on accepted endpoints only, one row at a time.
* You must have permission to change target settings. Without it, you can view the tab and its data, but the reject icon is not shown on any row.
* On a GraphQL target, the action is available only on rows that have an operation name.

To remove a URL from the reject list, navigate to **Target Settings** > **Scanner**.

{% hint style="info" %}
* Snyk records rejection reasons for web targets only. On API targets (OpenAPI, Postman, Bruno, and GraphQL), the scan engine does not record why a URL was skipped. The **Reason** column is empty, and changing **Coverage Detail** to **All URLs** makes no difference.
* The **Coverage** tab covers the crawl phase. It shows what the crawler found and what the crawler handed to the scanner. It does not show the individual test requests that the scanner then sent against each endpoint.
* Snyk does not record response times per endpoint, so the tab cannot show which URLs were slow to respond.
{% endhint %}
