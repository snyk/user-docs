# Filtering results

**INFO**

Filter rules can be used not only to **describe** resources, but also to **ignore** resources.

You can indeed use both inclusion and exclusion logics.

**Filter rules** allow you to build complex expression to include and exclude a set of resources in your workflow. Powered by expression language [JMESPath](https://jmespath.org) you could build a complex include and exclude expression.

Filter are applied on a normalized struct which contains the following fields:

* **Type**: Type of the resource, e.g. `aws_s3_bucket`
* **Id**: Id of the resource, e.g. `my-bucket-name`
* **Attr**: Contains every resource attributes (check [terraform attributes reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3\_bucket#attributes-reference) for a full list of supported attributes of a bucket)

**CAUTION**

If you want to filter on `Attr` you should enable deep mode otherwise you will not have access to resource's details.

#### Examples[​](https://docs.driftctl.com/0.22.0/usage/filtering/rules#examples) <a href="#examples" id="examples"></a>

```
# Will include only S3 bucket in the search
$ snyk iac describe --filter "Type=='aws_s3_bucket'"
# OR (beware of escape your shell special chars between double quotes)
$ snyk iac describe --filter $'Type==\'aws_s3_bucket\''
# Excludes only s3 bucket named 'my-bucket-name'
$ snyk iac describe --filter $'Type==\'aws_s3_bucket\' && Id!=\'my-bucket-name\''
# Ignore buckets that have tags terraform equal to 'false'
$ snyk iac describe --deep --filter $'!(Type==\'aws_s3_bucket\' && Attr.tags.terraform==\'false\')'
# Ignore buckets that don't have tag terraform
$ snyk iac describe --deep --filter $'!(Type==\'aws_s3_bucket\' && Attr.tags != null && !contains(keys(Attr.tags), \'terraform\'))'
# Ignore buckets with an ID prefix of 'terraform-'
$ snyk iac describe --filter $'!(Type==\'aws_s3_bucket\' && starts_with(Id, \'terraform-\'))'
# Ignore buckets with an ID suffix of '-test'
$ snyk iac describe --filter $'!(Type==\'aws_s3_bucket\' && ends_with(Id, \'-test\'))'
# Ignore GitHub archived repositories
$ snyk iac describe --to github+tf --deep --filter '!(Attr.Archived)'
```
