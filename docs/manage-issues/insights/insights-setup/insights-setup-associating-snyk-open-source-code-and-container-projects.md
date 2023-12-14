# Insights setup: Associating Snyk Open Source, Code, and Container Projects

After you have set up Insights, Snyk can set up the required linking for the chosen application.

To prioritize your Snyk Code and Snyk Open Source vulnerabilities, Snyk needs to understand the relationships between those imported Snyk Projects and the Container Project. Snyk uses the relationships between these Projects as a proxy to understand the composition of the container image.

The container image is the build artifact that is deployed and running on your Kubernetes cluster, so Snyk can map your application from its source code to its deployed state by understanding the following:

* The link between the Source Code and Open Source Dependencies and the **image**.
* What **images** are deployed to Kubernetes and how they are configured.

<figure><img src="../../../.gitbook/assets/Associating Snyk OS, Code and Container projects (1).png" alt="Mapping your application"><figcaption><p>Mapping your application</p></figcaption></figure>

## Use Project tags to link Projects

Add Snyk [Project tags](../../../snyk-admin/introduction-to-snyk-projects/project-tags.md) to all the Projects used by your application, to link these Projects together and allow Insights to represent the whole of the application that you are testing.

To associate two Projects together, add the exact same tag to both Projects. For example, add the same tag to your Snyk Open Source Projects and Snyk Container Project if they are related to each other.

See the examples at the end of this section.

## Requirements for Project tags

* The same tag must be applied to the container image and Code or Open Source Projects,
* The tag must follow the specified format,
* The Projects do not have to be in the same Snyk Organization to be mapped but MUST be in the same Snyk Group,

<figure><img src="../../../.gitbook/assets/Screenshot 2023-06-06 at 23.29.29.png" alt="Project tags"><figcaption><p>Project tags</p></figcaption></figure>

## Examples of Project tags

### **Single repo to a single image**

In this example, there is a single repository containing your package.json which is built into an image called image-A.

<figure><img src="../../../.gitbook/assets/Example - Single repo Single image.png" alt="Example: single repo to single image"><figcaption><p>Example: single repo to single image</p></figcaption></figure>

To map these associations, you would have the following tags set up:

<table><thead><tr><th width="165">Location</th><th width="161">Asset</th><th width="213">Snyk Project Location</th><th width="238">Tag(s)</th><th>Notes</th></tr></thead><tbody><tr><td>github.com/my-team/front-end</td><td><code>package.json</code></td><td><p>Snyk Org: my-team</p><p><br>Project name: package.json</p></td><td><code>component=pkg:github/my-team/front-end@main</code></td><td>The same tag is applied to both sides</td></tr><tr><td><br></td><td>Built container image called <code>image-A</code></td><td><p>Snyk Org: my-team<br></p><p>Project name: image-A</p></td><td><code>component=pkg:github/my-team/front-end@main</code></td><td></td></tr></tbody></table>

### **Multiple repos to a single image**

<figure><img src="../../../.gitbook/assets/Example - Multiple repos Single image.png" alt="Example: multiple repos to single image"><figcaption><p>Example: multiple repos to single image</p></figcaption></figure>

For this scenario, there are two contributing repositories.

There is a front-end repository containing the package.json scanned by Snyk Open Source and a back-end repository containing go code scanned by Snyk Code.

To map these associations, you would have the following tags set up:\\

| Location                     | Asset                                  | Snyk Project Location                                                | Tag(s)                                                                                                                        | Notes                                                                                                                                                       |
| ---------------------------- | -------------------------------------- | -------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| github.com/my-team/front-end | `package.json`                         | <p>Snyk Org: my-team</p><p><br></p><p>Project name: package.json</p> | `component=pkg:github/my-team/front-end@main`                                                                                 | <p><br></p>                                                                                                                                                 |
| github.com/my-team/back-end  | `Go source code`                       | <p>Snyk Org: my-team</p><p><br></p><p>Project name: Code</p>         | `component=pkg:github/my-team/back-end@main`                                                                                  | <p><br></p>                                                                                                                                                 |
| <p><br></p>                  | Built container image called `image-A` | <p>Snyk Org: my-team</p><p><br></p><p>Project name: image-A</p>      | <p><code>component=pkg:github/my-team/front-end@main</code></p><p><code>component=pkg:github/my-team/back-end@main</code></p> | <p>The image has two tags applied, as there are two upstream dependencies which have different tags.<br></p><p>You can apply multiple tags to an image.</p> |

