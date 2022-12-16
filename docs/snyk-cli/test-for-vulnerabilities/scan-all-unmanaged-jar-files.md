# Scan all unmanaged JAR files

The Snyk CLI can scan unmanaged JAR files in [Java applications](../../products/snyk-open-source/language-and-package-manager-support/snyk-for-java-gradle-maven.md) to identify which open source package they contain.

The CLI identifies the package name, version, and vulnerabilities only if the package is available in Maven Central, and the JAR file hash matches the hash in Maven Central.

Use the `snyk test --scan-all-unmanaged` CLI command to scan all JAR files in a single folder.

You can also scan each JAR file individually using the `snyk test --scan-unmanaged --file=/path/to/file` command.

Testing each JAR file individually shows the name of the JAR file that was scanned on the Snyk web UI. Using --`scan-all-unmanaged` shows the package name instead of the file name.

**WAR file support**: You can scan individual WAR files that are published in Maven Central. To scan open-source dependency JARs directly, you must extract (unzip) all other WAR files or JAR files containing other JARs.

**Prerequisite:** You must install a [supported version](../../products/snyk-open-source/language-and-package-manager-support/snyk-for-java-gradle-maven.md#supported-versions) of Maven to scan unmanaged JAR files.

## Recursively scanning all subfolders

Java apps often have JAR files in a number of different folders within an application.

Below is a Linux/Mac Bash script that recurses through all subfolders starting with the current folder and tests each JAR file found.&#x20;

Setting a value for the `REMOTE_REPO_URL` variable is important; it is used to combine all scan results under a single Snyk project in the UI using the `--remote-repo-url` parameter.

```
#!/bin/bash

SNYK_CLI_BINARY_NAME=snyk-cli
SNYK_CLI_BINARY_LOCATION=https://github.com/snyk/cli/releases/latest/download/
REMOTE_REPO_URL= #Insert desired project name in Snyk UI here

detected_jars=""
undetected_jars=""
detected_count=0
undetected_count=0

[[ -z "$REMOTE_REPO_URL" ]] && { echo "REMOTE_REPO_URL is empty. Please enter REMOTE_REPO_URL (line 6) and re-run script." ; exit 1; }

#Download Snyk binary specific to OS (MacOS or Linux)
case "$(uname -s)" in
   Darwin)
     curl -L -O $SNYK_CLI_BINARY_LOCATION/snyk-macos
     mv snyk-macos snyk-cli
     ;;
   Linux)
     curl -L -O $SNYK_CLI_BINARY_LOCATION/snyk-linux
     mv snyk-linux snyk-cli
     ;;
esac

chmod +x $SNYK_CLI_BINARY_NAME

#Loop through folders recursively to find all .jar files
#NOTE: will ERROR on files with whitespace in name or contained in directories with whitespace in name
for file in $(find . -type f -name '*.jar' | uniq)
do
echo ""
echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-="    
echo $file
echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=" 

#Execute Snyk monitor for each .jar
if (./$SNYK_CLI_BINARY_NAME monitor --scan-unmanaged --file=$file --project-name=$file --remote-repo-url=$REMOTE_REPO_URL) then
  detected_jars+=$file'\n'
  let detected_count++
else
  undetected_jars+=$file'\n'
  let undetected_count++
fi

done

#Output metrics to the console
echo ""
echo ""
echo ""
echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=" 
echo "Detected jars ($detected_count) - does not include transitive dependencies:"
echo ""
printf $detected_jars
echo ""
echo ""
echo ""
echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=" 
echo "Undetected jars ($undetected_count) - not found on Maven Central:"
echo ""
printf $undetected_jars
```

Here is a Windows batch script for scanning JARs in all subfolders, run from a `scanjar.bat` file.&#x20;

To use this script, you must have installed the Snyk CLI.

```batch
REM Usage:    
REM For example: scanjar.bat "C:\workspace\app" "myapp" 
SET WORKSPACE=%1 
SET REMOTE_REPO_URL=%2 
for /R %WORKSPACE% %%f in (*.jar) do cmd /c snyk monitor --scan-unmanaged --remote-repo-url=%REMOTE_REPO_URL% --file=%%f
```

Here is example in the Snyk UI of using these scripts with `REMOTE_REPO_URL` set to "econnect".

<figure><img src="../../.gitbook/assets/untitled.png" alt="Result of scanning unmanaged JAR files"><figcaption><p>Result of scanning unmanaged JAR files</p></figcaption></figure>
