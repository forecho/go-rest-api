package path

import (
	"fmt"
	"github.com/google/martian/log"
	"os"
	"path/filepath"
	"strings"
)

func RootPath() string {
	var (
		err error
	)

	defer func() {
		if err != nil {
			panic(fmt.Sprintf("GetProjectRoot error :%+v", err))
		}
	}()

	pwd, err := os.Getwd()
	for i := 0; i < 10; i = i + 1 {
		if strings.HasSuffix(pwd, "go-rest-api") || strings.HasSuffix(pwd, "srv") {
			log.Infof("\nfind project root path: %s\n", pwd)
			return pwd
		}

		pwd = filepath.Dir(pwd)
	}

	panic(fmt.Sprintf("GetProjectRoot error :%+v", err))
}

func RootPathWithPostfix(p string) string {
	d := RootPath() + "/" + p
	return d
}

func StoragePath() string {
	return RootPath() + "/storage"
}

func StoragePathWithPostfix(p string) string {
	d := StoragePath() + "/" + p
	return d
}
