(self.webpackChunksiren=self.webpackChunksiren||[]).push([[889],{1410:function(e,t,i){const n=i(7694),o=i(3618);e.exports={title:"Siren",tagline:"Universal data observability toolkit",url:"https://odpf.github.io",baseUrl:"/siren/",onBrokenLinks:"throw",onBrokenMarkdownLinks:"warn",favicon:"img/favicon.ico",organizationName:"odpf",projectName:"siren",customFields:{apiVersion:"v1beta1",defaultHost:"http://localhost:8080"},presets:[["@docusaurus/preset-classic",{docs:{sidebarPath:6679,editUrl:"https://github.com/odpf/siren/edit/master/docs/",sidebarCollapsed:!0,breadcrumbs:!1},blog:!1,theme:{customCss:[5308,2295]},gtag:{trackingID:"G-EPXDLH6V72"}}]],themeConfig:{colorMode:{defaultMode:"light",respectPrefersColorScheme:!0},navbar:{title:"Siren",logo:{src:"img/logo.svg"},hideOnScroll:!0,items:[{type:"doc",docId:"introduction",position:"right",label:"Docs"},{to:"docs/support",label:"Support",position:"right"},{href:"https://bit.ly/2RzPbtn",position:"right",className:"header-slack-link"},{href:"https://github.com/odpf/siren",className:"navbar-item-github",position:"right"}]},footer:{style:"light",links:[]},prism:{theme:n,darkTheme:o},announcementBar:{id:"star-repo",content:'\u2b50\ufe0f If you like Siren, give it a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/odpf/siren">GitHub</a>! \u2b50',backgroundColor:"#222",textColor:"#eee",isCloseable:!0}}}},5162:function(e,t,i){"use strict";i.d(t,{Z:function(){return a}});var n=i(7294),o=i(4334),r="tabItem_Ymn6";function a(e){let{children:t,hidden:i,className:a}=e;return n.createElement("div",{role:"tabpanel",className:(0,o.Z)(r,a),hidden:i},t)}},5488:function(e,t,i){"use strict";i.d(t,{Z:function(){return f}});var n=i(3117),o=i(7294),r=i(4334),a=i(2389),s=i(7392),l=i(7094),c=i(2466),u="tabList__CuJ",d="tabItem_LNqP";function p(e){var t;const{lazy:i,block:a,defaultValue:p,values:f,groupId:m,className:b}=e,g=o.Children.map(e.children,(e=>{if((0,o.isValidElement)(e)&&"value"in e.props)return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)})),h=f??g.map((e=>{let{props:{value:t,label:i,attributes:n}}=e;return{value:t,label:i,attributes:n}})),y=(0,s.l)(h,((e,t)=>e.value===t.value));if(y.length>0)throw new Error(`Docusaurus error: Duplicate values "${y.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`);const v=null===p?p:p??(null==(t=g.find((e=>e.props.default)))?void 0:t.props.value)??g[0].props.value;if(null!==v&&!h.some((e=>e.value===v)))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${v}" but none of its children has the corresponding value. Available values are: ${h.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);const{tabGroupChoices:k,setTabGroupChoices:N}=(0,l.U)(),[w,S]=(0,o.useState)(v),T=[],{blockElementScrollPositionUntilNextRender:_}=(0,c.o5)();if(null!=m){const e=k[m];null!=e&&e!==w&&h.some((t=>t.value===e))&&S(e)}const x=e=>{const t=e.currentTarget,i=T.indexOf(t),n=h[i].value;n!==w&&(_(t),S(n),null!=m&&N(m,String(n)))},C=e=>{var t;let i=null;switch(e.key){case"ArrowRight":{const t=T.indexOf(e.currentTarget)+1;i=T[t]??T[0];break}case"ArrowLeft":{const t=T.indexOf(e.currentTarget)-1;i=T[t]??T[T.length-1];break}}null==(t=i)||t.focus()};return o.createElement("div",{className:(0,r.Z)("tabs-container",u)},o.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,r.Z)("tabs",{"tabs--block":a},b)},h.map((e=>{let{value:t,label:i,attributes:a}=e;return o.createElement("li",(0,n.Z)({role:"tab",tabIndex:w===t?0:-1,"aria-selected":w===t,key:t,ref:e=>T.push(e),onKeyDown:C,onFocus:x,onClick:x},a,{className:(0,r.Z)("tabs__item",d,null==a?void 0:a.className,{"tabs__item--active":w===t})}),i??t)}))),i?(0,o.cloneElement)(g.filter((e=>e.props.value===w))[0],{className:"margin-top--md"}):o.createElement("div",{className:"margin-top--md"},g.map(((e,t)=>(0,o.cloneElement)(e,{key:t,hidden:e.props.value!==w})))))}function f(e){const t=(0,a.Z)();return o.createElement(p,(0,n.Z)({key:String(t)},e))}},6679:function(e){e.exports={docsSidebar:["introduction",{type:"category",label:"Tour",items:["tour/introduction","tour/startup_siren_server","tour/registering_provider","tour/registering_receivers","tour/sending_notifications_to_receiver","tour/configuring_provider_alerting_rules","tour/subscribing_notifications"]},{type:"category",label:"Concepts",items:["concepts/overview","concepts/plugin","concepts/schema"]},{type:"category",label:"Guides",items:["guides/overview","guides/provider_and_namespace","guides/receiver","guides/subscription","guides/rule","guides/template","guides/alert_history","guides/notification","guides/deployment"]},{type:"category",label:"Contribute",items:["contribute/contribution","contribute/receiver","contribute/provider","contribute/release"]},{type:"category",label:"Reference",items:["reference/api","reference/server_configuration","reference/client_configuration","reference/receiver","reference/cli"]}]}},629:function(e,t,i){"use strict";i.r(t),i.d(t,{apiVersion:function(){return b},assets:function(){return f},contentTitle:function(){return d},default:function(){return y},defaultHost:function(){return g},frontMatter:function(){return u},metadata:function(){return p},toc:function(){return m}});var n=i(3117),o=(i(7294),i(3905)),r=i(5488),a=i(5162),s=i(6066),l=i(1410),c=i.n(l);const u={},d="Notification",p={unversionedId:"guides/notification",id:"guides/notification",title:"Notification",description:"Notification is one of main features in Siren. Siren capables to send notification to various receivers (e.g. Slack, PagerDuty). Notification in Siren could be sent directly to a receiver or user could subscribe notifications by providing key-value label matchers. For the latter, Siren routes notification to specific receivers by matching notification key-value labels with the provided label matchers.",source:"@site/docs/guides/notification.md",sourceDirName:"guides",slug:"/guides/notification",permalink:"/siren/docs/guides/notification",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/guides/notification.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Alert History",permalink:"/siren/docs/guides/alert_history"},next:{title:"Deployment",permalink:"/siren/docs/guides/deployment"}},f={},m=[{value:"Sending a message/notification",id:"sending-a-messagenotification",level:2},{value:"Example: Sending Notification to Slack",id:"example-sending-notification-to-slack",level:3},{value:"Alerts Notification",id:"alerts-notification",level:2}],b=c().customFields.apiVersion,g=c().customFields.defaultHost,h={toc:m,apiVersion:b};function y(e){let{components:t,...i}=e;return(0,o.kt)("wrapper",(0,n.Z)({},h,i,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"notification"},"Notification"),(0,o.kt)("p",null,"Notification is one of main features in Siren. Siren capables to send notification to various receivers (e.g. Slack, PagerDuty). Notification in Siren could be sent directly to a receiver or user could subscribe notifications by providing key-value label matchers. For the latter, Siren routes notification to specific receivers by matching notification key-value labels with the provided label matchers."),(0,o.kt)("h2",{id:"sending-a-messagenotification"},"Sending a message/notification"),(0,o.kt)("p",null,"We could send a notification to a specific receiver by passing a ",(0,o.kt)("inlineCode",{parentName:"p"},"receiver_id")," in the path params and correct payload format in the body. The payload format needs to follow receiver type contract. "),(0,o.kt)("h3",{id:"example-sending-notification-to-slack"},"Example: Sending Notification to Slack"),(0,o.kt)("p",null,"If receiver is slack, the ",(0,o.kt)("inlineCode",{parentName:"p"},"payload.data")," should be within the expected ",(0,o.kt)("a",{parentName:"p",href:"#slack"},"slack")," payload format."),(0,o.kt)(r.Z,{groupId:"api",mdxType:"Tabs"},(0,o.kt)(a.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren receiver create --file receiver.yaml\n"))),(0,o.kt)(a.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,o.kt)(s.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request POST\n  --url ",g,"/",b,'/receivers/51/send\n  --header \'content-type: application/json\'\n  --data-raw \'{\n    "payload": {\n        "data": {\n            "channel": "siren-devs",\n            "text": "an alert or notification",\n            "icon_emoji": ":smile:"\n            "attachments": [\n                "blocks": [\n                    {\n                        "type": "section",\n                        "text": {\n                            "type": "mrkdwn",\n                            "text": "New Paid Time Off request from <example.com|Fred Enriquez>\n\n<https://example.com|View request>"\n                        }\n                    }\n                ]\n            ]\n        }\n    }\n}\''))),(0,o.kt)("p",null,"Above end the message to channel name ",(0,o.kt)("inlineCode",{parentName:"p"},"#siren-devs")," with ",(0,o.kt)("inlineCode",{parentName:"p"},"payload.data")," in ",(0,o.kt)("a",{parentName:"p",href:"#slack"},"slack")," payload format."),(0,o.kt)("h2",{id:"alerts-notification"},"Alerts Notification"),(0,o.kt)("p",null,"For all incoming alerts via Siren hook API, notifications are also generated and published via subscriptions. Siren will match labels from the alerts with label matchers in subscriptions. The assigned receivers for all matched subscriptions will get the notifications. More details are explained ",(0,o.kt)("a",{parentName:"p",href:"/siren/docs/guides/alert_history"},"here"),". Sending notification message requires notification message payload to be in the same format as what receiver expected. The format can be found in the detail in ",(0,o.kt)("a",{parentName:"p",href:"/siren/docs/reference/receiver"},"reference"),"."))}y.isMDXComponent=!0},3618:function(e,t,i){"use strict";i.r(t);t.default={plain:{color:"#F8F8F2",backgroundColor:"#282A36"},styles:[{types:["prolog","constant","builtin"],style:{color:"rgb(189, 147, 249)"}},{types:["inserted","function"],style:{color:"rgb(80, 250, 123)"}},{types:["deleted"],style:{color:"rgb(255, 85, 85)"}},{types:["changed"],style:{color:"rgb(255, 184, 108)"}},{types:["punctuation","symbol"],style:{color:"rgb(248, 248, 242)"}},{types:["string","char","tag","selector"],style:{color:"rgb(255, 121, 198)"}},{types:["keyword","variable"],style:{color:"rgb(189, 147, 249)",fontStyle:"italic"}},{types:["comment"],style:{color:"rgb(98, 114, 164)"}},{types:["attr-name"],style:{color:"rgb(241, 250, 140)"}}]}},7694:function(e,t,i){"use strict";i.r(t);t.default={plain:{color:"#393A34",backgroundColor:"#f6f8fa"},styles:[{types:["comment","prolog","doctype","cdata"],style:{color:"#999988",fontStyle:"italic"}},{types:["namespace"],style:{opacity:.7}},{types:["string","attr-value"],style:{color:"#e3116c"}},{types:["punctuation","operator"],style:{color:"#393A34"}},{types:["entity","url","symbol","number","boolean","variable","constant","property","regex","inserted"],style:{color:"#36acaa"}},{types:["atrule","keyword","attr-name","selector"],style:{color:"#00a4db"}},{types:["function","deleted","tag"],style:{color:"#d73a49"}},{types:["function-variable"],style:{color:"#6f42c1"}},{types:["tag","selector","keyword"],style:{color:"#00009f"}}]}}}]);