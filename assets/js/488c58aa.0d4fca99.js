"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[156],{3905:function(e,t,n){n.d(t,{Zo:function(){return p},kt:function(){return h}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),l=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=l(e.components);return r.createElement(c.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,c=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),d=l(n),h=a,m=d["".concat(c,".").concat(h)]||d[h]||u[h]||i;return n?r.createElement(m,o(o({ref:t},p),{},{components:n})):r.createElement(m,o({ref:t},p))}));function h(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=d;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var l=2;l<i;l++)o[l]=n[l];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},719:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return s},contentTitle:function(){return c},metadata:function(){return l},toc:function(){return p},default:function(){return d}});var r=n(7462),a=n(3366),i=(n(7294),n(3905)),o=["components"],s={},c="Receivers",l={unversionedId:"guides/receivers",id:"guides/receivers",isDocsHomePage:!1,title:"Receivers",description:"Receivers represent a notification medium, which can be used to define routing configuration in the monitoring",source:"@site/docs/guides/receivers.md",sourceDirName:"guides",slug:"/guides/receivers",permalink:"/siren/docs/guides/receivers",editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/guides/receivers.md",tags:[],version:"current",lastUpdatedBy:"Abhishek Sah",lastUpdatedAt:1646811080,formattedLastUpdatedAt:"3/9/2022",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Providers",permalink:"/siren/docs/guides/providers"},next:{title:"Subscriptions",permalink:"/siren/docs/guides/subscriptions"}},p=[{value:"API Interface",id:"api-interface",children:[{value:"Create a Receiver",id:"create-a-receiver",children:[]},{value:"Update a Receiver",id:"update-a-receiver",children:[]},{value:"Getting a receiver",id:"getting-a-receiver",children:[]},{value:"Getting all receivers",id:"getting-all-receivers",children:[]},{value:"Deleting a receiver",id:"deleting-a-receiver",children:[]},{value:"Sending a message/notification",id:"sending-a-messagenotification",children:[]}]},{value:"CLI Interface",id:"cli-interface",children:[{value:"Permissions and Auth Settings for Slack Receivers",id:"permissions-and-auth-settings-for-slack-receivers",children:[]}]}],u={toc:p};function d(e){var t=e.components,n=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"receivers"},"Receivers"),(0,i.kt)("p",null,"Receivers represent a notification medium, which can be used to define routing configuration in the monitoring\nproviders, to control the behaviour of how your alerts are notified. Few examples: Slack receiver, HTTP receiver,\nPagerduty receivers etc. Currently, Siren supports these 3 types of receivers. Configuration of each receiver depends on\nthe type."),(0,i.kt)("p",null,"You can use receivers to send notifications on demand as well as on certain matching conditions. Subscriptions use\nreceivers to define routing configuration in monitoring providers. For eg. Cortex-metrics uses alertmanager for routing\nalerts. With Siren subscriptions, you will be able to manage routing in Alertmanager using pre-registered receivers."),(0,i.kt)("h2",{id:"api-interface"},"API Interface"),(0,i.kt)("h3",{id:"create-a-receiver"},"Create a Receiver"),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Type: Slack")),(0,i.kt)("p",null,"Using a slack receiver you will be able to send out Slack notification using its send API. You can also use it to route\nalerts using Subscriptions whenever an alert matches the conditions of your choice. Check the required permissions of\nthe Slack App ",(0,i.kt)("a",{parentName:"p",href:"#permissions-and-auth-settings-for-slack-receivers"},"below"),"."),(0,i.kt)("p",null,"Creating a slack receiver involves exchanging the auth code for a token with slack oauth server. Siren will need the\nauth code, client id, client secret and optional label metadata."),(0,i.kt)("p",null,"Example"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'POST /v1beta1/receivers HTTP/1.1\nHost: localhost:3000\nContent-Type: application/json\nContent-Length: 228\n\n{\n    "name": "doc-slack-receiver",\n    "type": "slack",\n    "labels": {\n        "team": "siren-devs"\n    },\n    "configurations": {\n        "client_id": "abcd",\n        "client_secret": "xyz",\n        "auth_code": "123"\n    }\n}\n')),(0,i.kt)("p",null,"On success, this will store the app token for that particular slack workspace and use it for sending out notifications."),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Type: Pagerduty")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'POST /v1beta1/receivers HTTP/1.1\nHost: localhost:3000\nContent-Type: application/json\nContent-Length: 182\n\n{\n    "name": "doc-pagerduty-receiver",\n    "type": "http",\n    "labels": {\n        "team": "siren-devs"\n    },\n    "configurations": {\n        "url": "http://localhost:4000"\n    }\n}\n')),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Type: HTTP")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'POST /v1beta1/receivers HTTP/1.1\nHost: localhost:3000\nContent-Type: application/json\nContent-Length: 177\n\n{\n    "name": "doc-http-receiver",\n    "type": "http",\n    "labels": {\n        "team": "siren-devs"\n    },\n    "configurations": {\n        "url": "http://localhost:4000"\n    }\n}\n')),(0,i.kt)("h3",{id:"update-a-receiver"},"Update a Receiver"),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Note:")," While updating a receiver, you will have to make sure all subscriptions that are using this receivers get\nrefreshed(updated), since subscriptions use receivers to create routing configuration dynamically."),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Type: HTTP")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'PUT /v1beta1/receivers/61 HTTP/1.1\nHost: localhost:3000\nContent-Type: application/json\nContent-Length: 177\n\n{\n    "name": "doc-http-receiver",\n    "type": "http",\n    "labels": {\n        "team": "siren-devs"\n    },\n    "configurations": {\n        "url": "http://localhost:4001"\n    }\n}\n')),(0,i.kt)("h3",{id:"getting-a-receiver"},"Getting a receiver"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},"GET /v1beta1/receivers/61 HTTP/1.1\nHost: localhost:3000\n")),(0,i.kt)("h3",{id:"getting-all-receivers"},"Getting all receivers"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},"GET /v1beta1/receivers HTTP/1.1\nHost: localhost:3000\n")),(0,i.kt)("h3",{id:"deleting-a-receiver"},"Deleting a receiver"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},"DELETE /v1beta1/receivers/61 HTTP/1.1\nHost: localhost:3000\n")),(0,i.kt)("h3",{id:"sending-a-messagenotification"},"Sending a message/notification"),(0,i.kt)("p",null,"The types that supports sending messages using API are:"),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Type: Slack")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'GET /v1beta1/receivers/51/send HTTP/1.1\nHost: localhost:3000\nContent-Type: application/json\nContent-Length: 399\n\n{\n    "slack": {\n        "receiverName": "siren-devs",\n        "receiverType": "channel",\n        "blocks": [\n            {\n                "type": "section",\n                "text": {\n                    "type": "mrkdwn",\n                    "text": "New Paid Time Off request from <example.com|Fred Enriquez>\\n\\n<https://example.com|View request>"\n                }\n            }\n        ]\n    }\n}\n')),(0,i.kt)("p",null,"Here we are using slack builder kit to construct block of messages, to send the message to channel name #siren-devs."),(0,i.kt)("h2",{id:"cli-interface"},"CLI Interface"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'Receivers are the medium to send notification for which we intend to mange configuration.\n\nUsage:\n  siren receiver [command]\n\nAliases:\n  receiver, receivers\n\nAvailable Commands:\n  create      Create a new receiver\n  delete      Delete a receiver details\n  edit        Edit a receiver\n  list        List receivers\n  send        Send a receiver notification\n  view        View a receiver details\n\nFlags:\n  -h, --help   help for receiver\n\nUse "siren receiver [command] --help" for more information about a command.\n')),(0,i.kt)("h3",{id:"permissions-and-auth-settings-for-slack-receivers"},"Permissions and Auth Settings for Slack Receivers"),(0,i.kt)("p",null,"In order to send Slack notifications via Siren Apis, you need to create its receiver. This Slack app then must be\ninstalled in the required workspaces and added to the required channels."),(0,i.kt)("p",null,"Siren helps with the installation flow by automating the exchanging code for access token\nflow. ",(0,i.kt)("a",{parentName:"p",href:"https://api.slack.com/legacy/oauth#authenticating-users-with-oauth__the-oauth-flow"},"Reference"),"."),(0,i.kt)("p",null,"Here is the list of actions one need to take to attach a Slack app to Siren."),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"Create a Slack app with these permissions. Visit ",(0,i.kt)("a",{parentName:"li",href:"https://api.slack.com/apps"},"this"),". If you already have an app, make\nsure permissions mentioned below are there."),(0,i.kt)("li",{parentName:"ol"},"Configure these permissions in the app:",(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-text"}," channels:read\n chat:write\n groups:read\n im:read\n team:read\n users:read\n users:read.email\n"))),(0,i.kt)("li",{parentName:"ol"},"Enable Distribution"),(0,i.kt)("li",{parentName:"ol"},"Setup a redirection server. You can use localhost as well. This must be a https server. Slack will call this server\nonce we install the app in any workspace."),(0,i.kt)("li",{parentName:"ol"},"Install your app to a workspace. Visit ",(0,i.kt)("inlineCode",{parentName:"li"},"Manage Distribution")," section on the App Dashboard. Click the ",(0,i.kt)("inlineCode",{parentName:"li"},"Add to Slack"),"\nButton."),(0,i.kt)("li",{parentName:"ol"},"This will prompt you to the OAuth Consent screen. Make sure you have selected the correct Slack Workspace by\nverifying the dropdown in the top-right corner. Click Allow."),(0,i.kt)("li",{parentName:"ol"},"Copy the ",(0,i.kt)("inlineCode",{parentName:"li"},"code")," that you received from Slack redirection URL query params and use this inside create receiver\npayload.")))}d.isMDXComponent=!0}}]);