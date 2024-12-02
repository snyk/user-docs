# Invalid string length error when scanning projects

The invalid string length error can occur in the following situations:

* Scanning **many Projects** at once using the `--all-projects` option in combination with the `--json` or `--json-file-output=<OUTPUT_FILE_PATH>` options.
* Scanning a **large Project** using the `--json` or `--json-file-output=<OUTPUT_FILE_PATH>` options.

You can use the following workarounds to avoid the invalid string length error.

* Remove the `--json` or `--json-file-output=<OUTPUT_FILE_PATH>` option. This may not be possible in integrations using the CLI, for example, CI/CD integrations.
* If you require the JSON output, try increasing the severity threshold using the `--severity-threshold=<SEVERITY_LEVEL>` option. This is likely to reduce the number of less critical vulnerabilities reported and thus reduce the size of the JSON file output.
* If you are scanning multiple Projects using the `--all-projects` option, try removing this option and scanning Projects individually.
* Another method of reducing the JSON output is adding the `--prune-repeated-subdependencies` option. This may sufficiently reduce the size of the resulting scan so the JSON can be resolved.&#x20;

For detailed information about the CLI options, see the [CLI test command help](../commands/test.md).
