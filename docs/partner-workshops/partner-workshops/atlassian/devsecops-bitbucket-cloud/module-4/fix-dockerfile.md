# Fix Dockerfile

From the Snyk app, we will go to the **Projects** menu where we can expand each integration and have a holistic view of our project. Here we will select the _**container image**_ under our Amazon ECR repository.

![](../../../../.gitbook/assets/snyk-projects-02.png)

You will notice a message instructing you to complete one minor configuration item. Let's address that by clicking the **Settings** tab.

![](../../../../.gitbook/assets/snyk-docker-fix-01.png)

Proceed to click on the **Configure Dockerfile** button.

![](../../../../.gitbook/assets/snyk-docker-fix-02.png)

Select **Bitbucket Cloud** as your source.

![](../../../../.gitbook/assets/snyk-docker-fix-03.png)

Select your repository and click **Update Dockerfile**.

![](../../../../.gitbook/assets/snyk-docker-fix-04.png)

Update the default path with the path to our Dockerfile. In this case, the path is `/app/goof/Dockerfile` or as shown below:

![](../../../../.gitbook/assets/snyk-docker-fix-05.png)

You will receive a confirmation message stating that the settings have been successfully applied.

![](../../../../.gitbook/assets/snyk-docker-fix-06.png)

Recommendations for base image upgrade will be provided to you. Here, you will notice the **Current image** defined in your _Dockerfile_ and a **Major upgrade** suggestion to reduce the total number of vulnerabilities on your container image.

![](../../../../.gitbook/assets/snyk-docker-fix-07.png)

For the purpose of this exercise, we will keep things simple and use Bitbucket's built-in editor to make the change. Let's navigate to our _Dockerfile_ in our Bitbucket repo. The path will be `./app/goof/Dockerfile`. Here we can _Edit_ the file and save our changes.

![](../../../../.gitbook/assets/bitbucket-edit-dockerfile.png)

Let's update **Line 1** as follows:

* **OLD VALUE:** `node:6-stretch`
* **NEW VALUE:** `node:12.18-stretch`

Click **Commit**.

