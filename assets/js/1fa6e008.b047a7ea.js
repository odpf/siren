"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[556],{3905:(t,e,a)=>{a.d(e,{Zo:()=>o,kt:()=>N});var n=a(7294);function r(t,e,a){return e in t?Object.defineProperty(t,e,{value:a,enumerable:!0,configurable:!0,writable:!0}):t[e]=a,t}function l(t,e){var a=Object.keys(t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(t);e&&(n=n.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),a.push.apply(a,n)}return a}function i(t){for(var e=1;e<arguments.length;e++){var a=null!=arguments[e]?arguments[e]:{};e%2?l(Object(a),!0).forEach((function(e){r(t,e,a[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(a)):l(Object(a)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(a,e))}))}return t}function p(t,e){if(null==t)return{};var a,n,r=function(t,e){if(null==t)return{};var a,n,r={},l=Object.keys(t);for(n=0;n<l.length;n++)a=l[n],e.indexOf(a)>=0||(r[a]=t[a]);return r}(t,e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(t);for(n=0;n<l.length;n++)a=l[n],e.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(t,a)&&(r[a]=t[a])}return r}var m=n.createContext({}),d=function(t){var e=n.useContext(m),a=e;return t&&(a="function"==typeof t?t(e):i(i({},e),t)),a},o=function(t){var e=d(t.components);return n.createElement(m.Provider,{value:e},t.children)},k={inlineCode:"code",wrapper:function(t){var e=t.children;return n.createElement(n.Fragment,{},e)}},u=n.forwardRef((function(t,e){var a=t.components,r=t.mdxType,l=t.originalType,m=t.parentName,o=p(t,["components","mdxType","originalType","parentName"]),u=d(a),N=r,g=u["".concat(m,".").concat(N)]||u[N]||k[N]||l;return a?n.createElement(g,i(i({ref:e},o),{},{components:a})):n.createElement(g,i({ref:e},o))}));function N(t,e){var a=arguments,r=e&&e.mdxType;if("string"==typeof t||r){var l=a.length,i=new Array(l);i[0]=u;var p={};for(var m in e)hasOwnProperty.call(e,m)&&(p[m]=e[m]);p.originalType=t,p.mdxType="string"==typeof t?t:r,i[1]=p;for(var d=2;d<l;d++)i[d]=a[d];return n.createElement.apply(null,i)}return n.createElement.apply(null,a)}u.displayName="MDXCreateElement"},7796:(t,e,a)=>{a.r(e),a.d(e,{assets:()=>m,contentTitle:()=>i,default:()=>k,frontMatter:()=>l,metadata:()=>p,toc:()=>d});var n=a(3117),r=(a(7294),a(3905));const l={},i="Schema Design",p={unversionedId:"concepts/schema",id:"concepts/schema",title:"Schema Design",description:"Siren stores providers, namespaces, templates, rules and triggered alerts history, receivers and subscriptions in",source:"@site/docs/concepts/schema.md",sourceDirName:"concepts",slug:"/concepts/schema",permalink:"/siren/docs/concepts/schema",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/concepts/schema.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Plugin",permalink:"/siren/docs/concepts/plugin"},next:{title:"Usage",permalink:"/siren/docs/guides/overview"}},m={},d=[{value:"Providers table",id:"providers-table",level:2},{value:"Namespace table",id:"namespace-table",level:2},{value:"Templates table",id:"templates-table",level:2},{value:"Rules Table",id:"rules-table",level:2},{value:"Receivers",id:"receivers",level:2},{value:"Subscriptions",id:"subscriptions",level:2}],o={toc:d};function k(t){let{components:e,...l}=t;return(0,r.kt)("wrapper",(0,n.Z)({},o,l,{components:e,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"schema-design"},"Schema Design"),(0,r.kt)("p",null,"Siren stores providers, namespaces, templates, rules and triggered alerts history, receivers and subscriptions in\nPostgresDB."),(0,r.kt)("p",null,(0,r.kt)("img",{alt:"Siren Architecture",src:a(8548).Z,width:"574",height:"475"})),(0,r.kt)("p",null,"There are the tables as of now as described below:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"providers:")," Stores the info of monitoring providers.")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"namespaces:")," Stores the info of tenancy inside a monitoring provider")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"alerts:")," Stores the triggered alert history.")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"templates:")," Stores the templates uploaded via HTTP APIs.")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"rules:")," Stores the rules configured and their state and thresholds defined.")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"receivers:")," Stores the info of notification mediums e.g. Slack, HTTP Webhook, Pagerduty etc.")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"subscriptions:")," Stores notification routing logic based on matching conditions"))),(0,r.kt)("h2",{id:"providers-table"},"Providers table"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Column"),(0,r.kt)("th",{parentName:"tr",align:null},"Type"),(0,r.kt)("th",{parentName:"tr",align:null},"Description"),(0,r.kt)("th",{parentName:"tr",align:null},"Example"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"id"),(0,r.kt)("td",{parentName:"tr",align:null},"bigint"),(0,r.kt)("td",{parentName:"tr",align:null},"Primary key"),(0,r.kt)("td",{parentName:"tr",align:null},"1")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"created_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Creation timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"updated_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Last update timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"name"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"name of the provider"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"localhost-cortex"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"urn"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"urn of the provider, should be unique"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"localhost-cortex"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"type"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"Type of monitoring provider (cortex/influx etc)"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"cortex"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"labels"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"generic kv pair that can be used for searching for appropriate row"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'{"org":"odpf"}'))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"credentials"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"any configuration data for that provider e.g. auth"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'{"bearer_token": "abcd"}'))))),(0,r.kt)("h2",{id:"namespace-table"},"Namespace table"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Column"),(0,r.kt)("th",{parentName:"tr",align:null},"Type"),(0,r.kt)("th",{parentName:"tr",align:null},"Description"),(0,r.kt)("th",{parentName:"tr",align:null},"Example"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"id"),(0,r.kt)("td",{parentName:"tr",align:null},"bigint"),(0,r.kt)("td",{parentName:"tr",align:null},"Primary key"),(0,r.kt)("td",{parentName:"tr",align:null},"1")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"created_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Creation timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"updated_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Last update timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"name"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"name of the namespace"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"odpf"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"urn"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"urn of the namespace, should be unique within the provider"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"odpf"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"labels"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"generic kv pair that can be used for searching for appropriate row"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'{"org":"odpf"}'))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"credentials"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"any configuration data for that namespace e.g. auth"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'{"bearer_token": "abcd"}'))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"provider_id"),(0,r.kt)("td",{parentName:"tr",align:null},"int"),(0,r.kt)("td",{parentName:"tr",align:null},"foreign key of provider to which this namespace belongs"),(0,r.kt)("td",{parentName:"tr",align:null},"4")))),(0,r.kt)("h2",{id:"templates-table"},"Templates table"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Column"),(0,r.kt)("th",{parentName:"tr",align:null},"Type"),(0,r.kt)("th",{parentName:"tr",align:null},"Description"),(0,r.kt)("th",{parentName:"tr",align:null},"Example"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"id"),(0,r.kt)("td",{parentName:"tr",align:null},"bigint"),(0,r.kt)("td",{parentName:"tr",align:null},"Primary key"),(0,r.kt)("td",{parentName:"tr",align:null},"1")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"created_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Creation timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"updated_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Last update timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"name"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"name of the template, should be unique"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"cpuHigh"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"tags"),(0,r.kt)("td",{parentName:"tr",align:null},"text[]"),(0,r.kt)("td",{parentName:"tr",align:null},"Tags array represented which resource types can use this template"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"{kafka, airflow}"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"body"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"Alert or recording rule body"),(0,r.kt)("td",{parentName:"tr",align:null},"See examples body in ",(0,r.kt)("a",{parentName:"td",href:"/siren/docs/guides/template"},"here"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"variables"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"JSON variable listing all variables in the body with their data type, description and default value."),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'[{"name": "for", "type": "string", "default": "bar", "description": "group period"}]'))))),(0,r.kt)("h2",{id:"rules-table"},"Rules Table"),(0,r.kt)("p",null,"Rules belong to a provider namespace, identified using an optional namespace, optional group_name and mandatory template\nand variables and status."),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Column"),(0,r.kt)("th",{parentName:"tr",align:null},"Type"),(0,r.kt)("th",{parentName:"tr",align:null},"Description"),(0,r.kt)("th",{parentName:"tr",align:null},"Example"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"id"),(0,r.kt)("td",{parentName:"tr",align:null},"bigint"),(0,r.kt)("td",{parentName:"tr",align:null},"Primary key"),(0,r.kt)("td",{parentName:"tr",align:null},"1")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"created_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Creation timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"updated_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Last update timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"namespace"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"the ",(0,r.kt)("a",{parentName:"td",href:"https://cortexmetrics.io/docs/api/#get-rule-groups-by-namespace"},"ruler namespace")," in which this rule should be created(optional key for providers which doesn't have a need for namespace)"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"kafka"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"group_name"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"the ruler group in which this rule should be created, optional key for provider where doesn't apply"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"testGroup"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"providerNamespace"),(0,r.kt)("td",{parentName:"tr",align:null},"int"),(0,r.kt)("td",{parentName:"tr",align:null},"foreign key of namespace in which rule should be created"),(0,r.kt)("td",{parentName:"tr",align:null},"4")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"enabled"),(0,r.kt)("td",{parentName:"tr",align:null},"bool"),(0,r.kt)("td",{parentName:"tr",align:null},"running status of alert (true or false)"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"true"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"template"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"the template which should be used for rule body"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"CPUHigh"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"name"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"name of the rule, must be unique, constructed as per ",(0,r.kt)("inlineCode",{parentName:"td"},"siren_api_providerURN_namespaceURN_namespace_groupName_template")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"siren_api_localhost-cortex_odpf_kafka_testGroup_cpuHigh"))))),(0,r.kt)("p",null,(0,r.kt)("strong",{parentName:"p"},"Alerts table:")),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Column"),(0,r.kt)("th",{parentName:"tr",align:null},"Type"),(0,r.kt)("th",{parentName:"tr",align:null},"Description"),(0,r.kt)("th",{parentName:"tr",align:null},"Example"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"id"),(0,r.kt)("td",{parentName:"tr",align:null},"bigint"),(0,r.kt)("td",{parentName:"tr",align:null},"Primary key"),(0,r.kt)("td",{parentName:"tr",align:null},"1")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"triggerd_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Triggered timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"updated_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Last update timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"created_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Creation timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"resource_name"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"resource on which the alert was triggered"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"kafkaMachine1"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"rule"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"name of template which used for this rule"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"cpuHigh"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"metric_name"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"the metric on which alert was triggered"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"cpu usgae %"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"metric_value"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"value of above metric on which the alert was triggered"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"95%"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"severity"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"severity level of alert (CRITICAL, WARNING, RESOLVED)"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"CRITICAL"))))),(0,r.kt)("h2",{id:"receivers"},"Receivers"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Column"),(0,r.kt)("th",{parentName:"tr",align:null},"Type"),(0,r.kt)("th",{parentName:"tr",align:null},"Description"),(0,r.kt)("th",{parentName:"tr",align:null},"Example"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"id"),(0,r.kt)("td",{parentName:"tr",align:null},"bigint"),(0,r.kt)("td",{parentName:"tr",align:null},"Primary key"),(0,r.kt)("td",{parentName:"tr",align:null},"1")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"created_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Creation timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"updated_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Last update timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"name"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"Name of receiver"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"siren-devs-slack-receivers"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"type"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"Type of receivers (Slack/HTTP/Pagerduty)"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"slack"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"labels"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"generic kv pair that can be used for searching for appropriate row"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'{"team":"siren-devs"}'))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"configuration"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"configuration data for that receiver(depends on the type)"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'{"token": "abcd", "workspace": "Odpf"}'))))),(0,r.kt)("h2",{id:"subscriptions"},"Subscriptions"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Column"),(0,r.kt)("th",{parentName:"tr",align:null},"Type"),(0,r.kt)("th",{parentName:"tr",align:null},"Description"),(0,r.kt)("th",{parentName:"tr",align:null},"Example"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"id"),(0,r.kt)("td",{parentName:"tr",align:null},"bigint"),(0,r.kt)("td",{parentName:"tr",align:null},"Primary key"),(0,r.kt)("td",{parentName:"tr",align:null},"1")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"created_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Creation timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"updated_at"),(0,r.kt)("td",{parentName:"tr",align:null},"timestamp with time zone"),(0,r.kt)("td",{parentName:"tr",align:null},"Last update timestamp"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"2021-03-05 12:37:56.905618+05:30"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"urn"),(0,r.kt)("td",{parentName:"tr",align:null},"text"),(0,r.kt)("td",{parentName:"tr",align:null},"URN of subscription"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"siren-devs-slack-subscription"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"receiver"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"list of receivers which will be picked for notification"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'[{"id":1, "configuration":{"channel_name":"siren-devs-critical"}}]'))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"match"),(0,r.kt)("td",{parentName:"tr",align:null},"jsonb"),(0,r.kt)("td",{parentName:"tr",align:null},"generic kv pair that must hold true in the alert for notification to be sent via the receives described"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},'{"token": "abcd", "workspace": "Odpf"}'))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"namespace_id"),(0,r.kt)("td",{parentName:"tr",align:null},"int"),(0,r.kt)("td",{parentName:"tr",align:null},"foreign key of namespace to which this belongs to"),(0,r.kt)("td",{parentName:"tr",align:null},"10")))))}k.isMDXComponent=!0},8548:(t,e,a)=>{a.d(e,{Z:()=>n});const n=a.p+"assets/images/siren_schema-28f56dfb99f8c65119cbb3d0d8422387.svg"}}]);