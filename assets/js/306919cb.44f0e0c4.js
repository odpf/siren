"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[825],{3905:function(t,e,n){n.d(e,{Zo:function(){return d},kt:function(){return m}});var a=n(7294);function r(t,e,n){return e in t?Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}):t[e]=n,t}function l(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function i(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?l(Object(n),!0).forEach((function(e){r(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}function o(t,e){if(null==t)return{};var n,a,r=function(t,e){if(null==t)return{};var n,a,r={},l=Object.keys(t);for(a=0;a<l.length;a++)n=l[a],e.indexOf(n)>=0||(r[n]=t[n]);return r}(t,e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(t);for(a=0;a<l.length;a++)n=l[a],e.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(t,n)&&(r[n]=t[n])}return r}var u=a.createContext({}),p=function(t){var e=a.useContext(u),n=e;return t&&(n="function"==typeof t?t(e):i(i({},e),t)),n},d=function(t){var e=p(t.components);return a.createElement(u.Provider,{value:e},t.children)},g={inlineCode:"code",wrapper:function(t){var e=t.children;return a.createElement(a.Fragment,{},e)}},s=a.forwardRef((function(t,e){var n=t.components,r=t.mdxType,l=t.originalType,u=t.parentName,d=o(t,["components","mdxType","originalType","parentName"]),s=p(n),m=r,c=s["".concat(u,".").concat(m)]||s[m]||g[m]||l;return n?a.createElement(c,i(i({ref:e},d),{},{components:n})):a.createElement(c,i({ref:e},d))}));function m(t,e){var n=arguments,r=e&&e.mdxType;if("string"==typeof t||r){var l=n.length,i=new Array(l);i[0]=s;var o={};for(var u in e)hasOwnProperty.call(e,u)&&(o[u]=e[u]);o.originalType=t,o.mdxType="string"==typeof t?t:r,i[1]=o;for(var p=2;p<l;p++)i[p]=n[p];return a.createElement.apply(null,i)}return a.createElement.apply(null,n)}s.displayName="MDXCreateElement"},4480:function(t,e,n){n.r(e),n.d(e,{frontMatter:function(){return o},contentTitle:function(){return u},metadata:function(){return p},toc:function(){return d},default:function(){return s}});var a=n(7462),r=n(3366),l=(n(7294),n(3905)),i=["components"],o={},u="Configuration",p={unversionedId:"reference/configuration",id:"reference/configuration",isDocsHomePage:!1,title:"Configuration",description:"| Go struct                     | YAML path        | ENV var          | default   | Valid values                                                                                                     |",source:"@site/docs/reference/configuration.md",sourceDirName:"reference",slug:"/reference/configuration",permalink:"/siren/docs/reference/configuration",editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/reference/configuration.md",tags:[],version:"current",lastUpdatedBy:"Abhishek Sah",lastUpdatedAt:1649745604,formattedLastUpdatedAt:"4/12/2022",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Siren APIs",permalink:"/siren/docs/reference/api"}},d=[{value:"How to configure",id:"how-to-configure",children:[{value:"Using env variables",id:"using-env-variables",children:[]},{value:"Using a yaml file",id:"using-a-yaml-file",children:[]},{value:"Using a combinnation of both",id:"using-a-combinnation-of-both",children:[]}]}],g={toc:d};function s(t){var e=t.components,n=(0,r.Z)(t,i);return(0,l.kt)("wrapper",(0,a.Z)({},g,n,{components:e,mdxType:"MDXLayout"}),(0,l.kt)("h1",{id:"configuration"},"Configuration"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Go struct"),(0,l.kt)("th",{parentName:"tr",align:null},"YAML path"),(0,l.kt)("th",{parentName:"tr",align:null},"ENV var"),(0,l.kt)("th",{parentName:"tr",align:null},"default"),(0,l.kt)("th",{parentName:"tr",align:null},"Valid values"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.LogConfig.Level"),(0,l.kt)("td",{parentName:"tr",align:null},"log.level"),(0,l.kt)("td",{parentName:"tr",align:null},"LOG_LEVEL"),(0,l.kt)("td",{parentName:"tr",align:null},"info"),(0,l.kt)("td",{parentName:"tr",align:null},"debug,info,warn,error,dpanic,panic,fatal")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.Port"),(0,l.kt)("td",{parentName:"tr",align:null},"port"),(0,l.kt)("td",{parentName:"tr",align:null},"PORT"),(0,l.kt)("td",{parentName:"tr",align:null},"8080"),(0,l.kt)("td",{parentName:"tr",align:null},"0-65535")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.DBConfig.User"),(0,l.kt)("td",{parentName:"tr",align:null},"db.user"),(0,l.kt)("td",{parentName:"tr",align:null},"DB_USER"),(0,l.kt)("td",{parentName:"tr",align:null},"postgres"),(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("a",{parentName:"td",href:"https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS"},"PostgreSQL identifiers"))),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.DBConfig.Name"),(0,l.kt)("td",{parentName:"tr",align:null},"db.name"),(0,l.kt)("td",{parentName:"tr",align:null},"DB_NAME"),(0,l.kt)("td",{parentName:"tr",align:null},"postgres"),(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("a",{parentName:"td",href:"https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS"},"PostgreSQL identifiers"))),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.DBConfig.Port"),(0,l.kt)("td",{parentName:"tr",align:null},"db.port"),(0,l.kt)("td",{parentName:"tr",align:null},"DB_PORT"),(0,l.kt)("td",{parentName:"tr",align:null},"5432"),(0,l.kt)("td",{parentName:"tr",align:null},"0-65535")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.DBConfig.Password"),(0,l.kt)("td",{parentName:"tr",align:null},"db.password"),(0,l.kt)("td",{parentName:"tr",align:null},"DB_PASSWORD"),(0,l.kt)("td",{parentName:"tr",align:null}),(0,l.kt)("td",{parentName:"tr",align:null},"valid PostgreSQL password")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.DBConfig.SslMode"),(0,l.kt)("td",{parentName:"tr",align:null},"db.sslmode"),(0,l.kt)("td",{parentName:"tr",align:null},"DB_SSLMODE"),(0,l.kt)("td",{parentName:"tr",align:null},"disable"),(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("a",{parentName:"td",href:"https://www.postgresql.org/docs/9.1/libpq-ssl.html"},"libpq sslmode"))),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.DBConfig.LogLevel"),(0,l.kt)("td",{parentName:"tr",align:null},"db.log_level"),(0,l.kt)("td",{parentName:"tr",align:null},"DB_LOG_LEVEL"),(0,l.kt)("td",{parentName:"tr",align:null},"info"),(0,l.kt)("td",{parentName:"tr",align:null},"silent,error,warn,info")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.DBConfig.Host"),(0,l.kt)("td",{parentName:"tr",align:null},"db.host"),(0,l.kt)("td",{parentName:"tr",align:null},"DB_HOST"),(0,l.kt)("td",{parentName:"tr",align:null},"localhost"),(0,l.kt)("td",{parentName:"tr",align:null},"valid hostname name or IP address")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.NewRelicConfig.Enabled"),(0,l.kt)("td",{parentName:"tr",align:null},"newrelic.enabled"),(0,l.kt)("td",{parentName:"tr",align:null},"NEWRELIC_ENABLED"),(0,l.kt)("td",{parentName:"tr",align:null},"false"),(0,l.kt)("td",{parentName:"tr",align:null},"bool")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.NewRelicConfig.License"),(0,l.kt)("td",{parentName:"tr",align:null},"newrelic.license"),(0,l.kt)("td",{parentName:"tr",align:null},"NEWRELIC_LICENSE"),(0,l.kt)("td",{parentName:"tr",align:null}),(0,l.kt)("td",{parentName:"tr",align:null},"40 char NewRelic license key")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Config.NewRelicConfig.AppName"),(0,l.kt)("td",{parentName:"tr",align:null},"newrelic.appname"),(0,l.kt)("td",{parentName:"tr",align:null},"NEWRELIC_APPNAME"),(0,l.kt)("td",{parentName:"tr",align:null},"siren"),(0,l.kt)("td",{parentName:"tr",align:null},"string")))),(0,l.kt)("h2",{id:"how-to-configure"},"How to configure"),(0,l.kt)("p",null,"There are 3 ways to configure siren:"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},"Using env variables"),(0,l.kt)("li",{parentName:"ul"},"Using a yaml file"),(0,l.kt)("li",{parentName:"ul"},"or using a combination of both")),(0,l.kt)("h3",{id:"using-env-variables"},"Using env variables"),(0,l.kt)("p",null,"Example:"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-sh"},"export PORT=9999\ngo run main.go serve\n")),(0,l.kt)("p",null,"This will run the service on port 9999 instead of the default 8080"),(0,l.kt)("h3",{id:"using-a-yaml-file"},"Using a yaml file"),(0,l.kt)("p",null,"For default values and the structure of the yaml file please check file - ",(0,l.kt)("inlineCode",{parentName:"p"},"config.yaml.example")),(0,l.kt)("p",null,"Usage example:"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-sh"},"cp config/config.yaml config.yaml\n# make any modifications to the configs as required\ngo run main.go serve\n")),(0,l.kt)("h3",{id:"using-a-combinnation-of-both"},"Using a combinnation of both"),(0,l.kt)("p",null,"If any key that is set via both env vars and yaml the value set in env vars will take effect."))}s.isMDXComponent=!0}}]);