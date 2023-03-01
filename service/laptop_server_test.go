package service_test

import (
	"context"
	"grpctest/pb"
	"grpctest/sample"
	"grpctest/service"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel() //让它并行运行

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID) //测试保存之前先将laptopDuplicateID存入store
	require.Nil(t, err)

	//使用表驱动测试
	//先声明所有的测试用例
	testCases := []struct {
		name   string              //每个测试用例都有一个名称。
		laptop *pb.Laptop          //需要测试的laptop
		store  service.LaptopStore //接口，有save和find函数
		code   codes.Code          //测试预期的状态码
	}{
		{
			name:   "success_with_id", //第一种情况是使用客户端生成的Id
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID, //测试客户端传过来的电脑没有ID。
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id", //测试客户端发过来的电脑ID是无效的。
			laptop: laptopInvalidID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id", //测试客户端发过来的电脑ID是无效的。
			laptop: laptopDuplicateID,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases { //遍历进行测试
		//先将测试用例保存到局部，这对并发安全来说非常重要，因为我们要创建多个并行子测试
		tc := testCases[i]

		//运行测试
		t.Run(tc.name, func(t *testing.T) { //子测试
			t.Parallel()		//使其与其他测试并行运行

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			////////////////////////////////////////
			server := service.NewLaptopServer(tc.store,nil,nil) //////////////////
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)     //确保没有错误
				require.NotNil(t, res)      //响应不应该为空
				require.NotEmpty(t, res.Id) //响应的id不应该为空
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else { //不为OK的情况/////////////////////////
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err) //////////////////////////
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
