// JavaScript Document
var JL_pgShPicClass = function()
{
	 var IdAlert="fle_Alert";
	 var PassUrl="";
	 var ck =new JL_Check();
	 var cb=new JL_Base();
 	 var tx =new JL_Text();
	 var jc =new JL_Comm();
	 var regVars = {
			   isDoing: false,
			   timeoutProc: null
	 };
	 var doi=0; var TimeHm=0;
	this.checkActNoOk=function(){
		if(TimeHm==0) TimeHm= cb.time();
		doi ++ ; 
		if( doi <=5 ) return true;//5time
		var df=cb.time() -  TimeHm ;
		if(df <= 300000 ) return false;// 5*60*1000 5min can 5time
		if(df <= 600000 && doi <=10 ) return true;// 10 min 10time
		if(df <= 6000000 && doi <=30 ) return true;//100min 30time
		return false;  
   };
   this.DelCallBack =function(){
	    clearTimeout(regVars.timeoutProc);
		regVars.timeoutProc = null;
		regVars.isDoing =false ;
		jlb.FromAlertG("提交超时，请稍后再操作...",1,1);
   };
	 /* add */
	 this.AddEvAj =function(eid){
		var tit= cb.V("TitPic") ;
		if( cb.VL2(tit)<2 ){
			jlb.FromAlert("请选择图片");
			return false ;
		}
		if(regVars.isDoing) return ;
		var ty= $("#tyid").val(); 
		var oder= $("#aa_OrderID").val(); 
		regVars.isDoing = true ;
	    var pur= "kt2_picact.aspx?act=add&i="+eid  ;
		 $.post(pur,
		{
		  u:tit ,
		  od:oder,
		  ty:ty
		},
		function(data,status){
		  //alert("数据：" + data + "\n状态：" + status);
		  var ret=eval(data );// $.parseJSON( data );
		  ret = ret || {rt:-401,al:'系统繁忙，稍后操作！'};
			//cb.Log("dvLog",JSON.stringify(ret) );
			regVars.isDoing =false;
			if(ret.rt==1){
				jlb.FromAlertG(ret.al,2,1);
				//$("#h_alert").css("opacity", "1").animate({ "opacity": "0" }, 1000 ,function (){ jlb.FromAlertClose();location.reload();});
				location.reload();
				return ;
			}
			jlb.FromAlertG( ret.al  +","+ ret.rt,1,1 );
		  
		});
	
	};
	 //-----------------
	 /* EditEv */
	 this.EditEv =function(i){
		//var tit= cb.V("aa_url"+i) ;
		//if( cb.VL2(tit)<2 ){
		//	jlb.FromAlert("名称不能为空");
		//	return false ;
		//}
		var tit='';
		if(regVars.isDoing) return ;
		var oder=$("#aa_OrderID"+i).val(); 
		regVars.isDoing = true ;
	    var pur= "kt2_picact.aspx?act=edit&i="+i ;
		 $.post(pur,
		{
		  u:tit ,
		  od:oder
		},
		function(data,status){
		  //alert("数据：" + data + "\n状态：" + status);
		  var ret=eval(data );// $.parseJSON( data );
		  ret = ret || {rt:-401,al:'系统繁忙，稍后操作！'};
			//cb.Log("dvLog",JSON.stringify(ret) );
			regVars.isDoing =false;
			if(ret.rt==1){
				jlb.FromAlertG(ret.al,2,1);
				//$("#h_alert").css("opacity", "1").animate({ "opacity": "0" }, 1000 ,function (){ jlb.FromAlertClose();location.reload();});
				//location.reload();
				return ;
			}
			jlb.FromAlertG( ret.al  +","+ ret.rt,1,1 );
		  
		});
		
	};

	/* list del */ 
	this.ListActEv=function(){
		  if(regVars.isDoing ){jlb.FromAlertG("正在操作中...",1,0);return ;
		  }
		  var a= jc.CheckBoxList("fckb");
		  if(a==""){
			  jlb.FromAlertG("请选择操作数据",1,0);return ;
		  }
		   PassUrl= "act=list&ty=4&did=" + a +"&ttt="+cb.time();
		   var u= "kt2_picact.aspx?jsonp=jlpcl.ListActCallBack&"+PassUrl+"&ttt="+cb.time();
		  cb.newScript( u);
		  regVars.isDoing = true ;
		  regVars.timeoutProc = setTimeout(jlpcl.DelCallBack, 10000);
		  //jlb.FromAlertG("正在提交中...",2,1);
	  }
	  this.ListActCallBack =function( ret )  
	  {       
			clearTimeout(regVars.timeoutProc);
			regVars.timeoutProc = null;
			if(!regVars.isDoing) return ;
			regVars.isDoing =false ;
			ret = ret || {rt:-401,al:'系统繁忙，稍后操作！' };
			if(ret.rt==1){
				 jlb.FromAlertG(ret.al,2,1);
				//$("#h_alert").css("opacity", "1").animate({ "opacity": "0" }, 3000 ,function (){ jlb.FromAlertClose();location.reload();});
				location.reload();
				return ;
			};
			jlb.FromAlertG(ret.al,1,0);return ;
			 
	 }; 
	/* list del */  
	
};

var jlpcl = new JL_pgShPicClass();


function UpFileResult(rt,url2){
	$("#aa_url").val(url2);
	$("#imgShow").attr("src",url2);
	$("#imgShow").attr("style",'');
}