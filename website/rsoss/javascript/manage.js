/*
Version:	Joekoe CMS 3.0
*/

//window.onerror=function(a,b,c){return true;}

function re_load(){ location.reload(); }

var web_dir="../";
var web_dir_common="common/";
var web_dir_support="support/";
var web_dir_editor="common/editor/";

var client_browse="";
if (document.getElementById)	// IE5+,NN6+	document.getElementById("id");
{ client_browse="ie5"; }
else if (document.all)		// IE4		document.all("id");
{ client_browse="ie4"; }
else if (document.layers)	// NN4		document.layers["id"];
{ client_browse="ns4"; }

function get_id(strname)
{
  switch (client_browse)
  {
    case "ie5":
      return document.getElementById(strname);
      break;
    case "ns4":
      return document.layers[strname];
      break;
    default :	//"ie4"
      return document.all(strname);
      break;
  }
}

function form_value(strer)
{
  if (strer!="")
  {
    if (strer.indexOf(".")==-1) strer="all."+strer;
    return eval("document."+strer+".value;");
  }
}

function click_return(strvar,strt)
{
  var tmpvar='';
  switch (strt)
  {
    case 1:
      tmpvar='您确定要'+strvar+'吗？';
      break;
    case 5:
      tmpvar=strvar;
      break;
    default :
      tmpvar='您确定要'+strvar+'吗？\n\n执行该操作后将不可恢复！';
      break;
  }
  var tmptrue=window.confirm(tmpvar);
  if (tmptrue) { return true; }
  return false;
}

function open_wins(url,name,width,height,scroll)
{
  var nwidth=width;
  var nheight=height;
  var Left_size=0;
  var Top_size=0;
  if (nwidth==0 || nheight==0)
  {
    nwidth=screen.width-8;
    nheight=screen.height-55;
  }
  else
  {
    Left_size=(screen.width) ? (screen.width-nwidth)/2 : 0;
    Top_size=(screen.height) ? (screen.height-nheight)/2 : 0;
  }
  var open_wins=window.open(url,name,'width=' + nwidth + ',height=' + nheight + ',left=' + Left_size + ',top=' + Top_size + ',toolbar=no,location=no,directories=no,status=yes,menubar=no,scrollbars=' + scroll + ',resizable=yes' );
}

function open_win(url,name,width,height,scroll)
{
  var Left_size = (screen.width) ? (screen.width-width)/2 : 0;
  var Top_size = (screen.height) ? (screen.height-height)/2 : 0;
  var open_win=window.open(url,name,'width=' + width + ',height=' + height + ',left=' + Left_size + ',top=' + Top_size + ',toolbar=no,location=no,directories=no,status=yes,menubar=no,scrollbars=' + scroll + ',resizable=no' );
}

function disp(str1,str2,str3)
{
  return str1.replace(str2,str3);
}

function is_int(itp)
{
  var knum=event.keyCode;
  var ktrue=false;
  if (knum >= 48 && knum <= 57) { ktrue=true; }
  if (itp==1)
  {
    if (!ktrue && knum==46) { ktrue=true; }
  }
  return ktrue;
}

// onKeyPress="event.returnValue=on_int(0);"
function on_int(strer)
{
  var knum=event.keyCode;
  var ktrue=false;
  if (knum >= 48 && knum <= 57) { ktrue=true; }
  if (strer==1)
  {
    if (!ktrue && knum==46) { ktrue=true; }
  }
  return ktrue;
}

function insert_value(strinput,strvalue)
{
  var tmpstr="";
  switch (client_browse)
  {
    case "ie5":
      tmpstr="document.all."+strinput+".value='"+strvalue+"';";
      break;
    case "ns4":
      tmpstr="document.layers.[\""+strinput+"\"].value='"+strvalue+"';";
      break;
    default :	//"ie4"
      tmpstr="document.all."+strinput+".value='"+strvalue+"';";
      break;
  }
  if (tmpstr!="") eval(tmpstr);
}

function frm_delete_pic(strinput)
{
  
  var tmpfilename=get_id(strinput).value;
  alert(strinput +"="+tmpfilename);
}

function frm_preview(strinput,strcode)
{
  if (strcode==null) strcode=1;
  open_win('my_change.asp?action=preview&remark='+strinput+'&cod='+strcode,'win_preview',600,525,'yes');
}

function textarea_resize(strremark,strt,strform)
{
  if (strremark=='' || strremark==null) return;
  if (strform=='' || strform==null) strform='frm_input';
  var tmpnum=5;
  if (strt=='-') { tmpnum=-5; }
  var tmpobj=eval("document."+strform+"."+strremark);
  if (parseInt(tmpobj.rows)+tmpnum>=3) { tmpobj.rows = parseInt(tmpobj.rows) + tmpnum; }
  if (tmpnum>0) { tmpobj.width="90%"; }
}

