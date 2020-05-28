/*page base function ui,sh */
var JL_System = {
	 Width:0 ,
	 Height:0
	 ,Url:"/sh/" /*path url rs,v,shm */
	 ,UrlDb:"/sh/" /*out db url ca*/
	 ,UName:''
	 ,UID:0
	 ,UPsw:''
	 ,UFlag:1
	 ,SiteSt:"j" /*site flag string:cookies */
};

/*page base function */
var JL_pgBase = function()
{
	/*v1 */
	this.Page_init=function(){
		var cb=new JL_Base();
		var ieN=cb.IEv();/*ie v小于8时，不支持json格式化，要加入*/
		if( ieN >0 &&  ieN <8) cb.newScriptV(JL_System.Url+"rs/mp/js/o/json2.js","",1);
		this.footer_init();
	}
	
	/*user info */	
	var dbName=JL_System.SiteSt+"ui";
	this.UserDb=function(){
	    if( this.IsLogin()==false) return null;
	    var db= new JL_pgData();
		var a= db.Get(dbName);
		return a;
   };
   this.IsLogin=function(){
	 var db= new JL_pgData();
	  var n2= db.Get(dbName);
	  if(n2==null) return false ;
	  try{ 
	  	if( parseInt(n2.id )>=1 ) return true ; }
	  catch(Exception){ return false ; }
	  return false ;
   }
   this.BindUserInfo=function(){
	   if( this.IsLogin()==false) return false;
	   if(JL_System.UID >0) return true;
	   var a= this.UserDb();
	   JL_System.UName=a.n;
		JL_System.UID= parseInt(a.id );
		JL_System.UPsw=a.psw;
		return true;
   };
   this.PgCheck=function(){/*是否登录，没转到登录*/
	  if( this.IsLogin()) return true; 
	  var cb=new JL_Base();
	  var g= JL_System.Url+'user.htm?pg=logout&al=4';
	  cb.GoUrl(g);
	  return false;
   }
   /* alert top */
    this.FromAlert_b=function(info,ty,IsSHow){/*提示内容; 类型：1红，绿2; 1固定显示*/
		alert( info );
		return ;
		var o=$("#h_alert");
	    $("#h_alert_tit").text(info);
	   o.clearQueue().stop().show(0).css("opacity", "1");
	   if(ty==1)o.attr("class","h_alert1");//o.removeClass("h_alert2").addClass("h_alert1");
	   else if(ty==2)o.attr("class","h_alert2");//o.removeClass("h_alert1").addClass("h_alert2");
	   else o.attr("class","h_alert3");
	   //if(o.is(":animated") )return ;
	   $("#h_info").hide();
	   //o.fadeOut(4000,function (){ $("#h_info").show();});
	   if(IsSHow !=1) o.animate({ "opacity": "0" }, 4000 ,function (){ $("#h_info").show();o.hide();});
	}
	this.FromAlertG=function(info,ty,IsSHow){
	   this.FromAlert_b(info,ty,IsSHow);/*1 Alert red, 2 OK green */
   }
   this.FromAlert=function(info){
	   this.FromAlert_b(info,1,0);/*1 Alert red, 2 OK green */
   }
   this.FromAlertClose=function(){
	  var o=$("#h_alert");
	  o.clearQueue().stop().hide();
	  $("#h_info").show();
   }
   /* alert top end*/
   /* other*/
   this.GoTop=function(){
	   if (document.documentElement && document.documentElement.scrollTop) {  
            document.documentElement.scrollTop =0;  
        }  
        else if (document.body) {    document.body.scrollTop =0; }
   }
   this.Back=function(){
		 window.history.go(-1);
   }
   this.BackAndRl=function(){
		window.history.back();
   }
   /* page */
   this.footer_init=function(){
		var b=this.BindUserInfo();
		var cb=new JL_Base();
		if(b){ cb.Write("fn_uinfo",'<a href="'+JL_System.Url+'shm.htm?pg=info" >'+JL_System.UName+'</a><a href="'+JL_System.Url+'shm.htm?pg=logout" >退出管理</a>');}
		 else{cb.Write("fn_uinfo",'<a  href="'+JL_System.Url+'shm.htm?pg=shlogin" >登录</a><a href="'+JL_System.Url+'shm.htm?pg=shreg">注册</a>');}
   }
   this.header_init=function(tit){
	   var cb=new JL_Base();
	   if(tit !="") cb.Write("fn_uinfo",tit);// $("#hi_tit").text(tit);
   }
   
};
var jlb=new JL_pgBase();

