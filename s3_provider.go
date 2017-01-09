package fisherking

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	s3lib "github.com/aws/aws-sdk-go/service/s3"
)

type s3 struct {
	context.Context
}

// func (s3) GetWithContext(contect context.Context, path string) FileGetter {
// 	//Listern to cancel channel.
// 	return gcs{}.Get
// }
func (s s3) Get(path string) (io.Reader, error) {
	region := `us-west-2`
	ri := s.Value(`region`)
	if rg, ok := ri.(string); ok && len(rg) > 0 {
		region = rg
	}

	sess, err := session.NewSession(&aws.Config{Region: aws.String(region), Credentials: credentials.NewSharedCredentials("", "")})
	if err != nil {
		fmt.Println("failed to create s3 session,", err)
		return nil, err
	}

	svc := s3lib.New(sess)
	bucket, fileName := parseS3Bucket(path)
	params := &s3lib.GetObjectInput{
		Bucket: bucket,
		Key:    fileName, // Required
		//IfMatch:                    aws.String("IfMatch"),
		//IfModifiedSince:            aws.Time(time.Now()),
		//IfNoneMatch:                aws.String("IfNoneMatch"),
		//IfUnmodifiedSince:          aws.Time(time.Now()),
		//PartNumber:                 aws.Int64(1),
		//Range:                      aws.String("Range"),
		//RequestPayer:               aws.String("RequestPayer"),
		//ResponseCacheControl:       aws.String("ResponseCacheControl"),
		//ResponseContentDisposition: aws.String("ResponseContentDisposition"),
		//ResponseContentEncoding:    aws.String("ResponseContentEncoding"),
		//ResponseContentLanguage:    aws.String("ResponseContentLanguage"),
		//ResponseContentType:        aws.String("ResponseContentType"),
		//ResponseExpires:            aws.Time(time.Now()),
		//SSECustomerAlgorithm:       aws.String("SSECustomerAlgorithm"),
		//SSECustomerKey:             aws.String("SSECustomerKey"),
		//SSECustomerKeyMD5:          aws.String("SSECustomerKeyMD5"),
		//VersionId:                  aws.String("ObjectVersionId"),
	}
	resp, err := svc.GetObject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil, err
	}

	// Pretty-print the response data.
	return resp.Body, nil
}

func parseS3Bucket(path string) (bucket, file *string) {
	be := strings.Index(path[4:], linPathSeperator)
	return aws.String(path[4 : be+4]), aws.String(path[be+4+1:])
}

func (s s3) Put(destination string, data io.Reader) error {
	return nil
}
