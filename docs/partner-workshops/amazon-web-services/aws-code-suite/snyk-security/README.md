# Snyk Security

## Introduction

This module is designed to introduce scanning **open source dependencies** of the application and the **Docker** container that is created during this workshop. The files required for this module have already been created and reside in the /modules/snyk folder. You will copy these over while following the steps below.

### Background

**Snyk** is a **SaaS** offering that organizations use to **find, fix, prevent and monitor** open source dependencies. Snyk is a developer first platform that can be easily integrated into the Software Development Lifecycle \(SDLC\).

At this point of the module, the `Petstore` application is created, so we will look to insert **Snyk** as part of an important security gate during the build process.

This module will demonstrate how to fail a build when high severity issues are found so that remediation can take place.

### Snyk CLI

The Snyk command line interface \(CLI\) has three key commands for this exercise:

* `snyk auth` which links the CLI to your account and authorizes it to perform tests.
  * We will utilize Amazon's System Manager Parameters to store this token to avoid hard-coding tokens.
  * Alternatively to using **snyk auth**, you can also set an **environment variable**  `SNYK_TOKEN` which the CLI will automatically detect.
* `snyk test` performs the actual test and can fail a build
* `snyk monitor` posts a snapshot for continuous monitoring and reporting on the `snyk.io` interface where you created your account.

### Tasks for this Lab

For the purposes of this module, **Snyk** will be inserted into two key processes, when the **application** is being built, and when the **Docker container** is created.

This module has five sections: 

1. Obtain a token for testing from `https://snyk.io/` 
2. Setting up application scanning
3. Setting up Docker analysis
4. Running a test and fixing the issue\(s\) detected 
5. Viewing the status on the Snyk dashboard

