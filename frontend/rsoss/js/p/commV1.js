/*///常用一些js: md5 , JL_pgData ,JL_Check 
//var KK_LOGIN_REG = new jlCheck();
//onclick="KK_LOGIN_REG.regSubmit4();" */
/* every page cb */
var JL_Base = function(){
	this.pageUrl=function(){
		return (window.location.href.toLowerCase());
	};
	this.V = function(id2){
		try{
		return (document.getElementById(id2).value);}catch(Exception){};
		return "";
	};
	this.CB=function(id2){
		var o = document.getElementById(id2);
		if(o.checked) return true;
		return false;
	};
	this.CbSbyC=function(n,o ){//("fbb",this)
		var ckf = o.checked;
		if(ckf)this.CbS(n,1);else {this.CbS(n,0);};
	};
	this.CbS=function(n,ty){ //ty=1,0
		 var o = document.getElementsByName(n);
			 for(var i=0;i< o.length;i++){
                 if(ty==1){o[i].checked =true;}
				 else{ o[i].checked = false; };
             };
	};
	this.RV=function(n){
		 var o = document.getElementsByName(n);
         var chk=0;
			 for(var i=0;i< o.length;i++){
                    if(o[i].checked){   chk = o[i].value;  break;    };
             };
		return chk;
	};
	this.$ = function(id2){
		return (document.getElementById(id2));
	};
	this.G = function(id2){
		return document.getElementById(id2);
	};
	this.VL2 = function(o){
		return (o).length ;//.length ;
	};
	this.VL = function(id2){
		return (document.getElementById(id2).value).length ;//.length ;
	};
	this.Write= function(id,body2){
		document.getElementById(id).innerHTML = body2+"";
	};
	this.Log= function(id,body2){
		document.getElementById(id).innerHTML += body2+"<br>";
	};
	this.GoUrl =function(url){
		location.href=url ;
	};
	this.GoBack =function(){
		window.history.back();
	};
	this.GoBack2 =function(){
		window.history.go(-1);
	};
	this.newScript = function(url){
		var  divID="pg_script" ;
		this.newScriptV(url,divID,0);
	};
	this.newScriptB = function(url,divID){
		this.newScriptV(url,divID,0);
	};
	this.newScriptV = function(url,divID,ty){
		var obj = document.createElement("script");
		obj.language="javascript";
		obj.type="text/javascript";
		obj.charset = "utf-8";
		obj.src = url;
		if(ty==1) document.body.appendChild(obj);
		else {
			document.getElementById(divID).innerHTML = "";
			document.getElementById(divID).appendChild(obj);
		};
	};
	this.time = function(){ /*back: long */
		return (new Date).getTime();
	};
	this.timeOrderID = function(){ /*back: long */
		var d = new Date(); 
		 var ret = d.getFullYear() + "";
               ret += ("00" + (d.getMonth() + 1)).slice(-2) + "";
               ret += ("00" + d.getDate()).slice(-2) + "";
               ret += ("00" + d.getHours()).slice(-2) + "";
               ret += ("00" + d.getMinutes()).slice(-2) + "";
               ret += ("00" + d.getSeconds()).slice(-2);
               return ret;
	};
	this.GetQueryValue=function(sorStr, panStr,defaultValue) {
	    /*	在JavaScript内实现QueryString	    eg: var   fid   =   GetQueryValue(this_url, 'fid ',0);++*/
	    var vStr = defaultValue ;
	    if (sorStr == null || sorStr == "" || panStr == null || panStr == "") return vStr;
	    panStr = panStr.toLowerCase();
	    sorStr = sorStr.toLowerCase();
	    panStr += "=";
	    var itmp = sorStr.indexOf(panStr);
	    if (itmp < 0) { return vStr; };
	    sorStr = sorStr.substr(itmp + panStr.length);
	    itmp = sorStr.indexOf("&");
	    if (itmp < 0) {
	        return sorStr;
	    }else {
	        sorStr = sorStr.substr(0, itmp);
	        return sorStr;
	    };
	};
	this.Display=function(id,IsShow){
		try{
			if(IsShow==1) this.G(id).style.display="block";
			else this.G(id).style.display="none";
		}catch(Exception){}	;
	};
	this.IsIE=function(){
		if(!+[1,])return true;
		return false;	
	};
	this.IEv=function(){
		 if(navigator.appName != "Microsoft Internet Explorer") return -1;
          if(navigator.appVersion.match(/6./i)=='6.') return 6;
          if(navigator.appVersion.match(/7./i)=='7.')return 7;
          if(navigator.appVersion.match(/8./i)=='8.') return 8;
		 if(navigator.appVersion.match(/8./i)=='9.') return 9;
		 if(navigator.appVersion.match(/8./i)=='10.') return 10;
		 return 99;
	};
};
/*check-- ck */
var JL_Check = function()
{
	this.checkEmpty = function(str){
		return (/^ *$/.test(str));
	};
	
	this.checkMobile = function(name){
		return (  /^0?(13[0-9]|15[012356789]|18[0236789]|14[57])[0-9]{8}$/.test(name));
	};
	this.checkVerify= function(s1){
		return (/^[0-9a-zA-Z]{4,6}$/.test(s1) );
	};
	
	this.checkEmail = function(name){
		return (/^[0-9a-zA-Z][\w-]+@\w+\.\w+$/.test(name) );
	};
	this.checkPasswd=function(s1) {
	    var patrn = (/.{6,15}/);
	    if (patrn.test(s1)) {
	        return true;
	    }
	    else {
	        return false;
	    };
	};
	
	this.checkSafePassword=function (str,uname) {
	    if (str == uname)
	       { return 1;};
	    for (var i = 0; i < str.length; i++) {
	        if (str.charAt(0) != str.charAt(i))
	            break;
	    };
	    if (i == str.length) return 2;
	    var seqStr = "01234567890";	    if (seqStr.indexOf(str) != -1) return 3;
		var seqStr = "9876543210";
	    if (seqStr.indexOf(str) != -1) return 3;
	    var seqStr = "abcdefghijklmnopqrstuvwxyz";
	    if (seqStr.indexOf(str) != -1) return 4;
	    var seqStr = "zyxwvutsrqponmlkjihgfedcba";
	    if (seqStr.indexOf(str) != -1) return 4;
	    var seqStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	    if (seqStr.indexOf(str) != -1) return 5;
	    var seqStr = "ZYXWVUTSRQPONMLKJIHGFEDCBA";
	    if (seqStr.indexOf(str) != -1) return 5;
	    return 0;
	};
	
	this.checkUserName=function(s1) {
	    //var patrn = (/^[a-zA-Z]{1}[a-zA-Z0-9\-_]{5,15}$/);
		 var patrn = (/^[a-zA-Z0-9\-_@.]{6,16}$/);
	    if (patrn.test(s1)) {
	        return true;
	    }
	    else {
	        return false;
	    };
	};
	
	this.checkbadname =function (s1) {
	    var badname = new Array("admin", "administrator","jl","jlwww");
	    for (var i = 0; i < badname.length; i++) {
	        if (s1.toString().toLowerCase() == badname[i]) {
	            break;
	        };
	    };
	    if (i == badname.length)	    { return true ;}	    else	    { return false; };
	};
	
	this.isNumber=function(oNum) {
	    if (!oNum) return false;
	    var strP = (/^\d+(\.\d+)?$/);
	    if (!strP.test(oNum)) return false;
	    try {
	        if (parseFloat(oNum) != oNum) return false;
	    }
	    catch (ex) {
	        return false;
	    };
	    return true;
	};
	
	/*  public  */
	/*  public  */
	//检查汉字格式   
	this.TrueNane=function (s) {
	    //var patrn = /^[\\u4E00-\\u9FA5\\uF900-\\uFA2D]+$/ ;
	    //if (patrn.test(s)) {
	    if (s == "") return false;
	    var vv = s.replace(/[^\u4E00-\u9FA5]/g, "");
	    if (s.length == vv.length) {
	        return true;
	    }
	    else {
	        return false;
	    };
	};
	
	//校验身份证: return=0 ok.
	this.checkIdcard=function ( idcard ) {
		if(idcard.length <15 ){return 10;};
		var area = { 11: "北京", 12: "天津", 13: "河北", 14: "山西", 15: "内蒙古", 21: "辽宁", 22: "吉林", 23: "黑龙江", 31: "上海", 32: "江苏", 33: "浙江", 34: "安徽", 35: "福建", 36: "江西", 37: "山东", 41: "河南", 42: "湖北", 43: "湖南", 44: "广东", 45: "广西", 46: "海南", 50: "重庆", 51: "四川", 52: "贵州", 53: "云南", 54: "西藏", 61: "陕西", 62: "甘肃", 63: "青海", 64: "宁夏", 65: "新疆", 71: "台湾", 81: "香港", 82: "澳门", 91: "国外" };
	    //var idcard;
	    var Y, JYM;
	    var S, M;
	    var idcard_array = new Array();
	    idcard_array = idcard.split("");
	    //地区检验
	    if (area[parseInt(idcard.substr(0, 2))] == null){ return 4;};
	    //身份号码位数及格式检验
	    switch (idcard.length) {
	        case 15:
	            if ((parseInt(idcard.substr(6, 2)) + 1900) % 4 == 0 || ((parseInt(idcard.substr(6, 2)) + 1900) % 100 == 0 && (parseInt(idcard.substr(6, 2)) + 1900) % 4 == 0)) {
	                ereg = /^[1-9][0-9]{5}[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|[1-2][0-9]))[0-9]{3}$/; //测试出生日期的合法性

	            } else {
	                ereg = /^[1-9][0-9]{5}[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|1[0-9]|2[0-8]))[0-9]{3}$/; //测试出生日期的合法性

	            };
	            if (ereg.test(idcard)){ return 0;}
	            else{ return 2;};
	            break;
	        case 18:
	            //18位身份号码检测
	            //出生日期的合法性检查 
	            //闰年月日:((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|[1-2][0-9]))
	            //平年月日:((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|1[0-9]|2[0-8]))
	            if (parseInt(idcard.substr(6, 4)) % 4 == 0 || (parseInt(idcard.substr(6, 4)) % 100 == 0 && parseInt(idcard.substr(6, 4)) % 4 == 0)) {
	                ereg = /^[1-9][0-9]{5}19[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|[1-2][0-9]))[0-9]{3}[0-9Xx]$/; //闰年出生日期的合法性正则表达式
	            } else {
	                ereg = /^[1-9][0-9]{5}19[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|1[0-9]|2[0-8]))[0-9]{3}[0-9Xx]$/; //平年出生日期的合法性正则表达式
	            };
	            if (ereg.test(idcard)) {//测试出生日期的合法性

	                //计算校验位
	                S = (parseInt(idcard_array[0]) + parseInt(idcard_array[10])) * 7 + (parseInt(idcard_array[1]) + parseInt(idcard_array[11])) * 9 + (parseInt(idcard_array[2]) + parseInt(idcard_array[12])) * 10 + (parseInt(idcard_array[3]) + parseInt(idcard_array[13])) * 5 + (parseInt(idcard_array[4]) + parseInt(idcard_array[14])) * 8 + (parseInt(idcard_array[5]) + parseInt(idcard_array[15])) * 4 + (parseInt(idcard_array[6]) + parseInt(idcard_array[16])) * 2 + parseInt(idcard_array[7]) * 1 + parseInt(idcard_array[8]) * 6 + parseInt(idcard_array[9]) * 3;
	                Y = S % 11;
	                M = "F";
	                JYM = "10X98765432";
	                M = JYM.substr(Y, 1); //判断校验位

	                if (M == idcard_array[17]) return 0; //检测ID的校验位 ok
	                else {return 3;};
	            } else {return 2;};
	            break;
	        default:
	            return 1;
	            break;
	    };
	};
	
};

