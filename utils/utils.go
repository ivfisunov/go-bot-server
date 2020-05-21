package utils

import (
	"encoding/json"
	"os"
)

// ConfigLdapJSON json struct for ldap configuration
type ConfigLdapJSON struct {
	URL      string `json:"url" binding:"required"`
	BindDN   string `json:"bindDN" binding:"required"`
	Password string `json:"password" binding:"required"`
	SearchDN string `json:"searchDN" binding:"required"`
}

// LoadJSONConfig loads json config from file
func LoadJSONConfig(fileName string) (*ConfigLdapJSON, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := ConfigLdapJSON{}
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// ParseADStatus figure out the users AD status
func ParseADStatus(lockoutTime, userAccountControl, pwdLastSet string) int {
	if userAccountControl == "512" && lockoutTime == "0" {
		return 0
	} else
	// accont has blocked by administrator
	if userAccountControl == "514" {
		return 1
	} else
	// account is expired
	if pwdLastSet == "0" {
		return 2
	} else
	// account is blocked by reason of typing incorrect password multiple times (>5)
	if userAccountControl == "512" && lockoutTime != "0" && lockoutTime != "" {
		return 3
	}

	return 0
}
