// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opGenerateDataSet = "GenerateDataSet"

// GenerateDataSetRequest is a API request type for the GenerateDataSet API operation.
type GenerateDataSetRequest struct {
	*aws.Request
	Input *GenerateDataSetInput
}

// Send marshals and sends the GenerateDataSet API request.
func (r GenerateDataSetRequest) Send() (*GenerateDataSetOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GenerateDataSetOutput), nil
}

// GenerateDataSetRequest returns a request value for making API operation for
// AWS Marketplace Commerce Analytics.
//
// Given a data set type and data set publication date, asynchronously publishes
// the requested data set to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD.csv. If a file with the same name already
// exists (e.g. if the same data set is requested twice), the original file
// will be overwritten by the new file. Requires a Role with an attached permissions
// policy providing Allow permissions for the following actions: s3:PutObject,
// s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish, iam:GetRolePolicy.
//
//    // Example sending a request using the GenerateDataSetRequest method.
//    req := client.GenerateDataSetRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/GenerateDataSet
func (c *MarketplaceCommerceAnalytics) GenerateDataSetRequest(input *GenerateDataSetInput) GenerateDataSetRequest {
	op := &aws.Operation{
		Name:       opGenerateDataSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateDataSetInput{}
	}

	output := &GenerateDataSetOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GenerateDataSetRequest{Request: req, Input: input}
}

const opStartSupportDataExport = "StartSupportDataExport"

// StartSupportDataExportRequest is a API request type for the StartSupportDataExport API operation.
type StartSupportDataExportRequest struct {
	*aws.Request
	Input *StartSupportDataExportInput
}

// Send marshals and sends the StartSupportDataExport API request.
func (r StartSupportDataExportRequest) Send() (*StartSupportDataExportOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*StartSupportDataExportOutput), nil
}

// StartSupportDataExportRequest returns a request value for making API operation for
// AWS Marketplace Commerce Analytics.
//
// Given a data set type and a from date, asynchronously publishes the requested
// customer support data to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD'T'HH-mm-ss'Z'.csv. If a file with the
// same name already exists (e.g. if the same data set is requested twice),
// the original file will be overwritten by the new file. Requires a Role with
// an attached permissions policy providing Allow permissions for the following
// actions: s3:PutObject, s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish,
// iam:GetRolePolicy.
//
//    // Example sending a request using the StartSupportDataExportRequest method.
//    req := client.StartSupportDataExportRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/StartSupportDataExport
func (c *MarketplaceCommerceAnalytics) StartSupportDataExportRequest(input *StartSupportDataExportInput) StartSupportDataExportRequest {
	op := &aws.Operation{
		Name:       opStartSupportDataExport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartSupportDataExportInput{}
	}

	output := &StartSupportDataExportOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return StartSupportDataExportRequest{Request: req, Input: input}
}

