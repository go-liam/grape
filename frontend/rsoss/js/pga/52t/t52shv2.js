// JavaScript Document

var os52T = function()
{
	var ck =new jlCheck();
	 this.ShopEdit_Ckeck =function(){
	   ck.Write("dvInfo","");
	   //purl get url
	  
		if(ck.checkEmpty(ck.V("fInfoBody"))){
			ck.Write("dvInfo","内容不能为空");
			return false ;
		}
		if( ck.checkEmpty(ck.V("fShopName"))){
			ck.Write("dvInfo","名称不能为空");
			return false ;
		} 
  
	   return true;
   };
	
	
	
}


var jj= new os52T();