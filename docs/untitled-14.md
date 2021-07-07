# What does Snyk access and store when scanning a project?

##  What does Snyk access and store when scanning a project?

You can use Snyk to scan for vulnerabilities either locally via our CLI App, or through our GitHub, BitBucket, GitLab, or PaaS and Serverless integrations \(check Snyk's up-to-date [list of integrations\)](https://snyk.io/docs/).

We do not access or store any knowledge about your source code \(with the exception of [reachability analysis](https://support.snyk.io/hc/en-us/articles/360010554837-Reachable-Vulnerabilities-)\), but rather we only access and read your project's manifest files \([see the full list](https://support.snyk.io/hc/en-us/articles/360000911957-Language-support)\) to build a dependency tree that we can use to query against our database of active vulnerabilities. We also read Dockerfiles to detect base image vulnerabilities.

[How does Snyk handle your data?](https://snyk.io/wp-content/uploads/Snyk-dataflows-How-Snyk-handles-customer-data.pdf)

When integrating to a container registry, Snyk pulls the image and detects packages and binaries. This is done to find security vulnerabilities in your container image against the Snyk Vulnerability Database and to return those security results.

If you have enabled Snyk Infrastructure as Code, we will detect configuration files _\(currently YAML & Terraform\)_ and scan those for any issues. The files are not kept by Snyk, only a reference to the file is stored. When we display any detected issues in the Snyk UI, we reload the contents of the file directly from source control.  

A record of your latest dependencies is kept so we can notify you when a dependency is affected by a newly disclosed vulnerability.

More details can be found in section 9 of our [terms and conditions](https://snyk.io/policies).

