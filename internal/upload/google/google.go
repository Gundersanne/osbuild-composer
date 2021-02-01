package google

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

//creds, err := google.CredentialsFromJSON
//func CredentialsFromJSON(ctx context.Context, jsonData []byte, scopes ...string) (*Credentials, error)

type Google struct {
	client *storage.Client
}

// Credentials as described in https://pkg.go.dev/golang.org/x/oauth2/google#CredentialsFromJSON
func New(creds []byte) (*Google, error) {
	// Session credentials
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(creds))
	if err != nil {
		return nil, err
	}

	return &Google{
		client,
	}, nil
}

func (g *Google) Upload(source, dest, bucket string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	ctx := context.Background()

	b := g.client.Bucket(bucket)
	// // Will throw an error if the bucket doesn't exist
	// _, err = b.Attrs(ctx)
	// if err != nil {
	// 	return err
	// }

	writer := b.Object(dest).NewWriter(ctx)
	defer writer.Close()
	// writer.ContentType = "octetstream or something?"

	// TODO include checksum
	log.Printf("[Google] 🚀 Uploading image to Google cloud: %s", bucket)
	_, err = io.Copy(writer, sourceFile)
	if err != nil {
		return err
	}

	// If share with someone, do it here
	// err = b.Object(dest).ACL().Set(ctx,, storage.RoleReader)

	return nil
}

// func (g *Google) Share() error {

// }
