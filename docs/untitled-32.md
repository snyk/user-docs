# Visual Studio extension

* [ Eclipse plugin](/hc/en-us/articles/360004032337-Eclipse-plugin)
* [ JetBrains plugins](/hc/en-us/articles/360004032317-JetBrains-plugins)
* [ Visual Studio extension](/hc/en-us/articles/360020073958-Visual-Studio-extension)
* [ Visual Studio Code extension for Snyk Code](/hc/en-us/articles/360018585717-Visual-Studio-Code-extension-for-Snyk-Code-)

##  Visual Studio extension

This documentation describes the Visual Studio extension for Snyk Open Source.

### Prerequisites

* Windows OS
* Supported versions of Visual Studio: 2015, 2017, 2019. Compatible with Community and Professional.

### Running Snyk scan

Open your solution, then click **run scan** to run a Snyk scan:

  
When the scan completes, the results appear in the Snyk window:

You must have successfully built your solution for the Snyk to find the dependencies and vulnerabilities.

### Filtering vulnerabilities

You can filter vulnerabilities by name or by severity.

* Filter by name by typing the name of the vulnerability in the search bar:
* Filter by severity by selecting one or more of the the severities when you open the search bar filter.

### Configure Snyk extension \(Project settings\)

The **Scan all projects** option is enabled  by default. It adds the **--all-projects** option for Snyk CLI, scanning all projects by default:

If you see the message: **Could not detect supported target files**, then ensure **Scan all projects** is enabled.

### Access to log files

Run Visual Studio with **/log** parameter, and set the path for where to save the log file:

```text
devenv.exe /log "%DIRECTORY_PATH%\MyVSLog.xml"
```

For example, for Visual Studio 2019 :

```text
C:\Program Files (x86)\Microsoft Visual Studio\2019\Community\Common7\IDE\devenv.exe /log "C:\Temp\MyVSLog.xml"
```

