package app

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gopkg.in/ini.v1"
)

type Config struct {
	path       string
	cfg        *ini.File
	hmacSecret []byte
}

type addressRecord struct {
	Address string `json:"address"`
	HMAC    string `json:"hmac"`
}

func init() {
}

func NewConfig() (*Config, error) {
	path, err := getConfigPathFor("config")
	if err != nil {
		return nil, err
	}

	ini.DefaultHeader = true
	cfg, err := ini.Load(path)
	if err != nil {
		cfg = ini.Empty()
	}

	return &Config{
		path:       path,
		cfg:        cfg,
		hmacSecret: loadOrGenerateSalt(),
	}, nil
}

func (c *Config) SaveAddress(ticker, network, address string) error {
	section := c.cfg.Section(sectionName(ticker, network))

	if addressExists(section, address) {
		return nil
	}

	record := addressRecord{
		Address: address,
		HMAC:    c.generateHMAC(address),
	}

	jsonData, err := json.Marshal(record)
	if err != nil {
		return err
	}
	section.Key(currentTimestamp()).SetValue(string(jsonData))
	return c.save()
}

func (c *Config) GetLatestAddress(ticker, network string) string {
	section := c.cfg.Section(sectionName(ticker, network))
	keys := sortKeysByTimestamp(section.Keys())
	if len(keys) == 0 {
		return ""
	}

	var record addressRecord
	if err := json.Unmarshal([]byte(keys[0].Value()), &record); err != nil {
		return ""
	}

	if !c.validateHMAC(record.Address, record.HMAC) {
		return "invalid address"
	}

	return record.Address
}

func (c *Config) GetAllAddress(ticker, network string) []string {
	section := c.cfg.Section(sectionName(ticker, network))
	keys := sortKeysByTimestamp(section.Keys())
	if len(keys) == 0 {
		return []string{}
	}
	var addrList []string
	for _, key := range keys {
		var record addressRecord
		if err := json.Unmarshal([]byte(key.Value()), &record); err != nil {
			continue
		}
		if !c.validateHMAC(record.Address, record.HMAC) {
			continue
		}
		addrList = append(addrList, record.Address)
	}
	return addrList
}

func (c *Config) DeleteAddress(ticker, network, address string) {
	section := c.cfg.Section(sectionName(ticker, network))
	for _, key := range section.Keys() {
		var record addressRecord
		if err := json.Unmarshal([]byte(key.Value()), &record); err == nil {
			if record.Address == address {
				section.DeleteKey(key.Name())
			}
		}
	}
	_ = c.save()
}

func loadOrGenerateSalt() []byte {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("failed to get config directory: " + err.Error())
	}
	appConfigDir := filepath.Join(configDir, "swap")
	saltFile := filepath.Join(appConfigDir, "secret.bin")
	saltFile = filepath.Clean(saltFile)

	// #nosec G304 -- saltFile is constructed from os.UserConfigDir and hardcoded strings, not user input
	if salt, err := os.ReadFile(saltFile); err == nil {
		return salt
	}

	// Generate new 32-byte (256-bit) salt
	newSalt := make([]byte, 32)
	if _, err := rand.Read(newSalt); err != nil {
		panic("failed to generate secure salt: " + err.Error())
	}

	// Ensure config directory exists
	if err := os.MkdirAll(appConfigDir, 0700); err != nil {
		panic("failed to create config directory: " + err.Error())
	}

	// Save salt with restrictive permissions
	if err := os.WriteFile(saltFile, newSalt, 0600); err != nil {
		panic("failed to save salt file: " + err.Error())
	}

	return newSalt
}

func (c *Config) generateHMAC(address string) string {
	h := hmac.New(sha256.New, c.hmacSecret)
	h.Write([]byte(address))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Config) validateHMAC(address, receivedHMAC string) bool {
	expectedMAC := c.generateHMAC(address)
	return hmac.Equal([]byte(expectedMAC), []byte(receivedHMAC))
}

func (c *Config) save() error {
	dir := filepath.Dir(c.path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return c.cfg.SaveTo(c.path)
}

func sectionName(ticker, network string) string {
	return strings.ToLower(ticker) + "-" + strings.ToLower(network)
}

func sortKeysByTimestamp(keys []*ini.Key) []*ini.Key {
	sort.Slice(keys, func(i, j int) bool {
		t1, _ := time.Parse("20060102150405", keys[i].Name())
		t2, _ := time.Parse("20060102150405", keys[j].Name())
		return t1.After(t2)
	})
	return keys
}

func addressExists(section *ini.Section, address string) bool {
	for _, key := range section.Keys() {
		var record addressRecord
		if err := json.Unmarshal([]byte(key.Value()), &record); err == nil {
			if record.Address == address {
				return true
			}
		}
	}
	return false
}

func currentTimestamp() string {
	return time.Now().Format("20060102150405")
}

func getConfigPathFor(file string) (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "swap", file), nil
}
