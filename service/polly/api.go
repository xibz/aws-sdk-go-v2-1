// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opDeleteLexicon = "DeleteLexicon"

// DeleteLexiconRequest is a API request type for the DeleteLexicon API operation.
type DeleteLexiconRequest struct {
	*aws.Request
	Input *DeleteLexiconInput
}

// Send marshals and sends the DeleteLexicon API request.
func (r DeleteLexiconRequest) Send() (*DeleteLexiconOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DeleteLexiconOutput), nil
}

// DeleteLexiconRequest returns a request value for making API operation for
// Amazon Polly.
//
// Deletes the specified pronunciation lexicon stored in an AWS Region. A lexicon
// which has been deleted is not available for speech synthesis, nor is it possible
// to retrieve it using either the GetLexicon or ListLexicon APIs.
//
// For more information, see Managing Lexicons (http://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
//
//    // Example sending a request using the DeleteLexiconRequest method.
//    req := client.DeleteLexiconRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DeleteLexicon
func (c *Polly) DeleteLexiconRequest(input *DeleteLexiconInput) DeleteLexiconRequest {
	op := &aws.Operation{
		Name:       opDeleteLexicon,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/lexicons/{LexiconName}",
	}

	if input == nil {
		input = &DeleteLexiconInput{}
	}

	output := &DeleteLexiconOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DeleteLexiconRequest{Request: req, Input: input}
}

const opDescribeVoices = "DescribeVoices"

// DescribeVoicesRequest is a API request type for the DescribeVoices API operation.
type DescribeVoicesRequest struct {
	*aws.Request
	Input *DescribeVoicesInput
}

// Send marshals and sends the DescribeVoices API request.
func (r DescribeVoicesRequest) Send() (*DescribeVoicesOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DescribeVoicesOutput), nil
}

// DescribeVoicesRequest returns a request value for making API operation for
// Amazon Polly.
//
// Returns the list of voices that are available for use when requesting speech
// synthesis. Each voice speaks a specified language, is either male or female,
// and is identified by an ID, which is the ASCII version of the voice name.
//
// When synthesizing speech ( SynthesizeSpeech ), you provide the voice ID for
// the voice you want from the list of voices returned by DescribeVoices.
//
// For example, you want your news reader application to read news in a specific
// language, but giving a user the option to choose the voice. Using the DescribeVoices
// operation you can provide the user with a list of available voices to select
// from.
//
// You can optionally specify a language code to filter the available voices.
// For example, if you specify en-US, the operation returns a list of all available
// US English voices.
//
// This operation requires permissions to perform the polly:DescribeVoices action.
//
//    // Example sending a request using the DescribeVoicesRequest method.
//    req := client.DescribeVoicesRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DescribeVoices
func (c *Polly) DescribeVoicesRequest(input *DescribeVoicesInput) DescribeVoicesRequest {
	op := &aws.Operation{
		Name:       opDescribeVoices,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/voices",
	}

	if input == nil {
		input = &DescribeVoicesInput{}
	}

	output := &DescribeVoicesOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DescribeVoicesRequest{Request: req, Input: input}
}

const opGetLexicon = "GetLexicon"

// GetLexiconRequest is a API request type for the GetLexicon API operation.
type GetLexiconRequest struct {
	*aws.Request
	Input *GetLexiconInput
}

// Send marshals and sends the GetLexicon API request.
func (r GetLexiconRequest) Send() (*GetLexiconOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetLexiconOutput), nil
}

// GetLexiconRequest returns a request value for making API operation for
// Amazon Polly.
//
// Returns the content of the specified pronunciation lexicon stored in an AWS
// Region. For more information, see Managing Lexicons (http://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
//
//    // Example sending a request using the GetLexiconRequest method.
//    req := client.GetLexiconRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/GetLexicon
func (c *Polly) GetLexiconRequest(input *GetLexiconInput) GetLexiconRequest {
	op := &aws.Operation{
		Name:       opGetLexicon,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/lexicons/{LexiconName}",
	}

	if input == nil {
		input = &GetLexiconInput{}
	}

	output := &GetLexiconOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetLexiconRequest{Request: req, Input: input}
}

const opListLexicons = "ListLexicons"

// ListLexiconsRequest is a API request type for the ListLexicons API operation.
type ListLexiconsRequest struct {
	*aws.Request
	Input *ListLexiconsInput
}

// Send marshals and sends the ListLexicons API request.
func (r ListLexiconsRequest) Send() (*ListLexiconsOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*ListLexiconsOutput), nil
}

