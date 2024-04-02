# Import Project repository

## How importing repositories works

After you select repositories on the Snyk Web UI and click the **Add selected repositories** button, the import starts, and a progress bar appears on the Projects page.

<figure><img src="../../.gitbook/assets/importing projects.png" alt="Importing repositories progress bar"><figcaption><p>Importing repositories progress bar</p></figcaption></figure>

When the import is finished, a confirmation message appears on the Projects page. Your imported repositories appear as separate Target folders on the Projects page. Each Target folder has the name of your Git repository account and the imported repository and contains the Snyk Projects that were created for it.

<figure><img src="../../.gitbook/assets/target folders.png" alt="Target folders on Projects page"><figcaption><p>Target folders on Projects page</p></figcaption></figure>

If some of the files in the selected repositories were not imported, you receive a notification about the Projects that were not imported.&#x20;

<figure><img src="../../.gitbook/assets/import_failed.png" alt="Project could not be created notification"><figcaption><p>Project could not be created notification</p></figcaption></figure>

## Next steps?

* [Exclude directories and files from Project import](excluding-directories-and-files-from-the-import-process.md)
* [Remove imported repository from a Project](remove-imported-repository.md)
