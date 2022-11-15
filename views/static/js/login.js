function phoneController() {
	if (checkPhone("phone")) {//格式正确
		document.getElementById("phone-s-m").innerHTML = "电话号码长度正确"
		document.getElementById("phone-e-m").innerHTML = ""
		//改变输入框颜色
		document.getElementById("phone").className = "input success"
	} else {
		document.getElementById("phone-s-m").innerHTML = ""
		document.getElementById("phone-e-m").innerHTML = "电话号码长度必须为11位"
		//改变输入框颜色
		document.getElementById("phone").className = "input danger"
	}
}
function pswController() {
	if (checkPSW("psw")) {//格式正确
		document.getElementById("psw-s-m").innerHTML = "密码长度正确"
		document.getElementById("psw-e-m").innerHTML = ""
		//改变输入框颜色
		document.getElementById("psw").className = "input success"
	} else {
		document.getElementById("psw-s-m").innerHTML = ""
		document.getElementById("psw-e-m").innerHTML = "密码长度必须为6到15位"
		//改变输入框颜色
		document.getElementById("psw").className = "input danger"
	}
}
function btnController() {
	if (checkPSW("psw") == true && checkPhone("phone") == true) {
		document.getElementById("btn").disabled = false
	} else {
		document.getElementById("btn").disabled = true
	}
}