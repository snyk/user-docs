# Manifest not found

##  Manifest not found

### Why couldn't Snyk find the manifest file for commit status check? Why is my PR blocked and how do I fix it?

Snyk uses your project manifest file to analyze your projects for vulnerabilities. When you import a project for monitoring, Snyk scans the project to locate the manifest file and then remembers where that file is. Snyk then scans your project again whenever a webhook event occurs and also on a regular basis based on your configurations, based on the manifest file at its originally identified location.

At the same time, your code is changing all the time, and occasionally you might even move or delete the very same files that Snyk relies on for testing and analysis of your project. When a project manifest file is moved or deleted, we still try to look for in it in the last known location in order to run tests on commit statuses. If we can't find the file, the "Could not reach target file" status is returned.

Fear not though! You can quickly fix this.

Here's how:

1. Delete the matching project from your account in the Snyk app \(UI or CLI\).
2. Now import the same project from scratch.

   As during the original import, Snyk scans the project and locates the manifest file.

