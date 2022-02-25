# API rate limit control

### Azure DevOps

Azure DevOps has a unique way of limiting the API call rate with their own "TSTU" concept as described in this [guide](https://docs.microsoft.com/en-us/azure/devops/integrate/concepts/rate-limits?view=azure-devops).

The Snyk SCM contributors Count tool applies a strict limit of a max of 2 calls per second to deal with the rate limit.

### Bitbucket Cloud

On Bitbucket Cloud, the API rate limit is 1,000 calls per hour for authenticated users as described in this [guide](https://support.atlassian.com/bitbucket-cloud/docs/api-request-limits/).

The Snyk SCM contributors Count tool applies a strict limit of a max 1,000 calls per hour to deal with the rate limit and additional regulating mechanism to deal with 429 responses ("too many calls")

### Bitbucket Server

On Bitbucket Server, the system admin has full control of the API rate limiting as described in this [guide](https://confluence.atlassian.com/bitbucketserver/improving-instance-stability-with-rate-limiting-976171954.html).

The Snyk SCM contributors Count tool applies a moderate limit of a max 1000 calls per hour and additional regulating mechanism to deal with 429 responses ("too many calls")

### GitHub

On GithHub, API rate limit of 5,000 calls per hour for authenticated users as described in this [guide](https://docs.github.com/en/developers/apps/building-github-apps/rate-limits-for-github-apps).

The Snyk SCM contributors Count tool applies a strict limit of a max 4,500 calls per hour to deal with the rate limit and additional regulating mechanism to deal with 429 responses ("too many calls")

### GitHub Enterprise

On Github Enterprise, API rate limit of 15,000 calls per hour for authenticated users as described in this [guide](https://docs.github.com/en/developers/apps/building-github-apps/rate-limits-for-github-apps).

The Snyk SCM contributors Count tool applies a strict limit of a max 3 calls per second which amounts to 10,800 calls per hour, to deal with the rate limit and additional regulating mechanism to deal with 429 responses ("too many calls")

### GitLab and GitLab Server

On GitLab, API rate limit of 300 calls per minute for authenticated users as described in this [guide](https://docs.gitlab.com/ee/user/gitlab\_com/index.html#gitlabcom-specific-rate-limits) for GitLab and this [guide](https://docs.gitlab.com/ee/user/admin\_area/settings/rate\_limits\_on\_raw\_endpoints.html) for GitLab Server

The Snyk SCM contributors Count tool applies a strict limit of a max 120 calls per minute to deal with the rate limit and additional regulating mechanism to deal with 429 responses ("too many calls")

{% hint style="info" %}
On GitLab Server, the API rate control is configurable by the admin, as described in the guide above.
{% endhint %}
