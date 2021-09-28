---
description: >-
  To ensure we don't introduce more High Severity risks into our
  Production-ready branch, let's add a Snyk Security Gate to our Repo's GitHub
  Actions workflows.
---

# Section 2: Stop the Bleeding

In this section, we introduce a Security Gate using the [Snyk GitHub Actions](https://github.com/snyk/actions). The Snyk Action can fail GitHub checks, alerting you when Repo activity introduces risks, helping you catch issues early.

### Step 1: Create a Snyk Token

To use the Snyk Action, we need to create a Snyk API Token and store it as a GitHub Secret. 

#### Retrieve your Snyk Token

You can find your API Token one of two ways:

* If you have the Snyk CLI, retrieve it by running `snyk config get api`
* In the Snyk UI, head to your account [Settings Page](https://app.snyk.io/account) and retrieve it.

![](../../../../.gitbook/assets/snyk-token.png)

{% hint style="info" %}
Stuck? Check out [Revoking and regenerating Snyk API Tokens](https://support.snyk.io/hc/en-us/articles/360004008278-Revoking-and-regenerating-Snyk-API-tokens)
{% endhint %}

#### Store the Snyk Token in GitHub Secrets

Store the Token in the Forked Repo's secrets by navigating to Settings -&gt; Secrets -&gt; New Repository Secret. Name the Secret `SNYK_TOKEN`

![](../../../../.gitbook/assets/gh-secrets.png)

{% hint style="info" %}
Stuck? Check out [Creating Encrypted Secrets for a Repository](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets#creating-encrypted-secrets-for-a-repository)
{% endhint %}

Continue to Step 2 once your Snyk Token is saved.

### Step 2: Add the Snyk GitHub Action

Time to add a Snyk Security Gate to our workflow! The necessary files have already been created for you in the `oss-actions` branch. 

#### Inspect the Snyk Gate YML file 

Switch to the `oss-actions` branch, and navigate to the `.github/workflows` folder to see `snyk-gate.yml`. Note the following:

* Since this is designed to "stop the bleeding" of new vulnerabilities being introduced into `PROD`, this workflow only runs if a PR is opened against `PROD`. 
* The `--severity.threshold` and `--fail-on` arguments on the Snyk Action tell Snyk to only fail if any `high` severity risks that are `upgradable` \(a fix is available\) are present. 

{% hint style="info" %}
Learn more about these, and other CLI commands, in [Snyk CLI Reference](https://support.snyk.io/hc/en-us/articles/360003812578-CLI-reference).
{% endhint %}

![](../../../../.gitbook/assets/gh-snykgate.png)

#### Add the Snyk Gate to the develop branch

We'll create a Pull Request to add `snyk-gate.yml` to the `develop` branch. To open a Pull Request, first navigate to Pull Requests -&gt; New Pull Request.

![](../../../../.gitbook/assets/gh-newpr.png)

{% hint style="danger" %}
GitHub defaults to the upstream Snyk-Partners repo as the Base Repository as shown in the image below. Be sure to change that to your fork.
{% endhint %}

![Your PR should not look like this!](../../../../.gitbook/assets/gh-prcompare.png)

In the **Comparing Changes** window, select the `develop` branch from your forked repo as the base repository, and the `oss-actions` as the head repository, then open the PR.

![Ensure your fork is selected as the base repo before opening the PR.](../../../../.gitbook/assets/gh-oss-pr.png)

You should see that this PR adds the `snyk-gate.yml` file to the `develop` branch.

![](../../../../.gitbook/assets/gh-oss-pr-1-.png)

After input a name, description, and click **Create Pull Request** you'll be asked to **Merge Pull Request.**

![](../../../../.gitbook/assets/gh-mergepr.png)

### Step 3: Test the Security Gate

To test this, open a New Pull Request. This time, set `PROD` as the Base, and `develop` as the Head. 

![](../../../../.gitbook/assets/gh-mainpr.png)

In the Pull Request details, the Snyk Security Gate fails as expected because `develop` still contains the High Severity Vulnerabilities identified earlier.

![](../../../../.gitbook/assets/gh-snykgateworks.png)

You'll also notice two Snyk checks, `license/snyk` and `security/snyk`. These are making sure no **new** security or license issues are introduced. These **incremental** checks pass because we haven't introduced any risks that weren't already present. 

{% hint style="info" %}
Learn more about this functionality by reading [Snyk checks on Pull Requests](https://support.snyk.io/hc/en-us/articles/360006581938-Snyk-checks-on-pull-requests).
{% endhint %}

You now have a Snyk Security Gate! While you can merge at your discretion, Snyk fails this check to alert that there are unresolved, **Fixable,** **High** **Severity** risks. Leave the Pull Request open for now; in the next section, we'll secure our `develop` branch, clear this gate, and secure the PROD-ready version of our code!

