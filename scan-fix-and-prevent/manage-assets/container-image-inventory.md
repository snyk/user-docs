# Container image inventory

The Container image inventory is a unified view in the Snyk platform that consolidates all of your container images into a single, manageable list, regardless of where they were scanned.

Instead of navigating a fragmented set of per-scan Projects, you get one authoritative inventory of unique container image assets, each identified by `Registry + Repository + Config Digest`. A single image scanned from a registry import, from the CLI, and a Kubernetes workload all appear as **one asset** with a single, deduplicated issue count.

| Capability | Description |
| :--- | :--- |
| **Unified asset list** | View all unique container images across your Organization or Group in one place |
| **Deduplicated issue counts** | Snyk merges issues from multiple scan sources — no more inflated counts |
| **Version grouping** | Group images by repository to explore the full version history of an image |
| **Filter** | Filter by registry, image repository, tag, image labels, digest, class, and more |
| **Search** | Quickly find an image by matching against the beginning of the asset name |
| **Base image fix recommendations** | View base image upgrade recommendations with impact analysis |
| **Related Projects** | See all Snyk Projects linked to a single container image asset in one place |

## How it works

Container image inventory identifies each unique image by its `Registry + Repository + Config Digest`. Because this identity is based on the immutable config digest rather than a mutable tag, one image scanned from the CLI, a container registry, and a Kubernetes workload appears as a single asset — not three separate Projects.

Issues from all scan sources are merged and deduplicated, so you see one count per unique vulnerability rather than inflated totals from overlapping scans. Snyk groups images by image repository (registry and repository), giving you a version history view where you can compare build dates, risk scores, and issue counts across digests and spot regressions over time. Each asset also surfaces key metadata — tags, inferred base image, test surface, and last scan date — in one place.

## What you need to do

Depending on how you scan containers, take the following steps to ensure your images appear in the inventory:

- **CLI users** — Upgrade to Snyk CLI version 1.1303.0 or later (which bundles an updated snyk-docker-plugin) and re-run `snyk container monitor` for your images.
- **Container Registry integrations** — Newly imported images automatically populate the inventory. Existing Projects appear in the inventory when they are retested, either manually from the UI or on a recurring test schedule.
- **Kubernetes (`snyk-monitor`)** — Upgrade `snyk-monitor` in your cluster to a version bundled with the updated snyk-docker-plugin, then redeploy your application to the cluster.

{% hint style="warning" %}
Not all existing Projects have the metadata required to compute the new asset identity. To fully populate the inventory for existing Projects, CLI and Kubernetes users must upgrade and re-scan.
{% endhint %}

## Use Container image inventory

Container image inventory is accessible from the **Inventory** navigation item at both the Organization and Group levels.

### Navigate to the inventory

Select **Inventory** in the left-hand navigation to access Container image inventory.

**At the Organization level**, selecting Inventory opens the new Container image inventory directly. There is no previous inventory experience at the Organization level, so the new view is the only one available. The **Container Images** tab is selected by default, showing all unique container image assets within that Organization.

**At the Group level**, the existing Asset Inventory remains the default view. A banner at the top of the page invites you to try the new container image experience. Click **Try it now** to switch. You can switch back and forth between the existing Asset Inventory and the new Container image inventory at any time — both views remain available. Over time, Snyk expects to add more asset types to the new view, and the older experience will eventually be deprecated.

### Grouped view (default)

By default, the inventory groups assets by **Image repository** (Registry + Repository). Each group row shows the repository name, the number of images within the group, and aggregated metadata from the most recent image in that group (by build date), including risk score and issue counts broken down by severity.

![The default grouped inventory view. Each repository row shows the number of images, latest build date, risk score, and issue severity breakdown.](../.gitbook/assets/container-inventory-grouped-view.png)

{% hint style="info" %}
Risk score is shown in both the asset table and the issues table only if your Group or Organization has enabled **Risk score** under **Snyk Preview** settings and your Projects have been re-scanned since it was enabled. If risk score is not enabled, the score columns will not be populated.
{% endhint %}

By default, Snyk sorts groups by most recent build date. You can expand any group to reveal the individual image assets within it. Each asset row within the group displays:

