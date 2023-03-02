# Troubleshooting and known issues with Visual Studio extension

## Known issues

### Could not detect supported target files

**Solution** Open Visual Studio Options to go to the Project Settings of the Snyk extension and check Scan all projects.

![](../../../.gitbook/assets/readme\_image\_4\_1.png)

### The system cannot find the file specified

**Solution** This issue is related to CLI file. Close and open Snyk extension window to start CLI download.

### The specified executable is not a valid application for this OS platform

**Solution** This issue is related to CLI file and its integrity. Remove CLI from in\
`%HOMEPATH%\AppData\Local\Snyk\snyk-win.exe`. Close and open Snyk extension window to start CLI download.

### Snyk Code no supported code available

**Solution** Check .gitignore and .dcignore file rules. Check if there are any rules that exclude your project's source files.

## Troubleshooting

You can find logs in the user AppData directory:

```
%HOMEPATH%\AppData\Local\Snyk\snyk-extension.log
```
