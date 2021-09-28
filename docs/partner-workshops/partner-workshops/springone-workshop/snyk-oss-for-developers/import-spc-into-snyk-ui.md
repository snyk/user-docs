# Import SPC into Snyk UI

## **Add projects to Snyk for GitHub repositories**

Snyk tests and monitors GitHub repositories that are in any of our supported languages by evaluating root folders and custom file locations.

**Adding Projects to Snyk**

1\) Go to Projects and click Add projects. Choose the tool from which to import your projects

![](../../../.gitbook/assets/project_import.png)

2\) A popup screen opens with all the available repositories under the selected integration

![](../../../.gitbook/assets/select_repo.png)

{% hint style="info" %}
Select the Spring-Petclinic ****repository.
{% endhint %}

3\) Select the repository to import into Snyk. This will monitor the repositories for security and license issues. To import all repositories for a specific organization, checkmark the organization.

4\) Click add selected repositories. Snyk evaluates root folders and custom file locations. If no manifest files are found on the root level or in the paths you configure, Snyk notifies you that no files can be imported.

## Imported projects

As Snyk is importing your project you will see the import status bar followed by a **green box** asking you to refresh the screen. 

{% hint style="info" %}
Refresh green box not shown
{% endhint %}

![](../../../.gitbook/assets/import_bar.png)

After refreshing, you will see the imported projects in the Snyk UI. Expand the project to see the package manager file, in our case _**pom.xml**_ 

![](../../../.gitbook/assets/screen-shot-2020-08-21-at-4.43.05-pm%20%281%29.png)

