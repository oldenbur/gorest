package githubapi

import (
	"testing"

	log "github.com/cihub/seelog"
	T "github.com/oldenbur/sql-parser/testutil"
	. "github.com/smartystreets/goconvey/convey"
)

func init() { T.ConfigureTestLogger() }

func TestFuncCall(t *testing.T) {

	defer log.Flush()

	Convey("Test list config files\n", t, func() {
		c := NewLRGithubConfig()
		_, err := c.ConfigFilesForBranch("DataIndexer-7.1.8")
		So(err, ShouldBeNil)
	})
}