// ListLexiconsRequest returns a request value for making API operation for
// Amazon Polly.
//
// Returns a list of pronunciation lexicons stored in an AWS Region. For more
// information, see Managing Lexicons (http://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
//
//    // Example sending a request using the ListLexiconsRequest method.
//    req := client.ListLexiconsRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/ListLexicons
func (c *Polly) ListLexiconsRequest(input *ListLexiconsInput) ListLexiconsRequest {
	op := &aws.Operation{
		Name:       opListLexicons,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/lexicons",
	}

	if input == nil {
		input = &ListLexiconsInput{}
	}

	output := &ListLexiconsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return ListLexiconsRequest{Request: req, Input: input}
}

const opPutLexicon = "PutLexicon"

// PutLexiconRequest is a API request type for the PutLexicon API operation.
type PutLexiconRequest struct {
	*aws.Request
	Input *PutLexiconInput
}

// Send marshals and sends the PutLexicon API request.
func (r PutLexiconRequest) Send() (*PutLexiconOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*PutLexiconOutput), nil
}

// PutLexiconRequest returns a request value for making API operation for
// Amazon Polly.
//
// Stores a pronunciation lexicon in an AWS Region. If a lexicon with the same
// name already exists in the region, it is overwritten by the new lexicon.
// Lexicon operations have eventual consistency, therefore, it might take some
// time before the lexicon is available to the SynthesizeSpeech operation.
//
// For more information, see Managing Lexicons (http://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
//
//    // Example sending a request using the PutLexiconRequest method.
//    req := client.PutLexiconRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/PutLexicon
func (c *Polly) PutLexiconRequest(input *PutLexiconInput) PutLexiconRequest {
	op := &aws.Operation{
		Name:       opPutLexicon,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/lexicons/{LexiconName}",
	}

	if input == nil {
		input = &PutLexiconInput{}
	}

	output := &PutLexiconOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return PutLexiconRequest{Request: req, Input: input}
}

const opSynthesizeSpeech = "SynthesizeSpeech"

// SynthesizeSpeechRequest is a API request type for the SynthesizeSpeech API operation.
type SynthesizeSpeechRequest struct {
	*aws.Request
	Input *SynthesizeSpeechInput
}

// Send marshals and sends the SynthesizeSpeech API request.
func (r SynthesizeSpeechRequest) Send() (*SynthesizeSpeechOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*SynthesizeSpeechOutput), nil
}

