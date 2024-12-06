# Custom versioning schema for custom base images

The Custom Versioning Schema (CVS) is a way for Snyk to understand your company’s container image tag versioning schema. It enables Snyk to give more accurate base image upgrade recommendations.

## Prerequisites for Snyk CVS

CVS is part of the Snyk Container Custom Base Images feature. For more information, see [Use Custom base image recommendations](./).

## When to use CVS

If your container image's tags follow a versioning schema other than [Semantic Versioning](https://semver.org/) (SemVer), it is highly recommended that you select the Custom Versioning Schema for your image repositories.

## CVS expression guide

CVS is essentially a regular expression that groups the different parts of an image’s tag into comparable sections.

### Order in a CVS

As an example, consider the following image tags:

```
snyk/example:1.2_V3
snyk/example:1.2_V4
snyk/example:1.3_V1
```

Because this repository’s image tags do not follow the semantic versioning standard, it is necessary to describe the tags using a custom versioning schema.

The  `snyk/example` tag schema is defined by the following elements, in this order:

1. A number whose value has the highest significance (MAJOR part)
2. A period
3. Another number, whose significance is less than the first number number (MINOR part)
4. An underscore followed by the letter "V" (version)
5. A number whose value is the least significant.

For Snyk to understand the different parts and their role, it is necessary to define a schema. In this regular expression, named groups represent the significant variables.

The schema below is a translated version of the above example and its elements.

```regex
(?<SIGNIFICANT>\d+)\.(?<LESS_SIGNIFICANT>\d+)_V(?<LEAST_SIGNIFICANT>\d+)
```

Instead of naming a group "SIGNIFICANT", the name is changed to the letter "C" followed by a number. "C" stands for "compare", and the number represents the significance of that group, where 0 is the most significant.

