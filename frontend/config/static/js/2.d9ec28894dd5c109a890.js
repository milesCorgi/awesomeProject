webpackJsonp([2],{Vfgl:function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var a=i("PJh5"),s=i.n(a),r=i("TQvf"),l=i.n(r),o=i("bFR2"),n={name:"queryFrom",filters:{infoTypeName:function(e){var t={"":"全部",0:"视频",1:"图片",2:"同人"};return t.hasOwnProperty(e)?t[e]:"其他"},dateFormat:function(e){return s()(e).add(8,"hours").format("YYYY-MM-DD HH:mm:ss")},showFullLink:function(e){return e.length<=27?e:"点击跳转或复制获取链接"}},data:function(){return{options:[],allOptions:[],isAdmin:!1,lenternUrl:"https://github.com/getlantern/lantern",loading:!0,filterPara:{keyword:"",Info:"",BiliBiliID:"",YoutubeLink:"",OtherLink:"",InfoType:""},totalList:[],showList:[],page:{totalNum:0,currentPage:1,pageSize:10},PagerCount:7,PagerIsSmall:!1,MobileIsPc:!0,user:{UserName:"",PassWord:""}}},created:function(){this.MobileIsPc=!this.isMobile(),this.MobileIsPc||(this.PagerCount=5,this.PagerIsSmall=!0),this.getKeyWords(),this.user.UserName=o.a.get("two-set-info-username"),console.log(this.user.UserName),void 0===this.user.UserName&&""===this.user.UserName||(this.IsLogin=!0),this.getInfo()},mounted:function(){"1"===o.a.get("two-set-info-isAdmin")?this.isAdmin=!0:this.isAdmin=!1},methods:{getKayWordsToShow:function(){this.options=[];for(var e=[],t=0;t<this.allOptions.length;t++)""===this.filterPara.InfoType?-1===e.indexOf(this.allOptions[t].value)&&(e.push(this.allOptions[t].value),this.options.push(this.allOptions[t])):this.allOptions[t].InfoType===this.filterPara.InfoType&&this.options.push(this.allOptions[t])},getKeyWords:function(){var e=this;this.$http.post("/api/get_keywords").then(function(t){var i=t.data;if(e.options=[],e.allOptions=[],0===i.error_num)for(var a=0;a<i.keyWord_list.length;a++)e.allOptions.push({value:i.keyWord_list[a].KeyWord,label:i.keyWord_list[a].KeyWord,InfoType:i.keyWord_list[a].InfoType});else e.$message.error(i.msg)})},infoTypeStyle:function(e){var t={0:"",1:"success",2:"warning"};return t.hasOwnProperty(e)?t[e]:"info"},settingShowImg:function(e,t){return"img"===t?0!==e.length&&this.verifyImg(e):"info"===t?0!==e.length&&!this.verifyImg(e):0!==e.length&&this.verifyImg(e)},notShowIfNull:function(e){return 0!==e.length},isMobile:function(){return navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)},searchReset:function(){this.paraReset(),this.getInfo()},paraReset:function(){this.filterPara.Info="",this.filterPara.BiliBiliID="",this.filterPara.YoutubeLink="",this.filterPara.OtherLink="",this.filterPara.InfoType="",this.filterPara.keyword=""},disableThisInfo:function(e){var t=this,i=!0;this.user.UserName=o.a.get("two-set-info-username"),void 0!==this.user.UserName&&""!==this.user.UserName||(i=!1,this.$message({type:"waring",message:"没有获取到用户名，请找二狗"})),this.user.PassWord=o.a.get("two-set-info-password"),void 0!==this.user.PassWord&&""!==this.user.PassWord||(i=!1,this.$message({type:"waring",message:"用户验证没通过，请找二狗"})),"1"!==o.a.get("two-set-info-isAdmin")&&(this.$message({type:"waring",message:"管理员检查没通过，请找二狗",duration:500}),i=!1,this.isAdmin=!1),i&&this.$http.post("/api/set_disable_info",{UserName:this.user.UserName,PassWord:this.user.PassWord,InfoId:e,Enable:"0"},{emulateJSON:!0}).then(function(e){console.log(e);var i=e.data;0===i.error_num?t.$message({type:"success",message:"屏蔽成功，刷新后生效",duration:500}):t.$message.error(i.msg)})},getInfo:function(){var e=this;this.loading=!0,this.page.totalNum=0,this.page.currentPage=1,this.$http.post("/api/query_two_set_video_info",{KeyWord:this.filterPara.keyword,Info:this.filterPara.Info,BiliBiliID:this.filterPara.BiliBiliID,YoutubeLink:this.filterPara.YoutubeLink,OtherLink:this.filterPara.OtherLink,InfoType:this.filterPara.InfoType}).then(function(t){console.log(t);var i=t.data;0===i.error_num?(e.totalList=i.info_list,e.showList=i.info_list,e.page.totalNum=i.total_num,e.updateCurrentPageInfo(),e.loading=!1,console.log("this.page.totalNum:",e.page.totalNum)):e.$message.error(i.msg)})},openUrl:function(e){/^(https?|ftp|file):\/\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]/.test(e)?window.open(e):this.$message({message:"url格式不合法",duration:500}),console.log(e)},handleSizeChange:function(e){console.log("每页 "+e+" 条"),this.page.pageSize=e,this.updateCurrentPageInfo()},handleCurrentChange:function(e){this.page.currentPage=e,console.log("当前页: "+e),this.updateCurrentPageInfo()},copy:function(){var e=new l.a(".tag-read");e.on("success",function(t){e.destroy()}),e.on("error",function(t){e.destroy()}),this.$message({type:"success",message:"已复制",duration:500})},verifyImg:function(e){if(void 0!==e&&""!==e){var t=e.split("."),i=t[t.length-1].toLowerCase();return"png"===i||"jpg"===i||"jpeg"===i||"gif"===i}return!1},updateCurrentPageInfo:function(){var e=(this.page.currentPage-1)*this.page.pageSize,t=this.page.currentPage*this.page.pageSize;t=t>this.page.totalNum?this.page.totalNum:t,this.showList=this.totalList.slice(e,t);for(var i=0;i<this.showList.length;i++){var a=this.showList[i].ImgLinks.split("\n");this.showList[i].ShowImgLinks=[];for(var s=[],r=0;r<a.length;r++)this.verifyImg(a[r])&&s.push(a[r]);for(var l=0;l<a.length;l++){var o=this.verifyImg(a[l]);this.showList[i].ShowImgLinks.push({enableShow:o,value:a[l],srcList:s})}}console.log(this.showList)}}},c={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",[i("p",{directives:[{name:"show",rawName:"v-show",value:e.isAdmin,expression:"isAdmin"}]},[e._v("你好,管理员"+e._s(e.user.UserName))]),e._v(" "),i("div",{staticStyle:{"margin-top":"10px"}},[i("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:e.lenternUrl,expression:"lenternUrl",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document"}),e._v("\n      或者你需要一个小手电:\n      "),i("a",{attrs:{href:"javascript:;"},on:{click:function(t){return e.openUrl(e.lenternUrl)}}},[e._v(e._s(e.lenternUrl))])]),e._v(" "),i("el-collapse",{staticStyle:{"border-color":"white"}},[i("el-collapse-item",{staticStyle:{width:"200px","margin-left":"30px","margin-right":"30px","margin-top":"10px"},attrs:{title:"高级搜索"}},[i("el-row",[i("el-row",{staticStyle:{"margin-top":"10px"},attrs:{type:"flex",justify:"center"}},[i("el-col",[i("el-input",{attrs:{clearable:"",placeholder:"输入描述关键词"},model:{value:e.filterPara.Info,callback:function(t){e.$set(e.filterPara,"Info",t)},expression:"filterPara.Info"}})],1)],1),e._v(" "),i("el-row",[i("el-col",[i("el-input",{staticStyle:{"margin-top":"10px"},attrs:{clearable:"",placeholder:"输入b站号"},model:{value:e.filterPara.BiliBiliID,callback:function(t){e.$set(e.filterPara,"BiliBiliID",t)},expression:"filterPara.BiliBiliID"}})],1)],1)],1),e._v(" "),i("el-row",[i("el-row",{staticStyle:{"margin-top":"10px"},attrs:{type:"flex",justify:"center"}},[i("el-col",[i("el-input",{attrs:{clearable:"",placeholder:"输入油管链接"},model:{value:e.filterPara.YoutubeLink,callback:function(t){e.$set(e.filterPara,"YoutubeLink",t)},expression:"filterPara.YoutubeLink"}})],1)],1),e._v(" "),i("el-row",[i("el-col",[i("el-input",{staticStyle:{"margin-top":"10px"},attrs:{clearable:"",placeholder:"输入其他链接"},model:{value:e.filterPara.OtherLink,callback:function(t){e.$set(e.filterPara,"OtherLink",t)},expression:"filterPara.OtherLink"}})],1)],1)],1)],1)],1),e._v(" "),i("el-row",[i("el-row",{staticStyle:{"margin-top":"10px"},attrs:{type:"flex",justify:"center"}},[i("span",{staticStyle:{"margin-top":"10px","font-size":"small"}},[e._v("类型：")]),e._v(" "),i("el-radio-group",{attrs:{size:"mini"},model:{value:e.filterPara.InfoType,callback:function(t){e.$set(e.filterPara,"InfoType",t)},expression:"filterPara.InfoType"}},[i("el-radio-button",{attrs:{label:""}},[e._v("全部")]),e._v(" "),i("el-radio-button",{attrs:{label:"0"}},[e._v("视频")]),e._v(" "),i("el-radio-button",{attrs:{label:"1"}},[e._v("图片")]),e._v(" "),i("el-radio-button",{attrs:{label:"2"}},[e._v("同人")])],1)],1)],1),e._v(" "),i("el-row",[i("el-row",{staticStyle:{"margin-top":"10px"},attrs:{type:"flex",justify:"center"}},[i("span",{staticStyle:{"margin-top":"10px","font-size":"small"}},[e._v("关键词：")]),e._v(" "),i("el-select",{attrs:{clearable:"",filterable:"",placeholder:"关键词"},on:{focus:function(t){return e.getKayWordsToShow()}},model:{value:e.filterPara.keyword,callback:function(t){e.$set(e.filterPara,"keyword",t)},expression:"filterPara.keyword"}},e._l(e.options,function(e){return i("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})}),1)],1)],1),e._v(" "),i("el-row",{attrs:{type:"flex",justify:"center"}},[i("el-button",{staticStyle:{"margin-top":"10px"},attrs:{type:"primary"},on:{click:e.getInfo}},[e._v("查询")]),e._v(" "),i("el-button",{staticStyle:{"margin-top":"10px","margin-right":"7px"},on:{click:e.searchReset}},[e._v("重置")])],1),e._v(" "),i("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],staticStyle:{width:"100%"},attrs:{data:e.showList,"header-align":"center"}},[i("el-table-column",{attrs:{type:"expand",fixed:"left",label:"详情",width:"80"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("el-form",{staticClass:"demo-table-expand",attrs:{"label-position":"left",inline:""}},[i("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.notShowIfNull(t.row.User),expression:"notShowIfNull(props.row.User)"}],attrs:{label:"提交人"}},[i("span",[e._v(e._s(t.row.User))])]),e._v(" "),i("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.notShowIfNull(t.row.TheTime),expression:"notShowIfNull(props.row.TheTime)"}],attrs:{label:"在该视频的哪个时段"}},[i("span",[e._v(e._s(t.row.TheTime))])]),e._v(" "),i("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.notShowIfNull(t.row.BiliBiliLink),expression:"notShowIfNull(props.row.BiliBiliLink)"}],attrs:{label:"b站链接"}},[i("div",{staticStyle:{display:"inline-block","text-align":"center"},attrs:{slot:"label"},slot:"label"},[i("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.row.BiliBiliLink,expression:"props.row.BiliBiliLink",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}}),e._v(" "),i("span",[e._v("b站链接")])]),e._v(" "),i("a",{attrs:{href:"javascript:;"},on:{click:function(i){return e.openUrl(t.row.BiliBiliLink)}}},[e._v("\n                "+e._s(e._f("showFullLink")(t.row.BiliBiliLink))+"\n              ")])]),e._v(" "),i("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.notShowIfNull(t.row.BiliBiliID),expression:"notShowIfNull(props.row.BiliBiliID)"}],attrs:{label:""}},[i("div",{staticStyle:{display:"inline-block","text-align":"center"},attrs:{slot:"label"},slot:"label"},[i("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.row.BiliBiliID,expression:"props.row.BiliBiliID",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}}),e._v(" "),i("span",[e._v("b站号")])]),e._v(" "),i("a",{attrs:{href:"javascript:;"},on:{click:function(i){return e.openUrl(t.row.BiliBiliLink)}}},[e._v("\n                "+e._s(t.row.BiliBiliID)+"\n              ")])]),e._v(" "),i("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.notShowIfNull(t.row.YoutubeLink),expression:"notShowIfNull(props.row.YoutubeLink)"}],attrs:{label:"油管链接"}},[i("div",{staticStyle:{display:"inline-block","text-align":"center"},attrs:{slot:"label"},slot:"label"},[i("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.row.YoutubeLink,expression:"props.row.YoutubeLink",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}}),e._v(" "),i("span",[e._v("油管链接")])]),e._v(" "),i("a",{attrs:{href:"javascript:;"},on:{click:function(i){return e.openUrl(t.row.YoutubeLink)}}},[e._v("\n                "+e._s(e._f("showFullLink")(t.row.YoutubeLink))+"\n              ")])]),e._v(" "),i("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.notShowIfNull(t.row.OtherLink),expression:"notShowIfNull(props.row.OtherLink)"}],attrs:{label:"其余渠道的链接"}},[i("div",{staticStyle:{display:"inline-block","text-align":"center"},attrs:{slot:"label"},slot:"label"},[i("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.row.OtherLink,expression:"props.row.OtherLink",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}}),e._v(" "),i("span",[e._v("其余渠道的链接")])]),e._v(" "),i("a",{attrs:{href:"javascript:;"},on:{click:function(i){return e.openUrl(t.row.OtherLink)}}},[e._v("\n                "+e._s(e._f("showFullLink")(t.row.OtherLink))+"\n              ")])]),e._v(" "),i("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.notShowIfNull(t.row.FanFictionLink),expression:"notShowIfNull(props.row.FanFictionLink)"}],attrs:{label:"同人链接"}},[i("div",{staticStyle:{display:"inline-block","text-align":"center"},attrs:{slot:"label"},slot:"label"},[i("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.row.FanFictionLink,expression:"props.row.FanFictionLink",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}}),e._v(" "),i("span",[e._v("同人链接")])]),e._v(" "),i("a",{attrs:{href:"javascript:;"},on:{click:function(i){return e.openUrl(t.row.FanFictionLink)}}},[e._v("\n                "+e._s(e._f("showFullLink")(t.row.FanFictionLink))+"\n              ")])]),e._v(" "),e._l(t.row.ShowImgLinks,function(t,a){return e.settingShowImg(t.value,"img")?i("el-form-item",{key:a,attrs:{label:"图链",prop:t.value}},[i("div",{staticStyle:{display:"inline-block","text-align":"center"},attrs:{slot:"label"},slot:"label"},[i("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.value,expression:"item.value",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}}),e._v(" "),i("span",[e._v("←复制图链")])]),e._v(" "),i("div",{staticStyle:{display:"inline-block","text-align":"center"}},[i("el-image",{staticStyle:{width:"100px",height:"100px"},attrs:{src:t.value,fit:"scale-down",lazy:"","preview-src-list":t.srcList}},[i("div",{staticClass:"image-slot",attrs:{slot:"placeholder"},slot:"placeholder"},[i("img",{staticClass:"el-icon-picture-outline",staticStyle:{width:"100px",height:"100px"}})])])],1)]):e._e()}),e._v(" "),e._l(t.row.ShowImgLinks,function(t,a){return e.settingShowImg(t.value,"info")?i("el-form-item",{key:a,attrs:{prop:t.value}},[i("div",{staticStyle:{display:"inline-block","text-align":"center"},attrs:{slot:"label"},slot:"label"},[i("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.value,expression:"item.value",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}}),e._v(" "),i("span",[e._v("图链")])]),e._v(" "),i("div",{staticStyle:{display:"inline-block","text-align":"center"}},[i("a",{attrs:{href:"javascript:;"},on:{click:function(i){return e.openUrl(t.value)}}},[e._v("\n                    "+e._s(e._f("showFullLink")(t.value))+"\n                  ")])])]):e._e()}),e._v(" "),i("el-form-item",{attrs:{label:"添加时间"}},[i("span",[e._v("\n                "+e._s(e._f("dateFormat")(t.row.AddTime))+"\n              ")])]),e._v(" "),i("el-form-item",{attrs:{label:"更新时间"}},[i("span",[e._v("\n                "+e._s(e._f("dateFormat")(t.row.UpdateTime))+"\n              ")])]),e._v(" "),i("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.isAdmin,expression:"isAdmin"}],attrs:{label:"管理员操作"}},[i("el-button",{attrs:{size:"mini"},on:{click:function(i){return e.disableThisInfo(t.row.Id)}}},[e._v("屏蔽此条")])],1)],2)]}}])}),e._v(" "),i("el-table-column",{attrs:{fixed:"left",label:"类型",width:"60"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("el-tag",{attrs:{size:"mini",type:e.infoTypeStyle(t.row.InfoType)}},[e._v(e._s(e._f("infoTypeName")(t.row.InfoType)))])]}}])}),e._v(" "),i("el-table-column",{attrs:{fixed:"left",label:"详细描述",width:"230"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("div",{staticStyle:{"white-space":"pre-line"}},[i("p",[e._v(e._s(t.row.Info))])])]}}])}),e._v(" "),e.MobileIsPc?i("el-table-column",{attrs:{prop:"TheTime",label:"在该视频的哪个时段",width:"150"}}):e._e(),e._v(" "),e.MobileIsPc?i("el-table-column",{attrs:{prop:"BiliBiliLink",label:"b站链接",width:"120"}}):e._e(),e._v(" "),e.MobileIsPc?i("el-table-column",{attrs:{prop:"YoutubeLink",label:"油管链接",width:"120"}}):e._e(),e._v(" "),e.MobileIsPc?i("el-table-column",{attrs:{prop:"OtherLink",label:"其余渠道的链接",width:"120"}}):e._e(),e._v(" "),e.MobileIsPc?i("el-table-column",{attrs:{prop:"FanFictionLink",label:"同人链接",width:"120"}}):e._e(),e._v(" "),e.MobileIsPc?i("el-table-column",{attrs:{label:"创建时间",width:"120"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("span",[e._v("\n            "+e._s(e._f("dateFormat")(t.row.AddTime))+"\n          ")])]}}],null,!1,555091900)}):e._e(),e._v(" "),e.MobileIsPc?i("el-table-column",{attrs:{prop:"UpdateTime",label:"更新时间",width:"120"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("span",[e._v("\n            "+e._s(e._f("dateFormat")(t.row.UpdateTime))+"\n          ")])]}}],null,!1,3497406828)}):e._e()],1),e._v(" "),i("el-pagination",{attrs:{small:e.PagerIsSmall,"current-page":e.page.currentPage,"pager-count":e.PagerCount,"page-sizes":[1,10,20,50,100],"page-size":e.page.pageSize,layout:"total, sizes, prev, pager, next",total:e.page.totalNum},on:{"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}})],1)},staticRenderFns:[]};var p=i("VU/8")(n,c,!1,function(e){i("ZdlP")},"data-v-29411c1a",null);t.default=p.exports},ZdlP:function(e,t){}});