# Org identification and projects in Snyk APIs

The Org Id is defined by Snyk in Snyk and thus is the same in both Snyk API v1 and the Snyk REST API. For projects, the best identifier depends on the Snyk product.

In the Snyk REST API, you can:

* Get a list of groups identified by GroupId (**GET** /groups/{group\_id}).
* Get the list of Orgs that you have access to, identified by Orgid (**GET** /orgs/{org\_id}).

**Note:** These are [experimental](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#overview) endpoints.

To get more data, you can use Snyk API v1:

* Find the Groupid: ([List all the organizations a user belongs to](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to)).
* Request a list of Orgs (identified by Orgid) based on a Groupid: ([List all organizations in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-grouphttps://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)).