| Column | Description |
| :--- | :--- |
| **Asset name** | Registry, repository, and a short config digest identifier (for example, `alpine-base@a23a90a4`) |
| **Class** | The asset classification (A through D). This is a configurable value that you can change from the Overview tab once you click into an asset. |
| **Image tags** | The distinct set of tags seen across all discovery sources for this asset |
| **Build date** | The date the image was built |
| **Score** | The maximum risk score across all related Snyk Project discovery sources |
| **Issues** | The summed issue counts across discovery sources, broken down by severity (critical, high, medium, low) |
| **Coverage** | Security testing and scan engine coverage |
| **Test surface** | The distinct set of test surfaces (for example, CLI, container registry, Kubernetes) |
| **Last scan** | The most recent scan timestamp across all related discovery sources |

In an expanded group, Snyk sorts assets by build date (newest first) by default so you can easily identify the most recent version.

![An expanded repository group for "alpine-base" showing four image versions.](../.gitbook/assets/container-inventory-expanded-group.png)

#### Sort the grouped view

The grouped view uses a dual sort selector which lets you order the repository groups and the assets inside those groups independently. Click the sort control in the top-right corner to open it. The panel has two columns:

- **Sort Image Repositories** — controls the order of the group rows.
- **Sort Assets** — controls the order of the image assets within each expanded group.

The following options are available for repository groups:

| Option | Description |
| :--- | :--- |
| **Follow asset sorting** (default) | Orders repositories using the currently selected asset sort, applied to each repository's latest-built image |
| **Name** | Orders repositories alphabetically (A-Z or Z-A) |
| **Count** | Orders repositories by the number of assets they contain |
| **Last seen** | Orders repositories by when they were last seen |

You can sort assets within a group by build date, risk score (if enabled), discovered, last updated, or total issues.

### Flat (ungrouped) view

To see all assets in a single flat list instead of grouped by repository, click **Group by** and select **None**. This displays every individual image asset as its own row, with the full set of columns visible.

![The ungrouped flat view showing individual assets with all columns.](../.gitbook/assets/container-inventory-flat-view.png)

You can sort the flat view by build date, score, issue count, last scan, class, discovered, or updated — in ascending or descending order — using the sort control in the top-right corner.

### Filter and search

Click **Add filter** to open the filter panel. Filters combine using AND logic. The following filter dimensions are available:

| Filter | Description |
| :--- | :--- |
| **Asset name** | Filter by the full asset name (`registry/repository@digest`) |
| **Class** | Filter by asset classification (A, B, C, D) |
| **Config digest** | Filter by config digest |
| **Image labels** | Filter by image label key/value pairs |
| **Image tag** | Filter by image tags |
| **Index digest** | Filter by index digest |
| **Manifest digest** | Filter by manifest digest |
| **Registry** | Filter by container registry hostname |
| **Repository** | Filter by image repository |

![The filter panel showing all available filter dimensions.](../.gitbook/assets/container-inventory-filter-panel.png)

A **search bar** is also available in the top-right corner. Search uses prefix matching against the asset name field: it matches only from the beginning of the string, so entering text that appears in the middle of an asset name will not return a match. The asset name field includes the registry and repository (where present), so you must search from the start of that prefix — for example, `docker.io/snyk/kubernetes` matches the `kubernetes-monitor` asset, but `kubernetes` on its own does not.

{% hint style="info" %}
Because the search bar matches from the beginning of the asset name field (including the registry and repository prefix), searching for a partial name or tag from the middle of a string will not find it. If you cannot find a specific image, try using the **Repository** or **Image tag** filters instead.
{% endhint %}

### Asset details

Click any asset row to open a side drawer with detailed information. The drawer contains five tabs: **Overview**, **Image content**, **Issues**, **Fix Recommendations**, and **Related Projects**.

#### Overview tab

The Overview tab is split into two sections.

**Asset Information** (left side) displays general metadata:

