# Frequently Asked Questions

* How do I get support from Snyk on building a Snyk connector?
  * Our preferred communication is using Slack. We will create a shared channel in Slack and add anyone from your team to the channel. We prefer to use shared channels for communication and support if possible.
* What features are available in each of the different Snyk plans?
  * Snyk plans \(free, standard, pro, and enterprise\) contain different features. Please see the [Snyk plans](https://snyk.io/plans/) on line for feature details.
* Which API \(Project or Reporting\) should I use to get project issues?
  * The [Project](https://snyk.docs.apiary.io/#reference/projects) and [Reporting API](https://snyk.docs.apiary.io/#reference/reporting-api) both provide access to get a list of project issues. Each API takes a different set of inputs to generate a list of project issues. We recommend using the Project API.
* What is a [Project Type](https://snyk.docs.apiary.io/#reference/projects/retrieve-a-single-project), and which types are available in the Snyk API?
  * Each project in Snyk has an associated project type. The type is determined based on the package manager used for the project. SCA based projects will use the package manager used for the application. For example, a Java application using maven will have maven as the project type. For a list of supported languages and project types, please visit our language support page. If the project is a container image, the project type is the OS package manager used by the base image. The types are rpm, deb, apk, and linux. We also support Infrastructure as Code \(IaC\) projects and the type for these projects are k8sconfig, helmconfig, and terraformconfig.
* What is the origin field in a Project?
  * Each project has an origin field to identify the source of the project. The origin types align with the integration to specific DevOps toolchain. A few origin types include GitHub, Docker Hub, and Synk CLI and are based on what integration created the project.

