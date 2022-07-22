package error

type PropertyEmptyError struct {
	PropertyName string
}

func (e *PropertyEmptyError) Error() string {
	return e.PropertyName + " 属性为空"
}
