# Fix PR failed for Bitbucket Cloud integration

## Fix PR failed for Bitbucket Cloud integration

A fix PR opened from within Snyk Fails without an obvious error, you may receive a notification email saying failed to create automatic fix PR.

This may be due to a setting within Bitbucket cloud, 'Require issue keys in commit messages'. If you have this enabled for a repository in Bitbucket cloud then any PR for that repo from Snyk will fail. This option requires linked issue keys in commits, see [here](https://support.atlassian.com/bitbucket-cloud/docs/link-to-a-web-service/#Linktoawebservice-Requirelinkedissuekeysincommits) for more information. This is something we can't currently support so if you are experiencing the error and want to be able to open fix PRs then you will need to disable this option within Bitbucket.

If this option isn't enabled or you still get an error after turning it off then please contact Snyk [support](https://support.snyk.io/hc/en-us/requests/new) for more help.

