(function(e){function t(t){for(var a,r,o=t[0],l=t[1],c=t[2],d=0,m=[];d<o.length;d++)r=o[d],Object.prototype.hasOwnProperty.call(i,r)&&i[r]&&m.push(i[r][0]),i[r]=0;for(a in l)Object.prototype.hasOwnProperty.call(l,a)&&(e[a]=l[a]);u&&u(t);while(m.length)m.shift()();return s.push.apply(s,c||[]),n()}function n(){for(var e,t=0;t<s.length;t++){for(var n=s[t],a=!0,o=1;o<n.length;o++){var l=n[o];0!==i[l]&&(a=!1)}a&&(s.splice(t--,1),e=r(r.s=n[0]))}return e}var a={},i={app:0},s=[];function r(t){if(a[t])return a[t].exports;var n=a[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,r),n.l=!0,n.exports}r.m=e,r.c=a,r.d=function(e,t,n){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},r.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)r.d(n,a,function(t){return e[t]}.bind(null,a));return n},r.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="/RobotResearcher/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],l=o.push.bind(o);o.push=t,o=o.slice();for(var c=0;c<o.length;c++)t(o[c]);var u=l;s.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d76")},"35c1":function(e,t,n){},4678:function(e,t,n){var a={"./af":"2bfb","./af.js":"2bfb","./ar":"8e73","./ar-dz":"a356","./ar-dz.js":"a356","./ar-kw":"423e","./ar-kw.js":"423e","./ar-ly":"1cfd","./ar-ly.js":"1cfd","./ar-ma":"0a84","./ar-ma.js":"0a84","./ar-sa":"8230","./ar-sa.js":"8230","./ar-tn":"6d83","./ar-tn.js":"6d83","./ar.js":"8e73","./az":"485c","./az.js":"485c","./be":"1fc1","./be.js":"1fc1","./bg":"84aa","./bg.js":"84aa","./bm":"a7fa","./bm.js":"a7fa","./bn":"9043","./bn-bd":"9686","./bn-bd.js":"9686","./bn.js":"9043","./bo":"d26a","./bo.js":"d26a","./br":"6887","./br.js":"6887","./bs":"2554","./bs.js":"2554","./ca":"d716","./ca.js":"d716","./cs":"3c0d","./cs.js":"3c0d","./cv":"03ec","./cv.js":"03ec","./cy":"9797","./cy.js":"9797","./da":"0f14","./da.js":"0f14","./de":"b469","./de-at":"b3eb","./de-at.js":"b3eb","./de-ch":"bb71","./de-ch.js":"bb71","./de.js":"b469","./dv":"598a","./dv.js":"598a","./el":"8d47","./el.js":"8d47","./en-au":"0e6b","./en-au.js":"0e6b","./en-ca":"3886","./en-ca.js":"3886","./en-gb":"39a6","./en-gb.js":"39a6","./en-ie":"e1d3","./en-ie.js":"e1d3","./en-il":"7333","./en-il.js":"7333","./en-in":"ec2e","./en-in.js":"ec2e","./en-nz":"6f50","./en-nz.js":"6f50","./en-sg":"b7e9","./en-sg.js":"b7e9","./eo":"65db","./eo.js":"65db","./es":"898b","./es-do":"0a3c","./es-do.js":"0a3c","./es-mx":"b5b7","./es-mx.js":"b5b7","./es-us":"55c9","./es-us.js":"55c9","./es.js":"898b","./et":"ec18","./et.js":"ec18","./eu":"0ff2","./eu.js":"0ff2","./fa":"8df4","./fa.js":"8df4","./fi":"81e9","./fi.js":"81e9","./fil":"d69a","./fil.js":"d69a","./fo":"0721","./fo.js":"0721","./fr":"9f26","./fr-ca":"d9f8","./fr-ca.js":"d9f8","./fr-ch":"0e49","./fr-ch.js":"0e49","./fr.js":"9f26","./fy":"7118","./fy.js":"7118","./ga":"5120","./ga.js":"5120","./gd":"f6b4","./gd.js":"f6b4","./gl":"8840","./gl.js":"8840","./gom-deva":"aaf2","./gom-deva.js":"aaf2","./gom-latn":"0caa","./gom-latn.js":"0caa","./gu":"e0c5","./gu.js":"e0c5","./he":"c7aa","./he.js":"c7aa","./hi":"dc4d","./hi.js":"dc4d","./hr":"4ba9","./hr.js":"4ba9","./hu":"5b14","./hu.js":"5b14","./hy-am":"d6b6","./hy-am.js":"d6b6","./id":"5038","./id.js":"5038","./is":"0558","./is.js":"0558","./it":"6e98","./it-ch":"6f12","./it-ch.js":"6f12","./it.js":"6e98","./ja":"079e","./ja.js":"079e","./jv":"b540","./jv.js":"b540","./ka":"201b","./ka.js":"201b","./kk":"6d79","./kk.js":"6d79","./km":"e81d","./km.js":"e81d","./kn":"3e92","./kn.js":"3e92","./ko":"22f8","./ko.js":"22f8","./ku":"2421","./ku.js":"2421","./ky":"9609","./ky.js":"9609","./lb":"440c","./lb.js":"440c","./lo":"b29d","./lo.js":"b29d","./lt":"26f9","./lt.js":"26f9","./lv":"b97c","./lv.js":"b97c","./me":"293c","./me.js":"293c","./mi":"688b","./mi.js":"688b","./mk":"6909","./mk.js":"6909","./ml":"02fb","./ml.js":"02fb","./mn":"958b","./mn.js":"958b","./mr":"39bd","./mr.js":"39bd","./ms":"ebe4","./ms-my":"6403","./ms-my.js":"6403","./ms.js":"ebe4","./mt":"1b45","./mt.js":"1b45","./my":"8689","./my.js":"8689","./nb":"6ce3","./nb.js":"6ce3","./ne":"3a39","./ne.js":"3a39","./nl":"facd","./nl-be":"db29","./nl-be.js":"db29","./nl.js":"facd","./nn":"b84c","./nn.js":"b84c","./oc-lnc":"167b","./oc-lnc.js":"167b","./pa-in":"f3ff","./pa-in.js":"f3ff","./pl":"8d57","./pl.js":"8d57","./pt":"f260","./pt-br":"d2d4","./pt-br.js":"d2d4","./pt.js":"f260","./ro":"972c","./ro.js":"972c","./ru":"957c","./ru.js":"957c","./sd":"6784","./sd.js":"6784","./se":"ffff","./se.js":"ffff","./si":"eda5","./si.js":"eda5","./sk":"7be6","./sk.js":"7be6","./sl":"8155","./sl.js":"8155","./sq":"c8f3","./sq.js":"c8f3","./sr":"cf1e","./sr-cyrl":"13e9","./sr-cyrl.js":"13e9","./sr.js":"cf1e","./ss":"52bd","./ss.js":"52bd","./sv":"5fbd","./sv.js":"5fbd","./sw":"74dc","./sw.js":"74dc","./ta":"3de5","./ta.js":"3de5","./te":"5cbb","./te.js":"5cbb","./tet":"576c","./tet.js":"576c","./tg":"3b1b","./tg.js":"3b1b","./th":"10e8","./th.js":"10e8","./tk":"5aff","./tk.js":"5aff","./tl-ph":"0f38","./tl-ph.js":"0f38","./tlh":"cf75","./tlh.js":"cf75","./tr":"0e81","./tr.js":"0e81","./tzl":"cf51","./tzl.js":"cf51","./tzm":"c109","./tzm-latn":"b53d","./tzm-latn.js":"b53d","./tzm.js":"c109","./ug-cn":"6117","./ug-cn.js":"6117","./uk":"ada2","./uk.js":"ada2","./ur":"5294","./ur.js":"5294","./uz":"2e8c","./uz-latn":"010e","./uz-latn.js":"010e","./uz.js":"2e8c","./vi":"2921","./vi.js":"2921","./x-pseudo":"fd7e","./x-pseudo.js":"fd7e","./yo":"7f33","./yo.js":"7f33","./zh-cn":"5c3a","./zh-cn.js":"5c3a","./zh-hk":"49ab","./zh-hk.js":"49ab","./zh-mo":"3a6c","./zh-mo.js":"3a6c","./zh-tw":"90ea","./zh-tw.js":"90ea"};function i(e){var t=s(e);return n(t)}function s(e){if(!n.o(a,e)){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t}return a[e]}i.keys=function(){return Object.keys(a)},i.resolve=s,e.exports=i,i.id="4678"},"56d7":function(e,t,n){},"56d76":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var a=n("2b0e"),i=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-app",[e.$demoMode||e.$auth.ready()?n("v-main",[n("router-view")],1):n("v-main",[e._v(" Loading... ")])],1)},s=[],r={name:"App"},o=r,l=n("2877"),c=n("6544"),u=n.n(c),d=n("7496"),m=n("f6c4"),f=Object(l["a"])(o,i,s,!1,null,"fa4c820e",null),p=f.exports;u()(f,{VApp:d["a"],VMain:m["a"]});var h=n("8c4f"),v=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-container",{staticClass:"fill-height"},[n("v-row",{attrs:{justify:"center",align:"center"}},[n("v-form",{attrs:{align:"right"},on:{keydown:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.login(t)},submit:function(e){e.preventDefault(),e.stopPropagation()}}},[n("v-text-field",{attrs:{"prepend-icon":"mdi-console-line",label:"Code","error-messages":e.errorMessage,disabled:e.$demoMode},model:{value:e.code,callback:function(t){e.code=t},expression:"code"}}),n("v-checkbox",{attrs:{label:"I accept that this site stores a cookie on my computer. The cookie is used to anonymously identify the submission."},model:{value:e.cookie,callback:function(t){e.cookie=t},expression:"cookie"}}),n("v-btn",{attrs:{disabled:e.processing||""===e.code||!e.cookie,loading:e.processing,type:"submit"},on:{click:function(t){return t.preventDefault(),t.stopPropagation(),e.login(t)}}},[e._v("Enter")])],1)],1)],1)},b=[],g={name:"Join",data:function(){return{code:this.$demoMode?"demo":"",cookie:!1,rememberMe:!0,processing:!1,errorMessage:""}},methods:{login:function(){var e=this;if(this.$demoMode)return this.$router.push("/intro");this.processing||""===this.code||(this.processing=!0,this.$auth.login({data:{code:this.code},rememberMe:!0,redirect:"/intro",fetchUser:!0}).then((function(t){e.processing=!1})).catch((function(t){e.processing=!1,t.response&&401===t.response.status?(e.errorMessage="Invalid Code",e.code=""):(e.errorMessage="Something went wrong... Contact organizer.",e.code="")})))}},created:function(){this.$demoMode||this.$auth.check()&&this.$router.push("/intro")}},y=g,k=(n("9b63"),n("8336")),w=n("ac7c"),x=n("a523"),j=n("4bd4"),_=n("0fd9"),C=n("8654"),$=Object(l["a"])(y,v,b,!1,null,"3b0d4962",null),F=$.exports;u()($,{VBtn:k["a"],VCheckbox:w["a"],VContainer:x["a"],VForm:j["a"],VRow:_["a"],VTextField:C["a"]});var S=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-container",{staticClass:"fill-height"},[n("v-row",{attrs:{justify:"center",align:"center"}},[n("v-card",{attrs:{"max-width":"700px"}},[n("v-card-title",[e._v("Intro")]),n("v-card-text",[n("h3",[e._v(" Welcome and thanks for participating in our experiment. ")]),n("p",[e._v(" The experiment aims to investigate the evolution of software projects by asking the participants to complete tasks in preexisting projects. There will be two (2) tasks to be performed followed by a short survey. You will also be asked to answer some background questions before you start. Your participation is voluntary and you are free to drop out of the experiment at any time. "),n("strong",[e._v("We do greatly appreciate if you complete the whole experiment")]),e._v(" so that we may acquire some useful results. ")]),n("p",[e._v(" This experiment is part of a master thesis by William Leven ("),n("a",{attrs:{target:"_blank",href:"mailto:william@leven.id"}},[e._v("william@leven.id")]),e._v(") and Hampus Broman ("),n("a",{attrs:{target:"_blank",href:"mailto:bromanh@student.chalmers.se"}},[e._v("bromanh@student.chalmers.se")]),e._v(") studying at the Software Engineering master program at Chalmers University. The supervisors are Terese Besker and Richard Torkar. ")]),n("p",[e._v(" If you want to receive a copy of the thesis after its completion you may provide your email at: "),e.$demoMode?n("span",[e._v("DEMO MODE")]):n("a",{attrs:{target:"_blank",href:"https://forms.gle/WJ3kXBfBYjCV35E89"}},[e._v("https://forms.gle/WJ3kXBfBYjCV35E89")])]),n("v-divider"),n("h3",[e._v("Confidentiality agreement")]),n("p",[e._v(" During the experiment we will collect your solutions to the tasks, your completion time, and survey answers for further analysis to investigate the evolution of software projects. Responsible for the collection are William Leven ("),n("a",{attrs:{target:"_blank",href:"mailto:william@leven.id"}},[e._v("william@leven.id")]),e._v(") and Hampus Broman ("),n("a",{attrs:{target:"_blank",href:"mailto:bromanh@student.chalmers.se"}},[e._v("bromanh@student.chalmers.se")]),e._v("). ")]),n("p",[e._v(" The data will be continuously collected for the whole duration of the experiment and will be stored even if the experiment isn't completed. The data we collect from you will also be enriched by data about the place of your employment. After the completion of the thesis, the collected (anonymized) data will be made publicly available. ")]),n("p",[e._v(" None of the collected data can, by its nature, uniquely identify you as an individual. In those cases where the submitted data contains any uniquely identifying information, the dataset will be anonymized before further processing or publication to ensure confidentiality. ")]),e.$demoMode?n("p",[e._v(" Demo mode! No data will be stored. ")]):e._e(),n("v-checkbox",{attrs:{label:"I accept that my data is collected and processed in accordance to the agreement above."},model:{value:e.accepted,callback:function(t){e.accepted=t},expression:"accepted"}}),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"primary",disabled:!e.accepted},on:{click:e.submit}},[e._v("Start")])],1)],1)],1)],1)],1)},V=[],D={name:"Intro",data:function(){return{accepted:!1}},methods:{submit:function(){var e=this;if(!0===this.accepted){if(this.$demoMode)return this.$router.push("/experiment/0");this.axios.post("agreement/accept").then((function(){e.$router.push("/experiment/0")}))}}},mounted:function(){var e=this;this.$demoMode||this.$auth.fetch().then((function(t){switch(t.data.stage){case"agreement":break;case"background":e.$router.push("/experiment/0");break;case"experiment":e.$router.push("/experiment/"+t.data.experiment);break;case"survey":e.$router.push("/experiment/3");break;default:e.$router.push("/farewell");break}}))}},O=D,M=n("b0af"),E=n("99d9"),N=n("ce7e"),T=n("2fa4"),A=Object(l["a"])(O,S,V,!1,null,"a90a39a2",null),z=A.exports;u()(A,{VBtn:k["a"],VCard:M["a"],VCardActions:E["a"],VCardText:E["b"],VCardTitle:E["c"],VCheckbox:w["a"],VContainer:x["a"],VDivider:N["a"],VRow:_["a"],VSpacer:T["a"]});var P=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-stepper",{attrs:{value:e.step}},[n("v-stepper-header",[n("v-stepper-step",{attrs:{step:0}},[e._v(" Background info ")]),e._l(e.scenarios,(function(t){return n("v-stepper-step",{key:t,attrs:{step:t}},[e._v(" "+e._s(e.scenarioList[t-1])+" Task ")])})),n("v-stepper-step",{attrs:{step:e.scenarios+1}},[e._v(" Final (4) Questions ")])],2),n("v-stepper-items",[e.isSurvey?n("Survey"):e.isBackground?n("Background",[e._v("Background questions ")]):n("Scenario",{attrs:{number:e.step}})],1)],1)},R=[],B=(n("d81d"),n("fb6a"),n("a9e3"),function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-row",{staticClass:"fill-height",attrs:{justify:"center",align:"center"}},[n("v-card",{staticClass:"pa-4"},[n("v-form",{on:{submit:function(e){e.preventDefault(),e.stopPropagation()}}},[n("v-select",{attrs:{items:e.edLevels,label:"What is the highest degree or level of education you have completed?"},model:{value:e.surveyAnswers.EducationLevel,callback:function(t){e.$set(e.surveyAnswers,"EducationLevel",t)},expression:"surveyAnswers.EducationLevel"}}),n("v-text-field",{attrs:{label:"In what field did you attain that level of education?",placeholder:"e.g. computer science, electrical engineering, software engineering..."},model:{value:e.surveyAnswers.EducationField,callback:function(t){e.$set(e.surveyAnswers,"EducationField",t)},expression:"surveyAnswers.EducationField"}}),n("v-text-field",{attrs:{label:"How many years of professional programming experience do you have?",placeholder:"e.g. 1, 2, 3...",suffix:"years"},model:{value:e.surveyAnswers.ProgrammingExperience,callback:function(t){e.$set(e.surveyAnswers,"ProgrammingExperience",t)},expression:"surveyAnswers.ProgrammingExperience"}}),n("v-text-field",{attrs:{label:"How many years of professional experience do you have working with java?",placeholder:"e.g. 1, 2, 3...",suffix:"years"},model:{value:e.surveyAnswers.JavaExperience,callback:function(t){e.$set(e.surveyAnswers,"JavaExperience",t)},expression:"surveyAnswers.JavaExperience"}}),n("v-text-field",{attrs:{label:"Within what domain did you attain the most work-experience?",placeholder:"e.g. telecom, game development, automotive..."},model:{value:e.surveyAnswers.WorkDomain,callback:function(t){e.$set(e.surveyAnswers,"WorkDomain",t)},expression:"surveyAnswers.WorkDomain"}}),n("v-divider"),n("span",[e._v("My most recent place of employment...")]),n("v-select",{attrs:{items:[{text:"Yes",value:!0},{text:"No",value:!1}],label:"used peer code reviews."},model:{value:e.surveyAnswers.CompanyUsesCodeReviews,callback:function(t){e.$set(e.surveyAnswers,"CompanyUsesCodeReviews",t)},expression:"surveyAnswers.CompanyUsesCodeReviews"}}),n("v-select",{attrs:{items:[{text:"Yes",value:!0},{text:"No",value:!1}],label:"used pair programing."},model:{value:e.surveyAnswers.CompanyUsesPairProgramming,callback:function(t){e.$set(e.surveyAnswers,"CompanyUsesPairProgramming",t)},expression:"surveyAnswers.CompanyUsesPairProgramming"}}),n("v-select",{attrs:{items:[{text:"Yes",value:!0},{text:"No",value:!1}],label:"tracked technical debt."},model:{value:e.surveyAnswers.CompanyTracksTechnicalDebt,callback:function(t){e.$set(e.surveyAnswers,"CompanyTracksTechnicalDebt",t)},expression:"surveyAnswers.CompanyTracksTechnicalDebt"}}),n("v-select",{attrs:{items:[{text:"Yes",value:!0},{text:"No",value:!1}],label:"had established coding standards."},model:{value:e.surveyAnswers.CompanyHasCodingStandards,callback:function(t){e.$set(e.surveyAnswers,"CompanyHasCodingStandards",t)},expression:"surveyAnswers.CompanyHasCodingStandards"}}),n("v-row",[n("v-spacer"),n("v-btn",{attrs:{color:"primary",disabled:e.loading||e.invalid},on:{click:function(t){return t.preventDefault(),t.stopPropagation(),e.submit(t)}}},[e._v("Submit")])],1)],1)],1)],1)}),L=[],I=(n("45fc"),n("07ac"),{name:"Survey",data:function(){return{loading:!1,surveyAnswers:{EducationLevel:null,EducationField:null,ProgrammingExperience:null,JavaExperience:null,WorkDomain:null,CompanyUsesCodeReviews:null,CompanyUsesPairProgramming:null,CompanyTracksTechnicalDebt:null,CompanyHasCodingStandards:null},edLevels:["None","Some bachelor studies","Bachelor degree","Some master studies","Master degree","Some Ph. D. studies","Ph. D."]}},computed:{invalid:function(){return Object.values(this.surveyAnswers).some((function(e){return null===e||e===[]}))}},methods:{submit:function(){var e=this;if(this.$demoMode)return this.$router.push("/experiment/1");this.loading||(this.loading=!0,this.axios.post("background",this.surveyAnswers).then((function(){e.loading=!1,e.$router.push("/experiment/1")})).catch((function(){e.loading=!1})))}},mounted:function(){var e=this;this.$demoMode||this.$auth.fetch().then((function(t){switch(t.data.stage){case"agreement":e.$router.push("/intro");break;case"background":break;case"experiment":e.$router.push("/experiment/"+t.data.experiment);break;case"survey":e.$router.push("/experiment/3");break;default:e.$router.push("/farewell");break}}))}}),U=I,W=(n("8a2d"),n("b974")),H=Object(l["a"])(U,B,L,!1,null,"1cb849cc",null),q=H.exports;u()(H,{VBtn:k["a"],VCard:M["a"],VDivider:N["a"],VForm:j["a"],VRow:_["a"],VSelect:W["a"],VSpacer:T["a"],VTextField:C["a"]});var Q=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-row",{staticClass:"root"},[n("v-col",{staticClass:"code fill-height"},[n("multi-file-editor",{attrs:{files:e.files},on:{"update:file":function(t){var n=t.file,a=t.content;return e.files[n]=a},"delete:file":function(t){var n=t.file;return delete e.files[n]},save:e.run}})],1),n("v-col",{staticClass:"desc-toggler",on:{click:function(t){e.utilVisible=!e.utilVisible}}},[e.utilVisible?n("v-icon",[e._v("mdi-arrow-expand-right")]):n("v-icon",[e._v("mdi-arrow-expand-left")])],1),n("v-col",{staticClass:"util fill-height d-flex justify-space-between flex-column",class:{hidden:!e.utilVisible}},[n("div",{domProps:{innerHTML:e._s(e.htmlinfo)}}),n("div",[n("h4",[e._v("Run output:")]),n("v-progress-linear",{attrs:{indeterminate:"",color:"blue",active:e.loading}}),n("Output",{attrs:{code:e.runOutput+"Exit code: "+e.runCode}})],1),n("div",{staticClass:"d-flex justify-space-between"},[n("v-btn",{attrs:{color:"success",disabled:e.loading,width:"200px"},on:{click:e.run}},[e._v(" Run Code ")]),n("v-spacer"),n("v-dialog",{attrs:{"max-width":"600px"},on:{keydown:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.next(t)}},scopedSlots:e._u([{key:"activator",fn:function(t){var a=t.on;return[n("v-btn",e._g({attrs:{color:"primary",width:"200px",disabled:e.loading||e.nextDialog}},a),[e._v(" Submit ")])]}}]),model:{value:e.nextDialog,callback:function(t){e.nextDialog=t},expression:"nextDialog"}},[n("v-card",[n("v-card-title",[e._v("Continue")]),n("v-card-text",[e._v("You can not go back once you have submitted your solution")]),n("v-card-actions",[n("v-btn",{attrs:{disabled:e.loading},on:{click:function(t){e.nextDialog=!1}}},[e._v(" Go back ")]),n("v-spacer"),n("v-btn",{attrs:{color:"primary",disabled:e.loading},on:{click:e.next}},[e._v(" Submit ")])],1)],1)],1)],1)])],1)},Y=[],J=(n("4160"),n("4fad"),n("159b"),n("3835")),X=n("c1df"),G=n.n(X),Z=n("0e54"),K=n.n(Z),ee=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-row",{staticClass:"fill-height"},[n("v-col",{staticClass:"list-col fill-height"},[n("v-list",{attrs:{dense:""}},[n("div",{staticClass:"helper"},[e._v(" Right click on files to add new, rename and delete. Refreshing the page will reset the whole task. ")]),n("v-list-item-group",{attrs:{color:"primary",mandatory:""},model:{value:e.selected,callback:function(t){e.selected=t},expression:"selected"}},e._l(Object.keys(e.files).sort(),(function(t){return n("v-list-item",{key:t,attrs:{value:t},on:{contextmenu:function(n){return e.openContextMenu(t,n)}}},[e._v(e._s(t))])})),1)],1),n("v-menu",{attrs:{absolute:"","position-x":e.contextMenuX,"position-y":e.contextMenuY},model:{value:e.contextMenuOpen,callback:function(t){e.contextMenuOpen=t},expression:"contextMenuOpen"}},[n("v-list",{staticClass:"ctx-menu"},[n("v-list-item",{on:{click:function(t){e.newFileDialog=!0}}},[n("v-list-item-title",[e._v("New File")])],1),"Main.java"!==e.contextMenuFile?n("v-list-item",{on:{click:function(t){e.pendingRename=e.contextMenuFile}}},[n("v-list-item-title",[e._v("Rename")])],1):e._e(),"Main.java"!==e.contextMenuFile?n("v-list-item",{on:{click:function(t){e.pendingDelete=e.contextMenuFile}}},[n("v-list-item-title",[e._v("Delete")])],1):e._e()],1)],1),n("v-dialog",{attrs:{"max-width":"600px"},model:{value:e.newFileDialog,callback:function(t){e.newFileDialog=t},expression:"newFileDialog"}},[n("v-card",[n("v-card-title",[e._v("New File")]),n("v-card-text",[e._v(" Create a new file "),n("v-form",{attrs:{align:"right"},on:{submit:function(e){e.preventDefault(),e.stopPropagation()}}},[n("v-text-field",{ref:"newFileInput",attrs:{label:"File name",suffix:".java","error-messages":e.newFileError},on:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:(t.preventDefault(),t.stopPropagation(),e.addFile())}},model:{value:e.newFileName,callback:function(t){e.newFileName=t},expression:"newFileName"}}),n("v-btn",{attrs:{disabled:null!=e.newFileError||""===e.newFileName,color:"primary"},on:{click:function(t){return t.preventDefault(),t.stopPropagation(),e.addFile()}}},[e._v("Create")])],1)],1)],1)],1),n("v-dialog",{attrs:{"max-width":"600px"},model:{value:e.renameDialogOpen,callback:function(t){e.renameDialogOpen=t},expression:"renameDialogOpen"}},[n("v-card",[n("v-card-title",[e._v("Rename File")]),n("v-card-text",[e._v(" Rename "+e._s(e.pendingRename)+" "),n("v-form",{attrs:{align:"right"},on:{submit:function(e){e.preventDefault(),e.stopPropagation()}}},[n("v-text-field",{ref:"renameFileInput",attrs:{label:"New name",suffix:".java",autofocus:"","error-messages":e.renameFileError},on:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:(t.preventDefault(),t.stopPropagation(),e.renameFile())}},model:{value:e.renameFileName,callback:function(t){e.renameFileName=t},expression:"renameFileName"}}),n("v-btn",{attrs:{disabled:null!=e.renameFileError||""===e.renameFileName||e.normalizeFileName(e.renameFileName)===e.pendingRename,color:"primary"},on:{click:function(t){return t.preventDefault(),t.stopPropagation(),e.renameFile()}}},[e._v("Rename")])],1)],1)],1)],1),n("v-dialog",{attrs:{"max-width":"600px"},on:{keydown:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.deleteFile(t)}},model:{value:e.deleteDialogOpen,callback:function(t){e.deleteDialogOpen=t},expression:"deleteDialogOpen"}},[n("v-card",[n("v-card-title",[e._v("Delete file")]),n("v-card-text",[e._v(" Do you want to delete "+e._s(e.pendingDelete)+"? "),n("v-card-actions",[n("v-btn",{on:{click:function(t){e.deleteDialogOpen=!1}}},[e._v("Cancel")]),n("v-spacer"),n("v-btn",{attrs:{color:"error"},on:{click:e.deleteFile}},[e._v("Delete")])],1)],1)],1)],1)],1),n("v-col",[n("editor",{attrs:{file:{name:e.selected,code:e.files[e.selected]}},on:{"update:code":e.edited,save:function(t){return e.$emit("save")}}})],1)],1)},te=[],ne=(n("c975"),n("b64b"),function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"editor",on:{keydown:function(t){return(t.type.indexOf("key")||83===t.keyCode)&&t.ctrlKey?(t.preventDefault(),t.stopPropagation(),e.$emit("save")):null}}},[n("div",{attrs:{id:e._uid}})])}),ae=[],ie=(n("b0c0"),n("6d4f")),se=n.n(ie),re=(n("65d8"),{name:"Editor",data:function(){return{editor:null,editorOptions:{maxLines:1e3,minLines:50,printMargin:!1,value:"",mode:"ace/mode/java",theme:"ace/theme/monokai",fontSize:"12pt",tabSize:"4",useSoftTabs:!1,autoScrollEditorIntoView:!1,scrollPastEnd:!1},editorSessions:{},currentSession:"none"}},props:{file:Object},watch:{file:{handler:function(e){var t=e.name,n=e.code;null!=this.editor&&this.editor.getValue()!==n&&(this.editor.setValue(n),this.editor.clearSelection(),t!==this.currentSession&&this.editor.getSession().getUndoManager().reset())},deep:!0}},mounted:function(){var e=this;se.a.require("ace/ext/language_tools"),this.editor=se.a.edit(String(this._uid),this.editorOptions),this.editor.on("change",(function(){e.$emit("update:code",e.editor.getValue())}))}}),oe=re,le=(n("dfde"),Object(l["a"])(oe,ne,ae,!1,null,"0b89f51a",null)),ce=le.exports,ue="Main.java",de={name:"MultiFileEditor",components:{Editor:ce},data:function(){return{selected:ue,contextMenuFile:null,contextMenuX:0,contextMenuY:0,newFileDialog:!1,newFileName:"",renameFileName:"",pendingDelete:null,pendingRename:null}},props:{files:Object},computed:{newFileError:function(){return""===this.newFileName?null:this.fileNameErrors(this.normalizeFileName(this.newFileName))},renameFileError:function(){return""===this.renameFileName||this.normalizeFileName(this.renameFileName)===this.pendingRename?null:this.fileNameErrors(this.normalizeFileName(this.renameFileName))},deleteDialogOpen:{get:function(){return null!==this.pendingDelete},set:function(e){!1===e&&(this.pendingDelete=null)}},renameDialogOpen:{get:function(){return null!==this.pendingRename},set:function(e){!1===e&&(this.pendingRename=null)}},contextMenuOpen:{get:function(){return null!==this.contextMenuFile},set:function(e){!1===e&&(this.contextMenuFile=null)}}},watch:{files:function(){void 0===this.files[this.selected]&&(this.selected=ue)},pendingRename:function(e){var t=this;null!=e&&this.$nextTick((function(){setTimeout((function(){t.$refs.renameFileInput.focus()}),0)}))},newFileDialog:function(e){var t=this;e&&this.$nextTick((function(){setTimeout((function(){t.$refs.newFileInput.focus()}),0)}))}},methods:{renameFile:function(){this.$emit("update:file",{file:this.normalizeFileName(this.renameFileName),content:this.files[this.pendingRename]}),this.$emit("delete:file",{file:this.pendingRename}),this.selected=this.normalizeFileName(this.renameFileName),this.renameFileName="",this.pendingRename=null},deleteFile:function(){this.$emit("delete:file",{file:this.pendingDelete}),this.pendingDelete=null,this.selected=ue},normalizeFileName:function(e){return-1===e.indexOf(".java",e.length-".java".length)&&(e+=".java"),e[0].toUpperCase()+e.substring(1)},fileNameErrors:function(e){return-1===e.indexOf(".java",e.length-".java".length)?"only java files allowed":e.indexOf(".")!==e.length-".java".length?"dots are not allowed in file name":Object.keys(this.files).some((function(t){return t===e}))?"that file already exists":null},addFile:function(){this.newFileDialog=!1,this.$emit("update:file",{file:this.normalizeFileName(this.newFileName),content:"\n"}),this.selected=this.normalizeFileName(this.newFileName),this.newFileName=""},edited:function(e){this.$emit("update:file",{file:this.selected,content:e})},openContextMenu:function(e,t){t.preventDefault(),this.contextMenuX=t.clientX,this.contextMenuY=t.clientY,this.contextMenuFile=e}}},me=de,fe=(n("84ce"),n("62ad")),pe=n("169a"),he=n("8860"),ve=n("da13"),be=n("1baa"),ge=n("5d23"),ye=n("e449"),ke=Object(l["a"])(me,ee,te,!1,null,"466da5f0",null),we=ke.exports;u()(ke,{VBtn:k["a"],VCard:M["a"],VCardActions:E["a"],VCardText:E["b"],VCardTitle:E["c"],VCol:fe["a"],VDialog:pe["a"],VForm:j["a"],VList:he["a"],VListItem:ve["a"],VListItemGroup:be["a"],VListItemTitle:ge["b"],VMenu:ye["a"],VRow:_["a"],VSpacer:T["a"],VTextField:C["a"]});var xe=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:e._uid}})},je=[],_e={name:"Output",data:function(){return{editor:null,editorOptions:{maxLines:15,minLines:5,printMargin:!1,value:this.code,theme:"ace/theme/monokai",fontSize:"12pt",tabSize:"2",useSoftTabs:!1,autoScrollEditorIntoView:!1,scrollPastEnd:!1,readOnly:!0}}},props:{code:String},watch:{code:function(e){null!=this.editor&&this.editor.getValue()!==e&&(this.editor.setValue(e),this.editor.clearSelection())}},mounted:function(){se.a.require("ace/ext/language_tools"),this.editor=se.a.edit(String(this._uid),this.editorOptions)}},Ce=_e,$e=Object(l["a"])(Ce,xe,je,!1,null,"6f20921d",null),Fe=$e.exports,Se={name:"Scenario",data:function(){return{scenario:null,files:{},info:"",startedAt:null,loading:!1,runOutput:"",runCode:null,nextDialog:!1,utilVisible:!0}},props:{number:Number},components:{Output:Fe,MultiFileEditor:we},computed:{htmlinfo:function(){return K()(this.info)}},methods:{next:function(){var e=this;if(this.$demoMode)return this.nextDialog=!1,this.$router.push("/experiment/"+(Number(this.$route.params.number)+1));this.loading||(this.loading=!0,this.axios.post("scenario/"+(Number(this.$route.params.number)-1),{submission:this.prepareSubmission(this.files)}).then((function(){e.loading=!1,e.nextDialog=!1,e.$router.push("/experiment/"+(Number(e.$route.params.number)+1))})).catch((function(){e.loading=!1,e.nextDialog=!1})))},run:function(){var e=this;if(this.$demoMode)return this.runCode=-1,void(this.runOutput="No code execution in demo mode\n");this.loading||(this.loading=!0,this.axios.post("code/run",{submission:this.prepareSubmission(this.files)}).then((function(t){e.runOutput=t.data.out,e.runCode=t.data.code,e.loading=!1})).catch((function(t){t.response&&(e.runOutput=t.response.data.out,e.runCode=t.response.data.code,e.loading=!1)})))},prepareSubmission:function(e){var t={};return Object.entries(e).forEach((function(e){var n=Object(J["a"])(e,2),a=n[0],i=n[1];t[a]=btoa(i)})),t},loadScenario:function(){var e=this;this.loading=!0,this.axios.get("scenario/"+(Number(this.$route.params.number)-1)).then((function(t){e.scenario=t.data,e.info=e.scenario.info,e.files={},Object.entries(e.scenario.given).forEach((function(t){var n=Object(J["a"])(t,2),a=n[0],i=n[1];e.files[a]=atob(i)})),e.startedAt=G()(e.scenario.started_at),e.runOutput="",e.loading=!1})).catch((function(){e.loading=!1}))},loadDemoScenario:function(){this.info="# Demo task\nThis is just a `demo`",this.files={"Main.java":'public static class Main {\n\tpublic static main(String[] args) {\n\t\tSystem.out.println("Hello world!");\n\t}\n}',"Other.java":"another java class"},this.startedAt=G()(),this.runOutput=""}},mounted:function(){var e=this;this.$demoMode?this.loadDemoScenario():(this.$auth.fetch().then((function(t){switch(t.data.stage){case"agreement":e.$router.push("/intro");break;case"background":e.$router.push("/experiment/0");break;case"experiment":e.$route.params.number!==t.data.experiment&&e.$router.push("/experiment/"+t.data.experiment);break;case"survey":e.$router.push("/experiment/3");break;default:e.$router.push("/farewell");break}})),this.loadScenario())},watch:{number:function(e,t){e!==t&&(this.$demoMode?this.loadDemoScenario():this.loadScenario())}}},Ve=Se,De=(n("57b3"),n("132d")),Oe=n("8e36"),Me=Object(l["a"])(Ve,Q,Y,!1,null,"481cb38c",null),Ee=Me.exports;u()(Me,{VBtn:k["a"],VCard:M["a"],VCardActions:E["a"],VCardText:E["b"],VCardTitle:E["c"],VCol:fe["a"],VDialog:pe["a"],VIcon:De["a"],VProgressLinear:Oe["a"],VRow:_["a"],VSpacer:T["a"]});var Ne=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-row",{staticClass:"fill-height",attrs:{justify:"center",align:"center"}},[n("v-card",{staticClass:"pa-4"},[n("v-form",{on:{submit:function(e){e.preventDefault(),e.stopPropagation()}}},[e._l(e.scenarios,(function(t){return n("span",{key:t,staticClass:"sliders"},[n("span",[e._v("The quality (maintainability) of the preexisting code in the "),n("strong",[e._v(e._s(t))]),n("v-tooltip",{attrs:{right:""},scopedSlots:e._u([{key:"activator",fn:function(t){var a=t.on,i=t.attrs;return[n("v-icon",e._g(e._b({staticClass:"infoa",attrs:{tag:"span",small:""}},"v-icon",i,!1),a),[e._v("mdi-information")])]}}],null,!0)},[n("span",[e._v(e._s(e.scenarioHelp[t]))])]),e._v(" task was:")],1),n("v-slider",{attrs:{"tick-labels":["Very Bad","Bad","Somewhat Bad","Neutral","Somewhat Good","Good","Very Good"],ticks:"","tick-size":"7",step:"1",min:"-3",max:"3"},model:{value:e.surveyAnswers["Scenario"+t+"Quality"],callback:function(n){e.$set(e.surveyAnswers,"Scenario"+t+"Quality",n)},expression:"surveyAnswers['Scenario' + scenario + 'Quality']"}}),n("span",[e._v("The work I did in the "),n("strong",[e._v(e._s(t))]),n("v-tooltip",{attrs:{right:""},scopedSlots:e._u([{key:"activator",fn:function(t){var a=t.on,i=t.attrs;return[n("v-icon",e._g(e._b({staticClass:"infoa",attrs:{tag:"span",small:""}},"v-icon",i,!1),a),[e._v("mdi-information")])]}}],null,!0)},[n("span",[e._v(e._s(e.scenarioHelp[t]))])]),e._v(" task made the quality (maintainability) of the system:")],1),n("v-slider",{attrs:{"tick-labels":["Much Worse","Worse","Somewhat Worse","Neutral","Somewhat Better","Better","Much Better"],ticks:"","tick-size":"7",step:"1",min:"-3",max:"3"},model:{value:e.surveyAnswers["Scenario"+t+"QualityChange"],callback:function(n){e.$set(e.surveyAnswers,"Scenario"+t+"QualityChange",n)},expression:"surveyAnswers['Scenario' + scenario + 'QualityChange']"}})],1)})),n("v-row",[n("v-spacer"),n("v-btn",{attrs:{color:"primary",disabled:e.loading||e.invalid},on:{click:function(t){return t.stopPropagation(),t.preventDefault(),e.submit(t)}}},[e._v("Submit")])],1)],2)],1)],1)},Te=[],Ae={name:"Survey",data:function(){return{loading:!1,surveyAnswers:{ScenarioTicketsQuality:null,ScenarioTicketsQualityChange:null,ScenarioBookingQuality:null,ScenarioBookingQualityChange:null},scenarios:["Booking","Tickets"],scenarioHelp:{Tickets:"The task where you were asked to extend a buss ticket system",Booking:"The task where you were asked to implement a ComputerRoom class"}}},computed:{invalid:function(){return Object.values(this.surveyAnswers).some((function(e){return null===e||e===[]}))}},mounted:function(){var e=this;this.$demoMode||this.$auth.fetch().then((function(t){switch(a["a"].set(e,"scenarios",t.data.scenario_order.map((function(e){return e.charAt(0).toUpperCase()+e.slice(1)}))),t.data.stage){case"agreement":e.$router.push("/intro");break;case"background":e.$router.push("/experiment/0");break;case"experiment":e.$router.push("/experiment/"+t.data.experiment);break;case"survey":break;default:e.$router.push("/farewell");break}}))},methods:{submit:function(){var e=this;if(this.$demoMode)return this.$router.push("/farewell");this.loading||(this.loading=!0,this.axios.post("survey",this.surveyAnswers).then((function(){e.loading=!1,e.$router.push("/farewell")})).catch((function(){e.loading=!1})))}}},ze=Ae,Pe=(n("eb07"),n("ba0d")),Re=n("3a2f"),Be=Object(l["a"])(ze,Ne,Te,!1,null,"57f84091",null),Le=Be.exports;u()(Be,{VBtn:k["a"],VCard:M["a"],VForm:j["a"],VIcon:De["a"],VRow:_["a"],VSlider:Pe["a"],VSpacer:T["a"],VTooltip:Re["a"]});var Ie={name:"Experiment",components:{Scenario:Ee,Survey:Le,Background:q},data:function(){return{scenarios:2,scenarioList:this.$demoMode?["Booking","Tickets"]:["1","2"]}},computed:{step:function(){return Number(this.$route.params.number)},isSurvey:function(){return this.step===this.scenarios+1},isBackground:function(){return 0===this.step}},mounted:function(){var e=this;this.$demoMode||this.$auth.fetch().then((function(t){a["a"].set(e,"scenarioList",t.data.scenario_order.map((function(e){return e.charAt(0).toUpperCase()+e.slice(1)})))}))}},Ue=Ie,We=(n("8c15"),n("7e85")),He=n("9c54"),qe=n("56a4"),Qe=Object(l["a"])(Ue,P,R,!1,null,"d2719b20",null),Ye=Qe.exports;u()(Qe,{VStepper:We["a"],VStepperHeader:He["a"],VStepperItems:He["b"],VStepperStep:qe["a"]});var Je=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-row",{staticClass:"fill-height pa-16",attrs:{justify:"center",align:"center"}},[n("v-card",{attrs:{width:"650"}},[n("v-card-title",[e._v("Thank you for participating!")]),n("v-card-text",[n("p",[e._v(" If you want to receive a copy of the completed thesis or want to participate in a short follow up interview you can use the form bellow. ")])]),n("hr"),e.$demoMode?n("span",[e._v("No form in demo mode")]):n("iframe",{attrs:{src:"https://docs.google.com/forms/d/e/1FAIpQLSeCn-W2hhktZngqNB72XSI5a-d9N1awwuldXkXlm1gZyO7tVQ/viewform?embedded=true&hl=en",width:"640",height:"1100",seamless:"",frameborder:"0",marginheight:"0",marginwidth:"0"}},[e._v(" Läser in... ")])],1)],1)},Xe=[],Ge={name:"Farewell",mounted:function(){this.$demoMode||this.$auth.logout({makeRequest:!0,url:"auth/close",redirect:!1})}},Ze=Ge,Ke=Object(l["a"])(Ze,Je,Xe,!1,null,"48d84ad9",null),et=Ke.exports;u()(Ke,{VCard:M["a"],VCardText:E["b"],VCardTitle:E["c"],VRow:_["a"]}),a["a"].use(h["a"]);var tt=[{path:"/",name:"Join",component:F},{path:"/intro",name:"Intro",component:z,meta:{auth:!0}},{path:"/experiment/:number",name:"Experiment",component:Ye,meta:{auth:!0}},{path:"/farewell",name:"Farewell",component:et}],nt=new h["a"]({routes:tt});a["a"].router=nt;var at=nt,it=n("f309");a["a"].use(it["a"]);var st=new it["a"]({}),rt=n("bc3a"),ot=n.n(rt),lt=n("2106"),ct=n.n(lt);a["a"].use(ct.a,ot.a),a["a"].axios.defaults.baseURL="/api";var ut=n("dc21"),dt=n("e594"),mt=n("688d");ut["a"]({dsn:"https://c3812d572f7f47418eb2bc2dd5e32179@o408198.ingest.sentry.io/5516819",integrations:[new dt["a"]({Vue:a["a"],tracing:!0,logErrors:!0,attachProps:!0}),new mt["a"].BrowserTracing],environment:"production"}),a["a"].config.productionTip=!1,a["a"].prototype.$demoMode=!0,a["a"].prototype.$demoMode||n("c3a2"),new a["a"]({router:at,vuetify:st,render:function(e){return e(p)}}).$mount("#app")},"57b3":function(e,t,n){"use strict";n("5a07")},"5a07":function(e,t,n){},"84ce":function(e,t,n){"use strict";n("c0d7")},"8a2d":function(e,t,n){"use strict";n("d10d")},"8c15":function(e,t,n){"use strict";n("b628")},"9b63":function(e,t,n){"use strict";n("b057")},b057:function(e,t,n){},b628:function(e,t,n){},c0d7:function(e,t,n){},c3a2:function(e,t,n){"use strict";n.r(t);var a=n("2b0e"),i=n("4211"),s=n("a026"),r=n("680f"),o=n("4ae6");a["a"].use(i["a"],{auth:s["a"],http:r["a"],router:o["a"],refreshData:{enabled:!1},fetchData:{enabled:!0,url:"auth/status"},loginData:{url:"auth/register"},authRedirect:"/"})},d10d:function(e,t,n){},dfde:function(e,t,n){"use strict";n("35c1")},eb07:function(e,t,n){"use strict";n("56d7")}});
//# sourceMappingURL=app.c7dc3d11.js.map