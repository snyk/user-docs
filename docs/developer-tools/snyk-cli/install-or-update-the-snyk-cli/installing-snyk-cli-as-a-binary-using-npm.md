# Installing Snyk CLI as a binary using npm

As a part of the [evolution of the Snyk CLI towards an extensible approach](https://snyk.io/blog/evolving-the-snyk-cli-through-an-extensible-approach/), Snyk is adapting the deployment via npm. While Snyk is attempting to make this as smooth and transparent to users as possible, the change might still cause issues for some users.

Snyk is starting to deploy the new Extensible CLI as a binary via npm and has implemented a graceful degradation in case of errors. If graceful degradation is triggered, users will see a warning like the one that follows.

```bash
------------------------------ Warning -------------------------------
You are currently running a degraded version of the Snyk CLI.
As a result, some features of the CLI will be unavailable.
For information on how to resolve this, please see this article: https://docs.snyk.io/snyk-cli/installing-snyk-cli-as-a-binary-via-npm
For any assistance, please check http://support.snyk.io/.
------------------------------- Warning -------------------------------
```

This warning is typically accompanied by an error message indicating the nature of the error. Some of the errors will include resolution options as shown in the following example:

```bash
------------------------------ Warning -------------------------------
You don't have the permissions to install Snyk. Please try the following options:
* If you are installing with increased privileges (for example sudo), try adding unsafe-perm a parameter to npm install
* If you run NPM <= 6, please upgrade to a later version.
If the problems persist please check http://support.snyk.io/.
------------------------------ Warning -------------------------------
```

If you continue to experience issues and they are not resolved by following any of the suggestions, you have the following options

* Install the Snyk CLI from any of the other deployment channels suggested on the [installation page](./).
* If you experience an error that is not automatically handled, or you suspect that the Extensible CLI behaves incorrectly, you can use `--legacy-cli` to force use of the legacy Typescript CLI instead of the Extensible CLI.
