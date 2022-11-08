# Invalid string length error when scanning Gradle projects

If you use the `--json-file-output=<OUTPUT_FILE_PATH>` option when you scan a very large Gradle project, you may see the invalid string length error.

To complete the scan you can remove the `--json-file-output` option.

If you require JSON file output, try increasing the severity threshold, for example `--severity-threshold=critical`. This is likely to reduce the number of critical vulnerabilities reported and thus reduce the size of the JSON file output.
