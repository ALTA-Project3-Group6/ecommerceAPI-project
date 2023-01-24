package helper

import (
	"ecommerceapi/config"
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var ObjectURL string = "https://ecommapi.s3.ap-southeast-1.amazonaws.com/"

func UploadProductPhotoS3(file multipart.FileHeader, productId int) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// ext := filepath.Ext(file.Filename)

	cnv := strconv.Itoa(productId)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("socmedapibucket"),
		Key:    aws.String("files/products/" + cnv + "/" + file.Filename),
		Body:   src,
	})
	if err != nil {
		return "", errors.New("problem with upload post photo")
	}
	path := ObjectURL + "files/products/" + cnv + "/" + file.Filename
	return path, nil
}
