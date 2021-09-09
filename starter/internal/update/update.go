package update

import (
	"context"
	"log"
	"runtime/debug"
	"strings"

	"github.com/aerogear/charmil/cli/pkg/common/modname"
	"github.com/aerogear/charmil/core/utils/color"
	"github.com/aerogear/charmil/core/utils/localize"
	"github.com/aerogear/charmil/core/utils/logging"
	"github.com/google/go-github/github"
)

var (
	modName string
	Version = "dev"
	Owner   string
	Repo    string
)

func init() {
	// Extracts the module name from `go.mod` file and stores it
	modName, err := modname.GetModuleName()
	if err != nil {
		log.Fatal(err)
	}

	splitModName := strings.Split(modName, "/")
	if len(splitModName) < 3 {
		log.Fatal("Invalid module name format")
	}

	Owner = splitModName[1]
	Repo = splitModName[2]

	if isDevBuild() {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}

// CheckForUpdate checks if there is a newer version of the CLI than
// the version currently being used. If so, it logs this information
// to the console.
func CheckForUpdate(ctx context.Context, logger logging.Logger, localizer localize.Localizer) {
	releases, err := getReleases(ctx)
	if err != nil {
		return
	}

	var latestRelease *github.RepositoryRelease
	releaseTagIndexMap := map[string]int{}
	for i, release := range releases {
		// assign the latest non-pre release as the latest public release
		if latestRelease == nil && !release.GetPrerelease() {
			latestRelease = release
		}

		// create an tag:index map of the releases
		// the first index (0) is the latest release
		releaseTagIndexMap[release.GetTagName()] = i
		if release.GetTagName() == Version {
			break
		}
	}

	currentVersionIndex, ok := releaseTagIndexMap[Version]
	if !ok {
		// the currently used version does not exist as a public release
		// assume it to be an unpublished or dev release
		return
	}

	latestVersionIndex := releaseTagIndexMap[latestRelease.GetTagName()]

	// if the index of the current version is greater than the latest release
	// this means it is older, and therefore, an update is available.
	if currentVersionIndex > latestVersionIndex {
		logger.Infoln()
		logger.Infoln(color.Info("Update available!"), color.CodeSnippet(latestRelease.GetTagName()))
		logger.Infoln(color.Info(latestRelease.GetHTMLURL()))
		logger.Infoln()
	}
}

func getReleases(ctx context.Context) ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)

	releases, _, err := client.Repositories.ListReleases(ctx, Owner, Repo, nil)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

// isDevBuild returns true if the current build is "dev" (dev build)
func isDevBuild() bool {
	return Version == "dev"
}
