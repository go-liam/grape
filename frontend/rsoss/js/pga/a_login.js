
var jlAdminLogin = function()
{
	var Passurl="";
	var doi=0;
	var ck =new jlCheck();
	   var regVars = {
			   isDoing: false,
			   timeoutProc: null
		   };
		   
	this.user_login=function(){
		ck.Write("InfoAlert","");
		var un=$("#username1").val();
		var psw=$("#password1").val();
		var cc=$("#valcode1").val();
		if(un.length <3){
			ck.Write("InfoAlert","账号不能为空");//FromAlert(0,'账号不能为空',1,0);//alert("账号不能为空");
		return;}
		if(psw.length <5){
			ck.Write("InfoAlert","密码为6到20个字符");//FromAlert(0,'密码为6到20个字符',1,0);//alert("密码为6到20个字符");
		return;}
		if(!ck.checkVerify(cc)){
			ck.Write("InfoAlert","验证码有误");
			return false ;
		}
		psw = hex_md5(psw);
		 $.post("ca/main/login.aspx?act=login",
		  {
			u:un,
			p:psw,
			v:cc
		  },
		  function(data,status){
			//alert(data);
			var ret= eval(" "+data+" ");//eval("("+data+")");
			ret = ret || {rt:-401,al:'系统繁忙，稍后操作！'};
			if(ret.rt==1)location.href='main/default.aspx?'+(new Date).getTime();
			else{
				//alert(''+ ret.al);
				//FromAlert(0,ret.al,1,0);
				ck.Write("InfoAlert", ret.al  +","+ ret.rt );
			};
		  });//post end	
	};//reg end
	
	this.LoginSubmit =function(){
		if(this.MyCkeck()==false) return ;
		if(regVars.isDoing) return ;
			regVars.isDoing = true ;
			var u= "ca/main/login.aspx?jsonp=jjLogin.jsonpCallback&"+Passurl+"&ttt="+ck.time() ;
			ck.newScript( u);//返回 jsonpCallback({ rtn:'00000'} )
			//ck.Write("dvLog",u);
			regVars.timeoutProc = setTimeout(this.jsonpCallback, 10000);
	};
	
	
	this.MyCkeck =function(){
		ck.Write("InfoAlert","");
//		if(!ck.checkUserName(ck.V("username1"))){
//			ck.write("InfoAlert","用户名不合法");
//			return false ;
//		}
//		if(!ck.checkPasswd(ck.V("password1"))){
//			ck.write("InfoAlert","密码不合法");
//			return false ;
//		}
		if(!ck.checkVerify(ck.V("valcode1"))){
			ck.Write("InfoAlert","验证码有误");
			return false ;
		}
		Passurl = "act=login&u="+ escape(ck.V("username1")) +"&p="+ hex_md5(ck.V("password1")) +"&v="+ck.V("valcode1") ;
		return true;
	};
	
	this.jsonpCallback =function( ret )  //run one time    
	     {       
			clearTimeout(regVars.timeoutProc);
			regVars.timeoutProc = null;
			if(!regVars.isDoing) return ;
			regVars.isDoing =false ;
			ret = ret || {rt:-401,al:'系统繁忙，稍后操作！'};
			if(ret.rt==1){
				//alert("ok"); 
				top.location='main/default.aspx';
				return ;
			}
			//alert(""+ ret.al);
			//doi ++;
			//try{
			//Passurl = "u="+ ck.V("username1") +"&p="+ ck.V("password1") +"&v="+ck.V("valcode1") ;
			//document.getElementById("InfoAlert").innerHTML = doi + "-"+ret.rtn +"-=-"+Passurl ;//jsonpCallback({"rtn":1})
			ck.Write("InfoAlert", ret.al  +","+ ret.rt );
				//}catch( exception ){}
	     } ; 
};
var jjLogin =new jlAdminLogin();