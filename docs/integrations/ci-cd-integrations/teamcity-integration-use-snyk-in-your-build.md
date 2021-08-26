# TeamCity integration: use Snyk in your build

For any project, you can add Snyk to your build to scan the code while you build and to fail the build for vulnerabilities, based on your configurations.

We recommend running a build with the Snyk Security step before deployment, to ensure excellent security posture.

For additional information with TeamCity and its features, refer to their documentation.

## How to configure your build with a Snyk step

1. Add the step to a new or existing project:
   * For new projects, after configuring the Git repo from which to create the build, activate the auto-detect feature to automatically identify relevant steps for your project build.
   * For existing projects, navigate to edit the project build steps. When complete, Snyk Security appears in the list of suggested steps and the current test policy appears in the Parameters Description column:

     ![image2.png](../../.gitbook/assets/uuid-97395df2-f141-6f77-4551-f19397ac0781-en.png)

2. Navigate to configure the Snyk Security step as follows:
   * Click anywhere on the Snyk Security row to access the configuration screen, or
   * for existing projects, click Add build step to access the configuration screen.

     ![image3.png](../../.gitbook/assets/uuid-88e38280-121e-a17b-cfd3-9fde89305b5c-en.png)

3. Configure the TeamCity fields \(Runner type, Step name and Execute Step\).
4. Optionally, click Show advanced options. Additional Snyk parameters are revealed:

![image4.png](../../.gitbook/assets/uuid-8f294e8d-ca5e-123b-2992-a98c1e62fd6f-en.png)


5. Configure Snyk Settings and Snyk Tool Settings. For more information see TeamCity configuration parameters.
6. Once configured, run the build. When the Snyk Security step ends successfully, you can navigate to the Snyk Security Report tab to view results within TeamCity and to navigate seamlessly to the Snyk UI for further action:

![image5.png](../../.gitbook/assets/uuid-e8b1fd6f-3b49-069c-c9fe-c0948931b141-en.png)


7. From the top of the report, click View on Snyk.io to view the snapshot and vulnerability information directly from our app.

