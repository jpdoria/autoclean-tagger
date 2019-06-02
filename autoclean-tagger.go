package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func applyTag(b, k string) {
	// Get the bucket location first.
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	svc := s3.New(sess)
	in0 := &s3.GetBucketLocationInput{
		Bucket: aws.String(b),
	}
	res0, err0 := svc.GetBucketLocation(in0)
	var bucketRegion string

	if err0 != nil {
		log.Errorln(err0)
		return
	}

	if res0.LocationConstraint != nil {
		bucketRegion = *res0.LocationConstraint
	} else {
		bucketRegion = "us-east-1"
	}

	// Apply tag to object.
	sess = session.Must(session.NewSession(&aws.Config{Region: aws.String(bucketRegion)}))
	svc = s3.New(sess)
	in1 := &s3.PutObjectTaggingInput{
		Bucket: aws.String(b),
		Key:    aws.String(k),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String("autoclean"),
					Value: aws.String("true"),
				},
			},
		},
	}
	res1, err1 := svc.PutObjectTagging(in1)
	var msg string

	if err1 != nil {
		log.Errorln(err1)
		return
	}

	if res1.VersionId != nil {
		msg = fmt.Sprintf("Applied \"autoclean=true\" tag to %q in %q bucket. Version ID is %q.\n", k, b, *res1.VersionId)
	} else {
		msg = fmt.Sprintf("Applied \"autoclean=true\" tag to %q in %q bucket.\n", k, b)
	}

	log.Println(msg)
}

func handler(s3Event events.S3Event) {
	bucket := s3Event.Records[0].S3.Bucket.Name
	key := s3Event.Records[0].S3.Object.Key
	applyTag(bucket, key)
}

func main() {
	lambda.Start(handler)
}
