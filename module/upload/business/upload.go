package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"food-delivery/common"
	"food-delivery/component/uploadprovider"
	"image/jpeg"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image)
}

type UploadBusiness struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBusiness(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *UploadBusiness {
	return &UploadBusiness{provider: provider, imgStore: imgStore}
}

func (biz *UploadBusiness) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	// Get width, height
	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		//return nil, uploadmodel.ErrFileIsNotImage(err)
		fmt.Println("ErrFileIsNotImage: ", err)
		return nil, err
	}
	fmt.Println("width, height: ", w, h)

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "abc.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // "1257856852039812612.jpg"

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, err
		//return nil, uploadmodel.ErrCanNotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.CloudName = "s3" // should beset in provider
	img.Extension = fileExt

	fmt.Println(img)

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, err := jpeg.DecodeConfig(reader)

	fmt.Println("image dimensions: ", img)

	if err != nil {
		fmt.Println("error here !!!")
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
