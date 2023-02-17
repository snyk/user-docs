# Run an analysis with Visual Studio Code extension

In the IDE note that the extension is already picking up the files and uploading them for analysis.

Snyk Open Source requires the Snyk CLI, so it downloads in the background.

Snyk Code analysis runs quickly without the CLI, so results may already be available. Otherwise, you see the following screen while Snyk scans your workspace for vulnerabilities:

![Snyk Code scan](<../../.gitbook/assets/image (134) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png>)

Snyk analysis runs automatically when you open a folder or workspace.

* Snyk Code performs scans automatically on file saves (can be disabled by setting **Scanning Mode** to **manual**).
* Snyk Infrastructure as Code scans automatically on file saves (can be disabled by setting **Scanning Mode** to **manual**).
* Snyk Open Source does not automatically run on save by default, but you can enable it in settings:

![Snyk Open Source settings](<../../.gitbook/assets/image (143) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (5).png>)

**Tip**: if you do not like to manually save while working, enable [AutoSave](https://code.visualstudio.com/docs/editor/codebasics#\_save-auto-save).

## Rescan

To manually trigger a scan, either Save or manually rescan using the rescan icon:

![Rescan icon](<../../.gitbook/assets/image (120) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png>)

If you only need the Code Quality, Code Security, or Open Source Security portion of the findings, you can easily disable the feature with the results you do not want to see or collapse the view:

![Configure Features view](../../.gitbook/assets/configure-features.png)
