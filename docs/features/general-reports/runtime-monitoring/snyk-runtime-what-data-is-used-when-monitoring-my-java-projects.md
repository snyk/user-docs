# Snyk runtime: what data is used when monitoring my Java projects?

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

Once you have installed the Snyk agent to monitor your projects at runtime, Snyk takes the following data from your deployment in order to properly monitor and offer fixes.

| Category                                                            | Parameter                                                                                              | Description                                                   |
| ------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------- |
| Snyk project ID                                                     |                                                                                                        | The unique Snyk ID for the project that is monitored.         |
| Unique runtime agent ID                                             |                                                                                                        | An ID generated at runtime to identify the Snyk agent.        |
| System info                                                         | Runtime agent version                                                                                  | The Snyk agent version.                                       |
|                                                                     | Hostname                                                                                               | The name of the machine on which the Snyk agent is installed. |
|                                                                     | OS Details                                                                                             | Including info such as operating system type and version      |
|                                                                     | JVM details                                                                                            | <p>VM name</p><p>Vendor</p><p>Version</p>                     |
| List of third-party package vulnerable monitored functions          | Details about the third-party packages with vulnerable functions that are observed by Snyk at runtime. |                                                               |
|                                                                     | JAR name                                                                                               | Name of the JAR that was seen.                                |
|                                                                     | Relative file path                                                                                     | Of the vulnerable function within the JAR.                    |
|                                                                     | Function name                                                                                          |                                                               |
|                                                                     | Artifact ID                                                                                            |                                                               |
|                                                                     | Class names                                                                                            | Class names for the function.                                 |
| List of third-party package vulnerable functions, called at runtime | Details about the third-party packages with vulnerable functions that are called at runtime.           |                                                               |
|                                                                     | JAR name                                                                                               | Name of the JAR that was seen.                                |
|                                                                     | File path                                                                                              | Of the vulnerable function within the JAR.                    |
|                                                                     | Function name                                                                                          |                                                               |
|                                                                     | Artifact ID                                                                                            |                                                               |
|                                                                     | Class names                                                                                            | <p>Class names for the function.</p><p>Java only</p>          |
|                                                                     | JAR crc32                                                                                              | crc32 JAR checksum                                            |
| List of all JAR files loaded at runtime                             | JAR file path                                                                                          | The file path                                                 |
|                                                                     | Artifact ID                                                                                            |                                                               |
