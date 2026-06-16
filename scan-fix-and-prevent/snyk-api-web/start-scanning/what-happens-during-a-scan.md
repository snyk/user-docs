# What happens during a scan

During a target scan, Snyk API & Web goes through the target's URLs and interacts with every element found, filling out forms and clicking on buttons, among other actions, to perform extensive tests on your target to identify as many vulnerabilities as possible.

Due to the thorough interactions Snyk API & Web scans have with a target, you should expect many requests and an influx of information into the target.

## Scan components

There are three major components at play in a target scan, each one with a specific job:

* The **fingerprinter** identifies the technologies used on the target.
* The **crawler** goes through the target's URLs and interacts with every element found, clicking on buttons and filling in forms, among other things.
* The **scanner** finds vulnerabilities within the target's URLs.

## Scan states

A target scan has several possible states:

* As soon as a scan is requested, it gets **Queued**.
* Once a queued scan begins, its state is changed to **Started**.
* After the fingerprinter, the crawler, and the scanner have completed their jobs, the scan ends and its state is set to **Completed**.

### Additional scan states

There are some extra states:

* If an ongoing scan is stopped by a user, the scan state is changed to **Canceled**.
* If the target is unreachable or there is a connection timeout, the scan ends with **Failed**, with a message indicating the error. The same state is used if a scan fails during its execution.
* If some vulnerabilities need to be manually confirmed by the Snyk API & Web team, the scan is set to **Under Review**. After this manual review, the scan changes to **Completed**.

## After the scan completes

Once a scan is successfully finished, its scan reports and coverage reports can be generated.
