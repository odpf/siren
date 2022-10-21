(self.webpackChunksiren=self.webpackChunksiren||[]).push([[616],{1410:function(e,t,r){const n=r(7694),i=r(3618);e.exports={title:"Siren",tagline:"Universal data observability toolkit",url:"https://odpf.github.io",baseUrl:"/siren/",onBrokenLinks:"throw",onBrokenMarkdownLinks:"warn",favicon:"img/favicon.ico",organizationName:"odpf",projectName:"siren",customFields:{apiVersion:"v1beta1",defaultHost:"http://localhost:8080"},presets:[["@docusaurus/preset-classic",{docs:{sidebarPath:6679,editUrl:"https://github.com/odpf/siren/edit/master/docs/",sidebarCollapsed:!0,breadcrumbs:!1},blog:!1,theme:{customCss:[5308,2295]},gtag:{trackingID:"G-EPXDLH6V72"}}]],themeConfig:{colorMode:{defaultMode:"light",respectPrefersColorScheme:!0},navbar:{title:"Siren",logo:{src:"img/logo.svg"},hideOnScroll:!0,items:[{type:"doc",docId:"introduction",position:"right",label:"Docs"},{to:"docs/support",label:"Support",position:"right"},{href:"https://bit.ly/2RzPbtn",position:"right",className:"header-slack-link"},{href:"https://github.com/odpf/siren",className:"navbar-item-github",position:"right"}]},footer:{style:"light",links:[]},prism:{theme:n,darkTheme:i},announcementBar:{id:"star-repo",content:'\u2b50\ufe0f If you like Siren, give it a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/odpf/siren">GitHub</a>! \u2b50',backgroundColor:"#222",textColor:"#eee",isCloseable:!0}}}},5162:function(e,t,r){"use strict";r.d(t,{Z:function(){return o}});var n=r(7294),i=r(4334),a="tabItem_Ymn6";function o(e){let{children:t,hidden:r,className:o}=e;return n.createElement("div",{role:"tabpanel",className:(0,i.Z)(a,o),hidden:r},t)}},5488:function(e,t,r){"use strict";r.d(t,{Z:function(){return m}});var n=r(3117),i=r(7294),a=r(4334),o=r(2389),s=r(7392),l=r(7094),c=r(2466),u="tabList__CuJ",p="tabItem_LNqP";function d(e){var t;const{lazy:r,block:o,defaultValue:d,values:m,groupId:g,className:b}=e,h=i.Children.map(e.children,(e=>{if((0,i.isValidElement)(e)&&"value"in e.props)return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)})),f=m??h.map((e=>{let{props:{value:t,label:r,attributes:n}}=e;return{value:t,label:r,attributes:n}})),y=(0,s.l)(f,((e,t)=>e.value===t.value));if(y.length>0)throw new Error(`Docusaurus error: Duplicate values "${y.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`);const v=null===d?d:d??(null==(t=h.find((e=>e.props.default)))?void 0:t.props.value)??h[0].props.value;if(null!==v&&!f.some((e=>e.value===v)))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${v}" but none of its children has the corresponding value. Available values are: ${f.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);const{tabGroupChoices:k,setTabGroupChoices:N}=(0,l.U)(),[w,C]=(0,i.useState)(v),_=[],{blockElementScrollPositionUntilNextRender:T}=(0,c.o5)();if(null!=g){const e=k[g];null!=e&&e!==w&&f.some((t=>t.value===e))&&C(e)}const I=e=>{const t=e.currentTarget,r=_.indexOf(t),n=f[r].value;n!==w&&(T(t),C(n),null!=g&&N(g,String(n)))},x=e=>{var t;let r=null;switch(e.key){case"ArrowRight":{const t=_.indexOf(e.currentTarget)+1;r=_[t]??_[0];break}case"ArrowLeft":{const t=_.indexOf(e.currentTarget)-1;r=_[t]??_[_.length-1];break}}null==(t=r)||t.focus()};return i.createElement("div",{className:(0,a.Z)("tabs-container",u)},i.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,a.Z)("tabs",{"tabs--block":o},b)},f.map((e=>{let{value:t,label:r,attributes:o}=e;return i.createElement("li",(0,n.Z)({role:"tab",tabIndex:w===t?0:-1,"aria-selected":w===t,key:t,ref:e=>_.push(e),onKeyDown:x,onFocus:I,onClick:I},o,{className:(0,a.Z)("tabs__item",p,null==o?void 0:o.className,{"tabs__item--active":w===t})}),r??t)}))),r?(0,i.cloneElement)(h.filter((e=>e.props.value===w))[0],{className:"margin-top--md"}):i.createElement("div",{className:"margin-top--md"},h.map(((e,t)=>(0,i.cloneElement)(e,{key:t,hidden:e.props.value!==w})))))}function m(e){const t=(0,o.Z)();return i.createElement(d,(0,n.Z)({key:String(t)},e))}},6679:function(e){e.exports={docsSidebar:["introduction",{type:"category",label:"Tour",items:["tour/introduction","tour/startup_siren_server","tour/registering_provider","tour/registering_receivers","tour/sending_notifications_to_receiver","tour/configuring_provider_alerting_rules","tour/subscribing_notifications"]},{type:"category",label:"Concepts",items:["concepts/overview","concepts/plugin","concepts/schema"]},{type:"category",label:"Guides",items:["guides/overview","guides/provider_and_namespace","guides/receiver","guides/subscription","guides/rule","guides/template","guides/alert_history","guides/notification","guides/deployment"]},{type:"category",label:"Contribute",items:["contribute/contribution","contribute/receiver","contribute/provider","contribute/release"]},{type:"category",label:"Reference",items:["reference/api","reference/server_configuration","reference/client_configuration","reference/receiver","reference/cli"]}]}},6746:function(e,t,r){"use strict";r.r(t),r.d(t,{apiVersion:function(){return b},assets:function(){return m},contentTitle:function(){return p},default:function(){return y},defaultHost:function(){return h},frontMatter:function(){return u},metadata:function(){return d},toc:function(){return g}});var n=r(3117),i=(r(7294),r(3905)),a=r(5488),o=r(5162),s=r(6066),l=r(1410),c=r.n(l);const u={},p="6 - Subscribing Notifications",d={unversionedId:"tour/subscribing_notifications",id:"tour/subscribing_notifications",title:"6 - Subscribing Notifications",description:"Notifications can be subscribed and routed to the defined receivers by adding a subscription. In this part, we will simulate how Cortex Ruler trigger an alert to Cortex Alertmanager, and Cortex Alertmanager trigger webhook-notification and calling Siren alerts hook API. On Siren side, we expect a notification is published everytime the hook API is being called.",source:"@site/docs/tour/6_subscribing_notifications.md",sourceDirName:"tour",slug:"/tour/subscribing_notifications",permalink:"/siren/docs/tour/subscribing_notifications",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/tour/6_subscribing_notifications.md",tags:[],version:"current",sidebarPosition:6,frontMatter:{},sidebar:"docsSidebar",previous:{title:"5 - Configuring Provider Alerting Rules",permalink:"/siren/docs/tour/configuring_provider_alerting_rules"},next:{title:"Overview",permalink:"/siren/docs/concepts/overview"}},m={},g=[],b=c().customFields.apiVersion,h=c().customFields.defaultHost,f={toc:g,apiVersion:b};function y(e){let{components:t,...r}=e;return(0,i.kt)("wrapper",(0,n.Z)({},f,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"6---subscribing-notifications"},"6 - Subscribing Notifications"),(0,i.kt)("p",null,"Notifications can be subscribed and routed to the defined receivers by adding a subscription. In this part, we will simulate how Cortex Ruler trigger an alert to Cortex Alertmanager, and Cortex Alertmanager trigger webhook-notification and calling Siren alerts hook API. On Siren side, we expect a notification is published everytime the hook API is being called."),(0,i.kt)("p",null,"In this part we will create alerting rules for our Cortex monitoring provider. Rules in Siren relies on ",(0,i.kt)("a",{parentName:"p",href:"/siren/docs/guides/template"},"template")," for its abstraction. We need to create a rule's template first before uploading a rule."),(0,i.kt)("p",null,"The first thing that we should do is knowing what would be the labels sent by Cortex Alertmanager. The labels should be defined when we were defining ",(0,i.kt)("a",{parentName:"p",href:"/siren/docs/tour/configuring_provider_alerting_rules#creating-a-rule"},"rules"),". Assuming the labels sent by Cortex Alertmanager are these:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"severity: WARNING\nteam: odpf\nservice: some-service\nenvironment: integration\nresource_name: some-resource\n")),(0,i.kt)("p",null,"Later we will try to simulate triggering alert by calling Cortex Alertmanager ",(0,i.kt)("inlineCode",{parentName:"p"},"POST /alerts")," API directly. The way Cortex Ruler monitor a specific metric and trigger an alert to Cortex Alertmanager are out of this ",(0,i.kt)("inlineCode",{parentName:"p"},"tour")," scope."),(0,i.kt)("p",null,"We want to subscribe all notifications owned by ",(0,i.kt)("inlineCode",{parentName:"p"},"odpf")," team and has severity ",(0,i.kt)("inlineCode",{parentName:"p"},"WARNING")," regardless the service name related with the alerts and route the notification to ",(0,i.kt)("inlineCode",{parentName:"p"},"file")," with receiver id ",(0,i.kt)("inlineCode",{parentName:"p"},"1")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"2"),". Currently there is no CLI to create a subscription (this would need to be added in the future) so we could call Siren HTTP API direclty to create one."),(0,i.kt)("p",null,"Prepare a subscription detail:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"cat <<EOT >> cpu_subs.yaml\nurn: subscribe-cpu-odpf-warning\nnamespace: 1\nreceivers:\n  - id: 1\n  - id: 2\nmatch\n  team: odpf\n  severity: WARNING\nEOT\n")),(0,i.kt)(a.Z,{groupId:"api",mdxType:"Tabs"},(0,i.kt)(o.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"./siren subscription create --file cpu_subs.yaml\n"))),(0,i.kt)(o.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,i.kt)(s.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request POST\n  --url ",h,"/",b,'/subscriptions\'\n--header \'Content-Type: application/json\'\n--header \'Accept: application/json\'\n--data-raw \'{\n  "urn": "subscribe-cpu-odpf-warning",\n  "namespace": 1,\n  "receivers": [\n    {\n      "id": 1\n    },\n    {\n      "id": 2\n    }\n  ],\n  "match": {\n    "team": "odpf",\n    "severity": "WARNING"\n  }\n}\''))),(0,i.kt)("p",null,"Once a subscription is created, let's simulate on how Cortex Ruler trigger an alert by calling Cortex Alertmanager API directly with this cURL."),(0,i.kt)(a.Z,{groupId:"api",mdxType:"Tabs"},(0,i.kt)(o.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'curl --location --request POST \'http://localhost:9009/api/prom/alertmanager/api/v1/alerts\'\n--header \'X-Scope-OrgId: odpf-ns\'\n--header \'Content-Type: application/json\' \\\n--data-raw \'[\n    {\n        "state": "firing",\n        "value": 1,\n        "labels": {\n            "severity": "WARNING",\n            "team": "odpf",\n            "service": "some-service",\n            "environment": "integration"\n        },\n        "annotations": {\n            "resource": "test_alert",\n            "metricName": "test_alert",\n            "metricValue": "1",\n            "template": "alert_test"\n        }\n    }\n]\'\n')))),(0,i.kt)("p",null,"If succeed, the response should be like this."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-json"},'{"status":"success"}\n')),(0,i.kt)("p",null,"Now, we need to expect Cortex Alertmanager send alerts to our Siren API ",(0,i.kt)("inlineCode",{parentName:"p"},"/alerts/cortex/:providerId"),". If that is the case, the alert should also be stored and published to the receivers in the matching subscriptions. You might want to wait for a Cortex Alertmanager ",(0,i.kt)("inlineCode",{parentName:"p"},"group_wait")," (usually 30s) until alerts are triggered by Cortex Alertmanager."),(0,i.kt)("p",null,"Let's verify the alert is stored inside our DB."),(0,i.kt)(a.Z,{groupId:"api",mdxType:"Tabs"},(0,i.kt)(o.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"./siren alert list --provider-id 1 --provider-type cortex --resource-name test_alert\n")),(0,i.kt)("p",null,"The result would be something like this."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"Showing 1 of 1 alerts\n \nID      PROVIDER_ID     RESOURCE_NAME   METRIC_NAME     METRIC_VALUE    SEVERITY\n1       1               test_alert      test_alert      1               WARNING \n\nFor details on a alert, try: siren alert view <id>\n"))),(0,i.kt)(o.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,i.kt)(s.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request GET\n  --url ",h,"/",b,"/alerts?providerId=1&providerType=cortex&resourceName=test_alert"))),(0,i.kt)("p",null,"We also expect notifications have been published to the receiver id ",(0,i.kt)("inlineCode",{parentName:"p"},"1")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"2")," similar with the ",(0,i.kt)("a",{parentName:"p",href:"/siren/docs/tour/sending_notifications_to_receiver"},"previous part"),". You can check a new notification is already added in ",(0,i.kt)("inlineCode",{parentName:"p"},"./out-file-sink1.json")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"./out-file-sink2.json")," with this value."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-json"},'{"environment":"integration","generatorUrl":"","groupKey":"{}:{severity=\\"WARNING\\"}","metricName":"test_alert","metricValue":"1","numAlertsFiring":1,"resource":"test_alert","routing_method":"subscribers","service":"some-service","severity":"WARNING","status":"firing","team":"odpf","template":"alert_test"}\n')))}y.isMDXComponent=!0},3618:function(e,t,r){"use strict";r.r(t);t.default={plain:{color:"#F8F8F2",backgroundColor:"#282A36"},styles:[{types:["prolog","constant","builtin"],style:{color:"rgb(189, 147, 249)"}},{types:["inserted","function"],style:{color:"rgb(80, 250, 123)"}},{types:["deleted"],style:{color:"rgb(255, 85, 85)"}},{types:["changed"],style:{color:"rgb(255, 184, 108)"}},{types:["punctuation","symbol"],style:{color:"rgb(248, 248, 242)"}},{types:["string","char","tag","selector"],style:{color:"rgb(255, 121, 198)"}},{types:["keyword","variable"],style:{color:"rgb(189, 147, 249)",fontStyle:"italic"}},{types:["comment"],style:{color:"rgb(98, 114, 164)"}},{types:["attr-name"],style:{color:"rgb(241, 250, 140)"}}]}},7694:function(e,t,r){"use strict";r.r(t);t.default={plain:{color:"#393A34",backgroundColor:"#f6f8fa"},styles:[{types:["comment","prolog","doctype","cdata"],style:{color:"#999988",fontStyle:"italic"}},{types:["namespace"],style:{opacity:.7}},{types:["string","attr-value"],style:{color:"#e3116c"}},{types:["punctuation","operator"],style:{color:"#393A34"}},{types:["entity","url","symbol","number","boolean","variable","constant","property","regex","inserted"],style:{color:"#36acaa"}},{types:["atrule","keyword","attr-name","selector"],style:{color:"#00a4db"}},{types:["function","deleted","tag"],style:{color:"#d73a49"}},{types:["function-variable"],style:{color:"#6f42c1"}},{types:["tag","selector","keyword"],style:{color:"#00009f"}}]}}}]);