- **Asset ID** — The unique identifier for this asset
- **Type** — The asset type (Container Image)
- **Class** — The risk classification (A through D). This field is editable; see [Change an asset's class](container-image-inventory.md#change-an-assets-class).
- **Discovered** — When Snyk first identified this asset
- **Last seen** — When Snyk last confirmed the asset exists in your environment
- **Last tested** — When Snyk last scanned the asset
- **Source** — The origin of the scan that created the asset (for example, docker.io or localhost for a CLI scan)
- **Test surface** — The distinct set of test surfaces for this asset (for example, CLI, container registry, Kubernetes)

**Container Image Details** (right side) displays image-specific metadata:

- **Index digest** — The OCI index digest (if applicable)
- **Manifest digest** — The manifest digest, usable with `docker pull`
- **Size** — The size of the image
- **Config digest** — The immutable config digest that forms part of the asset identity
- **Build date** — When the image was built
- **Platform** — The OS and architecture (for example, `linux / amd64`)
- **Base image** — The base image used, inferred by Snyk
- **Image tags** — All tags associated with this image across discovery sources
- **Image labels** — Labels set in the Dockerfile or image config

A **Security testing and coverage** section at the bottom shows which scan engines have tested this asset and when.

![The Overview tab showing asset information, container image details, and security testing coverage.](../.gitbook/assets/container-inventory-overview-tab.png)

##### Change an asset's class

Class is a label you assign to an asset, from A to D, using whatever criteria suits your organization (for example, business criticality). Because you can filter and sort the inventory by class, it gives you a way to slice your images along a dimension that you control.

To edit asset class, select the **Class** value in the **Overview** tab to open a dropdown, then select **A**, **B**, **C**, or **D**. The change is saved immediately and is reflected anywhere the asset's class is shown. Images that have not been classified are class **C** by default.

#### Image content tab

The Image content tab shows the provenance attestations Snyk found for the image. A provenance attestation is a signed record, produced by your build system, of where an image came from — the source repository, the commit, and how it was built. Snyk reads these attestations during a scan and displays them alongside the image.

An image can have more than one attestation; the most recent is marked **Latest**. Each card shows:

- **Source repository** — The repository the image was built from, linked to the provider where available. Images built outside a repository show `localhost`.
- **Source commit** — The commit the build used, linked where available, with a copy button
- **Built** — The build date
- **Builder** — The identity of the build system that produced the attestation
- **Build type** — The build definition the builder ran, for example `BuildKit`
- **Attests manifest** — The manifest digest this attestation covers, which ties the attestation to a specific image
- **Dockerfile** — The Dockerfile the build used, if the attestation records it, with a copy button

![The Image content tab showing a provenance attestation, including source repository, build details, and Dockerfile.](../.gitbook/assets/container-inventory-image-content-tab.png)

#### Issues tab

The Issues tab shows all deduplicated vulnerabilities for the asset. Summary cards at the top provide:

- **Open Issues** — The total number of unique issues across all scan sources, not including ones that were ignored
- **Critical Issues** — The number of critical-severity issues and how many packages are affected
- **Fixable Issues** — Issues that can be resolved through updates or patches
- **Exploitable Issues** — Issues with known exploits

Below the summary, a table lists each issue with the following columns:

| Column | Description |
| :--- | :--- |
| **Severity** | The severity level (critical, high, medium, low) |
| **Score** | The risk score for this specific issue |
| **Issue** | The vulnerability name and CVE identifier |
| **Affected package** | The package introducing the vulnerability |
| **Exploitable** | Whether the issue has a known exploit |
| **Fixable** | Whether the issue can be fixed through an upgrade |
| **Test surface** | Which scan sources identified this issue |

![The Issues tab showing total issues with severity breakdown and vulnerability details.](../.gitbook/assets/container-inventory-issues-tab.png)

#### Fix recommendations tab

The Fix Recommendations tab displays base image upgrade recommendations. Each recommendation shows:

- **The upgrade path** — For example, "Upgrade from alpine:3.12.1 to alpine:3.24.1"
- **Impact summary** — The number of issues that would be resolved and the upgrade type (for example, "26 fewer issues - Minor upgrade")
- **Side-by-side comparison** — The current base image and suggested base image with their respective issue counts and severity breakdowns

Opening a fix PR directly from a base image recommendation is not available in this release. To open a fix PR for a specific package upgrade, use the associated Snyk Project.

![The Fix Recommendations tab showing an upgrade path with a side-by-side comparison of issue counts.](../.gitbook/assets/container-inventory-fix-recommendations-tab.png)

#### Related Projects tab

The Related Projects tab lists all Snyk Projects that are linked to this asset as discovery sources. Projects are grouped by target. The table includes:

| Column | Description |
| :--- | :--- |
| **Project** | The Snyk Project name, with a "Most Recent" badge on the latest-tested project |
| **Last Tested** | When the Project was last scanned |
| **Issues** | Issue counts for this specific Project |
| **Test Surface** | The scan source for this Project (for example, CLI, container registry) |
| **Target Reference** | The target reference associated with the Project |
| **Actions** | Links to the Project details page |

You can sort the list by date last tested or issue count, and select **Modify columns** to customize which columns are displayed.

![The Related Projects tab showing linked Snyk Projects with metadata and actions.](../.gitbook/assets/container-inventory-related-projects-tab.png)

## How it impacts your workflow

### AppSec and security engineers

- Get a consolidated inventory of all unique container image assets and their aggregate risk scores.
- Track the security posture of images over time by exploring version history within a repository grouping.
- Report on real issue counts — no more inflated numbers from duplicate scan sources.

### Developers

- Find the image you own quickly using search and filter (by tag, registry, repository, or label).
- Understand a vulnerability's full context in one place rather than chasing it across multiple Projects.
- Take direct action: view fix recommendations for base image upgrades without leaving the inventory.

### Administrators

- The inventory is available at both **Organization** and **Group** scopes, giving the right visibility level to the right team.
- At the Organization level, the new Container image inventory replaces the Inventory page (there was no prior inventory experience at this scope). At the Group level, the existing Asset Inventory remains available alongside the new experience — users can switch between them.
- Access requires **Snyk Essentials** enabled at the Group level, currently available only on an Enterprise plan.

## Prerequisites

Container image inventory requires **Snyk Essentials** at the Group level. This is available to customers on an Enterprise plan; it is not included in the self-serve Enterprise trial.

- **Scope:** Organization and Group level
- **Interface:** Accessible from the **Inventory** navigation item, under the **Container Images** tab
- **Organization level:** The new Container image inventory is the only inventory experience at this scope.
- **Group level:** The existing Asset Inventory remains the default. A banner allows users to switch to the new container image experience and back at any time.
- **Provenance attestations:** Requires Snyk CLI v1.1307.0 or later, or `snyk-monitor` v2.23.24 or later. Images scanned by earlier versions show no attestations until they are scanned again.

The existing Projects view for containers remains unchanged.

## Known limitations

| Limitation | Details |
| :--- | :--- |
| **Search scope** | The search bar matches from the beginning of the asset name field. Searching for a partial image name or tag within a longer string may not return expected results. Use the **Repository** or **Image tag** filters for more precise lookups. |
| **CLI scans without a registry hostname** | Images scanned directly from a tar file (`snyk container test image.tar`) do not associate with their registry counterpart because the registry field is null. |
| **Tag staleness** | Because tags are scoped per discovery source, two assets in the same repository can temporarily show the same tag if a discovery source has not yet been refreshed. This can cause the inventory to show images that no longer carry the `latest` tag in your Projects view. |
| **Backfilling existing Projects** | Not all existing Projects have the metadata required to compute the new asset identity. CLI and Kubernetes users must upgrade to the latest Snyk CLI or `snyk-monitor` and re-scan to populate the inventory for existing images. |
| **Base image inference** | Base image detection is heuristic-based (parsing the Dockerfile if present, or matching layer hashes against a known image index). Results may vary across Projects that scanned the same image differently (for example, with or without the Dockerfile). |
| **Asset class sync** | Changing an asset's class in the new Container image inventory does not propagate to the old inventory. The two inventories can therefore show different class values for the same asset. |
| **Provenance attestations require a re-scan** | Attestations are read during a scan. Images scanned before provenance support was released show no attestations until they are scanned again, and only images whose build system produces attestations will populate this tab. |

## FAQs

**Does this replace the Projects page for containers?**

No. Container image inventory is a new, complementary view. The Projects page remains fully functional. The inventory provides a unified, deduplicated view of your container assets on top of the underlying Projects.

**Will all my existing container Projects automatically appear in the inventory?**

Not immediately. Snyk creates assets when a scan runs using the updated snyk-docker-plugin. Existing Projects populate the inventory as images are rescanned — either through a scheduled recurring test or when a new image digest is pushed to a monitored tag. CLI and Kubernetes users may need to take additional steps (see [What you need to do](container-image-inventory.md#what-you-need-to-do)).

**What is the asset identity for a container image?**

An asset is uniquely identified by `Registry + Repository + Config Digest`. Image tags are intentionally excluded from the identity because tags are mutable — a single image can have many tags, but it is still one asset.

**Is this available at the Group level?**

Yes. Container image inventory is available at both Organization and Group scopes. The data shown reflects only the discovery sources within the selected scope.

**How are issue counts calculated?**

Snyk sums issue counts across the latest scan snapshot of each unique Project target file associated with the asset, then breaks them down by severity. This deduplication ensures that the same vulnerability found by multiple scan sources is not double-counted.

**How is the risk score calculated?**

The score is the maximum score across the latest scan snapshots of all related Snyk Project discovery sources. Risk score is available only after you enable it under **Snyk Preview** settings for your Group or Organization and re-scan your Projects.
