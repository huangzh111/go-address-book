package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func saveWallet(path string, wallets []Wallet) error {
	data, err := json.MarshalIndent(wallets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

type Wallet struct {
	Address string
	Name    string
}

func isValidAddress(address string) bool {
	if len(address) != 42 || !strings.HasPrefix(address, "0x") {

		return false
	}

	for _, ch := range address[2:] {
		if !strings.ContainsRune("0123456789abcdefABCDEF", ch) {

			return false
		}
	}

	return true
}
func main() {
	wallets := make([]Wallet, 0) //目前没有元素，容量为0，自动扩容

	data, err := os.ReadFile("addresses.json")
	if err == nil {
		err = json.Unmarshal(data, &wallets)
		if err != nil {
			fmt.Println("JSON解析失败 :", err)
			return
		}
	} else if !os.IsNotExist(err) {
		fmt.Println("文件读取失败 :", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("请输入命令")
		return
	}

	switch os.Args[1] {

	case "add":
		if len(os.Args) < 4 {
			fmt.Println("请输入正确指令")
			return
		}
		address := os.Args[2]
		if !isValidAddress(address) {
			fmt.Println("钱包地址包含非法字符")
			return
		}

		for _, wallet := range wallets {
			if strings.EqualFold(wallet.Address, address) {
				fmt.Println("该钱包地址已存在")
				return
			}
		}
		wallets = append(wallets, Wallet{
			Address: os.Args[2],
			Name:    os.Args[3],
		})
		err := saveWallet("addresses.json", wallets)
		if err != nil {
			fmt.Println("保存失败 : ", err)
			return
		}
		fmt.Println("钱包保存成功")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("请输入要删除的钱包地址")
			return
		}
		for i, wallet := range wallets {
			if strings.EqualFold(wallet.Address, os.Args[2]) {
				wallets = append(wallets[:i], wallets[i+1:]...)

				err := saveWallet("addresses.json", wallets)
				if err != nil {
					fmt.Println("删除失败 : ", err)
					return
				}

				fmt.Println("删除成功：", wallets)
				return
			}
		}
		fmt.Println("没有找到该钱包")

	case "find":
		if len(os.Args) < 3 {
			fmt.Println("请输入要查找的钱包地址")
			return
		}

		for _, wallet := range wallets {
			if strings.EqualFold(wallet.Address, os.Args[2]) {
				fmt.Println("找到了:", wallet.Address, wallet.Name)
				return
			}
		}
		fmt.Println("没有找到该钱包")
	case "list":

		fmt.Println(wallets)

	default:
		fmt.Println("未知命令")
	}

}
