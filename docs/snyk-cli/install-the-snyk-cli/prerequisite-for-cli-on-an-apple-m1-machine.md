# Prerequisite for CLI on an Apple M1 machine

To run Snyk CLI on an Apple M1 machine, you need to install Apple's [Rosetta 2](https://support.apple.com/en-gb/HT211861) software. You can do so by running the following in a terminal:

```
softwareupdate --install-rosetta
```

If you do not have the Rosetta 2 software installed you may see the following error:

```
$ snyk --version
bad CPU type in executable

```
