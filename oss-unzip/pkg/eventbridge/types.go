package eventbridge

const (
	ExtensionsKeyAliyunAccountID         = "aliyunaccountid"
	ExtensionsKeyAliyunEventbusName      = "aliyuneventbusname"
	ExtensionsKeyAliyunOriginalAccountID = "aliyunoriginalaccountid"
	ExtensionsKeyAliyunPublishTime       = "aliyunpublishtime"
	ExtensionsKeyAliyunRegionID          = "aliyunregionid"
	OSSSourceName                        = "acs:oss"
)

type OSSEventSpec struct {
	Region            string
	EventVersion      string
	EventSource       string
	EventName         string
	EventTime         string
	RequestParameters *RequestParameters
	UserIdentity      *UserIdentity
	ResponseElements  *ResponseElements
	Oss               *OSS
}

type RequestParameters struct {
	SourceIPAddress string
}

type UserIdentity struct {
	PrincipalId string
}

type ResponseElements struct {
	RequestId string
}

type OSS struct {
	Bucket           *OSSBucket
	OssSchemaVersion string
	Object           *OSSObject
}

type OSSBucket struct {
	Name          string
	Arn           string
	OwnerIdentity string
}

type OSSObject struct {
	Size      int64
	DeltaSize int64
	ETag      string
	Key       string
}
