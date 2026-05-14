# Vulnerabilities with Social Trends

Snyk Social Trends shows a **Trending** notification for issues that are being actively discussed on X (formerly known as Twitter).

Social Trends are calculated as part of the priority score by default. You can look at the top tweets mentioning this vulnerability by clicking on **View Tweets**.

Snyk determines that a vulnerability is trending and displays the Trending banner as follows: trending?

* Snyk monitors mentions of known vulnerabilities on X, calculating the trend of tweets and reactions.
* Bots and other noise are canceled out to guarantee accuracy.
* Unexpected peaks in the trend cause the **Trending** notification to be displayed.
* The **Trending** notification remains live until several days after the trend dissipates.

{% hint style="info" %}
See [Priority Score](priority-score.md) for more information on how vulnerabilities are prioritized by Snyk.
{% endhint %}
