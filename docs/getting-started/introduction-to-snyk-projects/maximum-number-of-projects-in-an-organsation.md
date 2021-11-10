# Maximum number of projects in an organization

Depending on the type of plan you have with Snyk you will have a limit on the number of projects you can have in a single organization.

| Plan                   | Number of Projects |
| ---------------------- | ------------------ |
| Free                   | 10,000             |
| Team                   | 25,000             |
| Business               | 25,000             |
| Enterprise             | 25,000             |
| Pro (legacy plan)      | 25,000             |
| Standard (legacy plan) | 25,000             |

### How you will know when you have reached the limit?

When you reach the limit, Snyk stops importing more projects into the organization, and an error message appears:&#x20;

In the Snyk UI:

![](<../../.gitbook/assets/image (14) (1).png>)

In the Snyk CLI:

> This organization has 10,000 of the maximum 10,000 projects. You will not be able to import more projects: http://docs.snyk.io/getting-started/introduction-to-snyk-projects/maximum-number-of-projects-in-an-organization

The CLI will stop processing any new projects in your organization once the maximum number of projects has been reached. This may leave you with an impartial import. Once you clear space, by  removing the projects you on longer need in the organisation, you can repeat the command and the remaining projects will be imported to Snyk.

In the Snyk API:

```
    "data":{
        "code":400,
        "message":"This organization has 25000 of the maximum 25000 projects. You will not be able to import more projects: https://docs.snyk.io/getting-started/introduction-to-snyk-projects/maximum-number-of-projects-in-an-organsation",
        "errorRef":"5bc3fb50-cbcd-4c15-81f6-b183fc95d10f"
    },
```



This limit is in place to protect your experience of Snyk. There are no restrictions on the number of organizations you can create. If you are getting close to these limits, you can create more organizations  and split your projects across them.
