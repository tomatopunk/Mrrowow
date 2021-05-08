package executor

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type AliExecutor struct {
	AccessKey       string
	AccessKeySecret string
	SdkClient       *ecs.Client
	language string
	Scheme string
}

func (execContext AliExecutor) CreateClient(regionId string) (*ecs.Client, error) {
	client, err := ecs.NewClientWithAccessKey(regionId, execContext.AccessKey, execContext.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (execContext AliExecutor) DescribeRegions(resourceType string, instanceChargeType string) ([]ecs.Region, error) {
	if execContext.SdkClient == nil {
		client, err := execContext.CreateClient("cn-hangzhou")
		if err != nil {
			return nil, err
		}
		execContext.SdkClient = client
	}

	req := ecs.CreateDescribeRegionsRequest()
	req.Scheme = execContext.Scheme
	req.InstanceChargeType = instanceChargeType
	req.ResourceType = resourceType
	req.AcceptLanguage = execContext.language

	res, err := execContext.SdkClient.DescribeRegions(req)

	if err != nil {
		return nil, err
	}

	return res.Regions.Region,nil
}
