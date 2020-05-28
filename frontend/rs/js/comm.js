/*///常用一些js: md5 , JL_pgData ,JL_Check
//var KK_LOGIN_REG = new JSBase();
//onclick="KK_LOGIN_REG.regSubmit4();" */
/* every page cb */
let JSBase = function(){
    /*------------localStorage:eg:var udb =  {"n":n,"id":uid,"psw":psw};		db.DbJsonSet("jlLogin",udb); */
    this.LocalStorageGet=function(n){
        /*localStorage.lastname="Smith";document.write("Last name: " + localStorage.lastname);	*/
        try{
            let a= localStorage.getItem(n);
            if(a==null || a=="")return null;
            return  JSON.parse(a);
        }catch(Exception){};
        return null ;
    };
    this.LocalStorageSet=function(n,v){
        try{ localStorage.setItem(n,JSON.stringify(v));	}catch(Exception){};
    };
    this.LocalStorageRemove=function(n){
        try{ localStorage.removeItem(n );	}catch(Exception){};
    };

};