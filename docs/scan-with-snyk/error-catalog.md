# Error Catalog

## Snyk Error Codes

The error codes in the table below describe the codes that you may encounter while working with the [Snyk API](../snyk-api/snyk-api.md) or [CLI](../snyk-cli/). When errors are encountered using the API, they will also have an appropriate [HTTP status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). If you encounter errors without an error code, use the HTTP status code to determine the appropriate action.

***

## Fix

### [PR-FAILURES-0001](#pr-failures-0001)
<a id="#PR-FAILURES-0001"></a>

**Fix scenario not supported**

Snyk failed to open a fix PR as the scenario is not supported.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [PR-FAILURES-0002](#pr-failures-0002)
<a id="#PR-FAILURES-0002"></a>

**SCM rate limit**

SCM rate limit exceeded due to too many requests.

**HTTP Status:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)

### [PR-FAILURES-0003](#pr-failures-0003)
<a id="#PR-FAILURES-0003"></a>

**Unauthorised access**

Request failed due to unathorised access. Please read documentation around adding users and permitted roles.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

**Help Links:**

* [https://docs.snyk.io/snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations](https://docs.snyk.io/snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations)

### [SNYK-PACKAGES-0001](#snyk-packages-0001)
<a id="#SNYK-PACKAGES-0001"></a>

**Unsupported ecosystem**

The language or package manager is not supported. Please refer to the supported package managers in the documentation.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/upgrade-open-source-dependencies-with-automatic-prs#supported-languages-and-scms](https://docs.snyk.io/scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/upgrade-open-source-dependencies-with-automatic-prs#supported-languages-and-scms)

### [SNYK-PACKAGES-0003](#snyk-packages-0003)
<a id="#SNYK-PACKAGES-0003"></a>

**Metadata not found**

Package metadata not or found or missing.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://docs.snyk.io/supported-languages-package-managers-and-frameworks#package-managers-and-frameworks](https://docs.snyk.io/supported-languages-package-managers-and-frameworks#package-managers-and-frameworks)

### [SNYK-PACKAGES-0005](#snyk-packages-0005)
<a id="#SNYK-PACKAGES-0005"></a>

**No mature versions found for package**

Unable to provide a recommended version as no mature versions were found.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-PACKAGES-0006](#snyk-packages-0006)
<a id="#SNYK-PACKAGES-0006"></a>

**No recommended version found**

Unable to provide a recommended version for package using this policy.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-PACKAGES-0007](#snyk-packages-0007)
<a id="#SNYK-PACKAGES-0007"></a>

**Package is already at latest version**

No newer version found for this package, as it is already to latest version.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-PACKAGES-0008](#snyk-packages-0008)
<a id="#SNYK-PACKAGES-0008"></a>

**Version downgrade is not supported**

Unable to suggest a downgrade for a package version.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-PACKAGES-0009](#snyk-packages-0009)
<a id="#SNYK-PACKAGES-0009"></a>

**Invalid version**

Not a valid version for semver format.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://semver.org/](https://semver.org/)

### [SNYK-PR-TEMPLATE-0001](#snyk-pr-template-0001)
<a id="#SNYK-PR-TEMPLATE-0001"></a>

**Failed to get pull request attributes**

Snyk could not get the custom pull request template attributes, using the given variables and the fetched pr template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0002](#snyk-pr-template-0002)
<a id="#SNYK-PR-TEMPLATE-0002"></a>

**Not found**

We could not find your pull request template, have you created one yet? Please check the attached link for instructions on how to setup your pull request template.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0003](#snyk-pr-template-0003)
<a id="#SNYK-PR-TEMPLATE-0003"></a>

**Failed to compile pull request template**

Could not compile your customize pull request template. Please check for syntax errors using the Snyk variables inside the template.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0004](#snyk-pr-template-0004)
<a id="#SNYK-PR-TEMPLATE-0004"></a>

**Failed to parse pull request attributes**

Snyk could not parse the custom pull request template, using the given variables and assigning them to the fetched pr template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0005](#snyk-pr-template-0005)
<a id="#SNYK-PR-TEMPLATE-0005"></a>

**Failed to load YAML file after substituting Snyk variables**

Could not load YAML file after substituting Snyk variables into the custom PR template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0006](#snyk-pr-template-0006)
<a id="#SNYK-PR-TEMPLATE-0006"></a>

**Failed to generate hash for custom PR template**

Snyk could not generate hash using the customer PR files and projects vulnIds.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0007](#snyk-pr-template-0007)
<a id="#SNYK-PR-TEMPLATE-0007"></a>

**Unable to create pull request template**

Snyk could not create pull request template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0008](#snyk-pr-template-0008)
<a id="#SNYK-PR-TEMPLATE-0008"></a>

**Unable to get pull request template**

Snyk could not get pull request template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0009](#snyk-pr-template-0009)
<a id="#SNYK-PR-TEMPLATE-0009"></a>

**Unable to delete pull request template**

Snyk could not delete pull request template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0010](#snyk-pr-template-0010)
<a id="#SNYK-PR-TEMPLATE-0010"></a>

**Invalid payload**

The pull request template payload is invalid.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-PR-TEMPLATE-0011](#snyk-pr-template-0011)
<a id="#SNYK-PR-TEMPLATE-0011"></a>

**Failed to load JSON file after substituting Snyk variables**

Could not load JSON file after substituting Snyk variables into the custom PR template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

### [SNYK-PR-TEMPLATE-0012](#snyk-pr-template-0012)
<a id="#SNYK-PR-TEMPLATE-0012"></a>

**Failed to render default PR template**

Could not render default PR template.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

***

## Snyk

### [SNYK-0001](#snyk-0001)
<a id="#SNYK-0001"></a>

**Service temporarily throttled**

The request rate limit has been exceeded. Wait a few minutes, then try again.

**HTTP Status:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)

### [SNYK-0002](#snyk-0002)
<a id="#SNYK-0002"></a>

**Server error response**

The server doesn't recognize the request method, or it cannot fulfill it. Review the request and try again.

**HTTP Status:** [501](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501)

**Help Links:**

