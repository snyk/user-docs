# Net new issues (delta) scan troubleshooting

## Problem&#x20;

**Net new issues scans fail to find issues on npm-based reference scans.**

The net new issues filter in the IDEs compares the software Projects in two folders or two branches from a Git project.\
\
Some package managers, such as NuGet, pip, npm, and so on, require specific manifest files to be generated before you can perform a scan analysis. These files, for example, `package-lock.json` or `obj/project.assets.json` for NuGet, are typically generated by the Project's compiler or running a specific installation command.

## Solution

Required manifest files must be generated before you can perform a scan analysis. These files, for example, `package-lock.json` or `obj/project.assets.json` for NuGet must have been generated on the current branch or folder and either:

* Be committed to the reference branch before the scan
* Be prepared (compiled) into a different directory, that will be used as a reference folder for Snyk scan

<figure><img src="../../../.gitbook/assets/image (695).png" alt=""><figcaption><p>Choose the reference branch or folder</p></figcaption></figure>

\
