"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[999],{3905:function(e,t,r){r.d(t,{Zo:function(){return p},kt:function(){return m}});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var l=n.createContext({}),c=function(e){var t=n.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},p=function(e){var t=c(e.components);return n.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),d=c(r),m=a,g=d["".concat(l,".").concat(m)]||d[m]||u[m]||i;return r?n.createElement(g,o(o({ref:t},p),{},{components:r})):n.createElement(g,o({ref:t},p))}));function m(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,o=new Array(i);o[0]=d;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var c=2;c<i;c++)o[c]=r[c];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},2822:function(e,t,r){r.r(t),r.d(t,{assets:function(){return l},contentTitle:function(){return o},default:function(){return u},frontMatter:function(){return i},metadata:function(){return s},toc:function(){return c}});var n=r(3117),a=(r(7294),r(3905));const i={},o="Overview",s={unversionedId:"concepts/overview",id:"concepts/overview",title:"Overview",description:"The following contains all the details about architecture, database schema, code structure and other technical concepts of Siren.",source:"@site/docs/concepts/overview.md",sourceDirName:"concepts",slug:"/concepts/overview",permalink:"/siren/docs/concepts/overview",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/concepts/overview.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"6 - Subscribing Notifications",permalink:"/siren/docs/tour/subscribing_notifications"},next:{title:"Plugin",permalink:"/siren/docs/concepts/plugin"}},l={},c=[{value:"Features",id:"features",level:2},{value:"Overall System Architecture",id:"overall-system-architecture",level:2},{value:"Siren Architecture",id:"siren-architecture",level:3},{value:"Provider: Cortex",id:"provider-cortex",level:3},{value:"Schema Design",id:"schema-design",level:2}],p={toc:c};function u(e){let{components:t,...i}=e;return(0,a.kt)("wrapper",(0,n.Z)({},p,i,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"overview"},"Overview"),(0,a.kt)("p",null,"The following contains all the details about architecture, database schema, code structure and other technical concepts of Siren."),(0,a.kt)("p",null,"Siren depends on monitoring and alerting provider for rule and alert creation. Siren expects provider to send alerts to a Siren webhook API and Siren manages the notification routing and publication as well as storing the alerts history. Siren stores templates, rules, subscriptions, and triggered alerts history in PostgresDB."),(0,a.kt)("h2",{id:"features"},"Features"),(0,a.kt)("p",null,(0,a.kt)("em",{parentName:"p"},(0,a.kt)("strong",{parentName:"em"},"Alerting Rules"))),(0,a.kt)("p",null,"Siren capables to manage alerting rules for various monitoring providers."),(0,a.kt)("p",null,(0,a.kt)("em",{parentName:"p"},(0,a.kt)("strong",{parentName:"em"},"Notification"))),(0,a.kt)("p",null,"Siren Notification provides easy to use commands to perform various actions. Currently, the actions supported are,\nstarting Siren Web Server, creating/updating templates and rules via a specified YAML file, migrating database\nschema, start a notification handler worker, and run a notification-related job. Read more about usage ",(0,a.kt)("a",{parentName:"p",href:"/siren/docs/guides/overview"},"here"),"."),(0,a.kt)("p",null,(0,a.kt)("em",{parentName:"p"},(0,a.kt)("strong",{parentName:"em"},"GRPC and HTTP API"))),(0,a.kt)("p",null,"GRPC Server exposes RPC APIs and RESTful APIs (via GRPC gateway) to allow configuration of rules, templates, alerting credentials and storing triggered alert history."),(0,a.kt)("p",null,(0,a.kt)("em",{parentName:"p"},(0,a.kt)("strong",{parentName:"em"},"CLI"))),(0,a.kt)("p",null,"Siren CLI provides easy to use commands to perform various actions. Currently, the actions supported are: starting Siren Server, creating/updating templates and rules via a specified YAML file, migrating database schema, start a notification handler worker, and run a notification-related. job Read more about usage ",(0,a.kt)("a",{parentName:"p",href:"/siren/docs/guides/overview"},"here"),"."),(0,a.kt)("h2",{id:"overall-system-architecture"},"Overall System Architecture"),(0,a.kt)("p",null,(0,a.kt)("img",{alt:"Siren Architecture",src:r(7792).Z,width:"1457",height:"1622"})),(0,a.kt)("p",null,"Let's have a look at the major components:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Provider:")," is a service/platform that does monitoring, observability, and alerting (e.g. Cortex, Influx). Provider is expected to send alerts information to siren via siren's webhook everytime alerts are triggered. Siren does alerting rules management and synchronize the rules with the provider.")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Upstream Services:")," are services that sends observability metrics (e.g. via telegraf, prometheus-exporter, open-telemetry) to a monitoring & alerting provider (e.g. Cortex, Influx). A provider will trigger an alert if the incoming metrics meet certain rules.")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Notification Vendor:")," is a service/platform that has capability to send notification (e.g. PagerDuty, Slack, etc). Siren has capability to store a specific notification vendor information and credentials, Siren call it ",(0,a.kt)("inlineCode",{parentName:"p"},"receiver"),". Siren ables to send notification to receivers."))),(0,a.kt)("h3",{id:"siren-architecture"},"Siren Architecture"),(0,a.kt)("p",null,(0,a.kt)("img",{alt:"Siren Detailed Architecture",src:r(3791).Z,width:"2227",height:"1684"})),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Server:")," is a main component in Siren that exposes GRPC & HTTP API within the same port. Client interacts with Siren ecosystem through Server. Server talks to a provider e.g. Cortex and DB to configure alerting rules using stored templates and configure alertmanager per tenant with the stored credentials per team.")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"PostgreSQL:")," is a main storage of Siren. Secret information is stored encrypted in the DB. ",(0,a.kt)("a",{parentName:"p",href:"/siren/docs/guides/notification"},"Notification")," in Siren requires a Queue and for current version Siren uses PostgreSQL as a queue. Siren maintains two schemas in the PostgreSQL, ",(0,a.kt)("inlineCode",{parentName:"p"},"public")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"notification"),". Siren uses ",(0,a.kt)("inlineCode",{parentName:"p"},"public")," schema as its main storage to store all data (e.g. templates, rules, receivers) except notification messages. Notification messages are stored in ",(0,a.kt)("inlineCode",{parentName:"p"},"notification")," schema and handled as a queue in PostgreSQL.")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Workers:")," are another instances in Siren that run detached (although possibly to run within the server too) from the server. Notification handler and dlq handler are workers that run with short period to dequeue notification messages and publish the messages to the notification vendors (e.g. slack, PagerDuty, etc)")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Job:")," is a task in Siren that could be executed and stopped once the task is done. The Job is usually run as a CronJob to be executed on a specified time."))),(0,a.kt)("h3",{id:"provider-cortex"},"Provider: Cortex"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Cortex Ruler:")," is a microservice of Cortex responsible to handle alerting rules. The configured rules are stored in Cortex Ruler. Siren Rules HTTP APIs call Cortex ruler to create/update/delete rule group in a particular namespace. You can create a ",(0,a.kt)("a",{parentName:"p",href:"/siren/docs/guides/provider_and_namespace"},"provider")," for that purpose and provide appropriate hostname.")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("strong",{parentName:"p"},"Cortex Alertmanager:")," is a microservice of Cortex responsible to handle alerting notification. Cortex sets up alert history webhook receiver to capture triggered alert history. Cortex Alertmanger hostname is fetched from ",(0,a.kt)("a",{parentName:"p",href:"/siren/docs/guides/provider_and_namespace"},"provider's")," host key."))),(0,a.kt)("h2",{id:"schema-design"},"Schema Design"),(0,a.kt)("p",null,"Siren uses PostgresDB to store rules, templates, triggered alerts history and alert configuration. Read in\nfurther detail ",(0,a.kt)("a",{parentName:"p",href:"/siren/docs/concepts/schema"},"here")))}u.isMDXComponent=!0},7792:function(e,t,r){t.Z=r.p+"assets/images/siren_architecture-d42a8532294fb14668be345823156f75.svg"},3791:function(e,t,r){t.Z=r.p+"assets/images/siren_detailed_architecture-7fb3c6cea11b9f2ebfd306e19c05eb0f.svg"}}]);