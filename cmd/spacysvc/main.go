// Package main provides a simple CLI wrapper around the spacy-service-client
// which communicates with an instance of spacy-service. You can pass some
// text as an argument, and the response from SpaCy is presented as JSON.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"scout.ac/go/spacy-service-client/spacysvc"
	pb "scout.ac/go/spacy-service-client/spacysvc/generated"

	"google.golang.org/protobuf/encoding/protojson"
)

var (
	addr          = flag.String("addr", "localhost:50051", "the address of the spacy_service server")
	text          = flag.String("text", "", "text to process")
	skipTokenize  = flag.Bool("skip-tokenize", false, "skip tokenize")
	skipAncestors = flag.Bool("skip-ancestors", false, "skip ancestors")
	skipChildren  = flag.Bool("skip-children", false, "skip children")
	skipConjuncts = flag.Bool("skip-conjuncts", false, "skip conjuncts")
	skipLefts     = flag.Bool("skip-lefts", false, "skip lefts")
	skipRights    = flag.Bool("skip-rights", false, "skip rights")
	skipSentiment = flag.Bool("skip-sentiment", false, "skip sentiment")
	skipEnts      = flag.Bool("skip-ents", false, "skip ents")
	skipSents     = flag.Bool("skip-sents", false, "skip sents")
	skipSpans     = flag.Bool("skip-spans", false, "skip spans")
	skipCoref     = flag.Bool("skip-coref", false, "skip coreference resolution")
	skipSubtree   = flag.Bool("skip-subtree", false, "skip subtree")
)

func main() {
	flag.Parse()
	if *text == "" {
		os.Exit(0)
	}

	client, err := spacysvc.New(*addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	request := pb.GetDocRequest{
		Text:          *text,
		Tokenize:      !*skipTokenize,
		Ancestors:     !*skipAncestors,
		Children:      !*skipChildren,
		Conjuncts:     !*skipConjuncts,
		Lefts:         !*skipLefts,
		Rights:        !*skipRights,
		SkipSentiment: *skipSentiment,
		SkipEnts:      *skipEnts,
		SkipSents:     *skipSents,
		SkipSpans:     *skipSpans,
		SkipCoref:     *skipCoref,
		Subtree:       !*skipSubtree,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	response, err := client.GetDoc(ctx, &request)
	if err != nil {
		panic(err)
	}

	doc, err := protojson.Marshal(&pb.Doc{
		Ents:        response.Ents,
		Sents:       response.Sents,
		Sentiment:   response.Sentiment,
		Spans:       response.Spans,
		Tokens:      response.Tokens,
		CorefChains: response.CorefChains,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(doc))
}
