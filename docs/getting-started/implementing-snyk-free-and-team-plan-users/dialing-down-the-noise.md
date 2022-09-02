# Dialing down the noise

### What is "noise"?&#x20;

In any discussion about developer tools, and especially security tools, it is important to understand how to reduce “noise” (ensure distractions aren’t being raised for development), and focus on the areas that matter.

Snyk does a lot to reduce noise by default, and also allows you to configure and adjust settings as needed to keep focus, especially as a company grows and the number of applications increases.

### Available dials

Here are some dials that can assist in targeting alerts and ensuring your success/engagement.

* Key alert setting categories:&#x20;
  * Organization Notification Settings&#x20;
    * In your settings for the organization and its members, you can determine if alerts are only for High and Critical, this is very useful to focus on the worst issues. This does not impact reporting, only alerts.&#x20;
  * Personal Notifications Settings&#x20;
    * Navigate to your profile, on the top right of the screen, select notifications. Here you can choose which projects you get notified for. This is critical to ensure you are not getting messages for projects you are not responsible for, and your developers will appreciate knowing they can set that.
* Filtering/Fail criteria&#x20;
  * Snyk provides a multitude of criteria that can help in many situations. For example:&#x20;
    * Pull request checks can fail if there’s Critical/High and if a fix is available (See integration settings)&#x20;
    * CI/CD & CLI - You can filter on --severity-threshold, but also if a fix is available (See --help in the CLI).&#x20;
    *   CI/CD & CLI - Using the [Snyk-Filter](https://github.com/snyk-tech-services/snyk-filter) plugin, you can report on issues if there’s an exploit available, CVSS score, Priority score etc. This plugin gives an incredible amount of flexibility!

        \
        It’s quite common not to fail the build in the first month and to only use "monitor" to start. When you do start failing the build, using the criteria shown above is really helpful
* Automatic Pull Requests and Pull Request Checks&#x20;
  * If you have a large number of projects, initially we recommend you turn this off for all but the important projects. As your teams become more accustomed to this type of capability, you can then slowly turn it on across your repositories.

