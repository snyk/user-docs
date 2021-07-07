# Project takes too long to import into the Snyk UI \(Yarn Workspaces\)

####  [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

###  [ ]() <a id="category-name"></a>

#### Importing projects

* [ Project import errors](/hc/en-us/articles/360001373118-Project-import-errors)
* [ Monitored projects are overwriting each other in the UI](/hc/en-us/articles/360017785398-Monitored-projects-are-overwriting-each-other-in-the-UI)
* [ How can I find which projects imported into Snyk have been archived in Github?](/hc/en-us/articles/360007937497-How-can-I-find-which-projects-imported-into-Snyk-have-been-archived-in-Github-)
* [ What are docker scratch based images?](/hc/en-us/articles/360004012857-What-are-docker-scratch-based-images-)
* [ My project appears as /file://nothing in Snyk](/hc/en-us/articles/360003897778-My-project-appears-as-file-nothing-in-Snyk)
* [ Project takes too long to import into the Snyk UI \(Yarn Workspaces\)](/hc/en-us/articles/360002865538-Project-takes-too-long-to-import-into-the-Snyk-UI-Yarn-Workspaces-)
* [ Project owner and importer information](/hc/en-us/articles/360002827197-Project-owner-and-importer-information)

1.  [Docs Library \| Snyk](/hc/en-us)
2.  [FAQs](/hc/en-us/categories/360000116697-FAQs)
3.  [Importing projects](/hc/en-us/sections/360000923478-Importing-projects)

##  Project takes too long to import into the Snyk UI \(Yarn Workspaces\)

At times it may take a very long time for your project to import into the Snyk User Interface. We have seen these cases with `Yarn workspaces`. You can check if its a `Yarn Workspace` project by inspecting your `package.json` and finding the keyword `workspaces`

 **Reason:** This is case specific and we build the tree by talking to the vulnerabilties which is an extremely slow process

 **Workaround:** As a workaround you can use the Snyk API[ \(here\)](https://snyk.docs.apiary.io/#reference/test/yarn/test-package.json-&-yarn.lock-file) and send `yarn.lock` with every `package.json` file

* 
