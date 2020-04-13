package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jedib0t/go-pretty/table"
	"os"
	"strings"
)

func DescribeEC2Instance(instanceidentifier string, fullOutput bool) {
	svc := ec2.New(SharedAWSSession())
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceidentifier),
		},
		// Filters: []*ec2.Filter{
		// 	{
		// 		Name: aws.String("instance-type"),
		// 		Values: []*string{
		// 			aws.String(instancetype),
		// 		},
		// 	},
		// },
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	// Display Formatting:
	// By default, if one instance result is returned, a table summary is printed
	// if --full option is passed, the struct is printed raw
	if fullOutput {
		// Print the raw nested data structure
		fmt.Println(result)
	} else {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleColoredDark)
		t.SetIndexColumn(1)

		instanceDetails := *result.Reservations[0].Instances[0]

		var instanceName string = ""
		var tagKeysDisplay []string
		var tagValuesDisplay []string
		for _, item := range instanceDetails.Tags {
			if *item.Key == "Name" {
				instanceName = *item.Value
			}
			tagKeysDisplay = append(tagKeysDisplay, *item.Key)
			tagValuesDisplay = append(tagValuesDisplay, *item.Value)
		}

		var securityGroupIDs []string
		var securityGroupNames []string
		for _, item := range instanceDetails.SecurityGroups {
			securityGroupIDs = append(securityGroupIDs, *item.GroupId)
			securityGroupNames = append(securityGroupNames, *item.GroupName)
		}

		if instanceName == "" {
			instanceName = instanceidentifier
		}

		t.SetTitle(instanceName + " (Account: " + AWSAccountDisplay() + ")")

		t.AppendRows([]table.Row{
			{"Instance-ID", instanceidentifier, ""},
			{"InstanceType", *instanceDetails.InstanceType, ""},
			{"KeyName", *instanceDetails.KeyName, ""},
			{"AvailabilityZone", *instanceDetails.Placement.AvailabilityZone, ""},
			{"PrivateIpAddress", *instanceDetails.PrivateIpAddress, ""},
			{"PrivateDnsName", *instanceDetails.PrivateDnsName, ""},
			{"PublicDnsName", *instanceDetails.PublicDnsName, ""},
			{"LaunchTime", instanceDetails.LaunchTime, ""},
			{"ImageId", *instanceDetails.ImageId, ""},
			{"Tags", strings.Join(tagKeysDisplay, "\n"), strings.Join(tagValuesDisplay, "\n")},
			{"SecurityGroups", strings.Join(securityGroupIDs, "\n"), strings.Join(securityGroupNames, "\n")},
		})

		t.Render()
	}
}
