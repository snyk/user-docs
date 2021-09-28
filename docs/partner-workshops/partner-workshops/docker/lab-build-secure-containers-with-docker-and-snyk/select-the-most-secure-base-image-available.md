---
description: >-
  One of the easiest ways we can secure containers is to ensure we're using a
  secure base image.
---

# Select the most secure base image available

Snyk makes selecting a secure base image for your container easy. We'll show you how to get base image suggestions in both the Snyk UI, and using Docker Desktop. 

## Base image suggestions in Docker Desktop

Developers can `docker scan` containers to get vulnerability information and base image upgrade guidance. Scan the image by running the following command. Passing the `Dockerfile` used to build the image using `--file` is needed to receive base image suggestions.

```text
docker scan $DockerId/goof:dev --file=Dockerfile
```

{% hint style="info" %}
Learn what else is possible with Docker Scan with [our handy cheat sheet](https://snyk.io/blog/docker-security-scanning-cheatsheet-2021/).
{% endhint %}

Snyk recommends less vulnerable base images grouped by how likely they are to be compatible:

* `Minor` upgrades are the most likely to be compatible with little work, 
* `Major` upgrades can introduce breaking changes depending on image usage,
* `Alternative` architecture images are shown for more technical users to investigate.

![](../../../.gitbook/assets/docker-scanvulns.png)

Next we'll show you how you can get these same base image suggestions inside the Snyk UI.

## Base image suggestions in Snyk

### Step 1: Configure Snyk's GitHub Integration

{% hint style="info" %}
If you've already configured the Snyk GitHub integration, continue to Step 2.
{% endhint %}

First we need to connect Snyk to GitHub so we can import our Repository. Do so by:

1. Logging in to Snyk.io. [Sign up](https://snyk.co/SnykGH) if you haven't already.
2. Navigating to Integrations -&gt; Source Control -&gt; GitHub
3. Fill in your Account Credentials to Connect your GitHub Account.

![](../../../.gitbook/assets/snyk-gh.png)

### Step 2: Import the forked Repo into Snyk

Now that Snyk is connected to your GitHub Account, import the Repo into Snyk as a Project.

1. Navigate to Projects
2. Click "Add Project" then select "GitHub"
3. Click on the Repo you created.

![](../../../.gitbook/assets/snyk-ghimport.png)

### Step 3: Explore Base Image suggestions for your app

When the Repo imports, Snyk shows the supported manifest files in the repo, including our container's `Dockerfile`. 

![](../../../.gitbook/assets/snyk-proj-dockerfile.png)

Clicking the Dockerfile brings you to the project view, where you can see the base image suggestions. 

![](../../../.gitbook/assets/snyk-baseimagerecs.png)

Up next, we'll scan the containers running in our Kubernetes cluster for vulnerabilities and base image upgrade guidance. This use case is helpful when running images not developed in-house. 

## Optional: Finding vulnerabilities in running workloads

Snyk's Kubernetes Monitor allows you to find vulnerabilities in workloads running in your cluster. In this section we'll deploy the Snyk Kubernetes Monitor to scan the goof application.

{% hint style="info" %}
This section is optional because the Kubernetes Monitor requires a Snyk Standard plan or higher. Contact us via Intercom if you're a free user and want to evaluate this functionality.
{% endhint %}

{% hint style="info" %}
The full [instructions for deploying the Kubernetes Monitor](https://support.snyk.io/hc/en-us/articles/360003916158#UUID-753328ea-3d73-0eeb-4301-c22522273797) can be found in our Docs.
{% endhint %}

To get started, first create a namespace for the Snyk Monitor:

```text
kubectl create ns snyk-monitor
```

Next, retrieve your Kubernetes Integration ID from Snyk's Integrations menu.

![](https://support.snyk.io/hc/article_attachments/360007147458/uuid-26f9c2cd-2755-07d5-61a0-bdb0261d87ab-en.gif)

Save the integration ID as an Environment Variable to copy-paste the next command.

```text
IntegrationId=<<your_integration_id>>
```

Create a secret in the cluster containing the Integration ID.

```text
kubectl create secret generic snyk-monitor -n snyk-monitor \
        --from-literal=dockercfg.json={} \
        --from-literal=integrationId=$IntegrationId 
```

Now deploy the Kubernetes monitor using Helm:

```text
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Docker Desktop" 
```

When it finishes deploying, you'll be able to import the cluster workloads into Snyk.

![](../../../.gitbook/assets/k8s-import.png)

Select to import both the `goof` and `mongo` deployments.

![](../../../.gitbook/assets/k8simport2.png)

When the import completes, they will show up next to the other projects in the Snyk dashboard.

![](../../../.gitbook/assets/k8sworkloads.png)

Since it's the same container, clicking the `snyk-docker/deployment.apps/goof` project will show the same vulnerabilities as the `Dockerfile` project from the previous section. 

You can point the Kubernetes project at the `Dockerfile` to see base image suggestions. 

![](../../../.gitbook/assets/k8ssettings.png)

This will provide the same base image suggestions for running workloads. We can see that using a different base image can address hundreds of vulnerabilities. This is a relatively low effort fix with huge payoff. Let's apply a more secure base for our application.

## Choose a more secure base image

Open the `Dockerfile` with a text editor, and replace, or comment out, the old base image with a new one. In this example, we'll use `node:10.23.1`.

{% code title="Dockerfile" %}
```bash
FROM node:10.23.1

RUN mkdir /usr/src/goof
RUN mkdir /tmp/extracted_files
COPY . /usr/src/goof
WORKDIR /usr/src/goof

RUN npm update
RUN npm install
EXPOSE 3001
EXPOSE 9229
ENTRYPOINT ["npm", "start"]
```
{% endcode %}

Save the changes. Now build and push the container to Docker Hub.

```bash
docker build -t $DockerId/goof:dev .
docker push $DockerId/goof:dev
```

Once it's done, test it in Kubernetes by scaling the `goof` deployment with `kubectl`. The deployment's `ImagePullPolicy` forces Kubernetes to pull the latest image from Docker Hub.

```bash
kubectl scale deployment goof --replicas=0
kubectl scale deployment goof --replicas=1
```

Navigate to [http://localhost:3001](http://localhost:3001) to test the app. Success! It likes the new base image. In Docker Hub, we can appreciate the `dev` tag has less vulnerabilities than `PROD`.  

![](../../../.gitbook/assets/dockerhub-tagvulns%20%281%29.png)

If you imported the project using the Kubernetes Monitor, you'll see the new results once the Snyk Monitor re-tests the project.

![](../../../.gitbook/assets/k8sproject.png)

Let's get these changes into our GitHub Repo's PROD branch!

## Commit the changes to GitHub, then into PROD

### Commit the changes to the develop branch

First we need to commit the changes to our Repo's `develop` branch. Use Git to do so.

```bash
git add Dockerfile
git commit -m "new base image"
git push
```

### Open a Pull Request from develop to PROD

In GitHub, create a Pull Request from the `develop` branch to the `PROD` branch.

![](../../../.gitbook/assets/gh-newpr.png)

![](../../../.gitbook/assets/gh-pr-newbase.png)

### See the PR checks we configured in action

The workflows we configured earlier will run as part of this PR. You can see the results of the workflow runs in the PR view before choosing to Merge the PR. To recap what each one does:

* **license/snyk** checks the incoming changes against the snapshot in Snyk for new license issues
* **security/snyk** checks the incoming changes against the snapshot in Snyk for new vulnerabilities
* **Check for Open Source Vulnerabilities with Snyk** checks if there are **any** vulnerable open source components in the application with fixes available
* **CI Task for PROD Branch** rebuilds and app and container to make sure it correctly builds. 
* **Code scanning results** pushes the container scan results to GitHub Security for consumption.

![](../../../.gitbook/assets/gh-prchecks%20%281%29.png)

Even though we have open source vulnerabilities, merge the PR. This triggers AutoBuild to build and re-scan the PROD container. When it completes, we see a reduction in the number of vulnerabilities.

![](../../../.gitbook/assets/dockerhub-tags.png)

Choosing a more secure base image was just the start! In the next section, we'll address the application's vulnerable application dependencies.

