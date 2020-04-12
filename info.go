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

func InfoCommand(instanceidentifier string, fullOutput bool) {
	fmt.Println("Searching for: " + instanceidentifier + "...")

	if strings.HasPrefix(instanceidentifier, "i-") {
		DescribeEC2Instance(instanceidentifier, fullOutput)

	} else if strings.HasPrefix(instanceidentifier, "mi-") {
		fmt.Println("TODO: querying managed instances via SSM is not yet supported")
	} else {
		fmt.Println("Unsupported query format: " + instanceidentifier)
	}
}

func DescribeEC2Instance(instanceidentifier string, fullOutput bool) {
	fmt.Println("Describing ec2 instance\n")

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

	if fullOutput {
		fmt.Println(result)
	} else {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleColoredDark)
		t.SetIndexColumn(1)

		t.SetTitle(instanceidentifier)

		// type instanceSummary struct {
		// 	InstanceID       string `header: instance-id`
		// 	InstanceType     string `header: instance-type`
		// 	AccountID        string `header: account-id`
		// 	ImageId          string `header: image-id`
		// 	LaunchTime       string `header: Launch Time`
		// 	KeyName          string
		// 	AvailabilityZone string
		// 	PrivateIpAddress string
		// 	PrivateDnsName   string
		// 	PublicDnsName    string
		// 	SecurityGroups   []string
		// 	// Tags map[string]string
		// }

		t.AppendRows([]table.Row{
			{"InstanceType", *result.Reservations[0].Instances[0].InstanceType},
			{"ImageId", *result.Reservations[0].Instances[0].ImageId},
		})

		t.Render()
	}
}
