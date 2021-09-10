package congo

import (
	"os"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	config := New()
	config.data["hostname"] = "blrkm101"

	got := config.Get("hostname")
	var want interface{} = "blrkm101"

	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func TestStringType(t *testing.T) {
	config := New()
	config.data["hostname"] = "blrkm101"

	got := config.GetString("hostname")
	want := "blrkm101"

	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func TestLoadFromReader(t *testing.T) {
	data := `{
"hostname": "blrkm101"
}`
	reader := strings.NewReader(data)
	config := New()
	config.load(reader)

	got := config.GetString("hostname")
	want := "blrkm101"

	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func TestLoadFromReaderOverride(t *testing.T) {
	config := New()
	config.load(strings.NewReader(`{"hostname": "blrkm101", "ring": "ring-1"}`))
	config.load(strings.NewReader(`{"hostname": "blrkm101-dev"}`))

	got := config.GetString("hostname")
	want := "blrkm101-dev"

	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func TestLoadFromReaderNonOverride(t *testing.T) {
	config := New()
	config.load(strings.NewReader(`{"hostname": "blrkm101", "ring": "ring-1"}`))
	config.load(strings.NewReader(`{"hostname": "blrkm101-dev"}`))

	got := config.GetString("hostname")
	want := "blrkm101-dev"
	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}

	got = config.GetString("ring")
	want = "ring-1"
	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func TestConfigFile(t *testing.T) {
	def, cleanup := createFile(t, `{"hostname": "blrkm101", "ring": "ring-1"}`)
	defer cleanup()
	config := New()
	err := config.LoadFile(def.Name())
	if err != nil {
		t.Fatal(err.Error())
	}

	got := config.GetString("ring")
	want := "ring-1"
	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func TestConfigFileOverride(t *testing.T) {
	def, f1 := createFile(t, `{"hostname": "blrkm101", "ring": "ring-1"}`)
	dev, f2 := createFile(t, `{"hostname": "blrkm101-dev"}`)
	defer func() {
		f1()
		f2()
	}()

	config := New()
	err := config.LoadFile(def.Name())
	if err != nil {
		t.Fatal(err.Error())
	}
	err = config.LoadFile(dev.Name())
	if err != nil {
		t.Fatal(err.Error())
	}

	got := config.GetString("ring")
	want := "ring-1"
	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}

	got = config.GetString("hostname")
	want = "blrkm101-dev"
	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func createFile(t *testing.T, data string) (*os.File, func()) {
	t.Helper()

	file, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = os.WriteFile(file.Name(), []byte(data), 0644)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return file, func() {
		os.Remove(file.Name())
	}
}
