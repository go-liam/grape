var osMyInfo = function()
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
   }
   
   
 //psd---------------------------
    this.EditPSW_Ckeck =function(){
	   ck.Write("dvInfo","正在提交中...");
	  if( ck.checkEmpty(ck.V("fPassWord"))){
			ck.Write("dvInfo","密码不能为空");
			return false ;
		}
	   var psw= hex_md5(ck.V("fPassWord"));
	   purl="act=editpsw&pw="+ psw+"&wid="+ck.V("fWorkerID")+"&ttt="+ck.time() ;
	  // ck.Log("dvLog", purl);
	   return true;
   };
    this.EditPSW_Bt=function(){
		if(this.EditPSW_Ckeck()==false) return ;
		if( regVars.isDoing) return ;
			regVars.isDoing = true ;
		var ajaxobj = new ck.AJAXRequest;    // 
		ajaxobj.method = "POST";   //
		ajaxobj.url = "../../ca/my/wpwm_mypsw.aspx" ; //
		ajaxobj.content =purl;//
		ajaxobj.callback = function(xmlobj) {
			//ck.Log("dvLog", xmlobj.responseText );
			j2.Json_CallBack(xmlobj.responseText,"dvInfo");
		};
		ajaxobj.send(); 
		regVars.timeoutProc = setTimeout(j2.DelCallBack, 15000);
   };
   //----------------------------------
   
   
};

var j2= new osMyInfo();