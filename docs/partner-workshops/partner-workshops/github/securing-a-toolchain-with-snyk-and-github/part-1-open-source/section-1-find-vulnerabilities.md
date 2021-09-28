---
description: >-
  Let's start by importing the forked Goof repo into your Snyk account to scan
  it for vulnerabilities.
---

# Section 1: Find Vulnerabilities

## Step 1: Configure Snyk's GitHub Integration

{% hint style="info" %}
If you've already configured the Snyk GitHub integration, continue to Step 2.
{% endhint %}

First we need to connect Snyk to GitHub so we can import our Repository. Do so by:

1. Logging in to Snyk.io. [Sign up](https://snyk.co/SnykGH) if you haven't already.
2. Navigating to Integrations -&gt; Source Control -&gt; GitHub
3. Fill in your Account Credentials to Connect your GitHub Account.

![](../../../../.gitbook/assets/snyk-gh.png)

## Step 2: Import the forked Goof Repo into Snyk

Now that Snyk is connected to your GitHub Account, import the Repo into Snyk as a Project.

1. Navigate to Projects
2. Click "Add Project" then select "GitHub"
3. Click on the Repo you forked.

![](../../../../.gitbook/assets/snyk-ghimport.png)

## Step 3: Explore your Repo's risks

When the Repo imports, Snyk has found the `package.json` file where open source components for our Goof application are declared. We can see that they contain 60 vulnerabilities, including 31 High Severity ones! 😳

![](../../../../.gitbook/assets/snyk-projvulns.png)

Before we start fixing Vulnerabilities, remember that our `PROD` branch contains the Production-ready version of the code. In the next section, we'll implement a Snyk Gate designed to "stop the bleeding", preventing any new vulnerabilities from making it into that branch.

