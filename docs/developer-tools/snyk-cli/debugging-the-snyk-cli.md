# Debugging the Snyk CLI

When working with the CLI on your Projects, you may encounter unexpected behavior that requires investigation. Such behavior can be caused by configuration, environment, or bugs.

You can run every CLI command in debug mode in order to understand the behavior of the CLI in your environment. To run in debug mode, use the `--debug` or `-d`option to display more verbose log messages on the console.

These messages include basic information about the environment and the command and more detailed information about network interaction with the Snyk API.

In some cases, even more debug information is required. To obtain this, you can change the log level either by setting the environment variable `SNYK_LOG_LEVEL` or by using the `--log-level` options. Both require the name of the log level. The following levels are supported:

* `debug`, the default level
* `trace`, more verbose logging that includes all the information from `debug`, and additional information

The following examples show the use of the `--debug` option and the `SNYK_LOG_LEVEL` environment variable:

1. Using command line arguments to enable trace level logging.

```
snyk test --debug --log-level=trace
```

2. Using command line arguments to enable debugging and environment variables to set the debug level.

```
export SNYK_LOG_LEVEL=trace 
snyk test --debug
```
