"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[268],{3905:function(e,t,n){n.d(t,{Zo:function(){return u},kt:function(){return d}});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},l=Object.keys(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=a.createContext({}),p=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=p(e.components);return a.createElement(s.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,l=e.originalType,s=e.parentName,u=o(e,["components","mdxType","originalType","parentName"]),m=p(n),d=r,f=m["".concat(s,".").concat(d)]||m[d]||c[d]||l;return n?a.createElement(f,i(i({ref:t},u),{},{components:n})):a.createElement(f,i({ref:t},u))}));function d(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var l=n.length,i=new Array(l);i[0]=m;var o={};for(var s in t)hasOwnProperty.call(t,s)&&(o[s]=t[s]);o.originalType=e,o.mdxType="string"==typeof e?e:r,i[1]=o;for(var p=2;p<l;p++)i[p]=n[p];return a.createElement.apply(null,i)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},3773:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return o},contentTitle:function(){return s},metadata:function(){return p},toc:function(){return u},default:function(){return m}});var a=n(7462),r=n(3366),l=(n(7294),n(3905)),i=["components"],o={},s="Bulk Rule management",p={unversionedId:"guides/bulk_rules",id:"guides/bulk_rules",isDocsHomePage:!1,title:"Bulk Rule management",description:"For org wide use cases, teams end up managing a lot of rules, often manually.",source:"@site/docs/guides/bulk_rules.md",sourceDirName:"guides",slug:"/guides/bulk_rules",permalink:"/siren/docs/guides/bulk_rules",editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/guides/bulk_rules.md",tags:[],version:"current",lastUpdatedBy:"Abhishek Sah",lastUpdatedAt:1641901609,formattedLastUpdatedAt:"1/11/2022",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Alert History Subscription",permalink:"/siren/docs/guides/alert_history"},next:{title:"Monitoring",permalink:"/siren/docs/guides/monitoring"}},u=[],c={toc:u};function m(e){var t=e.components,n=(0,r.Z)(e,i);return(0,l.kt)("wrapper",(0,a.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h1",{id:"bulk-rule-management"},"Bulk Rule management"),(0,l.kt)("p",null,"For org wide use cases, teams end up managing a lot of rules, often manually."),(0,l.kt)("p",null,"Siren CLI can be used with some Gitops setup to automate the rule creation, rule update, template update. By putting all\nthe rules and templates YAML files in a version controlled repository, and uploading them using CI Jobs, you get speed\nin managing hundreds and thousands of rules in a reliable and predictable manner."),(0,l.kt)("p",null,"The benefits that one gets via this is:"),(0,l.kt)("ol",null,(0,l.kt)("li",{parentName:"ol"},"Predictable state of alerts after each CI job run"),(0,l.kt)("li",{parentName:"ol"},"Easy to rollback if something goes wrong"),(0,l.kt)("li",{parentName:"ol"},"Version controlled alerting state, democratizing alert setup, removing dependency from a central team")),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Example setup")),(0,l.kt)("ol",null,(0,l.kt)("li",{parentName:"ol"},(0,l.kt)("p",{parentName:"li"},"Create a github repo, let's call it ",(0,l.kt)("inlineCode",{parentName:"p"},"rules"),".")),(0,l.kt)("li",{parentName:"ol"},(0,l.kt)("p",{parentName:"li"},"Let's create a directory inside it and call it ",(0,l.kt)("inlineCode",{parentName:"p"},"templates"),". This is where people will put the YAML files of\nTemplates.")),(0,l.kt)("li",{parentName:"ol"},(0,l.kt)("p",{parentName:"li"},"Let's create a template names ",(0,l.kt)("inlineCode",{parentName:"p"},"cpu.yaml")," and add the below content"),(0,l.kt)("pre",{parentName:"li"},(0,l.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: v2\ntype: template\nname: CPU\nbody:\n  - alert: CPUWarning\n    expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > [[.warning]]\n    for: "[[.for]]"\n    labels:\n      severity: WARNING\n    annotations:\n      description: CPU has been above [[.warning]] for last [[.for]] {{ $labels.host }}\n  - alert: CPUCritical\n    expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > [[.critical]]\n    for: "[[.for]]"\n    labels:\n      severity: CRITICAL\n    annotations:\n      description: CPU has been above [[.critical]] for last [[.for]] {{ $labels.host }}\nvariables:\n  - name: for\n    type: string\n    default: 10m\n    description: For eg 5m, 2h; Golang duration format\n  - name: warning\n    type: int\n    default: 80\n  - name: critical\n    type: int\n    default: 90\ntags:\n  - systems\n'))),(0,l.kt)("li",{parentName:"ol"},(0,l.kt)("p",{parentName:"li"},"Let's define a shell script which iterates over all files inside ",(0,l.kt)("inlineCode",{parentName:"p"},"templates/")," directory on github to upload templates\nto Siren."),(0,l.kt)("pre",{parentName:"li"},(0,l.kt)("code",{parentName:"pre",className:"language-shell"},'#!/bin/bash\nset -e\necho "------------------------------------------------------------"\necho "Uploading templates: $DIR"\necho "------------------------------------------------------------"\n\nfor FILE in templates/*; do\n  eval ./siren template upload $FILE\n  echo $\'\\n\'\ndone\n\n'))),(0,l.kt)("li",{parentName:"ol"},(0,l.kt)("p",{parentName:"li"},"Now as the last step we need to run this script using github action. Here we are pulling siren image and using\nthe ",(0,l.kt)("inlineCode",{parentName:"p"},"upload")," command to upload the templates to Siren Web service, denoted by ",(0,l.kt)("inlineCode",{parentName:"p"},"SIREN_SERVICE_HOST")," environment\nvariable. An example is:"))),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-yaml"},"// to be filled later\n")),(0,l.kt)("ol",{start:6},(0,l.kt)("li",{parentName:"ol"},(0,l.kt)("p",{parentName:"li"},"For rules, create a directory called ",(0,l.kt)("inlineCode",{parentName:"p"},"rules")," beside ",(0,l.kt)("inlineCode",{parentName:"p"},"templates")," and start define an example rule as given below."),(0,l.kt)("pre",{parentName:"li"},(0,l.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: v2\ntype: rule\nnamespace: demo\nprovider: production-cortex\nproviderNamespace: odpf\nrules:\n  TestGroup:\n    template: CPU\n    status: enabled\n    variables:\n      - name: for\n        value: 15m\n      - name: warning\n        value: 185\n      - name: critical\n        value: 195\n"))),(0,l.kt)("li",{parentName:"ol"},(0,l.kt)("p",{parentName:"li"},"We can upload the files inside ",(0,l.kt)("inlineCode",{parentName:"p"},"rules")," directory iteratively. Here is an example script. This can be called in github\nCI action."),(0,l.kt)("pre",{parentName:"li"},(0,l.kt)("code",{parentName:"pre",className:"language-shell"},'#!/bin/bash\nset -e\necho "------------------------------------------------------------"\necho "Uploading rules: $DIR"\necho "------------------------------------------------------------"\n\nfor FILE in rules/*; do\n  eval ./siren rule upload $FILE\n  echo $\'\\n\'\ndone\n')))))}m.isMDXComponent=!0}}]);