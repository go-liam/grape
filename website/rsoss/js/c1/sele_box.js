function CheckAll(checkAllBox) {
    var actVar = checkAllBox.checked;
    var r=document.getElementsByName("contract");  
	   for (i = 0; i < r.length; i++) {
		   r[i].checked = actVar;
	   }
}
//un check
function CheckAllUn() {
   // var actVar = checkAllBox.checked;
    var r=document.getElementsByName("contract");  
	   for (i = 0; i < r.length; i++) {
		   r[i].checked = false;
	   }
}

CheckAllUn();