package conda_test

import (
	"testing"

	"github.com/robocorp/rcc/conda"
	"github.com/robocorp/rcc/hamlet"
)

func second(_ interface{}, version string) string {
	return version
}

func TestCanParseMicromambaVersion(t *testing.T) {
	must_be, _ := hamlet.Specifications(t)

	must_be.Equal("0", second(conda.AsVersion("python")))
	must_be.Equal("0.19.1", second(conda.AsVersion("0.19.1")))
	must_be.Equal("0.19.0", second(conda.AsVersion("micromamba: 0.19.0")))
	must_be.Equal("0.19.0", second(conda.AsVersion("\n\n\tmicromamba: 0.19.0 \nlibmamba: 0.18.7\n\n\t")))
	must_be.Equal("0.20", second(conda.AsVersion("microrumba: 0.20")))
}