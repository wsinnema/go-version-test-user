package libtest

import "github.com/wsinnema/go-version-test/lib"
import "testing"

func TestVersion(t *testing.T) {
	lib.PackageVersion()
	t.Error("Test fails")
}
