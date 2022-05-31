# Mirroring GitHub.com and GitHub Enterprise organizations and repos in Snyk

You can use four commands in the available utils to import the entirety of GitHub and GitHub Enterprise repos into Snyk. You must configure both the GitHub token and Snyk token as environment variables to proceed.

## General steps for importing GitHub and GitHub Enterprise Orgs and repos

Refer to individual documentation pages for more detailed info, however the general steps are:

1. `export GITHUB_TOKEN=***` and `export SNYK_TOKEN=***`
2. Generate organization data, for example, `snyk-api-import orgs:data --source=github --groupId=<snyk_group_id>` Full instructions: [Creating organizations in Snyk](creating-orgs-in-snyk.md)
3. Create organizations in Snyk: `snyk-api-import orgs:create --file=orgs.json` If you follow the full instructions on [Creating organizations in Snyk](creating-orgs-in-snyk.md), you will create a `snyk-created-orgs.json` file with Snyk organization ids and integration ids that are needed for import.
4. Generate import data: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=github --integrationType=github` Full instructions: [Creating imports targets data for import](creating-import-targets-data-for-import.md)
5. Run import: `DEBUG=*snyk* snyk-api-import import`Full instructions: [Kicking off an import](kicking-off-an-import.md)

## Re-importing new repos and Orgs only while mirroring

Once initial import is complete you may want to periodically check for new repos and make sure they are added into Snyk. To do this a similar flow to what is described above with a few small changes can be used:

1. `export GITHUB_TOKEN=***` and `export SNYK_TOKEN=***`
2. Generate organization data in Snyk and skip any that do not have any repos via `--skipEmptyOrg` `snyk-api-import orgs:data --source=github --groupId=<snyk_group_id> --skipEmptyOrg` Full instructions [https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/orgs.md](https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/orgs.md)
3. Create organizations in Snyk and this time skip any that have been created already with `--noDuplicateNames` parameter `snyk-api-import orgs:create --file=orgs.json --noDuplicateNames` Full instructions will create a `snyk-created-orgs.json` file with Snyk organization ids and integration ids that are needed for import. [https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/orgs.md](https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/orgs.md)
4. Generate import data `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=github --integrationType=github` Full instructions [https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/import-data.md](https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/import-data.md)
5. Optional. Generate the previously imported log to skip all previously imported repos a Group (see full documentation): `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>` [`https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/import.md#to-skip-all-previously-imported-targets`](https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/import.md#to-skip-all-previously-imported-targets)``
6. Run import `DEBUG=*snyk* snyk-api-import import`Full instructions [https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/import.md](https://github.com/snyk-tech-services/snyk-api-import/blob/master/docs/import.md)
