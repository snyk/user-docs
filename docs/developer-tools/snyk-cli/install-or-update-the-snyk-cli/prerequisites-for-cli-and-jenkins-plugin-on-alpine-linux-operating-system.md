# Prerequisites for CLI and Jenkins plugin on Alpine Linux operating system

Before running Snyk CLI and Snyk Jenkins plugin on the Alpine Linux operating system, follow these steps to establish the prerequisites:

1. Ensure that `libstdc++` is installed. \
   To install `libstdc++`, open a terminal from within the relevant container and run `apk add libstdc++` .\
   The command either adds the shared library to your container environment or returns `OK` if the shared library is already installed in `/usr/lib`.&#x20;
2. Ensure that Snyk CLI version 1.185.5 or above is installed on your system. \
   Snyk recommends that you install the **latest** version. &#x20;
3. To use the Snyk Jenkins plugin, ensure also that the Snyk Jenkins plugin v2.10.0 or above is installed. \
   Snyk recommends that you install the **latest** version.&#x20;
4. After you install the Snyk CLI, you must [authenticate](../authenticate-to-use-the-cli.md).

For additional information, see [Install or update the Snyk CLI](./) and [Jenkins Plugin](../../snyk-ci-cd-integrations/jenkins-plugin-integration-with-snyk.md) in the CI/CD docs.
