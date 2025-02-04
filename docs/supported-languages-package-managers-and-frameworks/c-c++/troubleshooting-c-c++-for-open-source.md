# Troubleshooting C/C++ for open source

## **Is my Snyk Open Source code sent to Snyk servers?**

No. The files are converted to a list of hashes before they are sent for scanning.

## **Why did Snyk not find any dependencies?**

Snyk stores the official releases of many open-source components in the Snyk database but it is possible that the source code you scanned is not there or is not found. If your scan does not find any dependencies [submit a request to support](https://support.snyk.io).

Here are a few things that you can check on your own:

* The source code of the dependencies you scanned is available as source code (unpacked) in the folder that you scanned. If you use a package manager, such as Conan, the source code is likely to be in the Conan cache, along with the source code of other dependencies of your other Projects. To scan dependencies managed by a package manager, Snyk recommends that you do that in a clean environment (for example, during a build).
* The source code of the dependencies is not from an official release of the open source software (OSS) component, and Snyk does not have it in the database.
* The source code of the OSS has been modified too much, so Snyk cannot detect it. If there are too few files and you modify most of them, Snyk cannot match them to a component from the Snyk database. Examples of common modifications are whitespace formatting and adding license or copyright headers.
* Only a subset of source code from the OSS component is present. If you have only a selection of the overall number of files, and this is a small percentage of the overall number of files in the component, Snyk cannot match them to a component from the Snyk database.
* Symlinks are not followed when collecting files for hashing. However, a Linux source package unzipped in Windows will usually have in-package symlinks replaced by _copies_ of linked files, thus creating a situation where the Windows representation is different from the original source. If the difference is too large, this can lead to Snyk not detecting it.
* The source code of the OSS components is too new. The Snyk database is refreshed twice a month, but it takes time for the latest releases to get processed.
