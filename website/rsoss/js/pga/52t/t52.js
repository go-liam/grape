// JavaScript Document

var os52T = function()
{
	var ck =new jlCheck();
	 this.NoticeEdit_Ckeck =function(){
	   ck.Write("dvInfo","");
	   //purl get url
	  
		if(ck.checkEmpty(ck.V("fBody"))){
			ck.Write("dvInfo","内容不能为空");
			return false ;
		}
		if( ck.checkEmpty(ck.V("ft0"))){
			ck.Write("dvInfo","开始时间不能为空");
			return false ;
		} 
  
	   return true;
   };
	
	
	
}


var jj= new os52T();