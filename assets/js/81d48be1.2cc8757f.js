"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[576],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>m});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=r.createContext({}),c=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},u=function(e){var t=c(e.components);return r.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,s=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),d=c(n),m=o,f=d["".concat(s,".").concat(m)]||d[m]||p[m]||i;return n?r.createElement(f,a(a({ref:t},u),{},{components:n})):r.createElement(f,a({ref:t},u))}));function m(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,a=new Array(i);a[0]=d;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:o,a[1]=l;for(var c=2;c<i;c++)a[c]=n[c];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},4949:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>a,default:()=>p,frontMatter:()=>i,metadata:()=>l,toc:()=>c});var r=n(3117),o=(n(7294),n(3905));const i={},a="Introduction",l={unversionedId:"tour/introduction",id:"tour/introduction",title:"Introduction",description:"This tour introduces you to Siren. Along the way you will learn how to manage alerting rules, notification receivers, and subscribing to alert notifications.",source:"@site/docs/tour/introduction.md",sourceDirName:"tour",slug:"/tour/introduction",permalink:"/siren/docs/tour/introduction",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/tour/introduction.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Installation",permalink:"/siren/docs/installation"},next:{title:"Setup Server",permalink:"/siren/docs/tour/setup_server"}},s={},c=[{value:"Pre-requisites",id:"pre-requisites",level:2},{value:"Help",id:"help",level:2},{value:"Background for this tutorial",id:"background-for-this-tutorial",level:2}],u={toc:c};function p(e){let{components:t,...n}=e;return(0,o.kt)("wrapper",(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"introduction"},"Introduction"),(0,o.kt)("p",null,"This tour introduces you to Siren. Along the way you will learn how to manage alerting rules, notification receivers, and subscribing to alert notifications."),(0,o.kt)("h2",{id:"pre-requisites"},"Pre-requisites"),(0,o.kt)("p",null,"This tour requires you to have Siren CLI tool installed on your local machine. You can run ",(0,o.kt)("inlineCode",{parentName:"p"},"siren version")," to verify the installation. Please follow ",(0,o.kt)("a",{parentName:"p",href:"/siren/docs/installation"},"installation")," and ",(0,o.kt)("a",{parentName:"p",href:"/siren/docs/reference/server_configuration"},"configuration")," guides if you do not have it installed already."),(0,o.kt)("p",null,"Siren client CLI talks to Siren server to configure and fetch rules, subscriptions, and notifications. Please make sure you also have a Siren server running. You can also run server locally with ",(0,o.kt)("inlineCode",{parentName:"p"},"siren server start")," command. For more details check the ",(0,o.kt)("a",{parentName:"p",href:"/siren/docs/guides/deployment"},"deployment")," guide."),(0,o.kt)("h2",{id:"help"},"Help"),(0,o.kt)("p",null,"At any time you can run the following commands."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"# Check the installed version for Siren cli tool\n$ siren version\n\n# See the help for a command\n$ siren --help\n")),(0,o.kt)("p",null,"The list of all available commands are as follows:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"CORE COMMANDS\n  alert           Manage alerts\n  namespace       Manage namespaces\n  provider        Manage providers\n  receiver        Manage receivers\n  rule            Manage rules\n  subscription    Manage subscriptions\n  template        Manage templates\n\nADDITIONAL COMMANDS\n  completion      Generate shell completion scripts\n  config          Manage siren CLI configuration\n  help            Help about any command\n  job             Manage siren jobs\n  environment     List of supported environment variables\n  reference       Comprehensive reference of all commands\n  server          Run siren server\n  worker          Start or manage Siren's workers\n")),(0,o.kt)("p",null,"Help command can also be run on any sub command with syntax ",(0,o.kt)("inlineCode",{parentName:"p"},"siren <command> <subcommand> --help"),". Here is an example for the same."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"$ siren rule --help\n")),(0,o.kt)("p",null,"Check the reference for Siren cli commands."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"$ siren reference\n")),(0,o.kt)("h2",{id:"background-for-this-tutorial"},"Background for this tutorial"),(0,o.kt)("p",null,"This tour introduces you to two different scenarios"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("a",{parentName:"li",href:"/siren/docs/tour/1sending_notifications_overview"},"Sending on-demand notification to a receiver"),(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"Register a receiver"),(0,o.kt)("li",{parentName:"ul"},"Send notification to the receiver"))),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("a",{parentName:"li",href:"/siren/docs/tour/2alerting_rules_subscriptions_overview"},"Setting up alerting rules and subscribing to the alerts"),(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"Register a CortexMetrics provider"),(0,o.kt)("li",{parentName:"ul"},"Create a new namespace"),(0,o.kt)("li",{parentName:"ul"},"Register a receiver that we want to send the notification to"),(0,o.kt)("li",{parentName:"ul"},"Create a subscription to define the routing so alert notification will be routed to the registered receivers")))),(0,o.kt)("p",null,"The tour takes approximately 20 minutes to complete."))}p.isMDXComponent=!0}}]);