/*
 *  自定义工具函数
 **/

Array.prototype.remove = function (val) {
    var index = this.indexOf(val);
    if (index > -1) {
        this.splice(index, 1);
    }
}
Object.defineProperty(Array.prototype, "remove", { enumerable: false });

//region 比较两个数组是否相等 array1.equals(array2)
// Warn if overriding existing method
if (Array.prototype.equals)
    console.warn("Overriding existing Array.prototype.equals. Possible causes: New API defines the method, there's a framework conflict or you've got double inclusions in your code.");
// attach the .equals method to Array's prototype to call it on any array
Array.prototype.equals = function (array) {
    // if the other array is a falsy value, return
    if (!array)
        return false;
    // compare lengths - can save a lot of time
    if (this.length != array.length)
        return false;
    for (var i = 0, l = this.length; i < l; i++) {
        // Check if we have nested arrays
        if (this[i] instanceof Array && array[i] instanceof Array) {
            // recurse into the nested arrays
            if (!this[i].equals(array[i]))
                return false;
        }
        else if (this[i] != array[i]) {
            // Warning - two different object instances will never be equal: {x:20} != {x:20}
            return false;
        }
    }
    return true;
}
// Hide method from for-in loops
Object.defineProperty(Array.prototype, "equals", { enumerable: false });

//endregion

//获取URL传参
function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
}

//获取所有URL参数
function getUrlParamAll() {
    let params = {};

    let url = document.location.toString();
    let arrObj = url.split("?");
    if (arrObj.length > 1) {
        let arrPara = arrObj[1].split("&");
        let arr;
        for (let i = 0; i < arrPara.length; i++) {
            arr = arrPara[i].split("=");
            if (arr != null && arr.length > 1) {
                params[arr[0]] = arr[1];
            }
        }
    }
    return params;
}

function getCurrentTime() {
    var now = new Date();

    var year = now.getFullYear();       //年
    var month = now.getMonth() + 1;     //月
    var day = now.getDate();            //日

    var hh = now.getHours();            //时
    var mm = now.getMinutes();          //分
    var ss = now.getSeconds();          //秒

    var clock = year + "-";

    if (month < 10)
        clock += "0";

    clock += month + "-";

    if (day < 10)
        clock += "0";

    clock += day + " ";

    if (hh < 10)
        clock += "0";

    clock += hh + ":";
    if (mm < 10) clock += '0';
    clock += mm + ":";

    if (ss < 10) clock += '0';
    clock += ss;
    return (clock);
}

//克隆新的Object对象
function cloneObject(obj) {
    let temp = null;
    if (obj instanceof Array) {
        temp = obj.concat();
    } else if (obj instanceof Function) {
        //函数是共享的是无所谓的，js也没有什么办法可以在定义后再修改函数内容
        temp = obj;
    } else {
        temp = new Object();
        for (let item in obj) {
            let val = obj[item];
            temp[item] = typeof val == 'object' ? clone(val) : val; //这里也没有判断是否为函数，因为对于函数，我们将它和一般值一样处理
        }
    }
    return temp;
}

function getFormParams(formid, temp) {
    if (!temp)
        var temp = {};
    $.each($('#' + formid).serializeArray(), function (i, field) {
        temp[field.name] = field.value;
    });

    return temp;
}

//处理Get请求.变_的问题，后端处理替换__为.
function getFormParamsFormatInput(formid) {
    var temp = getFormParams(formid);
    var params = {};
    $.each(temp, function (key, item) {
        if (key.indexOf(".") != -1)
            key = key.replace('.', '__');

        params[key] = item;
    });
    return params;
}

function uuid() {
    var s = [];
    var hexDigits = "0123456789abcdef";
    for (var i = 0; i < 36; i++) {
        s[i] = hexDigits.substr(Math.floor(Math.random() * 0x10), 1);
    }
    s[14] = "4"; // bits 12-15 of the time_hi_and_version field to 0010
    s[19] = hexDigits.substr((s[19] & 0x3) | 0x8, 1); // bits 6-7 of the clock_seq_hi_and_reserved to 01
    s[8] = s[13] = s[18] = s[23] = "-";

    var uuid = s.join("");
    return uuid;
}

function getRandom(num) {
    var random = Math.floor((Math.random() + Math.floor(Math.random() * 9 + 1)) * Math.pow(10, num - 1));
    return random;
}

function formatUrlParams(params) {
    var urlParams = '';
    $.each(params, function (key, item) {
        if (urlParams != '')
            urlParams += '&';
        urlParams += key + '=' + item;
    });
    return urlParams;
}

function formatParamsToUrl(url, params) {
    var urlParams = '';
    if (url.indexOf("?") != -1) {
        urlParams = url.split("?")[1];
    }
    $.each(params, function (key, item) {
        if (urlParams != '')
            urlParams += '&';
        urlParams += key + '=' + item;
    });

    if (url.indexOf("?") != -1)
        return url.split("?")[0] + "?" + urlParams;
    return url + "?" + urlParams;
}

//动态执行函数 fn 函数名 字符串需要eval('function') ,args 参数集 数组  doCallback(eval('function'),{});
function doCallback(fn, args) {
    fn.apply(this, args);
}


function dateFormat(fmt, date) {
    let ret;
    const opt = {
        "Y+": date.getFullYear().toString(),        // 年
        "m+": (date.getMonth() + 1).toString(),     // 月
        "d+": date.getDate().toString(),            // 日
        "H+": date.getHours().toString(),           // 时
        "M+": date.getMinutes().toString(),         // 分
        "S+": date.getSeconds().toString()          // 秒
        // 有其他格式化字符需求可以继续添加，必须转化成字符串
    };
    for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
            fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
        };
    };
    return fmt;
}