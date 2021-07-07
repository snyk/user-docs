# Project import errors

While importing a project into the Snyk console, you may sometimes encounter an error; for example:

  
Some common errors and suggested resolutions are described below.

### Failed to process manifest FILE\_NAME

The manifest file is missing, has been moved or renamed. If this is intentional, please deactivate or delete this project in settings.

You may receive this error if you delete, move or rename the manifest file which Snyk reads during the initial project import.

To resolve this issue, restore the manifest file to the location as it was at during the initial project import. Or, if the file move was intentional, deactivate or delete the Snyk project, then re-import the project.

### Failed to parse package.json or package-lock.json as valid JSON

### Failed to parse package.json or yarn.lock as valid JSON

Please check the package.json and lockfile is valid JSON, any manual manipulation of these files may result in a missing/extra character where it is not expected.

Use any of the online tools to validate it is correct JSON. To test locally: **npm install** should pass without errors.

Note: Snyk supports Yarn v2 via the Snyk CLI only. More details can be found in our official [documentation](https://support.snyk.io/hc/en-us/articles/360004712477).

### Failed to access private module

Go Modules projects that depend on modules from private Git repositories are supported where the private repositories are in the same Git organisation as the main project repository.

Imports for projects with private modules from repos in other Git organisations will fail.

### Failed to process RubyGems lockfile contents

We were unable to process the Gemfile.lock to extract dependencies from it. Please contact support on support@snyk.io and if possible share the lockfile you used so we can identify the issue.

If you manually add a project to Snyk using **Add custom file location,** ensure the folder path and file name are in the correct case that matches your repo. 

For example, if you have on BBS **myproject / MyRepo** with a subfolder **GoofFolder**, adding manually **gooffolder/package.json** will fail; specify **GoofFolder/package.json** instead.

### Satisfying version not found for cousin dependencies

When importing a .Net project, Snyk creates the dependency tree according to the **.\*proj** file. According to the NuGet logic, you cannot use the same package with different versions in the same project.

For example, if Package A and Package B are direct dependencies of the same project, and both use Package C, then Package C is a transitive dependency. When resolving the version of Package C, it must be resolved into one version and to satisfy the requirements of Package A and Package B. Failing to do that results in the error message **Satisfying version not found for cousin dependencies**.

To troubleshoot the issue, please try to build the project. If it was built successfully, please reach out to [support@snyk.io](mailto:support@snyk.io) and we will look into it. If not, please make sure that all of the dependencies you use in your project are not conflicting. 

### Satisfying version not found

When importing a .Net project, Snyk creates the dependency tree according to the **packages.config** or **.\*proj** files. To resolve the version of all dependencies, they must be compatible with the target framework of the project. If one or more dependencies are not compatible, the error message appears **Satisfying version not found**.

To troubleshoot the issue, please try to build the project. If it was built successfully, please reach out to [support@snyk.io](mailto:support@snyk.io) and we will look into it. If it wasn't, please make sure that all of the dependencies you are using in your project are compatible with the target framework of the project. 

### Target framework not present in &lt;project\_FileName&gt;

To build a dependency tree, Snyk needs to determine the target framework of the project. Currently, we require the target framework to be part of the .**\*proj** file. 

### Failed to detect OS release

This error may appear while testing/importing a Docker image, if you have a Scratch based Container Image**.** This is a current known limitation for Scratch based container images.  
As a workaround, you can use the CLI to import/test these docker images.

See [Test your container images with our CLI tool](https://support.snyk.io/hc/en-us/articles/360003946917-Test-your-container-images-with-our-CLI-tool) for details of how to run a **Snyk test in** the CLI.

Also see [What are docker scratch based images](https://support.snyk.io/hc/en-us/articles/360004012857-What-are-docker-scratch-based-images-).

For any other reasons please do reach out to [support@snyk.io](mailto:support@snyk.io)

### Image size exceeds allowed size

When importing an image from AWS ECR you can come up against the following error when Importing large AWS ECR images:

**Image size exceeds allowed size \(4.XXXGb &gt; 2.000Gb\) when importing container from AWS ECR**

Unfortunately, there is currently a limit of 2gb for AWS ECR images.

## Infrastructure as Code \(Kubernetes / Helm\)

### Failed to download Helm chart file

This error may be caused by a failure to download the `Chart.yaml` file from your git repository. This might mean Snyk does not have the correct permissions to access the file or that it has been deleted.

If you are using a Snyk broker, make sure you follow the relevant [configuration instructions](https://support.snyk.io/hc/en-us/articles/360010797537-Detecting-Kubernetes-configuration-files-using-a-broker).

### Failed to find a proper values file

Snyk needs the `values.yaml` file to render the template. This file needs to be in the same folder as the `Charts.yaml` file.

### Failed to receive the content of the templates folder

### Failed to get the file: &lt;FileName&gt; in order to template the chart

Ensure there is a `templates` folder in the same directory as the `Chart.yaml` file. This folder should contain templates and helper files \(`*.tpl`\) that we access as part of the import flow. For more details about the expected format of Helm Charts, see [Scan and fix security issues in your Helm Charts](https://support.snyk.io/hc/en-us/articles/360007673117-Scan-and-fix-security-issues-in-your-Helm-Charts).

If you use a Snyk broker, follow the relevant [configuration instructions](https://support.snyk.io/hc/en-us/articles/360010797537-Detecting-Kubernetes-configuration-files-using-a-broker.).

### Failed to run Helm template using the chart

You can check whether your chart is valid by testing locally, by running `helm template`, which should pass without errors.

Snyk does not currently support Helm charts which use an internal Helm repository to fetch dependencies, or that require additional values files to render the Chart.

## Infrastructure as Code \(Terraform\)

### Unsupported Terraform file format

This error appears if the imported repository contains HCL version 1 files. Snyk only supports HCL version 2, used from Terraform version 0.12+.

Use the `terraform validate` command to check your configuration file version.

### Go Template placeholders found in Terraform File

We detected the usage of Go Template syntax \(e.g. {{...}}\) in this file, we do not support this syntax in Terraform files.

## Python

Problem with **--no-binary** flag. Use the following instead:

```text
psycopg2==2.7.5 --no-binary=psycopg2
```

Even though pip does not require "=" between the flag and the value, Snyk requires the character in order to prevent parsing failures.

