# Import Project with Snyk Code

Imported Projects are organized under Target folders on the Projects page, named after the Git repository account and specific repository. See [Import a Project to scan and identify issues](../../discover-snyk/getting-started/#import-a-project-to-scan-and-identify-issues).

Configure the test frequency by navigating to **Code Analysis Project** > **Settings**. See [Snyk Code Analysis test frequency settings](../../snyk-platform-administration/snyk-projects/#test-frequency-settings).

To do a manual test on your repository, you can use the [Retest now option](manage-code-vulnerabilities/#retesting-code-repository).

## Re-import repository to Snyk

If you want to test repositories already imported to Snyk before the [Snyk Code option was enabled](configure-snyk-code.md#enable-snyk-code-in-snyk-web-ui), you need to re-import them (see [Snyk Code conditions](configure-snyk-code.md#conditions)).

1. Log in to the Snyk Web UI and select your [Group and Organization](../../snyk-platform-administration/groups-and-organizations/).
2. Navigate to **Projects** > **Add projects**.
3. In the **Enable Snyk Code** section, change the setting to **Enabled**.
4. Select the Git repository containing the repositories you want to re-import.
5. On the **Personal and Organization repositories** page, select the repositories you want to re-import to Snyk. Repositories that have already been imported are indicated with a checkmark.
6. Select **Add selected repositories** to re-import the selected repositories.

Your selected repositories are re-imported to Snyk. During the re-import process, Snyk Code analyzes these repositories, and a Code analysis Project is added to each repository, containing the results of the Snyk Code test.

<figure><img src="../../.gitbook/assets/project-import-repository.png" alt="Re-test repositories." width="563"><figcaption><p>Re-test repositories</p></figcaption></figure>

## What's next?

* [How importing repositories work](../import-project-repository/#how-importing-repositories-works)
* [Exclude directories and files from the import process](../import-project-repository/exclude-directories-and-files-from-project-import.md)
* [Remove imported repository](../import-project-repository/remove-imported-repository-from-a-project.md)
