# Securing data at rest

The Snyk CLI stores its configuration in a JSON file in the local file system in a user-related path. Because the configuration file might include secrets like tokens, Snyk recommends that you secure the stored file.

Snyk recommends the following practices:

* Use ACLs to limit access to the [file path of the configuration file](access-requirements.md). The operating system might provide this.
* Use Hard Disk Encryption to secure persisted data (data at rest).
