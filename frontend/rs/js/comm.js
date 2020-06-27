/*///常用一些js: md5 , JL_pgData ,JL_Check
//var cm = new JSBase();
//onclick="KK_LOGIN_REG.regSubmit4();" */
/* every page cb */
let JSBase = function () {
    let o = this;
    /*------------eg:var udb =  {"n":n,"id":uid,"psw":psw};cm.LocalStorageSet("jlLogin",udb); */
    this.LocalStorageGet = function (n) {
        try {
            let a = localStorage.getItem(n);
            if (a == null || a == "") return null;
            return JSON.parse(a);
        } catch (Exception) {
        }
        return null;
    };
    this.LocalStorageSet = function (n, v) {
        try {
            localStorage.setItem(n, JSON.stringify(v));
        } catch (Exception) {
        }
    };
    this.LocalStorageRemove = function (n) {
        try {
            localStorage.removeItem(n);
        } catch (Exception) {
        }
    };
    /*jwt*/
    this.jwtCmsSet = function (db) {
        o.LocalStorageSet("cmsJwt", db);
    };
    this.jwtCmsGet = function () {
        return o.LocalStorageGet("cmsJwt")
    };
    this.jwtParse = function (db) {
        try {
            let a = decodeURIComponent(escape(window.atob(db.token.split('.')[1])))
            return JSON.parse(a);
        } catch (Exception) {
        }
        return null;
    };
};
let jsBase = new JSBase();