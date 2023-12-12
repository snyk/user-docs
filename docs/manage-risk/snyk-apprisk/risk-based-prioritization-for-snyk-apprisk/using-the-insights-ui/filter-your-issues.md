# Filter your issues

Snyk AppRisk operates at the Group level and provides a holistic view of all the issues within that Group. Those issues are also tied to specific Organizations. Use the top-level filter to choose which Organizations are relevant to you and see only the issues in those Organizations.

<figure><img src="https://lh4.googleusercontent.com/paOzhXTtTM0VSwRKzKplmg5SuWbbbK68uyiQPP6kQ2xilPjMfchcpactt13_3_Ok1ARsxX2YLdswXK0mWnv6YLZxSHXaAbEujEqiR-ewasTJvTO9N28oIATX6eovT6yTdGOHt2D3W5AM2IhQLFyIupc" alt="Issues prioritized by risk factor in Organizations"><figcaption><p>Issues prioritized by risk factor in Organizations</p></figcaption></figure>

## Funnel View

The visual representation of the issues and risk factors for your application is called the funnel view. It makes it possible to filter the list of issues by the ones with specific risk factors or a combination of them. The risk factors Os condition, Deployed, and Public facing are clickable filters.

## Table view filters

By using the filters above the table view, you can filter your issues by the following criteria or any combination of them.

* **Issue status** - view only issues that are open, resolved, or ignored
* **Risk factors** - view only issues with specific risk factors
* **Asset name** - view only issues related to specific assets
* **Asset class** - view only issues related to a specific asset class
* **Source code** - view only issues related to specific source code
* **Issue severity** - view only critical, high, medium, or low-severity issues
* **Snyk Product** - view only issues scanned by specific Snyk products
* **Asset class** - view only issues related to specific asset class

You can also add a variety of filters that you consider relevant for any particular scenario.

Note that in an initial triage, you want to look at OS condition, Deployed, and Public facing, and choose both critical and high severity. Snyk Open Source finds critical vulnerabilities, while Snyk Code finds vulnerabilities up to high severity only. When you filter on the issues of greatest concern, filter for both critical and high severity.

When you filter based on the Asset class and an issue is found in two different repositories with two different classes assigned, the class with the highest priority is displayed in Snyk AppRisk.

When you want to split out Open Source and Snyk Code issues, use a product filter.

![Add filter dropdown](<../../../../.gitbook/assets/Screenshot 2023-07-12 at 02.07.21.png>)\