* [https://docs.snyk.io/snyk-api-info](https://docs.snyk.io/snyk-api-info)

### [SNYK-0003](#snyk-0003)
<a id="#SNYK-0003"></a>

**Client request cannot be processed**

The server cannot process the request due to a client error, such as malformed request syntax, size too large, invalid request message framing, or deceptive request routing. Review the request and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-0004](#snyk-0004)
<a id="#SNYK-0004"></a>

**Server communication error**

The server timed out during the request. Check Snyk status, then try again.

**HTTP Status:** [504](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

### [SNYK-0005](#snyk-0005)
<a id="#SNYK-0005"></a>

**Authentication error**

Authentication credentials not recognized, or user access is not provisioned. Revise credentials and try again, or request access from your Snyk administrator.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

### [SNYK-0006](#snyk-0006)
<a id="#SNYK-0006"></a>

**Test limit reached**

You have reached the maximum number of tests in your Snyk plan. This causes Snyk tests on PRs and CLI to fail. Deactivate Snyk Test on your Project or upgrade your Snyk plan.

**HTTP Status:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)

**Help Links:**

* [https://support.snyk.io/s/article/Rate-limit-hit-while-testing-the-project](https://support.snyk.io/s/article/Rate-limit-hit-while-testing-the-project)
* [https://docs.snyk.io/scan-using-snyk/working-with-snyk-in-your-environment/what-counts-as-a-test](https://docs.snyk.io/scan-using-snyk/working-with-snyk-in-your-environment/what-counts-as-a-test)
* [https://support.snyk.io/s/article/Snyk-Test-of-PR-failing-due-to-test-limit](https://support.snyk.io/s/article/Snyk-Test-of-PR-failing-due-to-test-limit)

### [SNYK-0007](#snyk-0007)
<a id="#SNYK-0007"></a>

**Organization is not part of a group**

This error occures when trying to add tags to Organizations that that are part of a Group.

Verify with your Group Admin if the Organization should be in a Group.

If you have more than one Organization, you can set the Organization with which new Projects should be associated by running `snyk config set org=ORG_ID`.

If you want to override this global configuration for individual runs of snyk monitor, `run snyk test --org=ORG_ID` or `snyk monitor --org=ORG_ID`.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/snyk-admin/snyk-projects/project-tags](https://docs.snyk.io/snyk-admin/snyk-projects/project-tags)

### [SNYK-0008](#snyk-0008)
<a id="#SNYK-0008"></a>

**Unable to fulfill the request**

Due to a server gateway error, the server cannot process the request. Check Snyk status and try again.

**HTTP Status:** [502](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

### [SNYK-0009](#snyk-0009)
<a id="#SNYK-0009"></a>

**Unable to fulfill the request**

Due to service availability issues, Snyk cannot process the request. This issue is unexpected, and the service will recover shortly. If the error still occurs, contact Snyk Support.

**HTTP Status:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

### [SNYK-9999](#snyk-9999)
<a id="#SNYK-9999"></a>

**Unable to process request**

The server cannot process the request due to an unexpected error. Check Snyk status, then try again.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

***

## AiBom

### [SNYK-AIBOM-0001](#snyk-aibom-0001)
<a id="#SNYK-AIBOM-0001"></a>

**Unexpected error**

An unexpected error occurred in the AIBOM request. Review the request while providing the debug command flag `-d`. If the error persists, contact Snyk Support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-AIBOM-0002](#snyk-aibom-0002)
<a id="#SNYK-AIBOM-0002"></a>

**Forbidden**

You or your Organization do not have permission to use the AIBOM feature. Check your user permissions or contact Snyk support.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

### [SNYK-AIBOM-0003](#snyk-aibom-0003)
<a id="#SNYK-AIBOM-0003"></a>

**No supported files**

Snyk was unable to find any supported files for the AIBOM command. Ensure the directory you are scanning contains supported files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

***

## Custom Base Images

### [SNYK-CBI-0001](#snyk-cbi-0001)
<a id="#SNYK-CBI-0001"></a>

**Versioning schema does not support tag**

The versioning schema used does not support the given tag. Update the versioning schema to include the tag.

Once the tag of the custom base image is correct, the versioning schema must be modified.
You can use a different versioning schema that supports all tags in the repository or you can update the relevant properties of the versioning schema.

For example, if the repository currently uses Semver, and a new tag "1.2.5.7" needs to be added, then you can use a Custom versioning schema.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations/versioning-schema-for-custom-base-images](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations/versioning-schema-for-custom-base-images)

### [SNYK-CBI-0002](#snyk-cbi-0002)
<a id="#SNYK-CBI-0002"></a>

**Missing required parameter**

Provide an ORG ID or GROUP ID.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0003](#snyk-cbi-0003)
<a id="#SNYK-CBI-0003"></a>

**Project does not exist**

The project could not be found. Check that the project exists, that you have access to the project, and also check that the ID you have provided is the project ID and not a CBI ID.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-CBI-0004](#snyk-cbi-0004)
<a id="#SNYK-CBI-0004"></a>

**Project is not a container image**

The project is not a container image.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations)

### [SNYK-CBI-0005](#snyk-cbi-0005)
<a id="#SNYK-CBI-0005"></a>

**Unable to retrieve group**

The project's org does not belong to a group. In order to use a Custom Base Image, recreate the project and add it to a group or add a group to the org. Note that the group feature is not available to free users.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0006](#snyk-cbi-0006)
<a id="#SNYK-CBI-0006"></a>

**The values in the request do not match**

The request body ID and the request path ID do not match. Ensure that the values are the same and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0007](#snyk-cbi-0007)
<a id="#SNYK-CBI-0007"></a>

**The request body cannot be updated**

The request body does not contain any attributes that can be updated. Provide the necessary attributes and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0008](#snyk-cbi-0008)
<a id="#SNYK-CBI-0008"></a>

**Invalid pagination cursor**

The provided pagination cursor is invalid.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0009](#snyk-cbi-0009)
<a id="#SNYK-CBI-0009"></a>

**Unable to sort by version**

Snyk was unable to filter by version. Provide a repository filter and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0010](#snyk-cbi-0010)
<a id="#SNYK-CBI-0010"></a>

**Unable to update versioning schema**

The versioning schema could not be applied to all images in the repository. Therefore, no resources have been updated. Update the provided versioning schema so that all tags in the repository fit the new schema.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0011](#snyk-cbi-0011)
<a id="#SNYK-CBI-0011"></a>

**Project is already linked to a custom base image**

The project ID provided is already linked to another Custom Base Image.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0012](#snyk-cbi-0012)
<a id="#SNYK-CBI-0012"></a>

**No versioning schema for repository**

No versioning schema exists for the repository. This image is the first in its repository. Provide a versioning schema that fits the format of current and future images in this repository.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0013](#snyk-cbi-0013)
<a id="#SNYK-CBI-0013"></a>

**Unable to apply versioning schema**

A versioning schema already exists for repository. Remove the "versioning_schema" property or, if you want to update the versioning schema, use the PATCH endpoint.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CBI-0014](#snyk-cbi-0014)
<a id="#SNYK-CBI-0014"></a>

**Unable to find custom base image**

Unable to find the requested custom base image. Try again, and if the error persists, contact Snyk support.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-CBI-0015](#snyk-cbi-0015)
<a id="#SNYK-CBI-0015"></a>

**Custom base image does not exist**

The requested custom base image does not exist.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

### [SNYK-CBI-0016](#snyk-cbi-0016)
<a id="#SNYK-CBI-0016"></a>

**Unable to update custom base image**

An internal error occurred while trying to update a custom base image. Try again, and if the error persists, contact Snyk support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

### [SNYK-CBI-0017](#snyk-cbi-0017)
<a id="#SNYK-CBI-0017"></a>

**Unable to retrieve project properties**

An internal error occurred while trying to retrieve project properties. Try again, and if the error persists, contact Snyk support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

### [SNYK-CBI-0018](#snyk-cbi-0018)
<a id="#SNYK-CBI-0018"></a>

**Unable to retrieve image collection**

An internal error occurred while trying to retrieve the image collection. Try again, and if the error persists, contact Snyk support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

### [SNYK-CBI-0019](#snyk-cbi-0019)
<a id="#SNYK-CBI-0019"></a>

**Unable to create versioning schema**

The provided versioning schema is invalid and image could therefor not be created. Provide a properly formatted versioning schema and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

***

## CLI

### [SNYK-CLI-0000](#snyk-cli-0000)
<a id="#SNYK-CLI-0000"></a>

**Unspecified Error**

The encountered error only provides basic information, please take a look at the given details.If they do not help to resolve the issue, consider debugging or consulting support.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands](https://docs.snyk.io/snyk-cli/commands)
* [https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli](https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli)

### [SNYK-CLI-0001](#snyk-cli-0001)
<a id="#SNYK-CLI-0001"></a>

**Unable to set environment**

The specified environment cannot be used. As a result, the configuration remains unchanged.Provide the correct specifications for the environment and try again.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/config-environment](https://docs.snyk.io/snyk-cli/commands/config-environment)

### [SNYK-CLI-0002](#snyk-cli-0002)
<a id="#SNYK-CLI-0002"></a>

**Possible inconsistent configuration**

You can configure the CLI in different ways, for example via Environment Variables or configuration file.
If one parameter is configured multiple times, it is probably unintentional and might cause unexpected behavior.
Review configured environment variables and ensure that everything is intentional. If so, you can skip this check by using --no-check.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/config-environment](https://docs.snyk.io/snyk-cli/commands/config-environment)

### [SNYK-CLI-0003](#snyk-cli-0003)
<a id="#SNYK-CLI-0003"></a>

**Empty flag option**

A specified flag is missing an option value. Provide a correct option value and try again.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary)

### [SNYK-CLI-0004](#snyk-cli-0004)
<a id="#SNYK-CLI-0004"></a>

**Invalid flag option**

A specified flag option or combination is invalid. Provide a valid flag option or combination and try again.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary)

### [SNYK-CLI-0005](#snyk-cli-0005)
<a id="#SNYK-CLI-0005"></a>

**Unable to get vulnerabilities from resource**

If you are testing an npm package, check the version and package name and try running `snyk test` again. If you are testing a repository, try testing it at https://snyk.io/test/.For further assistance, run `snyk help` or see the Snyk docs.


### [SNYK-CLI-0006](#snyk-cli-0006)
<a id="#SNYK-CLI-0006"></a>

**Missing AUTH token**

When running your command, Snyk requires an authenticated account. You must include your API token as an environment value, or use `snyk auth` to authenticate.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/authenticate-to-use-the-cli](https://docs.snyk.io/snyk-cli/authenticate-to-use-the-cli)
* [https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli)

### [SNYK-CLI-0007](#snyk-cli-0007)
<a id="#SNYK-CLI-0007"></a>

**Incomplete command arguments**

The specified CLI command includes missing or misconfigured arguments. Provide the correct arguments and try again.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary)

### [SNYK-CLI-0008](#snyk-cli-0008)
<a id="#SNYK-CLI-0008"></a>

**No supported files found**

Snyk could not detect any supported target files. Ensure the files you are importing are supported, that you are in the right directory, and try again.


**Help Links:**

* [https://docs.snyk.io/supported-languages-package-managers-and-frameworks](https://docs.snyk.io/supported-languages-package-managers-and-frameworks)

### [SNYK-CLI-0009](#snyk-cli-0009)
<a id="#SNYK-CLI-0009"></a>

**Too many vulnerable paths to Project**

There are too many vulnerable paths to process the project. If your command supports it, consider the following:pruning repeated sub-dependencies (`snyk test -p`); excluding directories (`snyk test --all-projects --exclude=dir1,file2`); setting a detection depth (`snyk test --all-projects --detection-depth=3`). If the error still occurs, consider debugging or contact Snyk Support.

**HTTP Status:** [413](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413)

**Help Links:**

* [https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#options-for-multiple-commands](https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#options-for-multiple-commands)
* [https://docs.snyk.io/snyk-cli/commands/test#prune-repeated-subdependencies-p](https://docs.snyk.io/snyk-cli/commands/test#prune-repeated-subdependencies-p)
* [https://docs.snyk.io/snyk-cli/commands/test#detection-depth-less-than-depth-greater-than](https://docs.snyk.io/snyk-cli/commands/test#detection-depth-less-than-depth-greater-than)
* [https://docs.snyk.io/snyk-cli/commands/test#exclude-less-than-name-greater-than-less-than-name-greater-than-...greater-than](https://docs.snyk.io/snyk-cli/commands/test#exclude-less-than-name-greater-than-less-than-name-greater-than-...greater-than)

### [SNYK-CLI-0010](#snyk-cli-0010)
<a id="#SNYK-CLI-0010"></a>

**CLI validation failure**

CLI was unable to validate the required parameter. Provide the correct parameter and try again. If the error still occurs, consider debugging or contact Snyk Support.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CLI-0011](#snyk-cli-0011)
<a id="#SNYK-CLI-0011"></a>

**SCA failure**

CLI was unable to execute your SCA command, please take a look at the given details.If they do not help to resolve the issue, consider debugging or consulting support.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/test](https://docs.snyk.io/snyk-cli/commands/test)

### [SNYK-CLI-0012](#snyk-cli-0012)
<a id="#SNYK-CLI-0012"></a>

**IAC failure**

CLI was unable to execute your IAC command, please take a look at the given details.If they do not help to resolve the issue, consider debugging or consulting support.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/iac](https://docs.snyk.io/snyk-cli/commands/iac)

### [SNYK-CLI-0013](#snyk-cli-0013)
<a id="#SNYK-CLI-0013"></a>

**SAST failure**

CLI was unable to execute your SAST command, please take a look at the given details.If they do not help to resolve the issue, consider debugging or consulting support.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/code](https://docs.snyk.io/snyk-cli/commands/code)

### [SNYK-CLI-0014](#snyk-cli-0014)
<a id="#SNYK-CLI-0014"></a>

**Feature under development**

This feature is under development and is not yet available for public use.


### [SNYK-CLI-0015](#snyk-cli-0015)
<a id="#SNYK-CLI-0015"></a>

**Command is experimental**

This CLI command is experimental, which means it is provided "as-is" without warranty of any kind.
You must acknowledge this by specifying the --experimental flag to run the command.


### [SNYK-CLI-0016](#snyk-cli-0016)
<a id="#SNYK-CLI-0016"></a>

**Feature not enabled**

This feature is disabled for your current organization. You can enable it in the settings or switch to an organization where it's already enabled.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

### [SNYK-CLI-0017](#snyk-cli-0017)
<a id="#SNYK-CLI-0017"></a>

**DNS resolution failed**

Unable to resolve the hostname to an IP address. Troubleshooting steps.
1) Test DNS resolution: nslookup api.<instance>.snyk.io.
2) Try different DNS servers: Change DNS to 8.8.8.8 or 1.1.1.1.
3) Check corporate proxy/firewall DNS blocking.
4) Verify hostname spelling in your Snyk configuration.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli](https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli)
* [https://status.snyk.io/](https://status.snyk.io/)

### [SNYK-CLI-0018](#snyk-cli-0018)
<a id="#SNYK-CLI-0018"></a>

**Network request timeout**

The network request timed out. Troubleshooting steps.
1) Test connectivity: ping api.<instance>.snyk.io.
2) Check corporate proxy timeout settings.
3) Try different network: Mobile hotspot or different WiFi.
4) Check if firewall is blocking or throttling connections.


**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)
* [https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli](https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli)

### [SNYK-CLI-0019](#snyk-cli-0019)
<a id="#SNYK-CLI-0019"></a>

**Network unreachable**

Unable to reach the target network or host. Troubleshooting steps.
1) Test Snyk connectivity: ping api.<instance>.snyk.io.
2) Check corporate firewall blocks Snyk domains.
3) Check if VPN routing is blocking Snyk domains.
4) Try mobile hotspot to isolate network issues.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli](https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli)
* [https://status.snyk.io/](https://status.snyk.io/)

### [SNYK-CLI-0020](#snyk-cli-0020)
<a id="#SNYK-CLI-0020"></a>

**TLS certificate error**

There was an issue with the TLS/SSL certificate during the secure connection. Troubleshooting steps.
1) Check system time: Ensure your system clock is correct (certificates are time-sensitive).
2) Update system certificates: Windows Update or macOS Software Update.
3) Corporate firewall: Check if corporate firewall intercepts SSL traffic.
4) Custom certificates: Set NODE_EXTRA_CA_CERTS environment variable to path of your CA certificate file.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli](https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli)
* [https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli)

### [SNYK-CLI-0021](#snyk-cli-0021)
<a id="#SNYK-CLI-0021"></a>

