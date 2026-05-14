# TeamCity configuration parameters

This page provides information about [Snyk settings](teamcity-configuration-parameters.md#snyk-settings), [Additional parameters](teamcity-configuration-parameters.md#additional-parameters), and [Snyk tool settings](teamcity-configuration-parameters.md#snyk-tool-settings).

## Snyk settings

### Severity threshold

* Default: low
* Specify a threshold: `low`

### Monitor Project on build

* Default: ON
* Snyk runs the `snyk monitor` command during the build, sending a Project snapshot to the Snyk Web UI, and continuing to monitor the Project for vulnerabilities after this build.

### File

* **Optional**
* If the manifest file is not on the root level, enter the relative path to that file.

### Organization

* Optional
* The ID of the Snyk Organization to which this Project will be associated when imported to the UI.
* Copy the Organization ID from the Snyk UI in the Settings area.

### Project name

* Optional
* Enter any unique name for this Project to help with recognizing it when viewing it in the Snyk UI.

## Additional parameters

* Optional
* Enter additional CLI options as necessary. See the [CLI documentation](../../snyk-cli/) and [cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/).

## Snyk tool settings

### Snyk PAT or API token

* From the Settings area in the Snyk UI, copy the Organization, Personal API token, PAT, or create a service account.
* Use the token to authenticate your Snyk account when connecting to TeamCity.

### Snyk version

* Default: the most recent version
* If you would like an older Snyk CLI version to support the plugin, select the plugin version to use in your build.
* Snyk recommends configuring automatic upgrades and using the most recent version.

### Use custom build tool path:

* Specify which tool instance in your local environment Snyk is to use for this build.
* Without a specified tool, Snyk auto-detects the tool and locates it in your environment based on the Project type.
