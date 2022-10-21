"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[403],{3905:function(e,t,n){n.d(t,{Zo:function(){return l},kt:function(){return d}});var r=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var u=r.createContext({}),c=function(e){var t=r.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},l=function(e){var t=c(e.components);return r.createElement(u.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,o=e.originalType,u=e.parentName,l=s(e,["components","mdxType","originalType","parentName"]),m=c(n),d=i,f=m["".concat(u,".").concat(d)]||m[d]||p[d]||o;return n?r.createElement(f,a(a({ref:t},l),{},{components:n})):r.createElement(f,a({ref:t},l))}));function d(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=n.length,a=new Array(o);a[0]=m;var s={};for(var u in t)hasOwnProperty.call(t,u)&&(s[u]=t[u]);s.originalType=e,s.mdxType="string"==typeof e?e:i,a[1]=s;for(var c=2;c<o;c++)a[c]=n[c];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},5057:function(e,t,n){n.r(t),n.d(t,{assets:function(){return u},contentTitle:function(){return a},default:function(){return p},frontMatter:function(){return o},metadata:function(){return s},toc:function(){return c}});var r=n(3117),i=(n(7294),n(3905));const o={},a="Contribution Process",s={unversionedId:"contribute/contribution",id:"contribute/contribution",title:"Contribution Process",description:"The following is a set of guidelines for contributing to Siren. These are mostly guidelines, not rules. Use your best",source:"@site/docs/contribute/contribution.md",sourceDirName:"contribute",slug:"/contribute/contribution",permalink:"/siren/docs/contribute/contribution",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/contribute/contribution.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Deployment",permalink:"/siren/docs/guides/deployment"},next:{title:"Add a New Receiver Plugin",permalink:"/siren/docs/contribute/receiver"}},u={},c=[{value:"How can I contribute?",id:"how-can-i-contribute",level:2},{value:"Becoming a maintainer",id:"becoming-a-maintainer",level:2},{value:"Guidelines",id:"guidelines",level:2}],l={toc:c};function p(e){let{components:t,...n}=e;return(0,i.kt)("wrapper",(0,r.Z)({},l,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"contribution-process"},"Contribution Process"),(0,i.kt)("p",null,"The following is a set of guidelines for contributing to Siren. These are mostly guidelines, not rules. Use your best\njudgment, and feel free to propose changes to this document in a pull request. Here are some important resources:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"../concepts/overview"},"Concepts")," section will explain you about Siren architecture,"),(0,i.kt)("li",{parentName:"ul"},"Our ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/odpf/siren#readme"},"roadmap")," is the 10k foot view of where we're going, and"),(0,i.kt)("li",{parentName:"ul"},"Github ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/odpf/siren/issues"},"issues")," track the ongoing and reported issues.")),(0,i.kt)("h2",{id:"how-can-i-contribute"},"How can I contribute?"),(0,i.kt)("p",null,"We use RFCS and GitHub issues to communicate ideas."),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"You can report a bug or suggest a feature enhancement or can just ask questions. Reach out on Github discussions for\nthis purpose."),(0,i.kt)("li",{parentName:"ul"},"You are also welcome to improve error reporting and logging and improve code quality."),(0,i.kt)("li",{parentName:"ul"},"You can help with documenting new features or improve an existing documentation."),(0,i.kt)("li",{parentName:"ul"},"You can also review and accept other contributions if you are a maintainer.")),(0,i.kt)("p",null,"Please submit a PR to the main branch of the Siren repository once you are ready to submit your contribution. Code\nsubmission to Siren (including submission from project maintainers) require review and approval from maintainers or code\nowners. PRs that are submitted need to pass the build. Once build is passed community members will help to review the\npull request."),(0,i.kt)("h2",{id:"becoming-a-maintainer"},"Becoming a maintainer"),(0,i.kt)("p",null,"We are always interested in adding new maintainers. What we look for is a series of contributions, good taste, and an\nongoing interest in the project."),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"maintainers will have write access to the Siren repository."),(0,i.kt)("li",{parentName:"ul"},"There is no strict protocol for becoming a maintainer. Candidates for new maintainers are typically people that are\nactive contributors and community members."),(0,i.kt)("li",{parentName:"ul"},"Candidates for new maintainers can also be suggested by current maintainers."),(0,i.kt)("li",{parentName:"ul"},"If you would like to become a maintainer, you should start contributing to Siren in any of the ways mentioned. You\nmight also want to talk to other maintainers and ask for their advice and guidance.")),(0,i.kt)("h2",{id:"guidelines"},"Guidelines"),(0,i.kt)("p",null,"Please follow these practices for you change to get merged fast and smoothly:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Contributions can only be accepted if they contain appropriate testing (Unit and Integration Tests)."),(0,i.kt)("li",{parentName:"ul"},"If you are introducing a completely new feature or making any major changes in an existing one, we recommend to start\nwith an RFC and get consensus on the basic design first."),(0,i.kt)("li",{parentName:"ul"},"Make sure your local build is running with all the tests and checkstyle passing."),(0,i.kt)("li",{parentName:"ul"},"If your change is related to user-facing protocols / configurations, you need to make the corresponding change in the\ndocumentation as well."),(0,i.kt)("li",{parentName:"ul"},"Docs live in the code repo under ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/odpf/siren/tree/main/docs"},(0,i.kt)("inlineCode",{parentName:"a"},"docs"))," so that changes to that can be\ndone in the same PR as changes to the code.")))}p.isMDXComponent=!0}}]);