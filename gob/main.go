package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	//创建两个实例map
	map1 := make(map[string]interface{}, 10)
	map2 := make(map[string]interface{}, 10)
	//存入实例数据
	map1["one"] = 1
	map1["two"] = "哈哈"
	//创建编码器将实例数据编码
	buf := new(bytes.Buffer) //创建一个buffer 指针 io.write类型
	enc := gob.NewEncoder(buf)
	enc.Encode(map1)
	//创建解码器
	dec := gob.NewDecoder(buf)
	dec.Decode(&map2)
	fmt.Printf("value:%v,type:%T", map2, map2["one"])

}
