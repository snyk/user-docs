# Using CLI releases before version 1.1230.0 on an Apple M1 machine

To run releases of Snyk CLI before [version 1.1230.0](https://github.com/snyk/cli/releases/tag/v1.1230.0) on an Apple M1 machine, you may see the following error:

```
$ snyk --version
bad CPU type in executable
```

When trying to run releases of Snyk CLI before [version 1.1230.0](https://github.com/snyk/cli/releases/tag/v1.1230.0), you also may see the following error:

```
$ snyk --version
bad CPU type in executable
```

To avoid these errors with releaes of Snyk CLI before [version 1.1230.0](https://github.com/snyk/cli/releases/tag/v1.1230.0), install Apple's [Rosetta 2](https://support.apple.com/en-gb/HT211861) software: `softwareupdate --install-rosetta`.
