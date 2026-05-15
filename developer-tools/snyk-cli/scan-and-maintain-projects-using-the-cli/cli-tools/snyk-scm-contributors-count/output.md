---
description: The output result of running the SCM-Contributors-Count tool
---

# Output

## Summary

The Summary section appears at both the beginning and end of the output, for example:

```
#### Summary
Private Repos Contributors Count:: 1
Public Repos Contributors Count: 1
Total Unique Contributors Count for Private and Public repositories: 1
Private Repository Count: 1
Public Repository Count: 1
Total Repository Count: 2
Exclusion Count: 1
```

* `Private Repos Contributors Count` - The number of unique contributors for the private repos that were found or provided.
* `Public Repos Contributors Count` - The number of unique contributors for the public repos that were found or provided.
* `Total Unique Contributors Count for Private and Public repositories`- The total number of unique contributors across all probed repositories.
* `Private Repository Count` - Number of private repos that were scanned.
* `Public Repository Count` - Number of public repos that were scanned.
* `Total Repository Count` - The number of the total repos that were scanned (public and private).
* `Exclusion Count` - The number of contributors that were not counted according to the exclusion file that was provided.

## Details

```
### Details:
## Repository List

# Private Repositories:
someOrganization/someRepository(Private)

# Public Repositories:
anotherOrganization/anotherRepository(Public)
```

* `Private Repositories` - A list of the private repositories that were scanned.
* `Public Repositories` - A list of the public repositories that were scanned.

{% hint style="info" %}
Next to each repository name, there is an indication of its visibility, either (Private) or (Public).
{% endhint %}

## Contributors details

```
## Contributors details
[
    [
        "someUser",
        {
            "email": "someUser@company.io",
            "contributionsCount": 15,
            "reposContributedTo": [
                "someOrganization/someRepository(Private)"
                "anotherOrganization/anotherRepository(Public)"
            ]
        }
    ],
    [
        "anotheruser",
        {
            "email": "anotherUser@company.io",
            "contributionsCount": 11,
            "reposContributedTo": [
                "someOrganization/someRepository(Private)"
            ]
        }
    ],
    [
        "anotheruser(duplicate)",
        {
            "email": "anotherUser@otherCompany.com",
            "contributionsCount": 11,
            "reposContributedTo": [
                "someOrganization/someRepository(Private)"
            ]
        }
    ]
]
```

* `email` - The email of the contributor
* `contributionsCount` - The number of the times this contributor has contributed to the repo/s.
* `repoContributedTo` - A list of the repo/repos to which this contributor has contributed.
* `(duplicate)` - Indicator that the same user has been detected from different email addresses; note that they will be counted as **different committers**.

{% hint style="info" %}
As the output can be long, Snyk recommends sending the output to a file for better review and parsing options.
{% endhint %}
