# API - Unauthorized 401 error when importing a Github project with a service account

## API - Unauthorized 401 error when importing a Github project with a service account

**Problem**:

Users may get an Unauthorized 401 error \(Invalid credentials\) when importing a project into Github using the API with a service account token.

**Discussion**:

When using your \*user\* API token, there is a workflow that gets used by Github and oAuth. The problem lies that some actions \(like import project\) do not have a workflow for service accounts.

Unfortunately, using a Service Account token is not supported at this time.

**Resolution**:

Currently, the workaround is to use a user token or create a fully signed-in user account as the service account for Github.

