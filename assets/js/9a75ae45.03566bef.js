(self.webpackChunksiren=self.webpackChunksiren||[]).push([[255],{1410:(e,t,a)=>{const r=a(7694),n=a(3618);e.exports={title:"Siren",tagline:"Universal data observability toolkit",url:"https://odpf.github.io",baseUrl:"/siren/",onBrokenLinks:"throw",onBrokenMarkdownLinks:"warn",favicon:"img/favicon.ico",organizationName:"odpf",projectName:"siren",customFields:{apiVersion:"v1beta1",defaultHost:"http://localhost:8080"},presets:[["@docusaurus/preset-classic",{docs:{sidebarPath:6679,editUrl:"https://github.com/odpf/siren/edit/master/docs/",sidebarCollapsed:!0,breadcrumbs:!1},blog:!1,theme:{customCss:[5308,2295]},gtag:{trackingID:"G-EPXDLH6V72"}}]],themeConfig:{colorMode:{defaultMode:"light",respectPrefersColorScheme:!0},navbar:{title:"Siren",logo:{src:"img/logo.svg"},hideOnScroll:!0,items:[{type:"doc",docId:"introduction",position:"right",label:"Docs"},{to:"docs/support",label:"Support",position:"right"},{href:"https://bit.ly/2RzPbtn",position:"right",className:"header-slack-link"},{href:"https://github.com/odpf/siren",className:"navbar-item-github",position:"right"}]},footer:{style:"light",links:[]},prism:{theme:r,darkTheme:n},announcementBar:{id:"star-repo",content:'\u2b50\ufe0f If you like Siren, give it a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/odpf/siren">GitHub</a>! \u2b50',backgroundColor:"#222",textColor:"#eee",isCloseable:!0}}}},5162:(e,t,a)=>{"use strict";a.d(t,{Z:()=>l});var r=a(7294),n=a(4334);const i="tabItem_Ymn6";function l(e){let{children:t,hidden:a,className:l}=e;return r.createElement("div",{role:"tabpanel",className:(0,n.Z)(i,l),hidden:a},t)}},5488:(e,t,a)=>{"use strict";a.d(t,{Z:()=>m});var r=a(3117),n=a(7294),i=a(4334),l=a(2389),o=a(7392),s=a(7094),p=a(2466);const c="tabList__CuJ",u="tabItem_LNqP";function d(e){var t;const{lazy:a,block:l,defaultValue:d,values:m,groupId:g,className:h}=e,b=n.Children.map(e.children,(e=>{if((0,n.isValidElement)(e)&&"value"in e.props)return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)})),v=m??b.map((e=>{let{props:{value:t,label:a,attributes:r}}=e;return{value:t,label:a,attributes:r}})),k=(0,o.l)(v,((e,t)=>e.value===t.value));if(k.length>0)throw new Error(`Docusaurus error: Duplicate values "${k.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`);const f=null===d?d:d??(null==(t=b.find((e=>e.props.default)))?void 0:t.props.value)??b[0].props.value;if(null!==f&&!v.some((e=>e.value===f)))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${f}" but none of its children has the corresponding value. Available values are: ${v.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);const{tabGroupChoices:y,setTabGroupChoices:N}=(0,s.U)(),[w,T]=(0,n.useState)(f),C=[],{blockElementScrollPositionUntilNextRender:x}=(0,p.o5)();if(null!=g){const e=y[g];null!=e&&e!==w&&v.some((t=>t.value===e))&&T(e)}const I=e=>{const t=e.currentTarget,a=C.indexOf(t),r=v[a].value;r!==w&&(x(t),T(r),null!=g&&N(g,String(r)))},_=e=>{var t;let a=null;switch(e.key){case"Enter":I(e);break;case"ArrowRight":{const t=C.indexOf(e.currentTarget)+1;a=C[t]??C[0];break}case"ArrowLeft":{const t=C.indexOf(e.currentTarget)-1;a=C[t]??C[C.length-1];break}}null==(t=a)||t.focus()};return n.createElement("div",{className:(0,i.Z)("tabs-container",c)},n.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,i.Z)("tabs",{"tabs--block":l},h)},v.map((e=>{let{value:t,label:a,attributes:l}=e;return n.createElement("li",(0,r.Z)({role:"tab",tabIndex:w===t?0:-1,"aria-selected":w===t,key:t,ref:e=>C.push(e),onKeyDown:_,onClick:I},l,{className:(0,i.Z)("tabs__item",u,null==l?void 0:l.className,{"tabs__item--active":w===t})}),a??t)}))),a?(0,n.cloneElement)(b.filter((e=>e.props.value===w))[0],{className:"margin-top--md"}):n.createElement("div",{className:"margin-top--md"},b.map(((e,t)=>(0,n.cloneElement)(e,{key:t,hidden:e.props.value!==w})))))}function m(e){const t=(0,l.Z)();return n.createElement(d,(0,r.Z)({key:String(t)},e))}},6679:e=>{e.exports={docsSidebar:["introduction","use_cases","installation",{type:"category",label:"Tour",items:["tour/introduction","tour/setup_server","tour/1sending_notifications_overview","tour/2alerting_rules_subscriptions_overview"]},{type:"category",label:"Concepts",items:["concepts/overview","concepts/plugin","concepts/notification","concepts/glossary"]},{type:"category",label:"Guides",items:["guides/overview","guides/deployment","guides/provider_and_namespace","guides/receiver","guides/subscription","guides/rule","guides/template","guides/alert_history","guides/notification","guides/workers","guides/job"]},{type:"category",label:"Providers",items:["providers/cortexmetrics"]},{type:"category",label:"Receivers",items:["receivers/slack","receivers/pagerduty","receivers/http","receivers/file"]},{type:"category",label:"Reference",items:["reference/api","reference/server_configuration","reference/client_configuration","reference/cli"]},{type:"category",label:"Extend",items:["extend/adding_new_provider","extend/adding_new_receiver"]},{type:"category",label:"Contribute",items:["contribute/contribution","contribute/release"]}]}},916:(e,t,a)=>{"use strict";a.r(t),a.d(t,{apiVersion:()=>h,assets:()=>m,contentTitle:()=>u,default:()=>k,defaultHost:()=>b,frontMatter:()=>c,metadata:()=>d,toc:()=>g});var r=a(3117),n=(a(7294),a(3905)),i=a(5488),l=a(5162),o=a(6066),s=a(1410),p=a.n(s);const c={},u="2 Alerting Rules and Subscription",d={unversionedId:"tour/2alerting_rules_subscriptions_overview",id:"tour/2alerting_rules_subscriptions_overview",title:"2 Alerting Rules and Subscription",description:"This tour shows you how could we create alerting rules and we want to subscribe to a notification triggered by an alert. If you want to know how to send on-demand notification to a receiver, you could go to the first tour.",source:"@site/docs/tour/2alerting_rules_subscriptions_overview.md",sourceDirName:"tour",slug:"/tour/2alerting_rules_subscriptions_overview",permalink:"/siren/docs/tour/2alerting_rules_subscriptions_overview",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/tour/2alerting_rules_subscriptions_overview.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"1 Sending On-demand Notification",permalink:"/siren/docs/tour/1sending_notifications_overview"},next:{title:"Overview",permalink:"/siren/docs/concepts/overview"}},m={},g=[{value:"2.1 Register a Provider and Namespaces",id:"21-register-a-provider-and-namespaces",level:2},{value:"Register a Provider",id:"register-a-provider",level:3},{value:"Register Namespaces",id:"register-namespaces",level:3},{value:"Verify Created Provider and Namespaces",id:"verify-created-provider-and-namespaces",level:3},{value:"2.2 Register a Receiver",id:"22-register-a-receiver",level:2},{value:"2.3 Configuring Provider Alerting Rules",id:"23-configuring-provider-alerting-rules",level:2},{value:"Creating a Rule&#39;s Template",id:"creating-a-rules-template",level:3},{value:"Creating a Rule",id:"creating-a-rule",level:3},{value:"2.4 Subscribing to Alert Notifications",id:"24-subscribing-to-alert-notifications",level:2},{value:"What Next?",id:"what-next",level:2}],h=p().customFields.apiVersion,b=p().customFields.defaultHost,v={toc:g,apiVersion:h};function k(e){let{components:t,...a}=e;return(0,n.kt)("wrapper",(0,r.Z)({},v,a,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h1",{id:"2-alerting-rules-and-subscription"},"2 Alerting Rules and Subscription"),(0,n.kt)("p",null,"This tour shows you how could we create alerting rules and we want to subscribe to a notification triggered by an alert. If you want to know how to send on-demand notification to a receiver, you could go to the ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/tour/1sending_notifications_overview"},"first tour"),"."),(0,n.kt)("p",null,"As mentioned previously, we will be using ",(0,n.kt)("a",{parentName:"p",href:"https://cortexmetrics.io/docs/getting-started/"},"CortexMetrics")," as a provider. We need to register the provider and create a provider namespace in Siren first before creating any rule and subscription. "),(0,n.kt)("blockquote",null,(0,n.kt)("p",{parentName:"blockquote"},"Provider is implemented as a plugin in Siren. You can learn more about Siren Plugin concepts ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/concepts/plugin"},"here"),". We also welcome all contributions to add new provider plugins. Learn more about how to add a new provider plugin ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/extend/adding_new_provider"},"here"),".")),(0,n.kt)("p",null,"Once an alert triggered, the subscription labels will be matched with alert's labels. If all subscription labels matched, receiver's subscripton will get the alert notification."),(0,n.kt)("h2",{id:"21-register-a-provider-and-namespaces"},"2.1 Register a Provider and Namespaces"),(0,n.kt)("h3",{id:"register-a-provider"},"Register a Provider"),(0,n.kt)("p",null,"To create a new provider with CLI, we need to create a ",(0,n.kt)("inlineCode",{parentName:"p"},"yaml")," file that contains provider detail."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-yaml",metastring:"title=provider.yaml",title:"provider.yaml"},"host: http://localhost:9009\nurn: localhost-dev-cortex\nname: dev-cortex\ntype: cortex\n")),(0,n.kt)("p",null,"Once the file is ready, we can create the provider with Siren CLI."),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren provider create --file provider.yaml\n")),(0,n.kt)("p",null,"If succeed, you will got this message."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"Provider created with id: 1 \u2713\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request POST\n  --url ",b,"/",h,'/providers\n  --header \'content-type: application/json\'\n  --data-raw \'{\n    "host": "http://localhost:9009",\n    "urn": "localhost-dev-cortex",\n    "name": "dev-cortex",\n    "type": "cortex"\n}\''))),(0,n.kt)("p",null,"The ",(0,n.kt)("inlineCode",{parentName:"p"},"id")," we got from the provider creation is important to create a namespace later."),(0,n.kt)("h3",{id:"register-namespaces"},"Register Namespaces"),(0,n.kt)("p",null,"For multi-tenant scenario, which ",(0,n.kt)("a",{parentName:"p",href:"https://cortexmetrics.io/"},"CortexMetrics")," supports, we need to define namespaces in Siren. Assuming there are 2 tenants in Cortex, ",(0,n.kt)("inlineCode",{parentName:"p"},"odpf")," and ",(0,n.kt)("inlineCode",{parentName:"p"},"non-odpf"),", we need to create 2 namespaces. This could be done in similar way with how we created provider."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash",metastring:"title=ns1.yaml",title:"ns1.yaml"},"urn: odpf-ns\nname: odpf-ns\nprovider:\n    id: 1\n")),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash",metastring:"title=ns2.yaml",title:"ns2.yaml"},"urn: non-odpf-ns\nname: non-odpf-ns\nprovider:\n    id: 1\n")),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren namespace create --file ns1.yaml\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request POST\n  --url ",b,"/",h,'/namespaces\n  --header \'content-type: application/json\'\n  --data-raw \'{\n    "urn": "odpf-ns",\n    "name": "odpf-ns",\n    "provider": {\n        "id": 1\n    }\n}\''))),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren namespace create --file ns2.yaml\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request POST\n  --url ",b,"/",h,'/namespaces\n  --header \'content-type: application/json\'\n  --data-raw \'{\n    "urn": "non-odpf-ns",\n    "name": "non-odpf-ns",\n    "provider": {\n        "id": 2\n    }\n}\''))),(0,n.kt)("h3",{id:"verify-created-provider-and-namespaces"},"Verify Created Provider and Namespaces"),(0,n.kt)("p",null,"To make sure all provider and namespaces are properly created, we could try query Siren with Siren CLI."),(0,n.kt)("p",null,"See what providers exist in Siren."),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"$ siren provider list\n")),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"Showing 1 of 1 providers\n\nID      TYPE    URN                     NAME\n1       cortex  localhost-dev-cortex    dev-cortex\n\nFor details on a provider, try: siren provider view <id>\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request GET\n  --url ",b,"/",h,"/providers'"))),(0,n.kt)("p",null,"See what namespaces exist in Siren."),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"$ siren namespace list\n")),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"Showing 2 of 2 namespaces\n\nID      URN             NAME\n1       odpf-ns         odpf-ns\n2       non-odpf-ns     non-odpf-ns\n\nFor details on a namespace, try: siren namespace view <id>\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request GET\n  --url ",b,"/",h,"/namespaces'"))),(0,n.kt)("h2",{id:"22-register-a-receiver"},"2.2 Register a Receiver"),(0,n.kt)("p",null,"Siren supports several types of receiver to send notification to. For this tour, let's pick the simplest receiver: ",(0,n.kt)("inlineCode",{parentName:"p"},"file"),". If the receiver is not added in Siren yet, you could add one using ",(0,n.kt)("inlineCode",{parentName:"p"},"siren receiver create"),". See ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/guides/receiver"},"receiver guide")," to explore more on how to work with ",(0,n.kt)("inlineCode",{parentName:"p"},"siren receiver")," command."),(0,n.kt)("p",null,"With ",(0,n.kt)("inlineCode",{parentName:"p"},"file")," receiver, all published notifications will be written to a file. Let's create a receivers ",(0,n.kt)("inlineCode",{parentName:"p"},"file")," using Siren CLI."),(0,n.kt)("blockquote",null,(0,n.kt)("p",{parentName:"blockquote"},"We welcome all contributions to add new type of receiver plugins. See ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/extend/adding_new_receiver"},"Extend")," section to explore how to add a new type of receiver plugin to Siren")),(0,n.kt)("p",null,"Prepare receiver detail and register the receiver with Siren CLI."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash",metastring:"title=receiver_2.yaml",title:"receiver_2.yaml"},"name: file-sink-2\ntype: file\nlabels:\n    key1: value1\n    key2: value2\nconfigurations:\n    url: ./out-file-sink2.json\n")),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"$ siren receiver create --file receiver_2.yaml\n")),(0,n.kt)("p",null,"Once done, you will get a message."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash"},"Receiver created with id: 2 \u2713\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request POST\n  --url ",b,"/",h,'/receivers\n  --header \'content-type: application/json\'\n  --data-raw \'{\n    "name": "file-sink-2",\n    "type": "file",\n    "labels": {\n        "key1": "value1",\n        "key2": "value2"\n    },\n    "configurations": {\n        "url": "./out-file-sink2.json"\n    }\n}\''))),(0,n.kt)("h2",{id:"23-configuring-provider-alerting-rules"},"2.3 Configuring Provider Alerting Rules"),(0,n.kt)("p",null,"In this part, we will create alerting rules for our CortexMetrics monitoring provider. Rules in Siren relies on ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/guides/template"},"template")," for its abstraction. To create a rule, we need to create a template first."),(0,n.kt)("h3",{id:"creating-a-rules-template"},"Creating a Rule's Template"),(0,n.kt)("p",null,"We will create a rule's template to monitor CPU usage. "),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-yaml",metastring:"title=cpu_template.yaml",title:"cpu_template.yaml"},'apiVersion: v2\ntype: template\nname: CPU\nbody:\n    - alert: CPUWarning\n      expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > [[.warning]]\n      for: "[[.for]]"\n      labels:\n          severity: WARNING\n      annotations:\n          description: CPU has been above [[.warning]] for last [[.for]] {{ $labels.host }}\n    - alert: CPUCritical\n      expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > [[.critical]]\n      for: "[[.for]]"\n      labels:\n          severity: CRITICAL\n      annotations:\n          description: CPU has been above [[.critical]] for last [[.for]] {{ $labels.host }}\nvariables:\n    - name: for\n      type: string\n      default: 10m\n      description: For eg 5m, 2h; Golang duration format\n    - name: warning\n      type: int\n      default: 80\n    - name: critical\n      type: int\n      default: 90\ntags:\n    - systems\n')),(0,n.kt)("p",null,"We named the template above as ",(0,n.kt)("inlineCode",{parentName:"p"},"CPU"),", the body in the template is the data that will be interpolated with variables. Notice that template body format is similar with ",(0,n.kt)("a",{parentName:"p",href:"https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/"},"Prometheus alerting rules"),". This is because Cortex uses the same rules as prometheus and Siren will translate the rendered rules to the Cortex alerting rules. Let's save the template above into a file called ",(0,n.kt)("inlineCode",{parentName:"p"},"cpu_template.yaml")," and upload our template to Siren using "),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"$ siren template upload cpu_template.yaml\n")))),(0,n.kt)("p",null,"You could verify the newly created template using this command."),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"$ siren template list\n")),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"Showing 1 of 1 templates\n \nID      NAME    TAGS   \n1       CPU     systems\n\nFor details on a template, try: siren template view <name>\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request GET\n  --url ",b,"/",h,"/templates"))),(0,n.kt)("h3",{id:"creating-a-rule"},"Creating a Rule"),(0,n.kt)("p",null,"Now we already have a ",(0,n.kt)("inlineCode",{parentName:"p"},"CPU")," template, we can create a rule based on that template. Let's prepare a rule and save it in a file called ",(0,n.kt)("inlineCode",{parentName:"p"},"cpu_test.yaml"),"."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-yaml",metastring:"title=cpu_test.yaml",title:"cpu_test.yaml"},"apiVersion: v2\ntype: rule\nnamespace: odpf\nprovider: localhost-dev-cortex\nproviderNamespace: odpf-ns\nrules:\n    cpuGroup:\n        template: CPU\n        enabled: true\n        variables:\n            - name: for\n              value: 15m\n            - name: warning\n              value: 185\n            - name: critical\n              value: 195\n")),(0,n.kt)("p",null,"We defined a rule based on ",(0,n.kt)("inlineCode",{parentName:"p"},"CPU")," template for namespace urn ",(0,n.kt)("inlineCode",{parentName:"p"},"odpf-ns")," and provider urn ",(0,n.kt)("inlineCode",{parentName:"p"},"localhost-dev-cortex"),". The rule group name is ",(0,n.kt)("inlineCode",{parentName:"p"},"cpuGroup")," and there are also some variables to be assign to the template when the template is rendered. Let's upload the rule with Siren CLI."),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"$ siren rule upload cpu_test.yaml\n")),(0,n.kt)("p",null,"If succeed, you will get this message."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"Upserted Rule\nID: 4\n")))),(0,n.kt)("p",null,"You could verify the created rules by getting all registered rules in CortexMetrics with cURL."),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash"},"curl --location --request GET 'http://localhost:9009/api/v1/rules' \\\n--header 'X-Scope-OrgId: odpf-ns'\n")))),(0,n.kt)("p",null,"The response body should be in ",(0,n.kt)("inlineCode",{parentName:"p"},"yaml")," format and like this."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-yaml"},'odpf:\n    - name: cpuGroup\n      rules:\n        - alert: CPUWarning\n          expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > 185\n          for: 15m\n          labels:\n            severity: WARNING\n          annotations:\n            description: CPU has been above 185 for last 15m {{ $labels.host }}\n        - alert: CPUCritical\n          expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > 195\n          for: 15m\n          labels:\n            severity: CRITICAL\n          annotations:\n            description: CPU has been above 195 for last 15m {{ $labels.host }}\n')),(0,n.kt)("p",null,"If there is a response like above, that means the rule that we created in Siren was already synchronized to the provider. Next, we can add a subscription to the alert and try to trigger an alert to verify whether we got a notification alert or not."),(0,n.kt)("h2",{id:"24-subscribing-to-alert-notifications"},"2.4 Subscribing to Alert Notifications"),(0,n.kt)("p",null,"Notifications can be subscribed and routed to the defined receivers by adding a subscription. In this part, we will trigger an alert to CortexMetrics manually by calling CortexMetrics ",(0,n.kt)("inlineCode",{parentName:"p"},"POST /alerts")," API and expect CortexMetrics to trigger webhook-notification and calling Siren alerts hook API. On Siren side, we expect a notification is published everytime the hook API is being called."),(0,n.kt)("blockquote",null,(0,n.kt)("p",{parentName:"blockquote"},"If you are curious about how notification in Siren works, you can read the concepts ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/concepts/notification"},"here"),".")),(0,n.kt)("p",null,"The first thing that we should do is knowing what would be the labels sent by CortexMetrics. The labels should be defined when we were defining ",(0,n.kt)("a",{parentName:"p",href:"#23-configuring-provider-alerting-rules"},"rules"),". Assuming the labels sent by CortexMetrics are these:"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-yaml"},"severity: WARNING\nteam: odpf\nservice: some-service\nenvironment: integration\nresource_name: some-resource\n")),(0,n.kt)("p",null,"We want to subscribe all notifications owned by ",(0,n.kt)("inlineCode",{parentName:"p"},"odpf")," team and has severity ",(0,n.kt)("inlineCode",{parentName:"p"},"WARNING")," regardless the service name and route the notification to ",(0,n.kt)("inlineCode",{parentName:"p"},"file")," with receiver id ",(0,n.kt)("inlineCode",{parentName:"p"},"2")," (the one that we created in the ",(0,n.kt)("a",{parentName:"p",href:"#22-register-a-receiver"},"previous")," part)."),(0,n.kt)("p",null,"Prepare a subscription detail and create a new subscription with Siren CLI."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash",metastring:"title=cpu_subs.yaml",title:"cpu_subs.yaml"},"urn: subscribe-cpu-odpf-warning\nnamespace: 1\nreceivers:\n  - id: 1\n  - id: 2\nmatch\n  team: odpf\n  severity: WARNING\n")),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"$ siren subscription create --file cpu_subs.yaml\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request POST\n  --url ",b,"/",h,'/subscriptions\'\n--header \'Content-Type: application/json\'\n--header \'Accept: application/json\'\n--data-raw \'{\n  "urn": "subscribe-cpu-odpf-warning",\n  "namespace": 1,\n  "receivers": [\n    {\n      "id": 1\n    },\n    {\n      "id": 2\n    }\n  ],\n  "match": {\n    "team": "odpf",\n    "severity": "WARNING"\n  }\n}\''))),(0,n.kt)("p",null,"Once a subscription is created, let's manually trigger alert in CortexMetrics with this cURL. The way CortexMetrics monitor a specific metric and auto-trigger an alert are out of this ",(0,n.kt)("inlineCode",{parentName:"p"},"tour")," scope."),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-bash"},'curl --location --request POST \'http://localhost:9009/api/prom/alertmanager/api/v1/alerts\'\n--header \'X-Scope-OrgId: odpf-ns\'\n--header \'Content-Type: application/json\' \\\n--data-raw \'[\n    {\n        "state": "firing",\n        "value": 1,\n        "labels": {\n            "severity": "WARNING",\n            "team": "odpf",\n            "service": "some-service",\n            "environment": "integration"\n        },\n        "annotations": {\n            "resource": "test_alert",\n            "metricName": "test_alert",\n            "metricValue": "1",\n            "template": "alert_test"\n        }\n    }\n]\'\n')))),(0,n.kt)("p",null,"If succeed, the response should be like this."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-json"},'{"status":"success"}\n')),(0,n.kt)("p",null,"Now, we need to expect CortexMetrics to send alerts notification to our Siren API ",(0,n.kt)("inlineCode",{parentName:"p"},"/alerts/cortex/:providerId"),". If that is the case, the alert should also be stored and published to the receivers in the matching subscriptions. You might want to wait for a CortexMetrics ",(0,n.kt)("inlineCode",{parentName:"p"},"group_wait")," (usually 30s) until alerts are triggered by Cortex Alertmanager."),(0,n.kt)("p",null,"Let's verify the alert is stored inside our DB."),(0,n.kt)(i.Z,{groupId:"api",mdxType:"Tabs"},(0,n.kt)(l.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"$ siren alert list --provider-id 1 --provider-type cortex --resource-name test_alert\n")),(0,n.kt)("p",null,"The result would be something like this."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-shell"},"Showing 1 of 1 alerts\n \nID      PROVIDER_ID     RESOURCE_NAME   METRIC_NAME     METRIC_VALUE    SEVERITY\n1       1               test_alert      test_alert      1               WARNING \n\nFor details on a alert, try: siren alert view <id>\n"))),(0,n.kt)(l.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,n.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request GET\n  --url ",b,"/",h,"/alerts?providerId=1&providerType=cortex&resourceName=test_alert"))),(0,n.kt)("p",null,"We also expect notifications have been published to the receiver id ",(0,n.kt)("inlineCode",{parentName:"p"},"2"),". You can check a new notification is already added in ",(0,n.kt)("inlineCode",{parentName:"p"},"./out-file-sink2.json")," with this value."),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-json"},'{"environment":"integration","generatorUrl":"","groupKey":"{}:{severity=\\"WARNING\\"}","metricName":"test_alert","metricValue":"1","numAlertsFiring":1,"resource":"test_alert","routing_method":"subscribers","service":"some-service","severity":"WARNING","status":"firing","team":"odpf","template":"alert_test"}\n')),(0,n.kt)("h2",{id:"what-next"},"What Next?"),(0,n.kt)("p",null,"This is the end of ",(0,n.kt)("inlineCode",{parentName:"p"},"alerting rules and subscription")," tour. If you want to know how to send on-demand notification to a receiver, you could check the ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/tour/1sending_notifications_overview"},"first tour"),"."),(0,n.kt)("p",null,"Apart from the tour, we recommend completing the ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/guides/overview"},"guides"),". You could also check out the remainder of the documentation in the ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/reference/server_configuration"},"reference")," and ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/concepts/overview"},"concepts")," sections for your specific areas of interest. We've aimed to provide as much documentation as we can for the various components of Siren to give you a full understanding of Siren's surface area. If you are interested to contribute, check out the ",(0,n.kt)("a",{parentName:"p",href:"/siren/docs/contribute/contribution"},"contribution")," page."))}k.isMDXComponent=!0},3618:(e,t,a)=>{"use strict";a.r(t),a.d(t,{default:()=>r});const r={plain:{color:"#F8F8F2",backgroundColor:"#282A36"},styles:[{types:["prolog","constant","builtin"],style:{color:"rgb(189, 147, 249)"}},{types:["inserted","function"],style:{color:"rgb(80, 250, 123)"}},{types:["deleted"],style:{color:"rgb(255, 85, 85)"}},{types:["changed"],style:{color:"rgb(255, 184, 108)"}},{types:["punctuation","symbol"],style:{color:"rgb(248, 248, 242)"}},{types:["string","char","tag","selector"],style:{color:"rgb(255, 121, 198)"}},{types:["keyword","variable"],style:{color:"rgb(189, 147, 249)",fontStyle:"italic"}},{types:["comment"],style:{color:"rgb(98, 114, 164)"}},{types:["attr-name"],style:{color:"rgb(241, 250, 140)"}}]}},7694:(e,t,a)=>{"use strict";a.r(t),a.d(t,{default:()=>r});const r={plain:{color:"#393A34",backgroundColor:"#f6f8fa"},styles:[{types:["comment","prolog","doctype","cdata"],style:{color:"#999988",fontStyle:"italic"}},{types:["namespace"],style:{opacity:.7}},{types:["string","attr-value"],style:{color:"#e3116c"}},{types:["punctuation","operator"],style:{color:"#393A34"}},{types:["entity","url","symbol","number","boolean","variable","constant","property","regex","inserted"],style:{color:"#36acaa"}},{types:["atrule","keyword","attr-name","selector"],style:{color:"#00a4db"}},{types:["function","deleted","tag"],style:{color:"#d73a49"}},{types:["function-variable"],style:{color:"#6f42c1"}},{types:["tag","selector","keyword"],style:{color:"#00009f"}}]}}}]);