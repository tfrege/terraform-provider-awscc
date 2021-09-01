package naming_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
)

func TestCloudFormationPropertyToTerraformAttribute(t *testing.T) {
	testCases := []struct {
		TestName      string
		Value         string
		ExpectedValue string
	}{
		{
			TestName:      "empty string",
			Value:         "",
			ExpectedValue: "",
		},
		{
			TestName:      "whitespace string",
			Value:         "  ",
			ExpectedValue: "",
		},
		{
			TestName:      "short property name",
			Value:         "Arn",
			ExpectedValue: "arn",
		},
		{
			TestName:      "long property name",
			Value:         "GlobalReplicationGroupDescription",
			ExpectedValue: "global_replication_group_description",
		},
		{
			TestName:      "including digit",
			Value:         "S3Bucket",
			ExpectedValue: "s3_bucket",
		},
		{
			TestName:      "including multiple digits",
			Value:         "AWS99Thing",
			ExpectedValue: "aws99_thing",
		},
		{
			TestName:      "including replacement",
			Value:         "CloudWatchLogsLogGroupArn",
			ExpectedValue: "cloudwatch_logs_log_group_arn",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got := naming.CloudFormationPropertyToTerraformAttribute(testCase.Value)

			if got != testCase.ExpectedValue {
				t.Errorf("expected: %s, got: %s", testCase.ExpectedValue, got)
			}
		})
	}
}

func TestPluralize(t *testing.T) {
	testCases := []struct {
		TestName      string
		Value         string
		ExpectedValue string
	}{
		{
			TestName:      "empty string",
			Value:         "",
			ExpectedValue: "",
		},
		{
			TestName:      "name ending in s",
			Value:         "aws_cloudwatch_event_bus",
			ExpectedValue: "aws_cloudwatch_event_buses",
		},
		{
			TestName:      "special name ending in s",
			Value:         "awscc_ec2_network_insights_analysis",
			ExpectedValue: "awscc_ec2_network_insights_analyses",
		},
		{
			TestName:      "name ending in number",
			Value:         "aws_datasync_location_s3",
			ExpectedValue: "aws_datasync_location_s3s",
		},
		{
			TestName:      "singular name",
			Value:         "aws_wafv2_web_acl",
			ExpectedValue: "aws_wafv2_web_acls",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got := naming.Pluralize(testCase.Value)

			if got != testCase.ExpectedValue {
				t.Errorf("expected: %s, got: %s", testCase.ExpectedValue, got)
			}
		})
	}
}
