var JL_SystemPic = {
	SelLiOj:null ,
	dbTid:0 
}

var osPic = function()
{
   var ck =new jlCheck();  
  //------------------------------------Pic Select
   this.DialogClose =function(id2){
	   ck.Display(id2,0);
   }
   this.DialogOpen =function(oj){
	   var id2='pg_piclist' ;
	   ck.Display(id2,1);
	   JL_SystemPic.SelLiOj = oj;
   }
   
   this.Event_ClassChange = function(){
	   var checkValue=$("#fpClass").val(); 
	   alert("this.value= "+checkValue );//
   }
   
   this.Event_Ok = function(){
	   //alert("this.id--ok ");
	   this.DialogClose("pg_piclist");
   }
   
   this.Event_LiClick = function(txt,imgurl){
	   //alert("this.id--ok ");
	   //alert(txt+ ",imgurl="+ imgurl);
	   ck.Write("dvSelName",txt);
	   ck.G("DImgUrl").value= imgurl ;
   }
   
   this.Init=function(){
	   //document.getElementById("fpbtok").onclick=jp.Event_ClassChange() ;
	   ck.G("fmbShop").value= ck.G("dvMbShop").innerHTML ;
	   
   }
   //------------------------------------------------------Shop 
   this.Shop_Add=function(){
		var aa='<div class="dl" >' +ck.G("fmbShop").value +'</div>';//'<div><a onClick="jjpic.Shop_Del(this);" >删除</a></div>'
		//var old=$("#dvShopList").html() ;
		//ck.G("dvShopList").innerHTML += aa ; //.html( old+ aa );
		$("#dvShopList").append( aa) ;
   }
   
   this.Shop_Del=function(oj){
		 var feIid =$(oj).parent().find("[name='feIid']").val()  ;
		 //alert(feIid);
		 //del data
		 if(feIid >0){
			 $.post("../../ca/ufw/ktp_add.aspx",
  			 {
    			act:"delshop",
				iid:feIid 
				},
  				function(data,status){
					//alert("Data: " + data + "\nStatus: " + status);
 			 });//end post
		 }//end if
		 //remove html
		  $(oj).parent().parent(".dl").remove(); 
   }
   
   this.Shop_Save =function(){
		   //var shop='' ,WangW ;
		   $("#dvShopList").children(".dl").each(function(){
		      jjpic.Shop_Save_each(this) ;
		   });	   
   }
   this.Shop_Save_each =function(oj){
			   var feShop = $(oj).find("[name='feShop']").val()  ;
			   var feWangW = $(oj).find("[name='feWangW']").val()  ;
			   var feTit = $(oj).find("[name='feTit']").val()  ;
			   var fePrice = $(oj).find("[name='fePrice']").val()  ;
			   var fePriDec = $(oj).find("[name='fePriDec']").val()  ;
			   var feOther = $(oj).find("[name='feOther']").val()  ;
			   var feOrder = $(oj).find("[name='feOrder']").val()  ;
			   var fePicUrl = $(oj).find("[name='fePicUrl']").val()  ;
			   var feUrl =  $(oj).find("[name='feUrl']").val()  ;
			   var feIid =  $(oj).find("[name='feIid']").val()  ;
			   
			$.post("../../ca/ufw/ktp_add.aspx",
  			{
    			act:"addeach",
				feShop:ck.TextPost(feShop),
				feWangW:ck.TextPost(feWangW),
				feTit:ck.TextPost(feTit),
				fePrice:ck.TextPost(fePrice),
				fePriDec:ck.TextPost(fePriDec),
				feOther:ck.TextPost(feOther) ,
				feOrder:ck.TextPost(feOrder),
				fePicUrl:ck.TextPost(fePicUrl) ,
				tid:JL_SystemPic.dbTid ,
				feUrl:ck.TextPost(feUrl ),
				feIid:feIid 
  			},
  			function(data,status){
    			//alert("Data: " + data + "\nStatus: " + status);
				var  gJson = eval("(" + data + ")");  
				gJson = gJson || {rt:-1,al:'系统繁忙，稍后操作！'};
				//ck.Write("dvInfo", gJson.al  +","+ gJson.rt );
 			 });
			   
			   //alert( $(oj).html() );
			   //alert( o1 );
			   return true ;
	}
   //----------------------------------------------------- Add
   //save post
   this.Data_Add_check =function(){
	   if(  ck.checkEmpty( ck.V("fTitle") )){
		   ck.Write("dvInfo","标题不为空。");    return false;
	   }
	   if(  ck.checkEmpty( ck.V("fBody") )){
		   ck.Write("dvInfo","内容不为空。");    return false;
	   }
	   return true;
   }
   var FlagDo1=false;
   this.Data_Add_Save =function(){
	   if(FlagDo1) return ;
	   ck.Write("dvInfo","正在操作中...");  
	   if(!this.Data_Add_check()) return ;
	   var fBody= ck.V("fBody");
	   var fTitle= ck.V("fTitle");
	   var fTishi= ck.V("fTishi");
	   var fXuZhi= ck.V("fXuZhi");
	   var ft0= ck.V("ft0");
	   var ft0b= ck.V("ft0b");
	   var ft1= ck.V("ft1");
	   var ft1b= ck.V("ft1b"); 
	   FlagDo1=true ;  
	   $.post("../../ca/ufw/ktp_add.aspx",
  		{
    		act:"addbig",
    		fBody:ck.TextPost(fBody) ,
			fTitle:ck.TextPost(fTitle),
			fTishi:ck.TextPost(fTishi),
			fXuZhi:ck.TextPost(fXuZhi) ,
			ft0:ft0,
			ft0b:ft0b ,
			ft1:ft1 ,
			ft1b:ft1b 
  		},
  		function(data,status){
    		//alert("Data: " + data + "\nStatus: " + status);
			var  gJson = eval("(" + data + ")");  
	   		gJson = gJson || {rt:-1,al:'系统繁忙，稍后操作！'};
			FlagDo1=false ;
			if( gJson.rt >=1 ){
				JL_SystemPic.dbTid =gJson.rt ;
				jjpic.Shop_Save();
				ck.Write("dvInfo", "操作成功！" );
				alert("添加成功，你可以在管理页面进行修改。");
				return ;
			 }
			
	   		ck.Write("dvInfo", gJson.al  +","+ gJson.rt );
			
 		 });
   }
   
   //--------------------------------------------------Edit
   this.Data_Edit_Save =function(){
	   if(FlagDo1) return ;
	   ck.Write("dvInfo","正在操作中...");  
	   if(!this.Data_Add_check()) return ;
	   var fBody= ck.V("fBody");
	   var fTitle= ck.V("fTitle");
	   var fTishi= ck.V("fTishi");
	   var fXuZhi= ck.V("fXuZhi");
	   var ft0= ck.V("ft0");
	   var ft0b= ck.V("ft0b");
	   var ft1= ck.V("ft1");
	   var ft1b= ck.V("ft1b");
	   var gid= ck.V("gid") ;
	   JL_SystemPic.dbTid = gid ;
	   FlagDo1=true ;  
	   $.post("../../ca/ufw/ktp_add.aspx",
  		{
    		act:"editbig",
    		fBody:ck.TextPost(fBody) ,
			fTitle:ck.TextPost(fTitle),
			fTishi:ck.TextPost(fTishi),
			fXuZhi:ck.TextPost(fXuZhi) ,
			ft0:ft0,
			ft0b:ft0b ,
			ft1:ft1 ,
			ft1b:ft1b ,
			gid:gid 
  		},
  		function(data,status){
    		//alert("Data: " + data + "\nStatus: " + status);
			var  gJson = eval("(" + data + ")");  
	   		gJson = gJson || {rt:-1,al:'系统繁忙，稍后操作！'};
			FlagDo1=false ;
			if( gJson.rt >=1 ){ 
				jjpic.Shop_Save();
				ck.Write("dvInfo", "操作成功！" );
				alert("操作成功");
				location.reload() ;//刷新当前页面
				return ;
			}
			
	   		ck.Write("dvInfo", gJson.al  +","+ gJson.rt );
			
 		 });
   }
   
   
   //--------------------------------------------------
};

