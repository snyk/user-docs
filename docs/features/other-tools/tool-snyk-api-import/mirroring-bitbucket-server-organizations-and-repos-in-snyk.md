# Mirroring Bitbucket Server organizations & repos in Snyk

## Mirroring Bitbucket Server organizations and repos in Snyk

In order to import the entirety of Bitbucket Server repos into Snyk you can use the available utils to make it possible in 4 commands. You will need to configure both Bitbucket Server token and Snyk token as environment variable to proceed. Please refer to individual documentation pages for more detailed info, however the general steps are:

1. `export BITBUCKET_SERVER_TOKEN=***` and `export SNYK_TOKEN=***`
2. Generate organization data e.g. `snyk-api-import orgs:data --source=bitbucket-server --groupId=<snyk_group_id>` Full instructions
3. Create organizations in Snyk `snyk-api-import orgs:create --file=orgs.json` Full instructions will create a `snyk-created-orgs.json` file with Snyk organization ids and integration ids that are needed for import.
4. Generate import data `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=bitbucket-server --integrationType=bitbucket-server` Full instructions
5. Run import `DEBUG=*snyk* snyk-api-import import`Full instructions

### Re-importing new repos & orgs only while Mirroring

Once initial import is complete you may want to periodically check for new repos and make sure they are added into Snyk. To do this a similar flow to what is described above with a few small changes can be used:

1. `export BITBUCKET_SERVER_TOKEN=***` and `export SNYK_TOKEN=***`
2. Generate organization data in Snyk and skip any that do not have any repos via `--skipEmptyOrg` `snyk-api-import orgs:data --source=bitbucket-server --groupId=<snyk_group_id> --skipEmptyOrg` Full instructions
3. Create organizations in Snyk and this time skip any that have been created already with `--noDuplicateNames` parameter `snyk-api-import orgs:create --file=orgs.json --noDuplicateNames` Full instructions will create a `snyk-created-orgs.json` file with Snyk organization ids and integration ids that are needed for import.
4. Generate import data `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=bitbucket-server --integrationType=bitbucket-server` Full instructions
5. Optional. Generate the previously imported log to skip all previously imported repos a Group (see full documentation): `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>`
6. Run import `DEBUG=*snyk* snyk-api-import import`Full instructions
