
// This file contains the data structures used to receive lists of job templates.

package data

type WorkflowJobTemplatesGetResponse struct {
	ListGetResponse

	Results []*WorkflowJobTemplate `json:"results,omitempty"`
}
