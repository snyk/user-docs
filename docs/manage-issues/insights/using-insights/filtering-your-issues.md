# Filtering your issues

Insights operates at a Group level and offers you a holistic view of all the issues within that Group, but those issues are also tied to specific Organizations. By using the top-level filter, you can choose which Organizations are relevant to you and see only the issues in those Organizations.

<figure><img src="https://lh4.googleusercontent.com/paOzhXTtTM0VSwRKzKplmg5SuWbbbK68uyiQPP6kQ2xilPjMfchcpactt13_3_Ok1ARsxX2YLdswXK0mWnv6YLZxSHXaAbEujEqiR-ewasTJvTO9N28oIATX6eovT6yTdGOHt2D3W5AM2IhQLFyIupc" alt="Issues prioritized by risk factor"><figcaption><p>Issues prioritized by risk factor</p></figcaption></figure>

## Funnel View

The visual representation of your application’s issues and risk factors is called the funnel view. It provides you with the possibility of filtering the list of issues by the ones that have specific risk factors present or a combination of them. Os condition, Deployed, and Public facing risk factors are clickable filters.

## Table view filters

By using the filters above the table view, you can filter your issues by the following criteria or any combination of them.

* **Issue status** - view only issues that are open, resolved, or ignored
* **Risk factors** - view only issues with specific risk factors
* **Asset name** - view only issues related to specific assets
* **Source code** - view only issues related to specific source code
* **Issue severity** - view only critical, high, medium, or low severity issues
* **Snyk Product** - view only issues scanned by specific Snyk products

You can also dd a variety of filters that you consider relevant for any particular scenario.

Note that in an initial triage, you want to look at OS condition, Deployed, and Public facing, and choose BOTH critical and high severity. Snyk Open source finds critical vulnerabilities, while Snyk Code finds vulnerabilties up to high severity only. When you filter on the most concerning issues, filter for both critical and high severity.

When you want to split out Open Source and Snyk Code issues, use a product filter.

![Add filter](<../../../.gitbook/assets/Screenshot 2023-07-12 at 02.07.21.png>)\
