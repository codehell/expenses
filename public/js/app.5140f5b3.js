(function(e){function t(t){for(var n,o,i=t[0],l=t[1],u=t[2],c=0,p=[];c<i.length;c++)o=i[c],r[o]&&p.push(r[o][0]),r[o]=0;for(n in l)Object.prototype.hasOwnProperty.call(l,n)&&(e[n]=l[n]);d&&d(t);while(p.length)p.shift()();return a.push.apply(a,u||[]),s()}function s(){for(var e,t=0;t<a.length;t++){for(var s=a[t],n=!0,o=1;o<s.length;o++){var l=s[o];0!==r[l]&&(n=!1)}n&&(a.splice(t--,1),e=i(i.s=s[0]))}return e}var n={},r={app:0},a=[];function o(e){return i.p+"js/"+({about:"about"}[e]||e)+"."+{about:"8b907bfe"}[e]+".js"}function i(t){if(n[t])return n[t].exports;var s=n[t]={i:t,l:!1,exports:{}};return e[t].call(s.exports,s,s.exports,i),s.l=!0,s.exports}i.e=function(e){var t=[],s=r[e];if(0!==s)if(s)t.push(s[2]);else{var n=new Promise(function(t,n){s=r[e]=[t,n]});t.push(s[2]=n);var a,l=document.getElementsByTagName("head")[0],u=document.createElement("script");u.charset="utf-8",u.timeout=120,i.nc&&u.setAttribute("nonce",i.nc),u.src=o(e),a=function(t){u.onerror=u.onload=null,clearTimeout(c);var s=r[e];if(0!==s){if(s){var n=t&&("load"===t.type?"missing":t.type),a=t&&t.target&&t.target.src,o=new Error("Loading chunk "+e+" failed.\n("+n+": "+a+")");o.type=n,o.request=a,s[1](o)}r[e]=void 0}};var c=setTimeout(function(){a({type:"timeout",target:u})},12e4);u.onerror=u.onload=a,l.appendChild(u)}return Promise.all(t)},i.m=e,i.c=n,i.d=function(e,t,s){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:s})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var s=Object.create(null);if(i.r(s),Object.defineProperty(s,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)i.d(s,n,function(t){return e[t]}.bind(null,n));return s},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/",i.oe=function(e){throw console.error(e),e};var l=window["webpackJsonp"]=window["webpackJsonp"]||[],u=l.push.bind(l);l.push=t,l=l.slice();for(var c=0;c<l.length;c++)t(l[c]);var d=u;a.push([0,"chunk-vendors"]),s()})({0:function(e,t,s){e.exports=s("cd49")},"034f":function(e,t,s){"use strict";var n=s("64a9"),r=s.n(n);r.a},"04f0":function(e,t,s){"use strict";var n=s("bc09"),r=s.n(n);t["default"]=r.a},"0f29":function(e,t){e.exports=function(e){e.options.__i18n=e.options.__i18n||[],e.options.__i18n.push('{"en":{"errors":{"unique_violation":"This email is already registered","undefined":"Sorry, something is going wrong"},"messages":{"user_created":"User was successfully created"}}}'),delete e.options._Ctor}},"2f7b":function(e,t,s){"use strict";var n=s("42b0"),r=s.n(n);r.a},"2fa0":function(e,t,s){e.exports=s.p+"img/chlogo.b5e4ca11.png"},3024:function(e,t,s){},"42b0":function(e,t,s){},4678:function(e,t,s){var n={"./af":"2bfb","./af.js":"2bfb","./ar":"8e73","./ar-dz":"a356","./ar-dz.js":"a356","./ar-kw":"423e","./ar-kw.js":"423e","./ar-ly":"1cfd","./ar-ly.js":"1cfd","./ar-ma":"0a84","./ar-ma.js":"0a84","./ar-sa":"8230","./ar-sa.js":"8230","./ar-tn":"6d83","./ar-tn.js":"6d83","./ar.js":"8e73","./az":"485c","./az.js":"485c","./be":"1fc1","./be.js":"1fc1","./bg":"84aa","./bg.js":"84aa","./bm":"a7fa","./bm.js":"a7fa","./bn":"9043","./bn.js":"9043","./bo":"d26a","./bo.js":"d26a","./br":"6887","./br.js":"6887","./bs":"2554","./bs.js":"2554","./ca":"d716","./ca.js":"d716","./cs":"3c0d","./cs.js":"3c0d","./cv":"03ec","./cv.js":"03ec","./cy":"9797","./cy.js":"9797","./da":"0f14","./da.js":"0f14","./de":"b469","./de-at":"b3eb","./de-at.js":"b3eb","./de-ch":"bb71","./de-ch.js":"bb71","./de.js":"b469","./dv":"598a","./dv.js":"598a","./el":"8d47","./el.js":"8d47","./en-au":"0e6b","./en-au.js":"0e6b","./en-ca":"3886","./en-ca.js":"3886","./en-gb":"39a6","./en-gb.js":"39a6","./en-ie":"e1d3","./en-ie.js":"e1d3","./en-il":"7333","./en-il.js":"7333","./en-nz":"6f50","./en-nz.js":"6f50","./eo":"65db","./eo.js":"65db","./es":"898b","./es-do":"0a3c","./es-do.js":"0a3c","./es-us":"55c9","./es-us.js":"55c9","./es.js":"898b","./et":"ec18","./et.js":"ec18","./eu":"0ff2","./eu.js":"0ff2","./fa":"8df4","./fa.js":"8df4","./fi":"81e9","./fi.js":"81e9","./fo":"0721","./fo.js":"0721","./fr":"9f26","./fr-ca":"d9f8","./fr-ca.js":"d9f8","./fr-ch":"0e49","./fr-ch.js":"0e49","./fr.js":"9f26","./fy":"7118","./fy.js":"7118","./gd":"f6b4","./gd.js":"f6b4","./gl":"8840","./gl.js":"8840","./gom-latn":"0caa","./gom-latn.js":"0caa","./gu":"e0c5","./gu.js":"e0c5","./he":"c7aa","./he.js":"c7aa","./hi":"dc4d","./hi.js":"dc4d","./hr":"4ba9","./hr.js":"4ba9","./hu":"5b14","./hu.js":"5b14","./hy-am":"d6b6","./hy-am.js":"d6b6","./id":"5038","./id.js":"5038","./is":"0558","./is.js":"0558","./it":"6e98","./it.js":"6e98","./ja":"079e","./ja.js":"079e","./jv":"b540","./jv.js":"b540","./ka":"201b","./ka.js":"201b","./kk":"6d79","./kk.js":"6d79","./km":"e81d","./km.js":"e81d","./kn":"3e92","./kn.js":"3e92","./ko":"22f8","./ko.js":"22f8","./ky":"9609","./ky.js":"9609","./lb":"440c","./lb.js":"440c","./lo":"b29d","./lo.js":"b29d","./lt":"26f9","./lt.js":"26f9","./lv":"b97c","./lv.js":"b97c","./me":"293c","./me.js":"293c","./mi":"688b","./mi.js":"688b","./mk":"6909","./mk.js":"6909","./ml":"02fb","./ml.js":"02fb","./mn":"958b","./mn.js":"958b","./mr":"39bd","./mr.js":"39bd","./ms":"ebe4","./ms-my":"6403","./ms-my.js":"6403","./ms.js":"ebe4","./mt":"1b45","./mt.js":"1b45","./my":"8689","./my.js":"8689","./nb":"6ce3","./nb.js":"6ce3","./ne":"3a39","./ne.js":"3a39","./nl":"facd","./nl-be":"db29","./nl-be.js":"db29","./nl.js":"facd","./nn":"b84c","./nn.js":"b84c","./pa-in":"f3ff","./pa-in.js":"f3ff","./pl":"8d57","./pl.js":"8d57","./pt":"f260","./pt-br":"d2d4","./pt-br.js":"d2d4","./pt.js":"f260","./ro":"972c","./ro.js":"972c","./ru":"957c","./ru.js":"957c","./sd":"6784","./sd.js":"6784","./se":"ffff","./se.js":"ffff","./si":"eda5","./si.js":"eda5","./sk":"7be6","./sk.js":"7be6","./sl":"8155","./sl.js":"8155","./sq":"c8f3","./sq.js":"c8f3","./sr":"cf1e","./sr-cyrl":"13e9","./sr-cyrl.js":"13e9","./sr.js":"cf1e","./ss":"52bd","./ss.js":"52bd","./sv":"5fbd","./sv.js":"5fbd","./sw":"74dc","./sw.js":"74dc","./ta":"3de5","./ta.js":"3de5","./te":"5cbb","./te.js":"5cbb","./tet":"576c","./tet.js":"576c","./tg":"3b1b","./tg.js":"3b1b","./th":"10e8","./th.js":"10e8","./tl-ph":"0f38","./tl-ph.js":"0f38","./tlh":"cf75","./tlh.js":"cf75","./tr":"0e81","./tr.js":"0e81","./tzl":"cf51","./tzl.js":"cf51","./tzm":"c109","./tzm-latn":"b53d","./tzm-latn.js":"b53d","./tzm.js":"c109","./ug-cn":"6117","./ug-cn.js":"6117","./uk":"ada2","./uk.js":"ada2","./ur":"5294","./ur.js":"5294","./uz":"2e8c","./uz-latn":"010e","./uz-latn.js":"010e","./uz.js":"2e8c","./vi":"2921","./vi.js":"2921","./x-pseudo":"fd7e","./x-pseudo.js":"fd7e","./yo":"7f33","./yo.js":"7f33","./zh-cn":"5c3a","./zh-cn.js":"5c3a","./zh-hk":"49ab","./zh-hk.js":"49ab","./zh-tw":"90ea","./zh-tw.js":"90ea"};function r(e){var t=a(e);return s(t)}function a(e){var t=n[e];if(!(t+1)){var s=new Error("Cannot find module '"+e+"'");throw s.code="MODULE_NOT_FOUND",s}return t}r.keys=function(){return Object.keys(n)},r.resolve=a,e.exports=r,r.id="4678"},"49f8":function(e,t,s){var n={"./en.json":"edd4"};function r(e){var t=a(e);return s(t)}function a(e){var t=n[e];if(!(t+1)){var s=new Error("Cannot find module '"+e+"'");throw s.code="MODULE_NOT_FOUND",s}return t}r.keys=function(){return Object.keys(n)},r.resolve=a,e.exports=r,r.id="49f8"},"557b":function(e,t,s){},"64a9":function(e,t,s){},"9b6a":function(e,t,s){"use strict";var n=s("3024"),r=s.n(n);r.a},a2a1:function(e,t,s){"use strict";var n=s("fd18"),r=s.n(n);t["default"]=r.a},bc09:function(e,t){e.exports=function(e){e.options.__i18n=e.options.__i18n||[],e.options.__i18n.push('{"en":{"errors":{"user_not_found":"Incorrect user or password"}},"messages":{"login_ok":"Correct login"}}'),delete e.options._Ctor}},cd49:function(e,t,s){"use strict";s.r(t);var n=s("2b0e"),r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("nav",{staticClass:"nav-top"},[n("img",{staticClass:"logo",attrs:{src:s("2fa0"),alt:"chlogo"}}),n("router-link",{attrs:{to:"/"}},[e._v("Home")]),n("router-link",{attrs:{to:"/about"}},[e._v("About")]),n("router-link",{attrs:{to:"/auth/register"}},[e._v("Register")]),n("router-link",{attrs:{to:"/auth/login"}},[e._v("Login")]),n("router-link",{attrs:{to:{name:"dashboard/expenses"}}},[e._v("Dashboard")])],1),n("router-view")],1)},a=[],o=(s("034f"),s("2877")),i={},l=Object(o["a"])(i,r,a,!1,null,null,null);l.options.__file="App.vue";var u=l.exports,c=s("8c4f"),d=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home"},[n("img",{attrs:{alt:"Vue logo",src:s("cf05")}}),n("HelloWorld",{attrs:{msg:"Welcome to Your Vue.js + TypeScript App"}})],1)},p=[],f=s("9ab4"),m=s("60a3"),h=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",{staticClass:"hello"},[s("h1",[e._v(e._s(e.msg))]),e._m(0),s("h3",[e._v("Installed CLI Plugins")]),e._m(1),s("h3",[e._v("Essential Links")]),e._m(2),s("h3",[e._v("Ecosystem")]),e._m(3)])},v=[function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("p",[e._v("\n    For guide and recipes on how to configure / customize this project,"),s("br"),e._v("\n    check out the\n    "),s("a",{attrs:{href:"https://cli.vuejs.org",target:"_blank",rel:"noopener"}},[e._v("vue-cli documentation")]),e._v(".\n  ")])},function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("ul",[s("li",[s("a",{attrs:{href:"https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-typescript",target:"_blank",rel:"noopener"}},[e._v("typescript")])]),s("li",[s("a",{attrs:{href:"https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-eslint",target:"_blank",rel:"noopener"}},[e._v("eslint")])]),s("li",[s("a",{attrs:{href:"https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-unit-mocha",target:"_blank",rel:"noopener"}},[e._v("unit-mocha")])]),s("li",[s("a",{attrs:{href:"https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-e2e-cypress",target:"_blank",rel:"noopener"}},[e._v("e2e-cypress")])])])},function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("ul",[s("li",[s("a",{attrs:{href:"https://vuejs.org",target:"_blank",rel:"noopener"}},[e._v("Core Docs")])]),s("li",[s("a",{attrs:{href:"https://forum.vuejs.org",target:"_blank",rel:"noopener"}},[e._v("Forum")])]),s("li",[s("a",{attrs:{href:"https://chat.vuejs.org",target:"_blank",rel:"noopener"}},[e._v("Community Chat")])]),s("li",[s("a",{attrs:{href:"https://twitter.com/vuejs",target:"_blank",rel:"noopener"}},[e._v("Twitter")])]),s("li",[s("a",{attrs:{href:"https://news.vuejs.org",target:"_blank",rel:"noopener"}},[e._v("News")])])])},function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("ul",[s("li",[s("a",{attrs:{href:"https://router.vuejs.org",target:"_blank",rel:"noopener"}},[e._v("vue-router")])]),s("li",[s("a",{attrs:{href:"https://vuex.vuejs.org",target:"_blank",rel:"noopener"}},[e._v("vuex")])]),s("li",[s("a",{attrs:{href:"https://github.com/vuejs/vue-devtools#vue-devtools",target:"_blank",rel:"noopener"}},[e._v("vue-devtools")])]),s("li",[s("a",{attrs:{href:"https://vue-loader.vuejs.org",target:"_blank",rel:"noopener"}},[e._v("vue-loader")])]),s("li",[s("a",{attrs:{href:"https://github.com/vuejs/awesome-vue",target:"_blank",rel:"noopener"}},[e._v("awesome-vue")])])])}],b=function(e){function t(){return null!==e&&e.apply(this,arguments)||this}return f["b"](t,e),f["a"]([Object(m["c"])()],t.prototype,"msg",void 0),t=f["a"]([m["a"]],t),t}(m["d"]),g=b,_=g,j=(s("2f7b"),Object(o["a"])(_,h,v,!1,null,"ee849752",null));j.options.__file="HelloWorld.vue";var x=j.exports,y=function(e){function t(){return null!==e&&e.apply(this,arguments)||this}return f["b"](t,e),t=f["a"]([Object(m["a"])({components:{HelloWorld:x}})],t),t}(m["d"]),w=y,C=w,k=Object(o["a"])(C,d,p,!1,null,null,null);k.options.__file="Home.vue";var E=k.exports,O=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",[s("router-view")],1)},$=[],z={},P=Object(o["a"])(z,O,$,!1,null,null,null);P.options.__file="Auth.vue";var T=P.exports,S=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",{staticClass:"flex justify-center container mx-auto"},[s("div",{staticClass:"w-full max-w-xs"},[s("form",{staticClass:"bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4",on:{submit:function(t){return t.preventDefault(),e.send(t)}}},[s("div",{staticClass:"mb-4"},[s("label",{staticClass:"block text-grey-darker text-sm font-bold mb-2",attrs:{for:"email"}},[e._v("\n          Email\n        ")]),s("input",{directives:[{name:"model",rawName:"v-model",value:e.email,expression:"email"}],staticClass:"shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker leading-tight focus:outline-none focus:shadow-outline",attrs:{id:"email",type:"email",placeholder:"email",required:""},domProps:{value:e.email},on:{input:function(t){t.target.composing||(e.email=t.target.value)}}})]),s("div",{staticClass:"mb-6"},[s("label",{staticClass:"block text-grey-darker text-sm font-bold mb-2",attrs:{for:"password"}},[e._v("\n          Password\n        ")]),s("input",{directives:[{name:"model",rawName:"v-model",value:e.password,expression:"password"}],staticClass:"shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker mb-3 leading-tight focus:outline-none focus:shadow-outline",attrs:{id:"password",type:"password",placeholder:"******************",required:"",maxlength:"16",minlength:"6"},domProps:{value:e.password},on:{input:function(t){t.target.composing||(e.password=t.target.value)}}})]),e.error?s("div",{staticClass:"flex items-center bg-red text-white text-sm font-bold px-4 py-3 mb-2",attrs:{role:"alert"}},[s("svg",{staticClass:"fill-current w-4 h-4 mr-2",attrs:{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 20 20"}},[s("path",{attrs:{d:"M12.432 0c1.34 0 2.01.912 2.01 1.957 0 1.305-1.164 2.512-2.679 2.512-1.269 0-2.009-.75-1.974-1.99C9.789 1.436 10.67 0 12.432 0zM8.309 20c-1.058 0-1.833-.652-1.093-3.524l1.214-5.092c.211-.814.246-1.141 0-1.141-.317 0-1.689.562-2.502 1.117l-.528-.88c2.572-2.186 5.531-3.467 6.801-3.467 1.057 0 1.233 1.273.705 3.23l-1.391 5.352c-.246.945-.141 1.271.106 1.271.317 0 1.357-.392 2.379-1.207l.6.814C12.098 19.02 9.365 20 8.309 20z"}})]),s("p",[e._v(e._s(e.error))])]):e._e(),e.message?s("div",{staticClass:"flex items-center bg-green text-white text-sm font-bold px-4 py-3 mb-2",attrs:{role:"alert"}},[s("svg",{staticClass:"fill-current w-4 h-4 mr-2",attrs:{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 20 20"}},[s("path",{attrs:{d:"M12.432 0c1.34 0 2.01.912 2.01 1.957 0 1.305-1.164 2.512-2.679 2.512-1.269 0-2.009-.75-1.974-1.99C9.789 1.436 10.67 0 12.432 0zM8.309 20c-1.058 0-1.833-.652-1.093-3.524l1.214-5.092c.211-.814.246-1.141 0-1.141-.317 0-1.689.562-2.502 1.117l-.528-.88c2.572-2.186 5.531-3.467 6.801-3.467 1.057 0 1.233 1.273.705 3.23l-1.391 5.352c-.246.945-.141 1.271.106 1.271.317 0 1.357-.392 2.379-1.207l.6.814C12.098 19.02 9.365 20 8.309 20z"}})]),s("p",[e._v(e._s(e.message))])]):e._e(),e._m(0)]),s("p",{staticClass:"text-center text-grey text-xs"},[e._v("\n      ©2018 Acme Corp. All rights reserved.\n    ")])])])},N=[function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",{staticClass:"flex items-center justify-between"},[s("button",{staticClass:"bg-blue hover:bg-blue-dark text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline",attrs:{type:"submit"}},[e._v("\n          Sign In\n        ")])])}],D=s("65d9"),A=s.n(D),L={baseUrl:"http://localhost:3000/"},U=function(e){function t(){var t=null!==e&&e.apply(this,arguments)||this;return t.email="",t.password="",t.error="",t.message="",t}return f["b"](t,e),t.prototype.send=function(){var e=this;this.error="",this.message="";var t={email:this.email,password:this.password},s=new Headers;s.set("Content-Type","application/json"),fetch(L.baseUrl+"auth/register",{method:"POST",body:JSON.stringify(t),headers:s}).then(function(t){if(201!==t.status)throw new Error(t.statusText);e.message=e.$t("messages.user_created")}).catch(function(t){"Conflict"===t.message?e.error=e.$t("errors.unique_violation"):e.error=e.$t("errors.undefined")})},t=f["a"]([A.a],t),t}(n["default"]),q=U,M=q,F=s("e349"),H=Object(o["a"])(M,S,N,!1,null,null,null);"function"===typeof F["default"]&&Object(F["default"])(H),H.options.__file="Register.vue";var B=H.exports,I=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",{staticClass:"flex justify-center container mx-auto"},[s("div",{staticClass:"w-full max-w-xs"},[s("form",{staticClass:"bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4",on:{submit:function(t){return t.preventDefault(),e.send(t)}}},[s("div",{staticClass:"mb-4"},[s("label",{staticClass:"block text-grey-darker text-sm font-bold mb-2",attrs:{for:"email"}},[e._v("\n          Email\n        ")]),s("input",{directives:[{name:"model",rawName:"v-model",value:e.email,expression:"email"}],staticClass:"shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker leading-tight focus:outline-none focus:shadow-outline",attrs:{id:"email",type:"email",placeholder:"email",required:""},domProps:{value:e.email},on:{input:function(t){t.target.composing||(e.email=t.target.value)}}})]),s("div",{staticClass:"mb-6"},[s("label",{staticClass:"block text-grey-darker text-sm font-bold mb-2",attrs:{for:"password"}},[e._v("\n          Password\n        ")]),s("input",{directives:[{name:"model",rawName:"v-model",value:e.password,expression:"password"}],staticClass:"shadow appearance-none rounded w-full py-2 px-3 text-grey-darker mb-3 leading-tight focus:outline-none focus:shadow-outline",attrs:{id:"password",type:"password",placeholder:"******************",required:"",maxlength:"16",minlength:"6"},domProps:{value:e.password},on:{input:function(t){t.target.composing||(e.password=t.target.value)}}})]),e.error?s("div",{staticClass:"flex items-center bg-red text-white text-sm font-bold px-4 py-3 mb-2",attrs:{role:"alert"}},[s("p",[e._v(e._s(e.error))])]):e._e(),e.message?s("div",{staticClass:"flex items-center bg-green text-white text-sm font-bold px-4 py-3 mb-2",attrs:{role:"alert"}},[s("p",[e._v(e._s(e.message))])]):e._e(),e._m(0)]),s("p",{staticClass:"text-center text-grey text-xs"},[e._v("\n      ©2018 Acme Corp. All rights reserved.\n    ")])])])},J=[function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",{staticClass:"flex items-center justify-between"},[s("button",{staticClass:"bg-blue hover:bg-blue-dark text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline",attrs:{type:"submit"}},[e._v("\n          Login\n        ")]),s("a",{staticClass:"inline-block align-baseline font-bold text-sm text-blue hover:text-blue-darker",attrs:{href:"#"}},[e._v("\n          Forgot Password?\n        ")])])}],W=function(e){function t(){var t=null!==e&&e.apply(this,arguments)||this;return t.email="",t.password="",t.error="",t.message="",t}return f["b"](t,e),t.prototype.send=function(){var e=this,t=new Headers;t.set("Content-type","text/plain"),t.set("Authorization","Basic "+btoa(this.email+":"+this.password)),fetch(L.baseUrl+"auth/login",{method:"POST",headers:t}).then(function(t){if(e.error="",e.message="",200===t.status)t.json().then(function(t){window.localStorage.setItem("token",t.token),e.$router.push({name:"dashboard/expenses"})}),e.message=e.$t("messages.login_ok");else{if(401!==t.status)throw new Error(t.statusText);e.error=e.$t("errors.user_not_found")}}).catch(function(){e.error=e.$t("errors.undefined")})},t=f["a"]([A.a],t),t}(n["default"]),R=W,V=R,G=s("04f0"),Y=Object(o["a"])(V,I,J,!1,null,null,null);"function"===typeof G["default"]&&Object(G["default"])(Y),Y.options.__file="Login.vue";var Z=Y.exports,K=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",[s("h2",{staticClass:"mb-5 text-center"},[e._v("Expenses")]),e.error?s("div",[s("p",[e._v(e._s(e.error))])]):e._e(),e.message?s("div",[s("p",[e._v(e._s(e.message))])]):e._e(),s("div",{staticClass:"flex flex-wrap justify-evenly container mx-auto"},[s("div",{staticClass:"w-full max-w-xs"},[s("ExpensesSeeker")],1),s("div",{staticClass:"w-full max-w-xs"},[s("form",{staticClass:"bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4",on:{submit:function(t){return t.preventDefault(),e.create(t)}}},[s("div",{staticClass:"mb-4"},[s("input",{directives:[{name:"model",rawName:"v-model",value:e.amount,expression:"amount"}],staticClass:"ch-input",attrs:{type:"text",placeholder:e.$t("texts.amount"),required:""},domProps:{value:e.amount},on:{input:function(t){t.target.composing||(e.amount=t.target.value)}}})]),s("div",{staticClass:"mb-4"},[s("input",{directives:[{name:"model",rawName:"v-model",value:e.description,expression:"description"}],staticClass:"ch-input",attrs:{type:"text",placeholder:e.$t("texts.description")},domProps:{value:e.description},on:{input:function(t){t.target.composing||(e.description=t.target.value)}}})]),s("div",{staticClass:"mb-4"},[s("select",{directives:[{name:"model",rawName:"v-model",value:e.selectedTags,expression:"selectedTags"}],staticClass:"ch-input",attrs:{multiple:""},on:{change:function(t){var s=Array.prototype.filter.call(t.target.options,function(e){return e.selected}).map(function(e){var t="_value"in e?e._value:e.value;return t});e.selectedTags=t.target.multiple?s:s[0]}}},e._l(e.tags,function(t){return s("option",{key:t.id,domProps:{value:t.id}},[e._v(e._s(t.name))])}))]),s("button",{staticClass:"btn btn-indigo",attrs:{type:"submit"}},[e._v("Submit")]),s("button",{staticClass:"btn btn-indigo",attrs:{type:"button"},on:{click:function(t){e.refresh()}}},[e._v("Refresh")])])]),s("div",{staticClass:"w-full max-w-xs"},[s("form",{staticClass:"bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4",on:{submit:function(t){return t.preventDefault(),e.createTag(t)}}},[s("div",{staticClass:"mb-4"},[s("input",{directives:[{name:"model",rawName:"v-model",value:e.tag,expression:"tag"}],staticClass:"ch-input",attrs:{type:"text",placeholder:e.$t("texts.new_tag")},domProps:{value:e.tag},on:{input:function(t){t.target.composing||(e.tag=t.target.value)}}})]),s("button",{staticClass:"btn btn-indigo",attrs:{type:"submit"}},[e._v("Submit")])])])]),s("ExpensesList",{attrs:{expenses:e.filtered},on:{"delete-expense":e.deleteExpense}})],1)},Q=[],X=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",{staticClass:"flex flex-wrap justify-center md:justify-around container mx-auto"},e._l(e.expenses,function(t){return s("div",{key:t.id,staticClass:"expense"},[s("div",{staticClass:"expense-body"},[s("div",{staticClass:"expense-data"},[s("div",{staticClass:"expense-data-item"},[e._v(e._s(t.amount)+"€")]),s("div",{staticClass:"expense-data-item"},[e._v(e._s(t.description))]),s("div",{staticClass:"expense-data-item"},[e._v(e._s(e.parseDate(t.created_at)))])]),s("div",{staticClass:"expense-actions"},[s("div",{staticClass:"cursor-pointer"},[s("font-awesome-icon",{attrs:{icon:"pen"}})],1),s("div",{staticClass:"cursor-pointer ml-2",on:{click:function(s){e.deleteExpense(t)}}},[s("font-awesome-icon",{attrs:{icon:"trash"}})],1)])]),s("div",{staticClass:"w-full flex flex-wrap text-white mt-1"},e._l(t.tags,function(t){return s("div",{key:t.id,staticClass:"ml-1 mt-1 bg-indigo-dark px-1 rounded-sm h-5"},[e._v(e._s(t.name))])}))])}))},ee=[],te=s("c1df"),se=s.n(te),ne=function(e){function t(){return null!==e&&e.apply(this,arguments)||this}return f["b"](t,e),t.prototype.deleteExpense=function(e){return e},t.prototype.parseDate=function(e){return se()(e).format("lll")},f["a"]([Object(m["c"])()],t.prototype,"expenses",void 0),f["a"]([Object(m["b"])()],t.prototype,"deleteExpense",null),t=f["a"]([m["a"]],t),t}(m["d"]),re=ne,ae=re,oe=(s("9b6a"),Object(o["a"])(ae,X,ee,!1,null,"5523e06c",null));oe.options.__file="ExpensesList.vue";var ie=oe.exports,le=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",[s("form",{staticClass:"bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4",on:{submit:function(t){return t.preventDefault(),e.search(t)}}},[s("div",{staticClass:"mb-4"},[s("input",{directives:[{name:"model",rawName:"v-model",value:e.date,expression:"date"}],ref:"flp",staticClass:"ch-input",attrs:{name:"date",placeholder:"By date"},domProps:{value:e.date},on:{input:function(t){t.target.composing||(e.date=t.target.value)}}})]),s("div",{staticClass:"mb-4"},[s("input",{directives:[{name:"model",rawName:"v-model",value:e.description,expression:"description"}],staticClass:"ch-input",attrs:{name:"description",type:"text",placeholder:"By description"},domProps:{value:e.description},on:{input:function(t){t.target.composing||(e.description=t.target.value)}}})]),s("button",{staticClass:"btn btn-indigo",attrs:{type:"submit"}},[e._v("Search")]),s("button",{staticClass:"btn btn-indigo",attrs:{type:"button"},on:{click:e.clear}},[e._v("Clear")])])])},ue=[],ce=(s("7952"),s("cf06")),de=s.n(ce),pe=function(e){function t(){var t=null!==e&&e.apply(this,arguments)||this;return t.date="",t.description="",t}return f["b"](t,e),Object.defineProperty(t.prototype,"expenses",{get:function(){return this.$store.state.expenses},enumerable:!0,configurable:!0}),t.prototype.clear=function(){this.date="",this.description="",this.search()},t.prototype.search=function(){var e=this,t=this.expenses;if(this.date.length>0){var s=se()(this.date);t=t.filter(function(e){var t=se()(e.created_at).startOf("day"),n=s.diff(t,"days");return 0===n})}this.description.length>0&&(t=t.filter(function(t){return t.description.toLowerCase().includes(e.description)})),this.$store.commit("setFiltered",t)},t.prototype.mounted=function(){var e=this.$refs.flp;de()(e,{})},t=f["a"]([m["a"]],t),t}(m["d"]),fe=pe,me=fe,he=Object(o["a"])(me,le,ue,!1,null,null,null);he.options.__file="ExpensesSeeker.vue";var ve=he.exports,be=function(e){function t(){var t=null!==e&&e.apply(this,arguments)||this;return t.amount="",t.description="",t.error="",t.headers=new Headers,t.message="",t.tag="",t.tags=[],t.selectedTags=[],t}return f["b"](t,e),Object.defineProperty(t.prototype,"expenses",{get:function(){return this.$store.state.expenses},enumerable:!0,configurable:!0}),Object.defineProperty(t.prototype,"filtered",{get:function(){return this.$store.state.filtered},enumerable:!0,configurable:!0}),t.prototype.create=function(){var e=this;this.error="",this.message="";var t=[];this.selectedTags.forEach(function(e){t.push({id:e})});var s={amount:this.amount,description:this.description,tags:t};this.headers.set("Content-Type","application/json"),fetch(L.baseUrl+"expenses",{method:"Post",headers:this.headers,body:JSON.stringify(s)}).then(function(t){201===t.status?e.message=e.$t("messages.created"):e.message=e.$t("errors.undefined")})},t.prototype.createTag=function(){var e=this;this.error="",this.message="";var t={name:this.tag};this.headers.set("Content-Type","application/json"),fetch(L.baseUrl+"tags",{method:"Post",headers:this.headers,body:JSON.stringify(t)}).then(function(t){201===t.status?e.message=e.$t("messages.created"):e.message=e.$t("errors.undefined")})},t.prototype.deleteExpense=function(e){var t=this;this.error="",this.message="",this.headers.set("Content-Type","application/json"),fetch(L.baseUrl+"expenses",{method:"Delete",headers:this.headers,body:JSON.stringify(e)}).then(function(e){204===e.status?t.message=t.$t("messages.expense_deleted"):t.message=t.$t("errors.undefined")})},t.prototype.refresh=function(){var e=this;this.error="",this.message="",fetch(L.baseUrl+"expenses",{method:"GET",headers:this.headers}).then(function(t){401===t.status?e.$router.push({name:"login"}):200!==t.status?e.error=e.$t("errors.undefined"):t.json().then(function(t){e.$store.commit("setExpenses",t.expenses),e.$store.commit("setFiltered",t.expenses),e.tags=t.tags})}).catch(function(){e.error=e.$t("errors.undefined")})},t.prototype.parseDate=function(e){return se()(e).format("lll")},t.prototype.mounted=function(){this.headers.append("Authorization","Bearer "+window.localStorage.getItem("token")),this.refresh()},t=f["a"]([A()({components:{ExpensesList:ie,ExpensesSeeker:ve}})],t),t}(n["default"]),ge=be,_e=ge,je=s("a2a1"),xe=Object(o["a"])(_e,K,Q,!1,null,null,null);"function"===typeof je["default"]&&Object(je["default"])(xe),xe.options.__file="Expenses.vue";var ye=xe.exports,we=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",[s("router-view")],1)},Ce=[],ke={},Ee=Object(o["a"])(ke,we,Ce,!1,null,null,null);Ee.options.__file="Dashboard.vue";var Oe=Ee.exports;n["default"].use(c["a"]);var $e=new c["a"]({mode:"history",base:"/",routes:[{path:"/",name:"home",component:E},{path:"/about",name:"about",component:function(){return s.e("about").then(s.bind(null,"f820"))}},{path:"/auth",name:"auth",component:T,children:[{path:"register",name:"register",component:B},{path:"login",name:"login",component:Z}]},{path:"/dashboard",name:"dashboard",component:Oe,children:[{path:"expenses",name:"dashboard/expenses",component:ye}]}]}),ze=s("2f62");n["default"].use(ze["a"]);var Pe=new ze["a"].Store({state:{expenses:[],filtered:[]},mutations:{setExpenses:function(e,t){e.expenses=t},setFiltered:function(e,t){e.filtered=t}},actions:{}}),Te=(s("557b"),s("a925"));function Se(){var e=s("49f8"),t={};return e.keys().forEach(function(s){var n=s.match(/([A-Za-z0-9-_]+)\./i);if(n&&n.length>1){var r=n[1];t[r]=e(s)}}),t}n["default"].use(Te["a"]);var Ne=new Te["a"]({locale:"en",fallbackLocale:"en",messages:Se()}),De=s("ecee"),Ae=s("c074"),Le=s("ad3d");De["c"].add(Ae["b"],Ae["a"]),n["default"].component("font-awesome-icon",Le.FontAwesomeIcon),n["default"].config.productionTip=!1,new n["default"]({router:$e,store:Pe,i18n:Ne,render:function(e){return e(u)}}).$mount("#app")},cf05:function(e,t,s){e.exports=s.p+"img/logo.82b9c7a5.png"},e349:function(e,t,s){"use strict";var n=s("0f29"),r=s.n(n);t["default"]=r.a},edd4:function(e){e.exports={errors:{undefined:"Sorry, something is going wrong"},messages:{created:"The resource was created successfully"}}},fd18:function(e,t){e.exports=function(e){e.options.__i18n=e.options.__i18n||[],e.options.__i18n.push('{"en":{"texts":{"description":"Description","amount":"Amount","new_tag":"New tag"},"messages":{"expense_deleted":"The expense was deleted successfully"}}}'),delete e.options._Ctor}}});
//# sourceMappingURL=app.5140f5b3.js.map