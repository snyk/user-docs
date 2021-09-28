# Section 4: Containerize your app

## Step 1: Create a new branch for the container files

Let's start by creating a new branch from our `develop` branch, where we'll create the files for this section before merging them to `develop`. Call it `add-container-files` .

![](../../../../.gitbook/assets/gh-container-newbranch.png)

## Step 2: Create the files for the Container workflow

{% hint style="info" %}
In this step you'll create a Dockerfile and modify the GitHub Actions CI Workflows. The files are available in the `container-actions` branch if you want to copy-paste from there. These are there for reference; attempting to merge that branch to `develop` will cause Merge Conflicts.  
{% endhint %}

### Create a Dockerfile

The first file we'll create is the `Dockerfile` , the instructions to package the application into a container. Ensure you're on the new branch, and click on "Add File", then "Create new file".

![](../../../../.gitbook/assets/gh-container-createnewfile.png)

Call the new file `Dockerfile` as shown below.

![](../../../../.gitbook/assets/gh-container-createdockerfile.png)

Paste this in as the contents of the `Dockerfile`

```text
FROM node:6-stretch

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

When ready, commit the changes directly to the `add-container-files` branch. 

![](../../../../.gitbook/assets/gh-container-createdockerfile2.png)

### Modify the CI workflows

Now, we need to tell GitHub Actions how to build our container. Recall that the `.github/workflows` folder contains the CI workflows that re-build our application when Pull Requests are opened against our branches. Let's add the Container build steps to both CI workflows using the GitHub Editor.

![](../../../../.gitbook/assets/gh-container-editdevelopci.png)

Replace the contents of `.github/workflows/goof-ci-develop` with the following:

```text
# This workflow will do a clean install of node dependencies, build the source code and run tests across different versions of node
# For more information see: https://help.github.com/actions/language-and-framework-guides/using-nodejs-with-github-actions

name: Node.js CI task for develop branch

on:
  push:
    branches: [ develop ]
  pull_request:
    branches: [ develop ]

jobs:
  build_app:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Use Node.js 12.x
      uses: actions/setup-node@v1
      with:
        node-version: 12.x
    - run: npm ci
    - run: npm run build --if-present
  build_container:
    needs: [build_app]
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Setup up Docker Buildx
        uses: docker/setup-buildx-action@v1
      - name: Build Docker Image
        id: docker_build
        uses: docker/build-push-action@v2
        with:
          push: false
          load: true
          tags: goof:develop         
      - name: Snyk Container Test
        continue-on-error: true
        uses: snyk/actions/docker@master
        env:
          SNYK_TOKEN: ${{ Secrets.SNYK_TOKEN }}
        with:
          image: goof:develop
          args: --file=Dockerfile
      - name: Upload Container Scan results to GitHub Code Scanning
        uses: github/codeql-action/upload-sarif@v1
        with:
          sarif_file: snyk.sarif
```

Our CI workflow will now take the following actions:

1. First, it will build the app. This is the `build-app` job. 
2. If the app build succeeds, it will build the container.
3. If the container builds successfully, it will then scan it with Snyk container. 
4. When the scan is done, results are pushed to GitHub Security Code Scanning.

Now commit the changes directly to our new `add-container-files` branch. 

![](../../../../.gitbook/assets/gh-container-commitcontainerci.png)

Now, repeat the same process for the `PROD` workflow, replacing the image tag specified in lines 36 and 43 from `goof:develop` to `goof:PROD`. This will ensure the `PROD` and `develop` containers remain separate entities once uploaded to a container registry.

## Step 3: Create a Pull Request into develop

With the Dockerfile created and workflows modified, we'll merge the `add-container-files` branch into the `develop` branch to make it official. In GitHub, initiative a Pull Request. Like in the previous section, remember to select your fork as the Base repo.

![](../../../../.gitbook/assets/gh-container-addfilespr.png)

Once the PR is created, the checks will run, including the new one we added to build the container image. Once all checks complete, merge the Pull Request and delete the branch. 

![](../../../../.gitbook/assets/gh-container-addfileprchecks.png)

We have now packaged our application into a container! Even through we fixed all open and fixable issues with the Open Source components in our application, our decision to package it into a container introduced additional risks that must be addressed. Head on to Section 2 to see how Snyk Container helps you find and fix those issues. 

