module github.com/aws/aws-sdk-go-v2/credentials

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.24.1-0.20200924201622-61f73b1f46cb
	github.com/aws/aws-sdk-go-v2/ec2imds v0.0.0-20200924201622-61f73b1f46cb
	github.com/aws/aws-sdk-go-v2/service/sts v0.0.0-20200924201622-61f73b1f46cb
	github.com/awslabs/smithy-go v0.0.0-20200924163652-fc0366622e14
)

replace (
	github.com/aws/aws-sdk-go-v2 => ../
	github.com/aws/aws-sdk-go-v2/ec2imds => ../ec2imds/
	github.com/aws/aws-sdk-go-v2/service/sts => ../service/sts/
)
