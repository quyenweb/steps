package main

import (
	"encoding/json"
	"fmt"

	"github.com/stackpulse/steps-sdk-go/env"
	"github.com/stackpulse/steps-sdk-go/step"
	"github.com/stackpulse/steps/atlassian/jira/base"
	"golang.org/x/oauth2"
)

type JiraDeleteIssue struct {
	base.Args
	IssueID string `env:"ISSUE_ID,required"`
	token   *oauth2.Token
}

type output struct {
	Success bool `json:"success"`
	step.Outputs
}

func (j *JiraDeleteIssue) Init() error {
	err := env.Parse(j)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(j.JiraToken), &j.token); err != nil {
		return fmt.Errorf("unmarshal token: %w", err)
	}

	return nil
}

func (j *JiraDeleteIssue) Run() (int, []byte, error) {
	jiraClient, err := base.NewClient(j.token, j.SiteName)
	if err != nil {
		return step.ExitCodeFailure, nil, err
	}

	resp, err := jiraClient.Issue.Delete(j.IssueID)
	if err != nil {
		jsonOutput, err := json.Marshal(&output{
			Success: false,
			Outputs: step.Outputs{Object: base.ExtractResponse(resp)},
		})
		if err == nil {
			return step.ExitCodeFailure, jsonOutput, fmt.Errorf("delete issue: %w", err)
		}
		return step.ExitCodeFailure, nil, fmt.Errorf("marshal output: %w", err)
	}

	jsonOutput, err := json.Marshal(&output{
		Success: true,
		Outputs: step.Outputs{Object: base.ExtractResponse(resp)},
	})
	if err != nil {
		return step.ExitCodeFailure, nil, fmt.Errorf("marshal output: %w", err)
	}

	return step.ExitCodeOK, jsonOutput, nil
}

func main() {
	step.Run(&JiraDeleteIssue{})
}
