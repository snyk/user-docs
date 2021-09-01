# Snyk runtime: what data is used when monitoring my Node.js projects?

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

Once you have installed the Snyk agent to monitor your projects at runtime, Snyk takes the following data from your deployment in order to properly monitor and offer remediation.

| Category | Parameter | Description |
| :--- | :--- | :--- |
| Snyk project ID |  | The unique Snyk ID for the project that is monitored. |
| Unique runtime agent ID |  | An ID generated at runtime to identify the Snyk agent. |
| System info | Runtime agent version | The Snyk agent version. |
|  | Hostname | The name of the machine on which the Snyk agent is installed. |
|  | OS Details | Including info such as operating system type and version |
|  | Node.js version |  |
| List of third-party package vulnerable monitored functions | Details about the third-party packages with vulnerable functions that are observed by Snyk at runtime. |  |
|  | Package name | Name of the package that was seen. |
|  | Relative file path | Of the vulnerable function within the package. |
|  | Function name |  |
| List of third-party package vulnerable functions, called at runtime | Details about the third-party packages with vulnerable functions that are called at runtime. |  |
|  | File path | Of the vulnerable function within the package. |
|  | Function name |  |
| List of all packages loaded at runtime | Package name |  |
|  | Version | Package version |

