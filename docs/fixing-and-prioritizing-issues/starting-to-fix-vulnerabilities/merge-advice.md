# Merge advice

Merge Advice is a badge we display on pull requests to indicate how confident we are that merging the pull request will not result in any breaking changes.

### How it's calculated

We determine this advice based on how well that same change has performed on other Snyk users' pull requests – did their tests on the PR pass or fail? Was the change subsequently rolled back? Did they merge successfully?

### How it's shown:

Once we've gathered enough data, we show a badge on the PR - either giving the advice "review recommended", or "high chance of success".

![merge-advice-review-recommended.png](https://support.snyk.io/hc/article_attachments/360007616777/merge-advice-review-recommended.png)  ![merge-advice-high-chance-of-success.png](https://support.snyk.io/hc/article_attachments/360007695038/merge-advice-high-chance-of-success.png)

If we haven't yet been able to collect enough data to give trustworthy advice, we show the message "not enough data". Once we've gathered enough data, we update this badge automatically with our recommendation – for that reason, a badge that was displaying "not enough data" might later show advice.

![](https://support.snyk.io/hc/article_attachments/360007695018/merge-advice-not-enough-data.png)

### Availability:

At the moment, merge advice badges are only available for Yarn and npm, where a single package is being upgraded. More support is coming soon.

All Snyk-supported source control integrations are supported for merge advice.

