var osPicClass = function()
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
	   if( ck.checkEmpty(ck.V("fTitle"))){
			ck.Write("dvInfo","真实名不能为空");
			return false ;
		}
	   //purl="act=add&tn="+ ck.V("fTrueName")+"&ln="+ck.V("fLoginName") +"&pw="+psw +"&pf="+ck.V("fPowerFlag")+"&il="+ck.V("fIsLock") +"&ttt="+ck.time() ;
	   return true;
   };
   
   this.Add_callback =function(ret){
	   //alert("back=" + ret );
	  var  gJson = eval("(" + ret + ")");  
	   gJson = gJson || {rt:-1,al:'系统繁忙，稍后操作！'};
	   ck.Write("dvInfo", gJson.al  +","+ gJson.rt );
	   
   };
   
    this.PicUpload_Ckeck =function(){
	   ck.Write("dvInfo","");
	   //purl get url
	  
		 if( (ck.V("f_classId")<=0)){
			ck.Write("dvInfo","请选择分类");
			return false ;
		}
		 if( ck.checkEmpty(ck.V("a_title"))){
			ck.Write("dvInfo","标题不能为空");
			return false ;
		}
		 if( ck.checkEmpty(ck.V("upfile"))){
			ck.Write("dvInfo","请选择上传的图片");
			return false ;
		}
		
	   //purl="act=add&tn="+ ck.V("fTrueName")+"&ln="+ck.V("fLoginName") +"&pw="+psw +"&pf="+ck.V("fPowerFlag")+"&il="+ck.V("fIsLock") +"&ttt="+ck.time() ;
	   return true;
   };
   
   
};

var jj= new osPicClass();