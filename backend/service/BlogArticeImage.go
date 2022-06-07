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

	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)


func ArticleAllDeleteImageS3(c *gin.Context) (error) {
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("500 not read confg.env")
	}
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region: aws.String("ap-northeast-1"),
	})

	svc := s3.New(sess)
	bucket := os.Getenv("AWS_BUCKET")

	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})
	
	err = s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("500 bucket err")
	}
	return nil
}

// func ArticleDeleteImageS3(c *gin.Context, objectKey string) (error) {
// 	err := godotenv.Load("./config.env")
// 	if err != nil {
// 		return fmt.Errorf("not read confg.env")
// 	}
	
// }


func ArticleUploadImageS3(c *gin.Context, username string, blogID uint) (string, error) {

	err := godotenv.Load("./config.env")
	if err != nil {
		fmt.Println("not read confg.env")
		return "", err
	}
	
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("ファイルのうけとりができませんでした")
	}
	
	fileName, err := ReName(header.Filename, username, blogID)

	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region: aws.String("ap-northeast-1")},
	)
	uploader := s3manager.NewUploader(sess)

	Bucket := os.Getenv("AWS_BUCKET")
	objectKey := fileName


	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Bucket),
		Key: aws.String(objectKey),
		Body: file,
		ContentType:   aws.String("image/jpeg"),
	})
	if err != nil{
		log.Fatal("uploadに失敗しました。")
		return "", err
	}

	return objectKey, nil
}


func ReName(fileName, username string, blogID uint) (string, error) {
	FileNameArray := strings.Split(fileName, ".")
	if len(FileNameArray) >= 3 {
		return "", fmt.Errorf("不適切なファイル")
	}
	Filetype := FileNameArray[len(FileNameArray)-1]
	FileFirstName := FileNameArray[len(FileNameArray)-2]

	FileName := fmt.Sprintf("%s%d%s.%s", FileFirstName, blogID, username, Filetype)
	return FileName, nil

}