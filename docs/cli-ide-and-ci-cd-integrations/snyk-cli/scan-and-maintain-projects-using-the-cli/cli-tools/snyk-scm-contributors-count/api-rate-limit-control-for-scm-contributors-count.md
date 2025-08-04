# API rate limit control for scm-contributors-count

## Azure DevOps

Azure DevOps has a unique way of limiting the API call rate with their own "TSTU" concept as described in this [guide](https://docs.microsoft.com/en-us/azure/devops/integrate/concepts/rate-limits?view=azure-devops).

The `snyk-scm-contributors-count` tool applies a strict limit of a maximum of two calls per second to deal with the rate limit.

## Bitbucket Cloud

On Bitbucket Cloud, the API rate limit is 1,000 calls per hour for authenticated users as described in this [guide](https://support.atlassian.com/bitbucket-cloud/docs/api-request-limits/).

The `snyk-scm-contributors-count` tool applies a strict limit of a maximum of 1,000 calls per hour to deal with the rate limit and an additional regulating mechanism to deal with 429 responses ("too many calls").

## Bitbucket Server

On Bitbucket Server, the system admin has full control of the API rate limiting as described in this [guide](https://confluence.atlassian.com/bitbucketserver/improving-instance-stability-with-rate-limiting-976171954.html).

The `snyk-scm-contributors-count` tool applies a moderate limit of a max 1000 calls per hour and additional regulating mechanism to deal with 429 responses ("too many calls")

## GitHub

On GithHub, the API rate limit is 5,000 calls per hour for authenticated users as described in this [guide](https://docs.github.com/en/developers/apps/building-github-apps/rate-limits-for-github-apps).

The `snyk-scm-contributors-count` tool applies a strict limit of a maximum of 4,500 calls per hour to deal with the rate limit and an additional regulating mechanism to deal with 429 responses ("too many calls").

## GitHub Enterprise

On Github Enterprise, the API rate limit is 5,000 calls per hour for authenticated users as described in this [guide](https://docs.github.com/en/developers/apps/building-github-apps/rate-limits-for-github-apps).

The `snyk-scm-contributors-count` tool applies a strict limit of a maximum of 3 calls per second which amounts to 10,800 calls per hour, to deal with the rate limit and an additional regulating mechanism to deal with 429 responses ("too many calls").

## GitLab and GitLab Server

On GitLab, the API rate limit is 300 calls per minute for authenticated users as described in this [guide](https://docs.gitlab.com/ee/user/gitlab_com/index.html#gitlabcom-specific-rate-limits) and this [guide](https://docs.gitlab.com/ee/user/admin_area/settings/rate_limits_on_raw_endpoints.html) for GitLab Server.

The `snyk-scm-contributors-count` tool applies a strict limit of a maximum of 120 calls per minute to deal with the rate limit and an additional regulating mechanism to deal with 429 responses ("too many calls").

{% hint style="info" %}
On GitLab Server, the API rate control is configurable by the admin, as described in the [guide](https://docs.gitlab.com/ee/user/admin_area/settings/rate_limits_on_raw_endpoints.html).
{% endhint %}
