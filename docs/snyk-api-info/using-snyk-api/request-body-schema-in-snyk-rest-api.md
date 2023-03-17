# Request body schema in Snyk REST API

The Snyk REST API endpoint [Invite a user to an organization](https://apidocs.snyk.io/?version=2022-10-06%7Ebeta#post-/orgs/-org\_id-/invites) shows the REQUEST BODY SCHEMA OBJECT as follows:

```
{ 
  email*: string 
  role*: uuid 
}
```

`email` is the invitee email address. `role` is the role public ID that will be granted to an invitee on acceptance.

The variable type for both `email` and `role` is a `string`. The role is a `string` with the format of `uuid`. The Snyk REST API uses these formats; thus the role is defined with the string format.
