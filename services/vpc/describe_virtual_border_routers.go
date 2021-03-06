package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeVirtualBorderRouters(request *DescribeVirtualBorderRoutersRequest) (response *DescribeVirtualBorderRoutersResponse, err error) {
	response = CreateDescribeVirtualBorderRoutersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeVirtualBorderRoutersWithChan(request *DescribeVirtualBorderRoutersRequest) (<-chan *DescribeVirtualBorderRoutersResponse, <-chan error) {
	responseChan := make(chan *DescribeVirtualBorderRoutersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVirtualBorderRouters(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeVirtualBorderRoutersWithCallback(request *DescribeVirtualBorderRoutersRequest, callback func(response *DescribeVirtualBorderRoutersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVirtualBorderRoutersResponse
		var err error
		defer close(result)
		response, err = client.DescribeVirtualBorderRouters(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeVirtualBorderRoutersRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer                      `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string                                `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer                      `position:"Query" name:"PageNumber"`
	ResourceOwnerId      requests.Integer                      `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer                      `position:"Query" name:"OwnerId"`
	Filter               *[]DescribeVirtualBorderRoutersFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

type DescribeVirtualBorderRoutersFilter struct {
	Key   string    `name:"Key"`
	Value *[]string `name:"Value" type:"Repeated"`
}

type DescribeVirtualBorderRoutersResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	PageNumber             int    `json:"PageNumber" xml:"PageNumber"`
	PageSize               int    `json:"PageSize" xml:"PageSize"`
	TotalCount             int    `json:"TotalCount" xml:"TotalCount"`
	VirtualBorderRouterSet struct {
		VirtualBorderRouterType []struct {
			VbrId                            string `json:"VbrId" xml:"VbrId"`
			CreationTime                     string `json:"CreationTime" xml:"CreationTime"`
			ActivationTime                   string `json:"ActivationTime" xml:"ActivationTime"`
			TerminationTime                  string `json:"TerminationTime" xml:"TerminationTime"`
			RecoveryTime                     string `json:"RecoveryTime" xml:"RecoveryTime"`
			Status                           string `json:"Status" xml:"Status"`
			VlanId                           int    `json:"VlanId" xml:"VlanId"`
			CircuitCode                      string `json:"CircuitCode" xml:"CircuitCode"`
			RouteTableId                     string `json:"RouteTableId" xml:"RouteTableId"`
			VlanInterfaceId                  string `json:"VlanInterfaceId" xml:"VlanInterfaceId"`
			LocalGatewayIp                   string `json:"LocalGatewayIp" xml:"LocalGatewayIp"`
			PeerGatewayIp                    string `json:"PeerGatewayIp" xml:"PeerGatewayIp"`
			PeeringSubnetMask                string `json:"PeeringSubnetMask" xml:"PeeringSubnetMask"`
			PhysicalConnectionId             string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
			PhysicalConnectionStatus         string `json:"PhysicalConnectionStatus" xml:"PhysicalConnectionStatus"`
			PhysicalConnectionBusinessStatus string `json:"PhysicalConnectionBusinessStatus" xml:"PhysicalConnectionBusinessStatus"`
			PhysicalConnectionOwnerUid       string `json:"PhysicalConnectionOwnerUid" xml:"PhysicalConnectionOwnerUid"`
			AccessPointId                    string `json:"AccessPointId" xml:"AccessPointId"`
			Name                             string `json:"Name" xml:"Name"`
			Description                      string `json:"Description" xml:"Description"`
			AssociatedPhysicalConnections    struct {
				AssociatedPhysicalConnection []struct {
					CircuitCode                      string `json:"CircuitCode" xml:"CircuitCode"`
					VlanInterfaceId                  string `json:"VlanInterfaceId" xml:"VlanInterfaceId"`
					LocalGatewayIp                   string `json:"LocalGatewayIp" xml:"LocalGatewayIp"`
					PeerGatewayIp                    string `json:"PeerGatewayIp" xml:"PeerGatewayIp"`
					PeeringSubnetMask                string `json:"PeeringSubnetMask" xml:"PeeringSubnetMask"`
					PhysicalConnectionId             string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
					PhysicalConnectionStatus         string `json:"PhysicalConnectionStatus" xml:"PhysicalConnectionStatus"`
					PhysicalConnectionBusinessStatus string `json:"PhysicalConnectionBusinessStatus" xml:"PhysicalConnectionBusinessStatus"`
					PhysicalConnectionOwnerUid       string `json:"PhysicalConnectionOwnerUid" xml:"PhysicalConnectionOwnerUid"`
					VlanId                           string `json:"VlanId" xml:"VlanId"`
				} `json:"AssociatedPhysicalConnection" xml:"AssociatedPhysicalConnection"`
			} `json:"AssociatedPhysicalConnections" xml:"AssociatedPhysicalConnections"`
			AssociatedCbns struct {
				AssociatedCbn []struct {
					CbnStatus   string `json:"CbnStatus" xml:"CbnStatus"`
					CbnOwnerId  string `json:"CbnOwnerId" xml:"CbnOwnerId"`
					CbnOwnerUid int    `json:"CbnOwnerUid" xml:"CbnOwnerUid"`
				} `json:"AssociatedCbn" xml:"AssociatedCbn"`
			} `json:"AssociatedCbns" xml:"AssociatedCbns"`
		} `json:"VirtualBorderRouterType" xml:"VirtualBorderRouterType"`
	} `json:"VirtualBorderRouterSet" xml:"VirtualBorderRouterSet"`
}

func CreateDescribeVirtualBorderRoutersRequest() (request *DescribeVirtualBorderRoutersRequest) {
	request = &DescribeVirtualBorderRoutersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVirtualBorderRouters", "vpc", "openAPI")
	return
}

func CreateDescribeVirtualBorderRoutersResponse() (response *DescribeVirtualBorderRoutersResponse) {
	response = &DescribeVirtualBorderRoutersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