**Connection refused**

The connection to the server was refused. Troubleshooting steps.
1) Check Snyk status: Visit status.snyk.io for service outages.
2) Test connectivity: ping api.<instance>.snyk.io.
3) Check corporate proxy blocks HTTPS connections to Snyk.
4) Try mobile hotspot or different network.


**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)
* [https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli](https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli)

### [SNYK-CLI-0022](#snyk-cli-0022)
<a id="#SNYK-CLI-0022"></a>

**Network communication error**

An unexpected network error occurred during communication. Troubleshooting steps.
1) Test connectivity: ping api.<instance>.snyk.io.
2) Check proxy settings: HTTP_PROXY and HTTPS_PROXY environment variables.
3) Run with verbose logging: snyk command --debug.
4) Try mobile hotspot to isolate network issues.


**Help Links:**

* [https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli](https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli)
* [https://status.snyk.io/](https://status.snyk.io/)

### [SNYK-OS-7001](#snyk-os-7001)
<a id="#SNYK-OS-7001"></a>

**Request to Snyk API timeout**

A request to the Snyk API has unexpectedly timeout. Check Snyk status, then try again.

**HTTP Status:** [504](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504)

**Help Links:**

* [https://status.snyk.io/](https://status.snyk.io/)
* [https://privatecloudstatus.snyk.io](https://privatecloudstatus.snyk.io)

***

## Code

### [SNYK-CODE-0001](#snyk-code-0001)
<a id="#SNYK-CODE-0001"></a>

**Analysis file count limit exceeded**

This error occurs when the analysis target has a supported file count that exceeds current system limits.

To reduce the file count, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code)
* [https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process](https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process)
* [https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli](https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli)

### [SNYK-CODE-0002](#snyk-code-0002)
<a id="#SNYK-CODE-0002"></a>

**Analysis result size limit exceeded**

This error occurs when the analysis target generates a result with a byte size that exceeds current system limits.

To reduce the overall result size, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process](https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process)
* [https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli](https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli)

### [SNYK-CODE-0003](#snyk-code-0003)
<a id="#SNYK-CODE-0003"></a>

**Analysis target size limit exceeded**

This error occurs when the analysis target byte size exceeds current system limits.

To reduce the overall result size, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli](https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli)

### [SNYK-CODE-0004](#snyk-code-0004)
<a id="#SNYK-CODE-0004"></a>

**Analysis target includes a file with a name longer than 255 bytes**

This error occurs when the analysis target has a file name length that exceeds 255 bytes.

To be able to scan the analysis target, rename the file to a name that is 255 bytes or less.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/scan-with-snyk/supported-languages-and-frameworks/introduction-to-snyk-supported-languages-and-frameworks#filename-length-limitation](https://docs.snyk.io/scan-with-snyk/supported-languages-and-frameworks/introduction-to-snyk-supported-languages-and-frameworks#filename-length-limitation)

### [SNYK-CODE-0005](#snyk-code-0005)
<a id="#SNYK-CODE-0005"></a>

**Snyk Code is not enabled**

This error occurs when Snyk Code is not enabled for the current Organization. Activate Snyk Code and try again..

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

**Help Links:**

* [https://docs.snyk.io/scan-using-snyk/snyk-code/configure-snyk-code#enable-snyk-code-in-snyk-web-ui](https://docs.snyk.io/scan-using-snyk/snyk-code/configure-snyk-code#enable-snyk-code-in-snyk-web-ui)

### [SNYK-CODE-0006](#snyk-code-0006)
<a id="#SNYK-CODE-0006"></a>

**Project not supported**

Snyk was unable to find supported files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/getting-started/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code](https://docs.snyk.io/getting-started/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code)

### [SNYK-CODE-0007](#snyk-code-0007)
<a id="#SNYK-CODE-0007"></a>

**SAST Rule extension already exists for the Group**

A published SAST Rule extension with the same fully qualified name already exists for the given Group.

**HTTP Status:** [409](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409)

### [SNYK-CODE-0008](#snyk-code-0008)
<a id="#SNYK-CODE-0008"></a>

**Organization relationships must be unique**

Each Org relationship to a Snyk SAST Rule extension must be unique.

Make sure each Org in relationships has a different ID.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CODE-0009](#snyk-code-0009)
<a id="#SNYK-CODE-0009"></a>

**Group relationship must match the Group in the requested URL**

You cannot associate a Snyk SAST Rule extension to any other Group.

Make sure the Group ID under relationships matches the Group ID in the request path.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CODE-0010](#snyk-code-0010)
<a id="#SNYK-CODE-0010"></a>

**Organization outside of the administrating Group**

You cannot use the SAST Rule extensions feature with an Org outside of the administrating Group.

Make sure each Org in the request is within the requested Group.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CODE-0011](#snyk-code-0011)
<a id="#SNYK-CODE-0011"></a>

**SAST Rule extension limit reached**

You have hit the maximum number of published Snyk SAST Rule extensions allowed for a Group.

To create a new SAST Rule extension you will have to remove an existing one.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CODE-0012](#snyk-code-0012)
<a id="#SNYK-CODE-0012"></a>

**SAST Rule Extension already published for the Group**

The Rule Extension under test conflicts with an already published SAST Rule Extension.

A test cannot be performed if a SAST Rule Extension with the same fully qualified name
and type is already published for the Group. Either delete the already published SAST Rule Extension
or perform a test with a different fully qualified name or type.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-CODE-0013](#snyk-code-0013)
<a id="#SNYK-CODE-0013"></a>

**Requested test ID not found**

The requested Test ID for testing SAST Rule Extension was not found.

Make sure to provide a valid Test ID.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-CODE-0014](#snyk-code-0014)
<a id="#SNYK-CODE-0014"></a>

**Test results have expired**

The results for testing SAST Rule Extensions have expired and are no longer available.

Please trigger a new test.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

***

## Integration

### [SNYK-INTEGRATION-0001](#snyk-integration-0001)
<a id="#SNYK-INTEGRATION-0001"></a>

**SCM integration not found**

Ensure your SCM integration exists and that it is correctly set up.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations](https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations)

***

## OpenAPI

### [SNYK-OPENAPI-0001](#snyk-openapi-0001)
<a id="#SNYK-OPENAPI-0001"></a>

**Bad request**

The server cannot process the request due to invalid or corrupt data. Review the request, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api ](https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api )

### [SNYK-OPENAPI-0002](#snyk-openapi-0002)
<a id="#SNYK-OPENAPI-0002"></a>

**Forbidden**

Access to the requested resource is forbidden. Review the request, then try again.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

### [SNYK-OPENAPI-0003](#snyk-openapi-0003)
<a id="#SNYK-OPENAPI-0003"></a>

**Not acceptable**

The server cannot provide a response that matches the provided accept headers. Review the request, then try again.

**HTTP Status:** [406](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406)

### [SNYK-OPENAPI-0004](#snyk-openapi-0004)
<a id="#SNYK-OPENAPI-0004"></a>

**Not found**

