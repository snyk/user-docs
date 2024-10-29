# Mirroring Bitbucket Cloud organizations and repos in Snyk

You can use four commands in the available utils to import the entirety of Bitbucket Cloud repos into Snyk. You must configure the Bitbucket Cloud username and password and Snyk token as environment variables to proceed.

## General steps to import Bitbucket Cloud repos

Refer to individual documentation pages for detailed information. The general steps are as follows:

1. `export BITBUCKET_CLOUD_USERNAME=***`, `export BITBUCKET_CLOUD_PASSWORD=***` and `export SNYK_TOKEN=***`
2. Generate organization data, for example, `snyk-api-import orgs:data --source=bitbucket-cloud --groupId=<snyk_group_id>` Full instructions: [Creating organizations in Snyk](creating-organizations-in-snyk.md)
3. Create organizations in Snyk: `snyk-api-import orgs:create --file=orgs.json` By following the full instructions on [Creating organizations in Snyk](creating-organizations-in-snyk.md) you will create a `snyk-created-orgs.json` file with Snyk organization ids and integration ids that are needed for import.
4. Generate import data: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=bitbucket-cloud --integrationType=bitbucket-cloud` Full instructions: [Creating import targets data for import](creating-import-targets-data-for-import-command.md)
5. Run import: `DEBUG=*snyk* snyk-api-import import` Full instructions: [Kicking off an import](kicking-off-an-import.md)\`\`

## Re-importing new repos and Orgs only while mirroring

Once initial import is complete you can periodically check for new repos and make sure they are added into Snyk by following these steps, which are similar to the preceding steps to import repos.

1. `export BITBUCKET_CLOUD_USERNAME=***`, `export BITBUCKET_CLOUD_PASSWORD=***` and `export SNYK_TOKEN=***`
2. Generate organization data in Snyk and skip any that do not have any repos by using `--skipEmptyOrg` `snyk-api-import orgs:data --source=bitbucket-cloud --groupId=<snyk_group_id> --skipEmptyOrg` Full instructions: [Creating organizations in Snyk](creating-organizations-in-snyk.md)
3. Create organizations in Snyk and this time skip any that have been created already with `--noDuplicateNames` parameter `snyk-api-import orgs:create --file=orgs.json --noDuplicateNames` By following the full instructions on [Creating organizations in Snyk](creating-organizations-in-snyk.md) you will create a `snyk-created-orgs.json` file with Snyk organization ids and integration ids that are needed for import.
4. Generate import data: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=bitbucket-cloud --integrationType=bitbucket-cloud` Full instructions: [Creating import targets data for import](creating-import-targets-data-for-import-command.md)
5. Optional - Generate the previously imported log to skip all previously imported repos in a Group: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>` Full instructions: [Kicking off an import](kicking-off-an-import.md)\`\`
6. Run import `DEBUG=*snyk* snyk-api-import import` Full instructions: [Kicking off an import](kicking-off-an-import.md)
