package define

import "regexp"

const PhoneNumberReg = `^1[3-9]\d{9}$`
const AgeReg = `^(0?[1-9]|[1-9][0-9])|^120`

// 外部调用，以减少内存占用

var PhoneNumberRegCompile = regexp.MustCompile(PhoneNumberReg)
var AgeRegCompile = regexp.MustCompile(AgeReg)
