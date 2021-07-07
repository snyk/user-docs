# Snyk changed the "resolved" URL's in my Lock file

##  Snyk changed the "resolved" URL's in my Lock file

In a case, that snyk is changing your "resolved" URL's in my Lock file, it means that you are using a private repository for your JS packages, so snyk doesn't know what is the right URL of the hosted packages in your private repository.

In this case, you can set Snyk not relock the lock file in order not to change the resolved links in the lock file.

Then you can fetch the PR locally, relock the lock file and push the changes back to the repo.

Snyk is building interactions with private repositories, please contact [support@snyk.io](mailto:support@snyk.io,)  if you are requesting this feature. 

