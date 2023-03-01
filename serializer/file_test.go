package serializer_test //测试包应该以_test结尾

import (
	"grpctest/pb"
	"grpctest/sample"
	"grpctest/serializer"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jasonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	//将laptop1写入二进制文件
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err) //测试专用，期望err为空

	err = serializer.WriteProtobufToJSONFile(laptop1, jasonFile)
	require.NoError(t, err)

	//将刚保存的laptop.bin文件夹中的内容读出，和保存之前的值进行比较
	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(laptop2, binaryFile)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

}
