package main

import (
	"fmt"
	"os"

	pb "excel2proto/xlsp" // 路径与 go.mod 同名

	"google.golang.org/protobuf/proto"
)

func main() {
	raw, err := os.ReadFile("../output/activity_weight.bytes")
	if err != nil {
		panic(err)
	}

	var list pb.ActivityWeightList
	if err := proto.Unmarshal(raw, &list); err != nil {
		panic(err)
	}

	m := make(map[uint32]uint32)
	for _, v := range list.Items {
		m[v.Id] = v.Weight
	}

	fmt.Println("反序列化成功，内容：", m)
}
