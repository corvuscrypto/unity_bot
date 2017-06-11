package config

import (
  "os"
  "os/exec"
  "testing"
  )

// Test the basic defaults
func TestDefaults(T *testing.T) {
  LoadConfig("")
	// Realistically this is all we need to test here
	if GlobalConfig.BotName != DefaultBotName {
		T.Errorf("Default not set correctly. Got %s, expected %s", GlobalConfig.BotName, DefaultBotName)
	}
}

// Test loading in a .yml file. in the file, only BotName is set so therefore the other values should
// remain untouched. Note that this will not show up on code coverage. But that's fine.
func TestFileLoad(T *testing.T) {
	// Create a tmp file.
	file, err := os.OpenFile("test.yml", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		T.Log("Unable to perform this test because the temporary file could not be created.")
		T.SkipNow()
	}
  file.Write([]byte("botName: Rick"))
  file.Close()
  LoadConfig("test.yml")
  if GlobalConfig.BotName != "Rick" || GlobalConfig.UIPort != DefaultUIPort {
    T.Errorf("Loading of file specifications did not work. \n %+v", GlobalConfig)
  }
  os.Remove("test.yml")
}

func TestConfigErrors(T *testing.T) {
  if os.Getenv("UnityGoTest") == "1" {
    LoadConfig("./config.go")
    return
  }
  if os.Getenv("UnityGoTest") == "2" {
    LoadConfig("test")
    return
  }
  if os.Getenv("UnityGoTest") == "3" {
    file, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0600)
    file.Write([]byte("asd--as--=a= sd--= ")) // Corrupt yml data
    file.Close()
    LoadConfig("test")
    return
  }

  cmd := exec.Command(os.Args[0], "-test.run=TestConfigErrors")
  cmd.Env = append(os.Environ(), "UnityGoTest=1")
  err := cmd.Run()
  if e, ok := err.(*exec.ExitError); !ok || e.Success() {
    T.Error("Program did not exit as expected.")
  }

  cmd = exec.Command(os.Args[0], "-test.run=TestConfigErrors")
  cmd.Env = append(os.Environ(), "UnityGoTest=2")
  err = cmd.Run()
  if e, ok := err.(*exec.ExitError); !ok || e.Success() {
    T.Error("Program did not exit as expected.")
  }

  cmd = exec.Command(os.Args[0], "-test.run=TestConfigErrors")
  cmd.Env = append(os.Environ(), "UnityGoTest=3")
  err = cmd.Run()
  if e, ok := err.(*exec.ExitError); !ok || e.Success() {
    T.Error("Program did not exit as expected.")
  }
  os.Remove("test")
}
