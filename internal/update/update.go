package update

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/inconshreveable/go-update"
	"golang.org/x/mod/semver"
)

// Version contains the current version (injected at compile)
var Version string

var repoURL string = "https://repo.simse.io/qc"
var manifestURL string = repoURL + "/manifest.json"

type manifest struct {
	LatestRelease release   `json:"latest_release"`
	Releases      []release `json:"releases"`
}

type release struct {
	Version  string                                  `json:"version"`
	Binaries map[string]map[string]binary            `json:"binaries"`
	Patches  map[string]map[string]map[string]string `json:"patches"`
}

type binary struct {
	Path     string `json:"path"`
	Checksum string `json:"checksum"`
}

// ErrLatestVersion is an error returned when qc is already up to date
var ErrLatestVersion error = errors.New("Already one latest version")

// Do performs an update of qc (in-place no less!)
func Do() error {
	if Version == "" {
		return nil
	}

	// Fetch manifest
	resp, err := http.Get(manifestURL)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	qcManifest := manifest{}
	json.Unmarshal(body, &qcManifest)

	if semver.Compare(qcManifest.LatestRelease.Version, Version) == 1 {
		// Attempt to find patch
		patchURL, patchURLExists := qcManifest.LatestRelease.Patches[runtime.GOOS][runtime.GOARCH][Version]

		if patchURLExists {
			// fmt.Println("patching binary")
			patch, err := http.Get(patchURL)
			if err != nil {
				// handle error
			}
			defer patch.Body.Close()

			patchErr := update.Apply(patch.Body, update.Options{
				Patcher: update.NewBSDiffPatcher(),
			})
			if patchErr != nil {
				// error handling
			}

			return patchErr
		}

		// fmt.Println("replacing binary")
		// Patch did not exist, fall back to full binary replacement
		resp, err := http.Get(qcManifest.LatestRelease.Binaries[runtime.GOOS][runtime.GOARCH].Path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		binaryUpdateError := update.Apply(resp.Body, update.Options{})
		if binaryUpdateError != nil {
			if rerr := update.RollbackError(binaryUpdateError); rerr != nil {
				fmt.Println("Failed to rollback from bad update")
			}
		}

		return binaryUpdateError
	}

	return ErrLatestVersion
	// return nil
}
