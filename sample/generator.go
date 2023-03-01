package sample

import (
	"grpctest/pb"

	"github.com/golang/protobuf/ptypes"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

//返回一个CPU样例
func NewCPU() *pb.CPU {
	brand := randomCPUBrande()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:        brand,
		Name:         name,
		NumberCores:  uint32(numberCores),
		NumberThread: uint32(numberThreads),
		MinGhz:       minGhz,
		MaxGhz:       maxGhz,
	}

	return cpu
}

//返回一个GPU样例
func NewGPU() *pb.GPU {
	brand := randomGPUBrande()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Uint:  pb.Memory_GIGABYTE,
		},
	}

	return gpu
}

//返回一个Ram样例
func NewRam() *pb.Memory {
	memGB := randomInt(4, 64)

	ram := &pb.Memory{
		Value: uint64(memGB),
		Uint:  pb.Memory_GIGABYTE,
	}

	return ram
}

//返回一个SSD样例
func NewSSD() *pb.Storage {
	memGB := randomInt(128, 1024)

	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Uint:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

//返回一个HDD样例
func NewHDD() *pb.Storage {
	memTB := randomInt(1, 6)

	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(memTB),
			Uint:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

//返回一个屏幕样例
func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

//返回一个电脑样例
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRam(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdateAt:    ptypes.TimestampNow(),
	}

	return laptop
}

//评级
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
