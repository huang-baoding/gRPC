//图像存储
package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

type ImageStore interface {
	//保存电脑图像
	Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error)
}

//将图片保存到磁盘,并将其信息存储在内存中。
type DiskImageStore struct{
	mutex sync.RWMutex
	imageFolder string
	images map[string]*ImageInfo	//key是图像id,va是图片的信息。
}

//
type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string			//在磁盘上生成图像的路径。
}

func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}

func (store *DiskImageStore) Save(
	laptopID string,
	imageType string,
	imageData bytes.Buffer,
) (string, error) {
	imageID, err := uuid.NewRandom()		//为图像生成一个ID。
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}

	imagePath := fmt.Sprintf("%s%s%s", store.imageFolder, imageID, imageType)

	file, err := os.Create(imagePath)		//	创建文件夹。
	if err != nil {
		return "", fmt.Errorf("cannot create image file : %w", err)
	}

	_,err=imageData.WriteTo(file)			//将图像存入刚刚创建的文件。
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	store.mutex.Lock()						//写入内存之前需要获取写锁
	defer store.mutex.Unlock()

	store.images[imageID.String()]=&ImageInfo{		//将图片信息保存到内存中的map中
		LaptopID: laptopID,
		Type: imageType,
		Path: imagePath,
	}
	return imageID.String(),nil
}
