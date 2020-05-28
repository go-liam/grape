var osYuLan = function()
{
   var ck =new jlCheck();
   var purl="";//url
   
   this.Add_Bt=function(){
		if(this.Add_Ckeck()==false) return ;
		var ajaxobj = new AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "ca/account/wpw_add.jsp" ; //
		ajaxobj.content =purl;//
		//alert(ajaxobj.content);
		ajaxobj.callback = function(xmlobj) {
			//alert("jsGetInfo back =" + xmlobj.responseText + "<br>" + Date());
			Add_callback(xmlobj.responseText);
		};
		ajaxobj.send(); 
   };
   this.Add_Ckeck =function(){
	   ck.Write("dvInfo","");
	   //purl get url
	   if(!ck.checkEmpty(ck.V("fTrueName"))){
			ck.write("dvInfo","真实名不能为空");
			return false ;
		}
		if(!ck.checkEmpty(ck.V("fLoginName"))){
			ck.write("dvInfo","登录名不能为空");
			return false ;
		}
		if(!ck.checkEmpty(ck.V("fPassWord"))){
			ck.write("dvInfo","密码不能为空");
			return false ;
		}
		var psw= hex_md5(ck.V("fPassWord"));
	   purl="act=add&tn="+ ck.V("fTrueName")+"&ln="+ck.V("fLoginName") +"&pw="+psw +"&pf="+ck.V("fPowerFlag")+"&il="+ck.V("fIsLock") +"&ttt="+ck.time() ;
	   return true;
   };
   
   this.Add_callback =function(ret){
	   //alert("back=" + ret );
	  var  gJson = eval("(" + ret + ")");  
	   gJson = gJson || {rt:-1,al:'系统繁忙，稍后操作！'};
	   ck.write("dvInfo", gJson.al  +","+ gJson.rt );
   };
   
  //-----------------------------------------------
   this.Edit_Ckeck =function(){
	   ck.Write("dvInfo","");
	   //purl get url
	   if( ck.checkEmpty(ck.V("fTitle"))){
			ck.Write("dvInfo","标题不能为空");
			return false ;
		}
		if(ck.checkEmpty(ck.V("fBody"))){
			ck.Write("dvInfo","内容不能为空");
			return false ;
		}
		if( ck.checkEmpty(ck.V("ft0"))){
			ck.Write("dvInfo","开始时间不能为空");
			return false ;
		} 
		if( ck.checkEmpty(ck.V("ft1"))){
			ck.Write("dvInfo","结束时间不能为空");
			return false ;
		}  
		 if( !ck.isNumber(ck.V("yMoney"))){
			ck.Write("dvInfo","价钱为数值");
			return false ;
		}
		if( !ck.isNumber(ck.V("yMoneyOld"))){
			ck.Write("dvInfo","原价钱为数值");
			return false ;
		}
	   return true;
   };
   
   
};

var jj= new osYuLan();