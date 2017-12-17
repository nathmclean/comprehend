package comprehend

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/aws"
)

type BatchSentimentsResponse struct {
	Sentiments []BatchSentiment
	Errors []BatchError
}

type BatchSentiment struct {
	Sentiment Sentiment
	Index int64
}

type BatchEntitiesResponse struct {
	BatchEntities []BatchEntity
	Errors []BatchError
}

type BatchEntity struct {
	Entities []Entity
	Index int64
}

type BatchKeyPhrasesResponse struct {
	BatchKeyPhrases []BatchKeyPhrases
	Errors []BatchError
}

type BatchKeyPhrases struct {
	KeyPhrase []KeyPhrase
	Index int64
}

type BatchLanguageResponse struct {
	BatchLanguage []BatchLanguage
	Errors []BatchError
}

type BatchLanguage struct {
	Langs []Lang
	Index int64
}

type BatchError struct {
	Index int64
	ErrorMessage string
}

func (c* ComprehendClient) GetSentimentBatch(texts []string) (BatchSentimentsResponse, error) {
	var sentiments []BatchSentiment
	var batchErrors []BatchError

	batchResponse := BatchSentimentsResponse{
		Sentiments: sentiments,
		Errors:     batchErrors,
	}

	resp, err := c.Client.BatchDetectSentiment(&comprehend.BatchDetectSentimentInput{
		LanguageCode: aws.String(c.Language),
		TextList:     aws.StringSlice(texts),
	})
	if err != nil {
		return batchResponse, err
	}

	for _, respError := range resp.ErrorList {
		batchError := BatchError{
			Index:        *respError.Index,
			ErrorMessage: *respError.ErrorMessage,
		}
		batchErrors = append(batchErrors, batchError)
	}
	for _, respSentiment := range resp.ResultList {
		var sentiment Sentiment
		sentiment.SentimentClass = *respSentiment.Sentiment
		sentiment.Score.Mixed =  *respSentiment.SentimentScore.Mixed
		sentiment.Score.Negative =  *respSentiment.SentimentScore.Negative
		sentiment.Score.Positive =  *respSentiment.SentimentScore.Positive
		sentiment.Score.Neutral =  *respSentiment.SentimentScore.Neutral

		batchSentiment := BatchSentiment{
			Sentiment: sentiment,
			Index:     *respSentiment.Index,
		}
		sentiments = append(sentiments, batchSentiment)
	}

	return batchResponse, nil
}

func (c* ComprehendClient) GetEntitiesBatch(texts []string) (BatchEntitiesResponse, error) {
	var batchEntities []BatchEntity
	var batchErrors []BatchError

	batchResponse := BatchEntitiesResponse{
		BatchEntities: batchEntities,
		Errors:        batchErrors,
	}

	resp, err := c.Client.BatchDetectEntities(&comprehend.BatchDetectEntitiesInput{
		LanguageCode: aws.String(c.Language),
		TextList:     aws.StringSlice(texts),
	})
	if err != nil {
		return batchResponse, err
	}

	for _, respError := range resp.ErrorList {
		batchError := BatchError{
			Index:        *respError.Index,
			ErrorMessage: *respError.ErrorMessage,
		}
		batchErrors = append(batchErrors, batchError)
	}

	for _, respEntities := range resp.ResultList {
		var entities []Entity
		for _, entity := range respEntities.Entities {
			entities = append(entities, Entity{
				Text:       *entity.Text,
				EntityType: *entity.Type,
				Score:      *entity.Score,
			})
		}
		var batchEntity BatchEntity
		batchEntity.Index = *respEntities.Index
		batchEntity.Entities = entities
		batchEntities = append(batchEntities, batchEntity)
	}

	return batchResponse, nil
}

func (c *ComprehendClient) GetKeyPhrasesBatch(texts []string) (BatchKeyPhrasesResponse, error) {
	var batchKeyPhrases []BatchKeyPhrases
	var batchErrors []BatchError

	batchResponse := BatchKeyPhrasesResponse{
		BatchKeyPhrases: batchKeyPhrases,
		Errors:          batchErrors,
	}

	resp, err := c.Client.BatchDetectKeyPhrases(&comprehend.BatchDetectKeyPhrasesInput{
		LanguageCode: aws.String(c.Language),
		TextList:     aws.StringSlice(texts),
	})
	if err != nil {
		return batchResponse, err
	}

	for _, respError := range resp.ErrorList {
		batchError := BatchError{
			Index:        *respError.Index,
			ErrorMessage: *respError.ErrorMessage,
		}
		batchErrors = append(batchErrors, batchError)
	}

	for _, respKeyPhrases := range resp.ResultList {
		var keyPhrases []KeyPhrase
		for _, keyPhrase := range respKeyPhrases.KeyPhrases {
			keyPhrases = append(keyPhrases, KeyPhrase{
				Phrase: *keyPhrase.Text,
				Score:  *keyPhrase.Score,
			})
			batchKeyPhrases = append(batchKeyPhrases, BatchKeyPhrases{
				KeyPhrase: keyPhrases,
				Index:     *respKeyPhrases.Index,
			})
		}
	}

	return batchResponse, nil
}

func (c *ComprehendClient) GetLanguageBatch(texts []string) (BatchLanguageResponse, error) {
	var batchLanguages []BatchLanguage
	var batchErrors []BatchError

	batchResponse := BatchLanguageResponse{
		BatchLanguage: batchLanguages,
		Errors:        batchErrors,
	}

	resp, err := c.Client.BatchDetectDominantLanguage(&comprehend.BatchDetectDominantLanguageInput{
		TextList: aws.StringSlice(texts),
	})
	if err != nil {
		return batchResponse, err
	}

	for _, respError := range resp.ErrorList {
		batchError := BatchError{
			Index:        *respError.Index,
			ErrorMessage: *respError.ErrorMessage,
		}
		batchErrors = append(batchErrors, batchError)
	}

	for _, respLangResp := range resp.ResultList {
		var langs []Lang
		for _, respLang := range respLangResp.Languages {
			langs = append(langs, Lang{
				Score:    *respLang.Score,
				LangCode: *respLang.LanguageCode,
			})
			batchLanguages = append(batchLanguages, BatchLanguage{
				Langs: langs,
				Index: *respLangResp.Index,
			})
		}
	}

	return batchResponse, nil
}