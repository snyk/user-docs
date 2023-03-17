# Snyk API v1 Project issue paths endpoints

The following provides information in addition to the information in the [API v1 reference docs](https://snyk.docs.apiary.io/reference/projects/project-issue-paths/list-all-project-issue-paths).

The `paths` endpoints provide details of the paths through which an issue has been introduced.

**Requests** to the `paths` endpoint are `GET` requests. The endpoint is available at the following URLs:

* `https://snyk.io/api/v1/org/<orgId>/project/<projectId>/issue/<issueId>/paths` This returns the paths for an issue in the most recent test of the project.
* `https://snyk.io/api/v1/org/<orgId>/project/<projectId>/history/<snapshotId>/issue/<issueId>/paths` This returns the paths for an issue in a specific test of the project.

Both paths can take a query string allowing for pagination, for example,`?page=2&perPage=500`.

By default, the first page of 100 results is returned. Up to 1,000 results can be requested per page.

The **response** has the following structure:

```json
{
    "snapshotId": "6d5813be-7e6d-4ab8-80c2-1e3e2a454553",
    "paths": [...],
    "total": 193,
    "links": {
        "prev": "<https://snyk.io/>...",
        "next": "<https://snyk.io/>...",
        "last": "<https://snyk.io/>..."
    }
}
```

* `snapshotId` is the ID of the project snapshot from which the paths were returned.
* `total` is the total number of paths for the issue in the snapshot.
* `links` provides convenience links for navigating between pages of the response. `links.next` and `links.prev` are provided only if such a page exists. For example, when you retrieve the last page of results, it does not include a `next` link.
* `paths` is an array, each element of which is a path through the dependency tree. Each path is itself an array of package descriptors, for example:

```json
{
    "paths": [
        [
            { "name": "lodash", "version": "4.17.4", "fixVersion": "4.17.20" }
        ],
        [ 
            { "name": "babel-template", "version": "6.26.0", "fixVersion": "6.26.0" },
            { "name": "lodash", "version": "4.17.10" }
        ]
    ]
}
```

In this example, an issue is introduced through two different versions of the `lodash` package, one of which is a direct dependency while the other is indirect via `babel-template`.

The shortest path is provided first. If an issue applies to the project itself, it appears as the only element in the path. For issues that apply to dependencies, each path starts with a direct dependency.

The `fixVersion` attribute is provided on the first element of each path when that path is upgradable. If the `version` attribute and `fixVersion` attributes are the same, then the upgrade will only involve re-locking transitive dependencies.
