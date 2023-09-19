# Securing data at rest

The Snyk CLI stores its configuration in a JSON file in the local file system in a user-related path. Because the configuration file might include secrets like tokens, it is recommended that you secure the stored file.

The following practices are recommended:

* Use ACLs to limit access to the [file path of the configuration file](access-requirements.md). This might be provided by the operating system.
* Use Hard Disk Encryption to secure persisted data (data at rest).
