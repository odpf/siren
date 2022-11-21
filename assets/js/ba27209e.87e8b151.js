"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[55],{3905:(e,r,t)=>{t.d(r,{Zo:()=>u,kt:()=>f});var n=t(7294);function o(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function i(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function s(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?i(Object(t),!0).forEach((function(r){o(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function a(e,r){if(null==e)return{};var t,n,o=function(e,r){if(null==e)return{};var t,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)t=i[n],r.indexOf(t)>=0||(o[t]=e[t]);return o}(e,r);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)t=i[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var l=n.createContext({}),c=function(e){var r=n.useContext(l),t=r;return e&&(t="function"==typeof e?e(r):s(s({},r),e)),t},u=function(e){var r=c(e.components);return n.createElement(l.Provider,{value:r},e.children)},d={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},p=n.forwardRef((function(e,r){var t=e.components,o=e.mdxType,i=e.originalType,l=e.parentName,u=a(e,["components","mdxType","originalType","parentName"]),p=c(t),f=o,h=p["".concat(l,".").concat(f)]||p[f]||d[f]||i;return t?n.createElement(h,s(s({ref:r},u),{},{components:t})):n.createElement(h,s({ref:r},u))}));function f(e,r){var t=arguments,o=r&&r.mdxType;if("string"==typeof e||o){var i=t.length,s=new Array(i);s[0]=p;var a={};for(var l in r)hasOwnProperty.call(r,l)&&(a[l]=r[l]);a.originalType=e,a.mdxType="string"==typeof e?e:o,s[1]=a;for(var c=2;c<i;c++)s[c]=t[c];return n.createElement.apply(null,s)}return n.createElement.apply(null,t)}p.displayName="MDXCreateElement"},7947:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>l,contentTitle:()=>s,default:()=>d,frontMatter:()=>i,metadata:()=>a,toc:()=>c});var n=t(3117),o=(t(7294),t(3905));const i={},s="Workers",a={unversionedId:"guides/workers",id:"guides/workers",title:"Workers",description:"Siren has a notification features that utilizes queue to publish notification messages. More concept about notification could be found here. The architecture requires a detached worker running asynchronously and polling queue periodically to dequeue notification messages and publish them. By default, Siren server run this asynchronous worker inside it. However it is also possible to run the worker as a different process. Currently there are two possible workers to run",source:"@site/docs/guides/workers.md",sourceDirName:"guides",slug:"/guides/workers",permalink:"/siren/docs/guides/workers",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/guides/workers.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Notification",permalink:"/siren/docs/guides/notification"},next:{title:"Job",permalink:"/siren/docs/guides/job"}},l={},c=[{value:"Running Workers as a Different Process",id:"running-workers-as-a-different-process",level:2}],u={toc:c};function d(e){let{components:r,...t}=e;return(0,o.kt)("wrapper",(0,n.Z)({},u,t,{components:r,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"workers"},"Workers"),(0,o.kt)("p",null,"Siren has a notification features that utilizes queue to publish notification messages. More concept about notification could be found ",(0,o.kt)("a",{parentName:"p",href:"/siren/docs/concepts/notification"},"here"),". The architecture requires a detached worker running asynchronously and polling queue periodically to dequeue notification messages and publish them. By default, Siren server run this asynchronous worker inside it. However it is also possible to run the worker as a different process. Currently there are two possible workers to run"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("strong",{parentName:"li"},"Notification message handler:")," this worker periodically poll and dequeue messages from queue, process the messages, and then publish notification messages to the specified receivers. If there is a failure, this handler enqueues the failed messages to the dlq."),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("strong",{parentName:"li"},"Notification dlq handler:")," this worker periodically poll and dequeue messages from dead-letter-queue, process the messages, and then publish notification messages to the specified receivers. If there is a failure, this handler enqueues the failed messages back to the dlq.")),(0,o.kt)("h2",{id:"running-workers-as-a-different-process"},"Running Workers as a Different Process"),(0,o.kt)("p",null,"Siren has a command to start workers. Workers use the same config like server does."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"Command to start a siren worker.\n\nUsage\n  siren worker start <command> [flags]\n\nCore commands\n  notification_dlq_handler    A notification dlq handler\n  notification_handler        A notification handler\n\nInherited flags\n  --help   Show help for command\n\nExamples\n  $ siren worker start notification_handler\n  $ siren server start notification_handler -c ./config.yaml\n")),(0,o.kt)("p",null,"Starting up a worker could be done by executing."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren worker start notification_handler\n")))}d.isMDXComponent=!0}}]);