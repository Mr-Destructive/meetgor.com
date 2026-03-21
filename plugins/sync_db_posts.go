package plugins

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Mr-Destructive/meetgor.com/models"
)

type SyncDBPostsPlugin struct {
	PluginName string
}

var syncDBExecCommand = exec.Command

func (p *SyncDBPostsPlugin) Name() string {
	return p.PluginName
}

func (p *SyncDBPostsPlugin) Phase() Phase {
	return PhaseRead
}

func (p *SyncDBPostsPlugin) Requires() []string {
	return nil
}

func (p *SyncDBPostsPlugin) AdminPolicy() AdminPolicy {
	return AdminSkip
}

func (p *SyncDBPostsPlugin) Execute(ssg *models.SSG) error {
	if ssg.Config.AdminMode {
		return nil
	}

	if os.Getenv("SYNC_DB_ON_BUILD") != "1" {
		return nil
	}

	args := []string{"run", "cmd/sync_db/main.go"}
	if ssg.Config.Blog.PostsDir != "" {
		args = append(args, "--output-dir", ssg.Config.Blog.PostsDir)
	}
	if os.Getenv("SYNC_DB_SYNC_ALL") == "1" {
		args = append(args, "--sync-all")
	}
	if os.Getenv("SYNC_DB_DRY_RUN") == "1" {
		args = append(args, "--dry-run")
	}
	if os.Getenv("SYNC_DB_VERBOSE") == "1" {
		args = append(args, "--verbose")
	}
	if hours := os.Getenv("SYNC_DB_HOURS"); hours != "" {
		args = append(args, "--hours", hours)
	}

	cmd := syncDBExecCommand("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	fmt.Println("SyncDB: running go", args)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("sync db failed: %w", err)
	}
	return nil
}

func init() {
	RegisterPlugin("SyncDB", func() Plugin {
		return &SyncDBPostsPlugin{PluginName: "SyncDB"}
	})
}
