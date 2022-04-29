package main

import (
	"fmt"
	"github.com/samarec1812/crypto-analysis-vigenere/algorithm"
	"io/ioutil"
)

func main() {
MAINLOOP:
	for {
		fmt.Println("======================================================")
		fmt.Println("Program crypto analysis Vigenere cipher")
		fmt.Println("======================================================")
		fmt.Println("1: Start program")
		fmt.Println("2: Help")
		fmt.Println("3: Exit")
		fmt.Println("======================================================")
		var action string
		fmt.Printf("select: ")
		fmt.Scanln(&action)
		switch action {
		case "1":
			fmt.Print("Enter text file name with text for encoded: ")
			var textFileName string
			fmt.Scanln(&textFileName)
			txtBytes, err := ioutil.ReadFile(textFileName + ".txt")
			if err != nil {
				fmt.Println("Error: can't read file. Enter again")
				continue
			}
			var text string

			fmt.Println()
			fmt.Println("===================COMMENTARY=========================")
			fmt.Println("Comment: supported key length from 2 to 10.")
			fmt.Println("======================================================")
			fmt.Println()
			fmt.Print("Enter key: ")

			var key string
			fmt.Scanln(&key)

			text = algorithm.ChangeText(string(txtBytes))
			err = algorithm.IsCorrectText(string(txtBytes))
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			key = algorithm.ChangeKey(key)
			err = algorithm.IsCorrectKey(key)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println()
			fmt.Println("===================COMMENTARY=========================")
			fmt.Println("Comment: To encrypt the text, only the letters remain.")
			fmt.Println("All other characters are subject to deletion")
			fmt.Println("======================================================")
			fmt.Println()
			fmt.Println("TEXT: ", text)
			fmt.Println("KEY: ", key)

		MENUPROGRAM:
			for {
				fmt.Println("======================================================")
				fmt.Println("1: Encrypted ")
				fmt.Println("2: Analysis")
				fmt.Println("3: Return to menu")
				fmt.Println("4: Exit")
				fmt.Println("======================================================")
				fmt.Printf("select: ")
				fmt.Scanln(&action)
				switch action {
				case "1":
					size := 26
					encrypted := algorithm.Encrypt(text, key, size)
					fmt.Println("Encrypted:", encrypted)
					err := ioutil.WriteFile("encrypted.txt", []byte(encrypted), 0644)
					if err != nil {
						fmt.Println("Error: can't write encrypted text")
						continue
					}
					fmt.Println("Create file \"encrypted.txt\" with encrypted text")
				case "2":
					txtBytes, err = ioutil.ReadFile("encrypted.txt")
					if err != nil {
						fmt.Println("Error: text isn't encrypted yet. Enter 1 for encrypted")
						continue
					}
					encrypted := string(txtBytes)
					rangeKey := algorithm.FindKeyLength(encrypted)
				CHOOSEKEY:
					for {
						fmt.Print("Enter key length: ")
						var keyLength int
						fmt.Scanln(&keyLength)
						err = algorithm.CheckLengthKey(keyLength, rangeKey)
						if err != nil {
							fmt.Printf("Error: %s\n", err.Error())
							continue
						}
						findKey, table := algorithm.FindKey(keyLength, encrypted)
						fmt.Println("FIND KEY: ", findKey)
						fmt.Println("----------------------------------------------")
						fmt.Println("Decrypted your ciphertext with this key [y/n]?")
						fmt.Scanln(&action)

						switch action {
						case "y":
							size := 26
							decrypted := algorithm.Decrypt(encrypted, findKey, size)
							fmt.Println("------------------------------------------------")
							fmt.Println("Decrypted:", decrypted)
							fmt.Println("------------------------------------------------")
							fmt.Println("The result of the cryptoanalysis mathes your expection? [y/n]")
							fmt.Scanln(&action)
							switch action {
							case "y":
								fmt.Println("Atack on the Vigenere chiper was carried out successfully")
								continue MAINLOOP
							case "n":
								fmt.Println("Probably the key was found incorrectly. Find key of this length? [y/n]?")
								fmt.Scanln(&action)
								switch action {
								case "y":
								OTHERKEYS:
									for {
										keys := algorithm.AllKeys(table)
										fmt.Println("------------------")
										fmt.Printf("Number |Key\t|\n")
										for i, oneKey := range keys {
											fmt.Printf("%d | %s\t|\n", i+1, oneKey)
											fmt.Println("------------------")
										}
										fmt.Print("Enter key number: ")
										// fmt.Printf("select: ")
										var numberKeys int
										fmt.Scanln(&numberKeys)
										decrypted = algorithm.Decrypt(text, keys[numberKeys-1], size)
										fmt.Println("------------------------------------------------")
										fmt.Println("Decrypted:", decrypted)
										fmt.Println("------------------------------------------------")
										fmt.Println("The result of the cryptoanalysis mathes your expection? [y/n]")
										fmt.Scanln(&action)
										switch action {
										case "y":
											fmt.Println("Atack on the Vigenere chiper was carried out successfully")
											continue MAINLOOP
										case "n":
											fmt.Println("Probably the key was found incorrectly. Find key of this length? [y/n]?")
											fmt.Scanln(&action)
											switch action {
											case "y":
												continue OTHERKEYS
											case "n":
												fmt.Println("Choose key another length")
												fmt.Println("------------------------------------------------")
												continue CHOOSEKEY
											default:
												fmt.Println("Error: unsupported action")
												continue OTHERKEYS
											}
										default:
											fmt.Println("Error: unsupported action")
											continue OTHERKEYS
										}
									}
								case "n":
									fmt.Println("Choose key another length")
									fmt.Println("------------------------------------------------")
									continue CHOOSEKEY
								default:
									fmt.Println("Error: unsupported action")
									continue MENUPROGRAM
								}
							}
						case "n":
							continue CHOOSEKEY
						default:
							fmt.Println("Error: unsupported action")
							continue MENUPROGRAM
						}
					}

				case "3":
					continue MAINLOOP
				case "4":
					fmt.Println("Program exit successful")
					return
				default:
					fmt.Println("Error: unsupported action. Enter again")
				}
			}

			// encrypted := algorithm.Encrypt(*text, *key, size)
			//fmt.Println("Encrypted:", encrypted)
			return

		case "2":
			fmt.Println("+++++++++++++++++INFO+++++++++++++++++++++")
			fmt.Println("This program developed by Kondratev Alexey")
			fmt.Println("Group: A-13-18")
		case "3":
			fmt.Println("Program exit successful")
			return
		default:
			fmt.Println("Error: unsupported action. Enter again")
		}
		//fmt.Println("Введите текст:")
		//var text string
		//fmt.

	}
	//text := flag.String("text", "", "source text")
	//key := flag.String("key", "", "key")
	//flag.Parse()
	//
	//size := 26
	//
	//fmt.Println("Source:", *text)
	//fmt.Println("Key:", *key)
	//
	//encrypted := algorithm.Encrypt(*text, *key, size)
	//fmt.Println("Encrypted:", encrypted)
	//
	//
	//decrypted := algorithm.Decrypt(encrypted, *key, size)
	//fmt.Println("Decrypted:", decrypted)
	//
	//algorithm.FindKeyLength(encrypted)
	//fmt.Println(algorithm.IndexIC(encrypted))
	//findKey := algorithm.FindKey(3, encrypted)
	//decrypted2 := algorithm.Decrypt(encrypted, findKey, 26)
	//fmt.Println("Decrypted:", decrypted2)
}
