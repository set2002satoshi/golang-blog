package service

import (

	"github.com/gin-gonic/gin"

	// imgupload "github.com/olahol/go-imageupload"
	// バケット一覧表示
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/s3"
	// "fmt"
	// "os"

	// アップロードandダウンロード
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	// "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)


func ArticleImageUploadS3(c *gin.Context, username string, blogID uint) (string, error) {

	err := godotenv.Load("./config.env")
	if err != nil {
		fmt.Println("not read confg.env")
		return "", err
	}

	
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal("ファイルのうけとりができませんでした")
		return "", err
	}
	filename := strings.Split(header.Filename, ".")
	if len(filename) >= 3 {
		log.Fatal("不適切なファイル")
		return "", err
	}
	filetype := filename[len(filename)-1]
	FileFirstName := filename[len(filename)-2]
	if string(filetype) != "jpeg" {
		log.Fatal("無効なファイル")
		return "", err
	}

	// fileName := FileFirstName+ "/" + s + username +  "." + filetype
	fileName := fmt.Sprintf("%s/%d%s.%s", FileFirstName, blogID, username, filetype)

	saveImage, err := os.Create(header.Filename)
	if err != nil {
		log.Fatal("ファイル作成に失敗")
		return "", err
	}
	defer saveImage.Close()
	_, err = io.Copy(saveImage, file)
	if err != nil {
		log.Fatal("ファイルコピーができませんでした。")
		return "", err
	}

	defer os.Remove(header.Filename)

	
	creds := credentials.NewStaticCredentials(os.Getenv("aws_access_key_id"), os.Getenv("aws_secret_access_key"), "")

	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region: aws.String("ap-northeast-1")},
	)
	uploader := s3manager.NewUploader(sess)

	Bucket := "blog0601"
	objectKey := fileName

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Bucket),
		Key: aws.String(objectKey),
		Body: saveImage,
	})
	if err != nil{
		log.Fatal("uploadに失敗しました。")
		return "", err
	}

	return objectKey, nil
}

func ArticleImageDownload(c *gin.Context, URL string) {
	
}




