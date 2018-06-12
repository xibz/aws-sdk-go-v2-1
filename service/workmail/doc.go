// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package workmail provides the client and types for making API
// requests to Amazon WorkMail.
//
// Amazon WorkMail is a secure, managed business email and calendaring service
// with support for existing desktop and mobile email clients. You can access
// your email, contacts, and calendars using Microsoft Outlook, your browser,
// or their native iOS and Android email applications. You can integrate Amazon
// WorkMail with your existing corporate directory and control both the keys
// that encrypt your data and the location in which your data is stored.
//
// The Amazon WorkMail API is designed for the following scenarios:
//
//    * Listing and describing organizations
//
//    * Managing users
//
//    * Managing groups
//
//    * Managing resources
//
// All Amazon WorkMail API actions are Amazon-authenticated and certificate-signed.
// They not only require the use of the AWS SDK, but also allow for the exclusive
// use of IAM users and roles to help facilitate access, trust, and permission
// policies. By creating a role and allowing an IAM user to access the Amazon
// WorkMail site, the IAM user gains full administrative visibility into the
// entire Amazon WorkMail organization (or as set in the IAM policy). This includes,
// but is not limited to, the ability to create, update, and delete users, groups,
// and resources. This allows developers to perform the scenarios listed above,
// as well as give users the ability to grant access on a selective basis using
// the IAM model.
//
// See https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01 for more information on this service.
//
// See workmail package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/workmail/
//
// Using the Client
//
// To Amazon WorkMail with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon WorkMail client WorkMail for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/workmail/#New
package workmail
