/*
Copyright (c) 2018 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file contains the implementation of the resource that manages launching of jobs from job
// templates.

package awx

import (
	"encoding/json"

	"github.com/bluearchive/awx-client-go/awx/internal/data"
)

type WorkflowJobTemplateLaunchResource struct {
	Resource
}

func NewWorkflowJobTemplateLaunchResource(connection *Connection, path string) *WorkflowJobTemplateLaunchResource {
	resource := new(WorkflowJobTemplateLaunchResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkflowJobTemplateLaunchResource) Get() *WorkflowJobTemplateLaunchGetRequest {
	request := new(WorkflowJobTemplateLaunchGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *WorkflowJobTemplateLaunchResource) Post() *WorkflowJobTemplateLaunchPostRequest {
	request := new(WorkflowJobTemplateLaunchPostRequest)
	request.resource = &r.Resource
	return request
}

type WorkflowJobTemplateLaunchGetRequest struct {
	Request
}

func (r *WorkflowJobTemplateLaunchGetRequest) Send() (response *WorkflowJobTemplateLaunchGetResponse, err error) {
	output := new(data.WorkflowJobTemplateLaunchGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(WorkflowJobTemplateLaunchGetResponse)
	if output.WorkflowJobTemplateData != nil {
		response.workflowJobTemplateData = new(WorkflowJobTemplate)
		response.workflowJobTemplateData.id = output.WorkflowJobTemplateData.Id
		response.workflowJobTemplateData.name = output.WorkflowJobTemplateData.Name
	}
	return
}

type WorkflowJobTemplateLaunchGetResponse struct {
	workflowJobTemplateData *WorkflowJobTemplate
}

func (r *WorkflowJobTemplateLaunchGetResponse) WorkflowJobTemplateData() *WorkflowJobTemplate {
	return r.workflowJobTemplateData
}

type WorkflowJobTemplateLaunchPostRequest struct {
	Request

	extraVars map[string]interface{}
	inventory string
}

// ExtraVars set a map or external variables sent to the AWX job.
func (r *WorkflowJobTemplateLaunchPostRequest) ExtraVars(value map[string]interface{}) *WorkflowJobTemplateLaunchPostRequest {
	r.extraVars = value
	return r
}

// ExtraVar adds a single external variable to extraVars map.
func (r *WorkflowJobTemplateLaunchPostRequest) ExtraVar(name string, value interface{}) *WorkflowJobTemplateLaunchPostRequest {
	if r.extraVars == nil {
		r.extraVars = make(map[string]interface{})
	}
	r.extraVars[name] = value
	return r
}

func (r *WorkflowJobTemplateLaunchPostRequest) Inventory(value string) *WorkflowJobTemplateLaunchPostRequest {
	r.inventory = value
	return r
}

func (r *WorkflowJobTemplateLaunchPostRequest) Send() (response *WorkflowJobTemplateLaunchPostResponse, err error) {
	// Generate the input data:
	input := new(data.WorkflowJobTemplateLaunchPostRequest)

	if r.extraVars != nil {
		// convert extravars json to string
		var bytes []byte
		bytes, err = json.Marshal(r.extraVars)
		if err != nil {
			return
		}
		input.ExtraVars = string(bytes)
	}

	input.Inventory = r.inventory

	// Send the request:
	output := new(data.WorkflowJobTemplateLaunchPostResponse)
	err = r.post(input, output)
	if err != nil {
		return
	}

	// Analyze the output data:
	response = new(WorkflowJobTemplateLaunchPostResponse)
	response.Job = output.Job

	return
}

type WorkflowJobTemplateLaunchPostResponse struct {
	Job int
}