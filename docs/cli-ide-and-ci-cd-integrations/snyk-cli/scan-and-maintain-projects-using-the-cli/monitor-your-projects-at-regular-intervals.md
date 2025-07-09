# Monitor your projects at regular intervals

By using the `test` command and the `@snyk/protect` [package](https://github.com/snyk/snyk/tree/master/packages/snyk-protect) (replaced the snyk `protect` command), you are well set up to address known vulnerabilities. To address new vulnerabilities, which are constantly disclosed, use snyk monitor.

Just before you deploy, run `snyk monitor` in your project directory:

`cd ~/projects/myproject/`\
`snyk monitor`

This sends a snapshot of your current dependencies to your Snyk dashboard, so Snyk can notify you about newly disclosed vulnerabilities in those dependencies, or when a previously unavailable patch or upgrade path is created. If you take multiple snapshots of the same project, Snyk alerts you only to new information about the latest snapshot.

Log in to the Snyk Web UI, and go to your organization's project page ([https://app.snyk.io/monitor/](https://app.snyk.io/monitor/)) to see the latest snapshot and history of your project.
