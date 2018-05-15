package loader

import (
	"os"
	"os/exec"
	"testing"
)

func SetTimeStamp() {
	exec.Command("sh", "../tests/set_timestamp_song.sh").Run()
}

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	SetTimeStamp()

	e := m.Run()

	teardown()

	os.Exit(e)
}
