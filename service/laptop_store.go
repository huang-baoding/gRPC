package service

import (
	"context"
	"errors"
	"fmt"
	"grpctest/pb"
	"log"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

//因为有不同的store，我们用接口来定义的它的功能
type LaptopStore interface {
	//save the laptop to the store
	Save(laptop *pb.Laptop) error
	//通过id查找store里是否存在此laptop
	Find(id string) (*pb.Laptop, error)
	//实现检索功能,输入一个过滤器和和一个回调函数（用于在找到时报告），返回一个错误
	Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

//定义一个结构体封装实现LapStore中定义的函数（保存到store中，和查找...）
type InMemoryLaptopStore struct {
	// 因为会有大量的并发存储，我们使用互斥锁保证并发安全
	mutex sync.RWMutex
	//key存储laptop的id
	//值是laptop对象（值里同样包含id）
	data map[string]*pb.Laptop
}

//如果之后我们想将laptop保存到数据库中，我们可以实现另一个DBLLaptopStore来做到
//type DBLLaptopStore struct{ ... }

//返回一个&InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

//将laptop保存到store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()         //先加锁
	defer store.mutex.Unlock() //别忘了加锁的同时defer解锁

	if store.data[laptop.Id] != nil { //如果store中已经有了一样的id，则不存储并返回错误
		return ErrAlreadyExists
	}

	//id不存在的情况，我们保存到data字典中
	//为了安全起见，我们对laptop对象进行深度复制
	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}
	//保存
	store.data[other.Id] = other
	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	//访问内存的时候需要先上锁
	store.mutex.Lock()
	defer store.mutex.Unlock()

	//通过id从map里找到laptop
	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	//找到后我们应该将其深刻复制给另一个对象
	return deepCopy(laptop)
}

//实现接口中的检索函数
func (store *InMemoryLaptopStore) Search(
	ctx context.Context,
	filter *pb.Filter,
	found func(laptop *pb.Laptop) error) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, laptop := range store.data { //遍历查看哪一台电脑符合过滤条件。

		//模拟超时
		//time.Sleep(time.Second)
		//log.Print("checking laptop id:",laptop.GetId())		//记录轨迹

		//检查电脑是否合格之前，我们先检查上下文的状态,错误是canceled还是deadlineExceeded
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is cancelled")
			return errors.New("context is cancelled")
		}

		if isQualified(filter, laptop) { //检查此电脑是否符合条件
			//假如此电脑合格，我们需要在调用回调函数之前对其进行深度复制
			other, err := deepCopy(laptop)
			if err != nil {
				return err
			}

			err = found(other) //使用found函数将其发送给调用方
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}
	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}
	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}
	//toBit将内存转为最小单位的函数
	//如果笔记本内存小于过滤器的，则返回false
	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}
	return true
}

func toBit(memory *pb.Memory) uint64 { //toBit将内存转为最小单位的函数
	value := memory.GetValue() //先得到内存

	switch memory.GetUint() {
	case pb.Memory_BIT: //如果只是位直接返回
		return value
	case pb.Memory_BYTE: //如果是字节，我们需返回value*8（因为一字节里面有8位）
		return value << 3 //又因为8=2^3,我们可以在这里使用位移运算左移3来避免算法
	case pb.Memory_KYLOBYTE:
		return value << 13 //1kilobyte = 1024 byte
	case pb.Memory_MEGABYTE:
		return value << 23 //
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return value
	}
}

func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data : %w", err)
	}
	return other, nil
}
