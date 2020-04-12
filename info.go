package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jedib0t/go-pretty/table"
	"os"
	"strings"
)

func InfoCommand(instanceidentifier string, fullOutput bool) {
	fmt.Println("Searching for: " + instanceidentifier + "\n")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	if strings.HasPrefix(instanceidentifier, "i-") {
		// Handle queries for an aws ec2 instance-identifier

		svc := ec2.New(sess)
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
			t.AppendRows([]table.Row{
				{"InstanceType", *result.Reservations[0].Instances[0].InstanceType},
				{"ImageId", *result.Reservations[0].Instances[0].ImageId},
			})

			t.Render()
		}

	} else if strings.HasPrefix(instanceidentifier, "mi-") {
		fmt.Println("TODO: querying managed instances via SSM is not yet supported")
	} else {
		fmt.Println("Unsupported query format: " + instanceidentifier)
	}
}
