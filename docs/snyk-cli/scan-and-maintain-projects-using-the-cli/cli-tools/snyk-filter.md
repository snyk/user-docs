# snyk-filter

The `snyk-filter` tool provides **custom filtering for Snyk CLI output**. `snyk-filter` takes the JSON-formatted output from the [Snyk CLI](https://support.snyk.io/hc/en-us/articles/360003812578-CLI-reference), for example, `snyk test --json` and applies custom filtering of the results, as well as options to fail your build.

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
When using multi-step approach like snyk test --json>result-opensource.json and then piping results to a plugin, you’ll note that the [exit code](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#exit-codes-for-cli-commands) may stop/break the process on your build system before you get to the next step of using the output file with a plugin like snyk-to-html or snyk-filter. There are several options, dependant on the capabilities of your build tools:\
1\) capture the [exit code](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#exit-codes-for-cli-commands) in a parameter, preventing it from being returned to the process, in addition to checking [exit code](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#exit-codes-for-cli-commands) for an error state.\
2\) use ||true or some form of logic to prevent the [exit code](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#exit-codes-for-cli-commands) from terminating the process. Note by doing this, any return is ignored, such as error codes (network/Snyk platform issues or other non-scan result issue). Subsequently the next step of using the JSON will be likely to fail. Reviewing the  [exit code](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#exit-codes-for-cli-commands) , prior to proceeding to the next step in your script, is the recommended approach.

3\) set the step to "continue on failure", if such option exists.
{% endhint %}

### Example with Snyk CLI (using .snyk-filter/snyk.yml by default)

`snyk test --json | snyk-filter`

### Example with Snyk CLI and custom yml file location

`snyk test --json | snyk-filter -f /path/to/example-cvss-9-or-above.yml`

### Example to input a JSON file

`snyk-filter -i snyk_results.json`

### Example with custom yml file location

`snyk-filter -i snyk_results.json -f /path/to/example-high-upgradeable-vulns.yml`

## Options

`--json` to output JSON

## License

License: Apache License, Version 2.0
