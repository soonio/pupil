package migrate

import (
	"errors"
	"fmt"
	"github.com/soonio/pupil/pkg/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/soonio/pupil/app"

	v4 "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
)

//迁移文件所在相对目录
const dir = "database"

func Version(cli *cli.Context) error {
	version, dirty, err := driver().Version()
	if err == nil {
		fmt.Printf("version: %d, dirty: %v\n", version, dirty)
	}
	return err
}
func Create(cli *cli.Context) error {
	wrap := fmt.Sprintf("%s/%s", app.Home, dir)

	matches, _ := filepath.Glob(filepath.Join(dir, "*sql"))

	version, _ := nextSeqVersion(matches)

	var name string
	if len(os.Args) > 2 {
		name = os.Args[2]
	} else {
		return errors.New("请填入数据表名参数")
	}

	for _, direction := range []string{"up", "down"} {
		err := utils.Touch(fmt.Sprintf("%s/%s_%s.%s.sql", wrap, version, name, direction))
		if err != nil {
			return err
		}
	}
	return nil
}
func Up(cli *cli.Context) error {
	err := driver().Up()
	return err
}

func Down(cli *cli.Context) error {

	var confirm string
	fmt.Print("请确认是否删除所有数据表数据[yes/no]: ")
	if _, err := fmt.Scanf("%s", &confirm); err != nil {
		return err
	}

	if confirm == "yes" {
		err := driver().Down()
		return err
	}
	fmt.Println("确认失败，不进行回滚")
	return nil
}

func Steps(cli *cli.Context) error {
	step := cli.Int("number")

	var confirm string
	fmt.Print("请确认回退迁移[yes/no]: ")
	if _, err := fmt.Scanf("%s", &confirm); err != nil {
		return err
	}

	if confirm == "yes" {
		err := driver().Steps(step)
		return err
	}
	fmt.Println("确认失败，不进行任何迁移")
	return nil
}
func driver() *v4.Migrate {
	m, err := v4.New(fmt.Sprintf("file://%s/%s", app.Home, dir), "mysql://"+app.Config.DB.Dsn())
	if err != nil {
		panic(err)
	}
	m.Log = new(Log)
	return m
}

type Log struct{}

// Printf 打印信息
func (l Log) Printf(format string, v ...any) {
	fmt.Printf(format, v...)
}

// Verbose 是否输出详细的调试信息
func (l Log) Verbose() bool {
	return false
}

func nextSeqVersion(matches []string) (string, error) {
	nextSeq := uint64(1)

	if len(matches) > 0 {
		filename := matches[len(matches)-1]
		matchSeqStr := filepath.Base(filename)
		idx := strings.Index(matchSeqStr, "_")

		if idx < 1 { // Using 1 instead of 0 since there should be at least 1 digit
			return "", fmt.Errorf("Malformed migration filename: %s ", filename)
		}

		var err error
		matchSeqStr = matchSeqStr[0:idx]
		nextSeq, err = strconv.ParseUint(matchSeqStr, 10, 64)

		if err != nil {
			return "", err
		}

		nextSeq++
	}

	version := fmt.Sprintf("%0[2]*[1]d", nextSeq, 6)
	if len(version) > 6 {
		return "", fmt.Errorf("Next sequence number %s too large. At most %d digits are allowed ", version, 6)
	}
	return version, nil
}
