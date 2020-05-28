////常用一些js
//var KK_LOGIN_REG = new jlCheck();
//onclick="KK_LOGIN_REG.regSubmit4();"

var jlCheck = function()
{
	this.pageUrl=function(){
		return window.location.href.toLowerCase();
	};
	this.V = function(id2){
		try{
		return document.getElementById(id2).value;
		}catch(Exception){}
		return "";
	};
	this.$ = function(id2){
		return document.getElementById(id2);
	};
	this.G = function(id2){
		return document.getElementById(id2);
	};
	this.VH = function(id2){
		return document.getElementById(id2).innerHTML;
	};
	this.Write= function(id,body2){
		document.getElementById(id).innerHTML = body2+"";
	};
	this.Log= function(id,body2){
		document.getElementById(id).innerHTML += body2+"<br>";
	};
	this.Display=function(id,IsShow){
		try{
			if(IsShow==1) this.G(id).style.display="block";
			else this.G(id).style.display="none";
		}catch(Exception){}	
	}
//	var regVars = {
//			   isDoing: false,
//			   timeoutProc: null
//		   };
	this.newScript = function(url){
		var obj = document.createElement("script");
		obj.language="javascript";
		obj.type="text/javascript";
		obj.charset = "utf-8";
		obj.src = url;
		document.body.appendChild(obj);
	};
	this.time = function(){
		return (new Date).getTime();
	};
	this.checkEmpty = function(str){
		return /^ *$/.test(str);
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
	    var patrn = /.{6,15}/;
	    if (patrn.test(s1)) {
	        return true;
	    }
	    else {
	        return false;
	    };
	};
	
	this.checkSafePassword=function (str,uname) {
	    if (str == uname)
	        return 1;
	    for (var i = 0; i < str.length; i++) {
	        if (str.charAt(0) != str.charAt(i))
	            break;
	    }
	    if (i == str.length) return 2;
	    var seqStr = "01234567890";
	    if (seqStr.indexOf(str) != -1) return 3;
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
	    var patrn = /^[a-zA-Z]{1}[a-zA-Z0-9\-_]{5,15}$/;
	    if (patrn.test(s1)) {
	        return true;
	    }
	    else {
	        return false;
	    }
	};
	
	this.checkbadname =function (s1) {
	    var badname = new Array("admin", "administrator");
	    for (var i = 0; i < badname.length; i++) {
	        if (s1.toString().toLowerCase() == badname[i]) {
	            break;
	        }
	    }
	    if (i == badname.length)
	    { return true ;}
	    else
	    { return false; }
	};
	
	this.isNumber=function(oNum) {
	   // if (!oNum) return false;
	    var strP = /^\d+(\.\d+)?$/;
	    if (!strP.test(oNum)) return false;
	    try {
	        if (parseFloat(oNum) != oNum) return false;
	    }
	    catch (ex) {
	        return false;
	    }
	    return true;
	};
	
	/*  public  */
	this.GetQueryValue=function(sorStr, panStr,defaultValue) {
	    /*++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	    Description:     在JavaScript内实现QueryString
	    eg: var   fid   =   GetQueryValue(this_url, 'fid ');
	    ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*/
	    var vStr = defaultValue ;
	    if (sorStr == null || sorStr == "" || panStr == null || panStr == "") return vStr;
	    //panStr = panStr.toLowerCase();
	    //sorStr = sorStr.toLowerCase();
	    panStr += "=";
	    var itmp = sorStr.indexOf(panStr);
	    if (itmp < 0) { return vStr; }
	    sorStr = sorStr.substr(itmp + panStr.length);
	    itmp = sorStr.indexOf("&");
	    if (itmp < 0) {
	        return sorStr;
	    }
	    else {
	        sorStr = sorStr.substr(0, itmp);
	        return sorStr;
	    }
	};
	/*
	this.getCookieStat=function (Name) {
	    var search = Name + "=";
	    if (document.cookie.length > 0) {
	        offset = document.cookie.indexOf(search);
	        if (offset != -1) {
	            offset += search.length;
	            end = document.cookie.indexOf(";", offset);
	            if (end == -1) {
	                end = document.cookie.length;
	            }
	            return unescape(document.cookie.substring(offset, end));
	        } else {
	            return ("");
	        }
	    } else {
	        return ("");
	    }
	};
	 */
	
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
	    }
	};
	
	//校验身份证: return=0 ok.
	this.checkIdcard=function ( idcard ) {
		if(idcard.length <15 )return 10;
		var area = { 11: "北京", 12: "天津", 13: "河北", 14: "山西", 15: "内蒙古", 21: "辽宁", 22: "吉林", 23: "黑龙江", 31: "上海", 32: "江苏", 33: "浙江", 34: "安徽", 35: "福建", 36: "江西", 37: "山东", 41: "河南", 42: "湖北", 43: "湖南", 44: "广东", 45: "广西", 46: "海南", 50: "重庆", 51: "四川", 52: "贵州", 53: "云南", 54: "西藏", 61: "陕西", 62: "甘肃", 63: "青海", 64: "宁夏", 65: "新疆", 71: "台湾", 81: "香港", 82: "澳门", 91: "国外" };
	    //var idcard;
	    var Y, JYM;
	    var S, M;
	    var idcard_array = new Array();
	    idcard_array = idcard.split("");
	    //地区检验
	    if (area[parseInt(idcard.substr(0, 2))] == null) return 4;
	    //身份号码位数及格式检验
	    switch (idcard.length) {
	        case 15:
	            if ((parseInt(idcard.substr(6, 2)) + 1900) % 4 == 0 || ((parseInt(idcard.substr(6, 2)) + 1900) % 100 == 0 && (parseInt(idcard.substr(6, 2)) + 1900) % 4 == 0)) {
	                ereg = /^[1-9][0-9]{5}[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|[1-2][0-9]))[0-9]{3}$/; //测试出生日期的合法性

	            } else {
	                ereg = /^[1-9][0-9]{5}[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|1[0-9]|2[0-8]))[0-9]{3}$/; //测试出生日期的合法性

	            }
	            if (ereg.test(idcard)) return 0;
	            else return 2;
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
	            }
	            if (ereg.test(idcard)) {//测试出生日期的合法性

	                //计算校验位
	                S = (parseInt(idcard_array[0]) + parseInt(idcard_array[10])) * 7 + (parseInt(idcard_array[1]) + parseInt(idcard_array[11])) * 9 + (parseInt(idcard_array[2]) + parseInt(idcard_array[12])) * 10 + (parseInt(idcard_array[3]) + parseInt(idcard_array[13])) * 5 + (parseInt(idcard_array[4]) + parseInt(idcard_array[14])) * 8 + (parseInt(idcard_array[5]) + parseInt(idcard_array[15])) * 4 + (parseInt(idcard_array[6]) + parseInt(idcard_array[16])) * 2 + parseInt(idcard_array[7]) * 1 + parseInt(idcard_array[8]) * 6 + parseInt(idcard_array[9]) * 3;
	                Y = S % 11;
	                M = "F";
	                JYM = "10X98765432";
	                M = JYM.substr(Y, 1); //判断校验位

	                if (M == idcard_array[17]) return 0; //检测ID的校验位 ok
	                else return 3;
	            }
	            else return 2;
	            break;
	        default:
	            return 1;
	            break;
	    }
	};


	this.delCookie = function(name, domain2) {
		var expireDate = new Date(new Date().getTime());
		var domain = domain2 ? domain2 : ((window.location.host.indexOf("xxx1.com") != -1) ? 'xxx1.com' : 'xxx2.com');
		document.cookie = name + "= ; path=/; domain=" + domain + "; expires=" + expireDate.toGMTString();
	};
	this.getCookie = function(name) {
		try {
			return (document.cookie.match(new RegExp("(^" + name + "| " + name + ")=([^;]*)")) == null) ? "" : decodeURIComponent(RegExp.$2);
		} catch (e) {
			return '';
		}
	};
	
	
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
		if(sContent.indexOf("\n") >=0 ){
			sContent=sContent.replace(/\n/g ,"ζǎ");
		};
		if(sContent.indexOf("<br>") >=0 ){
			sContent=sContent.replace(/\<br>/g ,"ζǎ");
		};
		//sContent=sContent.Trim();
 		return sContent; 
 	};
	
	/////crean :去掉特殊字符串
 	this.JsonStringCrean = function(sContent) {
		if(sContent==null){
			return sContent;
	    };
		if(sContent.Contains("\\")){
			sContent=sContent.Replace("\\","\\\\");
		};
		if(sContent.Contains("\'")){
			sContent=sContent.Replace("\'","\\\'");
		};
		if(sContent.Contains("\"")){
			sContent=sContent.Replace("\"","\\\"");
		};
		////去掉字符串的回车换行符 
		sContent = Regex.Replace(sContent, "[\n\r]", "");
		sContent=sContent.Trim();
 		return sContent; 
 	};
	
	//-------------------------------------
	this.AJAXRequest=function()  {
			var xmlObj = false;
			//var CBfunc ;
			var ObjSelf;
			ObjSelf=this;
			try { xmlObj=new XMLHttpRequest; }
			catch(e) {
				try { xmlObj=new ActiveXObject("MSXML2.XMLHTTP"); }
				catch(e2) {
					try { xmlObj=new ActiveXObject("Microsoft.XMLHTTP"); }
					catch(e3) { xmlObj=false; }
				}
			}
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
						}
					}
				};
				if(this.method=="POST") xmlObj.send(this.content);
				else xmlObj.send(null);
			};
		};
			//------------------------------
   			
		//-----------------------------------
};

