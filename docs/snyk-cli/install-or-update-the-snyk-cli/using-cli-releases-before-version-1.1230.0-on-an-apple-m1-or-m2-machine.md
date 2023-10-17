# Using CLI releases before version 1.1230.0 on an Apple M1 or M2 machine

If you run releases of the Snyk CLI before [version 1.1230.0](https://github.com/snyk/cli/releases/tag/v1.1230.0) on an Apple M1 machine, you may see the following error:

```
$ snyk --version
bad CPU type in executable
```

On an M2 machine, you may see this error:

`Unknown system error -86`

To avoid these errors with releases of Snyk CLI before [version 1.1230.0](https://github.com/snyk/cli/releases/tag/v1.1230.0), install the Apple [Rosetta 2](https://support.apple.com/en-gb/HT211861) software:

&#x20;`softwareupdate --install-rosetta`

Beginning with version 1.1230.0, the Snyk CLI supports Apple silicon, including M1 and M2 machines, so you no longer need the Rosetta 2 software.
