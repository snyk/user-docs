# Run an analysis with Visual Studio extension

Open your solution and run Snyk scan. Depending on the size of your solution and the time needed to build a dependency graph, it takes less than one or two minutes to get the vulnerabilities.

The extension provides two kinds of results:

* Open Source vulnerabilities
* Snyk Code issues

## Open Source vulnerabilities

Note that your solution must be built successfully in order to allow the CLI to pick up the dependencies and find the vulnerabilities.

If you see only npm vulnerabilities or vulnerabilities that are not related to your C#/.NET projects, that can mean your project was not built successfully and was not detected by the CLI. If you have difficulty or questions, submit a request to [Snyk Support](https://support.snyk.io/hc/en-us/requests/new).

<figure><img src="../../../.gitbook/assets/readme_image_3_1_1.png" alt="Run scan"><figcaption><p>Run scan</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/readme_image_3_1_2.png" alt="Open source vulnerabilities"><figcaption><p>Open Source vulnerabilities</p></figcaption></figure>

## Snyk Code issues

Snyk Code analysis shows a list of security vulnerabilities and code issues found in the application code. For more details and examples of how others fixed the issue, select a security vulnerability or a code security issue and examine the suggestions that Snyk provides.

<figure><img src="../../../.gitbook/assets/readme_image_3_1_3.png" alt="Snyk suggestion panel"><figcaption><p>Snyk suggestion panel</p></figcaption></figure>

The suggestions from Snyk include the recommendation of the Snyk engine using, for example, variable names in your code and the line numbers in red. You can also see:

* Links to external resources that explain the bug pattern in more detail.
* Tags that were assigned by Snyk, such as Security (the issue found is a security issue), Database (the issue is related to database interaction), or In Test (the issue is within the test code).
* Code from open source repositories that can be of help to see how others have fixed the issue.
