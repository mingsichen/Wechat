package main

import (
    "encoding/binary"
    "net"
    "fmt"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:3563")
    if err != nil {
        panic(err)
    }

    // Hello 消息（JSON 格式）
    // 对应游戏服务器 Hello 消息结构体
    data := []byte(`{
        "UserInfo": {
            "Cmd": "get_user_id_info",
	    "UserId":668
        }
    }`)
fmt.Println(data)
    // len + data
    m := make([]byte, 2+len(data))
fmt.Println(m)
    // 默认使用大端序
    binary.BigEndian.PutUint16(m, uint16(len(data)))
fmt.Println(m)
    copy(m[2:], data)
	fmt.Println(m)
    // 发送消息
    conn.Write(m)
}
