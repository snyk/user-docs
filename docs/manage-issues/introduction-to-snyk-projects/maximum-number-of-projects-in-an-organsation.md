# Maximum number of Projects in an Organization

## Introduction

The number of Projects you can have in a single Organization depends on the type of [Snyk Pricing Plan](../../more-info/plans.md) you are using:

| Plan       | Number of Projects |
| ---------- | ------------------ |
| Free       | 10,000             |
| Team       | 25,000             |
| Enterprise | 25,000             |

When you reach that limit, Snyk stops importing more Projects into the Organization.

### **Notification of reaching the limit**

You will know you have hit this limit when you see the following

#### Snyk Web UI

This warning banner appears:

<figure><img src="../../.gitbook/assets/Maximum number of projects.png" alt=".Banner at top of Projects listing showing how many Projects you have over the limit"><figcaption><p>Banner at top of Projects listing showing how many Projects you have over the limit</p></figcaption></figure>

#### Snyk CLI

In the CLI, the `snyk monitor` command returns the following error

> `Maximum number of projects reached for this organization. You cannot import more projects.`

#### Snyk API

In the API, an import request returns this error:

```
"data":{
        "code":400,
        "message":"This organization has 25000 of the maximum 25000 projects. You will not be able to import more projects: https://docs.snyk.io/getting-started/introduction-to-snyk-projects/maximum-number-of-projects-in-an-organsation",
        "errorRef":"5bc3fb50-cbcd-4c15-81f6-b183fc95d10f"
    },
```

### Actions to take

This limit is in place to protect your experience with Snyk. There are no restrictions on the number of Organizations you can create.&#x20;

If you are getting close to these limits, you can create more Organizations and split your Projects across them.
