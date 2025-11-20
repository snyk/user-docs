# Risk factor: OS condition

{% hint style="info" %}
The OS condition risk factor is available only for Snyk AppRisk users.
{% endhint %}

Some vulnerabilities have specific constraints that must be met for the problem to be exploitable. One such constraint is the operating system. Some vulnerabilities are exploitable only when the affected package is executing on a specific operating system, for example, Windows. Snyk can compare the vulnerability condition specified in the package with the operating system used by the runtime environment executing the package.

## Vulnerability OS conditions and risk factors

If the vulnerability condition matches the operating system used by the runtime environment or if there is no condition, meaning that the vulnerability applies to all operating systems, the OS Condition risk factor is generated.

| OS condition value | Runner OS value | Result                 |
| ------------------ | --------------- | ---------------------- |
| Any OS             | Unknown         | No risk factor         |
| Linux              | Unknown         | No risk factor         |
| Windows            | Unkown          | No risk factor         |
| Any OS             | Linux           | Risk factor identified |
| Linux              | Linux           | Risk factor identified |
| Windows            | Linux           | No risk factor         |
| Any OS             | Windows         | Risk factor identified |
| Linux              | Windows         | No risk factor         |
| Windows            | Windows         | Risk factor identified |

When an image is scanned by Snyk Container, the information about which operating system the base of the image is running on is exposed. Therefore, whenever Snyk Issues can determine that a package is a dependency of the image or that the package has been included in the image, it compares the vulnerability information with the image specification. The same applies to problems identified directly in the image entity.

## Kubernetes Connector integration

The OS condition risk factor works with your [Kubernetes Connector](broken-reference) integrations.

The Kubernetes Connector leverages the OS condition risk factors to enhance the identification of vulnerabilities within containerized applications. It continuously checks and compares the operating systems of running containers with the known risk factors from the Snyk database. This integration helps to detect potentially vulnerable packages or images in real-time based on the operating system conditions. It allows proactive security measures within Kubernetes environments.

## Technical details for OS condition risk factor

Every hour, the data pipeline takes a snapshot of all Snyk Projects and data sources and extrapolates packages and images. This snapshot determines which images and packages are known to Snyk for any customer.

Snyk Project tags are used to enable the customer to relate image assets to packages. This information is extracted from the hourly data snapshot, and a basic composition graph is generated between images and packages.

The data pipeline examines all the issues reported for the package and image and attempts to map the vulnerability condition between them.
