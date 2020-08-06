function CheckAll(checkAllBox) {
    var actVar = checkAllBox.checked;
    var r=document.getElementsByName("contract");  
	   for (i = 0; i < r.length; i++) {
		   r[i].checked = actVar;
	   }
}


var frm = document.getElementById("fmList"); //  document.form1;
function CheckAll99(checkAllBox) {
    var actVar = checkAllBox.checked;
    for (i = 0; i < frm.length; i++) {
        e = frm.elements[i];
        if (e.type == 'checkbox' && e.name.indexOf("contract") != -1)
            e.checked = actVar;
    }
}
function UnCheck() {
    for (i = 0; i < frm.length; i++) {
        e = frm.elements[i];
        if (e.type == 'checkbox' && e.name.indexOf("checkAll") != -1) {
            e.checked = false;
            break;
        }
    }
}