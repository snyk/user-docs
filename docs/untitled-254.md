# How to install a Beta version of the Eclipse plugin?

##  How to install a Beta version of the Eclipse plugin?

More details: [https://stackoverflow.com/questions/31553376/eclipse-how-to-install-a-plugin-manually](https://stackoverflow.com/questions/31553376/eclipse-how-to-install-a-plugin-manually)

### Uninstall the Snyk Plugin from the Eclipse Marketplace

**Help** &gt; **Eclipse Marketplace** &gt; Navigate to the **Installed** tab to find **Snyk Vuln scanner**

Click **Uninstall** for the Snyk Plugin that you have installed, if applicable

Accept any changes and click Finish

Restart your IDE

### Download the Beta version of the Eclipse Plugin from Github

[https://github.com/snyk/snyk-eclipse-plugin](https://github.com/snyk/snyk-eclipse-plugin)

Click on Releases and select the version of the plugin that you wish to install

Unzip the file to a folder on your system

### Install the plugin in Eclipse

**Help** &gt; **Install New Software...**

Click **Add**, Select **Local** and browse to the downloaded Zip file

Click **Add**

Check the box next to the **Snyk Vuln Scanner** and click on **Next**

Click **Next**, review any prompts and continue.

Once Completed, it should prompt you to restart your IDE once more.

Success!