var JL_Text= function(){ //tx
	/////crean :还原
 	this.TextBack = function( sContent ) {
		sContent = sContent || '' ;
		if(sContent== ''){
			return sContent;
	    };
		if(sContent.indexOf("ζā") >=0){ //'
			sContent=sContent.replace(/\ζā/g,"\'");
		};
		if(sContent.indexOf("ζá") >=0){ //"
			sContent=sContent.replace(/\ζá/g,"\"");
		};
		if(sContent.indexOf("ζǎ") >=0){ //br
			sContent=sContent.replace(/\ζǎ/g,"<br>");
		};
		//sContent=sContent.Trim();
 		return sContent; 
 	};
	this.TextBackE = function( sContent ) {
		sContent = sContent || '' ;
		if(sContent== ''){
			return sContent;
	    };
		if(sContent.indexOf("ζā") >=0){ //'
			sContent=sContent.replace(/\ζā/g,"\'");
		};
		if(sContent.indexOf("ζá") >=0){ //"
			sContent=sContent.replace(/\ζá/g,"\"");
		};
		if(sContent.indexOf("ζǎ") >=0){ //br
			sContent=sContent.replace(/\ζǎ/g,"/n");
		};
		//sContent=sContent.Trim();
 		return sContent; 
 	};
	//传前加工
	this.TextPost = function(sContent) {
		sContent = sContent || '' ;
		if(sContent== ''){
			return sContent;
	    };
		if(sContent.indexOf("\'")>=0){ //'
			sContent=sContent.replace(/\'/g,"ζā");
		};
		if(sContent.indexOf("\"")>=0){
			sContent=sContent.replace(/\"/g,"ζá");
		};
		if(sContent.indexOf("<br>") >=0 ){
			sContent=sContent.replace(/\<br>/g ,"ζǎ");
		};
		//sContent=sContent.Trim();
 		return sContent; 
 	};
	
	this.replaceAll = function(oldString,newString){    
		return this.replace(new RegExp(oldString ,"gm"),newString);    
	};

	/*////crean :去掉特殊字符串:ā á ǎ à、ō ó ǒ ò、ê ē é ě è、ī í ǐ ì、ū ú ǔ ù、ǖ ǘ ǚ ǜ ü */
 	this.JsonStringCrear = function(sContent) {
		sContent = sContent || '' ;
		if(sContent== ''){
			return sContent;
	    };
		if(sContent.indexOf("\\")>=0){
			sContent=sContent.replace(/\\/g,"ζō");
		};
		if(sContent.indexOf("\'")>=0){
			sContent=sContent.replace(/\'/g,"ζā");
		};
		if(sContent.indexOf("\"")>=0 ){
			sContent=sContent.replace(/\"/g,"ζá");
		};
		/*///去掉字符串的回车换行符*/ 
		sContent = sContent.replace(/\\n\r/g, "ζǎ");
		/*/sContent=sContent.Trim();*/
 		return sContent; 
 	};
};
/* comm  cm */
var JL_Comm = function(){
    /*-var aj=new cm.AJAXRequest();aj.method = "POST"; aj.url= u; aj.content=PassUrl ; aj.callback = function(o){ jlbk.CallBack(o.responseText );}; aj.send(); --*/
	this.AJAXRequest=function()  {
			var xmlObj = false;
			var ObjSelf;
			ObjSelf=this;
			try { xmlObj=new XMLHttpRequest; }
			catch(e) {
				try { xmlObj=new ActiveXObject("MSXML2.XMLHTTP"); }
				catch(e2) {
					try { xmlObj=new ActiveXObject("Microsoft.XMLHTTP"); }
					catch(e3) { xmlObj=false; };
				};
			};
			if (!xmlObj) return false;
			this.method="POST";
			this.url;
			this.async=true;
			this.content="";
			this.callback=function(cbobj) {return;};
			this.send=function() {
				if(!this.method||!this.url||!this.async) return false;
				xmlObj.open (this.method, this.url, this.async);
				if(this.method=="POST") xmlObj.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
				xmlObj.onreadystatechange=function() {
					if(xmlObj.readyState==4) {
						if(xmlObj.status==200) {
							ObjSelf.callback(xmlObj);
						};
					};
				};
				if(this.method=="POST"){ xmlObj.send(this.content);}
				else{ xmlObj.send(null);}
			};
		};
	/*------------------------------*/
	//直接js返页
	this.pageList_UiJs=function( nowPage, pageno, AllRecord,jsClick ){/*{}*/
		if(AllRecord<1)return;		if(nowPage <=0) nowPage=1;
		var o1 = this.pageList_db(nowPage, pageno, AllRecord);
		var fir='',last='',next='',pre='',pgNoSt='';
		//var noM= 3 ;
		var  page = "nowpage";
		var pMaxNo = o1.MaxNo ;
		//if( nowPage >noM)fir='<li class=g1><a href="?' + page + '=1' + oldUrl + '"  >首页</a></li>';
		//if( pMaxNo -nowPage >=noM)last='<li class=g1><a href="?' + page + '=' +pMaxNo+ oldUrl + '"  >最后页</a></li>';
		if( o1.Next >0)next='<li class=g1><a  onClick="'+jsClick+'('+o1.Next+')"  >下页</a></li>';
		if( o1.Pre >0)pre='<li class=g1><a onClick="'+jsClick+'('+o1.Pre+')" >前页</a></li>';
		if(pMaxNo >1){
			var t = o1.No ;
			for(var i=0; i<t.length ;i++) {
				if(t[i] >0){	
					if(t[i] !=nowPage){pgNoSt += '<li class=n0><a  onClick="'+jsClick+'('+t[i]+')"  >'+t[i]+'</a></li>';}
					else {pgNoSt += '<li class=n1>'+t[i]+'</li>';};
				};
			};
		};
		var other= '<li>('+nowPage+'/'+pMaxNo+')</li>';
		var st= fir + pre + pgNoSt + next+ last + other ;
		return st  ;
	};
	/* v1 page string top  */
	this.pageList_Ui=function( nowPage, pageno, AllRecord,oldUrl ){/*{}*/
		if(AllRecord<1)return;
		if(nowPage <=0) nowPage=1;
		var o1 = this.pageList_db(nowPage, pageno, AllRecord);
		var fir='',last='',next='',pre='',pgNoSt='';
		//var noM= 3 ;
		var  page = "nowpage";
		var pMaxNo = o1.MaxNo ;
		//if( nowPage >noM)fir='<li class=g1><a href="?' + page + '=1' + oldUrl + '"  >首页</a></li>';
		//if( pMaxNo -nowPage >=noM)last='<li class=g1><a href="?' + page + '=' +pMaxNo+ oldUrl + '"  >最后页</a></li>';
		if( o1.Next >0)next='<li class=g1><a href="?' + page + '=' + o1.Next + oldUrl + '"  >下页</a></li>';
		if( o1.Pre >0)pre='<li class=g1><a href="?' + page + '=' + o1.Pre + oldUrl + '"  >前页</a></li>';
		if(pMaxNo >1){
			var t = o1.No ;
			for(var i=0; i<t.length ;i++) {
				if(t[i] >0){	
					if(t[i] !=nowPage){pgNoSt += '<li class=n0><a href="?' + page + '=' + t[i] + oldUrl + '"  >'+t[i]+'</a></li>';}
					else {pgNoSt += '<li class=n1>'+t[i]+'</li>';};
				};
			};
		};
		var other= '<li>('+nowPage+'/'+pMaxNo+')</li>';
		var st= fir + pre + pgNoSt + next+ last + other ;
		return st  ;
	};
	this.pageList_db=function( nowPage, pageno, AllRecord){/*{}*/
		if(nowPage <=0) nowPage=1;
		//var pMaxNo =  (AllRecord % pageno == 0)? (AllRecord / pageno): ( AllRecord / pageno+1) ;
		var pMaxNo =parseInt(AllRecord / pageno) ;
		if( AllRecord % pageno != 0)pMaxNo=pMaxNo +1;
		var next= (nowPage+1) <= pMaxNo ? (nowPage+1):0 ;
		var pre= (nowPage-1) >= 1 ? (nowPage-1):0 ;
		var first = nowPage==1?0:1 ;
		var last = nowPage==pMaxNo ?0:pMaxNo ;
		var no =[0,0,0 ];
		var noM= 3 ;
		if(pMaxNo <=noM){
			for(var i=0; i<pMaxNo ;i++){ 
				if(i<pMaxNo)no[i]= i+1 ;
				else {	no[i]= -1 ;};
			};
		};
		if( pMaxNo > noM && (nowPage -1)<=1 ){  /*边界左*/
			for(var i=0; i<noM ;i++){
				no[i]= i+1 ;
			};
		};
		if( pMaxNo > noM && (nowPage +1)>=pMaxNo ){ /*边界右*/
			for(var i=0; i<noM ;i++){
				no[i]= pMaxNo -3+i ;
			};
		};
		if( pMaxNo > noM && (nowPage +1)< pMaxNo && (nowPage -1)>1){ /*边界*/
			for(var i=0; i<noM ;i++){
				no[i]= nowPage -1+i ;
			};
		};
		var j={"MaxNo":pMaxNo,"Next":next,"Pre":pre,"First":first,"Last":last,"No":no };
		return j;
	};
	
	this.pageList_UiPc=function( nowPage, pageno, AllRecord,oldUrl ){/*{}*/
		if(AllRecord<1)return;
		if(nowPage <=0) nowPage=1;
		var o1 = this.pageList_dbPc(nowPage, pageno, AllRecord);
		var fir='',last='',next='',pre='',pgNoSt='';
		var noM= 5 ;
		var  page = "nowpage";
		var pMaxNo = o1.MaxNo ;
		if( nowPage >noM)fir='<li class=g1><a href="?' + page + '=1' + oldUrl + '"  >首页</a></li>';
		if( pMaxNo -nowPage >=noM)last='<li class=g1><a href="?' + page + '=' +pMaxNo+ oldUrl + '"  >最后页</a></li>';
		if( o1.Next >0)next='<li class=g1><a href="?' + page + '=' + o1.Next + oldUrl + '"  >下一页</a></li>';
		if( o1.Pre >0)pre='<li class=g1><a href="?' + page + '=' + o1.Pre + oldUrl + '"  >前一页</a></li>';
		if(pMaxNo >1){
			var t = o1.No ;
			for(var i=0; i<t.length ;i++) {
				if(t[i] >0){	
					if(t[i] !=nowPage){pgNoSt += '<li class=n0><a href="?' + page + '=' + t[i] + oldUrl + '"  >'+t[i]+'</a></li>';}
					else {pgNoSt += '<li class=n1>'+t[i]+'</li>';};
				};
			};
		};
		var other= '<li>('+nowPage+'/'+pMaxNo+')</li>';
		var st= fir + pre + pgNoSt + next+ last + other ;
		return st  ;
	};
	this.pageList_dbPc=function( nowPage, pageno, AllRecord){/*{}*/
		if(nowPage <=0) nowPage=1;
		//var pMaxNo =  (AllRecord % pageno == 0)? (AllRecord / pageno): ( AllRecord / pageno+1) ;
		var pMaxNo =parseInt(AllRecord / pageno) ;
		if( AllRecord % pageno != 0)pMaxNo=pMaxNo +1;
		var next= (nowPage+1) <= pMaxNo ? (nowPage+1):0 ;
		var pre= (nowPage-1) >= 1 ? (nowPage-1):0 ;
		var first = nowPage==1?0:1 ;
		var last = nowPage==pMaxNo ?0:pMaxNo ;
		var no =[0,0,0,0,0];
		var noM= 5 ;
		if(pMaxNo <=noM){
			for(var i=0; i<pMaxNo ;i++){ 
				if(i<pMaxNo)no[i]= i+1 ;
				else {	no[i]= -1 ;};
			};
		};
		if(pMaxNo > noM && (nowPage-2)>=1 && (nowPage+2) <= pMaxNo){
			for(var i=0; i<noM ;i++) no[i]=(nowPage-2) +i;
		};
		if(pMaxNo > noM && (nowPage-2) <=0 ){
			for(var i=0; i<noM ;i++) no[i]= i+1 ;
		};
		if(pMaxNo > noM && (nowPage+2) > pMaxNo ){
			for(var i=0; i<noM ;i++) no[i]= pMaxNo-4+i ;
		};
		var j={"MaxNo":pMaxNo,"Next":next,"Pre":pre,"First":first,"Last":last,"No":no };
		return j;
	};
	/* page string no*/
	this.CheckBoxList=function(ckName){
		var aID="";
	   try{
		   var r=document.getElementsByName(ckName);  
		   for (i = 0; i < r.length; i++) {
			   if(r[i].checked){
				   if(aID =="")aID = r[i].value ;
				   else {aID +=","+ r[i].value ;};
			   };
		   } ;
	   }catch(e){};
	   return aID;
	};
	/* js ad gun dong */
	
		/* other */
};
/* data */
var JL_pgData = function()
{
	this.IsIE=function(){
		if(!+[1,])return true;
		return false;	
	};
	/* this 2way */
	this.Save=function(n, jsdb ){
		if(!this.IsIE())this.DbJsonSet(n, jsdb );
		//else
		 this.CookieJsonSet(n, jsdb, 30*24*60 );
	};
	this.Get=function(n){
		if(!this.IsIE()){
			var v =this.DbJsonGet(n);
			if( v!=null) return v ;
		};
		//else
		 return this.CookieJsonGet(n);
		return null;
	};
	this.Del=function(n){
		if(!this.IsIE())this.DbDel(n);
		//else 
		this.CookieJsonDel(n);
	};
	/*------------localStorage:eg:var udb =  {"n":n,"id":uid,"psw":psw};		db.DbJsonSet("jlLogin",udb); */
	this.DbJsonGet=function(n){
		/*localStorage.lastname="Smith";document.write("Last name: " + localStorage.lastname);	*/
		try{
		   var a= localStorage.getItem(n);
		   if(a==null || a=="")return null;
		   return  JSON.parse(a);
		}catch(Exception){};
		return null ;
	};
	this.DbJsonSet=function(n,v){
		try{ localStorage.setItem(n,JSON.stringify(v));	}catch(Exception){};
	};
	this.DbDel=function(n){
		try{ localStorage.removeItem(n );	}catch(Exception){};
	};
	/* //cookies json type */
	this.CookieJsonDel = function(name ) {
		this.CookieJsonSet(name,"",1);
		this.CookieDel(name);
	};

	this.CookieJsonGet = function(name) {
		try {
			var a= this.CookieGet(name);
			//var a= (document.cookie.match(new RegExp("(^" + name + "| " + name + ")=([^;]*)")) == null) ? "" : decodeURIComponent(RegExp.$2);
			if(a==null || a=="")return null;
			return JSON.parse(a);
		} catch (e) {return null;		};
	};
	this.CookieJsonSet =function (name2, value2,minuTime)
	{
		var v =JSON.stringify(value2) ;
		this.CookieSet(name2, v,minuTime);
	};
	
	//cookies 
	this.CookieDel = function(name2 ) {
	try{ 
		//var exp2 = new Date();
		//exp2.setTime(exp2.getTime() -60000 );//5 min
		var exp2 = new Date(new Date().getTime());
			//document.cookie = name + "= ; path=/; domain=" + domain + "; expires=" + expireDate.toGMTString();
		if(!+[1,])document.cookie = name2 + "= ; expires=" + exp2.toGMTString();
		else document.cookie = name2 + "= ; path=/; expires=" + exp2.toGMTString();
		}catch(Exception){};
	};
	this.CookieGet = function(name2) {
	 try {
			return (document.cookie.match(new RegExp("(^" + name2 + "| " + name2 + ")=([^;]*)")) == null) ? "" : decodeURIComponent(RegExp.$2);
		} catch (e) {return '';	};
	};
	this.CookieGetG = function(name2,eName) {
		try {
			var sorStr=this.CookieGet(name2);
			// (document.cookie.match(new RegExp("(^" + name2 + "| " + name2 + ")=([^;]*)")) == null) ? "" : decodeURIComponent(RegExp.$2);
			if(sorStr=="")return '';
			var panStr = eName+ "=";
			var itmp = sorStr.indexOf(panStr);
			    if (itmp < 0) { return ""; }
				sorStr = sorStr.substr(itmp + panStr.length);
				itmp = sorStr.indexOf("&");
				if (itmp < 0) {	return sorStr;}
				else {
					sorStr = sorStr.substr(0, itmp);
					return sorStr;
				};
			return "";
		} catch (e) {return '';	};
	};
	this.CookieSet =function (name2, value2,minuTime)
	{
		try {
			 var exp2 = new Date();
		exp2.setTime(exp2.getTime() + minuTime*60*1000);//5 min
		if(!+[1,]) document.cookie = name2 + "="+ escape (value2.toString()) + ";expires=" + exp2.toGMTString();
		else document.cookie = name2 + "="+ escape (value2.toString()) + "; path=/;expires=" + exp2.toGMTString();
		}catch(Exception){};
	};

};

/*------   */

var JL_pgComm = function()/*use jquery */
{
	/* scrol gun dong wp0.wpNameNo */
	var wp0={boxW_min:302,boxW_max:540,time:3000,current:0,total:3,maxW:1,wp_ImgOld_w:1,wp_ImgOld_h:1,wpName:"#scroller_",wpNameNo:"0",wp_imgs:null,wp_pos:null, wp_wrap:null};
	this.scrol_init=function(no,w_min,w_max){
		wp0.wpNameNo=no;
		wp0.boxW_min= w_min;
		wp0.boxW_max =w_max ;
		var n1= wp0.wpName+ "lis"+ no ;
		var n2= wp0.wpName+ "pt"+ no ;
		var n3=wp0.wpName+ "db"+ no ;
		wp0.wp_imgs =$(n1).find('li');
		wp0.wp_pos= $(n2).find('em');
		wp0.total= wp0.wp_imgs.length  ;
		wp0.wp_wrap = wp0.wp_imgs.eq(wp0.current);
		wp0.wp_ImgOld_w=$(n3).attr("w");
		wp0.wp_ImgOld_h=$(n3).attr("h") ;
		var oo= $(n1).find('img').eq(0);//aaa
		this.scrol_getSize(oo);//aaa
		
		this.scrol_resize();
		
		//var cb=new JL_Base();
		//cb.Log("dvLog", oo.attr("width")  +"==="+ oo.attr("height")  +"; ");
		//cb.Log("dvLog",  img.width  +"==="+  img.height  +"; ");
	};
	this.scrol_getSize =function(oo){
		var img = new Image(); 
		img.src =oo.attr("src") ;
		var o2= this ;
		img.onload=function(){
			wp0.wp_ImgOld_w=img.width ;
		    wp0.wp_ImgOld_h=img.height ;
			o2.scrol_resize();
			//var cb=new JL_Base();
			//cb.Log("dvLog",  img.width  +"==="+  img.height  +"; ");
		};
	};
	
	this.scrol_Tatal=function(){
	  var n1= wp0.wpName+ "lis"+ wp0.wpNameNo ;
	  return $(n1).find('li').length;//	wp0.total ;
	};
	/*屏幕尺寸变化时*/
	this.scrol_resize=function(){
		wp0.maxW=this.scrol_imgWH();
		var imgH= parseInt(wp0.maxW * wp0.wp_ImgOld_h / wp0.wp_ImgOld_w );
		wp0.wp_imgs.width(wp0.maxW ).height(imgH );
		var n1= wp0.wpName + "lis"+ wp0.wpNameNo ;
		$(n1).width(wp0.maxW ).height(imgH );//.css("width", imgW+'');//.width(imgW+'px');
		//var jb=new JL_Base();
		//jb.Log("dvLog",imgW +","+ imgH +"; ");
	};
	this.scrol_run=function(){
		if(wp0.total <= 1) return ;
		var i= wp0.current ;
		i++;
		this.scrol_goto(i);
	};
	this.scrol_run2=function(){
		//var jb=new JL_Base();
		//jb.Log("dvLog",wp0.total+","+ wp0.current+"; ");
		var ixo = wp0.current ;//wp_wrap= wp_imgs.eq(ixo);
		wp0.current++ ;
		var ix=wp0.current ;
		if(ix >= wp0.total )ix=0 ;
		wp0.current= ix;
		//try{
			var o= this ;
			var wp_wrap_next =wp0.wp_imgs.eq(ix);// $wp_imgs.eg(current);
			var aniNext = {left: 0}  , aniPrev = {left: -wp0.maxW};
			wp0.wp_wrap.css({left:0 }).show().fadeIn().attr("class","on").css("opacity", "1").css("z-index", "2");
			wp_wrap_next.css({left :wp0.maxW }).show().fadeIn().attr("class","on").css("opacity", "1").css("z-index", "1");;
			wp_wrap_next.animate(aniNext, 1000);//wp_wrap_next.stop().animate(aniNext, 1000);
			wp0.wp_wrap.animate(aniPrev, 1000, function(){ //wp_wrap.stop().animate(aniPrev, 1000, function(){
				$(this).hide().attr("class","");
				wp0.wp_wrap = wp_wrap_next;
				wp0.wp_pos.eq(ixo).attr("class","");
				wp0.wp_pos.eq(ix).attr("class","on");
				
				//setTimeout(function () {o.scrol_run() } , wp0.time );
			});
			//if(!+[1,]) setTimeout( 'this.scrol_run()'  ,wp0.time );
			//else setTimeout(function () {o.scrol_run() } , wp0.time );
		//}catch(e){}
	};
	/*直接到某一个*/
	this.scrol_goto=function(i){
		var ixo = wp0.current ;//wp_wrap= wp_imgs.eq(ixo);
		//wp0.current =i;
		if(i==ixo) return ;
		var ix= i;// wp0.current ;
		if(ix >= wp0.total || ix<0 )ix=0 ;
		wp0.current= ix;
		//try{
			var o= this ;
			var wp_wrap_next =wp0.wp_imgs.eq(ix);// $wp_imgs.eg(current);
			var aniNext = {left: 0}  , aniPrev = {left: -wp0.maxW};
			wp0.wp_wrap.css({left:0 }).show().fadeIn().attr("class","on").css("opacity", "1").css("z-index", "2");
			wp_wrap_next.css({left :wp0.maxW }).show().fadeIn().attr("class","on").css("opacity", "1").css("z-index", "1");;
			wp_wrap_next.animate(aniNext, 1000);//wp_wrap_next.stop().animate(aniNext, 1000);
			wp0.wp_wrap.animate(aniPrev, 1000, function(){ //wp_wrap.stop().animate(aniPrev, 1000, function(){
				$(this).hide().attr("class","");
				wp0.wp_wrap = wp_wrap_next;
				wp0.wp_pos.eq(ixo).attr("class","");
				wp0.wp_pos.eq(ix).attr("class","on");
				//setTimeout(function () {o.scrol_run() } , wp0.time );
			});
		//}catch(e){}
	};
	
	this.scrol_imgWH=function(){
		  var w = $(window).width() ;
		 if(w<= wp0.boxW_min ) w=wp0.boxW_min ;//320 
		 if(w>= wp0.boxW_max ) w=wp0.boxW_max ;//540 
		 return w;// { width:w };
	};
	
	
	/*
	this.scrol_start=function(nn){//nn 为id 最后一位相同代号
		var c1= new JL_Comm(); 
		c1.scrol_init(nn);
		setTimeout(  c1.scrol_run() ,1000); 
		//this.scrol_init(nn);	
		//var o=this ;
		//setTimeout(  o.scrol_run() ,wp0.time); 
		//setInterval(this.scrol_run(),scrol_time); 
	}*/
	/* scrol gun dong wp0.wpNameNo */
	
	
};

/*  ----------------------*/

