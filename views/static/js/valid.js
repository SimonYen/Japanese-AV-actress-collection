function checkPhone(id) {
	//检查电话号码长度，如果不是11位就返回false
	phone = document.getElementById(id).value
	if (phone.length != 11) {
		return false
	}
	return true
}

function checkPSW(id) {
	//密码长度需要大于5位同时小于16位
	psw = document.getElementById(id).value
	if (psw.length > 5 && psw.length < 16) {
		return true
	}
	return false
}