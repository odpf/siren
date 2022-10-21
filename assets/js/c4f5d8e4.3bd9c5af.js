/*! For license information please see c4f5d8e4.3bd9c5af.js.LICENSE.txt */
(self.webpackChunksiren=self.webpackChunksiren||[]).push([[195],{4035:function(e,t,n){"use strict";n.r(t),n.d(t,{default:function(){return f}});var r=n(7294),a=n(4704);function l(e){var t,n,r="";if("string"==typeof e||"number"==typeof e)r+=e;else if("object"==typeof e)if(Array.isArray(e))for(t=0;t<e.length;t++)e[t]&&(n=l(e[t]))&&(r&&(r+=" "),r+=n);else for(t in e)e[t]&&(r&&(r+=" "),r+=t);return r}function i(){for(var e,t,n=0,r="";n<arguments.length;)(e=arguments[n++])&&(t=l(e))&&(r&&(r+=" "),r+=t);return r}var s=n(2263),o=n(4184),c=n.n(o);const u=e=>{const t=c()(e.className,{darkBackground:"dark"===e.background,highlightBackground:"highlight"===e.background,lightBackground:"light"===e.background,paddingAll:e.padding.indexOf("all")>=0,paddingBottom:e.padding.indexOf("bottom")>=0,paddingLeft:e.padding.indexOf("left")>=0,paddingRight:e.padding.indexOf("right")>=0,paddingTop:e.padding.indexOf("top")>=0});let n;return n=e.wrapper?r.createElement("div",{className:"container"},e.children):e.children,r.createElement("div",{className:t,id:e.id},n)};u.defaultProps={background:null,padding:[],wrapper:!0};var m=u;class d extends r.Component{renderBlock(e){const t={imageAlign:"left",...e},n=c()("blockElement",this.props.className,{alignCenter:"center"===this.props.align,alignRight:"right"===this.props.align,fourByGridBlock:"fourColumn"===this.props.layout,threeByGridBlock:"threeColumn"===this.props.layout,twoByGridBlock:"twoColumn"===this.props.layout});return r.createElement("div",{className:n,key:t.title},r.createElement("div",{className:"blockContent"},this.renderBlockTitle(t.title),t.content))}renderBlockTitle(e){return e?r.createElement("h2",null,e):null}render(){return r.createElement("div",{className:"gridBlock"},this.props.contents.map(this.renderBlock,this))}}d.defaultProps={align:"left",contents:[],layout:"twoColumn"};var p=d,g=n(4996);const h=()=>{const{siteConfig:e}=(0,s.Z)();return r.createElement("div",{className:"homeHero"},r.createElement("div",{className:"logo"},r.createElement("img",{src:(0,g.Z)("img/pattern.svg")})),r.createElement("div",{className:"container banner"},r.createElement("div",{className:"row"},r.createElement("div",{className:i("col col--5")},r.createElement("div",{className:"homeTitle"},e.tagline),r.createElement("small",{className:"homeSubTitle"},"Siren provides an easy-to-use universal alert, notification, channels management framework for the entire observability infrastructure."),r.createElement("a",{className:"button",href:"docs/introduction"},"Documentation")),r.createElement("div",{className:i("col col--1")}),r.createElement("div",{className:i("col col--6")},r.createElement("div",{className:"text--right"},r.createElement("img",{src:(0,g.Z)("img/banner.svg")}))))))};function f(){const{siteConfig:e}=(0,s.Z)();return r.createElement(a.Z,{title:e.tagline,description:"Siren provides an easy-to-use universal alert, notification, channels management framework for the entire observability infrastructure."},r.createElement(h,null),r.createElement("main",null,r.createElement(m,{className:"textSection wrapper",background:"light"},r.createElement("h1",null,"Built for scale"),r.createElement("p",null,"Siren provides an easy-to-use universal alert, notification, channels management framework for the entire observability infrastructure.."),r.createElement(p,{layout:"threeColumn",contents:[{title:"Rule Templates",content:r.createElement("div",null,"Siren provides a way to define templates over prometheus Rule, which can be reused to create multiple instances of same rule with configurable thresholds.")},{title:"Multi-tenancy",content:r.createElement("div",null,"Rules created with Siren are by default multi-tenancy aware.")},{title:"DIY Interface",content:r.createElement("div",null,"Siren can be used to easily create/edit prometheus rules. It also provides soft delete(disable) so that you can preserve thresholds in case you need to reuse the same alert.")},{title:"Managing bulk rules",content:r.createElement("div",null,"Siren enables users to manage bulk alerts using YAML files in specified format using simple CLI.")},{title:"Credentials Management",content:r.createElement("div",null,"Siren can store slack and pagerduty credentials, sync them with Cortex alertmanager to deliver alerts on proper channels, in a multi-tenant fashion. It gives a simple interface to rotate the credentials on demand via HTTP API.")},{title:"Alert History",content:r.createElement("div",null,"Siren can store alerts triggered via Cortex Alertmanager, which can be used for audit purposes.")}]})),r.createElement(m,{className:"textSection wrapper",background:"light"},r.createElement("h1",null,"Trusted by"),r.createElement("p",null,"Siren was originally created for the Gojek data processing platform, and it has been used, adapted and improved by other teams internally and externally."),r.createElement(p,{className:"logos",layout:"fourColumn",contents:[{content:r.createElement("img",{src:(0,g.Z)("users/gojek.png")})},{content:r.createElement("img",{src:(0,g.Z)("users/midtrans.png")})},{content:r.createElement("img",{src:(0,g.Z)("users/mapan.png")})},{content:r.createElement("img",{src:(0,g.Z)("users/moka.png")})}]}))))}},4184:function(e,t){var n;!function(){"use strict";var r={}.hasOwnProperty;function a(){for(var e=[],t=0;t<arguments.length;t++){var n=arguments[t];if(n){var l=typeof n;if("string"===l||"number"===l)e.push(n);else if(Array.isArray(n)){if(n.length){var i=a.apply(null,n);i&&e.push(i)}}else if("object"===l)if(n.toString===Object.prototype.toString)for(var s in n)r.call(n,s)&&n[s]&&e.push(s);else e.push(n.toString())}}return e.join(" ")}e.exports?(a.default=a,e.exports=a):void 0===(n=function(){return a}.apply(t,[]))||(e.exports=n)}()}}]);