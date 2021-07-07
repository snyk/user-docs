# Badge Support for Repositories

##  Badge Support for Repositories

We currently only have support for Node.js, Ruby or Java GitHub repositories

To show a badge for a given Node.js, Ruby or Java GitHub repository, copy the relevant snippet below and replace “{username}/{repo}” with the GitHub username and repo you want to test

**HTML:**

`<*a href="https://snyk.io/test/github/{username}/{repo}">`![image2.png](https://support.snyk.io/hc/article_attachments/360004882698/uuid-2678b218-659f-61c6-8e14-57964dfb2bc6-en.png)

**Markdown:**

`[![Known Vulnerabilities](https://snyk.io/test/github/{username}/{repo}/badge.svg)]`

`(https://snyk.io/test/github/{username}/{repo})`

The badge will reflect the vulnerability state of the latest commit on the default branch. To show the vulnerability state of a specific branch, release or tag, simply add its name after the repo name in the URL.

