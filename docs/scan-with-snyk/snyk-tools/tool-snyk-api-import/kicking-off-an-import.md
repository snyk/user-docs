# Kicking off an import

`snyk-api-import` supports the same Project sources that you can import using the Snyk API: Git repositories, Docker images, containers, configuration files and much more. You can configure integrations using the Integrations settings on your Snyk Organization settings page. For more information, see the definition of [Target](../../../snyk-platform-administration/snyk-projects/#target) on the Snyk Projects documentation page.

Note that any logs will be generated at `SNYK_LOG_PATH` directory.

The steps to start an import follow.

## Create the `import-projects.json` file

The file is expected to have a required `targets` top-level key, which is an array of import targets.

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

* `orgId` - Can be found on your Organization settings page.
* `integrationId` - Can be found in the Integrations menu for each SCM on your Organization settings page.
* `target`, `files`, `exclusionGlobs` - See the [Snyk Import API documentation](../../../snyk-api/reference/import-projects-v1.md) for more information.
  * `exclusionGlobs` - Comma-separated list of up to ten folder names to exclude from scanning (each folder name must not exceed 100 characters). If not specified, defaults to "fixtures, tests, tests, node\_modules". If an empty string is provided, no folders will be excluded.
  * `files` - An object array. Each path must be the full relative path to the file from the root of the Target. Only those files found at that location will be imported.

Note that for a repository that may have 200+ manifest files, Snyk recommends that you split the import into multiple imports by targeting specific files. Importing hundreds of files at once from one repository can cause the import to result in some errors or failures.

Splitting the import to import some files or some folders only will benefit from the re-tries and produce a smaller load on the source control management system being used. Populate the `files` property to accomplish this in the import JSON.

If you have any tests or fixtures that should be ignored, set the `exclusionGLobs` property:

> a comma-separated list of up to ten folder names to exclude from scanning. If not specified, it defaults to "fixtures, tests, tests, node\_modules". If an empty string is provided, no folders will be excluded

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

## Set the environment variables

* `SNYK_IMPORT_PATH`- the path to the import file or use `--file` parameter
* `SNYK_TOKEN` - your [Snyk api token](https://app.snyk.io/account)
* `SNYK_LOG_PATH` - the path to the folder where all logs should be saved. Snyk recommends creating a dedicated logs folder for each import you have running. Note: all logs will be appended.
* `CONCURRENT_IMPORTS` (optional) - defaults to 15 repositories at a time, which is the recommended maximum to import at once. One repository alone may have many Projects inside, which can trigger a request for many files at once from the user's SCM instance. Some may have rate limiting in place. This script aims to help reduce the risk of hitting a rate limit.
* `SNYK_API` (optional) defaults to `https://api.snyk.io/v1`

## Download and run

Download a binary from the [releases page](https://github.com/snyk/snyk-api-import/releases) and run it with `DEBUG=snyk* snyk-api-import-macos import --file=path/to/imported-targets.json`

### **Skip all previously imported targets**

This utility can be used to skip previously imported Targets (repos), so only remaining Targets will be imported.

The utility helps generate the `imported-targets.log` file by analyzing the Projects already in a given Snyk Group. When present in the logging path, this file is used to look up Targets that should be skipped during the import.

### Example

* All GitHub repositories have been imported into Snyk into their respective Organizations during initial onboarding.
* New GitHub repositories have since been added and now need to be added to Snyk.
* To avoid importing everything again, you can use this utility and run `import` again to import only "new" GitHub repositories. This is much faster and removes unnecessary calls to Snyk and GitHub to fetch files and do the import for everything again.

### Importing the same Target

* The same Target imported into a different Organization can be imported.
* The same Target from a different source can be imported; for example, the same repository that is present in GitHub can now be imported through GitHub Enterprise into the same Organization.

### Command to run

* Skip all previously imported into all Organizations in a Group: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>`
* Skip all previously imported into a specific Organization: `snyk-api-import-macos list:imported --integrationType=<integration-type> --orgId=<snyk_org_id>`
* Import a single integration and Project source: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>`
* Import multiple integrations and Project sources: `snyk-api-import-macos list:imported --integrationType=<integration-type> --integrationType=<integration-type> --orgId=<snyk_org_id>`

### Supported integration types

* GitHub.com `github`
* GitHub Enterprise `github-enterprise`
* GitLab `gitlab`
* Bitbucket Cloud `bitbucket-cloud`
* Google Cloud Registry `gcr`
* DockerHub registry `docker-hub`
* Azure repos `azure-repos`
