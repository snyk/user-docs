# Satisfying version not found for cousin dependencies

When importing a .Net project, Snyk is creating the dependency tree according to the.\*proj file. According to the NuGet logic, it is not allowed to use the same package with different versions in the same project. For example, let's assume that Package A and Package B are direct dependencies of the same project and both use Package C, meaning Package C is a transitive dependency. When resolving the version of Package C, it must be resolved into 1 version and to satisfy the requirements of Package A and Package B. Failing to do that will result in the error "Satisfying version not found for cousin dependencies error".

To troubleshoot the issue, please try to build the project. If it was built successfully, please reach out to [support@snyk.io](mailto:support@snyk.io) and we will look into it. If it wasn't, please make sure that all of the dependencies you are using in your project are not conflicting. 

