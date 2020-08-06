var osTuank  = function()
{
   var regVars = {
			   isDoing: false,
			   timeoutProc: null
		   };
		   
   var ck =new jlCheck();
   var purl="";//url
   
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
		  this.CheckAllUn();
		  alert( gJson.al );
		  location.reload() ;//刷新当前页面
		  return ;
	  };
	   ck.Write(divID, gJson.al  +","+ gJson.rt );
   };
   //-----------------------------------------
    this.CheckAllUn =function() {
   		try{
			var r=document.getElementsByName("contract");  
			for (i = 0; i < r.length; i++) {
			   r[i].checked = false;
		   }
	   }catch(Exception){}	
	}
   //slist---------------------------
   this.slistAct_Ckeck =function(){
	   ck.Write("dvInfo","正在提交中...");
	   if( (ck.V("acttyle")=="-1" || ck.V("acttyle")=="" )){
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

  
  this.list_yuiLAct_Bt=function(){
		if(this.slistAct_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/ufw/52t/kt1_list.aspx" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jtk.Json_CallBack2(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(jtk.DelCallBack, 15000);
   };
   //-----------------------------------
       this.list_tkAct_Bt=function(){
		if(this.slistAct_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/ufw/ktp_list.aspx" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jtk.Json_CallBack2(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(jtk.DelCallBack, 15000);
   };
   //-----------------------------------
       this.list_TreeAct_Bt=function(){
		if(this.slistAct_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/ufw/ktr_list.aspx" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jtk.Json_CallBack2(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(jtk.DelCallBack, 15000);
   };
   //-----------------------------------
       this.list_PicClassAct_Bt=function(){
		if(this.slistAct_Ckeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/ufw/kpc_list.aspx" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			jtk.Json_CallBack2(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(jtk.DelCallBack, 15000);
   };
   
};

var jtk= new osTuank();