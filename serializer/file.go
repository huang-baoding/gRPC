package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

//proto.Message是官方包，指在proto里定义的结构体都可以作为参数
func WriteProtobufToBinaryFile(message proto.Message, fileName string) error {
	//将proto.Message转化字节切片类型的数据
	data, err := proto.Marshal(message)//message必须是指针类型
	if err != nil {
		return fmt.Errorf("can not marshal proto message to binary:%w", err)
	}
	//将data写入文件
	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("can not write to file:%w", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(message proto.Message, fileName string) error {
	data, err := ioutil.ReadFile(fileName)	//从文件中读取消息（读到的格式为字符切片）
	if err != nil {
		return fmt.Errorf("can not read the binary file:%w", err)
	}
	//Unmarshal将读到的data数据反格式化后写入message
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("can not unmarshal binary to proto message:%w", err)
	}
	return nil
}

func WriteProtobufToJSONFile(message proto.Message, fileName string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}

	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}
	return nil
}
