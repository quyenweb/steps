module github.com/stackpulse/public-steps/aws/redshift/base

go 1.16

require (
	cloud.google.com/go/storage v1.14.0 // indirect
	github.com/Jeffail/gabs/v2 v2.6.0
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/aws/aws-sdk-go-v2 v1.2.0
	github.com/aws/aws-sdk-go-v2/config v1.1.1
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.1.1
	github.com/olekukonko/tablewriter v0.0.5
	github.com/stackpulse/public-steps/common v0.0.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace github.com/stackpulse/public-steps/common v0.0.0 => ./../../../common