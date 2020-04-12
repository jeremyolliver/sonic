package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/sts"
)

var sharedsession *session.Session = nil

// Enabling SharedConfigEnable, ensures defaults are loaded from ~/.aws/config as well as ~/.aws/credentials
// This ensures auto-discovery of items such as region, and profiles
func SharedAWSSession() *session.Session {
	if sharedsession != nil {
		return sharedsession
	} else {
		sharedsession = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	}
	return sharedsession
}

// TODO: memoize display of more than one account for multi-account support
// var accountDisplay map[string]string
var accountDisplay *string = nil

// Display the current account
// Uses the given IAM Alias, if there is one (and has permission to query)
// Otherwise uses the account-id from checking the current callery identity
func AWSAccountDisplay() string {
	if accountDisplay != nil {
		// quick return memoized value if we can
		return *accountDisplay
	} else {
		svc := iam.New(SharedAWSSession())
		input := &iam.ListAccountAliasesInput{}

		result, err := svc.ListAccountAliases(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case iam.ErrCodeServiceFailureException:
					fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			// something went wrong - skip storing friendly IAM alias, and just use the account id
		} else {
			// list account aliases succeeded
			if len(result.AccountAliases) > 0 {
				// we found an IAM alias. memoize it and return it
				accountDisplay = result.AccountAliases[0]
			} else {
				// there is no IAM alias
				accountDisplay = nil
			}
		}

		if accountDisplay == nil {
			// fall-back to just showing the account-id if we didn't succeed in finding an alias
			accountDisplay = AWSAccountID()
		}

		return *accountDisplay
	}
}

var accountID *string = nil

func AWSAccountID() *string {
	if accountID != nil {
		return accountID
	} else {
		stssvc := sts.New(SharedAWSSession())
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
		} else {
			accountID = result.Account
		}
	}
	return accountID
}
