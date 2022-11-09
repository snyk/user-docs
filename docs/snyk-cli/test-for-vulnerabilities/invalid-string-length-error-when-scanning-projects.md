# Invalid string length error when scanning projects

If you use the `--json` or `--json-file-output=<OUTPUT_FILE_PATH>` option when you scan using the `--all-projects` option, you may see the invalid string length error.

This error can also occur when you scan a very large Gradle project using the `--json` or `--json-file-output=<OUTPUT_FILE_PATH>` option.

To complete the scan you can remove the `--all-projects` option, or remove the `--json` or `--json-file-output` option.

If you require JSON file output, try increasing the severity threshold, for example `--severity-threshold=critical`. This is likely to reduce the number of less critical vulnerabilities reported and thus reduce the size of the JSON file output.
