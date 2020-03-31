package check_phone

import (
	"regexp"
)

// 增加190 号段 166号段
var mobilePattern = regexp.MustCompile(`^((\+86)|(86))?(1(([35][0-9])|[8][0-9]|[6][6]|[7][01356789]|[4][579]|[9][09]))\d{8}$`)

var nicknamePattern = regexp.MustCompile(`^[a-z0-9_-]{3,16}$`)

var ipPattern = regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)

var telPattern = regexp.MustCompile(`^(0\d{2,3}(\-)?)?\d{7,8}$`)

var passwordPattern = regexp.MustCompile(`^[a-z0-9_-]{6,18}$`)

// var emailPattern = regexp.MustCompile(`^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`)
var emailPattern = regexp.MustCompile(`^[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[a-zA-Z0-9](?:[\w-]*[\w])?$`)

var urlPattern = regexp.MustCompile(`^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?`)

func IsValuePhone(p string) bool {
	return mobilePattern.MatchString(p)
}
