---
description: How to scan a Flat/Namespace monorepo with Snyk, where peer folders rely on naming conventions
nav_context: agnostic
---

# Flat/Namespace monorepo

```
Root/
├── main-ci-pipeline.yml       # Orchestrator: runs "npm build" on changed folders
├── @app-dashboard/            # Micro frontend (local package.json build script)
├── @app-mobile-api/           # Application (local package.json build script)
├── @lib-ui-core/              # Internal package
├── @lib-utils/                # Internal package
├── @tool-dev-proxy/           # Local dev tool
└── pnpm-workspace.yaml        # Defines workspace members
```

## Layout

* **Organized by naming convention**: No nested folders. Every project is a peer at the root level.
* **Naming as metadata**: Prefixes such as `@app-` or `@lib-` describe the purpose of a folder at a glance.
* **Generic task runner**: The pipeline loops over every changed folder and runs a standard command, such as `npm run build`, if the folder contains a manifest.
* **Shared source code**: Source code is likely shared across folders. Because all projects are peers, `@app-` folders frequently import logic directly from `@lib-` folders.
* **Typical consumption model**: Raw source code.

## Snyk best practices

* **Root-level execution for Snyk Code**: Run `snyk code test` from the root directory only if source code is shared between workspaces. Snyk must resolve workspace aliases, such as `@lib-utils`, and follow cross-folder data flows.
* **Full-context pull request gate for Snyk Code**: In the pull request pipeline, run a full-context scan from the root using `snyk code test`. This catches a vulnerability when a source (user input) is added to an `@app-` folder but the sink (vulnerable function) is in an unchanged `@lib-` folder.

{% hint style="info" %}
In namespace monorepos where many teams share the same root, consider a non-blocking (soft) gate for Snyk Code. Because this scan analyzes the entire repository context, a hard gate could prevent Team A from merging a clean pull request if existing vulnerabilities are in a peer folder owned by Team B. A soft gate provides visibility without forcing one team to fix another team's technical debt to unblock its own release.
{% endhint %}

* **Preventive gating for libraries**: Run a gating `snyk test` only if that specific `@lib-` or `@app-` folder changes. This blocks the pull request if a library change introduces a third-party vulnerability, which prevents poisoning all peer applications that consume the library.
* **Snyk Code reporting**: Run the following command from the root. This pushes the security state of the full repository to the Snyk web UI as a single ledger while maintaining branch context.

  ```bash
  snyk code test --report --project-name="<NAMESPACE_PREFIX>-<PROJECT>" --target-reference="$(git branch --show-current)" --org=<ORG_ID>
  ```
* **Application-level Snyk Open Source monitoring**: Run the following command for changed `@app-` folders only, on merge to `main`. Applications bundle the libraries they consume, so scanning the changed application gives the most accurate production heartbeat for that deployable.

  ```bash
  snyk monitor --project-name="<NAMESPACE_PREFIX>-<PROJECT>" --target-reference="$(git branch --show-current)" --org=<ORG_ID>
  ```
* **Unified Organization strategy**: Keep all Project data for the monorepo in a single, dedicated Snyk Organization (`ORG_ID`).
* **Namespace Project naming**: Use the folder name in the `--project-name` flag, for example `@app-dashboard`. This prevents results from different namespaces from overwriting each other in the Snyk web UI.
