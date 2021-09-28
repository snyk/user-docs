# Send Message

Lastly, create an **Action** and name it **Send Channel Message in Microsoft Teams**:

![](../../../../.gitbook/assets/zappier-teams-main.png)

Select **Send Channel Message** for the **Microsoft Teams** app:

![](../../../../.gitbook/assets/zappier-teams-channel.png)

Choose your Microsoft Teams account:

![](../../../../.gitbook/assets/zappier-teams-account.png)

Choose your Channel and select **Markdown** for the **Message Text Format**:

![](../../../../.gitbook/assets/zappier-teams-setup.png)

Copy and paste the following snippet into the **Message Text** field:

```markup
# New Issue
We found vulnerabilities that affect [{{project__name}}]({{project__browseUrl}}) project in the [{{org__name}}]({{org__url}}).

## Details:
<div class="issue">
  <p>{{newIssues}}</p>
</div>
```

{% hint style="info" %}
You will need to update the fields accordingly to correspond with those mapped from your payload.
{% endhint %}

Test your action:

![](../../../../.gitbook/assets/zappier-teams-test.png)

