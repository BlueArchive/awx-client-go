
// This file contains the data structures used to launch jobs from job templates.

package data

type WorkflowJobTemplateLaunchGetResponse struct {
	WorkflowJobTemplateData *WorkflowJobTemplateGetResponse `json:"workflow_job_template_data,omitempty"`
}

type WorkflowJobTemplateLaunchPostRequest struct {
	ExtraVars string `json:"extra_vars,omitempty"`
	Limit     string `json:"limit,omitempty"`
}

type WorkflowJobTemplateLaunchPostResponse struct {
	Job int `json:"job,omitempty"`
}
