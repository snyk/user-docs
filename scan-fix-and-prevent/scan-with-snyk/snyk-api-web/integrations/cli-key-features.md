# CLI key features for Snyk API & Web

Learn about the CLI key features of Snyk API & Web.

## Overview

The CLI lets you run faster, automated operations with Snyk API & Web, independently of the platform you work on (Windows, macOS, or Linux) or when integrating with your CI/CD pipelines.

You can run operations on your targets, scans, and findings one by one or in bulk by providing the identifiers or applying a filter. For example, you can pause a single scan or a group of scans that match specific criteria, such as all running scans.

## Targets

To learn the available operations for your targets, run the command `probely targets -h`. The operations are the following:

* List targets in your account. To indicate the targets to list, provide the identifiers or apply a filter. Run `probely targets get -h` to see all the options.
* Add targets to your account. Run `probely targets add -h` to see all the options.
* Update targets in your account. To indicate the targets to update, provide the identifiers or apply a filter. Run `probely targets update -h` to see all the options.
* Delete targets from your account. To indicate the targets to delete, provide the identifiers or apply a filter. Run `probely targets delete -h` to see all the options.
* Start scans on your targets. To indicate the targets to scan, provide the identifiers or apply a filter. Run `probely targets start-scan -h` to see all the options.

## Scans

To learn the available operations for your scans, run the command `probely scans -h`. The operations are the following:

* List scans in your account. To indicate the scans to list, provide the identifiers or apply a filter. Run `probely scans get -h` to see all the options.
* Pause running scans. To indicate the scans to pause, provide the identifiers or apply a filter. Run `probely scans pause -h` to see all the options.
* Cancel running scans. To indicate the scans to cancel, provide the identifiers or apply a filter. Run `probely scans cancel -h` to see all the options.
* Resume paused scans. To indicate the scans to resume, provide the identifiers or apply a filter. Run `probely scans resume -h` to see all the options.

## Findings

To learn the available operations for your findings, run the command `probely findings -h`. The operations are the following:

* List findings in your account. To indicate the findings to list, provide the identifiers or apply a filter. Run `probely findings get -h` to see all the options.

## Getting started

To run these commands, you must install the CLI and get an authorization token. Visit these topics in the Developers Portal to learn more:

* Quick Start
* Authentication