// Container for the parameters to the GenerateDataSet operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/GenerateDataSetRequest
type GenerateDataSetInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file. These key-value
	// pairs can be used to correlated responses with tracking information from
	// other systems.
	CustomerDefinedValues map[string]string `locationName:"customerDefinedValues" min:"1" type:"map"`

	// The date a data set was published. For daily data sets, provide a date with
	// day-level granularity for the desired day. For weekly data sets, provide
	// a date with day-level granularity within the desired week (the day value
	// will be ignored). For monthly data sets, provide a date with month-level
	// granularity for the desired month (the day value will be ignored).
	//
	// DataSetPublicationDate is a required field
	DataSetPublicationDate *time.Time `locationName:"dataSetPublicationDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The desired data set type.
	//
	// customer_subscriber_hourly_monthly_subscriptionsFrom 2014-07-21 to present:
	// Available daily by 5:00 PM Pacific Time.
	//
	// customer_subscriber_annual_subscriptionsFrom 2014-07-21 to present: Available
	// daily by 5:00 PM Pacific Time.
	//
	// daily_business_usage_by_instance_typeFrom 2015-01-26 to present: Available
	// daily by 5:00 PM Pacific Time.
	//
	// daily_business_feesFrom 2015-01-26 to present: Available daily by 5:00 PM
	// Pacific Time.
	//
	// daily_business_free_trial_conversionsFrom 2015-01-26 to present: Available
	// daily by 5:00 PM Pacific Time.
	//
	// daily_business_new_instancesFrom 2015-01-26 to present: Available daily by
	// 5:00 PM Pacific Time.
	//
	// daily_business_new_product_subscribersFrom 2015-01-26 to present: Available
	// daily by 5:00 PM Pacific Time.
	//
	// daily_business_canceled_product_subscribersFrom 2015-01-26 to present: Available
	// daily by 5:00 PM Pacific Time.
	//
	// monthly_revenue_billing_and_revenue_dataFrom 2015-02 to 2017-06: Available
	// monthly on the 4th day of the month by 5:00pm Pacific Time. Data includes
	// metered transactions (e.g. hourly) from two months prior.
	//
	// From 2017-07 to present: Available monthly on the 15th day of the month by
	// 5:00pm Pacific Time. Data includes metered transactions (e.g. hourly) from
	// one month prior.
	//
	// monthly_revenue_annual_subscriptionsFrom 2015-02 to 2017-06: Available monthly
	// on the 4th day of the month by 5:00pm Pacific Time. Data includes up-front
	// software charges (e.g. annual) from one month prior.
	//
	// From 2017-07 to present: Available monthly on the 15th day of the month by
	// 5:00pm Pacific Time. Data includes up-front software charges (e.g. annual)
	// from one month prior.
	//
	// disbursed_amount_by_productFrom 2015-01-26 to present: Available every 30
	// days by 5:00 PM Pacific Time.
	//
	// disbursed_amount_by_product_with_uncollected_fundsFrom 2012-04-19 to 2015-01-25:
	// Available every 30 days by 5:00 PM Pacific Time.
	//
	// From 2015-01-26 to present: This data set was split into three data sets:
	// disbursed_amount_by_product, disbursed_amount_by_age_of_uncollected_funds,
	// and disbursed_amount_by_age_of_disbursed_funds.
	//
	// disbursed_amount_by_instance_hoursFrom 2012-09-04 to present: Available every
	// 30 days by 5:00 PM Pacific Time.
	//
	// disbursed_amount_by_customer_geoFrom 2012-04-19 to present: Available every
	// 30 days by 5:00 PM Pacific Time.
	//
	// disbursed_amount_by_age_of_uncollected_fundsFrom 2015-01-26 to present: Available
	// every 30 days by 5:00 PM Pacific Time.
	//
	// disbursed_amount_by_age_of_disbursed_fundsFrom 2015-01-26 to present: Available
	// every 30 days by 5:00 PM Pacific Time.
	//
	// customer_profile_by_industryFrom 2015-10-01 to 2017-06-29: Available daily
	// by 5:00 PM Pacific Time.
	//
	// From 2017-06-30 to present: This data set is no longer available.
	//
	// customer_profile_by_revenueFrom 2015-10-01 to 2017-06-29: Available daily
	// by 5:00 PM Pacific Time.
	//
	// From 2017-06-30 to present: This data set is no longer available.
	//
	// customer_profile_by_geographyFrom 2015-10-01 to 2017-06-29: Available daily
	// by 5:00 PM Pacific Time.
	//
	// From 2017-06-30 to present: This data set is no longer available.
	//
	// sales_compensation_billed_revenueFrom 2016-12 to 2017-06: Available monthly
	// on the 4th day of the month by 5:00pm Pacific Time. Data includes metered
	// transactions (e.g. hourly) from two months prior, and up-front software charges
	// (e.g. annual) from one month prior.
	//
	// From 2017-06 to present: Available monthly on the 15th day of the month by
	// 5:00pm Pacific Time. Data includes metered transactions (e.g. hourly) from
	// one month prior, and up-front software charges (e.g. annual) from one month
	// prior.
	//
	// us_sales_and_use_tax_recordsFrom 2017-02-15 to present: Available monthly
	// on the 15th day of the month by 5:00 PM Pacific Time.
	//
	// DataSetType is a required field
	DataSetType DataSetType `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"true"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	//
	// DestinationS3BucketName is a required field
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	//
	// RoleNameArn is a required field
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	//
	// SnsTopicArn is a required field
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GenerateDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GenerateDataSetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateDataSetInput"}
	if s.CustomerDefinedValues != nil && len(s.CustomerDefinedValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomerDefinedValues", 1))
	}

	if s.DataSetPublicationDate == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetPublicationDate"))
	}
	if len(s.DataSetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DataSetType"))
	}

	if s.DestinationS3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationS3BucketName"))
	}
	if s.DestinationS3BucketName != nil && len(*s.DestinationS3BucketName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationS3BucketName", 1))
	}

	if s.RoleNameArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleNameArn"))
	}
	if s.RoleNameArn != nil && len(*s.RoleNameArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleNameArn", 1))
	}

	if s.SnsTopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the result of the GenerateDataSet operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/GenerateDataSetResult
type GenerateDataSetOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// A unique identifier representing a specific request to the GenerateDataSet
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation
func (s GenerateDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GenerateDataSetOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GenerateDataSetOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Container for the parameters to the StartSupportDataExport operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/StartSupportDataExportRequest
type StartSupportDataExportInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file.
	CustomerDefinedValues map[string]string `locationName:"customerDefinedValues" min:"1" type:"map"`

	// Specifies the data set type to be written to the output csv file. The data
	// set types customer_support_contacts_data and test_customer_support_contacts_data
	// both result in a csv file containing the following fields: Product Id, Product
	// Code, Customer Guid, Subscription Guid, Subscription Start Date, Organization,
	// AWS Account Id, Given Name, Surname, Telephone Number, Email, Title, Country
	// Code, ZIP Code, Operation Type, and Operation Time.
	//
	// customer_support_contacts_data Customer support contact data. The data set
	// will contain all changes (Creates, Updates, and Deletes) to customer support
	// contact data from the date specified in the from_date parameter.
	// test_customer_support_contacts_data An example data set containing static
	// test data in the same format as customer_support_contacts_data
	//
	// DataSetType is a required field
	DataSetType SupportDataSetType `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"true"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	//
	// DestinationS3BucketName is a required field
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The start date from which to retrieve the data set in UTC. This parameter
	// only affects the customer_support_contacts_data data set type.
	//
	// FromDate is a required field
	FromDate *time.Time `locationName:"fromDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	//
	// RoleNameArn is a required field
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	//
	// SnsTopicArn is a required field
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartSupportDataExportInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartSupportDataExportInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSupportDataExportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSupportDataExportInput"}
	if s.CustomerDefinedValues != nil && len(s.CustomerDefinedValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomerDefinedValues", 1))
	}
	if len(s.DataSetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DataSetType"))
	}

	if s.DestinationS3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationS3BucketName"))
	}
	if s.DestinationS3BucketName != nil && len(*s.DestinationS3BucketName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationS3BucketName", 1))
	}

	if s.FromDate == nil {
		invalidParams.Add(aws.NewErrParamRequired("FromDate"))
	}

	if s.RoleNameArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleNameArn"))
	}
	if s.RoleNameArn != nil && len(*s.RoleNameArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleNameArn", 1))
	}

	if s.SnsTopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the result of the StartSupportDataExport operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/StartSupportDataExportResult
type StartSupportDataExportOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// A unique identifier representing a specific request to the StartSupportDataExport
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation
func (s StartSupportDataExportOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartSupportDataExportOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s StartSupportDataExportOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

type DataSetType string

// Enum values for DataSetType
const (
	DataSetTypeCustomerSubscriberHourlyMonthlySubscriptions DataSetType = "customer_subscriber_hourly_monthly_subscriptions"
	DataSetTypeCustomerSubscriberAnnualSubscriptions        DataSetType = "customer_subscriber_annual_subscriptions"
	DataSetTypeDailyBusinessUsageByInstanceType             DataSetType = "daily_business_usage_by_instance_type"
	DataSetTypeDailyBusinessFees                            DataSetType = "daily_business_fees"
	DataSetTypeDailyBusinessFreeTrialConversions            DataSetType = "daily_business_free_trial_conversions"
	DataSetTypeDailyBusinessNewInstances                    DataSetType = "daily_business_new_instances"
	DataSetTypeDailyBusinessNewProductSubscribers           DataSetType = "daily_business_new_product_subscribers"
	DataSetTypeDailyBusinessCanceledProductSubscribers      DataSetType = "daily_business_canceled_product_subscribers"
	DataSetTypeMonthlyRevenueBillingAndRevenueData          DataSetType = "monthly_revenue_billing_and_revenue_data"
	DataSetTypeMonthlyRevenueAnnualSubscriptions            DataSetType = "monthly_revenue_annual_subscriptions"
	DataSetTypeDisbursedAmountByProduct                     DataSetType = "disbursed_amount_by_product"
	DataSetTypeDisbursedAmountByProductWithUncollectedFunds DataSetType = "disbursed_amount_by_product_with_uncollected_funds"
	DataSetTypeDisbursedAmountByInstanceHours               DataSetType = "disbursed_amount_by_instance_hours"
	DataSetTypeDisbursedAmountByCustomerGeo                 DataSetType = "disbursed_amount_by_customer_geo"
	DataSetTypeDisbursedAmountByAgeOfUncollectedFunds       DataSetType = "disbursed_amount_by_age_of_uncollected_funds"
	DataSetTypeDisbursedAmountByAgeOfDisbursedFunds         DataSetType = "disbursed_amount_by_age_of_disbursed_funds"
	DataSetTypeCustomerProfileByIndustry                    DataSetType = "customer_profile_by_industry"
	DataSetTypeCustomerProfileByRevenue                     DataSetType = "customer_profile_by_revenue"
	DataSetTypeCustomerProfileByGeography                   DataSetType = "customer_profile_by_geography"
	DataSetTypeSalesCompensationBilledRevenue               DataSetType = "sales_compensation_billed_revenue"
	DataSetTypeUsSalesAndUseTaxRecords                      DataSetType = "us_sales_and_use_tax_records"
)

func (enum DataSetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DataSetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SupportDataSetType string

// Enum values for SupportDataSetType
const (
	SupportDataSetTypeCustomerSupportContactsData     SupportDataSetType = "customer_support_contacts_data"
	SupportDataSetTypeTestCustomerSupportContactsData SupportDataSetType = "test_customer_support_contacts_data"
)

func (enum SupportDataSetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SupportDataSetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
