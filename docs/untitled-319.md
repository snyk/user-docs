# What do we do when Snyk finds vulnerabilities for which there is no fix available?

## What do we do when Snyk finds vulnerabilities for which there is no fix available?

Sometimes you will find vulnerabilities for which there is no current known fix available yet and some folks come to us asking what to do in these cases.

Ultimately, the decision is up to the end user or developer - Snyk is an information tool first and foremost. Not all "fixes" should or can be performed by an upgrade. One of the reasons why we tell developers about vulnerabilities early in such a way that is formatted to them is so that they can make informed, actionable decisions themselves.

Even _if_ a dependency has an upgrade available, the developer might still choose to remediate by removing the library altogether or changing it out with a different package entirely. Maybe you mitigate using other services such as a WAF or other reverse proxy configuration.

Remember that Snyk is short for "_**S**o **N**ow **Y**ou **K**now_" :-\)

