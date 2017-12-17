package comprehend

import (
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

type ComprehendClient struct {
	Client comprehendiface.ComprehendAPI
	Language string
}

func NewClient(lang string) ComprehendClient {
	client := ComprehendClient{
		Language: lang,
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := comprehend.New(sess)

	client.Client = svc

	return client
}