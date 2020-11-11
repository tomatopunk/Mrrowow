package context

type Config interface {
	Get(string) interface{}
	Set(string,interface{}) error
}

var dic map[string]interface {}

func Get(key string) interface{} {
	return dic[key]
}

func Set(key string,value interface{}) error{
	if dic == nil {
		dic = make(map[string]interface{})
	}
	dic[key]=value
	return nil
}