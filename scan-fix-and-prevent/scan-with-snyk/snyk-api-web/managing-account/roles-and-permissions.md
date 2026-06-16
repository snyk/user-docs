# Roles and permissions

Roles and permissions overview.

The access to the main features of Snyk API & Web is ruled by permissions. For instance, in order to add a target, the user needs to have the permission `Create Target`.

You can group permissions using [Roles](https://plus.probely.app/roles). For instance, the built-in role `Developer` can view targets, change target settings, change findings, and start target scans, but cannot add targets.

You can then map roles to users at an account level, team level, or at a target level. When you are [adding a user](https://plus.probely.app/users), you can set the user's:

* Global role, which is applied to all targets of the account (the role is set at an account level or global scope).
* Team role, which is applied to all targets of that team.
* Target role, which is applied solely to that target.

As such, you can have a user that has, for instance, the role `Developer` at an account level, the role `SecOps` on a specific team, and the role `Admin` on a specific target.

Learn more about this by visiting [Get started with teams](get-started-with-teams.md).
