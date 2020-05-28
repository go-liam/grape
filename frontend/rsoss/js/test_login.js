var jlLoginReg = function()
{
   var regVars = {
	   isDoing: false,
	   timeoutProc: null
   };
   
   var doi=0;
   
   this.newScript = function(url){
		var obj = document.createElement("script");
		obj.language="javascript";
		obj.type="text/javascript";
		obj.charset = "utf-8";
		obj.src = url;
		document.body.appendChild(obj);
	};
	
	 this.regSubmit4 = function(){
		if(regVars.isDoing) return ;
		regVars.isDoing = true ;
		this.newScript("http://payserver.youxi.kankan.com/user/register?callback=KK_LOGIN_REG.jsonpCallback&r=");//·µ»Ø jsonpCallback({"rtn":1})
		regVars.timeoutProc = setTimeout(this.jsonpCallback, 10000);
	};
	
    this.regSubmit3 =function(){
		if(regVars.isDoing) return ;
		regVars.isDoing = true ;
		this.newScript("/webJLp3/c/user/login.jsp?jsonp=KK_LOGIN_REG.jsonpCallback&t=");//·µ»Ø jsonpCallback({ rtn:'00000'} )
		regVars.timeoutProc = setTimeout(this.jsonpCallback, 10000);
		
		
	};
	
     this.jsonpCallback =function( ret )  //run one time    
     {       
		clearTimeout(regVars.timeoutProc);
		regVars.timeoutProc = null;
		if(!regVars.isDoing) return ;
		regVars.isDoing =false ;
		ret = ret || {rtn:-1};
		doi ++;
		//try{
			document.getElementById("djl").innerHTML = doi + " -- "+ret.rtn ;//jsonpCallback({"rtn":1})
		//}catch( exception ){}
		
     } ; 
};

var KK_LOGIN_REG = new jlLoginReg();