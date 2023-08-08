# Automatically link between Dockerfile and images using labels

Snyk allows you to link manually or automatically from a Dockerfile to all container images built from it. You can use this to understand the security impact on your running applications and understand which images can be better secured or need to be rebuilt when you take action and update the Dockerfile base image.

## Viewing linked images

These linked images are identified in the **LINKED IMAGES** section of the details for a Project, as shown in the following image.

<figure><img src="../../.gitbook/assets/mceclip3.png" alt="Project showing linked images"><figcaption><p>Project showing linked images</p></figcaption></figure>

Using container registry integration, you can get automatic links between imported images to existing Dockerfile Projects. To do this, check whether the OCI label in the image matches the path of a Dockerfile in the Organization in Snyk.

## How linked images work

At the point of import or re-test, the image is analyzed and scanned for vulnerabilities. Image labels are also retrieved from the image manifest. Snyk then checks whether:

* Image labels defining the Dockerfile location exist:
  * `org.opencontainers.image.source` - URL to the Project repo (mandatory)
  * `io.snyk.containers.image.dockerfile` - path to the Dockerfile, for example,  "/Dockerfile-prod" (optional)
* The Dockerfile Project exists in the same Organization, with a matching repo (and path or /Dockerfile) from the image labels.

If these conditions apply, Snyk automatically creates a link between the image and Dockerfile Projects.

## Automatic update and removal of links

Links are automatically updated if the Dockerfile labels are updated and are targeting a new location. This can happen at the time of re-test or when a recurring test runs.

Links are removed if either the image Project or Dockerfile Project is deleted.

* If the Dockerfile labels are updated so that they target the Dockerfile location without an existing Project in Snyk
* Or if the Dockerfile labels are removed.

## Linking in brokered SCM integrations

For a link to be created, Snyk needs to be able to map the Dockerfile repository URL to the right SCM Organization source. For brokered integrations, this is a bit more complex, as this URL is not available by default.

To create automatic links between container images to Dockerfiles stored in brokered SCMs, enter the URL in the integration page settings:

<figure><img src="../../.gitbook/assets/mceclip0-4-.png" alt="Integration page settings with integration URL"><figcaption><p>Integration page settings with integration URL</p></figcaption></figure>

When the URL is available, Snyk can use it for generating links.
