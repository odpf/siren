"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[75],{3905:function(e,n,t){t.d(n,{Zo:function(){return p},kt:function(){return m}});var r=t(7294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function a(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function s(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?a(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function o(e,n){if(null==e)return{};var t,r,i=function(e,n){if(null==e)return{};var t,r,i={},a=Object.keys(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var c=r.createContext({}),l=function(e){var n=r.useContext(c),t=n;return e&&(t="function"==typeof e?e(n):s(s({},n),e)),t},p=function(e){var n=l(e.components);return r.createElement(c.Provider,{value:n},e.children)},u={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},d=r.forwardRef((function(e,n){var t=e.components,i=e.mdxType,a=e.originalType,c=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),d=l(t),m=i,g=d["".concat(c,".").concat(m)]||d[m]||u[m]||a;return t?r.createElement(g,s(s({ref:n},p),{},{components:t})):r.createElement(g,s({ref:n},p))}));function m(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var a=t.length,s=new Array(a);s[0]=d;var o={};for(var c in n)hasOwnProperty.call(n,c)&&(o[c]=n[c]);o.originalType=e,o.mdxType="string"==typeof e?e:i,s[1]=o;for(var l=2;l<a;l++)s[l]=t[l];return r.createElement.apply(null,s)}return r.createElement.apply(null,t)}d.displayName="MDXCreateElement"},2001:function(e,n,t){t.r(n),t.d(n,{frontMatter:function(){return o},contentTitle:function(){return c},metadata:function(){return l},toc:function(){return p},default:function(){return d}});var r=t(7462),i=t(3366),a=(t(7294),t(3905)),s=["components"],o={},c="Subscriptions",l={unversionedId:"guides/subscriptions",id:"guides/subscriptions",isDocsHomePage:!1,title:"Subscriptions",description:"Siren lets you subscribe to the rules when they are triggered. You can define custom matching conditions and use",source:"@site/docs/guides/subscriptions.md",sourceDirName:"guides",slug:"/guides/subscriptions",permalink:"/siren/docs/guides/subscriptions",editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/guides/subscriptions.md",tags:[],version:"current",lastUpdatedBy:"Abhishek Sah",lastUpdatedAt:1646811080,formattedLastUpdatedAt:"3/9/2022",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Receivers",permalink:"/siren/docs/guides/receivers"},next:{title:"Rules",permalink:"/siren/docs/guides/rules"}},p=[{value:"API Interface",id:"api-interface",children:[{value:"Create a subscription",id:"create-a-subscription",children:[]},{value:"Update a subscription",id:"update-a-subscription",children:[]},{value:"Get all subscriptions",id:"get-all-subscriptions",children:[]},{value:"Get a subscriptions",id:"get-a-subscriptions",children:[]},{value:"Delete subscriptions",id:"delete-subscriptions",children:[]}]}],u={toc:p};function d(e){var n=e.components,t=(0,i.Z)(e,s);return(0,a.kt)("wrapper",(0,r.Z)({},u,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"subscriptions"},"Subscriptions"),(0,a.kt)("p",null,"Siren lets you subscribe to the rules when they are triggered. You can define custom matching conditions and use\n",(0,a.kt)("a",{parentName:"p",href:"/siren/docs/guides/receivers"},"receivers")," to describe which medium you want to use for getting the notifications when those rules are\ntriggered. Siren syncs this configuration in the respective monitoring provider."),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"Example Subscription:")),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-json"},'{\n  "id": "385",\n  "urn": "siren-dev-prod-critical",\n  "namespace": "10",\n  "receivers": [\n    {\n      "id": "2"\n    },\n    {\n      "id": "1",\n      "configuration": {\n        "channel_name": "siren-dev-critical"\n      }\n    }\n  ],\n  "match": {\n    "environment": "production",\n    "severity": "CRITICAL"\n  },\n  "created_at": "2021-12-10T10:38:22.364353Z",\n  "updated_at": "2021-12-10T10:38:22.364353Z"\n}\n')),(0,a.kt)("p",null,"The above means whenever any alert which has labels matching the ",(0,a.kt)("inlineCode",{parentName:"p"},"match"),"viz:\n",(0,a.kt)("inlineCode",{parentName:"p"},'"environment": "production", "severity": "CRITICAL"'),", send this alert to two medium defined by receivers with id: ",(0,a.kt)("inlineCode",{parentName:"p"},"2"),"\nand ",(0,a.kt)("inlineCode",{parentName:"p"},"1"),". Assuming the receivers id ",(0,a.kt)("inlineCode",{parentName:"p"},"2")," to be of Pagerduty type, a PD call will be invoked and assuming the receiver with\nid ",(0,a.kt)("inlineCode",{parentName:"p"},"1")," to be slack type, a message will be sent to the channel #siren-dev-critical."),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"Upstream sync example")),(0,a.kt)("p",null,"The logical equivalence of this routing configuration is put in the respective monitoring provider by Siren. For ex: if\nthe provider is cortex, an alertmanager configuration will be created depicting the above routing logic."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},'templates:\n  - "helper.tmpl"\nglobal:\n  pagerduty_url: https://events.pagerduty.com/v2/enqueue\n  resolve_timeout: 5m\n  slack_api_url: https://slack.com/api/chat.postMessage\nreceivers:\n  - name: default\n  - name: slack_siren-dev-prod-critical_receiverId_1_idx_0\n    slack_configs:\n      - channel: "siren-dev-critical"\n        http_config:\n          bearer_token: "secret-taken-from-receiver-config"\n        icon_emoji: ":eagle:"\n        link_names: false\n        send_resolved: true\n        color: \'{{ template "slack.color" . }}\'\n        title: ""\n        pretext: \'{{template "slack.pretext" . }}\'\n        text: \'{{ template "slack.body" . }}\'\n        actions:\n          - type: button\n            text: "Runbook :books:"\n            url: \'{{template "slack.runbook" . }}\'\n          - type: button\n            text: "Dashboard :bar_chart:"\n            url: \'{{template "slack.dashboard" . }}\'\n  - name: pagerduty_siren-dev-prod-critical_receiverId_2_idx_1\n    pagerduty_configs:\n      - service_key: "secret-taken-from-receiver-config"\nroute:\n  group_by:\n    - alertname\n    - severity\n    - owner\n    - service_name\n    - time_stamp\n    - identifier\n  group_wait: 30s\n  group_interval: 5m\n  repeat_interval: 4h\n  receiver: default\n  routes:\n    - receiver: slack_siren-dev-prod-critical_receiverId_1_idx_0\n      match:\n        environment: production\n        severity: CRITICAL\n      continue: true\n    - receiver: pagerduty_siren-dev-prod-critical_receiverId_2_idx_1\n      match:\n        environment: production\n        severity: CRITICAL\n      continue: true\n\n')),(0,a.kt)("p",null,"As you can see, Siren dynamically defined two receivers: ",(0,a.kt)("inlineCode",{parentName:"p"},"slack_siren-dev-prod-critical_receiverId_1_idx_0"),"\nand ",(0,a.kt)("inlineCode",{parentName:"p"},"pagerduty_siren-dev-prod-critical_receiverId_2_idx_1")," and used them in the routing tree as per the match\nconditions."),(0,a.kt)("p",null,"This alertmanager config is for the tenant defined by namespace with id ",(0,a.kt)("inlineCode",{parentName:"p"},"10")," as mentioned in the example. This is an\nexample config, the actual config will contain all subscriptions that belong to namespace with id ",(0,a.kt)("inlineCode",{parentName:"p"},"10")),(0,a.kt)("h2",{id:"api-interface"},"API Interface"),(0,a.kt)("h3",{id:"create-a-subscription"},"Create a subscription"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},'POST /v1beta1/subscriptions HTTP/1.1\nHost: localhost:3000\nContent-Type: application/json\nContent-Length: 363\n\n{\n    "urn": "siren-dev-prod-critical",\n    "receivers": [\n        {\n            "id": "1",\n            "configuration": {\n                "channel_name": "siren-dev-critical"\n            }\n        },\n        {\n            "id": "2"\n        }\n    ],\n    "match": {\n        "severity": "CRITICAL",\n        "environment": "production"\n    },\n    "namespace": "10"\n}\n')),(0,a.kt)("h3",{id:"update-a-subscription"},"Update a subscription"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},'POST /v1beta1/subscriptions HTTP/1.1\nHost: localhost:3000\nContent-Type: application/json\nContent-Length: 392\n\n{\n    "urn": "siren-dev-prod-critical",\n    "receivers": [\n        {\n            "id": "1",\n            "configuration": {\n                "channel_name": "siren-dev-critical"\n            }\n        },\n        {\n            "id": "2"\n        }\n    ],\n    "match": {\n        "severity": "CRITICAL",\n        "environment": "production",\n        "team": "siren-dev"\n    },\n    "namespace": "10"\n}\n')),(0,a.kt)("h3",{id:"get-all-subscriptions"},"Get all subscriptions"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},"GET /v1beta1/subscriptions HTTP/1.1\nHost: localhost:3000\n")),(0,a.kt)("h3",{id:"get-a-subscriptions"},"Get a subscriptions"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},"GET /v1beta1/subscriptions/10 HTTP/1.1\nHost: localhost:3000\n")),(0,a.kt)("h3",{id:"delete-subscriptions"},"Delete subscriptions"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},"DELETE /v1beta1/subscriptions/10 HTTP/1.1\nHost: localhost:3000\n")))}d.isMDXComponent=!0}}]);