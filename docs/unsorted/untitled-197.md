# How can I find which projects imported into Snyk have been archived in Github?

* [ Project import errors](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360001373118-Project-import-errors/README.md)
* [ Monitored projects are overwriting each other in the UI](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360017785398-Monitored-projects-are-overwriting-each-other-in-the-UI/README.md)
* [ How can I find which projects imported into Snyk have been archived in Github?](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360007937497-How-can-I-find-which-projects-imported-into-Snyk-have-been-archived-in-Github-/README.md)
* [ What are docker scratch based images?](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004012857-What-are-docker-scratch-based-images-/README.md)
* [ My project appears as /file://nothing in Snyk](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003897778-My-project-appears-as-file-nothing-in-Snyk/README.md)
* [ Project takes too long to import into the Snyk UI \(Yarn Workspaces\)](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360002865538-Project-takes-too-long-to-import-into-the-Snyk-UI-Yarn-Workspaces-/README.md)
* [ Project owner and importer information](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360002827197-Project-owner-and-importer-information/README.md)

## How can I find which projects imported into Snyk have been archived in Github?

Thanks to **Jake Champion** for sharing this great tip!

Projects that have been archived in Github but that have already been imported into Snyk will continue to notify you of vulnerabilities. Some users may want to remove Archived projects from your Snyk Organization in bulk to prevent unwanted notifications or too much bloat in your Organization.

Currently, Snyk does not have a feature to do this in an automated fashion, but you can leverage Chrome's developer tools to get a list of all Imported \(in Snyk\) and Archived \(in Github\) projects as follows:

1. Navigate to your Org, and click on **Add Project** &gt; **Github**
2. Expand the relevant Github Repo\(s\) and launch developer tools \(**Command** + **Option** + **J** on Mac\)
3. Click **Console** and paste:

```text
Array.from(document.querySelectorAll('label [d="M3.5 7.52h17V20.1a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1V7.52zm6.5 4.53h4a1 1 0 0 0 1-1 1 1 0 0 0-1-1.01h-4a1 1 0 0 0 0 2.01zM21 3a1 1 0 0 1 1 1v2.02a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4A1 1 0 0 1 3 3h18z"]')).filter(a => a.parentElement.parentElement.parentElement.parentElement.parentElement.parentElement.querySelectorAll('[d="M9.3 19.7L3 13.4l1.9-2L9.3 16 20.1 5.1l1.9 2z"]').length).map(a=>a.parentElement.parentElement.parentElement.parentElement.parentElement.parentElement.querySelector('span').innerText)
```

1. Press Enter

You'll be presented with a list of Orgs that are Imported and Archived:

