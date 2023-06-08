# Custom Rego Builtins

The SDK also registers some helper Rego functions that can be used while testing.

These are:

* `hcl2.unmarshal_file`: takes in the path to a file and parses it from HCL2 format to JSON
* `yaml.unmarshal_file`: takes in the path to a file and parses it from YAML format to JSON

These functions are used by the testing framework at `lib/testing/main.rego` so that you can pass in the path to fixture files instead of having to generate the JSON fixture yourself.
