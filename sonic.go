package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/devfacet/gocmd"
	"math"
)

// Enabling shared config, ensures defaults are loaded from ~/.aws/config as well as ~/.aws/credentials
// This ensures auto-discovery of items such as region, and profiles
// Always initiate sessions like:
// sess := session.Must(session.NewSessionWithOptions(session.Options{
// 	SharedConfigState: session.SharedConfigEnable,
// }))

type instanceSummary struct {
	InstanceID       string `header: instance-id`
	InstanceType     string `header: instance-type`
	AccountID        string `header: account-id`
	ImageId          string `header: image-id`
	LaunchTime       string `header: Launch Time`
	KeyName          string
	AvailabilityZone string
	PrivateIpAddress string
	PrivateDnsName   string
	PublicDnsName    string
	SecurityGroups   []string
	// Tags map[string]string
}

func main() {
	flags := struct {
		Help    bool `short:"h" long:"help" description:"Display usage" global:"true"`
		Version bool `short:"v" long:"version" description:"Display version"`
		Debug   struct {
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		} `command:"debug" description:"Print arguments"`
		Info struct {
			Full     bool `short:"f" long:"full" required:"false" description:"show full output"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		} `command:"info" description:"Show information about a compute resource"`
		Math struct {
			Sqrt struct {
				Number float64 `short:"n" long:"number" required:"true" description:"Number"`
			} `command:"sqrt" description:"Calculate square root"`
			Pow struct {
				Base     float64 `short:"b" long:"base" required:"true" description:"Base"`
				Exponent float64 `short:"e" long:"exponent" required:"true" description:"Exponent"`
			} `command:"pow" description:"Calculate base exponential"`
		} `command:"math" description:"Math functions" nonempty:"true"`
	}{}

	// Echo command
	gocmd.HandleFlag("Debug", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println("Sonic Debug info: aws sts get-caller-identity:")

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		stssvc := sts.New(sess)
		input := &sts.GetCallerIdentityInput{}

		result, err := stssvc.GetCallerIdentity(input)
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
			return nil
		}

		fmt.Println(result)

		return nil
	})

	gocmd.HandleFlag("Info", func(cmd *gocmd.Cmd, args []string) error {
		InfoCommand(args[1], flags.Info.Full)

		// var instanceidentifier = args[1]
		// fmt.Println("Searching for: " + instanceidentifier + "\n")
		//
		// sess := session.Must(session.NewSessionWithOptions(session.Options{
		// 	SharedConfigState: session.SharedConfigEnable,
		// }))
		//
		// if strings.HasPrefix(instanceidentifier, "i-") {
		// 	// Handle queries for an aws ec2 instance-identifier
		//
		// 	svc := ec2.New(sess)
		// 	input := &ec2.DescribeInstancesInput{
		// 		InstanceIds: []*string{
		// 			aws.String(instanceidentifier),
		// 		},
		// 		// Filters: []*ec2.Filter{
		// 		// 	{
		// 		// 		Name: aws.String("instance-type"),
		// 		// 		Values: []*string{
		// 		// 			aws.String(instancetype),
		// 		// 		},
		// 		// 	},
		// 		// },
		// 	}
		//
		// 	result, err := svc.DescribeInstances(input)
		// 	if err != nil {
		// 		if aerr, ok := err.(awserr.Error); ok {
		// 			switch aerr.Code() {
		// 			default:
		// 				fmt.Println(aerr.Error())
		// 			}
		// 		} else {
		// 			// Print the error, cast err to awserr.Error to get the Code and
		// 			// Message from an error.
		// 			fmt.Println(err.Error())
		// 		}
		// 		return nil
		// 	}
		//
		// 	if flags.Info.Full {
		// 		fmt.Println(result)
		// 	} else {
		// 		t := table.NewWriter()
		// 		t.SetOutputMirror(os.Stdout)
		// 		t.SetStyle(table.StyleColoredDark)
		// 		t.SetIndexColumn(1)
		//
		// 		t.SetTitle(instanceidentifier)
		// 		t.AppendRows([]table.Row{
		// 			{"InstanceType", *result.Reservations[0].Instances[0].InstanceType},
		// 			{"ImageId", *result.Reservations[0].Instances[0].ImageId},
		// 		})
		//
		// 		t.Render()
		// 	}
		//
		// } else if strings.HasPrefix(instanceidentifier, "mi-") {
		// 	fmt.Println("TODO: querying managed instances via SSM is not yet supported")
		// } else {
		// 	fmt.Println("Unsupported query format: " + instanceidentifier)
		// }

		return nil
	})

	// Math commands
	gocmd.HandleFlag("Math.Sqrt", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println(math.Sqrt(flags.Math.Sqrt.Number))
		return nil
	})
	gocmd.HandleFlag("Math.Pow", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println(math.Pow(flags.Math.Pow.Base, flags.Math.Pow.Exponent))
		return nil
	})

	// Init the app
	gocmd.New(gocmd.Options{
		Name:        "sonic",
		Version:     "1.0.0",
		Description: "Find and connect wherever you need to go",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}