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

// This file contains the implementation of the resource that manages a specific job
// template.

package awx

import (
	"github.com/bluearchive/awx-client-go/awx/internal/data"
)

type WorkflowJobTemplateResource struct {
	Resource
}

func NewWorkflowJobTemplateResource(connection *Connection, path string) *WorkflowJobTemplateResource {
	resource := new(WorkflowJobTemplateResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkflowJobTemplateResource) Get() *WorkflowJobTemplateGetRequest {
	request := new(WorkflowJobTemplateGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *WorkflowJobTemplateResource) Launch() *WorkflowJobTemplateLaunchResource {
	return NewWorkflowJobTemplateLaunchResource(r.connection, r.path+"/launch")
}

type WorkflowJobTemplateGetRequest struct {
	Request
}

func (r *WorkflowJobTemplateGetRequest) Send() (response *WorkflowJobTemplateGetResponse, err error) {
	output := new(data.WorkflowJobTemplateGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(WorkflowJobTemplateGetResponse)
	response.result = new(WorkflowJobTemplate)
	response.result.id = output.Id
	response.result.name = output.Name
	response.result.askInventoryOnLaunch = output.AskInventoryOnLaunch
	response.result.askVarsOnLaunch = output.AskVarsOnLaunch

	return
}

type WorkflowJobTemplateGetResponse struct {
	result *WorkflowJobTemplate
}

func (r *WorkflowJobTemplateGetResponse) Result() *WorkflowJobTemplate {
	return r.result
}
