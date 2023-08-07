# Kicking off an import

Note that any logs will be generated at `SNYK_LOG_PATH` directory.

## 1. Create the `import-projects.json` file

The file is expected to have a **required** `targets` top level key which is an array of **import targets**.

```
{
  targets: [
    {..},
    {..}
  ],
}
```

Each **import target** has the following keys:

```
{
  // required
  "orgId": "<public_snyk_org_id>",
  "integrationId": <"public_snyk_integration_id>",
  "target": {..} // the identifier of where the projects can be found (for example branch, repo name and owner for Github)

   // optional
  "files": [ { path: "full/path/to/file1"} , { path: "full/path/to/file2" }],
  "exclusionGlobs": "fixtures, tests, __tests__, node_modules",
}
```

* `orgId` - Can be found in https://app.snyk.io/org/YOUR\_ORG/manage/settings
* `integrationId` - Can be found in the Integrations menu for each SCM https://app.snyk.io/org/YOUR\_ORG/manage/settings
* `target`, `files`, `exclusionGlobs` - See the Snyk [Import API documentation](https://snyk.docs.apiary.io/#reference/integrations/import-projects/import) for more information.
  * `exclusionGlobs` - Comma-separated list of up to ten folder names to exclude from scanning (each folder name must not exceed 100 characters). If not specified, defaults to "fixtures, tests, **tests**, node\_modules". If an empty string is provided, no folders will be excluded.
  * `files` - An object array. Each path must be the full relative path to the file from the root of the target. Only those files will be imported if found at that location.

**Note:** For a repo that may have 200+ manifest files it is recommended to split this import into multiple imports by targeting specific files. Importing hundreds of files at once from one repo can cause the import to result in some errors or failures.

Splitting the import to target some files or some folders only will benefit from the re-tries and yield a smaller load on the source control management system being used. Populate the `files` property to accomplish this in the import JSON.

If you have any tests or fixtures that should be ignored, please set the `exclusionGLobs` property:

> a comma-separated list of up to ten folder names to exclude from scanning. If not specified, it defaults to "fixtures, tests, **tests**, node\_modules". If an empty string is provided, no folders will be excluded

**Note: snyk-api-import supports all of the same integration types and project sources as identified in the** [**Import API documentation (Import Projects, Import)**](https://snyk.docs.apiary.io/#reference/import-projects)**. Examples follow. If there is no example for your use case, see the API documentation.**

### **Example: GitLab**

```
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "id": 13,
        "branch": "master"
      },
    },
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "id": 2,
        "branch": "master"
      },
    },
  ]
}
```

### **Example: Bitbucket Server**

```
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "repoSlug": "api-import-circle-test",
        "name": "Snyk api-import-circle-test",
        "projectKey": "PROJECT"
      },
      "files": [
        {
          "path": "package.json"
        },
        {
          "path": "package/package.json"
        },
        {
          "path": "Gemfile.lock"
        }
      ],
      "exclusionGlobs": "fixtures, test"
    },
  ]
}
```

### **Example: GitHub.com, GitHub Enterprise, dev.azure.com, Hosted Azure Repos**

```
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "name": "shallow-goof-policy",
        "owner": "api-import-circle-test",
        "branch": "master"
      },
      "exclusionGlobs": "fixtures, test"
    }
  ]
}
```

### **Example: Google Container Registry**

```
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "name": "projectId/repository:tag"
      },
    }
  ]
}
```

### **Example: Azure Container Registry, Elastic Container Registry, Artifactory Container Registry**

```
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "name": "repository:tag"
      },
    }
  ]
}
```

## 2. Set the env vars

* `SNYK_IMPORT_PATH`- the path to the import file or use `--file` parameter
* `SNYK_TOKEN` - your [Snyk api token](https://app.snyk.io/account)
* `SNYK_LOG_PATH` - the path to folder where all logs should be saved. Snyk recommends creating a dedicated logs folder for each import you have running. Note: all logs will be appended.
* `CONCURRENT_IMPORTS` (optional) - defaults to 15 repos at a time, which is the recommended amount to import at once as a maximum. Just one repo may have many projects inside which can trigger many files at once to be requested from the user's SCM instance and some may have rate limiting in place. This script aims to help reduce the risk of hitting a rate limit.
* `SNYK_API` (optional) defaults to `https://snyk.io/api/v1`

## 3. Download and run

Grab a binary from the [releases page](https://github.com/snyk-tech-services/snyk-api-import/releases) and run with `DEBUG=snyk* snyk-api-import-macos import --file=path/to/imported-targets.json`

### **Skip all previously imported targets**

This util can be used to skip previously imported targets (repos) so only remaining targets will be imported.

The util helps generate the `imported-targets.log` file by analyzing the projects already in a given Snyk Group. When present in the logging path, this file is used to look up targets that should be skipped during the import.

### Example

* All GitHub repos have been imported into Snyk into their respective organizations during initial onboarding.
* New GitHub repos have since been added and now need to be added to Snyk.
* To avoid importing everything again, you can use this util and run import again to import only "new" Github repos. This is much faster and removes unnecessary calls to Snyk and GitHub to fetch files and do the import for everything again.

### Importing the same target

* The same target imported into a different organization can be imported.
* The same target from a differed source can be imported, for example, the same repo is present in GitHub and is now is being imported via GitHub Enterprise into the same org.

### Command to run

* skip all previously imported into all orgs in a Group: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>`
* skip all previously imported for a specific organization: `snyk-api-import-macos list:imported --integrationType=<integration-type> --orgId=<snyk_org_id>`
* import a single integration / projects source: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>`
* import multiple integrations / projects sources: `snyk-api-import-macos list:imported --integrationType=<integration-type> --integrationType=<integration-type> --orgId=<snyk_org_id>`

### Supported integration types

* GitHub.com `github`
* GitHub Enterprise `github-enterprise`
* GitLab `gitlab`
* Bitbucket Cloud `bitbucket-cloud`
* Google Cloud Registry `gcr`
* DockerHub registry `docker-hub`
* Azure repos `azure-repos`