The server cannot find the requested resource. Review the request, then try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OPENAPI-0005](#snyk-openapi-0005)
<a id="#SNYK-OPENAPI-0005"></a>

**Method not allowed**

The target endpoint does not support your request method. Review the request, then try again.

**HTTP Status:** [405](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405)

### [SNYK-OPENAPI-0006](#snyk-openapi-0006)
<a id="#SNYK-OPENAPI-0006"></a>

**Request entity too large**

The request entity exceeds server limitations. Reduce the size of the request entity, then try again.

**HTTP Status:** [413](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413)

### [SNYK-OPENAPI-0007](#snyk-openapi-0007)
<a id="#SNYK-OPENAPI-0007"></a>

**Unauthorized**

The request lacks authentication credentials for the requested resource. Ensure you are sending valid credentials, then try again.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**Help Links:**

* [https://docs.snyk.io/snyk-api-info/authentication-for-api](https://docs.snyk.io/snyk-api-info/authentication-for-api)

### [SNYK-OPENAPI-0008](#snyk-openapi-0008)
<a id="#SNYK-OPENAPI-0008"></a>

**Unsupported media type**

The media format of the request is not supported. Change media format, then try again.

**HTTP Status:** [415](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415)

### [SNYK-OPENAPI-0009](#snyk-openapi-0009)
<a id="#SNYK-OPENAPI-0009"></a>

**Conflict**

The request could not be completed due to a conflict with the current state of the target resource. Review the request, then try again.

**HTTP Status:** [409](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409)

***

## Open Source Languages & Package Managers

### [SNYK-OS-0001](#snyk-os-0001)
<a id="#SNYK-OS-0001"></a>

**Unable to parse manifest file**

The provided manifest file could not be parsed as it has invalid syntax or does not match the expected schema. Review the manifest file, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OS-0002](#snyk-os-0002)
<a id="#SNYK-OS-0002"></a>

**Unable to parse lock file**

The provided lock file could not be parsed as it has invalid syntax or does not match the expected schema. Review the lock file, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OS-0003](#snyk-os-0003)
<a id="#SNYK-OS-0003"></a>

**Unknown dependency version**

Dependency version could not be resolved.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://support.snyk.io/s/article/Could-not-determine-version-for-dependencies](https://support.snyk.io/s/article/Could-not-determine-version-for-dependencies)

### [SNYK-OS-0004](#snyk-os-0004)
<a id="#SNYK-OS-0004"></a>

**Missing required request header**

The server encountered a request that is missing a mandatory request header.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-0005](#snyk-os-0005)
<a id="#SNYK-OS-0005"></a>

**Payload missing required elements**

The server could not process the request.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-0006](#snyk-os-0006)
<a id="#SNYK-OS-0006"></a>

**Files cannot be processed**

The dependency service could not process the files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-0007](#snyk-os-0007)
<a id="#SNYK-OS-0007"></a>

**Cannot get file from source**

Could not get the file from the source URL.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-0008](#snyk-os-0008)
<a id="#SNYK-OS-0008"></a>

**Missing environment variable**

The server encountered a critical operation that requires a specific environment variable, but the variable is not set or is not accessible within the current environment.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-0009](#snyk-os-0009)
<a id="#SNYK-OS-0009"></a>

**Brokered connections not currently supported**

The service encountered a permissions or credentials error most likely related to an import through a brokered connection for a scanner that does not yet support that.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-0010](#snyk-os-0010)
<a id="#SNYK-OS-0010"></a>

**Snyk failed to clone your repository**

We encountered a fatal error from Git while trying to clone your code using your provided credentials. Please verify that:

* Your provided credentials are correct or not scoped too narrowly.
* The branch you've asked us to clone exists.
* The repository you've provided is accessible from the internet and you are not connected through a broker.

And try the operation again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-0011](#snyk-os-0011)
<a id="#SNYK-OS-0011"></a>

**Unsupported platform**

The specified platform is not supported.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-DOTNET-0001](#snyk-os-dotnet-0001)
<a id="#SNYK-OS-DOTNET-0001"></a>

**Unsupported manifest file type for remediation**

The provided manifest file is not supported by Snyk for .NET.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net)

### [SNYK-OS-DOTNET-0002](#snyk-os-dotnet-0002)
<a id="#SNYK-OS-DOTNET-0002"></a>

**Target framework not supported**

The provided manifest file defines a `<TargetFramework>` or `<TargetFrameworks>` that is not currently supported by Snyk's .NET scanning solution.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-DOTNET-0003](#snyk-os-dotnet-0003)
<a id="#SNYK-OS-DOTNET-0003"></a>

**Your C# code is missing a static Main function**

This error occurs when no static Main method with a correct signature is found in the code that produces an executable file. 
It also occurs if the entry point function, `Main`, is defined with the wrong case, such as lower-case main.

In order to fix this issue, ensure that your program has a .cs file that contains a main function, such as
```c#
namespace Example
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("hello world");
        }
    }
}
```

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://learn.microsoft.com/en-us/dotnet/csharp/misc/cs5001](https://learn.microsoft.com/en-us/dotnet/csharp/misc/cs5001)

### [SNYK-OS-DOTNET-0004](#snyk-os-dotnet-0004)
<a id="#SNYK-OS-DOTNET-0004"></a>

**The dotnet CLI is unable to generate a self-contained binary**

This error occurs when running `dotnet publish --sc --framework <your-target-framework>` fails to generate a 
self-contained binary. Snyk needs to run this command in order to adequately determine the dependency tree for your project. If this command fails, Snyk cannot continue.

Steps to determine why this happened:

* Checkout a clean version of your project in a temporary folder
* Run `dotnet publish --sc --framework <your-target-framework> ` on your project, and confirm this step fails.

If this step is successful locally, it is possible that Snyk is running another version of the .NET SDK. To tell Snyk which version of the .NET SDK to use, consider using the [global.json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json) solution provided by Microsoft.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/](https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/)
* [https://learn.microsoft.com/en-us/dotnet/core/tools/global-json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json)
* [https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L67](https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L67)

### [SNYK-OS-DOTNET-0005](#snyk-os-dotnet-0005)
<a id="#SNYK-OS-DOTNET-0005"></a>

**The dotnet CLI was unable to restore from private package sources**

This error occurs when running `dotnet restore` fails to access dependencies stored in a private package source that Snyk does not have access to. 

This means that your `.csproj` file or files refer to a dependency hosted on a private package store or Nuget Artifact Registry defined in your `NuGet.config` file, such as:

```xml
<?xml version="1.0" encoding="utf-8"?>
<configuration>
  <packageSources>
    <clear />
    <add key="AzureFeed" value="https://pkgs.dev.azure.com/your-org/_packaging/your-repo/nuget/v3/index.json" />
    <add key="nuget.org" value="https://api.nuget.org/v3/index.json" />
  </packageSources>
</configuration>
```

In order to allow Snyk to access your private dependency package source, you must supply Snyk with a valid JSON object as a private registry token in the .NET language settings.

You can set up a connection to your private Nuget repository in your Snyk integration settings.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**Help Links:**

* [https://github.com/microsoft/artifacts-credprovider#environment-variables](https://github.com/microsoft/artifacts-credprovider#environment-variables)

### [SNYK-OS-DOTNET-0006](#snyk-os-dotnet-0006)
<a id="#SNYK-OS-DOTNET-0006"></a>

**Missing MSBuild Condition Construct in project file**

The `dotnet` tool was unable to locate the `.targets`, `.csproj` or `.props` file responsible for one or more MSBuild conditions in your project file.

The tool encountered an error like 
```
/path/to/file/project.csproj(33,13): error MSB4100: Expected "$(SomeCondition)" to evaluate to a boolean instead of "", in condition "!$(SomeCondition)".
```

This means the condition definition is missing in the project file that is currently being restored and in any project linked to it from there.      

Snyk can scan only the project files accessible in the current repository or the private dependencies available to Snyk.

For example, if your code has the following structure:

```title=project.targets
<Project>
  <PropertyGroup>
    <SomeCondition Condition="'$(SomeCondition)' == ''">false</SomeCondition>
  </PropertyGroup>
</Project>
```

And

```title=project.csproj
<Project Sdk='Microsoft.NET.Sdk'>
  <Import Project='..\external-libraries\some-library\project.targets' />
  <PropertyGroup>
    <TargetFrameworks>net8.0</TargetFrameworks>
  </PropertyGroup>
  <ItemGroup Condition='!$(SomeCondition)'>
    <PackageReference Include='Newtonsoft.Json' Version='13.0.3' />
  </ItemGroup>
</Project>
```

And `external-libraries` is not a part of your repository currently being scanned, Snyk is not able to find it.

This error occurs when your code depends on external libraries that are added to or generated from your source code using external tools unknown to Snyk or as part of a build step in your build or a deployment pipeline.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://learn.microsoft.com/en-us/visualstudio/msbuild/msbuild-conditional-constructs](https://learn.microsoft.com/en-us/visualstudio/msbuild/msbuild-conditional-constructs)

### [SNYK-OS-DOTNET-0007](#snyk-os-dotnet-0007)
<a id="#SNYK-OS-DOTNET-0007"></a>

**No target frameworks found in manifest files**

Snyk was unable to detect any `<TargetFramework>`s in the supplied manifest files. 

If you are using `Directory.Build.props` files to determine the target framework, ensure that it is named as such. Due to performance considerations on the customer's SCM network, Snyk does not perform case-insensitive searches for `.props` files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://learn.microsoft.com/en-us/visualstudio/msbuild/customize-by-directory?view=vs-2022#directorybuildprops-and-directorybuildtargets](https://learn.microsoft.com/en-us/visualstudio/msbuild/customize-by-directory?view=vs-2022#directorybuildprops-and-directorybuildtargets)

### [SNYK-OS-DOTNET-0008](#snyk-os-dotnet-0008)
<a id="#SNYK-OS-DOTNET-0008"></a>

**Your global.json is targeting an outdated SDK version**

Snyk supports the latest channels of .NET which is currently [supported by Microsoft](https://dotnet.microsoft.com/en-us/download/dotnet), but does **not** guarantee to support all SDK versions within each currently supported channel.

Within the supported channels, Snyk aims to support most, if not all, of the SDK versions currently released under the **newest** of the channels.

If the channels currently supported by Microsoft are `8.0`, `7.0` and `6.0`, Snyk **will** support all of the *latest* SDKs released for these channels.

If the SDK versions released under `8.0.3` are: `8.0.203`, `8.0.202` and `8.0.103`, Snyk **cannot** guarantee to support *all* of them, but makes an effort to do so. Snyk **will** support the latest of the SDK versions currently released by Microsoft. 

If channel `8.0` is the newest channel currently supported, Snyk **cannot** guarantee that multiple, specific SDK versions for older, still supported channels such as .NET 6. 

### Example support matrix

If:

* .NET channels currently supported by Microsoft are `.NET 8.0`,  `.NET 7.0` and  `.NET 6.0`
* Newest SDK version under `.NET 8.0` is `8.0.203`

Then:

| Channel |              SDK             | End-of-Life |  Supported  |
|:-------:|:----------------------------:|:-----------:|:-----------:|
|   8.0   | 8.0.203  (latest in channel) |      No     |     Yes     |
|   8.0   |            8.0.202           |      No     |     Yes     |
|   8.0   |            8.0.103           |      No     |     Yes     |
|         |             (...)            |             |             |
|   7.0   | 7.0.407  (latest in channel) |      No     |     Yes     |
|   7.0   |            7.0.314           |      No     |      No     |
|         |             (...)            |             |             |
|   6.0   |            6.0.420           |      No     |     Yes     |
|   6.0   |            6.0.128           |      No     |      No     |
|         |             (..)             |             |             |
|   5.0   |  5.0.408 (latest in channel) |     Yes     |      No     |
|   5.0   |            5.0.214           |     Yes     |      No     |
|         |             (..)             |             |             |

### Workarounds

This limitation can lead to scan failures for customers that are pinning SDK versions in their `global.json` files without a [rollForward](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward) directive, such as:
```json
{
  "sdk": {
    "version": "6.0.101"
  }
}
```
Since as `6.0` is not the newest .NET channel. 

To work around this issue, we recommend that customers employ some flexibility in their `global.json` file by employing the `rollFoward` directive to be `latestMajor`, as such:
```json
{
  "sdk": {
    "version": "6.0.101",
    "rollForward": "latestMajor"
  }
}
```

Which will allow Snyk to scan your code using a newer version of the SDK, despite your version pinning.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://versionsof.net/core/](https://versionsof.net/core/)
* [https://dotnet.microsoft.com/en-us/download/dotnet](https://dotnet.microsoft.com/en-us/download/dotnet)
* [https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward)

### [SNYK-OS-DOTNET-0009](#snyk-os-dotnet-0009)
<a id="#SNYK-OS-DOTNET-0009"></a>

**Project failed to build due to missing type or namespace references**

While attempting to build your solution for scanning, the `dotnet` SDK was unable to restore one or more projects referenced in your manifest files.

Please note that Snyk runs these builds on a **case-sensitive** filesystem, meaning that `<ProjectReference>../src/NS.Project.csproj</ProjectReference>` and `<ProjectReference>../src/ns.project.csproj</ProjectReference>` are not referring to the same thing.

This can present itself as a problem for customers that are using Mac or Windows build pipeline where file systems are not case-sensitive. In this case, verify you're referring to the right manifest files and check the Snyk import logs for more details.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-messages/assembly-references#missing-references](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-messages/assembly-references#missing-references)

### [SNYK-OS-DOTNET-0010](#snyk-os-dotnet-0010)
<a id="#SNYK-OS-DOTNET-0010"></a>

**The 10 GB space limit for downloaded Nuget dependencies has been exceeded**

The total size of the downloaded Nuget dependencies in the manifest files exceeds the 10GB limit.
This often happens due to the large number or the large size of Nuget dependencies. 

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/supported-languages-package-managers-and-frameworks/.net/improved-.net-scanning#limitations-on-improved-.net-scanning-for-scm-integrations](https://docs.snyk.io/supported-languages-package-managers-and-frameworks/.net/improved-.net-scanning#limitations-on-improved-.net-scanning-for-scm-integrations)

### [SNYK-OS-DOTNET-0011](#snyk-os-dotnet-0011)
<a id="#SNYK-OS-DOTNET-0011"></a>

**The dotnet CLI is unable to download and install all the required dependencies**

This error occurs when running `dotnet restore <path-to-csproj>` fails to generate a 
manifest of the resolved dependencies. Snyk needs to run this command in order to adequately determine the dependency tree for your project. 
If this command fails, Snyk cannot continue.

Steps to determine why this happened:

* Checkout a clean version of your project in a temporary folder
* Run `dotnet restore <path-to-csproj> ` on your project, and confirm this step fails.

If this step is successful locally, it is possible that Snyk is running another version of the .NET SDK. To tell Snyk which version of the .NET SDK to use, consider using the [global.json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json) solution provided by Microsoft.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/](https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/)
* [https://learn.microsoft.com/en-us/dotnet/core/tools/global-json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json)
* [https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L50](https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L50)

### [SNYK-OS-GO-0001](#snyk-os-go-0001)
<a id="#SNYK-OS-GO-0001"></a>

**Failed to access private module**

Snyk could not access the private modules within your go.mod files.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go)

### [SNYK-OS-GO-0002](#snyk-os-go-0002)
<a id="#SNYK-OS-GO-0002"></a>

**Go mod file not found**

A go.mod file was not found in the current directory or any parent directory.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go)

### [SNYK-OS-GO-0003](#snyk-os-go-0003)
<a id="#SNYK-OS-GO-0003"></a>

**OAuth re-authorization required**

Your code is cloned on an isolated environment using Git as it is required by Snyk to analyze its dependencies.

Your Organization has enabled or enforced SAML SSO after you authorized Snyk to access your code, and a re-authentication is therefore required.

The error you're seeing is usually reproducible by attempting to do a `git clone` of your repository with incorrectly configured credentials.
Verify your authentication configuration with your Git cloud provider and try again.

{% hint style="warning" %}
**Error has been deprecated**

Reason: This error has been moved to a more generalized namespace to avoid repetition.

Going forward, [SNYK-OS-8004](#snyk-os-8004) will be used instead.
{% endhint %}

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso](https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso)

### [SNYK-OS-GO-0004](#snyk-os-go-0004)
<a id="#SNYK-OS-GO-0004"></a>

**Your project repository is missing required files**

Generating the dependency graph requires Snyk to run go list `go list -deps -json` inside the project. If the operation fails, creating a full dependency graph cannot continue.  

This error means that you need some cleanup, (such as `go mod tidy`) or your project deployment process contains a code generation step such as `protobuf` or similar that is not currently supported by Snyk. 

To verify if this is the case, clone your project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. 

If Snyk cannot process your code successfully, insert the Snyk CLI as part of your deployment pipeline.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/snyk-cli](https://docs.snyk.io/snyk-cli)
* [https://github.com/snyk/snyk-go-plugin](https://github.com/snyk/snyk-go-plugin)
* [https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go](https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go)

### [SNYK-OS-GO-0005](#snyk-os-go-0005)
<a id="#SNYK-OS-GO-0005"></a>

**Your project repository has inconsistent vendoring information**

Generating the dependency graph requires Snyk to run `go list -deps -json` inside the project. If the operation fails, creating a full dependency graph cannot continue.  

This error means that there is inconsistency between your `vendor/modules.txt` file and your `go.mod` file. To remediate, you need to:

* `go mod vendor`
* `go mod tidy`

Next, commit those changes to your repo. Snyk does not manipulate with your code on our end by design, which is why this is not done automatically.

To verify if this is the case, clone your project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. 
Then try and run the above mentioned commands and see if your SCM system reports changes in files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://go.dev/ref/mod#go-mod-vendor](https://go.dev/ref/mod#go-mod-vendor)

### [SNYK-OS-GO-0006](#snyk-os-go-0006)
<a id="#SNYK-OS-GO-0006"></a>

**Unsupported external file generation**

Snyk currently does not support external file generation in your project. This limitation is due to Snyk's lack of visibility into the third-party generator tools you may be using and the specific commands required to generate these files.

Snyk can only work with the files available in your repository and does not have insight into the generation process for external files.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-GO-0007](#snyk-os-go-0007)
<a id="#SNYK-OS-GO-0007"></a>

**Unable to access private dependencies**

The Go tool encountered a `DepsError` while trying to download a private dependency. Private repositories that are not accessible to the public internet and are not available on the official Go proxy mirror are cloned with a version control system and built on demand. 
This requires the VCS to have the correct access rights to that repository.

Snyk supports private repositories that are hosted in the same Organization and on the same Project that is scanned for vulnerabilities. The authentication to the private repository is the same as the authentication used to integrate that repository with Snyk. 

This error appears when the authorization credentials do not allow access to the requested private dependency. 

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://go.dev/ref/mod#vcs](https://go.dev/ref/mod#vcs)

### [SNYK-OS-GO-0008](#snyk-os-go-0008)
<a id="#SNYK-OS-GO-0008"></a>

**Unable to fetch private dependencies**

The Go tool encountered a permissions error while fetching one of the private dependencies. Ensure that the integration token you used to sign in to Snyk is properly configured so that Snyk can access the private dependencies.

The Snyk Go integration only supports private dependencies that are used inside the same Organization as the Project you are scanning.

This error appears when Snyk is unable to properly access the authorization credentials for the requested private dependency. 

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

### [SNYK-OS-GO-0009](#snyk-os-go-0009)
<a id="#SNYK-OS-GO-0009"></a>

**Toolchain not available**

Could not download Go toolchain.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-GO-0010](#snyk-os-go-0010)
<a id="#SNYK-OS-GO-0010"></a>

**The 10 GB space limit for downloaded Golang dependencies has been exceeded**

The total size of the downloaded Golang dependencies in the manifest file exceeds the 10GB limit.
This often happens due to the large size or large number of Golang dependencies.

Currently this is a product limitation for SCM. As a workaround, use the 'snyk monitor' command via Snyk CLI.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/monitor](https://docs.snyk.io/snyk-cli/commands/monitor)

### [SNYK-OS-GO-0011](#snyk-os-go-0011)
<a id="#SNYK-OS-GO-0011"></a>

**No secure protocol found for repository**

The Go toolchain could not find a secure protocol (HTTPS) to access the repository.

This error typically occurs when the repository URL is not configured to use a secure protocol, or the necessary credentials for accessing the repository securely are not provided.

Ensure that the repository URL uses a secure protocol (https://) and verify that the necessary authentication credentials (such as HTTPS credentials) are correctly configured and accessible by the Go toolchain.
 
Note: SSH is not supported.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-GO-0012](#snyk-os-go-0012)
<a id="#SNYK-OS-GO-0012"></a>

**Connection reset by peer**

The Go toolchain encountered a connection reset error while trying to access the repository.

This error typically occurs when the connection to the repository is unexpectedly closed by the remote server. This can be due to network issues, server configuration, or other transient problems.

Try to reimport the project, if that does not work, please reach out to Snyk support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-GO-0013](#snyk-os-go-0013)
<a id="#SNYK-OS-GO-0013"></a>

**Invalid zip file**

The Go toolchain encountered an error while trying to download and extract a Go package version. The downloaded file is not a valid zip file.

This error typically occurs when there is an issue with the downloaded file, such as corruption or an incomplete download.

If the package that fails is a private package that's hosted in a private network, make sure the network is up and running and retry the import. If the issue persists, verify the integrity of the downloaded file and check for any issues with the source of the download.

If the issue persists, please reach out to Snyk support.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-GO-0014](#snyk-os-go-0014)
<a id="#SNYK-OS-GO-0014"></a>

**Go version mismatch**

The Go toolchain encountered a version mismatch error while trying to process the module.

This usually happens when the version of Go that was used in the go.mod file is not yet supported by Snyk.

We usually try and add support for new Golang versions shortly after a new one was released.

If the Go version used in the go.mod file is supported based on our Golang documentation, please reach out to Snyk support.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/supported-languages-package-managers-and-frameworks/go/go-for-open-source#go-for-snyk-open-source-support](https://docs.snyk.io/supported-languages-package-managers-and-frameworks/go/go-for-open-source#go-for-snyk-open-source-support)

### [SNYK-OS-GO-0015](#snyk-os-go-0015)
<a id="#SNYK-OS-GO-0015"></a>

**Invalid Go version in go.mod**

The Go toolchain encountered an error while parsing the go.mod file. The specified Go version '__GOLANG_VERSION__' is invalid and must match the format 1.23.4.

This error typically occurs when the Go version in the go.mod file is not correctly formatted.

Ensure that the Go version specified in the go.mod file matches the required format (e.g., 1.23.4). Update the go.mod file with the correct Go version and try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-GO-0016](#snyk-os-go-0016)
<a id="#SNYK-OS-GO-0016"></a>

**Dial TCP timeout**

The Go toolchain encountered a timeout error while trying to fetch a module/package.

This error typically occurs when the connection to the server hosting the module/package times out. This can be due to network issues, server configuration, or the server being unreachable.

Ensure that the server hosting the module is reachable and that there are no network issues. If the issue persists, check if there are any restrictions or firewalls blocking access to the server.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-GO-0017](#snyk-os-go-0017)
<a id="#SNYK-OS-GO-0017"></a>

**Host key verification failed**

The Go toolchain encountered an error while trying to fetch a module/package. The host key verification failed, which means the key used to connect to the remote repository could not be verified.

This error typically occurs when the key is not recognized or is incorrect. It can also happen if the remote repository's host key has changed.

Ensure that the key used to connect to the remote repository is correct and properly configured. You may need to update the known_hosts file to include the correct host key for the remote repository. Verify your repository access configuration and try again. 

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.github.com/en/authentication/troubleshooting-ssh/error-host-key-verification-failed](https://docs.github.com/en/authentication/troubleshooting-ssh/error-host-key-verification-failed)

### [SNYK-OS-GO-0018](#snyk-os-go-0018)
<a id="#SNYK-OS-GO-0018"></a>

**Missing module declaration in go.mod**

The Go toolchain encountered an error while reading the go.mod file. The error indicates that the module declaration is missing.

This error typically occurs when the go.mod file does not specify the module path. To resolve this issue, you need to add a module declaration to the go.mod file.

To specify the module path, run the following command: go mod edit -module=example.com/mod

Update the go.mod file with the correct module path and try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-GO-0019](#snyk-os-go-0019)
<a id="#SNYK-OS-GO-0019"></a>

**Go module version constraint not met**

The Go toolchain encountered a version constraint violation while resolving dependencies.

This error typically occurs when a module or dependency specifies a version requirement that is not satisfied by the current Go version in use. For example, a module may require "go >= 1.22.0", but your environment is running "go 1.21.0".

To resolve this, ensure your Go toolchain version meets the module's specified constraint. This might involve upgrading or downgrading your Go version from Settings -> Snyk Open Source -> Edit settings for Go, or finding an alternative version of the module that is compatible with your current Go environment.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0001](#snyk-os-maven-0001)
<a id="#SNYK-OS-MAVEN-0001"></a>

**Missing property**

The required property is missing from the pom object.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0002](#snyk-os-maven-0002)
<a id="#SNYK-OS-MAVEN-0002"></a>

**Unable to resolve value for property**

The targeted property could not be resolved with a valid value.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0003](#snyk-os-maven-0003)
<a id="#SNYK-OS-MAVEN-0003"></a>

**Unable to resolve version for property**

The targeted property could not be resolved with a valid version.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-MAVEN-0004](#snyk-os-maven-0004)
<a id="#SNYK-OS-MAVEN-0004"></a>

**Cyclic property detected in POM file**

There is circular dependency among properties in the Maven project's configuration file (POM), preventing proper resolution and causing an error.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0005](#snyk-os-maven-0005)
<a id="#SNYK-OS-MAVEN-0005"></a>

**Error parsing the XML file**

There is an error parsing the XML file. This could be referring to either pom.xml or maven-metadata.xml.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0006](#snyk-os-maven-0006)
<a id="#SNYK-OS-MAVEN-0006"></a>

**Invalid coordinates provided**

The coordinates provided for a project were invalid.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0007](#snyk-os-maven-0007)
<a id="#SNYK-OS-MAVEN-0007"></a>

**Skipping group**

Skipping a specific groupId starting due to remapped coordinates.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0008](#snyk-os-maven-0008)
<a id="#SNYK-OS-MAVEN-0008"></a>

**Pom file not found**

The pom file was not found in Maven repository.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0009](#snyk-os-maven-0009)
<a id="#SNYK-OS-MAVEN-0009"></a>

**Missing project from POM**

A project element is missing from POM.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0010](#snyk-os-maven-0010)
<a id="#SNYK-OS-MAVEN-0010"></a>

**Cannot resolve the target POM from the input XML**

Cannot resolve the targeted POM from the input XML.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0011](#snyk-os-maven-0011)
<a id="#SNYK-OS-MAVEN-0011"></a>

**Cannot resolve the target POM from the repository**

Cannot resolve the targeted POM from the repository.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-MAVEN-0012](#snyk-os-maven-0012)
<a id="#SNYK-OS-MAVEN-0012"></a>

**Cannot get the build file repository**

Cannot get the build file repository.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-MAVEN-0013](#snyk-os-maven-0013)
<a id="#SNYK-OS-MAVEN-0013"></a>

**Unable to create hosted git info**

Cannot create source URL.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-MAVEN-0014](#snyk-os-maven-0014)
<a id="#SNYK-OS-MAVEN-0014"></a>

**No released version for versions range**

There was no version released for the specified versions range.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0015](#snyk-os-maven-0015)
<a id="#SNYK-OS-MAVEN-0015"></a>

**Source is not supported**

The source used is not supported by fetcher. The supported sources are: github, bitbucket, gitlab.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0016](#snyk-os-maven-0016)
<a id="#SNYK-OS-MAVEN-0016"></a>

**Timeout when processing the dependency tree**

There was an timeout when processing the dependency tree.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-MAVEN-0017](#snyk-os-maven-0017)
<a id="#SNYK-OS-MAVEN-0017"></a>

**Cannot reach one or more Maven repositories configured under your Snyk organisations language settings**

One or more of the Maven repositories configured under your organisations language settings cannot be reached.

This error can happen for a variety of reasons:

* If using broker it could be a misconfiguration in your broker client. Double check the username and password. 
* It could be network connectivity between the broker client and Snyk or between the broker client and the configured repository, check your firewall rules.

In order to solve this issue, refer to the specific details of this error message to identify which repository is causing issues. 

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://docs.snyk.io/integrate-with-snyk/package-repository-integrations](https://docs.snyk.io/integrate-with-snyk/package-repository-integrations)

### [SNYK-OS-NODEJS-0001](#snyk-os-nodejs-0001)
<a id="#SNYK-OS-NODEJS-0001"></a>

**No repository found for A NPM package**

No repository found for the NPM package.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0002](#snyk-os-nodejs-0002)
<a id="#SNYK-OS-NODEJS-0002"></a>

**Could not parse NPM registry URL**

Could not parse NPM registry URL.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0003](#snyk-os-nodejs-0003)
<a id="#SNYK-OS-NODEJS-0003"></a>

**Could not find a broker resolved URL**

Could not find a broker resolved URL.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0004](#snyk-os-nodejs-0004)
<a id="#SNYK-OS-NODEJS-0004"></a>

**Unable to replace broker URL**

Unable to replace all broker urls in lock file.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0005](#snyk-os-nodejs-0005)
<a id="#SNYK-OS-NODEJS-0005"></a>

**Bad NPM version**

The NPM version is not supported.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0006](#snyk-os-nodejs-0006)
<a id="#SNYK-OS-NODEJS-0006"></a>

**Unknown blob encoding on Github**

Unknown blob encoding on Github.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0007](#snyk-os-nodejs-0007)
<a id="#SNYK-OS-NODEJS-0007"></a>

**No result from forked process**

No result from forked process.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-NODEJS-0008](#snyk-os-nodejs-0008)
<a id="#SNYK-OS-NODEJS-0008"></a>

**Child Process Execution Error**

The child process encountered an error during execution.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-NODEJS-0009](#snyk-os-nodejs-0009)
<a id="#SNYK-OS-NODEJS-0009"></a>

**No valid package upgrades**

The system attempted to find valid upgrades for the packages specified in the lock file, but none were available.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0010](#snyk-os-nodejs-0010)
<a id="#SNYK-OS-NODEJS-0010"></a>

**No dependency updates**

There are no available updates for the dependencies.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0011](#snyk-os-nodejs-0011)
<a id="#SNYK-OS-NODEJS-0011"></a>

**Could not parse JSON file**

An error occurred while attempting to parse a JSON file.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0012](#snyk-os-nodejs-0012)
<a id="#SNYK-OS-NODEJS-0012"></a>

**Could not Base64 encode**

An error occurred while attempting to perform Base64 encoding.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0013](#snyk-os-nodejs-0013)
<a id="#SNYK-OS-NODEJS-0013"></a>

**Could not Base64 decode**

An error occurred while attempting to perform Base64 decoding.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0014](#snyk-os-nodejs-0014)
<a id="#SNYK-OS-NODEJS-0014"></a>

**Missing supported file**

Could not find supported file.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OS-NODEJS-0015](#snyk-os-nodejs-0015)
<a id="#SNYK-OS-NODEJS-0015"></a>

**Invalid configuration**

The configuration parameter does not meet the expected data type. Please ensure the provided value is of the correct data type.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OS-NODEJS-0016](#snyk-os-nodejs-0016)
<a id="#SNYK-OS-NODEJS-0016"></a>

**Out of Sync Error**

Sometimes a project may become out of sync between the lockfile and the manifest file. This might happen if the package.json is modified or updated but the pnpm-lock.yaml is not. 

This can be resolved by ensuring the lockfile and manifest file are correctly synced, by executing pnpm install.

In some cases, it may be necessary to delete the node_modules folder and the pnpm-lock.yaml and run pnpm install again to force a full reinstall. 

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://support.snyk.io/s/article/Out-of-sync-manifest--lockfile-in-the-project](https://support.snyk.io/s/article/Out-of-sync-manifest--lockfile-in-the-project)

### [SNYK-OS-NODEJS-0017](#snyk-os-nodejs-0017)
<a id="#SNYK-OS-NODEJS-0017"></a>

**Unsupported pnpm lockfile version**

The lockfile version is not supported. Supported lockfile versions for pnpm include v5 and v6.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-NODEJS-0019](#snyk-os-nodejs-0019)
<a id="#SNYK-OS-NODEJS-0019"></a>

**Yarn package not found**

Snyk could not find the package in the Yarn registry.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-NODEJS-0020](#snyk-os-nodejs-0020)
<a id="#SNYK-OS-NODEJS-0020"></a>

**Unable to reach package registry**

Snyk could not reach the node package registry.

**HTTP Status:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

### [SNYK-OS-NODEJS-0021](#snyk-os-nodejs-0021)
<a id="#SNYK-OS-NODEJS-0021"></a>

**Lock file is outdated**

The lock file is outdated. Update the lock file and try again.

**HTTP Status:** [409](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409)

### [SNYK-OS-NODEJS-0022](#snyk-os-nodejs-0022)
<a id="#SNYK-OS-NODEJS-0022"></a>

**Unable to read from remote repository**

Snyk does not have sufficient permissions to access the repository, or the repository does not exist.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

### [SNYK-OS-PYTHON-0001](#snyk-os-python-0001)
<a id="#SNYK-OS-PYTHON-0001"></a>

**Unsupported manifest file type for remediation**

The provided requirements file is not supported by Snyk for Python.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/python](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/python)

### [SNYK-OS-PYTHON-0002](#snyk-os-python-0002)
<a id="#SNYK-OS-PYTHON-0002"></a>

**Received more manifests than expected**

Too many manifest files were provided in the request body.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0003](#snyk-os-python-0003)
<a id="#SNYK-OS-PYTHON-0003"></a>

**Failed to apply dependency updates**

An error occurred while updating dependencies.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0004](#snyk-os-python-0004)
<a id="#SNYK-OS-PYTHON-0004"></a>

**Python package not found**

A package listed in the manifest file cannot be found in the Python Package Index(PyPI).
Make sure all packages included in the manifest file are public existing ones.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0005](#snyk-os-python-0005)
<a id="#SNYK-OS-PYTHON-0005"></a>

**Syntax errors found in manifest file**

The manifest file has syntax issues like incorrect package names or unsupported characters.
Make sure the manifest file follows the syntax stardards and can be installed locally as well.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0006](#snyk-os-python-0006)
<a id="#SNYK-OS-PYTHON-0006"></a>

**Python version not supported**

At least one of the packages requires a Python version that doesn't match the one used in the project scan.
Make sure to select a suitable Python version from the organization Python language settings.
Alternatively, add a `.snyk` file for Python version selection override.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0007](#snyk-os-python-0007)
<a id="#SNYK-OS-PYTHON-0007"></a>

**Packages versions caused conflicts**

Two or more packages have conflicting version requirements that cannot be resolved.
Make sure no two packages and their requirements cause conflicts and that the manifest file can be installed locally.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0008](#snyk-os-python-0008)
<a id="#SNYK-OS-PYTHON-0008"></a>

**No matching distribution found for one or more of the packages**

At least one of the packages requires a Python version that doesn't match the one used in the project scan.
Make sure to select a suitable Python version from the organization Python language settings.
Alternatively, add a `.snyk` file for Python version selection override.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0009](#snyk-os-python-0009)
<a id="#SNYK-OS-PYTHON-0009"></a>

**Packages installation failed**

Some packages failed during installation due to missing system dependencies, compilation errors, or other package-specific issues.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0010](#snyk-os-python-0010)
<a id="#SNYK-OS-PYTHON-0010"></a>

**Python version not supported**

At least one of the packages requires a Python version that doesn't match the one used in the project scan.
Make sure to use the correct python version in the requires section of the Pipfile.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0011](#snyk-os-python-0011)
<a id="#SNYK-OS-PYTHON-0011"></a>

**No matching distribution found for one or more of the packages**

At least one of the packages requires a Python version that doesn't match the one used in the project scan.
Make sure to use the correct python version in the requires section of the Pipfile.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0012](#snyk-os-python-0012)
<a id="#SNYK-OS-PYTHON-0012"></a>

**The 10 GB space limit for downloaded Python dependencies has been exceeded**

The total size of the downloaded Python dependencies in the manifest file exceeds the 10GB limit.
This often happens due to the large size of some of the Python dependencies and is usually the case for Python packages that require NVIDIA drivers like PyTorch. 

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/git-repositories-and-python#pip-and-git-repositories](https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/git-repositories-and-python#pip-and-git-repositories)

### [SNYK-OS-RUBY-0001](#snyk-os-ruby-0001)
<a id="#SNYK-OS-RUBY-0001"></a>

**Cyclic dependency detected in lockfile**

Cyclic dependency detected in lockfile.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-RUBY-0002](#snyk-os-ruby-0002)
<a id="#SNYK-OS-RUBY-0002"></a>

**Gem not found**

A gem listed in the Gemfile cannot be found in the RubyGems repository or locally.
Make sure all gems included in the Gemfile are publicly available or properly configured in your gem sources.
Verify that the gem name and version are correct, and check that you have access to any private gem repositories.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-RUBY-0003](#snyk-os-ruby-0003)
<a id="#SNYK-OS-RUBY-0003"></a>

**Gem version conflict**

Bundler was unable to resolve compatible versions for the gems specified in the Gemfile.
This occurs when multiple gems have conflicting version requirements that cannot be satisfied simultaneously.
Review your Gemfile and Gemfile.lock to identify conflicting dependencies, and consider updating gem versions or constraints to resolve the conflict.

**HTTP Status:** [409](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409)

### [SNYK-OS-SETTINGS-0001](#snyk-os-settings-0001)
<a id="#SNYK-OS-SETTINGS-0001"></a>

**Reachability settings not enabled**

The reachability settings are not enabled for your Organization. You can enable them from the Setttings page or you can switch to an Organization where the reachability settings are already enabled.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

**Help Links:**

* [https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/reachability-analysis](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/reachability-analysis)

***

## Builds

### [SNYK-OS-8001](#snyk-os-8001)
<a id="#SNYK-OS-8001"></a>

**Invalid request**

The provided request payload is not valid for the selected ecosystem. Please review the API documentation.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://apidocs.snyk.io/](https://apidocs.snyk.io/)

### [SNYK-OS-8002](#snyk-os-8002)
<a id="#SNYK-OS-8002"></a>

**Build environment not found**

The build environment for the provided context could not be found. Please ensure you have created the build environment first.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-8003](#snyk-os-8003)
<a id="#SNYK-OS-8003"></a>

**Unsupported Ecosystem**

The language or package manager is not supported. Please refer to the supported package managers in the links.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source)

### [SNYK-OS-8004](#snyk-os-8004)
<a id="#SNYK-OS-8004"></a>

**OAuth re-authorization required**

Your code is cloned on an isolated environment using Git as it is required by Snyk to analyze its dependencies.

Your Organization has enabled or enforced SAML SSO after you authorized Snyk to access your code, and a re-authentication is therefore required.

The error you're seeing is usually reproducible by attempting to do a `git clone` of your repository with incorrectly configured credentials.
Verify your authentication configuration with your Git cloud provider and try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso](https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso)

### [SNYK-OS-8005](#snyk-os-8005)
<a id="#SNYK-OS-8005"></a>

**Project too large to be processed**

The project cannot be built or processed due to requiring more memory than available. 
For node projects, please try again after removing requirement to generate a lockfile when opening a fix PR.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#howtonotrequestalockfiletobegenerated](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#howtonotrequestalockfiletobegenerated)

### [SNYK-OS-8006](#snyk-os-8006)
<a id="#SNYK-OS-8006"></a>

**No default image found in repository**

Unable to find the default image. Please try again, and contact Snyk support if the error persists.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

***

## SBOM Export

### [SNYK-OS-9000](#snyk-os-9000)
<a id="#SNYK-OS-9000"></a>

**SBOM generation export server error**

An unexpected error occurred during the SBOM generation. Review the request, then try again. If the error persists, contact Snyk Support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-9001](#snyk-os-9001)
<a id="#SNYK-OS-9001"></a>

**Dependency graph error**

An unexpected dependency graph error occurred. Review the request, then try again. If the error persists, contact Snyk Support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-9002](#snyk-os-9002)
<a id="#SNYK-OS-9002"></a>

**Error parsing dependency graph**

The dependency graph cannot be parsed due to an unexpected error. Review the request, then try again. If the error persists, contact Snyk Support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-9003](#snyk-os-9003)
<a id="#SNYK-OS-9003"></a>

**SBOM not supported due to project type**

Only SBOMs for Snyk Open Source or Snyk Container projects are supported.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-9004](#snyk-os-9004)
<a id="#SNYK-OS-9004"></a>

**SBOM not supported**

Only SBOMs for open source projects are supported (Snyk Open Source).

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-9005](#snyk-os-9005)
<a id="#SNYK-OS-9005"></a>

**Dependency graph request cannot be processed**

The server cannot process the request due to incomplete data. Review the request, then try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-9006](#snyk-os-9006)
<a id="#SNYK-OS-9006"></a>

**Authorization failed due to missing API token**

The API token is misconfigured or expired. Configure or generate the API token, then try again.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**Help Links:**

* [https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens](https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens)

### [SNYK-OS-9007](#snyk-os-9007)
<a id="#SNYK-OS-9007"></a>

**Client request cannot be processed**

The body of the request is empty. Review the request, then try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OS-9008](#snyk-os-9008)
<a id="#SNYK-OS-9008"></a>

**Invalid dependency graph**

The supplied dependency graph was not valid. Review the request, then try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

***

## Open Source Unmanaged

### [SNYK-OSJVM-001](#snyk-osjvm-001)
<a id="#SNYK-OSJVM-001"></a>

**Maven search service unavailable**

The upstream Maven search service is not available.

**HTTP Status:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

**Help Links:**

* [https://search.maven.org](https://search.maven.org)
* [https://status.maven.org](https://status.maven.org)

### [SNYK-OSJVM-002](#snyk-osjvm-002)
<a id="#SNYK-OSJVM-002"></a>

**SHA1 not found**

Unable to find the coordinates for the provided SHA1. Please verify the data you are sending and try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files)

***

## PURL Vulnerabilities

### [SNYK-OSSI-1040](#snyk-ossi-1040)
<a id="#SNYK-OSSI-1040"></a>

**Your Organisation is not authorized to perform this action**

You likely don’t have access to the features in Beta. To get access, you can request access to features in Beta through your account manager or team.

**HTTP Status:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

### [SNYK-OSSI-1050](#snyk-ossi-1050)
<a id="#SNYK-OSSI-1050"></a>

**Authorization request failure**

Unexpected error when authenticating. Try again, and if the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-2010](#snyk-ossi-2010)
<a id="#SNYK-OSSI-2010"></a>

**Invalid purl**

Make sure that the purl is valid. See the Package URL specification link for further information.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst](https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst)

### [SNYK-OSSI-2011](#snyk-ossi-2011)
<a id="#SNYK-OSSI-2011"></a>

**Namespace not specified**

You have requested a package type that requires a namespace (e.g. maven group id). Provide the namespace to retrieve the package.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst](https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst)

