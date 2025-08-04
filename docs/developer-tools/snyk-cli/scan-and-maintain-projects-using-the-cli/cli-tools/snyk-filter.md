# snyk-filter

The `snyk-filter` tool provides **custom filtering for Snyk CLI output**. `snyk-filter` takes the JSON-formatted output from the [Snyk CLI](../../), for example, `snyk test --json` and applies custom filtering of the results, as well as options to fail your build.

## Clone and install snyk-filter

First, clone the [repo](https://github.com/snyk-labs/snyk-filter). Then run:

`npm install -g`

**`snyk-filter` uses the `node-jq` library**, which requires that a [`jq`](https://stedolan.github.io/jq/) binary be installed. This typically happens transparently via `npm install -g`, but on some systems `jq` is not properly installed locally. If you receive an error after installation regarding `node-jq`, then install `jq` manually to avoid this error.

```
# install jq ahead of time (ubuntu example)
sudo apt-get install -y jq

# tell node-jq to skip trying to install it on its own
export NODE_JQ_SKIP_INSTALL_BINARY=true

# tell node-jq where the existing jq binary is
export JQ_PATH=$(which jq)

# finally, install snyk-filter (does not work with node version > 12)
sudo npm install -g
```

## Usage

1. Implement your custom `jq` filters in a .snyk-filter/snyk.yml file relative to your current working directory where you are running `snyk test`. See [sample-filters](https://github.com/snyk-tech-services/snyk-filter/tree/develop/sample-filters) and tweak things from there; use [JQPlay](https://jqplay.org/)
2. Then pipe your `snyk test --json` output into `snyk-filter` or use the `-i` argument to input a JSON file. Use the `-f` argument to point to the yml file containing your custom filters if you are not using the default location (.snyk-filter/snyk.yml).
3. The exit code from `snyk-filter` is 0 for pass (no issues) and 1 for fail (issues found).

{% hint style="info" %}
When you use a multi-step approach like `snyk test --json > result-opensource.json` and then pass the results to a plugin, the [exit code](../../cli-commands-and-options-summary.md#exit-codes-for-cli-commands) may stop or break the process on your build system before you get to the step of passing the output file to a tool like `snyk-to-html` or `snyk-filter`. You have several options, depending on the capabilities of your build tools:\
\
1\) Capture the [exit code](../../cli-commands-and-options-summary.md#exit-codes-for-cli-commands) in a parameter to prevent it from being returned to the process in addition to checking for an error state.\
2\) Use `||true` or some form of logic to prevent the [exit code](../../cli-commands-and-options-summary.md#exit-codes-for-cli-commands) from terminating the process.\
Note that when you do this, any return code is ignored, such as error codes signifying network or Snyk platform issues or another non-scan result issue. The next step in using the JSON is likely to fail. It is recommended that you review the exit code before you proceed to the next step in your script.\
3\) Set the step to `continue on failure`, if such an option exists.
{% endhint %}

### Example with Snyk CLI (using .snyk-filter/snyk.yml by default)

`snyk test --json | snyk-filter`

### Example with Snyk CLI and custom yml file location

`snyk test --json | snyk-filter -f /path/to/example-cvss-9-or-above.yml`

### Example to input a JSON file

snyk test --json-file-output=results-opensource.json

`snyk-filter -i` results-opensource.json

### Example with custom yml file location

`snyk-filter -i snyk_results.json -f /path/to/example-high-upgradeable-vulns.yml`

## Options

`--json` to output JSON

## License

License: Apache License, Version 2.0
