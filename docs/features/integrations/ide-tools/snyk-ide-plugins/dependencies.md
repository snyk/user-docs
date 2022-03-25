# Dependencies and IDE plugins

Developers specify a list of dependencies they want included in their application inside a manifest file. The dependencies explicitly specified are called `top level dependencies`. In turn, each one of those dependencies uses other dependencies and they are typically referred to as `nested dependencies` or `transitive dependencies`. All the package managers install the nested dependencies without the developer having to explicitly ask for that to happen.

Snyk CLI automatically scans the nested dependencies as well as the top level dependencies, and those scan results are included in the returned `json` file.

The output `json` file includes a `from` field which shows the chain of dependencies leading to the vulnerability displayed. The way to interpret the information that follows is: my application _goof_ has a top level dependency _tap_ (version 5.8.0) that has a nested dependency _nyc_ that further depends on _istanbul_ that ultimately depends on _handlebars_ (version 4.0.5), which is the nested dependency that is carrying the vulnerability in this case.

```yaml
"from": [
  "goof@0.0.3",
  "tap@5.8.0",
  "nyc@6.6.1",
  "istanbul@0.4.3",
  "handlebars@4.0.5"
],
```

If you want to associate this vulnerability back to a line in the manifest file, the vulnerability must be mapped to the line that includes the top level dependency _tap._
