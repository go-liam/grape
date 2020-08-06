var osWorker = function()
{
   var ck =new jlCheck();
   var purl="";//url
   
   var regVars = {
			   isDoing: false,
			   timeoutProc: null
		   };
		   
   this.Add_Bt=function(){
		if(this.Add_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
			
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/account/wpw_add.jsp" ; //
		ajaxobj.content =purl;//
		//alert(ajaxobj.content);
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jj.Json_CallBack(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(this.DelCallBack, 15000);
   };
   this.Add_Ckeck =function(){
	   ck.Write("dvInfo","正在提交中...");
	   //purl get url
	   if( ck.checkEmpty(ck.V("fTrueName"))){
			ck.Write("dvInfo","真实名不能为空");
			return false ;
		}
		if( ck.checkEmpty(ck.V("fLoginName"))){
			ck.Write("dvInfo","登录名不能为空");
			return false ;
		}
		if( ck.checkEmpty(ck.V("fPassWord"))){
			ck.Write("dvInfo","密码不能为空");
			return false ;
		}
		var psw= hex_md5(ck.V("fPassWord"));
	   purl="act=add&tn="+ ck.V("fTrueName")+"&ln="+ck.V("fLoginName") +"&pw="+psw +"&pf="+ck.V("fPowerFlag")+"&il="+ck.V("fIsLock") +"&ttt="+ck.time() ;
	   //ck.Log("dvLog", purl);
	   return true;
   };
   
  
   this.test= function(ret){
      ret = ret || "{rt:-1,al:'系统繁忙，稍后操作！'}";
	  //ret="{ rt:'1',al:'操作成功'}";
	  var  gJson = eval("(" + ret + ")");  
	   //gJson = gJson || {rt:-1,al:'系统繁忙，稍后操作！'};
	   ck.Write("dvInfo", gJson.al  +","+ gJson.rt );
   };
   
   // edit---------------------------
    this.Edit_Ckeck =function(){
	   ck.Write("dvInfo","正在提交中...");
	   if( ck.checkEmpty(ck.V("fTrueName"))){
			ck.Write("dvInfo","真实名不能为空");
			return false ;
	   }
	   purl="act=edit&tn="+ ck.V("fTrueName")+"&wid="+ck.V("fWorkerID") +"&pf="+ck.V("fPowerFlag")+"&il="+ck.V("fIsLock")+"&ttt="+ck.time() ;
	  // ck.Log("dvLog", purl);
	   return true;
   };
    this.Edit_Bt=function(){
		if(this.Edit_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/account/wpw_edit.jsp" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jj.Json_CallBack(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(this.DelCallBack, 15000);
   };
   //psd---------------------------
    this.EditPSW_Ckeck =function(){
	   ck.Write("dvInfo2","正在提交中...");
	  if( ck.checkEmpty(ck.V("fPassWord"))){
			ck.Write("dvInfo2","密码不能为空");
			return false ;
		}
	   var psw= hex_md5(ck.V("fPassWord"));
	   purl="act=editpsw&pw="+ psw+"&wid="+ck.V("fWorkerID")+"&ttt="+ck.time() ;
	  // ck.Log("dvLog", purl);
	   return true;
   };
    this.EditPSW_Bt=function(){
		if(this.EditPSW_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/account/wpw_psw.jsp" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jj.Json_CallBack(xmlobj.responseText,"dvInfo2");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(jj.DelCallBack, 15000);
   };
   //----------------------------------
   
   this.Json_CallBack =function(ret,divID){
	  ret = ret || "{rt:-400,al:'系统繁忙，稍后操作！'}";
	  divID=divID || "dvInfo";
	  clearTimeout(regVars.timeoutProc);
			regVars.timeoutProc = null;
			if(!regVars.isDoing) return ;
			regVars.isDoing =false ;
	    var  gJson = eval("(" + ret + ")");  
	  if(gJson.rt==1){
		  ck.Write(divID, gJson.al);
		  return ;
	  }
	  else{
		   ck.Write(divID, gJson.al  +","+ gJson.rt );
	  }
   };
   this.DelCallBack =function(){
	       clearTimeout(regVars.timeoutProc);
			regVars.timeoutProc = null;
			regVars.isDoing =false ;
   }
   
   //list---------------------------
   this.listAct_Ckeck =function(){
	   ck.Write("dvInfo","正在提交中...");
	  if( (ck.V("acttyle")=="0")){
			ck.Write("dvInfo","请选择操作类型");
			return false ;
		}
	   var aID="";
	   var r=document.getElementsByName("contract");  
	   for (i = 0; i < r.length; i++) {
		   if(r[i].checked){
			   if(aID =="")aID = r[i].value ;
			   else aID +=","+ r[i].value ;
		   }
	   }
  		//ck.Log("dvLog", aID);
	   purl="act=list&ty="+ ck.V("acttyle")+"&did="+aID+"&ttt="+ck.time() ;
	   //ck.Log("dvLog", purl);
	   return true;
   };
    this.listAct_Bt=function(){
		if(this.listAct_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/account/wpw_list.jsp" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jj.Json_CallBack2(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(jj.DelCallBack, 15000);
   };
   this.Json_CallBack2 =function(ret,divID){
	  ret = ret || "{rt:-400,al:'系统繁忙，稍后操作！'}";
	  divID=divID || "dvInfo";
	  clearTimeout(regVars.timeoutProc);
			regVars.timeoutProc = null;
			if(!regVars.isDoing) return ;
			regVars.isDoing =false ;
	    var  gJson = eval("(" + ret + ")");  
	  if(gJson.rt==1){
		  //ck.Write(divID, gJson.al);
		  alert( gJson.al );
		  location.reload() ;//刷新当前页面
		  return ;
	  };
	   ck.Write(divID, gJson.al  +","+ gJson.rt );
   };
   //-----------------------------------power menu
   this.PowerEdit_Ckeck =function(){
	   ck.Write("dvInfo","正在提交中...");
	  var flagPower =ck.V("fPowerFlag") ;
	  var allID ="";
	  if(flagPower=="1"){
		  allID ="";
	  }else{
		  allID = tree.getAllChecked() ; //get
	  }
  	   //ck.Log("dvLog", allID);
	   purl="act=power&dt=1&pf="+ck.V("fPowerFlag")+"&wid="+ck.V("fWorkerID")+"&st="+ allID +"&ttt="+ck.time() ;
	  // ck.Log("dvLog", purl);
	   return true;
   };
    this.PowerEdit_Bt=function(){
		if(this.PowerEdit_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/account/wpw_poweredit.jsp" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jj.Json_CallBack(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(jj.DelCallBack, 15000);
   };
   //-----------------------------------end
   
   
   
   //-----------------------------------end
};

var jj= new osWorker();