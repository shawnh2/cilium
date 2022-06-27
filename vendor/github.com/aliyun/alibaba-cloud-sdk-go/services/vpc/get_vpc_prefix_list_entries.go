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

// GetVpcPrefixListEntries invokes the vpc.GetVpcPrefixListEntries API synchronously
func (client *Client) GetVpcPrefixListEntries(request *GetVpcPrefixListEntriesRequest) (response *GetVpcPrefixListEntriesResponse, err error) {
	response = CreateGetVpcPrefixListEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// GetVpcPrefixListEntriesWithChan invokes the vpc.GetVpcPrefixListEntries API asynchronously
func (client *Client) GetVpcPrefixListEntriesWithChan(request *GetVpcPrefixListEntriesRequest) (<-chan *GetVpcPrefixListEntriesResponse, <-chan error) {
	responseChan := make(chan *GetVpcPrefixListEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVpcPrefixListEntries(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetVpcPrefixListEntriesWithCallback invokes the vpc.GetVpcPrefixListEntries API asynchronously
func (client *Client) GetVpcPrefixListEntriesWithCallback(request *GetVpcPrefixListEntriesRequest, callback func(response *GetVpcPrefixListEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVpcPrefixListEntriesResponse
		var err error
		defer close(result)
		response, err = client.GetVpcPrefixListEntries(request)
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

// GetVpcPrefixListEntriesRequest is the request struct for api GetVpcPrefixListEntries
type GetVpcPrefixListEntriesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PrefixListId         string           `position:"Query" name:"PrefixListId"`
	NextToken            string           `position:"Query" name:"NextToken"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
}

// GetVpcPrefixListEntriesResponse is the response struct for api GetVpcPrefixListEntries
type GetVpcPrefixListEntriesResponse struct {
	*responses.BaseResponse
	RequestId       string            `json:"RequestId" xml:"RequestId"`
	NextToken       string            `json:"NextToken" xml:"NextToken"`
	TotalCount      int64             `json:"TotalCount" xml:"TotalCount"`
	Count           int64             `json:"Count" xml:"Count"`
	PrefixListEntry []PrefixListCidrs `json:"PrefixListEntry" xml:"PrefixListEntry"`
}

// CreateGetVpcPrefixListEntriesRequest creates a request to invoke GetVpcPrefixListEntries API
func CreateGetVpcPrefixListEntriesRequest() (request *GetVpcPrefixListEntriesRequest) {
	request = &GetVpcPrefixListEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "GetVpcPrefixListEntries", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVpcPrefixListEntriesResponse creates a response to parse from GetVpcPrefixListEntries response
func CreateGetVpcPrefixListEntriesResponse() (response *GetVpcPrefixListEntriesResponse) {
	response = &GetVpcPrefixListEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}