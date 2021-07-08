# Monitored projects are overwriting each other in the UI

When you run `snyk monitor`, we use the `name` of the root package to find the appropriate project. The CLI will use the directory base-name if the root package has no name so if you are using the same default working directory for all your CI jobs then this could mean they are overwritten in the UI due to Snyk matching them to the wrong existing project.

In order to avoid this you'll need to use a different working directory for each project monitored in your CI.

For example, Azure DevOps have a default working directory which will often lead to projects monitored here being prefixed with an 's' in the Snyk UI. To get around this you can use the Azure specific argument **Test \(Working\) Directory** \(testDirectory\) and pass in the path to the directory you are currently testing. Please see the table with the parameters in our docs [here](https://support.snyk.io/hc/en-us/articles/360004127677-Azure-Pipelines-integration). Otherwise you will need to change the default directory in Azure.