### **Monorepo to many images**

<figure><img src="../../../.gitbook/assets/Example - Multiple repos Single image-2.png" alt="Example: monorepo to multiple images"><figcaption><p>Example: monorepo to multiple images</p></figcaption></figure>

In this example, the application team is using a monorepo approach. The contents of the repository are built into different container images as they may be run separately.

Here we need to further differentiate the tags by scoping them more accurately.

| Location                   | Asset                                                 | Snyk Project Location                                                          | Tag(s)                                                |                                                               |
| -------------------------- | ----------------------------------------------------- | ------------------------------------------------------------------------------ | ----------------------------------------------------- | ------------------------------------------------------------- |
| github.com/my-team/service | `package.json`                                        | <p>Snyk Org: my-team</p><p><br></p><p>Project name: package.json</p>           | `component=pkg:github/my-team/service/front-end@main` | The tag is further scoped by specifying /front-end at the end |
| <p><br></p>                | Built container image called `my-app-frontend:latest` | <p>Snyk Org: my-team</p><p><br></p><p>Project name: my-app-frontend:latest</p> | `component=pkg:github/my-team/service/front-end@main` |                                                               |

| Location                   | Asset                                                | Snyk Project Location                                                         | Tag(s)                                               | Notes                                                        |
| -------------------------- | ---------------------------------------------------- | ----------------------------------------------------------------------------- | ---------------------------------------------------- | ------------------------------------------------------------ |
| github.com/my-team/service | `Go source code`                                     | <p>Snyk Org: my-team</p><p><br></p><p>Project name: Code</p>                  | `component=pkg:github/my-team/service/back-end@main` | The tag is further scoped by specifying /back-end at the end |
| <p><br></p>                | Built container image called `my-app-backend:latest` | <p>Snyk Org: my-team</p><p><br></p><p>Project name: my-app-backend:latest</p> | `component=pkg:github/my-team/service/back-end@main` |                                                              |

## Tags application

Project Tags can be applied as follows:

<table><thead><tr><th width="185.33333333333331">Product</th><th width="197">Method</th><th>Project Source</th></tr></thead><tbody><tr><td>Snyk Code</td><td><ul><li>API</li></ul></td><td><p>For projects created by</p><ul><li>Git Import</li></ul></td></tr><tr><td>Snyk Open Source</td><td><ul><li>API</li><li>UI</li><li>CLI</li></ul></td><td><p>For projects created by</p><ul><li>Git Import</li><li>CLI Monitor</li></ul></td></tr><tr><td>Snyk Container</td><td><ul><li>API</li><li>UI</li><li>CLI</li></ul></td><td><p>For projects created by</p><ul><li>Git Import</li><li>CLI Monitor</li><li>Container Registry Integration</li></ul></td></tr></tbody></table>

{% hint style="info" %}
Snyk recommends applying the tags through the API because the process can be automated using the API.
{% endhint %}

## UI example for Project tags

<div align="center">

<figure><img src="../../../.gitbook/assets/uiExample.png" alt="UI example for Project tags"><figcaption><p>UI example for Project tags</p></figcaption></figure>

</div>

### CLI Example

`snyk monitor --project-tags=component=pkg:github/my-team/back-end@main`

### API Example

See the [Project Update documentation](https://snyk.docs.apiary.io/#reference/integrations/integration-settings/update-a-project).

Example:

```
 "tags": [
    {
      "key": "component",
      "value": "pkg:github/my-team/back-end@main"
    }
  ],
```

{% hint style="warning" %}
To ensure you have set up your Kubernetes Connector properly, navigate to the **Set up Insights** tab on the **Insights** page and check the **Image composition** section to view the data Insights has access to.
{% endhint %}
