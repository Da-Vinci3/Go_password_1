package main

import (
	"fmt"
	"os"
	"strings"

	"generatepass/counter"
	"generatepass/display"
	"generatepass/encrypt"
	"generatepass/key"
	"generatepass/password"
	"generatepass/storage"
	"generatepass/utils"
)

// Global variables for commonly used values
var (
	secretkey    string
	passfilename = "passwords.json"
)

func main() {
	if err := handleSecretKey(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for {
		displayMenu()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1: // View all passwords
			if err := viewPasswords(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case 2: // Add new password
			newData, err := addNewPassword()
			if err != nil {
				fmt.Printf("Error creating password: %v\n", err)
				continue
			}

			if err := saveNewPassword(newData); err != nil {
				fmt.Printf("Error saving password: %v\n", err)
			}

		case 3: // Add new and view all
			newData, err := addNewPassword()
			if err != nil {
				fmt.Printf("Error creating password: %v\n", err)
				continue
			}

			if err := saveNewPassword(newData); err != nil {
				fmt.Printf("Error saving password: %v\n", err)
				continue
			}

			if err := viewPasswords(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case 4: // Update Password
			if err := updatePassword(); err != nil {
				fmt.Printf("Error updating password: %v\n", err)
			}

		case 5: //Delete Password
			if err := deletePassword(); err != nil {
				fmt.Printf("Error deleting password: %v\n", err)
			}

		case 6:
			if err := searchPasswords(); err != nil {
				fmt.Printf("Error searching passwords: %v\n", err)
			}

		case 7: // Exit
			fmt.Printf("%s\nGoodbye!%s%s%s", utils.Bold, utils.Red, utils.Italic, utils.Reset)
			os.Exit(0)

		default:
			fmt.Printf("%s\nInvalid choice. Please try again.%s%s%s", utils.Bold, utils.Red, utils.Italic, utils.Reset)
		}

		// Add a pause before showing the menu again
		fmt.Printf("%s\nPress Enter to continue...%s%s%s", utils.Bold, utils.Red, utils.Italic, utils.Reset)
		fmt.Scanln()
	}
}

func displayMenu() {
	fmt.Printf("\n%s%sPassword Manager Menu%s\n", utils.Bold, utils.Blue, utils.Reset)
	fmt.Println("1. View all passwords")
	fmt.Println("2. Add new password")
	fmt.Println("3. Add new password and view all")
	fmt.Println("4. Update existing password")
	fmt.Println("5. Delete existing password")
	fmt.Println("6. Search existing password")
	fmt.Println("7. Exit")
	fmt.Print("\nEnter your choice (1-7): ")
}

func handleSecretKey() error {
	filename := "counter.txt"
	currentNum, err := counter.ReadAndIncrementCounter(filename)
	if err != nil {
		return fmt.Errorf("error with counter: %v", err)
	}

	if currentNum == 1 {
		fmt.Print("\nYour unique key is: ")
		secretkey = key.GenerateKey(32)
		fmt.Println(secretkey)
		fmt.Println("\nPlease save this key securely! You'll need it to access your passwords.")
		fmt.Print("Press Enter to continue...")
		fmt.Scanln() // Wait for user to acknowledge
	} else {
		fmt.Print("\nEnter your secret key: ")
		fmt.Scanln(&secretkey)
	}
	return nil
}

func addNewPassword() (map[string]string, error) {
	var site string
	fmt.Print("\nWhat site would you like to make a password for: ")
	fmt.Scanln(&site)

	result := password.GeneratePass(14)

	pass1, err := encrypt.EncryptAES([]byte(secretkey), result)
	if err != nil {
		return nil, err
	}

	site1, err := encrypt.EncryptAES([]byte(secretkey), site)
	if err != nil {
		return nil, err
	}

	newData := map[string]string{
		site1: pass1,
	}

	// Display the newly created password
	fmt.Printf("\n%sNewly Created Password:%s\n", utils.Bold, utils.Reset)
	fmt.Printf("%sSite     :%s %s\n", utils.Green, utils.Reset, site)
	fmt.Printf("%sPassword :%s %s\n", utils.Green, utils.Reset, result)

	return newData, nil
}

func viewPasswords() error {
	existingData, err := storage.ReadPasswordFile(passfilename)
	if err != nil {
		return fmt.Errorf("error reading passwords: %v", err)
	}

	format := display.ColorfulFormat()
	display.DisplayPasswords(existingData, []byte(secretkey), encrypt.DecryptAES, format)
	return nil
}

func saveNewPassword(newData map[string]string) error {
	existingData, err := storage.ReadPasswordFile(passfilename)
	if err != nil {
		return fmt.Errorf("error reading passwords: %v", err)
	}

	existingData = append(existingData, newData)

	err = storage.SavePasswords(passfilename, existingData)
	if err != nil {
		return fmt.Errorf("error saving passwords: %v", err)
	}

	return nil
}

// Update function:
func updatePassword() error {
	// Display current passwords first
	if err := viewPasswords(); err != nil {
		return err
	}

	var site string
	fmt.Print("\nEnter site to update: ")
	fmt.Scanln(&site)

	// Find and update the password
	existingData, err := storage.ReadPasswordFile(passfilename)
	if err != nil {
		return err
	}

	found := false
	for i, entry := range existingData {
		// For each entry, get the encrypted site name (there's only one key per map)
		for encSite := range entry {
			// Decrypt the site name
			decryptedSite, err := encrypt.DecryptAES([]byte(secretkey), encSite)
			if err != nil {
				continue
			}

			// Compare the decrypted site name with input
			if decryptedSite == site {
				// Generate new password and encrypt it
				newPass := password.GeneratePass(14)
				encPass, err := encrypt.EncryptAES([]byte(secretkey), newPass)
				if err != nil {
					return err
				}

				// Use the original encrypted site name as the key
				existingData[i][encSite] = encPass
				found = true

				// Display update confirmation
				fmt.Printf("\n%sPassword Updated:%s\n", utils.Bold, utils.Reset)
				fmt.Printf("%sSite     :%s %s\n", utils.Green, utils.Reset, site)
				fmt.Printf("%sPassword :%s %s\n", utils.Green, utils.Reset, newPass)
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return fmt.Errorf("site not found")
	}

	return storage.SavePasswords(passfilename, existingData)
}

func searchPasswords() error {
	var searchTerm string
	fmt.Print("\nEnter search term: ")
	fmt.Scanln(&searchTerm)

	existingData, err := storage.ReadPasswordFile(passfilename)
	if err != nil {
		return err
	}

	matchingEntries := []map[string]string{}

	for _, entry := range existingData {
		for encSite, encPass := range entry {
			site, err := encrypt.DecryptAES([]byte(secretkey), encSite)
			if err != nil {
				continue
			}

			// Case-insensitive search
			if strings.Contains(strings.ToLower(site),
				strings.ToLower(searchTerm)) {
				matchingEntries = append(matchingEntries,
					map[string]string{encSite: encPass})
			}
		}
	}

	if len(matchingEntries) == 0 {
		fmt.Printf("\n%sNo matching entries found%s\n",
			utils.Bold, utils.Reset)
		return nil
	}

	fmt.Printf("\n%sMatching Entries:%s\n", utils.Bold, utils.Reset)
	format := display.ColorfulFormat()
	display.DisplayPasswords(matchingEntries, []byte(secretkey),
		encrypt.DecryptAES, format)

	return nil
}

func deletePassword() error {
	// Display current passwords first
	if err := viewPasswords(); err != nil {
		return err
	}

	var site string
	fmt.Print("\nEnter site to delete: ")
	fmt.Scanln(&site)

	// Find and delete the password
	existingData, err := storage.ReadPasswordFile(passfilename)
	if err != nil {
		return err
	}

	found := false
	newData := make([]map[string]string, 0, len(existingData))

	for _, entry := range existingData {
		shouldKeep := true
		// For each entry, check the encrypted site name
		for encSite := range entry {
			// Decrypt the site name
			decryptedSite, err := encrypt.DecryptAES([]byte(secretkey), encSite)
			if err != nil {
				continue
			}

			// If we find a match, mark for deletion
			if decryptedSite == site {
				found = true
				shouldKeep = false
				break
			}
		}

		// Keep all entries except the one marked for deletion
		if shouldKeep {
			newData = append(newData, entry)
		}
	}

	if !found {
		return fmt.Errorf("site not found")
	}

	fmt.Printf("\n%sPassword for %s%s %shas been deleted%s\n",
		utils.Bold, utils.Green, site, utils.Red, utils.Reset)

	return storage.SavePasswords(passfilename, newData)
}