function select_color_box(strinput,strcolor)
{
  var tmpid="tbl_color_"+strinput;
  if (strcolor=="" || strcolor==null) strcolor="000000";
  document.write("<table border=0 cellspacing=0 cellpadding=0>");
  document.write("<tr>");
  document.write("<td width=20 align=center>");
  document.write("<table id="+tmpid+" border=0 cellspacing=0 cellpadding=0 bgcolor=#"+strcolor+">");
  document.write("<tr><td onClick=\"javascript:select_color('"+strinput+"','"+tmpid+"');\"><img border=0 src=../../hfwadmin/rs/javascript/""+web_dir+web_dir_editor+"images/color_box.gif/" width=18 height=17></td></tr>");
  document.write("</table>");
  document.write("</td>");
  document.write("<td width=5></td>");
  document.write("<td><a href='javascript:;' onClick=\"javascript:select_color('"+strinput+"','"+tmpid+"');\">点击选取颜色</a></td>");
  document.write("</tr>");
  document.write("</table>");
}

function select_color(strinput,strtable)
{
  var tmpcolor=window.showModalDialog(web_dir+web_dir_editor+"win_select_color.asp","win_select_color","dialogWidth=300px;dialogHeight=295px;status=0");
  if (tmpcolor && tmpcolor!="")
  {
    if (tmpcolor.length==7) { tmpcolor=tmpcolor.substr(1,tmpcolor.length); }
    eval("window.document.all."+strinput+".value=tmpcolor;");
    //insert_value(strinput,tmpcolor);
  }
    if (strtable!="" && strtable != null)
    {
      eval("window.document.all."+strtable+".bgColor=tmpcolor;");
    }
}

function frm_submitonce(theform)
{
  if (document.all||document.getElementById)
  {
    for (var i=0;i<theform.length;i++)
    {
      var tempobj=theform.elements[i]
      if(tempobj.type.toLowerCase()=="submit"||tempobj.type.toLowerCase()=="reset") { tempobj.disabled=true; }
    }
  }
}

function frm_quicksubmit(eventobject) { if (event.keyCode==13 && event.ctrlKey) write_frm.submit(); }

//<input type=test name=tim value='2003-5-19' size=12 maxlength=10 readonly>
//<input type=button class=btns name=btn_select_tim value="选择" onclick="javascript:select_time(tim);">

function select_time(st_obj,st_type)
{
  var showx=event.screenX-event.offsetX-14;
  var showy=event.screenY-event.offsetY-168;
  var retval=window.showModalDialog(web_dir+web_dir_editor+"win_select_time.asp?data_tim="+st_obj.value+"&data_type="+st_type,"","dialogWidth:197px; dialogHeight:235px; dialogLeft:"+showx+"px; dialogTop:"+showy+"px; status:no; directories:yes;scrollbars:no;Resizable=no;");
  if( retval!=null ) { st_obj.value=retval; }
}

function pagecute_go(strtxt,strnum,strpage)
{
  if (strtxt=="" || strtxt==null) strtxt="page_url";
  if (strnum=="" || strnum==null) strnum="page_num";
  if (strpage=="" || strpage==null) strpage="{$page}";
  var tmpurl=form_value(strtxt);
  var tmppage=form_value(strnum);
  tmpurl=disp(tmpurl,strpage,tmppage);
  document.location.href=tmpurl;
}

//*****************选择框 开始******************
var _select_frm="frm_select";

function select_all()
{
  var frm=eval("document."+_select_frm);
  var slength=0;
  if (frm.sel_id==null) { return false; }
  var sall=frm.sel_all.checked;
  if (frm.sel_id.length)
  {
    slength=frm.sel_id.length;
    for (var i=0;i<slength;i++) { frm.sel_id[i].checked=sall; }
  }
  else { frm.sel_id.checked=sall; }
}

function select_click()
{
  var frm=eval("document."+_select_frm);
  var slength=0;
  var issel=false;
  if (frm.sel_id!=null)
  {
    if (frm.sel_id.length)
    {
      slength=frm.sel_id.length;
      for (var i=0;i<slength;i++) { if (frm.sel_id[i].checked==true) { issel=true; break; } }
    }
    else { if (frm.sel_id.checked==true) { issel=true; } }
  }
  if (issel==true)
  {
    if (confirm("执行此操作后可能无法恢复！你确定吗？")) { return true; }
    else { return false; }
  }
  else { alert("没有选择任何记录！"); return false; }
}
//*****************选择框 结束******************
