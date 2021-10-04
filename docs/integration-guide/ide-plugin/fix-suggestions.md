# Fix Suggestions

Snyk CLI's test capability also includes information about whether a vulnerability has been fixed \(typically in a newer version of this dependency\) and separately what the upgrade required to implement a fix in the context of this application. There could be instances in which a nested dependency has a vulnerability which is fixed is some newer version, however there is no change to the top level dependency that a developer can do to enjoy it.

To discern if a particular vulnerability has been fixed, inspect the `fixedIn` field. In the case below this particular vulnerability in _handlebar_ seemed to exist in both the 4.x and 3.x streams of this library, and therefore Snyk lets you know the specific version in each stream in which the vulnerability has been fixed. Since the _handlebar_ version currently in use is 4.0.5 \(i.e on the 4.x stream\), the appropriate version to mention it was fixed in is 4.5.3. Note that although handlebar has this version fixed the user does not directly control its version! Instead the user will need to find a version of the top level dependency _tap_ that will indirectly use a new version of handlebars \(see below\):

```yaml
"fixedIn": [
   "4.5.3",
   "3.5.8"
 ],
```

To recommend to the user what change they need to implement in the manifest file to get the fix follow the `upgradePath` field, which shows if the user upgrades the top level dependency _tap_ to version 5.9.1 it will have a trickle effect on the rest of the nested dependencies ultimately including the fixed version of _handlebars_ 4.5.3 :

```yaml
"upgradePath": [
  false,
  "tap@5.9.1",
  "nyc@6.6.2",
  "istanbul@0.5.13",
  "handlebars@4.5.3"
],
```

In a case like below, we can see something different. In this case the upgrade path looks identical to the `from` field, except for the version of the _handlebars._ What this means that in order to get a fixed version of handlebars the developer does NOT need to edit the manifest file, but rather only needs to rebuild the dependencies \(usually by forcing an install command\). This is possible as many dependency managers support semvers which allow for flexibility in the resolved versions. Read more about [semvers here](https://nodesource.com/blog/semver-a-primer/).

```yaml
"upgradePath": [
  false,
  "tap@5.8.0",
  "nyc@6.6.1",
  "istanbul@0.4.3",
  "handlebars@4.5.3"
],
```

## Tracking <a id="c6c77acb-7d74-4439-9659-07723eab4a4b"></a>

In order for Snyk to track the efficacy of your integration we require you to include UTM parameters in various calls. You will need to obtain your UTM tracking information from Snyk.

