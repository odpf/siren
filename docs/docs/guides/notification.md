import Tabs from "@theme/Tabs";
import TabItem from "@theme/TabItem";
import CodeBlock from '@theme/CodeBlock';
import siteConfig from '/docusaurus.config.js';

# Notification

export const apiVersion = siteConfig.customFields.apiVersion
export const defaultHost = siteConfig.customFields.defaultHost

To understand more concepts of notification in Siren, you can visit this [page](../concepts/notification.md).

## Sending a message/notification

We could send a notification with `POST /notifications` API to a specific receiver by passing receiver labels information in the `receivers` field in the body. The payload format needs to follow receiver type contract. 

| Field Name 	| Type 	| Required? 	| Description 	|
|---	|---	|---	|---	|
| receivers 	| list of json 	| yes 	| Selector of receivers using receiver labels <br/>[<br/> {<br/>   "id": 3,    <br/>  },<br/> {<br/>   "type": "slack_channel",<br/>   "team": "de-infra"<br/> },<br/> {<br/>   "type": "email",<br/>   "team": "de-experience" <br/> }<br/>]<br/>
 This will fetch all receivers that have the labels. 	|
| data 	| json 	| yes 	| any data that we want to pass to the message. The data will populate the corresponding template or content variables. 	|
| template 	| string 	| no 	| template name that will be used to populate the message. default template will be used if this value is empty. errors might be thrown if there are errors when parsing template. 	|
| labels 	| json 	| no 	| If populated, labels would be used by subscription matchers to find the subscribers that listen to specific labels. e.g. <br/>{<br/>  "team": "de-infra",<br/>  "severity": "CRITICAL"<br/>}	|


### Example: Sending Notification to Slack

Assuming there is a slack receiver registered in Siren with ID `51`. Sending to that receiver would require us to have a `payload.data` that have the same format as the expected [slack](../receivers/slack.md#message-payload) payload format.

```yaml title=payload.yaml
payload:
  data:
    text: |-
      New Paid Time Off request from <example.com|Fred Enriquez>

      <https://example.com|View request>
  template: sample-slack-msg
```

<Tabs groupId="api">
  <TabItem value="cli" label="CLI" default>

```bash
$ siren notifications send --file payload.yaml
```

  </TabItem>
  <TabItem value="http" label="HTTP">
    <CodeBlock className="language-bash">
    {`$ curl --request POST
  --url `}{defaultHost}{`/`}{apiVersion}{`/notifications
  --header 'content-type: application/json'
  --data-raw '{
    "payload": {
        "data": {
          "text": "New Paid Time Off request from <example.com|Fred Enriquez>\n\n<https://example.com|View request>"
        }
    }
}'`}
    </CodeBlock>
  </TabItem>
</Tabs>

Above end the message to channel name `#siren-devs` with `payload.data` in [slack](#slack) payload format.


## Alerts Notification

For all incoming alerts via Siren hook API, notifications are also generated and published via subscriptions. Siren will match labels from the alerts with label matchers in subscriptions. The assigned receivers for all matched subscriptions will get the notifications. More details are explained [here](./alert_history.md). 

Siren has a default template for alerts notification for each receiver. Go to the [Receivers](../receivers/slack.md#default-alert-template) section to explore the default template defined by Siren.