<pre class="language-regex"><code class="lang-regex">(?&#x3C;<a data-footnote-ref href="#user-content-fn-1">C0</a>>\d+)\.(?&#x3C;<a data-footnote-ref href="#user-content-fn-2">C1</a>>\d+)_V(?&#x3C;<a data-footnote-ref href="#user-content-fn-3">C2</a>>\d+)
</code></pre>

Snyk then:

* Parses all of the tags in a repository using this expression.
* Compares the values in order of significance.
* Generates a set of images ordered by their tags.

Snyk can then use this ordered set to give better recommendations.

### Filter with CVS

The example below expands on the repository in the previous example and includes a small modification to add a slim version of each image.

```
snyk/example:1.2_V3-full
snyk/example:1.2_V3-slim
snyk/example:1.2_V4-full
snyk/example:1.2_V4-slim
snyk/example:1.3_V1-full
snyk/example:1.3_V1-slim
```

The schema below ensures that you receive only recommendations for the flavor you are currently using.

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)_V(?&#x3C;C2>\d+)<a data-footnote-ref href="#user-content-fn-4">\-(?&#x3C;M0>.*)</a>
</code></pre>

This includes a new group M0, where "M" stands for "match." Snyk then uses this group to filter out possible image recommendations where the MATCH group's value is not equal.

{% hint style="info" %}
The number following "M" differentiates the multiple categories that your tag can have. In this case, there is only one category of things to match against, so there is only one MATCH group.
{% endhint %}

For example, recommendations for `snyk/example:1.2_V3-slim` do not include images whose M0 group does not equal "slim".

## Examples of expressions

### Standard format expression

The example below shows how to develop the expression for this registry.

```
snyk/example:1.2.5_deb9_2023061209
snyk/example:1.2.5_deb9_2023090208
snyk/example:1.2.5_deb10_2023110208
snyk/example:1.2.6_deb10_2023110508
snyk/example:1.4.1_deb10_2023110808
snyk/example:1.4.1_deb10_2023110208-slim
```

Start with defining the simple SemVer element:

```regex
(?<C0>\d+)\.(?<C1>\d+)\.(?<C2>\d+)
```

This groups the major, minor, and patch parts of the tag.

Next, there is a section that indicates the underlying OS distribution. Here, there are two options:

* If the OS that the base image is using is not a concern, and only the version's magnitude is a concern, "ignore" this section by not grouping it in the regex: `_.*_`.
* To avoid a mismatch between distributions, add a MATCH group: `(?<M0>deb\d+)`

Now, the expression looks like this:

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)\.(?&#x3C;C2>\d+)<a data-footnote-ref href="#user-content-fn-5">_(?&#x3C;M0>deb\d+)_</a>
</code></pre>

Next is the date element. Sometimes dates are there only to provide more information and need not be taken into consideration when comparing versions. In this case, skip over it.

If the date element is important, decide how significant each date element is relative to the SemVer. For example, is the year more significant than a major version?

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

The optional flavor is last. Add another MATCH group here and make it optional:&#x20;

```regex
(?:\_(?<M1>.*))?
```

This avoids getting `slim` recommended if it is not in use and only gets `slim` recommended if it is being used.

The complete custom versioning schema expression looks like this:&#x20;

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)\.(?&#x3C;C2>\d+)_(?&#x3C;M0>deb\d+)_(?&#x3C;C3>\d{10})<a data-footnote-ref href="#user-content-fn-7">(?:\-(?&#x3C;M1>.*))?</a>
</code></pre>

### Inconsistent tag format (optional groups)

In cases where a repository does not have a consistent tagging format, you can use non-capture groups.

```
snyk/example:1.1
snyk/example:1.1.2
snyk/example:1.2
snyk/example:1.2.4
snyk/example:1.2.8
snyk/example:1.3
snyk/example:1.3.5
```

The above repository contains an inconsistent number of capture groups. To handle this, use the following expression:

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)<a data-footnote-ref href="#user-content-fn-8">(?:\.(?&#x3C;C2>\d+))?</a>
</code></pre>

In part `(?:\.(?<C2>\d+))?` , the expression optionally includes a COMPARE group.

When there is an inconsistent number of COMPARE groups, Snyk filters out tags that do not contain enough information to compare them accurately. That is, to get a recommendation for `snyk/example:1.2.4`, Snyk does not consider `snyk/example:1.2` to be a newer version because it is not possible to know whether `1.2` is the same as `1.2.0` or whether it is a rolling tag that points to `1.2.8`. However, `1.3`, and `1.3.5` are both higher than `1.2.4` and are taken into account as a possible recommendation.

Since `1.3` and `1.3.5` have the same issue as `1.2` and `1.2.4`, the recommendation is either `1.3` or `1.3.5`. At this point, the specific version that is recommended depends on undefined internal factors. However, Snyk aims to improve this logic.

## Technical rules of CVS

### Regular expression syntax flavor

* CVS uses a subset of the ECMAScript regex. See the [guide on ECMAScript regex syntax by MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions).
* Backreferences and lookahead assertions are not supported. Internally, Snyk uses the RE2 library. For a full list of unsupported features, see [Syntax on re2 GitHub](https://github.com/google/re2/wiki/Syntax).
* Note that the regular expression string is parsed as an ECMAScript regex and then internally converted to RE2 syntax. For example, use the `(?<name>re)` syntax for grouping. `(?P<name>re)` will not parse correctly.
* In the list of [supported features](https://github.com/google/re2/wiki/Syntax), take into consideration only the featur&#x65;**,** not the syntax.

### Size limitations

The maximum length of the expression is 1,000 characters, with up to 100 COMPARE groups and 100 MATCH groups.

### Groups

* All named groups must start with either the letter "C" or the letter "M"**,** followed by an index.
* There must be at least one COMPARE group.
* If an unnamed group is required, it must be explicitly marked as non-capturing.

### Compare logic

* If a compare group is an unsigned integer, compare its numeric value. Else, it will be treated as a string.
* Comparison of non-numeric characters is done by comparing their sequences of UTF-16 code unit values. For more information, see [ECMAScript® 2023 Language Specification, 7.2.13 IsLessThan](https://tc39.es/ecma262/#sec-abstract-relational-comparison).

### Matcher logic

* MATCH groups are case-sensitive.
* If a MATCH group exists on one tag and not the other, due to the use of optional groups, the tags are considered not matching.

## Errors in CVS

`Use of unsupported or invalid regex syntax`

#### Why does it occur?

This error can happen if an expression that does not conform to ECMAScript syntax is passed.

An example is using Python’s named capture group syntax `(?P<C0>.*)`.

#### How to fix it

In the example, change the capture group syntax from `(?P<C0>.*)` to `(?<C0>.*)`.

For information on the regex syntax, see [Regular expression syntax flavor](custom-versioning-schema-for-custom-base-images.md#regular-expression-syntax-flavor).



`Use of an unsupported regex feature`

#### Why does it occur?

This error can happen if an expression is passed that contains unsupported regex features such as lookahead assertions and backreferences.

For example `(?<thing>.+)_\\k<thing>/`.

#### How to fix it

There is no simple fix. The expression must be redesigned to work without such features.

See [supported features](https://github.com/google/re2/wiki/Syntax).



`Group name format is incorrect`

#### Why does it occur?

Custom Versioning Schema uses named groups that follow a specific naming format.

This error can happen if a named group exists that does not follow the format of either a `C` or an `M`, followed by a positive number, such as `C2` or `M51.`

#### How to fix it

Change the group name to fit the format. If you need only a group, you can use non-capture groups.



`Operator index is too big`

#### Why does it occur?

This error can happen if the number in your named group is larger than 100, such as `(?<C101>.*)`.

#### How to fix it

Use a number between 0 and 100.

Take into consideration that each character does not require its own separate group. If it makes logical sense, bunch the characters together to end up with fewer groups.

If you still require more than 100 groups, CVS might not be a great fit.



`Group is missing a name`

#### Why does it occur?

This error can happen if a group other than a named group or non-capture group is used.

For example `(debian)|(ubuntu)`.

#### How to fix it

Often the solution is to explicitly define a non-capture group.

The example `(debian)|(ubuntu)` becomes `(?:debian)|(?:ubuntu)`.



`No groups were found`

#### Why does it occur?

CVS uses named groups to categorize the parts of the tag into comparable sections.

This error can happen if no named groups are found, in which case Snyk cannot compare different image tags to each other.

#### How to fix it

See the [CVS expression guide](custom-versioning-schema-for-custom-base-images.md#cvs-expression-guide) to better understand how Snyk compares and filters out tags.



`Expression length is not supported`

#### Why does it occur?

As with every user input, a maximum length must be set. In this case, the limit is 1,000 characters.

If your expression string is longer than 1,000 characters, Snyk is not able to parse it.

#### How to fix it

Shorten the string by using [character classes](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions/Character_Classes) and [quantifiers](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions/Quantifiers) instead of specific characters.

If you still require more than 1,000 characters to describe your tags, CVS might not be a great fit.&#x20;

For information on other options, contact [Snyk support](https://support.snyk.io).

[^1]: MOST\_SIGNIFICANT

[^2]: LESS\_SIGNIFICANT

[^3]: LEAST\_SIGNIFICANT

[^4]: This is what we added

[^5]: What we added

[^6]: What we added

[^7]: What we added

[^8]: compare group within an optional non capture group
