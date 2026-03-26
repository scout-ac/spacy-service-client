# spacy-service-client

Client for [spacy-service](https://scout.ac/go/spacy-service/), a GRPC server wrapped around SpaCy.


## Install the CLI

```bash
go install scout.ac/go/spacy-service-client/cmd/spacysvc@latest
```

Start your `spacy-service`, then you can send text to it and get JSON back:

```bash
spacysvc --text "Hello, world!" | jq .
{
  "sents": [...],
  "tokens": [
    {
      "endChar": 5,
      "tag": "UH",
      "lemma": "hello",
      "dep": "ROOT",
      "isAlpha": true,
      "isAscii": true,
      "isSentStart": true,
      "isTitle": true,
      "lang": "en",
      "lower": "hello",
      "ancestors": {},
      "children": {
        "i": [...]
      },
      "lefts": {},
      "rights": {
        "i": [...]
      },
      "conjuncts": {},
      "subtree": {
        "i": [...]
      },
      "text": "Hello"
    },
	{...etc...}
  ]
}
```


## Use as a Go library

```go
package main

import (
	"context"
	"fmt"

	"scout.ac/go/spacy-service-client/spacysvc"              // the client
	pb "scout.ac/go/spacy-service-client/spacysvc/generated" // the protobufs
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client, err := spacysvc.New("localhost:50051")
	check(err)
	defer client.Close()

	ctx := context.Background()

	req := pb.GetDocRequest{
		Text: "Hello, world!",
		Tokenize:  true,
		Ancestors: true,
		Children:  true,
		Conjuncts: true,
		Lefts:     true,
		Rights:    true,
		Subtree:   true,
	}
	doc, err := client.GetDoc(ctx, &req)
	check(err)

	// The following are all available:
	// doc.Doc
	// doc.Ents
	// doc.Sentiment
	// doc.Sents
	// doc.Spans
	// doc.Text
	// doc.Token

	// Or you can simply get a JSON byte slice:
	j, err := doc.ToJSON()
	check(err)
	fmt.Println(string(j))
}
```


## Contributing

See [CONTRIBUTING.md](https://github.com/scout-ac/spacy-service/blob/main/CONTRIBUTING.md)
