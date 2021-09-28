# Section 7: Create Deployment Manifests

## Step 1: Create a new branch for the container files

Let's start by creating a new branch from our `develop` branch, where we'll create the deployment manifests for our application before merging them to `develop`. Call it `add-iac-files` .

![](../../../../.gitbook/assets/gh-iac-createbranch.png)

## Step 2: Create the application manifests

{% hint style="warning" %}
In this step we create the YAML files with the deployment config for our application. The files in the `iac-actions` branch if you want to copy-paste from there. These are for reference; attempting to merge that branch to `develop` will cause Merge Conflicts.  
{% endhint %}

### Create a Deployment definition

First we create the `Deployment` , the configuration file that tells orchestration environments how to spin up the Goof container. Switch to the new branch, click on "Add File", then "Create new file".

![](../../../../.gitbook/assets/gh-iac-createfile.png)

Call the new file `goof-deployment.yaml` and paste this in as the contents of the file. 

```text
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goof
spec:
  replicas: 1
  selector:
    matchLabels:
      app: goof
      tier: frontend
  template:
    metadata:
      labels:
        app: goof
        tier: frontend
    spec:
      containers:
        - name: goof-app
          image: goof:PROD
          ports:
            - containerPort: 3001
            - containerPort: 9229
          env:
            - name: DOCKER
              value: "1"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goof-mongo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: goof
      tier: backend
  template:
    metadata:
      labels:
        app: goof
        tier: backend
    spec:
      containers:
        - name: goof-mongo
          image: mongo
          ports:
            - containerPort: 27017
```

When ready, commit the changes directly to the `add-iac-files` branch. 

![](../../../../.gitbook/assets/gh-iac-commitdeployment.png)

### Create the Service definition

Next, we create a `Service` to define how our application is exposed. In the new branch, add a file, call it `goof-service.yaml` , and paste this in as the contents. 

```text
apiVersion: v1
kind: Service
metadata:
  name: goof
spec:
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 3001
      name: "http"
    - protocol: TCP
      port: 9229
      targetPort: 9229
      name: "debug"
  selector:
    app: goof
    tier: frontend
---
apiVersion: v1
kind: Service
metadata:
  name: goof-mongo
spec:
  ports:
    - protocol: TCP
      port: 27017
      targetPort: 27017
      name: "mongo"
  selector:
    app: goof
    tier: backend

```

Commit the file directly to the `add-iac-files` branch.

![](../../../../.gitbook/assets/gh-iac-commitservice.png)

## Step 3: Create a Snyk IaC Test workflow

To ensure our manifest files are checked for configuration issues, we'll create a workflow that checks them with Snyk IaC. Create a file called `snyk-iac-check.yaml` in the `.github/workflows` folder and paste in the following as its contents:

```text
name: Test Infrastructure as Code with Snyk
on:
  push:
    paths: 
      - '**.yaml'
jobs:
  snyk:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Check configuration files for security issues with Snyk
        continue-on-error: true
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          file: goof-deployment.yaml goof-service.yaml
      - name: Upload result to GitHub Code Scanning
        uses: github/codeql-action/upload-sarif@v1
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
The `files` section tells GitHub Actions to run the workflow when changes are made to the YAML files we created. After it runs, results will be pushed to GitHub Security Code Scanning. For any GitHub activity not relating to these files, nothing happens.
{% endhint %}

Commit this file directly to the `add-iac-files` branch. 

## Step 4: Add IaC checks to the Snyk Gate

In Part 1, we established that our `PROD` branch contains the deploy-ready state of our code. With GitHub Actions, it's possible to create workflows to continuously deploy our application to an orchestration environment. For example, there are Actions to:

* [Deploy to Azure Kubernetes ](https://github.com/marketplace/actions/deploy-to-kubernetes-cluster)
* [Deploy to DigitalOcean Kubernetes](https://github.com/do-community/example-doctl-action)
* [Set up the `oc` tool to deploy to OpenShift](https://github.com/marketplace/actions/openshift-client-installer)

These actions take the manifest files we created as input, and can re-deploy the application when code is pushed into the `PROD` branch. We will configure our Snyk gate to run during the Pull Request process, to check our deployment manifests for configuration issues and catch issues before they are introduced into our Production environments in those Continuous Deployment scenarios. 

### Edit the Snyk Gate

Our Snyk Gate is located in `.github/workflows/snyk-gate.yml` . While in the `add-iac-files` branch,  open it with the GitHub web editor. 

![](../../../../.gitbook/assets/gh-iac-editgate.png)

Replace the contents of the file with the following. Notice the new `iac-security` job. In our job definition, we've decided to fail the check if any medium severity or higher issues are present. 

```text
name: Snyk Security Gate
on: 
  pull_request:
    branches:
      PROD
jobs:
  oss-security:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: Check for High Severity OSS Vulnerabilities
      uses: snyk/actions/node@master
      env:
        SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
      with:
        args: --severity-threshold=high --fail-on=upgradable
  iac-security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Snyk IaC Test
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: ${{ Secrets.SNYK_TOKEN }}
        with:
          file: goof-deployment.yaml goof-service.yaml
          args: --severity-threshold=medium
```

When ready, commit the changed file directly to the `add-iac-files` branch. 

![](../../../../.gitbook/assets/gh-iac-commitgatechanges.png)

### Understand the new Snyk Gate workflow

The Snyk Gate now checks for two things when a Pull Request is opened against the PROD branch:

* Are there application dependencies with High Severity issues that are fixable?
* Are there Medium severity configuration issues in our deployment manifests?

These are tolerances our team settled on for our Production environment. Automating this check as part of the PR process gives us confidence that what we merge into PROD conforms to these rules. 

## Step 5: Create a Pull Request into develop

We're ready to merge the `add-iac-files` branch into the `develop` branch. In GitHub, initiate a Pull Request from `develop` to `PROD`. Remember to select your fork as the Base repo.

![](../../../../.gitbook/assets/gh-iac-developpr.png)

Once the PR is created, we can see the new Snyk IaC check workflow run as part of the Pull Request, which uploads the results into GitHub Security Code Scanning. When it completes, merge the changes.  

## Step 6: Create a Pull Request into PROD

Our `develop` branch now has what we need to deploy our application into Kubernetes! Let's open a Pull Request to merge our changes, and our deployment manifests, into `PROD`.

![](../../../../.gitbook/assets/gh-iac-prodpr.png)

In the PR Checks, we see that our IaC gate failed!  

![](../../../../.gitbook/assets/gh-iac-iacgatefail.png)

Leave the Pull Request open for now. In the next section, we'll more about the risks Snyk identified and prevented us from introducing into our Production branch.

