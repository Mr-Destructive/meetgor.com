package plugins

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func createFileInGitHub(repoOwner, repoName, filePath, fileContent, commitMessage, token string) error {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	encodedContent := base64.StdEncoding.EncodeToString([]byte(fileContent))

	_, _, err := client.Repositories.CreateFile(ctx, repoOwner, repoName, filePath, &github.RepositoryContentFileOptions{
		Message: &commitMessage,
		Content: []byte(encodedContent),
	})

	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	return nil
}
