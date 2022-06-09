package method

import "github.com/zii/web/service"

func Hello(md *service.Meta) (interface{}, error) {
	id := md.Get("id").Int()
	if id == 1 {

	}
	return id, nil
}
