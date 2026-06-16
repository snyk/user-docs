# CLI key features for Snyk API & Web

Learn about the CLI key features of Snyk API & Web.

## Overview

The features provided by the CLI allow you to execute faster and automated operations with Snyk API & Web independently of the platform you're working on (Windows, macOS, or Linux), or if integrating with your CI/CD pipelines.

The CLI allows you to execute operations on your targets, scans, and findings. You can do it one by one or in bulk by providing the identifiers or applying a filter. For example, you can pause a single scan or a group of scans given by specific criteria (for example, all scans that are running).

## Targets

For your targets, you can learn the available operations using the command `probely targets -h`. They are basically the following:

* **List targets in your account.** To indicate the targets to list, you can provide the identifiers or apply a filter. Use `probely targets get -h` to see all the possible options.
* **Add targets to your account.** Use `probely targets add -h` to see all the possible options.
* **Update targets in your account.** To indicate the targets to update, you can provide the identifiers or apply a filter. Use `probely targets update -h` to see all the possible options.
* **Delete targets from your account.** To indicate the targets to delete, you can provide the identifiers or apply a filter. Use `probely targets delete -h` to see all the possible options.
* **Start scans on your targets.** To indicate the targets to start the scan, you can provide the identifiers or apply a filter. Use `probely targets start-scan -h` to see all the possible options.

## Scans

For your scans, you can learn the available operations using the command `probely scans -h`. They are basically the following:

* **List scans in your account.** To indicate the scans to list, you can provide the identifiers or apply a filter. Use `probely scans get -h` to see all the possible options.
* **Pause running scans.** To indicate the scans to pause, you can provide the identifiers or apply a filter. Use `probely scans pause -h` to see all the possible options.
* **Cancel running scans.** To indicate the scans to cancel, you can provide the identifiers or apply a filter. Use `probely scans cancel -h` to see all the possible options.
* **Resume paused scans.** To indicate the scans to resume, you can provide the identifiers or apply a filter. Use `probely scans resume -h` to see all the possible options.

## Findings

For your findings, you can learn the available operations using the command `probely findings -h`. They are basically the following:

* **List findings in your account.** To indicate the findings to list, you can provide the identifiers or apply a filter. Use `probely findings get -h` to see all the possible options.

## Getting started

To run these commands, you need to install the CLI and get an authorization token to use them. Check out these topics in the Developers Portal to learn more about it:

* Quick Start
* Authentication
