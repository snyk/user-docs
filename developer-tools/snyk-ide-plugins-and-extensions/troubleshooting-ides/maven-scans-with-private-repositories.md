# Maven scans with private repositories

## Problem <a href="#problem" id="problem"></a>

Maven fails due to dependencies that reference private repositories in the IDE plugin or CI/CD pipeline.

## Investigation <a href="#investigation" id="investigation"></a>

* Try running CLI standalone with the repository.
* Check to see if Maven profiles or Maven settings are used for running the build (for example, located in `.mvn` ).
* Check to see if the pipeline has access to the private sources.

Dependencies need to be downloaded from a private repository and the Snyk CLI shows that the download fails in the error message.

## Solution <a href="#solution" id="solution"></a>

You can establish Maven settings or a profile containing the connection information, for example, with `-- -s your-home-directory/.m2/settings.xml` or `-- -Dprofile=my-profile`.

You can pass these parameters on to the CLI using the `Additional Parameters` plugin setting in the IDE. For CI/CD, pass them as a pipeline argument if you use the Snyk CLI directly, or provide a correct path mapping if you run it in a Docker image.
