webpackJsonp([4],{OUQW:function(e,t){},bfvG:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var i=a("TQvf"),r=a.n(i),o=a("PJh5"),s=a.n(o),n=a("bFR2"),l=a("mg6h"),c={name:"YoutubeList",components:{DateTime:a.n(l).a},data:function(){return{setData:"from",loading:!0,PagerIsSmall:!1,PagerCount:7,currentRow:"",page:{totalNum:0,currentPage:1,pageSize:10},totalList:[],showList:[],MobileIsPc:!0,filterPara:{from:"2014-01-01",to:"2050-01-01",keyword:"",ifDesc:!1}}},filters:{dateFormat:function(e){return s()(e).add(8,"hours").format("YYYY-MM-DD HH:mm:ss")}},created:function(){this.MobileIsPc=!this.isMobile(),this.MobileIsPc||(this.PagerCount=5,this.PagerIsSmall=!0);var e=n.a.get("two-set-info-youtube-list-filterPara");void 0!==e&&(this.filterPara=e),this.page.pageSize=n.a.get("two-set-info-youtube-list-pageSize"),void 0!==this.page.pageSize&&""!==this.page.pageSize||(this.page.pageSize=10),this.page.currentPage=n.a.get("two-set-info-youtube-list-currentPage"),void 0!==this.page.currentPage&&""!==this.page.currentPage||(this.page.currentPage=1),this.getInfo(),this.currentRow=n.a.get("two-set-info-youtube-list-currentRow")},methods:{dateTimeShow:function(e){this.setData=e,this.$refs.dateTime.show()},dateTimeConfirm:function(e){"to"===this.setData?(this.filterPara.to=e,this.filterPara.from>this.filterPara.to&&(this.$message({type:"warn",message:"输入日期【从】要早于【到】"}),this.filterPara.to="2050-01-01")):(this.filterPara.from=e,this.filterPara.from>this.filterPara.to&&(this.$message({type:"warn",message:"输入日期【从】要早于【到】"}),this.filterPara.from="2014-01-01")),console.log("get date"+e)},getTableRowIndex:function(e){var t=e.row,a=e.rowIndex;t.index=a},openUrl:function(e){/^(https?|ftp|file):\/\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]/.test(e)?window.open(e):this.$message({message:"url格式不合法",duration:500}),console.log(e)},handleHighlightCurrentChange:function(e){this.currentRow=e.index,n.a.set("two-set-info-youtube-list-currentRow",this.currentRow),console.log("当前行:"+this.currentRow)},handleSizeChange:function(e){this.currentRow="",n.a.set("two-set-info-youtube-list-currentRow",this.currentRow),console.log("每页 "+e+" 条"),this.page.pageSize=e,n.a.set("two-set-info-youtube-list-pageSize",e),this.updateCurrentPageInfo()},handleCurrentPageChange:function(e){this.currentRow="",n.a.set("two-set-info-youtube-list-currentRow",this.currentRow),this.page.currentPage=e,n.a.set("two-set-info-youtube-list-currentPage",e),console.log("当前页: "+e),this.updateCurrentPageInfo()},isMobile:function(){return navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)},copy:function(){var e=new r.a(".tag-read");e.on("success",function(t){e.destroy()}),e.on("error",function(t){e.destroy()}),this.$message({type:"success",message:"已复制",duration:500})},paraReset:function(){this.filterPara.from="2014-01-01",this.filterPara.to="2050-01-01",this.filterPara.keyword="",this.filterPara.ifDesc=!1},getInfo:function(){var e=this;this.loading=!0,this.page.totalNum=0,this.$http.post("/api/query_youtube_video_list",{from:this.filterPara.from,to:this.filterPara.to,keyword:this.filterPara.keyword,ifDesc:this.filterPara.ifDesc?"1":"0"},{emulateJSON:!0}).then(function(t){var a=t.data;0===a.error_num?(n.a.set("two-set-info-youtube-list-filterPara",e.filterPara),e.totalList=a.info_list,e.showList=a.info_list,e.page.totalNum=a.total_num,e.updateCurrentPageInfo(),e.loading=!1,console.log("this.page.totalNum:",e.page.totalNum)):e.$message.error(a.msg)})},updateCurrentPageInfo:function(){var e=this,t=(this.page.currentPage-1)*this.page.pageSize,a=this.page.currentPage*this.page.pageSize;a=a>this.page.totalNum?this.page.totalNum:a,this.showList=this.totalList.slice(t,a);for(var i=0;i<this.showList.length;i++)this.showList[i].Title=this.showList[i].Title.replace(/&lt;/g,"<").replace(/&gt;/g,">").replace(/&amp;/g,"&").replace(/&quot;/g,'"').replace(/&#39;/g,"'").replace(/&nbsp;/g,"").replace(/\n/g,"<br/>").replace(/\r/g,"<br/>");console.log(this.showList),void 0!==this.currentRow&&""!==this.currentRow&&(console.log("高亮指定行"),this.$nextTick(function(){e.$refs.aTable.setCurrentRow(e.showList[e.currentRow])}))}}},u={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("el-collapse",[a("el-collapse-item",{staticStyle:{width:"360px"},attrs:{title:"搜索",name:"1"}},[a("el-row",[a("el-input",{attrs:{size:"mini",placeholder:"请输入标题关键字"},model:{value:e.filterPara.keyword,callback:function(t){e.$set(e.filterPara,"keyword",t)},expression:"filterPara.keyword"}})],1),e._v(" "),a("el-row",{staticStyle:{"margin-top":"10px"}},[a("el-switch",{attrs:{"inactive-color":"#13ce66","active-text":"新到旧","inactive-text":"旧到新"},model:{value:e.filterPara.ifDesc,callback:function(t){e.$set(e.filterPara,"ifDesc",t)},expression:"filterPara.ifDesc"}}),e._v(" "),a("span",{staticStyle:{"font-size":"small","margin-left":"10px"}},[e._v("从")]),e._v(" "),a("el-button",{attrs:{type:"text"},on:{click:function(t){return e.dateTimeShow("from")}}},[e._v(e._s(e.filterPara.from))]),e._v(" "),a("span",{staticStyle:{"font-size":"small"}},[e._v("到")]),e._v(" "),a("el-button",{attrs:{type:"text"},on:{click:function(t){return e.dateTimeShow("to")}}},[e._v(e._s(e.filterPara.to))])],1),e._v(" "),a("el-row",[a("el-button",{attrs:{size:"mini",type:"primary",icon:"el-icon-refresh",circle:""},on:{click:e.paraReset}}),e._v(" "),a("el-button",{attrs:{size:"mini",type:"primary",icon:"el-icon-search",circle:""},on:{click:e.getInfo}})],1)],1)],1),e._v(" "),a("date-time",{ref:"dateTime",attrs:{type:"date",format:"yyyy-MM-dd"},on:{confirm:e.dateTimeConfirm}}),e._v(" "),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],ref:"aTable",staticStyle:{width:"100%"},attrs:{"highlight-current-row":"","row-class-name":e.getTableRowIndex,data:e.showList,"header-align":"center"},on:{"row-click":e.handleHighlightCurrentChange}},[e.MobileIsPc?e._e():a("el-table-column",{attrs:{fixed:"left",width:"30"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.row.VideoUrl,expression:"scope.row.VideoUrl",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}})]}}],null,!1,1921819720)}),e._v(" "),e.MobileIsPc?e._e():a("el-table-column",{attrs:{fixed:"left",label:"标题",width:"200"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("div",[a("span",{staticStyle:{"white-space":"normal","word-break":"keep-all"},attrs:{href:"javascript:;"},on:{click:function(a){return e.openUrl(t.row.VideoUrl)}}},[e._v("\n            "+e._s(t.row.Title)+"\n          ")])])]}}],null,!1,2250145198)}),e._v(" "),e.MobileIsPc?a("el-table-column",{attrs:{prop:"Title",label:"标题",width:"250"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("div",[a("a",{staticStyle:{"white-space":"normal","word-break":"keep-all"},attrs:{href:"javascript:;"},on:{click:function(a){return e.openUrl(t.row.VideoUrl)}}},[e._v("\n            "+e._s(t.row.Title)+"\n          ")])])]}}],null,!1,1077567587)}):e._e(),e._v(" "),a("el-table-column",{attrs:{label:"发布时间",width:"160"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v("\n          "+e._s(e._f("dateFormat")(t.row.PublishedAt))+"\n        ")])]}}])}),e._v(" "),e.MobileIsPc?a("el-table-column",{attrs:{label:"点击保存",width:"120"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("i",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.row.VideoUrl,expression:"scope.row.VideoUrl",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.copy,expression:"copy",arg:"success"}],staticClass:"el-icon-copy-document",staticStyle:{"margin-right":"10px"}})]}}],null,!1,1921819720)}):e._e()],1),e._v(" "),a("el-pagination",{attrs:{small:e.PagerIsSmall,"current-page":e.page.currentPage,"pager-count":e.PagerCount,"page-sizes":[1,10,20,50,100],"page-size":e.page.pageSize,layout:"total, sizes, prev, pager, next",total:e.page.totalNum},on:{"size-change":e.handleSizeChange,"current-change":e.handleCurrentPageChange}})],1)},staticRenderFns:[]};var f=a("VU/8")(c,u,!1,function(e){a("OUQW")},"data-v-06dfaf7e",null);t.default=f.exports}});