# Custom Versioning Schema for Custom Base Images

The Custom Versioning Schema (CVS) is a way for Snyk to understand your company’s container image tag versioning scheme, enabling Snyk to give more accurate base image upgrade recommendations.

## Prerequisites for Snyk CVS

CVS is part of the Snyk Container **Custom Base Images** feature. For more information, see [Custom Base Image recommendations](../../../scan-applications/snyk-container/use-snyk-container/use-custom-base-image-recommendations/).

{% hint style="info" %}
The Custom Versioning Schema is available only through the [Snyk REST API.](https://apidocs.snyk.io/a?#get-/custom\_base\_images)
{% endhint %}

## What is a versioning schema?

A versioning schema is a system for identifying and organizing different versions of a Project. It is used to track changes and updates to the Project over time and to help users identify which version they are using.

A versioning schema typically consists of a series of numbers or labels that are incremented to reflect the progression of versions.

For example, a versioning schema might use a series of numbers, such as "1.0", "1.1", "2.0", and so on, to indicate major and minor releases of a product.

A consistent and well-defined versioning schema helps users and tools understand and track the development of a Project.

## When to use the Custom Versioning Schema (CVS)

If your container image's tags follow a versioning scheme other than [Semantic Versioning](https://semver.org/) (SemVer), it is highly recommended that you select the Custom Versioning Schema for your image repositories.

## Custom Versioning Schema expression guide

CVS is essentially a regular expression that groups the different parts of an image’s tag into comparable sections.

### Ordering in a CVS

As an example, consider the following image tags:

```
snyk/example:1.2_V3
snyk/example:1.2_V4
snyk/example:1.3_V1
```

Because this repository’s image tags are not following the semantic versioning standard, it is necessary to describe the tags using a custom versioning schema.

The following defines the `snyk/example` tag schema:

* First is a **number** whose value has the highest significance (MAJOR part)
* Followed by a **period**
* Followed by another **number** whose significance is less than the first number (MINOR part)
* Followed by an **underscore** and the letter "**V**"
* And lastly, a **number** whose value is the least significant

For Snyk to understand the different parts and their role, it is necessary to define a schema. In this regular expression, named groups represent the significant variables.

```regex
(?<SIGNIFICANT>\d+)\.(?<LESS_SIGNIFICANT>\d+)_V(?<LEAST_SIGNIFICANT>\d+)
```

This schema is a translated version of the bullet point explanation. A small change to the names of the groups is needed:

<pre class="language-regex"><code class="lang-regex">(?&#x3C;<a data-footnote-ref href="#user-content-fn-1">C0</a>>\d+)\.(?&#x3C;<a data-footnote-ref href="#user-content-fn-2">C1</a>>\d+)_V(?&#x3C;<a data-footnote-ref href="#user-content-fn-3">C2</a>>\d+)
</code></pre>

Instead of naming a group "**SIGNIFICANT**", the name is changed to the letter "C" followed by a number. "**C**" stands for "**compare**", and the number represents the significance of that group, where 0 is the most significant.

Snyk will then do the following:

* Parse all of the tags in a repository using this expression
* Compare the values in order of significance
* Generate a set of images ordered by their tags.

Snyk can then use this ordered set to give better recommendations.

### Filtering with CVS

The following expands on the example repository and includes a small modification to add a slim version of each image.

```
snyk/example:1.2_V3-full
snyk/example:1.2_V3-slim
snyk/example:1.2_V4-full
snyk/example:1.2_V4-slim
snyk/example:1.3_V1-full
snyk/example:1.3_V1-slim
```

The following ensures that you will receive only recommendations for the flavor you are currently using.

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)_V(?&#x3C;C2>\d+)<a data-footnote-ref href="#user-content-fn-4">\-(?&#x3C;M0>.*)</a>
</code></pre>

This includes a new group **M0**. "**M**" stands for "**match**." Snyk will use this group to filter out possible image recommendations where the match group's value is not equal.

{% hint style="info" %}
The number following the "M" differentiates the multiple categories that your tag might have. In this case, there is only one category of things to match against, so there is only one match group.
{% endhint %}

For example, recommendations for `snyk/example:1.2_V3-slim` will not include images whose **M0** group does not equal "slim".

## Examples of expressions

### Standard format expression

The following demonstrates how to develop the expression for this registry.

```
snyk/example:1.2.5_deb9_2023061209
snyk/example:1.2.5_deb9_2023090208
snyk/example:1.2.5_deb10_2023110208
snyk/example:1.2.6_deb10_2023110508
snyk/example:1.4.1_deb10_2023110808
snyk/example:1.4.1_deb10_2023110208-slim
```

Start with defining the **simple semver part**:

```regex
(?<C0>\d+)\.(?<C1>\d+)\.(?<C2>\d+)
```

This groups the major, minor, and patch parts of the tag.

Next there is a section that indicates the underlying OS distribution. Here, there are two options:

* If the OS that the base image is using is not a concern, and only the version's magnitude is a concern, "ignore" this section by not grouping it in the regex: `_.*_`.
* To avoid a mismatch between distributions, add a **MATCH** group: `(?<M0>deb\d+)`

Now the expression looks like this:

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)\.(?&#x3C;C2>\d+)<a data-footnote-ref href="#user-content-fn-5">_(?&#x3C;M0>deb\d+)_</a>
</code></pre>

Next is **the date part**. Sometimes dates are there only to provide more information and need not be taken into consideration when comparing versions. In this case, skip over it.

If the date is important, decide how significant each date part is relative to the “semver” part. Is the **year** part more significant than a **major** version?

To keep the significance order, use the regex:&#x20;

```regex
(?<C3>\d{4})(?<C4>\d{2})(?<C5>\d{2})(?<C6>\d{2})
```

Since the date is ordered in such a way that the number produced by concatenating the year, month, day, and hour can be compared to another concatenated date correctly, the long regex above can be replaced with a simpler one:&#x20;

```regex
(?<C3>\d{10})
```

Now the regex looks like this:

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)\.(?&#x3C;C2>\d+)_(?&#x3C;M0>deb\d+)<a data-footnote-ref href="#user-content-fn-6">_(?&#x3C;C3>\d{10})</a>
</code></pre>

The **optional flavor** is last. Add another **MATCH** group here and make it optional:&#x20;

```regex
(?:\_(?<M1>.*))?
```

This avoids getting `slim` recommended if it is not in use and will only get `slim` recommended if it is being used.

The complete custom versioning schema expression looks like this:&#x20;

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)\.(?&#x3C;C2>\d+)_(?&#x3C;M0>deb\d+)_(?&#x3C;C3>\d{10})<a data-footnote-ref href="#user-content-fn-7">(?:\-(?&#x3C;M1>.*))?</a>
</code></pre>

### Inconsistent tag format (optional groups)

In cases where a repository does not have a consistent tagging format, non-capture groups can be used.

```
snyk/example:1.1
snyk/example:1.1.2
snyk/example:1.2
snyk/example:1.2.4
snyk/example:1.2.8
snyk/example:1.3
snyk/example:1.3.5
```

The above repository contains an inconsistent number of capture groups. Use the following expression to handle this case:

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)<a data-footnote-ref href="#user-content-fn-8">(?:\.(?&#x3C;C2>\d+))?</a>
</code></pre>

The interesting part here is `(?:\.(?<C2>\d+))?`. The expression optionally includes a **compare** group.

When there is an inconsistent number of **compare** groups, Snyk will filter out tags that do not contain enough information to compare them accurately. That is, to get a recommendation for `snyk/example:1.2.4`, Snyk does not consider `snyk/example:1.2` to be a newer version because it is not possible to know whether `1.2` is the same as `1.2.0` or whether it is a rolling tag that points to `1.2.8`. However `1.3`, and `1.3.5` are both definitively higher than `1.2.4` and will be taken into account as a possible recommendation.

Since `1.3` and `1.3.5` have the same issue as `1.2` and `1.2.4`, the recommendation will be either `1.3` or `1.3.5`. At this point, the specific version that will be recommended depends on undefined internal factors. However, we are looking to improve this logic.

## Technical details of the custom versioning schema

### Regular expression syntax flavor

Custom Versioning Schema uses a **subset** of the ECMAScript regex.

There is an excellent [guide on ECMAScript regex syntax by MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular\_Expressions).

Some noteworthy features that are not supported are backreferences and lookahead assertions. Snyk uses the RE2 library under the hood. For a full list of unsupported features, see [Syntax on re2 GitHub](https://github.com/google/re2/wiki/Syntax).

Note that the regular expression string is parsed as an ECMAScript regex and then internally converted to RE2 syntax. For example, use the `(?<name>re)` syntax for grouping. `(?P<name>re)` will not parse correctly.

When looking at the [list of supported features](https://github.com/google/re2/wiki/Syntax), pay attention only to the **feature,** not the **syntax**.

### Size limitations

The maximum length of the expression is 1000 characters, with up to 100 **COMPARE** groups and 100 **MATCH** groups.

### Groups

All named groups **must** start with either the letter **C** or the letter **M,** followed by an index.

There **must** exist at least one (1) **COMPARE** group.

If an unnamed group is required, it must be explicitly marked as **non-capturing**.

### Compare logic

If a compare group is an unsigned integer, compare its numeric value. Else, it will be treated as a string.

Comparison of non-numeric characters is done by comparing their sequences of UTF-16 code unit values. For more information, see [ECMAScript® 2023 Language Specification, 7.2.13 IsLessThan](https://tc39.es/ecma262/#sec-abstract-relational-comparison).

### Matcher logic

Match groups are case-sensitive.

If a match group exists on one tag and not the other, due to the use of optional groups, the tags are considered **not** matching.

## Errors

### `Use of unsupported or invalid regex syntax`

#### How can this occur?

This error can happen if an expression that does not conform to ECMAScript syntax is passed.

An example is using Python’s named capture group syntax `(?P<C0>.*)`.

#### How to fix it

See [Regular expression syntax flavor](custom-versioning-schema-for-custom-base-images.md#regular-expression-syntax-flavor) for information on the regex syntax.

In the example case, change the capture group syntax from `(?P<C0>.*)` to `(?<C0>.*)`.

### `Use of an unsupported regex feature`

#### How can this occur?

This error can happen if an expression is passed that contains unsupported regex features such as lookahead assertions and backreferences.

An example is `(?<thing>.+)_\\k<thing>/`.

#### How to fix it

There is no simple fix. The expression must be redesigned to work without such features.

See [Regular expression syntax flavor](custom-versioning-schema-for-custom-base-images.md#regular-expression-syntax-flavor) for more information about the supported and unsupported features.

### `Group name format is incorrect`

#### How can this occur?

Custom Versioning Schema uses named groups that follow a specific naming format.

This error can happen if a named group exists that does not follow the format of either a `C` or an `M`, followed by a positive number, such as `C2` or `M51`

#### How to fix it

Change the group name to fit the format. If you need only a group, you can use non-capture groups.

### `Operator index is too big`

#### How can this occur?

This can error can happen if the number in your named group is larger than **100**, such as `(?<C101>.*)`.

#### How to fix it

Use a number between **0** and **100**.

Remember, each character does not require its own separate group. If it makes logical sense, bunch the characters together to end up with fewer groups.

If you still require more than 100 groups, Custom Versioning Schema might not be a great fit.

### `Group is missing a name`

#### How can this occur?

This error can happen if a group other than a named group or non-capture group is used.

Example `(debian)|(ubuntu)`.

#### How to fix it

Often the solution is to explicitly define a non-capture group.

The example `(debian)|(ubuntu)` becomes `(?:debian)|(?:ubuntu)`.

### `No groups were found`

#### How can this occur?

Custom Versioning Schema uses named groups to categorize the parts of the tag into comparable sections.

This error can happen if no named groups were found, in which case Snyk cannot compare different image tags to each other.

#### How to fix it

Look at the [Custom Versioning schema expression guide](custom-versioning-schema-for-custom-base-images.md#custom-versioning-schema-expression-guide) on this page to better understand how Snyk compares and filters out tags.

### `Expression length is not supported`

#### How can this occur?

As with every user input, a maximum length must be set. In this case, the limit is 1000 characters.

If your expression string is longer than 1000 characters, Snyk will not be able to parse it.

#### How to fix it

Try shortening the string by using [character classes](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular\_Expressions/Character\_Classes) and [quantifiers](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular\_Expressions/Quantifiers) instead of specific characters.

If you still require more than 1000 characters to describe your tags, Custom Versioning Schema might not be a great fit. For information about options, contact [Snyk support](https://support.snyk.io/hc/en-us/requests/new).

[^1]: MOST\_SIGNIFICANT

[^2]: LESS\_SIGNIFICANT

[^3]: LEAST\_SIGNIFICANT

[^4]: This is what we added

[^5]: What we added

[^6]: What we added

[^7]: What we added

[^8]: compare group within an optional non capture group
