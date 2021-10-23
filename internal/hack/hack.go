package hack

import (
	"os"
	"path"
)

type Hack struct {
	ID            string
	BlurFilepaths []string
	OrigFilepaths []string
}

var hacksDir = "assets/img/hacks"

func Parse() ([]Hack, error) {
	hacks := []Hack{}

	entries, err := os.ReadDir(hacksDir)
	if err != nil {

	}

	for _, entry := range entries {
		hack := Hack{}

		if entry.IsDir() {
			hack.ID = entry.Name()

			dirEntries, err := os.ReadDir(path.Join(hacksDir, hack.ID))
			if err != nil {
				return nil, err
			}
			for _, dentry := range dirEntries {
				if dentry.Name() == "blur" {
					photos, err := os.ReadDir(path.Join(hacksDir, hack.ID, dentry.Name()))
					if err != nil {
						return nil, err
					}
					for _, p := range photos {
						hack.BlurFilepaths = append(
							hack.BlurFilepaths,
							path.Join(hacksDir, hack.ID, dentry.Name(), p.Name()))
					}
				} else if dentry.Name() == "orig" {
					photos, err := os.ReadDir(path.Join(hacksDir, hack.ID, dentry.Name()))
					if err != nil {
						return nil, err
					}
					for _, p := range photos {
						hack.OrigFilepaths = append(
							hack.OrigFilepaths,
							path.Join(hacksDir, hack.ID, dentry.Name(), p.Name()))
					}
				}
			}
			hacks = append(hacks, hack)
		}
	}
	return hacks, nil
}
