# Serverless files

With Snyk Infrastructure as Code, you can test your configuration files using the CLI.

Snyk Infrastructure as Code for Serverless framework supports scanning the packaged output of Serverless files in Cloudformation JSON format files.

You can **test for an issue on specified Serverless files** as explained on this page.

Be sure you have the [Serverless CLI installed](https://www.serverless.com/framework/docs/getting-started).

After installing the Serverless CLI, **navigate** to the directory that contains your Serverless files and **generate** the Serverless artifacts by entering the following:

```
serverless package --package serverless-artifacts
```

Then you can enter the following Snyk CLI command to test for vulnerabilities:

```
snyk iac test
```

You can also upload the results into the Snyk UI with the command:

```
snyk iac test --report
```
