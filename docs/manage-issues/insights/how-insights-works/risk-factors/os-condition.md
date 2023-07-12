# OS condition

Some vulnerabilities have specific constraints which must be met for the problem to be exploitable. One such constraint is the operating system. Some vulnerabilities are exploitable only when the affected package is executing on a specific operating system, for example, Windows. Insights is able to compare the vulnerability condition specified in the package with the operating system used by the runtime environment executing the package.

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

When an image is scanned by Snyk Container, the information about which operating system the base of the image is running on is exposed. Therefore, whenever Insights is able to determine that a package is a dependency of the image or that the package has been included in the image, it compares the vulnerability information with the image specification. The same applies to problems identified directly in the image entity.

Note the following **technical details**.

On a regular schedule, every hour, the data pipeline takes a snapshot of all Snyk Projects and data sources and extrapolates packages and images. This snapshot is used to determine which images and packages are known to Snyk for any given customer.&#x20;

Snyk Project tags are used to enable the customer to relate image assets to packages. This information is extracted from the hourly data snapshot, and a basic composition graph is generated between images and packages.

The data pipeline looks at all the issues reported for the package and image and attempts to map the vulnerability condition between the entities.
