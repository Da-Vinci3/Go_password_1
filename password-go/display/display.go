package display

import (
	"fmt"
	"generatepass/utils"
	"strings"
)

type DisplayFormat struct {
	BorderChar    string
	BorderColor   string
	TitleColor    string
	SiteColor     string
	PasswordColor string
	ShowBorder    bool
	ShowIndex     bool
}

// Define a type for the decrypt function
type DecryptFunc func([]byte, string) (string, error)

func DefaultFormat() DisplayFormat {
	return DisplayFormat{
		BorderChar:    "=",
		BorderColor:   utils.Blue,
		TitleColor:    utils.Bold + utils.Purple,
		SiteColor:     utils.Green,
		PasswordColor: utils.Cyan,
		ShowBorder:    true,
		ShowIndex:     true,
	}
}

func DisplayPasswords(passwords []map[string]string, key []byte, decryptFn DecryptFunc, format DisplayFormat) {
	if len(passwords) == 0 {
		fmt.Println("\nNo passwords stored.")
		return
	}

	borderLine := strings.Repeat(format.BorderChar, 40)

	fmt.Printf("\n%s%sStored Passwords%s\n", format.TitleColor, utils.Bold, utils.Reset)

	if format.ShowBorder {
		fmt.Printf("%s%s%s\n", format.BorderColor, borderLine, utils.Reset)
	}

	for i, entry := range passwords {
		if format.ShowIndex {
			fmt.Printf("\n%sEntry %d:%s\n", utils.Bold, i+1, utils.Reset)
		}

		for encSite, encPass := range entry {
			site, err := decryptFn(key, encSite)
			if err != nil {
				fmt.Printf("Error decrypting site for entry %d: %v\n", i+1, err)
				continue
			}

			pass, err := decryptFn(key, encPass)
			if err != nil {
				fmt.Printf("Error decrypting password for entry %d: %v\n", i+1, err)
				continue
			}

			fmt.Printf("%sSite     :%s %s\n", format.SiteColor, utils.Reset, site)
			fmt.Printf("%sPassword :%s %s\n", format.PasswordColor, utils.Reset, pass)
		}

		if format.ShowBorder {
			fmt.Printf("%s%s%s\n", format.BorderColor, borderLine, utils.Reset)
		}
	}

	fmt.Printf("\n%sTotal passwords stored: %d%s\n", utils.Bold, len(passwords), utils.Reset)
}

func MinimalFormat() DisplayFormat {
	format := DefaultFormat()
	format.ShowBorder = false
	format.ShowIndex = false
	return format
}

func CompactFormat() DisplayFormat {
	format := DefaultFormat()
	format.BorderChar = "-"
	return format
}

func ColorfulFormat() DisplayFormat {
	format := DefaultFormat()
	format.BorderChar = "*"
	format.BorderColor = utils.Purple
	format.SiteColor = utils.Yellow
	format.PasswordColor = utils.Green
	return format
}