// SynthesizeSpeechRequest returns a request value for making API operation for
// Amazon Polly.
//
// Synthesizes UTF-8 input, plain text or SSML, to a stream of bytes. SSML input
// must be valid, well-formed SSML. Some alphabets might not be available with
// all the voices (for example, Cyrillic might not be read at all by English
// voices) unless phoneme mapping is used. For more information, see How it
// Works (http://docs.aws.amazon.com/polly/latest/dg/how-text-to-speech-works.html).
//
//    // Example sending a request using the SynthesizeSpeechRequest method.
//    req := client.SynthesizeSpeechRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/SynthesizeSpeech
func (c *Polly) SynthesizeSpeechRequest(input *SynthesizeSpeechInput) SynthesizeSpeechRequest {
	op := &aws.Operation{
		Name:       opSynthesizeSpeech,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/speech",
	}

	if input == nil {
		input = &SynthesizeSpeechInput{}
	}

	output := &SynthesizeSpeechOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return SynthesizeSpeechRequest{Request: req, Input: input}
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DeleteLexiconInput
type DeleteLexiconInput struct {
	_ struct{} `type:"structure"`

	// The name of the lexicon to delete. Must be an existing lexicon in the region.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"LexiconName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLexiconInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteLexiconInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLexiconInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLexiconInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DeleteLexiconOutput
type DeleteLexiconOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response
}

// String returns the string representation
func (s DeleteLexiconOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteLexiconOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DeleteLexiconOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DescribeVoicesInput
type DescribeVoicesInput struct {
	_ struct{} `type:"structure"`

	// The language identification tag (ISO 639 code for the language name-ISO 3166
	// country code) for filtering the list of voices returned. If you don't specify
	// this optional parameter, all available voices are returned.
	LanguageCode LanguageCode `location:"querystring" locationName:"LanguageCode" type:"string" enum:"true"`

	// An opaque pagination token returned from the previous DescribeVoices operation.
	// If present, this indicates where to continue the listing.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeVoicesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVoicesInput) GoString() string {
	return s.String()
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/DescribeVoicesOutput
type DescribeVoicesOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The pagination token to use in the next request to continue the listing of
	// voices. NextToken is returned only if the response is truncated.
	NextToken *string `type:"string"`

	// A list of voices with their properties.
	Voices []Voice `type:"list"`
}

// String returns the string representation
func (s DescribeVoicesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVoicesOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DescribeVoicesOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/GetLexiconInput
type GetLexiconInput struct {
	_ struct{} `type:"structure"`

	// Name of the lexicon.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"LexiconName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLexiconInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetLexiconInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLexiconInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLexiconInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/GetLexiconOutput
type GetLexiconOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// Lexicon object that provides name and the string content of the lexicon.
	Lexicon *Lexicon `type:"structure"`

	// Metadata of the lexicon, including phonetic alphabetic used, language code,
	// lexicon ARN, number of lexemes defined in the lexicon, and size of lexicon
	// in bytes.
	LexiconAttributes *LexiconAttributes `type:"structure"`
}

// String returns the string representation
func (s GetLexiconOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetLexiconOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetLexiconOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Provides lexicon name and lexicon content in string format. For more information,
// see Pronunciation Lexicon Specification (PLS) Version 1.0 (https://www.w3.org/TR/pronunciation-lexicon/).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/Lexicon
type Lexicon struct {
	_ struct{} `type:"structure"`

	// Lexicon content in string format. The content of a lexicon must be in PLS
	// format.
	Content *string `type:"string"`

	// Name of the lexicon.
	Name *string `type:"string"`
}

// String returns the string representation
func (s Lexicon) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Lexicon) GoString() string {
	return s.String()
}

// Contains metadata describing the lexicon such as the number of lexemes, language
// code, and so on. For more information, see Managing Lexicons (http://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/LexiconAttributes
type LexiconAttributes struct {
	_ struct{} `type:"structure"`

	// Phonetic alphabet used in the lexicon. Valid values are ipa and x-sampa.
	Alphabet *string `type:"string"`

	// Language code that the lexicon applies to. A lexicon with a language code
	// such as "en" would be applied to all English languages (en-GB, en-US, en-AUS,
	// en-WLS, and so on.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// Date lexicon was last modified (a timestamp value).
	LastModified *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Number of lexemes in the lexicon.
	LexemesCount *int64 `type:"integer"`

	// Amazon Resource Name (ARN) of the lexicon.
	LexiconArn *string `type:"string"`

	// Total size of the lexicon, in characters.
	Size *int64 `type:"integer"`
}

// String returns the string representation
func (s LexiconAttributes) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s LexiconAttributes) GoString() string {
	return s.String()
}

// Describes the content of the lexicon.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/LexiconDescription
type LexiconDescription struct {
	_ struct{} `type:"structure"`

	// Provides lexicon metadata.
	Attributes *LexiconAttributes `type:"structure"`

	// Name of the lexicon.
	Name *string `type:"string"`
}

// String returns the string representation
func (s LexiconDescription) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s LexiconDescription) GoString() string {
	return s.String()
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/ListLexiconsInput
type ListLexiconsInput struct {
	_ struct{} `type:"structure"`

	// An opaque pagination token returned from previous ListLexicons operation.
	// If present, indicates where to continue the list of lexicons.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListLexiconsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListLexiconsInput) GoString() string {
	return s.String()
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/ListLexiconsOutput
type ListLexiconsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// A list of lexicon names and attributes.
	Lexicons []LexiconDescription `type:"list"`

	// The pagination token to use in the next request to continue the listing of
	// lexicons. NextToken is returned only if the response is truncated.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListLexiconsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListLexiconsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s ListLexiconsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/PutLexiconInput
type PutLexiconInput struct {
	_ struct{} `type:"structure"`

	// Content of the PLS lexicon as string data.
	//
	// Content is a required field
	Content *string `type:"string" required:"true"`

	// Name of the lexicon. The name must follow the regular express format [0-9A-Za-z]{1,20}.
	// That is, the name is a case-sensitive alphanumeric string up to 20 characters
	// long.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"LexiconName" type:"string" required:"true"`
}

// String returns the string representation
func (s PutLexiconInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutLexiconInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLexiconInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLexiconInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/PutLexiconOutput
type PutLexiconOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response
}

// String returns the string representation
func (s PutLexiconOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutLexiconOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s PutLexiconOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/SynthesizeSpeechInput
type SynthesizeSpeechInput struct {
	_ struct{} `type:"structure"`

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon
	// is the same as the language of the voice. For information about storing lexicons,
	// see PutLexicon (http://docs.aws.amazon.com/polly/latest/dg/API_PutLexicon.html).
	LexiconNames []string `type:"list"`

	// The format in which the returned output will be encoded. For audio stream,
	// this will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json.
	//
	// OutputFormat is a required field
	OutputFormat OutputFormat `type:"string" required:"true" enum:"true"`

	// The audio frequency specified in Hz.
	//
	// The valid values for mp3 and ogg_vorbis are "8000", "16000", and "22050".
	// The default value is "22050".
	//
	// Valid values for pcm are "8000" and "16000" The default value is "16000".
	SampleRate *string `type:"string"`

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []SpeechMarkType `type:"list"`

	// Input text to synthesize. If you specify ssml as the TextType, follow the
	// SSML format for the input text.
	//
	// Text is a required field
	Text *string `type:"string" required:"true"`

	// Specifies whether the input text is plain text or SSML. The default value
	// is plain text. For more information, see Using SSML (http://docs.aws.amazon.com/polly/latest/dg/ssml.html).
	TextType TextType `type:"string" enum:"true"`

	// Voice ID to use for the synthesis. You can get a list of available voice
	// IDs by calling the DescribeVoices (http://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html)
	// operation.
	//
	// VoiceId is a required field
	VoiceId VoiceId `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s SynthesizeSpeechInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SynthesizeSpeechInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SynthesizeSpeechInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SynthesizeSpeechInput"}
	if len(s.OutputFormat) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("OutputFormat"))
	}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if len(s.VoiceId) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("VoiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/SynthesizeSpeechOutput
type SynthesizeSpeechOutput struct {
	_ struct{} `type:"structure" payload:"AudioStream"`

	responseMetadata aws.Response

	// Stream containing the synthesized speech.
	AudioStream io.ReadCloser `type:"blob"`

	// Specifies the type audio stream. This should reflect the OutputFormat parameter
	// in your request.
	//
	//    *  If you request mp3 as the OutputFormat, the ContentType returned is
	//    audio/mpeg.
	//
	//    *  If you request ogg_vorbis as the OutputFormat, the ContentType returned
	//    is audio/ogg.
	//
	//    *  If you request pcm as the OutputFormat, the ContentType returned is
	//    audio/pcm in a signed 16-bit, 1 channel (mono), little-endian format.
	//
	//
	//    * If you request json as the OutputFormat, the ContentType returned is
	//    audio/json.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Number of characters synthesized.
	RequestCharacters *int64 `location:"header" locationName:"x-amzn-RequestCharacters" type:"integer"`
}

// String returns the string representation
func (s SynthesizeSpeechOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SynthesizeSpeechOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s SynthesizeSpeechOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Description of the voice.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/Voice
type Voice struct {
	_ struct{} `type:"structure"`

	// Gender of the voice.
	Gender Gender `type:"string" enum:"true"`

	// Amazon Polly assigned voice ID. This is the ID that you specify when calling
	// the SynthesizeSpeech operation.
	Id VoiceId `type:"string" enum:"true"`

	// Language code of the voice.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// Human readable name of the language in English.
	LanguageName *string `type:"string"`

	// Name of the voice (for example, Salli, Kendra, etc.). This provides a human
	// readable voice name that you might display in your application.
	Name *string `type:"string"`
}

// String returns the string representation
func (s Voice) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Voice) GoString() string {
	return s.String()
}

type Gender string

// Enum values for Gender
const (
	GenderFemale Gender = "Female"
	GenderMale   Gender = "Male"
)

type LanguageCode string

// Enum values for LanguageCode
const (
	LanguageCodeCyGb    LanguageCode = "cy-GB"
	LanguageCodeDaDk    LanguageCode = "da-DK"
	LanguageCodeDeDe    LanguageCode = "de-DE"
	LanguageCodeEnAu    LanguageCode = "en-AU"
	LanguageCodeEnGb    LanguageCode = "en-GB"
	LanguageCodeEnGbWls LanguageCode = "en-GB-WLS"
	LanguageCodeEnIn    LanguageCode = "en-IN"
	LanguageCodeEnUs    LanguageCode = "en-US"
	LanguageCodeEsEs    LanguageCode = "es-ES"
	LanguageCodeEsUs    LanguageCode = "es-US"
	LanguageCodeFrCa    LanguageCode = "fr-CA"
	LanguageCodeFrFr    LanguageCode = "fr-FR"
	LanguageCodeIsIs    LanguageCode = "is-IS"
	LanguageCodeItIt    LanguageCode = "it-IT"
	LanguageCodeKoKr    LanguageCode = "ko-KR"
	LanguageCodeJaJp    LanguageCode = "ja-JP"
	LanguageCodeNbNo    LanguageCode = "nb-NO"
	LanguageCodeNlNl    LanguageCode = "nl-NL"
	LanguageCodePlPl    LanguageCode = "pl-PL"
	LanguageCodePtBr    LanguageCode = "pt-BR"
	LanguageCodePtPt    LanguageCode = "pt-PT"
	LanguageCodeRoRo    LanguageCode = "ro-RO"
	LanguageCodeRuRu    LanguageCode = "ru-RU"
	LanguageCodeSvSe    LanguageCode = "sv-SE"
	LanguageCodeTrTr    LanguageCode = "tr-TR"
)

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatJson      OutputFormat = "json"
	OutputFormatMp3       OutputFormat = "mp3"
	OutputFormatOggVorbis OutputFormat = "ogg_vorbis"
	OutputFormatPcm       OutputFormat = "pcm"
)

type SpeechMarkType string

// Enum values for SpeechMarkType
const (
	SpeechMarkTypeSentence SpeechMarkType = "sentence"
	SpeechMarkTypeSsml     SpeechMarkType = "ssml"
	SpeechMarkTypeViseme   SpeechMarkType = "viseme"
	SpeechMarkTypeWord     SpeechMarkType = "word"
)

type TextType string

// Enum values for TextType
const (
	TextTypeSsml TextType = "ssml"
	TextTypeText TextType = "text"
)

type VoiceId string

// Enum values for VoiceId
const (
	VoiceIdGeraint   VoiceId = "Geraint"
	VoiceIdGwyneth   VoiceId = "Gwyneth"
	VoiceIdMads      VoiceId = "Mads"
	VoiceIdNaja      VoiceId = "Naja"
	VoiceIdHans      VoiceId = "Hans"
	VoiceIdMarlene   VoiceId = "Marlene"
	VoiceIdNicole    VoiceId = "Nicole"
	VoiceIdRussell   VoiceId = "Russell"
	VoiceIdAmy       VoiceId = "Amy"
	VoiceIdBrian     VoiceId = "Brian"
	VoiceIdEmma      VoiceId = "Emma"
	VoiceIdRaveena   VoiceId = "Raveena"
	VoiceIdIvy       VoiceId = "Ivy"
	VoiceIdJoanna    VoiceId = "Joanna"
	VoiceIdJoey      VoiceId = "Joey"
	VoiceIdJustin    VoiceId = "Justin"
	VoiceIdKendra    VoiceId = "Kendra"
	VoiceIdKimberly  VoiceId = "Kimberly"
	VoiceIdMatthew   VoiceId = "Matthew"
	VoiceIdSalli     VoiceId = "Salli"
	VoiceIdConchita  VoiceId = "Conchita"
	VoiceIdEnrique   VoiceId = "Enrique"
	VoiceIdMiguel    VoiceId = "Miguel"
	VoiceIdPenelope  VoiceId = "Penelope"
	VoiceIdChantal   VoiceId = "Chantal"
	VoiceIdCeline    VoiceId = "Celine"
	VoiceIdMathieu   VoiceId = "Mathieu"
	VoiceIdDora      VoiceId = "Dora"
	VoiceIdKarl      VoiceId = "Karl"
	VoiceIdCarla     VoiceId = "Carla"
	VoiceIdGiorgio   VoiceId = "Giorgio"
	VoiceIdMizuki    VoiceId = "Mizuki"
	VoiceIdLiv       VoiceId = "Liv"
	VoiceIdLotte     VoiceId = "Lotte"
	VoiceIdRuben     VoiceId = "Ruben"
	VoiceIdEwa       VoiceId = "Ewa"
	VoiceIdJacek     VoiceId = "Jacek"
	VoiceIdJan       VoiceId = "Jan"
	VoiceIdMaja      VoiceId = "Maja"
	VoiceIdRicardo   VoiceId = "Ricardo"
	VoiceIdVitoria   VoiceId = "Vitoria"
	VoiceIdCristiano VoiceId = "Cristiano"
	VoiceIdInes      VoiceId = "Ines"
	VoiceIdCarmen    VoiceId = "Carmen"
	VoiceIdMaxim     VoiceId = "Maxim"
	VoiceIdTatyana   VoiceId = "Tatyana"
	VoiceIdAstrid    VoiceId = "Astrid"
	VoiceIdFiliz     VoiceId = "Filiz"
	VoiceIdVicki     VoiceId = "Vicki"
	VoiceIdTakumi    VoiceId = "Takumi"
	VoiceIdSeoyeon   VoiceId = "Seoyeon"
	VoiceIdAditi     VoiceId = "Aditi"
)
