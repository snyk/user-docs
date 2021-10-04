---
description: >-
  Top level vs. nested dependencies, and how to map them back to the manifest
  file
---

# Dependencies

Developers specify a list of dependencies they want included in their application inside a manifest file. Those dependencies explicitly specified are called `top level dependencies`. However each one of those dependencies uses other dependencies and are typically referred to as `nested dependencies` or `transitive dependencies`. All the package managers install the nested dependencies without the developer having to explicitly ask for it.

Snyk CLI automatically scans the nested dependencies as well and those scan results are included in the returned json.

The output json file includes a `from` field which shows the chain of dependencies leading the vulnerability displayed. The way to interpret the information below is that my application _goof_ has a top level dependency _tap_ \(version 5.8.0\) that has a nested dependency _nyc_ that further depends on _istanbul_ that ultimately depends on _handlebars_ \(version 4.0.5\) which if the nested dependency that is carrying the vulnerability in this case.

```yaml
"from": [
  "goof@0.0.3",
  "tap@5.8.0",
  "nyc@6.6.1",
  "istanbul@0.4.3",
  "handlebars@4.0.5"
],
```

If you want to associate this vulnerability back to a line in the manifest file, it needs to be mapped to the line that includes the top level dependency _tap._

