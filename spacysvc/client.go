// Package spacysvc defines the client.
package spacysvc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	pb "scout.ac/go/spacy-service-client/spacysvc/generated"
)

// Doc wraps pb.Doc and provides a ToJSON method.
type Doc struct{ *pb.Doc }

// ToJSON serialises a Doc to a JSON byte slice.
func (d *Doc) ToJSON() ([]byte, error) {
	return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(d.Doc)
}

// Client holds a live connection to the spacy gRPC service.
type Client struct {
	conn *grpc.ClientConn
	svc  pb.SpacyServiceClient
}

// New connects to addr (e.g. "localhost:50051") and returns a Client.
// The caller is responsible for calling Close when done.
func New(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
		svc:  pb.NewSpacyServiceClient(conn),
	}, nil
}

// Close releases the underlying gRPC connection.
func (c *Client) Close() error {
	return c.conn.Close()
}

// GetDoc sends a request to the spacy service. Timeout is the caller's
// responsibility via ctx.
func (c *Client) GetDoc(ctx context.Context, req *pb.GetDocRequest) (Doc, error) {
	doc, err := c.svc.GetDoc(ctx, req)
	if err != nil {
		return Doc{}, err
	}
	return Doc{doc}, nil
}
