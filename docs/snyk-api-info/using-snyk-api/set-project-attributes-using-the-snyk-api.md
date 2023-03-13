# Set project attributes using the Snyk API

By using the Snyk API v1 endpoint [Applying attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes) you can set attributes for Snyk Projects including business criticality, lifecycle stage, and environment once the project has been created . To do so:

* Import the project using the Snyk API v1 endpoint [Import targets](https://snyk.docs.apiary.io/#reference/import-projects/import/import-targets).
* Get the status API ID from Import targets.
* Poll using [Import job details](https://snyk.docs.apiary.io/#reference/import-projects/import-job/get-import-job-details) until all imports have completed.
* Parse the project IDs from the projectURL field.
* Use the [Applying attributes](https://snyk.docs.apiary.io/#reference/projects/project-attributes/applying-attributes) endpoint to set the project attributes.
