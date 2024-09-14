package generator

import "reflect"

// StructCopy 参数: 第一个为结构体，第二个为指针（结构体地址）
func StructCopy(source, target interface{}) {
	sValue := reflect.ValueOf(source)
	tValue := reflect.ValueOf(target).Elem()

	for i := 0; i < sValue.NumField(); i++ {
		sField := sValue.Field(i)
		tField := tValue.FieldByName(sValue.Type().Field(i).Name)

		if tField.IsValid() && tField.CanSet() {
			tField.Set(sField)
		}
	}
}
