package spacysvc_test

import (
	"context"
	"testing"

	"scout.ac/go/spacy-service-client/spacysvc"
	pb "scout.ac/go/spacy-service-client/spacysvc/generated"
)

func TestNew(t *testing.T) {
	client, err := spacysvc.New("localhost:50051")
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("Close() returned error: %v", err)
	}
}

func TestGetDoc(t *testing.T) {
	client, err := spacysvc.New("localhost:50051")
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}
	defer client.Close()

	doc, err := client.GetDoc(context.Background(), &pb.GetDocRequest{
		Text:          "The cat sat on the mat.",
		Tokenize:      true,
		Ancestors:     true,
		Children:      true,
		Conjuncts:     true,
		Lefts:         true,
		Rights:        true,
		SkipSentiment: false,
		SkipEnts:      false,
		SkipSents:     false,
		SkipSpans:     false,
		Subtree:       true,
	})
	if err != nil {
		t.Fatalf("GetDoc() returned error: %v", err)
	}

	j, err := doc.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() returned error: %v", err)
	}
	if len(j) == 0 {
		t.Fatal("ToJSON() returned empty output")
	}
}
