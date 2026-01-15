# Debugging the Snyk CLI

When working with the CLI on your Projects, it is possible that you encounter unexpected behavior that requires investigation. Such behavior can be caused by configuration, environment, or bugs.

## Debug mode

You can run every CLI command in debug mode to understand the behavior of the CLI in your environment. To run in debug mode, use the `--debug` or `-d` option to display more verbose log messages on the console.

These messages include basic information about the environment and the command, along with more detailed information about network interactions with the Snyk API.

### Configuring log levels

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

## Error Catalog

The CLI implements error codes from the Snyk [Error Catalog](../../scan-with-snyk/error-catalog.md) for many issues that can occur during CLI execution. These are designed to make issues during CLI execution clearer and, where possible, easier to address.

For example, when using the CLI while unauthorized, it outputs the appropriate Error Catalog error.

```
> snyk test

 ERROR   Authentication error (SNYK-0005)
         Authentication credentials not recognized, or user access is not provisioned.  
         Revise credentials and try again, or request access from your Snyk             
         administrator.                                                                 
                                                                                        
           Use `snyk auth` to authenticate.                                             

Status:  401 Unauthorized 
Docs:    https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-0005 
                                                                     
ID:      urn:snyk:interaction:11111111-2222-3333-4444-555555555555
```

## Exit codes

The CLI returns different exit codes depending on the outcome of a command. In general, there are three standard exit codes:

| Exit code | Meaning                                                                                     |
| --------- | ------------------------------------------------------------------------------------------- |
| 0         | Command executed successfully and no issues found                                           |
| 1         | Command executed successfully and issues found                                              |
| 2         | Command did not execute successfully, review debug output, error codes for more information |

Additionally, error codes from the Snyk Error Catalog may result in different exit codes:

| Exit code | Meaning                                |
| --------- | -------------------------------------- |
| 3         | No supported files/projects detected   |
| 69        | Service unavailable                    |
| 75        | Temporary failure / Maintenance window |
| 77        | No permission                          |
