---
description: Best practices for scanning a monorepo with the Snyk CLI, including ignores, reporting, Organizations, and pipeline scans
nav_context: agnostic
---

# Apply CLI best practices

If you scan your monorepo with the Snyk CLI, apply these practices to keep results accurate, organized, and fast.

## Use a unified ignore strategy

* **Snyk Code**: Use [Consistent Ignores](https://docs.snyk.io/scan-fix-and-prevent/fix/prioritize-issues-for-fixing/ignore-issues/consistent-ignores-for-snyk-code) and the [Ignore Approval Workflow](https://docs.snyk.io/scan-fix-and-prevent/fix/prioritize-issues-for-fixing/ignore-issues/ignore-approval-workflow). Ignores propagate across the CLI and IDE for all branches on subsequent rescans.
* **Snyk Open Source**: Use multiple [`.snyk` policy files](https://docs.snyk.io/scan-fix-and-prevent/prevent/policies/the-.snyk-file), with one in each folder. The `.snyk` file must be in the same folder as the manifest files.
  * Add `--policy-path=<PATH_TO_POLICY_FILE>` to Snyk Open Source commands so that ignores are applied during the scan.
  * If you use one `.snyk` policy file at the root level and then scan from a subfolder, the ignores are not applied.
  * If Snyk Open Source scans must run at the root level, use one centralized `.snyk` file.
  * Ignores propagate across the CLI and IDE, and are applied as long as the `.snyk` file exists in a branch.

## Choose the right reporting output

If you need code snippets, use [snyk-to-html](https://docs.snyk.io/developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html), the IDE, or SARIF or JSON output for triage.

## Scan from a directory that contains all referenced code

If Folder A references Folder B, run the Snyk Code scan from a directory that includes both, usually the root. Snyk needs both folders to complete the source-to-sink analysis.

## Isolate each monorepo in its own Organization

Use a dedicated Snyk Organization for each monorepo to stay within the [Project limit for your plan](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects/maximum-number-of-projects-in-an-organization). This approach:

* Simplifies rolling out CLI commands in CI pipelines, because you only need a single `ORG_ID`.
* Avoids duplicate imports of the same monorepo, which keeps reporting clean.

## Manage language and dependency resolution settings

* **Organization-level standardization**: Manage global language settings and package manager behaviors at the Organization level to ensure a uniform security baseline.
* **Project-level overrides**: Snyk supports Project-level configurations for specific languages. Use these overrides when individual Projects in the monorepo deviate from the Organization standard.
* **Strategic segmentation**: If a monorepo contains Projects that require conflicting scan configurations, and Snyk cannot differentiate them at the Project level, segment those Projects into separate Snyk Organizations. For example, Projects might use different package manager versions or resolution behaviors. Segmenting ensures the Snyk Open Source engine applies the correct resolution logic to every manifest file.

## Run scans in your pipeline based on what changed

Choose the model that matches how your pipelines are organized:

* **Centralized pipeline**: Identify the changed folders. For Snyk Open Source, `cd` into each changed folder, unless the monorepo has centralized lockfiles. For Snyk Code, stay at the root if code is shared between folders.
* **Individual pipelines**: Trigger a scan only when files in that folder change. Run the scan inside the folder only if the code in it is isolated for Snyk Code scans.

## Plan your preventive gating strategy

The goal of preventive gating in a monorepo is to prevent vulnerability drift, where a change in one shared library unknowingly creates a security hole in the applications that consume it.

Your gating logic depends on your consumption model. Identify whether your applications consume shared logic as raw source code (direct imports) or as published artifacts (versioned packages).

For examples, see [Monorepo structures](monorepo-structures/).

## Categorize your Projects

* **Tags**: Use `--tag` (for example, `team=checkout`, `env=prod`) to allow granular filtering. For details, see [Project tags](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects/project-tags).
* **Collections**: Group Projects into [collections](https://docs.snyk.io/scan-fix-and-prevent/scan-with-snyk/snyk-projects/project-collections-groupings/project-collections) by monorepo name or business domain for improved reporting.

## Handle centralized lockfiles

If a monorepo has a centralized lockfile, run `snyk test` from the root directory.

To reduce Snyk Open Source scan times, run `snyk test --all-projects --exclude=folder1,folder2`. Use your orchestrator's built-in logic to compare all folders in the repository with the changed folders, and list the unchanged folders in the `--exclude` flag.

## Scan workspace-based monorepos from the root

In monorepos that use workspace-aware package managers, such as Yarn, uv, or Rust, manifest files exist at both the root (the lockfile and workspace definition) and in individual folders (project-specific dependencies).

* **Execution level**: Run Snyk Open Source scans at the root directory. This ensures the Snyk CLI can access the centralized lockfile to resolve the full, flattened dependency tree and any workspace-level aliases.
* **Policy consolidation**: Keep a single, consolidated `.snyk` policy file at the root. This ensures that ignores and patch instructions are applied consistently across the entire workspace during the root-level scan.

## Scan Maven multi-module monorepos with the aggregate flag

For Maven multi-module monorepos, run `snyk test --maven-aggregate-project` instead of `--all-projects`. Run the scan in the same directory as the root `pom.xml`.

The `--maven-aggregate-project` flag runs a compilation step to ensure all submodules and their dependencies are fully resolved before scanning. This gives more complete results than `--all-projects`.
