module github.com/aws/aws-sdk-go-v2/service/internal/benchmark

go 1.15

require (
	github.com/aws/aws-sdk-go v1.34.33
	github.com/aws/aws-sdk-go-v2 v0.26.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v0.26.0
	github.com/aws/aws-sdk-go-v2/service/lexruntimeservice v0.26.0
	github.com/awslabs/smithy-go v0.1.1
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v0.1.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/aws/aws-sdk-go-v2/service/lexruntimeservice => ../../../service/lexruntimeservice/
