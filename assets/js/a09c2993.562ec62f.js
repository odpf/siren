"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[128],{3905:function(e,t,n){n.d(t,{Zo:function(){return u},kt:function(){return f}});var r=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=r.createContext({}),c=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},u=function(e){var t=c(e.components);return r.createElement(l.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,a=e.originalType,l=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),d=c(n),f=i,m=d["".concat(l,".").concat(f)]||d[f]||p[f]||a;return n?r.createElement(m,o(o({ref:t},u),{},{components:n})):r.createElement(m,o({ref:t},u))}));function f(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var a=n.length,o=new Array(a);o[0]=d;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:i,o[1]=s;for(var c=2;c<a;c++)o[c]=n[c];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},8495:function(e,t,n){n.r(t),n.d(t,{assets:function(){return l},contentTitle:function(){return o},default:function(){return p},frontMatter:function(){return a},metadata:function(){return s},toc:function(){return c}});var r=n(3117),i=(n(7294),n(3905));const a={},o="Introduction",s={unversionedId:"introduction",id:"introduction",title:"Introduction",description:"Siren orchestrates alerting rules of your applications using a monitoring and alerting provider e.g. Cortex metrics and sending notifications in a simple DIY configuration. With Siren, you can define templates (using go templates standard), create/edit/enable/disable alerting rules on demand, and sending out notifications. It also gives flexibility to manage bulk of rules via YAML files. Siren can be integrated with any client such as CI/CD pipelines, Self-Serve UI, microservices etc.",source:"@site/docs/introduction.md",sourceDirName:".",slug:"/introduction",permalink:"/siren/docs/introduction",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/introduction.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",next:{title:"Introduction",permalink:"/siren/docs/tour/introduction"}},l={},c=[{value:"Key Features",id:"key-features",level:2},{value:"Usage",id:"usage",level:2}],u={toc:c};function p(e){let{components:t,...a}=e;return(0,i.kt)("wrapper",(0,r.Z)({},u,a,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"introduction"},"Introduction"),(0,i.kt)("p",null,"Siren orchestrates alerting rules of your applications using a monitoring and alerting provider e.g. ",(0,i.kt)("a",{parentName:"p",href:"https://cortexmetrics.io/"},"Cortex metrics")," and sending notifications in a simple DIY configuration. With Siren, you can define templates (using go templates standard), create/edit/enable/disable alerting rules on demand, and sending out notifications. It also gives flexibility to manage bulk of rules via YAML files. Siren can be integrated with any client such as CI/CD pipelines, Self-Serve UI, microservices etc."),(0,i.kt)("p",null,(0,i.kt)("img",{alt:"Siren Overview",src:n(4470).Z,width:"590",height:"336"})),(0,i.kt)("h2",{id:"key-features"},"Key Features"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"Rule Templates:")," Siren provides a way to define templates over alerting rule which can be reused to create multiple instances of the same rule with configurable thresholds."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"Multi-tenancy:")," Rules created with Siren are by default multi-tenancy aware."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"DIY Interface:")," Siren can be used to easily create/edit alerting rules. It also provides soft delete(disable) so that you can preserve thresholds in case you need to reuse the same alert."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"Managing bulk rules:")," Siren enables users to manage bulk alerting rules using YAML files in specified format with simple CLI."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"Receivers:")," Siren can be used to send out notifications to several channels (e.g. slack, pagerduty, email etc)."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"Subscriptions")," Siren can be used to subscribe to notifications (with desired matching conditions) via the channel of your choice."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"Alert History:")," Siren can store alerts triggered by monitoring & alerting provider e.g. Cortex Alertmanager, which can be used for audit purposes.")),(0,i.kt)("h2",{id:"usage"},"Usage"),(0,i.kt)("p",null,"Explore the following resources to get started with Siren:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/siren/docs/tour/introduction"},"Tour")," allows you to explore Siren features quickly."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/siren/docs/concepts/overview"},"Concepts")," describes all important Siren concepts including system architecture."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/siren/docs/guides/overview"},"Guides")," provides guidance on usage."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/siren/docs/reference/server_configuration"},"Reference")," contains the details about configurations and other aspects of Siren."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/siren/docs/contribute/contribution"},"Contribute")," contains resources for anyone who wants to contribute to Siren.")))}p.isMDXComponent=!0},4470:function(e,t,n){t.Z=n.p+"assets/images/overview-640dcd08ea55323369bae78b6055feec.svg"}}]);