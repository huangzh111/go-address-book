package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSaveWalletError(t *testing.T) {
	path := t.TempDir()
	wallets := []Wallet{
		{Address: "0x123", Name: "主钱包"},
	}
	err := saveWallet(path, wallets)

	if err == nil {
		t.Error("预期保存失败，但没有返回错误")
	}

}
func TestSaveWallet(t *testing.T) {
	want := []Wallet{
		{Address: "0x123", Name: "主钱包"},
	}

	path := filepath.Join(t.TempDir(), "test.json")

	err := saveWallet(path, want)
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}

	var got []Wallet
	err = json.Unmarshal(data, &got)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("实际结果：%v,预期结果：%v", got, want)
	}

}
func TestValidAddress(t *testing.T) {

	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{"合法地址", "0x1111111111111111111111111111111111111111", true},
		{"地址太短", "0x123", false},
		{"包含非法字符", "0xgggggggggggggggggggggggggggggggggggggggg", false},
		{"缺少0x前缀", "1111111111111111111111111111111111111111", false},
		{"前缀错误", "ab1111111111111111111111111111111111111111", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidAddress(tc.address)

			if got != tc.want {
				t.Errorf("实际结果: %v,预期结果: %v ", got, tc.want)
			}
		})
	}
}
