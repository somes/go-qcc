package qcc

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"strings"
)

/*
var i = (0, a.default)(t, e.data),
l = (0, r.default)(t, e.data, (0, s.default)());
*/
// s.default)() = "00000000000000000000000000000000"
func Calculate(t, edata, tid string) (key, value string, err error) {
	i, err := l1aDefault([]string{t, edata})
	if err != nil {
		return
	}

	l := rDefault([]string{t, edata, tid})

	return i, l, nil
}

/*
	function() {
		var e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {}
			, t = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : ""
			, n = (arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "/").toLowerCase()
			, i = JSON.stringify(e).toLowerCase();
		return (0,
		o.default)(n + "pathString" + i + t, (0,
		a.default)(n))
	}
*/
// t, e.data, (0, s.default)()
func rDefault(tAndedataAndfuncret []string) string {
	var e, t, n, i string
	if len(tAndedataAndfuncret) > 1 && tAndedataAndfuncret[1] != "" {
		e = tAndedataAndfuncret[1]
	} else {
		e = "{}"
	}
	if len(tAndedataAndfuncret) > 2 && tAndedataAndfuncret[2] != "" {
		t = tAndedataAndfuncret[2]
	} else {
		t = ""
	}
	if len(tAndedataAndfuncret) > 0 && tAndedataAndfuncret[0] != "" {
		n = tAndedataAndfuncret[0]
	} else {
		n = "/"
	}
	n = strings.ToLower(n)
	i = strings.ToLower(e)
	return oDefault([]string{n + "pathString" + i + t, l2aDefault([]string{n})})
}

/*
	function() {
		var e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {}
			, t = (arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "/").toLowerCase()
			, n = JSON.stringify(e).toLowerCase();
		return (0,
		o.default)(t + n, (0,
		a.default)(t)).toLowerCase().substr(8, 20)
	}
*/
func l1aDefault(tAndedata []string) (string, error) {
	var e, t, n string
	if len(tAndedata) > 1 && tAndedata[1] != "" {
		e = tAndedata[1]
	} else {
		e = "{}"
	}
	if len(tAndedata) > 0 && tAndedata[0] != "" {
		t = tAndedata[0]
	} else {
		t = "/"
	}
	t = strings.ToLower(t)
	n = strings.ToLower(e)

	return strings.ToLower(oDefault([]string{t + n, l2aDefault([]string{t})}))[8:28], nil
}

/*
	function(e, t) {
		return (0,
		o.default)(e, t).toString()
	}

// o.default

	function(e, n) {
		return new h.HMAC.init(t,n).finalize(e)
	}
*/
func oDefault(eAndn []string) string {
	var e, n = eAndn[0], eAndn[1]
	hmac := hmac.New(sha512.New, []byte(n))
	hmac.Write([]byte(e))

	return hex.EncodeToString(hmac.Sum([]byte(nil)))
}

/*
	function() {
		for (var e = (arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "/").toLowerCase(), t = e + e, n = "", i = 0; i < t.length; ++i) {
			var a = t[i].charCodeAt() % o.default.n;
			n += o.default.codes[a]
		}
		return n
	}
*/
func l2aDefault(t []string) (n string) {
	var e string
	if len(t) > 0 && t[0] != "" {
		e = t[0]
	} else {
		e = "/"
	}
	e = strings.ToLower(e)
	t2 := e + e

	codes := map[int]string{0: "W", 1: "l", 2: "k", 3: "B", 4: "Q", 5: "g", 6: "f", 7: "i", 8: "i", 9: "r", 10: "v", 11: "6", 12: "A", 13: "K", 14: "N", 15: "k", 16: "4", 17: "L", 18: "1", 19: "8"}

	for i := 0; i < len(t2); i++ {
		a := int(rune(t2[i])) % 20
		n += codes[a]
	}

	return
}
