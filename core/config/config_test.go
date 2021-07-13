package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

var (
	fPath = "./testdata/test_config.json"
	fExt  = ".json"
	cfg   = struct {
		Key1    string
		Key2    string
		Key3    string
		Key4    string
		Plugins map[string]interface{}
	}{
		Key1: "val1",
		Key2: "val2",
		Key3: "val3",
	}

	cHandler = &CfgHandler{
		cfg:      &cfg,
		filePath: fPath,
		fileExt:  fExt,
	}
)

func TestNewHandler(t *testing.T) {
	h := NewHandler(fPath, cfg)

	if !reflect.DeepEqual(h.cfg, cfg) {
		t.Errorf("Test failed: config structs do not match. Expected: %+v, Received: %+v", cfg, h.cfg)
	}

	if h.filePath != fPath {
		t.Errorf("Test failed: file paths do not match. Expected: %v, Received: %v", fPath, h.filePath)
	}

	if h.fileExt != fExt {
		t.Errorf("Test failed: config file extensions do not match. Expected: %v, Received: %v", fExt, h.fileExt)
	}
}

func TestMarshal(t *testing.T) {
	actualBuf, err := Marshal(cfg, fExt)
	if err != nil {
		t.Error(err)
	}

	// Limited to JSON files as of now. Plan to extend to other formats later
	expectedBuf, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(actualBuf, expectedBuf) {
		t.Errorf("Test failed: outputs do not match. Expected: %v, Received: %v", string(expectedBuf), string(actualBuf))
	}
}

func TestUnmarshal(t *testing.T) {
	buf := []byte(`{"Key1":"val1","Key2":"val2","Key3":"newVal3"}`)

	err := Unmarshal(buf, &cfg, fExt)
	if err != nil {
		t.Error(err)
	}

	if cfg.Key1 != "val1" || cfg.Key2 != "val2" || cfg.Key3 != "newVal3" || cfg.Key4 != "" {
		t.Errorf("Test failed: unmarshal unsuccessful")
	}
}

func TestReadFile(t *testing.T) {
	actualBuf, err := readFile("./testdata/testReadFile.txt")
	if err != nil {
		t.Error(err)
	}

	if string(actualBuf) != "Test Read Data" {
		t.Errorf("Test failed: read unsuccessful")
	}
}

func TestWriteFile(t *testing.T) {
	writeFilePath := "./testdata/testWriteFile.txt"
	expectedBuf := []byte("Test Write Data")

	err := writeFile(writeFilePath, expectedBuf)
	if err != nil {
		t.Error(err)
	}

	actualBuf, err := ioutil.ReadFile(writeFilePath)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(actualBuf, expectedBuf) {
		t.Errorf("Test failed: contents do not match. Expected: %v, Received: %v", string(expectedBuf), string(actualBuf))
	}
}

func TestLoad(t *testing.T) {
	cHandler.filePath = "./testdata/load_config.json"

	err := cHandler.Load()
	if err != nil {
		t.Error(err)
	}

	if cfg.Key4 != "newVal4" {
		t.Errorf("Test failed: load unsuccessful")
	}
}

func TestSave(t *testing.T) {
	cHandler.filePath = "./testdata/save_config.json"

	expectedBuf := []byte(`{
  "Key1": "val1",
  "Key2": "val2",
  "Key3": "newVal3",
  "Key4": "newVal4",
  "Plugins": null
}`)

	err := cHandler.Save()
	if err != nil {
		t.Error(err)
	}

	actualBuf, err := readFile(cHandler.filePath)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(actualBuf, expectedBuf) {
		t.Errorf("Test failed: contents do not match. Expected: %v, Received: %v", string(expectedBuf), string(actualBuf))
	}
}

func TestMergePluginCfg(t *testing.T) {
	pluginCfg := struct {
		Key5 string
		Key6 string
		Key7 string
	}{"val5", "val6", "val7"}

	expectedVal := `{
  "Key1": "val1",
  "Key2": "val2",
  "Key3": "newVal3",
  "Key4": "newVal4",
  "Plugins": {
    "testPlugin": {
      "Key5": "val5",
      "Key6": "val6",
      "Key7": "val7"
    }
  }
}`

	err := MergePluginCfg("testPlugin", cHandler, pluginCfg)
	if err != nil {
		t.Error(err)
	}

	actualBuf, err := json.MarshalIndent(cHandler.cfg, "", "  ")
	if err != nil {
		t.Error(err)
	}

	if string(actualBuf) != expectedVal {
		t.Errorf("Test failed: merge unsuccessful. Expected: %v, Received: %v", expectedVal, string(actualBuf))
	}
}
