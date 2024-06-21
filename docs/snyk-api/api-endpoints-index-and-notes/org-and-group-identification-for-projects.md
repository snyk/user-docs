# Org and group identification for Projects

The Org Id is defined by Snyk and is the same in both the V1 and REST API. For Projects, the best identifier depends on the Snyk product.

In the Snyk REST API, you can:

* Get a list of groups identified by GroupId using [Get a group](https://apidocs.snyk.io/?version=2024-06-18%7Ebeta#get-/groups/-group\_id-) (**GET** /groups/{group\_id}):&#x20;
* Get the list of Orgs that you have access to, identified by Orgid, using [Get organization](https://apidocs.snyk.io/?#get-/orgs/-org\_id-) (**GET** /orgs/{org\_id}).

To get more data, you can use Snyk API v1:

* Find the Groupid: ([List all the organizations a user belongs to](https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to)).
* Request a list of Orgs (identified by Orgid) based on a Groupid: ([List all organizations in a group](https://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-grouphttps://snyk.docs.apiary.io/#reference/groups/list-all-organizations-in-a-group/list-all-organizations-in-a-group)).
