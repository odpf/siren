"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[217],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>m});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),p=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=p(e.components);return r.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},c=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,s=e.parentName,u=o(e,["components","mdxType","originalType","parentName"]),c=p(n),m=a,h=c["".concat(s,".").concat(m)]||c[m]||d[m]||l;return n?r.createElement(h,i(i({ref:t},u),{},{components:n})):r.createElement(h,i({ref:t},u))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,i=new Array(l);i[0]=c;var o={};for(var s in t)hasOwnProperty.call(t,s)&&(o[s]=t[s]);o.originalType=e,o.mdxType="string"==typeof e?e:a,i[1]=o;for(var p=2;p<l;p++)i[p]=n[p];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}c.displayName="MDXCreateElement"},9803:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>i,default:()=>d,frontMatter:()=>l,metadata:()=>o,toc:()=>p});var r=n(3117),a=(n(7294),n(3905));const l={},i="Installation",o={unversionedId:"installation",id:"installation",title:"Installation",description:"There are several approaches to install Siren CLI",source:"@site/docs/installation.md",sourceDirName:".",slug:"/installation",permalink:"/siren/docs/installation",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/installation.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Use Cases",permalink:"/siren/docs/use_cases"},next:{title:"Introduction",permalink:"/siren/docs/tour/introduction"}},s={},p=[{value:"Binary (Cross-platform)",id:"binary-cross-platform",level:4},{value:"macOS",id:"macos",level:4},{value:"Linux",id:"linux",level:4},{value:"Windows",id:"windows",level:4},{value:"Building from source",id:"building-from-source",level:3},{value:"Prerequisites",id:"prerequisites",level:4},{value:"Build",id:"build",level:4},{value:"Use the Docker image",id:"use-the-docker-image",level:3},{value:"Use the Helm chart",id:"use-the-helm-chart",level:3},{value:"Verifying the installation\u200b",id:"verifying-the-installation",level:3},{value:"Dockerized dependencies",id:"dockerized-dependencies",level:3}],u={toc:p};function d(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"installation"},"Installation"),(0,a.kt)("p",null,"There are several approaches to install Siren CLI"),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("a",{parentName:"li",href:"#binary-cross-platform"},"Using a pre-compiled binary")),(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("a",{parentName:"li",href:"#macOS"},"Installing with package manager")),(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("a",{parentName:"li",href:"#building-from-source"},"Installing from source")),(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("a",{parentName:"li",href:"#use-the-docker-image"},"Using the Docker image")),(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("a",{parentName:"li",href:"#use-the-helm-chart"},"Using the Helm Chart"))),(0,a.kt)("h4",{id:"binary-cross-platform"},"Binary (Cross-platform)"),(0,a.kt)("p",null,"Download the appropriate version for your platform from ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/odpf/siren/releases"},"releases")," page. Once downloaded, the binary can be run from anywhere.\nYou don\u2019t need to install it into a global location. This works well for shared hosts and other systems where you don\u2019t have a privileged account.\nIdeally, you should install it somewhere in your PATH for easy use. ",(0,a.kt)("inlineCode",{parentName:"p"},"/usr/local/bin")," is the most probable location."),(0,a.kt)("h4",{id:"macos"},"macOS"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"siren")," is available via a Homebrew Tap, and as downloadable binary from the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/odpf/siren/releases/latest"},"releases")," page:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-sh"},"brew install odpf/tap/siren\n")),(0,a.kt)("p",null,"To upgrade to the latest version:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"brew upgrade siren\n")),(0,a.kt)("p",null,"Check for installed siren version"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-sh"},"siren version\n")),(0,a.kt)("h4",{id:"linux"},"Linux"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"siren")," is available as downloadable binaries from the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/odpf/siren/releases/latest"},"releases")," page. Download the ",(0,a.kt)("inlineCode",{parentName:"p"},".deb")," or ",(0,a.kt)("inlineCode",{parentName:"p"},".rpm")," from the releases page and install with ",(0,a.kt)("inlineCode",{parentName:"p"},"sudo dpkg -i")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"sudo rpm -i")," respectively."),(0,a.kt)("h4",{id:"windows"},"Windows"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"siren")," is available via ",(0,a.kt)("a",{parentName:"p",href:"https://scoop.sh/"},"scoop"),", and as a downloadable binary from the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/odpf/siren/releases/latest"},"releases")," page:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"scoop bucket add siren https://github.com/odpf/scoop-bucket.git\n")),(0,a.kt)("p",null,"To upgrade to the latest version:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"scoop update siren\n")),(0,a.kt)("h3",{id:"building-from-source"},"Building from source"),(0,a.kt)("h4",{id:"prerequisites"},"Prerequisites"),(0,a.kt)("p",null,"Siren requires the following dependencies:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Golang (version 1.16 or above)"),(0,a.kt)("li",{parentName:"ul"},"Git")),(0,a.kt)("h4",{id:"build"},"Build"),(0,a.kt)("p",null,"Run either of the following commands to clone and compile Siren from source"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-sh"},"git clone git@github.com:odpf/siren.git  (Using SSH Protocol) Or\ngit clone https://github.com/odpf/siren.git (Using HTTPS Protocol)\n")),(0,a.kt)("p",null,"Install all the golang dependencies"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"make setup\n")),(0,a.kt)("p",null,"Build siren binary file"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"make build\n")),(0,a.kt)("p",null,"Init server config. Customise with your local configurations."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"make config\n")),(0,a.kt)("p",null,"Run database migrations"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"$ siren server migrate -c config.yaml\n")),(0,a.kt)("p",null,"Start siren server"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"$ siren server start -c config.yaml\n")),(0,a.kt)("p",null,"Initialize client configurations"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"$ siren config init\n")),(0,a.kt)("h3",{id:"use-the-docker-image"},"Use the Docker image"),(0,a.kt)("p",null,"We provide ready to use Docker container images. To pull the latest image:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"docker pull odpf/siren:latest\n")),(0,a.kt)("p",null,"To pull a specific version:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"docker pull odpf/siren:v0.4.1\n")),(0,a.kt)("h3",{id:"use-the-helm-chart"},"Use the Helm chart"),(0,a.kt)("p",null,"Siren can be installed in Kubernetes using the Helm chart from ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/odpf/charts"},"https://github.com/odpf/charts"),"."),(0,a.kt)("p",null,"Ensure that the following requirements are met:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Kubernetes 1.14+"),(0,a.kt)("li",{parentName:"ul"},"Helm version 3.x is ",(0,a.kt)("a",{parentName:"li",href:"https://helm.sh/docs/intro/install/"},"installed"))),(0,a.kt)("p",null,"Add ODPF chart repository to Helm:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"helm repo add odpf https://odpf.github.io/charts/\n")),(0,a.kt)("p",null,"You can update the chart repository by running:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"helm repo update\n")),(0,a.kt)("p",null,"And install it with the helm command line:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"helm install my-release odpf/siren\n")),(0,a.kt)("h3",{id:"verifying-the-installation"},"Verifying the installation\u200b"),(0,a.kt)("p",null,"To verify if Siren is properly installed, run ",(0,a.kt)("inlineCode",{parentName:"p"},"siren --help")," on your system. You should see help output. If you are executing it from the command line, make sure it is on your PATH or you may get an error about Siren not being found."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"$ siren --help\n")),(0,a.kt)("h3",{id:"dockerized-dependencies"},"Dockerized dependencies"),(0,a.kt)("p",null,"  You will notice there is a ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/odpf/siren/blob/main/docker-compose.yaml"},(0,a.kt)("inlineCode",{parentName:"a"},"docker-compose.yaml"))," file contains all dependencies that you need to bootstrap Siren. Inside, it has ",(0,a.kt)("inlineCode",{parentName:"p"},"postgresql")," as a main storage, ",(0,a.kt)("inlineCode",{parentName:"p"},"cortex ruler")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"cortex alertmanager")," as monitoring provider, and ",(0,a.kt)("inlineCode",{parentName:"p"},"minio")," as a backend storage for ",(0,a.kt)("inlineCode",{parentName:"p"},"cortex"),"."))}d.isMDXComponent=!0}}]);