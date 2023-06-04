package ginupload

import (
	"fmt"
	"food-delivery/common"
	"food-delivery/component/appcontext"
	uploadbusiness "food-delivery/module/upload/business"
	"github.com/gin-gonic/gin"
)

func Upload(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img") // Phân biệt loại image

		fmt.Println("folder: ", folder)

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // We can close here

		// Read file
		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		business := uploadbusiness.NewUploadBusiness(appCtx.UploadProvider(), nil)
		img, err := business.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
