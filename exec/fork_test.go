package exec

import (
	"testing"

	"github.com/juhovuori/builder/project"
)

type mockBuild struct{}

func (b *mockBuild) ID() string               { return "" }
func (b *mockBuild) Project() project.Project { return nil }
func (b *mockBuild) Completed() bool          { return false }
func (b *mockBuild) Error() error             { return nil }

func TestForkExecutor(t *testing.T) {
	dir := ""
	b := mockBuild{}
	e := forkExecutor{dir, &b}
	e.Run()

	e.Cleanup()
}
