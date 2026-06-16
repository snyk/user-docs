# Use seeds and reject lists

Control scan behavior by including or excluding specific areas of your target using seeds and reject lists.

## Understanding seeds and reject lists

**Seeds list**: Add areas of the target that would otherwise be hidden and not discovered during scans. This ensures scans cover everything within the target scope. For example, add administrative areas only accessible through direct links. The **seeds list applies only to Web targets**.

**Reject list**: Limit what the scanner visits on the target. This is a collection of URLs the crawler is instructed to avoid. Use this to exclude unnecessary or sensitive content such as logout options, private user data, or repetitive structures like paginated lists.

{% hint style="info" %}
The reject list takes precedence over the seeds list. Areas in the seeds list are ignored if they match areas in the reject list.
{% endhint %}


## Add paths to the seeds list

1. Navigate to **Targets** and find the target in the list.
2. Click the **gear icon** to edit the target settings.
3. Go to the **Scanner** tab.
4. In the **Seeds list** section, enter the path to include in **Add path**.
5. Click **Add**.

You can add multiple paths. Paths must be relative to the target URL.

Scans follow these paths and explore them, expanding the scan reach across the target.

## Add URLs to the reject list

1. Navigate to **Targets** and find the target in the list.
2. Click the **gear icon** to edit the target settings.
3. Navigate to the **Scanner** tab.
4. In the **Reject list** section, enter the area to exclude in **Add URL**.
5. Click **Add**.

URLs must be absolute. You can use `*` as a wildcard.

Scans do not follow these paths, avoiding these areas during target scans.

## Remove entries from lists

To remove an entry from either list, locate the entry and click the trash icon next to it.

This tailored approach ensures hidden areas are thoroughly tested while sensitive or unnecessary areas are excluded, optimizing your scanning process.
