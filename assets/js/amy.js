
function hasClass(elem, className) {
	return (' ' + elem.className + ' ').indexOf(' ' + className + ' ') > -1;
}

function addEventListener(elem, on, cb, useCapture) {
	if(!elem) {
		return false;
	}
	if(typeof elem.addEventListener === "function") {
		console.log("DEBUG: elem.addEventListener");
		if(useCapture === "undefined") {
			useCapture = false;
		}
		elem.addEventListener(on, cb, useCapture);
	} else if(typeof elem.attachEvent === "function") {
		console.log("DEBUG: elem.attachEvent");
		elem.attachEvent("on" + on, cb);
	} else {
		console.log("DEBUG: elem[]");
		elem["on" + on] = cb;
	}
	return true;
}

function removeEventListener(elem, on, cb, useCapture) {
	if(!elem) {
		return false;
	}
	if(typeof elem.removeEventListener === "function") {
		if(useCapture === "undefined") {
			useCapture = false;
		}
		elem.removeEventListener(on, cb, useCapture);
	} else if(typeof elem.detachEvent === "function") {
		elem.detachEvent("on" + on, cb);
	} else {
		elem["on" + on] = null;
	}
	return true;
}

function replaceState(url) {
	if(typeof history.replaceState === "function") {
		history.replaceState(null, null, url);
		return true;
	}
	return false;
}

function onDOMReadyHead() {
	var content = document.getElementById("content");
	if(content) {
		content.style.display = "block";
	}
	var lightSwitch = document.getElementById("lights-off");
	if(lightSwitch) {
		var lights = cookieGet("lights");
		switch(lights) {
		case "off":
			lightSwitch.checked = false;
			console.log("DEBUG: found lights on");
			break;
		case "on":
			lightSwitch.checked = true;
			console.log("DEBUG: found lights off");
			break;
		default:
			console.log("DEBUG: found nothing");
			break;
		}
		console.log("DEBUG: lights.onchange was: "+lights.onchange);
		addEventListener(lights, "change", onLightsChanged);
		console.log("DEBUG: lights.onchange is now: "+lights.onchange);
		console.log("DEBUG: document.getElementById('lights-off').onchange is now: "+document.getElementById("lights-off").onchange);
	}
}

function onDOMReadyTail() {
	if(location.hash.length >= 2) {
		var elem = document.getElementById(location.hash.slice(1))
		if(elem) {
			if(hasClass(elem, "fragment-block")) {
				elem.style.display = "block";
			} else if(hasClass(elem, "fragment-inline")) {
				elem.style.display = "inline";
			} else if(hasClass(elem, "fragment-inline-block")) {
				elem.style.display = "inline-block";
			}
			if(!replaceState(location.href.split('#')[0])) {
				location.hash = '';
			}
		}
	} else if(location.href.slice(-1) == '#') {
		replaceState(location.href.slice(0, -1));
	}
}

function onLightsChanged(event) {
	var lightSwitch = document.getElementById("lights-off");
	if(lightSwitch) {
		if(lightSwitch.checked) {
			cookieSet("lights", "on");
			console.log("DEBUG: set lights to on");
		} else {
			cookieSet("lights", "off");
			console.log("DEBUG: set lights to off");
		}
	}
}
