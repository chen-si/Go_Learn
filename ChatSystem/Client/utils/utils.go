package utils

import (
	"ChatSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message ,err error){
	//defer conn.Close()
	buf := make([]byte,8096)
	fmt.Println("读取服务器发送的消息....")
	_,err = conn.Read(buf[:4])
	if err != nil{
		//err = errors.New("conn.Read(buf[:4]) error")
		return
	}
	//将读取到的buf[:4]转换成uint32类型
	var PkgLen uint32
	PkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据PkgLen读取数据和话剧
	n,err := conn.Read(buf[:PkgLen])
	if err != nil || n != int(PkgLen){
		err = errors.New("conn.Read(buf[:PkgLen]) error")
		return
	}

	//将读取到的数据反序列化
	err = json.Unmarshal(buf[:PkgLen],&mes)
	if err != nil{
		err = errors.New("json.Unmarshal(buf[:PkgLen],&mes) error")
		return
	}

	return
}

func writePkg(conn net.Conn,data []byte)(err error){
	//先发送一个长度给对方
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkglen)
	//发送长度
	n,err :=conn.Write(buf[0:4])
	if n != 4 || err != nil{
		fmt.Println("conn.Write(buf) error:",err)
		return
	}
	//发送消息本身
	//发送消息本身
	n,err = conn.Write(data)
	if n != 4 || err != nil{
		fmt.Println("conn.Write(data) error:",err)
		return
	}
	return
}
