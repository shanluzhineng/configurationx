package configurationx

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func discoverFileFromPath(path string, fileExt []string) ([]string, error) {
	if _, err := os.Stat(path); err != nil {
		// is not exist
		return []string{}, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config: Open failed on %s. %s", path, err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("config: Stat failed on %s. %s", path, err)
	}

	if !fi.IsDir() {
		if !pathIsFileExtension(path, fileExt) {
			return []string{}, nil
		}
		return []string{path}, nil
	}

	var sources []string
	fis, err := f.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("config: Readdir failed on %s. %s", path, err)
	}

	// sort files by name
	sort.Sort(byName(fis))

	for _, fi := range fis {
		fp := filepath.Join(path, fi.Name())
		// check for a symlink and resolve the path
		if fi.Mode()&os.ModeSymlink > 0 {
			var err error
			fp, err = filepath.EvalSymlinks(fp)
			if err != nil {
				return nil, err
			}
			fi, err = os.Stat(fp)
			if err != nil {
				return nil, err
			}
		}
		// do not recurse into sub dirs
		if fi.IsDir() {
			continue
		}

		if !pathIsFileExtension(fp, fileExt) {
			continue
		}
		sources = append(sources, fp)
	}
	return sources, nil
}

func pathIsFileExtension(path string, extList []string) bool {
	for _, eachExt := range extList {
		result := strings.HasSuffix(path, eachExt)
		if result {
			return true
		}
	}
	return false
}

type byName []os.FileInfo

// #region byName sort.Interface members

func (a byName) Len() int {
	return len(a)
}
func (a byName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a byName) Less(i, j int) bool {
	return a[i].Name() < a[j].Name()
}

// #endregion
