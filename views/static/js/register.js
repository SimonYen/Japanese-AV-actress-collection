function checkSame(id1, id2) {
	psw1 = document.getElementById(id1).value
	psw2 = document.getElementById(id2).value
	return psw1 === psw2
}
function phoneController() {
	if (checkPhone("phone-re")) {//格式正确
		document.getElementById("phone-s-m-re").innerHTML = "电话号码长度正确"
		document.getElementById("phone-e-m-re").innerHTML = ""
		//改变输入框颜色
		document.getElementById("phone-re").className = "input success"
	} else {
		document.getElementById("phone-s-m-re").innerHTML = ""
		document.getElementById("phone-e-m-re").innerHTML = "电话号码长度必须为11位"
		//改变输入框颜色
		document.getElementById("phone-re").className = "input danger"
	}
}
function pswController() {
	if (checkPSW("psw-re")) {//格式正确
		document.getElementById("psw-s-m-re").innerHTML = "密码长度正确"
		document.getElementById("psw-e-m-re").innerHTML = ""
		//改变输入框颜色
		document.getElementById("psw-re").className = "input success"
	} else {
		document.getElementById("psw-s-m-re").innerHTML = ""
		document.getElementById("psw-e-m-re").innerHTML = "密码长度必须为6到15位"
		//改变输入框颜色
		document.getElementById("psw-re").className = "input danger"
	}
}
function pswConfirmController() {
	if (checkSame("psw-re", "psw-re-2")) {//格式正确
		document.getElementById("psw-s-m-re-2").innerHTML = "密码一致"
		document.getElementById("psw-e-m-re-2").innerHTML = ""
		//改变输入框颜色
		document.getElementById("psw-re-2").className = "input success"
	} else {
		document.getElementById("psw-s-m-re-2").innerHTML = ""
		document.getElementById("psw-e-m-re-2").innerHTML = "密码不一致"
		//改变输入框颜色
		document.getElementById("psw-re-2").className = "input danger"
	}
}
function btnController() {
	if (checkPSW("psw-re") == true && checkPhone("phone-re") == true && checkSame("psw-re", "psw-re-2")) {
		document.getElementById("btn-re").disabled = false
	} else {
		document.getElementById("btn-re").disabled = true
	}
}