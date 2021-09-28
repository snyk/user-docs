# CircleCI Project

Let's quickly review the preceding steps. If you skipped or missed any of these, you will want to go back and make sure you've completed them:

1. A valid Snyk, CircleCI, and AWS account
2. An Amazon EKS cluster provisioned
3. A `fork` or `clone` of the GitHub repository

### Setting up CircleCI

If this is your first time using CircleCI, I recommend reviewing the [step-by-step instructions](https://circleci.com/docs/2.0/getting-started/#setting-up-circleci) on setting up your project. You will need to navigate to the CircleCI [Project Page](https://app.circleci.com/projects/) and configure things so that you will automatically trigger a pipeline run against your Git repo.

Once configured, subsequent `commits` or `pull requests` will trigger runs as shown in the examples below:

#### Using git push

```bash
git push --force-with-lease origin develop
```

```text
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 16 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (5/5), 1.31 KiB | 1.31 MiB/s, done.
Total 5 (delta 2), reused 0 (delta 0)
remote: Resolving deltas: 100% (2/2), completed with 2 local objects.
To github.com:snyk-partners/snyk-circleci-eks.git
   4a09bea..9025276  develop -> develop
```

#### Pipeline is triggered

![](http://g.recordit.co/PgPkhEpdy1.gif)

#### Pipeline success

![](../../../.gitbook/assets/circleci_success.png)

