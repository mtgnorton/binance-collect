(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-8d719210"],{"2e05":function(e,t,r){"use strict";r.r(t);var s=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("el-form",{ref:"form",attrs:{model:e.user,rules:e.rules,"label-width":"80px"}},[r("el-form-item",{attrs:{label:"手机号码",prop:"phonenumber"}},[r("el-input",{attrs:{maxlength:"11"},model:{value:e.user.phonenumber,callback:function(t){e.$set(e.user,"phonenumber",t)},expression:"user.phonenumber"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"邮箱",prop:"email"}},[r("el-input",{attrs:{maxlength:"50"},model:{value:e.user.email,callback:function(t){e.$set(e.user,"email",t)},expression:"user.email"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"性别"}},[r("el-radio-group",{model:{value:e.user.sex,callback:function(t){e.$set(e.user,"sex",t)},expression:"user.sex"}},[r("el-radio",{attrs:{label:"0"}},[e._v("男")]),e._v(" "),r("el-radio",{attrs:{label:"1"}},[e._v("女")])],1)],1),e._v(" "),r("el-form-item",[r("el-button",{attrs:{type:"primary",size:"mini"},on:{click:e.submit}},[e._v("保存")]),e._v(" "),r("el-button",{attrs:{type:"danger",size:"mini"},on:{click:e.close}},[e._v("关闭")])],1)],1)},a=[],o=r("c0c7"),i={props:{user:{type:Object}},data:function(){return{rules:{user_name:[{required:!0,message:"用户昵称不能为空",trigger:"blur"}],email:[{required:!0,message:"邮箱地址不能为空",trigger:"blur"},{type:"email",message:"'请输入正确的邮箱地址",trigger:["blur","change"]}],phonenumber:[{required:!0,message:"手机号码不能为空",trigger:"blur"},{pattern:/^1[3|4|5|6|7|8|9][0-9]\d{8}$/,message:"请输入正确的手机号码",trigger:"blur"}]}}},methods:{submit:function(){var e=this;this.$refs["form"].validate((function(t){t&&Object(o["h"])(e.user).then((function(t){0===t.code?e.msgSuccess("修改成功"):e.msgError(t.msg)}))}))},close:function(){this.$store.dispatch("tagsView/delView",this.$route),this.$router.push({path:"/index"})}}},n=i,l=r("2877"),u=Object(l["a"])(n,s,a,!1,null,null,null);t["default"]=u.exports},"6d40":function(e,t,r){"use strict";r.r(t);var s=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"app-container"},[r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:6,xs:24}},[r("el-card",{staticClass:"box-card"},[r("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[r("span",[e._v("个人信息")])]),e._v(" "),r("div",[r("div",{staticClass:"text-center"},[r("userAvatar",{attrs:{user:e.user}})],1),e._v(" "),r("ul",{staticClass:"list-group list-group-striped"},[r("li",{staticClass:"list-group-item"},[r("svg-icon",{attrs:{"icon-class":"user"}}),e._v("\n              用户名称\n              "),r("div",{staticClass:"pull-right"},[e._v(e._s(e.user.username))])],1),e._v(" "),r("li",{staticClass:"list-group-item"},[r("svg-icon",{attrs:{"icon-class":"date"}}),e._v("\n              创建日期\n              "),r("div",{staticClass:"pull-right"},[e._v(e._s(e.user.created_at))])],1)])])])],1),e._v(" "),r("el-col",{attrs:{span:18,xs:24}},[r("el-card",[r("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[r("span",[e._v("基本资料")])]),e._v(" "),r("el-form",{ref:"form",attrs:{model:e.user,rules:e.rules,"label-width":"80px"}},[r("el-form-item",{attrs:{label:"用户昵称",prop:"nickname"}},[r("el-input",{model:{value:e.user.nickname,callback:function(t){e.$set(e.user,"nickname",t)},expression:"user.nickname"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"旧密码",prop:"old_password"}},[r("el-input",{attrs:{placeholder:"请输入旧密码",type:"password"},model:{value:e.user.old_password,callback:function(t){e.$set(e.user,"old_password",t)},expression:"user.old_password"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"新密码",prop:"password"}},[r("el-input",{attrs:{placeholder:"请输入新密码",type:"password"},model:{value:e.user.password,callback:function(t){e.$set(e.user,"password",t)},expression:"user.password"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"确认密码",prop:"password_confirmed"}},[r("el-input",{attrs:{placeholder:"请确认密码",type:"password"},model:{value:e.user.password_confirmed,callback:function(t){e.$set(e.user,"password_confirmed",t)},expression:"user.password_confirmed"}})],1),e._v(" "),r("el-form-item",[r("el-button",{attrs:{type:"primary",size:"mini"},on:{click:e.submit}},[e._v("保存")]),e._v(" "),r("el-button",{attrs:{type:"danger",size:"mini"},on:{click:e.close}},[e._v("关闭")])],1)],1)],1)],1)],1)],1)},a=[],o=r("9179"),i=r("2e05"),n=r("8a24"),l=r("c0c7"),u={name:"Profile",components:{userAvatar:o["default"],userInfo:i["default"],resetPwd:n["default"]},data:function(){var e=this,t=function(t,r,s){e.user.newPassword!==r?s(new Error("两次输入的密码不一致")):s()};return{test:"1test",user:{oldPassword:void 0,newPassword:void 0,confirmPassword:void 0},rules:{oldPassword:[{required:!0,message:"旧密码不能为空",trigger:"blur"}],newPassword:[{required:!0,message:"新密码不能为空",trigger:"blur"},{min:6,max:20,message:"长度在 6 到 20 个字符",trigger:"blur"}],confirmPassword:[{required:!0,message:"确认密码不能为空",trigger:"blur"},{required:!0,validator:t,trigger:"blur"}]}}},created:function(){this.getUser()},methods:{getUser:function(){var e=this;Object(l["e"])().then((function(t){e.user=t.data,console.log(t.data)}))},submit:function(){var e=this;this.$refs["form"].validate((function(t){t&&Object(l["h"])(e.user).then((function(t){0===t.code?e.msgSuccess("修改成功"):e.msgError(t.msg)}))}))},close:function(){this.$store.dispatch("tagsView/delView",this.$route),this.$router.push({path:"/index"})}}},c=u,p=r("2877"),d=Object(p["a"])(c,s,a,!1,null,null,null);t["default"]=d.exports},"8a24":function(e,t,r){"use strict";r.r(t);var s=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div")},a=[],o=(r("c0c7"),{data:function(){},methods:{}}),i=o,n=r("2877"),l=Object(n["a"])(i,s,a,!1,null,null,null);t["default"]=l.exports},9179:function(e,t,r){"use strict";r.r(t);var s=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",[r("img",{staticClass:"img-circle img-lg",attrs:{src:e.options.img,title:"点击上传头像"},on:{click:function(t){return e.editCropper()}}}),e._v(" "),r("el-dialog",{attrs:{title:e.title,visible:e.open,width:"800px","append-to-body":""},on:{"update:visible":function(t){e.open=t}}},[r("el-row",[r("el-col",{style:{height:"350px"},attrs:{xs:24,md:12}},[r("vue-cropper",{ref:"cropper",attrs:{img:e.options.img,info:!0,autoCrop:e.options.autoCrop,autoCropWidth:e.options.autoCropWidth,autoCropHeight:e.options.autoCropHeight,fixedBox:e.options.fixedBox},on:{realTime:e.realTime}})],1),e._v(" "),r("el-col",{style:{height:"350px"},attrs:{xs:24,md:12}},[r("div",{staticClass:"avatar-upload-preview"},[r("img",{style:e.previews.img,attrs:{src:e.previews.url}})])])],1),e._v(" "),r("br"),e._v(" "),r("el-row",[r("el-col",{attrs:{lg:2,md:2}},[r("el-upload",{attrs:{action:"#","http-request":e.requestUpload,"show-file-list":!1,"before-upload":e.beforeUpload}},[r("el-button",{attrs:{size:"small"}},[e._v("\n            上传\n            "),r("i",{staticClass:"el-icon-upload el-icon--right"})])],1)],1),e._v(" "),r("el-col",{attrs:{lg:{span:1,offset:2},md:2}},[r("el-button",{attrs:{icon:"el-icon-plus",size:"small"},on:{click:function(t){return e.changeScale(1)}}})],1),e._v(" "),r("el-col",{attrs:{lg:{span:1,offset:1},md:2}},[r("el-button",{attrs:{icon:"el-icon-minus",size:"small"},on:{click:function(t){return e.changeScale(-1)}}})],1),e._v(" "),r("el-col",{attrs:{lg:{span:1,offset:1},md:2}},[r("el-button",{attrs:{icon:"el-icon-refresh-left",size:"small"},on:{click:function(t){return e.rotateLeft()}}})],1),e._v(" "),r("el-col",{attrs:{lg:{span:1,offset:1},md:2}},[r("el-button",{attrs:{icon:"el-icon-refresh-right",size:"small"},on:{click:function(t){return e.rotateRight()}}})],1),e._v(" "),r("el-col",{attrs:{lg:{span:2,offset:6},md:2}},[r("el-button",{attrs:{type:"primary",size:"small"},on:{click:function(t){return e.uploadImg()}}},[e._v("提 交")])],1)],1)],1)],1)},a=[],o=(r("7f7f"),r("4360")),i=r("7e79"),n=r("c0c7"),l={components:{VueCropper:i["VueCropper"]},props:{user:{type:Object}},data:function(){return{open:!1,title:"修改头像",filename:"",options:{img:o["a"].getters.avatar,autoCrop:!0,autoCropWidth:200,autoCropHeight:200,fixedBox:!0},previews:{}}},methods:{editCropper:function(){this.open=!0,this.filename=""},requestUpload:function(){},rotateLeft:function(){this.$refs.cropper.rotateLeft()},rotateRight:function(){this.$refs.cropper.rotateRight()},changeScale:function(e){e=e||1,this.$refs.cropper.changeScale(e)},beforeUpload:function(e){var t=this;if(-1==e.type.indexOf("image/"))this.msgError("文件格式错误，请上传图片类型,如：JPG，PNG后缀的文件。");else{this.filename=e.name;var r=new FileReader;r.readAsDataURL(e),r.onload=function(){t.options.img=r.result}}},uploadImg:function(){var e=this;this.$refs.cropper.getCropBlob((function(t){var r=new FormData;r.append("avatar_file",t,e.filename),Object(n["i"])(r).then((function(t){0===t.code?(e.open=!1,e.options.img=t.data.avatar_url,e.msgSuccess("修改成功")):e.msgError(t.msg),e.$refs.cropper.clearCrop()}))}))},realTime:function(e){this.previews=e}}},u=l,c=r("2877"),p=Object(c["a"])(u,s,a,!1,null,null,null);t["default"]=p.exports},c0c7:function(e,t,r){"use strict";r.d(t,"f",(function(){return o})),r.d(t,"d",(function(){return i})),r.d(t,"a",(function(){return n})),r.d(t,"g",(function(){return l})),r.d(t,"c",(function(){return u})),r.d(t,"b",(function(){return c})),r.d(t,"e",(function(){return p})),r.d(t,"h",(function(){return d})),r.d(t,"i",(function(){return m}));var s=r("b775"),a=r("c38a");function o(e){return Object(s["a"])({url:"/administrator-list",method:"get",params:e})}function i(e){return Object(s["a"])({url:"/administrator-info",params:{id:Object(a["e"])(e)},method:"get"})}function n(e){return Object(s["a"])({url:"/administrator-store",method:"post",data:e})}function l(e){return Object(s["a"])({url:"/administrator-update",method:"put",data:e})}function u(e){return Object(s["a"])({url:"/administrator-destroy",params:{Ids:e},method:"delete"})}function c(e,t){var r={userId:e,status:t};return Object(s["a"])({url:"/system/user/changeStatus",method:"put",data:r})}function p(){return Object(s["a"])({url:"/personal-info",method:"get"})}function d(e){return Object(s["a"])({url:"/personal-update",method:"put",data:e})}function m(e){return Object(s["a"])({url:"/personal-avatar",method:"post",data:e})}}}]);