var jjpic= new osPic();


//jquery
$(document).ready(function(){
	var jp = new osPic();
	 var ck =new jlCheck();
	 jp.Init();
	//dialog show ------------------------------- picList
	$("#fpbtok").click (function(){	Event_Bt();});//click 
	$("#fpClass").change (function(){	Event_ClassChange();});//click 
	
	function bindLi_event(){
		var OldSel=null ;
		$("#dvPlist").children("li").each(function(){
				$(this).click(function(){
					var txt= $(this).children("dt").text();
					var url= $(this).children("dd").children("img").attr("src");
					 //addClass toggleClass
					$(this).removeClass("off");
					$(this).addClass("on");
					if( OldSel !=null && OldSel != this ){ $(OldSel).addClass("off"); $(OldSel).removeClass("on");}
					OldSel = this ;
					//alert(txt+ ",url="+ url);
					jp.Event_LiClick(txt,url);
				});
			});
	}//bindLi_event()
	
	function Event_Bt(){
		var ff= $(JL_SystemPic.SelLiOj).parent();
		var ff0= $(JL_SystemPic.SelLiOj).parent().parent();
		var purl= ck.V("DImgUrl")  ;
		var txt ="【"+ ck.VH("dvSelName")  + "】" ;
		
		//alert("V="+purl+",T="+ txt );
		if(purl=="" || purl.length < 8 ) {//|| purl.length < 8
			alert("请先选择图片");
			return ;
			}
		$(ff0).find(".l-img").attr("src",purl) ;
		$(ff).children(".c-imgtxt").text( txt );
		$(ff).children(".i1pic").val(purl);
		//alert($(ff).text( );
		jp.Event_Ok();
	}
	
	function getPicList(cid2 ){
  		$.post("../../ca/ufw/ktp_picget.aspx",
  		{
    		act:"top",
    		cid:cid2 ,
			path:"../../../"
  		},
  		function(data,status){
    		//alert("Data: " + data + "\nStatus: " + status);
			$("#dvPlist").html( data );
			bindLi_event();
 		 });
	};
	//class change
	function Event_ClassChange(){
	   var cid=$("#fpClass").val(); 
	   //alert("this.value= "+cid);//
	   getPicList(cid);
   }
	//bind data
	getPicList(0);
	//-----------------------------------------------------picList end
	
	
	
 })//big function

