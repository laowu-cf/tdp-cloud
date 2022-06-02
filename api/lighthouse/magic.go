package lighthouse

import (
	"sync"

	"github.com/gin-gonic/gin"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/qcloud"
)

// 获取所有地域和实例列表

func getAllRegionsInstances(c *gin.Context) {

	var result = make(map[string]interface{})

	// 获取所有地域
	regionsClient := qcloud.NewLighthouseClient(c, "")
	regionsRequest := lighthouse.NewDescribeRegionsRequest()
	regionsResponse, err := regionsClient.DescribeRegions(regionsRequest)

	if err != nil || regionsResponse.Response.RegionSet == nil {
		c.Set("Error", err)
		return
	}

	var wg sync.WaitGroup
	var instanceSet []*lighthouse.Instance

	// 获取所有地域的实例
	for _, region := range regionsResponse.Response.RegionSet {
		wg.Add(1)

		go func(r string) {
			regionsClient := qcloud.NewLighthouseClient(c, r)
			instancesRequest := lighthouse.NewDescribeInstancesRequest()
			instanceResponse, er2 := regionsClient.DescribeInstances(instancesRequest)

			if er2 == nil && instanceResponse.Response.InstanceSet != nil {
				instanceSet = append(instanceSet, instanceResponse.Response.InstanceSet...)
			}

			wg.Done()
		}(*region.Region)
	}

	wg.Wait()

	result["RegionSet"] = regionsResponse.Response.RegionSet
	result["InstanceSet"] = instanceSet

	c.Set("Payload", result)

}
