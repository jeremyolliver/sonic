package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/jedib0t/go-pretty/table"
	// "os"
	"strings"
)

// Returns boolean for true if online, and table rows displaying status
func GetSSMStatusRows(instanceidentifier string) (online bool, statusrows []table.Row) {
	// fmt.Println("Listing SSM Managed Instances")
	svc := ssm.New(SharedAWSSession())

	input := &ssm.DescribeInstanceInformationInput{
		Filters: []*ssm.InstanceInformationStringFilter{
			{
				Key: aws.String("InstanceIds"),
				Values: []*string{
					aws.String(instanceidentifier),
				},
			},
		},
	}

	result, err := svc.DescribeInstanceInformation(input)
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
		return false, []table.Row{}
	}

	if len(result.InstanceInformationList) == 0 {
		// No SSM Results for this instance, skip displaying information
		return false, []table.Row{}
	}

	// Print the raw nested data structure
	fmt.Println(result)

	var statusitemkeys []string
	var statusitemvalues []string

	statusitemkeys = append(statusitemkeys,
		"Ping Status",
		"Platform",
	)
	statusitemvalues = append(statusitemvalues,
		*result.InstanceInformationList[0].PingStatus,
		*result.InstanceInformationList[0].PlatformName+" "+*result.InstanceInformationList[0].PlatformVersion,
	)

	return (*result.InstanceInformationList[0].PingStatus == "Online"), []table.Row{
		{"SSM", strings.Join(statusitemkeys, "\n"), strings.Join(statusitemvalues, "\n")},
	}

}
