# Vulnerabilities with Social Trends

Snyk Social Trends shows a **Trending** notification for issues that are being actively discussed on X (formerly known as Twitter).

Snyk calculates Social Trends as part of the priority score by default. To see the top tweets mentioning this vulnerability, click **View Tweets**.

Snyk determines that a vulnerability is trending and displays the Trending banner as follows:

* Snyk monitors mentions of known vulnerabilities on X, calculating the trend of tweets and reactions.
* Bots and other noise are canceled out to guarantee accuracy.
* Unexpected peaks in the trend cause the **Trending** notification to be displayed.
* The **Trending** notification remains live until several days after the trend dissipates.

{% hint style="info" %}
See [Priority Score](priority-score.md) for more information on how Snyk prioritizes vulnerabilities.
{% endhint %}
