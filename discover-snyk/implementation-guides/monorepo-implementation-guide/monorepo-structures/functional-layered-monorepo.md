---
description: How to scan a Functional/Layered monorepo with Snyk, where deployable applications are separated from shared libraries
nav_context: agnostic
---

# Functional/Layered monorepo

```
Root/
├── pipeline-main.yml          # Uses "affected" logic to trigger builds
├── apps/                      # Deployables
│   ├── portal-web/            # Micro frontend (Webpack/Vite)
│   └── analytics-api/         # Microservice (Docker)
├── libs/                      # Shared code
│   ├── ui-shared/             # Components
│   └── db-schemas/            # Models
├── common-config/             # Shared lint and TSConfig
└── package.json               # Root workspace
```

## Layout

* **Organized by technical stack layer**: Separation of concerns based on what each folder is, for example an application versus a library.
* **One-way dependency flow**: Applications depend on libraries. Libraries generally do not depend on applications.
* **Centralized pipeline**: A single pipeline at the root uses "affected" logic to skip building projects that a change did not touch.
* **Shared source code**: Folders reference each other, so source code is likely shared between folders.
* **Typical consumption model**: Raw source code.

## Snyk best practices

* **Root-level execution for Snyk Code**: Run `snyk code test` from the root directory if source code is shared between folders. Snyk must resolve relative file paths in import statements, for example `import {..} from '../../libs/ui-shared'`. This lets Snyk follow the data flow from the application's source code into the library's source code to complete the source-to-sink analysis.
* **Dependency-aware pull request gate**: Use your orchestrator's "affected" logic to trigger Snyk scans for changed folders and their dependents.
  * **Snyk Code**: If a library or application changes, run `snyk code test` from the root for that library and all consuming applications. This validates new sinks in a library or new sources in an application against each other.
  * **Snyk Open Source**: Run a blocking `snyk test` on changed folders only. Gating a library is a preventive gate that stops vulnerable third-party packages from being merged into the shared codebase before applications inherit them. Gating an application ensures the vulnerability is not merged into `main`.

{% hint style="info" %}
If multiple teams contribute to this monorepo, consider making the Snyk Code gate non-blocking (a soft gate). Snyk Code requires a root-level scan to stay accurate, so a hard gate could block a team from merging because of a vulnerability in a shared library owned by a different team. A soft gate provides visibility and reporting without cross-team delivery bottlenecks.
{% endhint %}

* **Snyk Code reporting**: For all affected applications and libraries, run the following command. Run it from the root if source files reference each other. This pushes the full analysis to the Snyk web UI.

  ```bash
  snyk code test --report --project-name="<LAYER>/<PROJECT_NAME>" --target-reference="$(git branch --show-current)" --org=<ORG_ID>
  ```
* **Snyk Open Source monitoring**: Run the following command for affected applications only. Applications bundle libraries and their third-party dependencies at build time, so an application-level scan gives the most accurate production dependency graph and prevents duplicate alerts for shared library vulnerabilities.

  ```bash
  snyk monitor --project-name="<LAYER>/<PROJECT_NAME>" --target-reference="$(git branch --show-current)" --org=<ORG_ID>
  ```
* **Unified Organization strategy**: Keep all Project data for the monorepo in a single, dedicated Snyk Organization (`ORG_ID`).
* **Naming hierarchy**: Use the folder path in `--project-name`, for example `apps/portal-web` or `libs/ui-shared`, so that the Snyk web UI mirrors the physical and functional layout of the repository.
