package certificate

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Certificate struct {
  Id               uint64 `yaml:"id"`
  Name             string `yaml:"name"`
  ContactGate      string `yaml:"contact_gate"`
  RegistrationDate uint64 `yaml:"registration_date"`
}

func LoadAbilityCertificate() (*Certificate, error) {
  yamlFile, err := os.ReadFile("ability/certificate.yml")
  if err != nil {
    return nil, err
  }

  var certi *Certificate

  err = yaml.Unmarshal(yamlFile, certi)
  if err != nil {
    return nil, err
  }

  return certi, nil
}
