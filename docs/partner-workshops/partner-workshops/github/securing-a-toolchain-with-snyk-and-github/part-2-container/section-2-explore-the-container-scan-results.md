# Section 5: Explore the Container Scan results

In our Repo, we now have all the files necessary to package Goof into a container. We also introduced the Snyk Container scan into the workflow, to help identify issues introduced into our codebase by our decision to package the application into a container. Let's dig into the scan results before taking action.

## Step 1: Review the Container Scan results in GitHub

### GitHub Security Code Scanning

In the CI files we modified in the previous step, you might have noticed this section:

```text
      - name: Upload Container Scan results to GitHub Code Scanning
        uses: github/codeql-action/upload-sarif@v1
        with:
          sarif_file: snyk.sarif
```

This uploads the results from the Snyk Container scan into GitHub's Security Code Scanning panel, allowing you to view vulnerabilities present right within the GitHub UI. To view the results from the previous scan, head to Security -&gt; Code Scanning Alerts.

![](../../../../.gitbook/assets/gh-container-codescanning.png)

This gives you a fast initial glimpse into the risks introduced by your choice of container base image.

## Step 2: Review Container Scan results in the Snyk UI

When looking at risks introduced by a container base image, it's best to start by making sure you're using the best base image possible for your container. When using the Snyk CLI, the Docker CLI, or the Snyk UI, Snyk will recommend a more secure base image if one is available. 

### Import the Dockerfile into Snyk

In the Snyk UI, open the project created in Part 1. To keep everything neat, we'll add the `Dockerfile` into this Project. To do this, click the + sign on the Project entry, and provide the path to the `Dockerfile`. 

![](../../../../.gitbook/assets/snyk-container-adddockerfile.png)

Once it's imported, you should see the Dockerfile show up in the Project list. Click into it.

![](../../../../.gitbook/assets/snyk-container-dockerfile.png)

### Review Base Image Suggestions

In the Project view, you'll notice a new box that comes up specifically when scanning containers. This shows you how your current base image compares to other options that might contain less vulnerabilities.

These are grouped by how likely they are to be compatible with your application:

* `Minor` upgrades are the most likely to be compatible with little work, 
* `Major` upgrades can introduce breaking changes depending on image usage,
* `Alternative` architecture images are shown for more technical users to investigate.

![](../../../../.gitbook/assets/snyk-container-baseimagerecs.png)

In the next section, we'll take action and remedy some of the vulnerabilities in our application by upgrading the base image we're using to build our container image. Proceed to Section 3 when ready!

