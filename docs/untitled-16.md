# Could not determine version for dependencies

We always try to extract all possible dependencies from project files, but sometimes we can't find the relevant information we need in order to determine what the exact dependency version this project has. We only process manifest files so any specially configured version & variables may be outside our reach.

**Some common issues include:**

1. Dependencies relying on a version that is configured as a variable in an external file like \*.props.

2. Sometimes we have to resolve the dependency version by reaching out to an external API like npmjs.com and follow that ecosystem's dependency resolutions rules, occasionally we won't be able to determine the version due to a combination of the rules / the api being unreachable / dependency being unknown on that version of the API.

