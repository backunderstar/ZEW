package upload

import (
	"fmt"
	"strconv"

	"github.com/backunderstar/zew/global"
	"github.com/backunderstar/zew/models"
	"github.com/backunderstar/zew/utils"
	"github.com/backunderstar/zew/utils/res"
	"github.com/gin-gonic/gin"
)

type UploadResponse struct {
	Filename  string `json:"filename"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

func (u *UploadApi) UploadImage(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		global.Log.Error("不存在的图片", err)
		res.FailWithMsg(err.Error(), c)
		return
	}

	var resFileList []res.FileUpload
	files := form.File["files"]
	global.Log.Infof("本次上传文件数量为%v", len(files))
	for _, file := range files {
		suffix, suffixErr := utils.CheckFileSuffixIsRight(file)
		if suffixErr != nil {
			resFileList = append(resFileList, res.FileUpload{
				FileName:  file.Filename,
				Url:       "",
				IsSuccess: false,
				ErrMsg:    suffixErr.Error(),
			})
			continue
		}
		if !utils.CheckFileSizeIsRight(float64(file.Size)) {
			// 超出文件大小限制
			resFileList = append(resFileList, res.FileUpload{
				FileName:  file.Filename,
				Url:       "",
				IsSuccess: false,
				ErrMsg:    fmt.Sprintf("当前文件大小为%vM,已超出%vM限制", strconv.FormatFloat(float64(file.Size)/float64(1024*1024), 'f', 2, 64), strconv.FormatFloat(global.Config.Upload.Size, 'f', 2, 64)),
			})
		} else {
			filePath := utils.GenerationFilePath(file.Filename)
			err = c.SaveUploadedFile(file, filePath)
			if err != nil {
				resFileList = append(resFileList, res.FileUpload{
					FileName:  file.Filename,
					Url:       "",
					IsSuccess: false,
					ErrMsg:    err.Error(),
				})
			} else {
				imgRes := models.FileHashToDb(file, filePath, suffix)
				resFileList = append(resFileList, imgRes)
			}
		}
	}
	res.OkWithData(resFileList, c)

}
