package comprehend

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

type ComprehendClient struct {
	Client comprehendiface.ComprehendAPI
}

type Sentiment struct {
	SentimentClass string
	Score SentinmentScore
}

type SentinmentScore struct {
	Mixed float64
	Negative float64
	Positive float64
	Neutral float64
}

func NewClient() ComprehendClient {
	client := ComprehendClient{}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := comprehend.New(sess)

	client.Client = svc

	return client
}

func (c* ComprehendClient) GetSentiment(text string) (Sentiment, error) {
	var sentiment Sentiment

	resp, err := c.Client.DetectSentiment(&comprehend.DetectSentimentInput{
		LanguageCode: aws.String("en"),
		Text:         aws.String(text),
	})

	if err != nil {
		return sentiment, err
	}

	sentiment.SentimentClass = *resp.Sentiment
	sentiment.Score.Mixed =  *resp.SentimentScore.Mixed
	sentiment.Score.Negative =  *resp.SentimentScore.Negative
	sentiment.Score.Positive =  *resp.SentimentScore.Positive
	sentiment.Score.Neutral =  *resp.SentimentScore.Neutral

	return sentiment, nil
}

type Entity struct {
	Text string
	EntityType string
	Score float64
}

func (c* ComprehendClient) GetEntities(text string) ([]Entity, error) {
	var entities []Entity

	resp, err := c.Client.DetectEntities(&comprehend.DetectEntitiesInput{
		LanguageCode: aws.String("en"),
		Text:         aws.String(text),
	})
	if err != nil {
		return entities, err
	}

	for _, entity := range resp.Entities {
		entities = append(entities, Entity{
			Text:       *entity.Text,
			EntityType: *entity.Type,
			Score:      *entity.Score,
		})
	}

	return entities, nil
}

type KeyPhrase struct {
	Phrase string
	Score float64
}

func (c* ComprehendClient) GetKeyPhrases(text string) ([]KeyPhrase, error) {
	var keyPhrases []KeyPhrase


	resp, err := c.Client.DetectKeyPhrases(&comprehend.DetectKeyPhrasesInput{
		LanguageCode: aws.String("en"),
		Text:         aws.String(text),
	})

	if err != nil {
		return keyPhrases, err
	}

	for _, phrase := range resp.KeyPhrases {
		keyPhrases = append(keyPhrases, KeyPhrase{
			Phrase: *phrase.Text,
			Score:  *phrase.Score,
		})
	}

	return keyPhrases, nil
}