# Prioritizing Snyk issues

##  Prioritizing Snyk issues

Snyk offers a priority score to make the prioritization of issues as quick and easy as possible, allowing you to have the biggest impact with the least effort.

**How it works**

Snyk processes a number of issue characteristics in a proprietary algorithm that weights each to produce a score. Scores take into account factors such as the CVSS score, whether a fix is available, whether there are known exploits, and how new the vulnerability is.

Scores are calculated and shown for all issues, vulnerabilities and licenses and range from 0 to 1000, giving a high degree of granularity that reflects the many considerations taken into account. The wide range also helps to avoid many issues ending up with the same score.

Issues of all kinds \(vulnerability or license\) will be scored, with the scores being visible in the places listed below.

**View scores in projects**

Scores can be seen on each issue in the projects view, with all issues now sorted by the priority score, to show you the most pressing issues first.

Issues can be filtered on the left.

**View scores in the Reports**

The "Issues" tab within the reports includes the priority score as it's own sortable column. By default the table is already sorted by the score, to show you the most pressing issues first.

Issues can also be filtered by the score.

**View scores in the Snyk API**

Various issue-related API calls now include the scores in the response, and support filtering by the score.

Read more about the relevant API calls:

* [https://snyk.docs.apiary.io/\#reference/reporting-api/latest-issues/get-list-of-latest-issues](https://snyk.docs.apiary.io/#reference/reporting-api/latest-issues/get-list-of-latest-issues)
* [https://snyk.docs.apiary.io/\#reference/reporting-api/get-list-of-issues](https://snyk.docs.apiary.io/#reference/reporting-api/get-list-of-issues)
* [https://snyk.docs.apiary.io/\#reference/projects/all-projects/list-all-issues](https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-issues)

**Settings**

There are no settings related to the priority score. They have no active impact, so are just extra metadata, so they cannot be disabled or hidden.  


