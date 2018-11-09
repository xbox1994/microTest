package service

import (
	"microTest/dao"
	"microTest/request"
)

func Do(request *request.HelloRequest) (string, error) {
	dao.NewHelloDAO().GetInfo()
	return request.Name, nil
}
