package comprehend

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/aws"
)

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

type Entity struct {
	Text string
	EntityType string
	Score float64
}

type KeyPhrase struct {
	Phrase string
	Score float64
}

type Lang struct {
	Score float64
	LangCode string
}

func (c* ComprehendClient) GetSentiment(text string) (Sentiment, error) {
	var sentiment Sentiment

	resp, err := c.Client.DetectSentiment(&comprehend.DetectSentimentInput{
		LanguageCode: aws.String(c.Language),
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

func (c* ComprehendClient) GetEntities(text string) ([]Entity, error) {
	var entities []Entity

	resp, err := c.Client.DetectEntities(&comprehend.DetectEntitiesInput{
		LanguageCode: aws.String(c.Language),
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

func (c* ComprehendClient) GetKeyPhrases(text string) ([]KeyPhrase, error) {
	var keyPhrases []KeyPhrase


	resp, err := c.Client.DetectKeyPhrases(&comprehend.DetectKeyPhrasesInput{
		LanguageCode: aws.String(c.Language),
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

func (c *ComprehendClient) GetDominantLanguage(text string) ([]Lang, error) {
	var langs []Lang

	resp, err := c.Client.DetectDominantLanguage(&comprehend.DetectDominantLanguageInput{
		Text: aws.String(text),
	})
	if err != nil {
		return langs, err
	}

	for _, lang := range resp.Languages {
		langs = append(langs, Lang{
			Score:    *lang.Score,
			LangCode: *lang.LanguageCode,
		})
	}

	return langs, nil
}