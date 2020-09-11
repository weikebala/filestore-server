package rpc

import (
	"context"
	"crypto/md5"
	"fmt"

	weiekbalaProto "filestore-server/service/weikebala/proto"
)

//Weikebala
type Weikebala struct {
}

func (w Weikebala) DoMD5(ctx context.Context, req *weiekbalaProto.Req, res *weiekbalaProto.Res) error {
	fmt.Println("MD5方法请求JSON:" + req.JsonStr)
	res.BackJson = "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(req.JsonStr)))
	//return &test.Res{BackJson: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(in.JsonStr)))}, nil
	return nil
}
