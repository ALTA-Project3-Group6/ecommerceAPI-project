package helper

import (
	"ecommerceapi/config"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
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
	ext := filepath.Ext(file.Filename)

	cnv := strconv.Itoa(productId)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("ecommapi"),
		Key:    aws.String("files/products/" + cnv + "/product_photo_" + fmt.Sprint(productId) + ext),
		Body:   src,
	})
	if err != nil {
		return "", errors.New("problem with upload post photo")
	}
	path := ObjectURL + "files/products/" + cnv + "/product_photo_" + fmt.Sprint(productId) + ext
	return path, nil
}

func UploadProfilePhotoS3(file multipart.FileHeader, userID uint) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	cnv := strconv.Itoa(int(userID))
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("ecommapi"),
		Key:    aws.String("files/user/" + cnv + "/profile-photo" + ext),
		Body:   src,
	})
	if err != nil {
		return "", errors.New("problem with upload profile photo")
	}
	path := ObjectURL + "files/user/" + cnv + "/profile-photo" + ext
	return path, nil
}
