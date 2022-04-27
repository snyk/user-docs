# Fix suggestions and IDE plugins

The Snyk CLI test capability also includes information about whether a vulnerability has been fixed, typically in a newer version of the dependency, and separately what upgrade is required to implement a fix in the context of the application. There may be instances in which a nested dependency has a vulnerability which is fixed is some newer version but there is no change to the top level dependency that a developer can do to employ the fix.

To discern if a particular vulnerability has been fixed, inspect the `fixedIn` field. In the case that follows, the particular vulnerability in _handlebars_ seemed to exist in both the 4.x and 3.x streams of the library, and therefore Snyk lets you know the specific version in each stream in which the vulnerability has been fixed. Since the _handlebars_ version currently in use is 4.0.5 (that is, on the 4.x stream), the appropriate version to mention the vulnerability was fixed in is 4.5.3. Note that although _handlebars_ has this version fixed, the user does not directly control its version. Instead the user must find a version of the top level dependency _tap_ that indirectly uses a new version of _handlebars_. Refer to the code that follows:

```yaml
"fixedIn": [
   "4.5.3",
   "3.5.8"
 ],
```

To find the change the user must implement in the manifest file to get the fix, follow the `upgradePath` field, which shows that, if the user upgrades the top level dependency _tap_ to version 5.9.1, it will have a trickle effect on the rest of the nested dependencies ultimately including the fixed version of _handlebars_ 4.5.3 :

```yaml
"upgradePath": [
  false,
  "tap@5.9.1",
  "nyc@6.6.2",
  "istanbul@0.5.13",
  "handlebars@4.5.3"
],
```

In a case like the following, you see something different. The upgrade path looks identical to the `from` field, except for the version of _handlebars._ This means that in order to get a fixed version of _handlebars_ the developer does NOT need to edit the manifest file, but rather only needs to rebuild the dependencies (usually by forcing an install command). This is possible as many dependency managers support semvers which allow for flexibility in the resolved versions. Read more about [semvers.](https://nodesource.com/blog/semver-a-primer/)

```yaml
"upgradePath": [
  false,
  "tap@5.8.0",
  "nyc@6.6.1",
  "istanbul@0.4.3",
  "handlebars@4.5.3"
],
```

**Note:** In order for Snyk to track the efficacy of your integration you must include UTM parameters in various calls. You must obtain your UTM tracking information from Snyk.
