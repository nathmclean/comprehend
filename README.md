# Comprehend

A Go library for interacting with the AW Comprehend service.

## Install

`go get github.com/nathmclean/comprehend`

## Usage

```
import "github.com/nathmclean/appStore/comprehend"

// This will use credentials within your AWS credential chain
comprehendClient := comprehend.NewClient()

// Submit a string for Sentiment analysis
txt = "some test text"
sentiment, err := comprehendClient.GetSentiment(txt)

// get the overall sentiment
fmt.Println("Sentiment:" sentiment.SentimentClass)

// get the value (0-1) for a specific Sentiment Class (POSITIVE, NEGATIVE, MIXED, NEUTRAL)
fmt.Println("Mixed:", sentiment.Score.Mixed)
fmt.Println("Positive:", sentiment.Score.Positive)
fmt.Println("Negative:", sentiment.Score.Negative)
fmt.Println("Neutral:", sentiment.Score.Neutral)


// submit a string to recieve entities
entities, err := comprehendClient.GetEntities(txt)
if err != nil {
	fmt.Println(err)
}
fmt.Println(entities)



// submit a string for phrase analysis
phrases, _ := comprehendClient.GetKeyPhrases(comment)
if err != nil {
	log.Println(err)
}
log.Println(phrases)
```