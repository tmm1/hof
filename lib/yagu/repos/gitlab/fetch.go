package gitlab

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-billy/v5"
	"github.com/xanzy/go-gitlab"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/lib/yagu/repos/git"
)

func Fetch(FS billy.Filesystem, owner, repo, tag string, private bool) (error) {
	if private && os.Getenv(TokenEnv) == "" {
		fmt.Println("gitlab git fallback")
		return git.Fetch(FS, "gitlab.com", owner, repo, tag, private)
	}

	client, err := NewClient(private)
	if err != nil {
		return err
	}

	pid := path.Join(owner, repo)

	var sha string

	if tag == "v0.0.0" {
		bs, _, err := client.Branches.ListBranches(pid, nil)
		if err != nil {
			return  err
		}

		var branch *gitlab.Branch

		for _, candidate := range bs {
			if candidate.Default {
				branch = candidate

				break
			}
		}

		if branch == nil {
			return fmt.Errorf("Could not find default branch for repository %s", pid)
		}

		sha = branch.Commit.ID
	} else {
		t, _, err := client.Tags.GetTag(pid, tag)
		if err != nil {
			return err
		}

		sha = t.Commit.ID
	}

	zReader, err := fetchShaZip(client, pid, sha)
	if err != nil {
		return fmt.Errorf("While fetching from GitLab\n%w\n", err)
	}

	if err := yagu.BillyLoadFromZip(zReader, FS, true); err != nil {
		return fmt.Errorf("While reading zipfile\n%w\n", err)
	}

	return nil
}

func fetchShaZip(client *gitlab.Client, pid, sha string) (*zip.Reader, error) {
	format := "zip"
	data, _, err := client.Repositories.Archive(pid, &gitlab.ArchiveOptions{
		Format: &format,
		SHA:    &sha,
	})
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(data)

	return zip.NewReader(r, int64(len(data)))
}