### [SNYK-OSSI-2020](#snyk-ossi-2020)
<a id="#SNYK-OSSI-2020"></a>

**Unsupported ecosystem**

The package type is not supported. Check the List issues for a package in Snyk API.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2021](#snyk-ossi-2021)
<a id="#SNYK-OSSI-2021"></a>

**Purl components required**

A list of components of the purl spec is required. The purl did not specify all the required components. Please add the missing components to the purl and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2022](#snyk-ossi-2022)
<a id="#SNYK-OSSI-2022"></a>

**Unsupported purl components**

Remove the unsupported component and retry the request.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2030](#snyk-ossi-2030)
<a id="#SNYK-OSSI-2030"></a>

**Requested package not found**

The package you specified in the purl cannot be found in the vulnerability database. Check the package name, ecosystem, and version, then try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OSSI-2031](#snyk-ossi-2031)
<a id="#SNYK-OSSI-2031"></a>

**Vulnerability service not available**

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

### [SNYK-OSSI-2032](#snyk-ossi-2032)
<a id="#SNYK-OSSI-2032"></a>

**This issue is unexpected and the service should recover quickly if not please contact support**

An unexpected error occurred. Please try again, and if you continue to experience issues please contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-2033](#snyk-ossi-2033)
<a id="#SNYK-OSSI-2033"></a>

**This issue is unexpected and the service should recover quickly if not please contact support**

An unexpected error occurred with the vulnerability service. Please try again, and if you continue to experience issues please contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-2040](#snyk-ossi-2040)
<a id="#SNYK-OSSI-2040"></a>

**Request not processed due to unexpected error**

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-2041](#snyk-ossi-2041)
<a id="#SNYK-OSSI-2041"></a>

**Invalid pagination parameters**

The pagination limit is > 1 and ≤ 1000, and the offset is ≥0.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2042](#snyk-ossi-2042)
<a id="#SNYK-OSSI-2042"></a>

**purls exceed limit**

The number of purls sent in the request exceeds the limit of 1000 set by the service.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2043](#snyk-ossi-2043)
<a id="#SNYK-OSSI-2043"></a>

**Number of issues exceeds limit**

The number of issues found for the provided purls exceeds the limit defined by the API. Reduce the number of purls sent in a single request.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2044](#snyk-ossi-2044)
<a id="#SNYK-OSSI-2044"></a>

**Expected distro to be present**

The given Package URL does not have a required distro qualifier.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**Help Links:**

* [https://docs.snyk.io/scan-containers/how-snyk-container-works/supported-operating-system-distributions#debian](https://docs.snyk.io/scan-containers/how-snyk-container-works/supported-operating-system-distributions#debian)

### [SNYK-OSSI-2045](#snyk-ossi-2045)
<a id="#SNYK-OSSI-2045"></a>

**Unsupported Debian distro**

This Debian distro is currently not supported.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2046](#snyk-ossi-2046)
<a id="#SNYK-OSSI-2046"></a>

**Expected namespace to be present**

The given Package URL does not have a required namespace.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2047](#snyk-ossi-2047)
<a id="#SNYK-OSSI-2047"></a>

**Unsupported vendor**

The given Package URL does not contain a supported vendor. Please use one of the listed vendors and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2048](#snyk-ossi-2048)
<a id="#SNYK-OSSI-2048"></a>

**Unsupported Alpine distro**

This Alpine distro is currently not supported.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

***

## Open Source Project Issues

### [SNYK-OSSI-OSPI-1001](#snyk-ossi-ospi-1001)
<a id="#SNYK-OSSI-OSPI-1001"></a>

**Invalid request**

Check the body of your request and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-OSPI-1002](#snyk-ossi-ospi-1002)
<a id="#SNYK-OSSI-OSPI-1002"></a>

**Unable to return valid API response**

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPI-2001](#snyk-ossi-ospi-2001)
<a id="#SNYK-OSSI-OSPI-2001"></a>

**Failed to process data**

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPI-3001](#snyk-ossi-ospi-3001)
<a id="#SNYK-OSSI-OSPI-3001"></a>

**Failed to store issue data**

Check inputs and then try again. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPI-4001](#snyk-ossi-ospi-4001)
<a id="#SNYK-OSSI-OSPI-4001"></a>

**Internal server error**

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

***

## Open Source Project Snapshots

### [SNYK-OSSI-OSPSS-1001](#snyk-ossi-ospss-1001)
<a id="#SNYK-OSSI-OSPSS-1001"></a>

**Invalid request**

Check the body of your request and try again.

**HTTP Status:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-OSPSS-1002](#snyk-ossi-ospss-1002)
<a id="#SNYK-OSSI-OSPSS-1002"></a>

**Unable to return valid API response**

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPSS-2001](#snyk-ossi-ospss-2001)
<a id="#SNYK-OSSI-OSPSS-2001"></a>

**Failed to process data**

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPSS-3001](#snyk-ossi-ospss-3001)
<a id="#SNYK-OSSI-OSPSS-3001"></a>

**Failed to store snapshot data**

Check inputs and then try again. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPSS-4001](#snyk-ossi-ospss-4001)
<a id="#SNYK-OSSI-OSPSS-4001"></a>

**Internal server error**

This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

***

## Policies

### [SNYK-POLICY-0001](#snyk-policy-0001)
<a id="#SNYK-POLICY-0001"></a>

**Unable to apply a policy with an invalid configuration**

Snyk could not apply a policy whilst executing a test because the configuration for the policy was invalid.
You may be able to fix the policy and try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/manage-risk/policies](https://docs.snyk.io/manage-risk/policies)

***

## PRChecks

### [SNYK-PR-CHECK-0001](#snyk-pr-check-0001)
<a id="#SNYK-PR-CHECK-0001"></a>

**Error reading manifest**

Snyk failed to read 1 or more manifest files.
Sometimes things go wrong: a flaky connection, 3rd party services go down and Snyk is unable to read the files needed in order to test your project. 

If this happens, you could try:

- Closing and re-opening your Pull Request / Merge Request, to kick off a new test
- Change the status of Pull Request / Merge Request to draft, then change it back to "Ready for review", to kick off a new test
- Removing and re-adding the repo to Snyk

Ultimately, you should contact support@snyk.io if the issue persists

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://support.snyk.io/s/article/Failed-to-read-manifest-file---Commit-Status](https://support.snyk.io/s/article/Failed-to-read-manifest-file---Commit-Status)

### [SNYK-PR-CHECK-0002](#snyk-pr-check-0002)
<a id="#SNYK-PR-CHECK-0002"></a>

**Manifest not found**

Snyk uses your project manifest file to analyze your projects for vulnerabilities. When you import a project for monitoring, Snyk scans the project to locate the manifest file and then remembers where that file is. 
When a project manifest file is moved or deleted, we still try to look for in it in the last known location in order to run tests on commit statuses. If we can't find the file, this error can occur.

If this happens, you could try the following:
1. Delete the matching project from your account in the Snyk app (UI or CLI).
2. Now import the same project from scratch.

As during the original import, Snyk scans the project and locates the manifest file.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://support.snyk.io/s/article/Manifest-not-found](https://support.snyk.io/s/article/Manifest-not-found)

### [SNYK-PR-CHECK-0003](#snyk-pr-check-0003)
<a id="#SNYK-PR-CHECK-0003"></a>

**Rate limit hit while testing project**

Snyk makes requests to your SCM when testing a project, in order to analyze your projects for vulnerabilities. If we need to make a lot of requests in a short time period, we may encounter third party rate limits, and this error can occur.

If you receive any of these errors, try re-running the tests, by closing and reopening the pull request.

**HTTP Status:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)

### [SNYK-PR-CHECK-0004](#snyk-pr-check-0004)
<a id="#SNYK-PR-CHECK-0004"></a>

**Out of Sync Error**

Sometimes a project may become out of sync between the lockfile and the manifest file. This might happen if the package.json is modified or updated but the lockfile is not. 

This can be resolved by ensuring the lockfile and manifest file are correctly synced, by executing npm install or yarn install.

In some cases, it may be necessary to delete the node_modules folder and the package-lock.json and run npm install again to force a full reinstall. 

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://support.snyk.io/s/article/Out-of-sync-manifest--lockfile-in-the-project](https://support.snyk.io/s/article/Out-of-sync-manifest--lockfile-in-the-project)

### [SNYK-PR-CHECK-0005](#snyk-pr-check-0005)
<a id="#SNYK-PR-CHECK-0005"></a>

**Failed determining project target**

An internal error occurred, whereby Snyk was unable to determine the correct target for a given project in your PR Check.

If you receive this error, try re-running the tests, by closing and reopening the pull request.

Ultimately, you should contact support@snyk.io if the issue persists.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-PR-CHECK-0006](#snyk-pr-check-0006)
<a id="#SNYK-PR-CHECK-0006"></a>

**Failed to complete the test**

A "Failed to complete testing check status" appears in your commit checks when an unknown error occurs while Snyk was trying to test your projects for vulnerabilities or license issues.

If you receive this error, try re-running the tests, by closing and reopening the pull request.

Ultimately, you should contact support@snyk.io if the issue persists.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://support.snyk.io/s/article/Unknown-PR-test-error](https://support.snyk.io/s/article/Unknown-PR-test-error)

### [SNYK-PR-CHECK-0007](#snyk-pr-check-0007)
<a id="#SNYK-PR-CHECK-0007"></a>

**Failed to fetch merge commit SHA**

In order for snyk test to run, we need the merge commit SHA from the GitHub. For some reason, we couldn’t get it.

Try closing and then reopening the pull request, or you can Skip the Pull Request Check if it is consistent.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-PR-CHECK-0008](#snyk-pr-check-0008)
<a id="#SNYK-PR-CHECK-0008"></a>

**Merge conflict error**

Merge Conflict Error is not a Snyk specific issue but rather some issues on your SCM environment. As an example, merge conflicts could happen when people make different changes to the same line of the same file, or when one person edits a file and another person deletes the same file.

To resolve this, you might need to figure out all the merge conflicts on your SCM environment and resolve them to fully remediate these types of errors on Snyk. As a note, this cannot be modified/changed on Snyk's side.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://support.snyk.io/s/article/Merge-conflict-error](https://support.snyk.io/s/article/Merge-conflict-error)

### [SNYK-PR-CHECK-0009](#snyk-pr-check-0009)
<a id="#SNYK-PR-CHECK-0009"></a>

**Failed to detect issues**

Snyk is always trying to check for new issues and vulnerabilities to keep you safe. We do so by testing on your code on webhook Pull Request events and Push events.

Occasionally you might see a "Failed to detect issues" commit status which may block your PR. This means that we tried to run a test against your changes but unfortunately something went wrong / we encountered an internal problem. If this happens to you try recreating the pull request and if it still occurs reach out and let us know which user, organization and project and commit sha you experienced the issue with on support@snyk.io

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://support.snyk.io/s/article/Failed-to-detect-issues](https://support.snyk.io/s/article/Failed-to-detect-issues)

### [SNYK-PR-CHECK-0010](#snyk-pr-check-0010)
<a id="#SNYK-PR-CHECK-0010"></a>

**No valid credentials to process PR check**

Snyk uses credentials configured on your integration to test your code and to update your PR Check.

If this error occurs, please ensure your integration and credentials are correctly set up, by following the instructions for your SCM here: https://docs.snyk.io/integrate-with-snyk/git-repository-scm-integrations

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

### [SNYK-PR-CHECK-0011](#snyk-pr-check-0011)
<a id="#SNYK-PR-CHECK-0011"></a>

**Failed to generate a commit status**

Snyk is always trying to check for new issues and vulnerabilities to keep you safe. We do so by testing on your code on webhook Pull Request events and Push events.

Occasionally you might see a "Failed to generate a commit status" which may block your PR. This means that we tried to run a test against your changes but unfortunately something went wrong / we encountered an internal problem. If this happens to you try recreating the pull request and if it still occurs reach out and let us know which user, organization and project and commit sha you experienced the issue with on support@snyk.io

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

***

## SBOM Test

### [SNYK-SBOM-0001](#snyk-sbom-0001)
<a id="#SNYK-SBOM-0001"></a>

**SBOM test error**

An unexpected error occurred. Review the request, then try again. If the error persists, contact Snyk Support.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-SBOM-0002](#snyk-sbom-0002)
<a id="#SNYK-SBOM-0002"></a>

**Organization ID mismatch**

The requested organization ID does not match the owner of the SBOM test ID.

This error occurs when the supplied organization ID is different to the one used when creating an SBOM test run.
Ensure the organization ID used to make the request is the same as the the organization ID used to create the SBOM test.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-SBOM-0003](#snyk-sbom-0003)
<a id="#SNYK-SBOM-0003"></a>

**Unable to find SBOM test**

Snyk was unable to find the requested SBOM test.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-SBOM-0004](#snyk-sbom-0004)
<a id="#SNYK-SBOM-0004"></a>

**SBOM test failed**

The SBOM test failed and does not have any results.

This error occurs when results for a failed SBOM test are being requested.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-SBOM-0005](#snyk-sbom-0005)
<a id="#SNYK-SBOM-0005"></a>

**SBOM test results still pending**

The SBOM test is still processing and does not have any results yet.

This error occurs when the results for an SBOM test have been requested, but the SBOM test is still being processed.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-SBOM-0006](#snyk-sbom-0006)
<a id="#SNYK-SBOM-0006"></a>

**Unknown SBOM format**

Snyk does not recognize the SBOM file format.

Provide an SBOM document with a supported format.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/sbom-test](https://docs.snyk.io/snyk-cli/commands/sbom-test)

### [SNYK-SBOM-0007](#snyk-sbom-0007)
<a id="#SNYK-SBOM-0007"></a>

**Unable to process SBOM input**

Snyk is unable to decode the SBOM file. Provide a valid SBOM document and try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/sbom-test](https://docs.snyk.io/snyk-cli/commands/sbom-test)

### [SNYK-SBOM-0008](#snyk-sbom-0008)
<a id="#SNYK-SBOM-0008"></a>

**SBOM format not supported**

Provide a supported format of the SBOM document and try again.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/snyk-cli/commands/sbom-test](https://docs.snyk.io/snyk-cli/commands/sbom-test)

### [SNYK-SBOM-0009](#snyk-sbom-0009)
<a id="#SNYK-SBOM-0009"></a>

**SBOM analysis failed**

Snyk was unable to process the provided SBOM input and is unable to scan it for issues.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-SBOM-0010](#snyk-sbom-0010)
<a id="#SNYK-SBOM-0010"></a>

**No testable packages found**

The SBOM document you provided does not contain any packages supported by Snyk vulnerability analysis, or it does not contain any package.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

***

## SCM

### [SNYK-SCM-0001](#snyk-scm-0001)
<a id="#SNYK-SCM-0001"></a>

**Integration type not supported**

The integration you provided does not support SCM repository access.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**Help Links:**

* [https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations](https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations)

### [SNYK-SCM-0002](#snyk-scm-0002)
<a id="#SNYK-SCM-0002"></a>

**Revision cannot be resolved**

Snyk was unable to resolve the SCM revision you provided. Provide a valid revision, either a full commit ID or an existing commit reference.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-SCM-0003](#snyk-scm-0003)
<a id="#SNYK-SCM-0003"></a>

**Integration authentication failed**

Snyk was unable to authenticate with your SCM provider. Ensure you are using valid credentials for the Snyk integration.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

### [SNYK-SCM-0004](#snyk-scm-0004)
<a id="#SNYK-SCM-0004"></a>

**Integration authorization failed**

Snyk was unable to authorize with your SCM provider. If your Organization has SAML SSO enabled or enforced, re-authorize the OAuth Application `Snyk`.

**HTTP Status:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

### [SNYK-SCM-0005](#snyk-scm-0005)
<a id="#SNYK-SCM-0005"></a>

**Too many files**

Snyk was unable to retrieve the repository because the overall file count exceeds the limit of 40000.

To reduce the file count, use a `.snyk` file to ignore certain directories or files. Alternatively, analyze individual work subdirectories separately.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-SCM-0006](#snyk-scm-0006)
<a id="#SNYK-SCM-0006"></a>

**Repository size too large**

Snyk was unable to retrieve the repository because the size of the repository exceeds 15 GB.

To reduce the overall size of the repository, use a a `.snyk` file to ignore certain directories or files. Alternatively, analyze individual work subdirectories separately.

**HTTP Status:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-SCM-0010](#snyk-scm-0010)
<a id="#SNYK-SCM-0010"></a>

**Specified resource cannot be found**

Snyk was unable to resolve the SCM resource you provided. Provide a valid resource path and revision.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

***

## Target

### [SNYK-TARGET-0001](#snyk-target-0001)
<a id="#SNYK-TARGET-0001"></a>

**Target not found**

Snyk was unable to resolve the imported target. Ensure that Snyk created the target and try again.

**HTTP Status:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**Help Links:**

* [https://docs.snyk.io/snyk-admin/snyk-projects#target](https://docs.snyk.io/snyk-admin/snyk-projects#target)

### [SNYK-TARGET-0002](#snyk-target-0002)
<a id="#SNYK-TARGET-0002"></a>

**No unique target found**

Snyk was unable to resolve a single target. Snyk found multiple targets configured for the same integration and repository URL pair. Ensure a unique target exists.

**HTTP Status:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**Help Links:**

* [https://docs.snyk.io/snyk-admin/snyk-projects#target](https://docs.snyk.io/snyk-admin/snyk-projects#target)
