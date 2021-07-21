package main

import (
	"strconv"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucketList := []string{"bucket-pulumi-1", "bucket-pulumi-2", "bucket-pulumi-3"}

		for i, v := range bucketList {

			// Create an AWS resource (S3 Bucket)
			bucket, err := s3.NewBucket(ctx, v, nil)
			if err != nil {
				return err
			}

			// Export the name of the bucket
			ctx.Export("bucketName_"+strconv.Itoa(i), bucket.ID())
		}

		return nil
	})
}
