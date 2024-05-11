package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/**
 * Affiche le contenu du fichier donner en paramètre
 * @params file: string
 */
func readFile(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Le fichier n'existe pas ", err)
		} else {
			log.Fatal("Erreur de lecture du fichier")
		}
	}

	fmt.Println("Contenu du fichier :")
	fmt.Println(string(data))
}

/**
 * Crée le fichier donner en paramètre
 * @params create: string
 */
func createFile(create string) {
	file, err := os.Create(create)

	tab := strings.Split(file.Name(), ".")
	if len(tab) != 2 {
		log.Fatal("Le fichier ne contient pas d'extension")
	}

	defer file.Close()

	if err != nil {
		if os.IsExist(err) {
			log.Fatal("Le fichier ", file.Name() ," existe déjà")
		} else {
			log.Fatal("Erreur lors de la création du fichier : ", err)
		}
	}

	fmt.Println("Le fichier : ", file.Name(), " à été créer")
}

func writeFile(file string, write string) {
	f, err := os.OpenFile(file, os.O_RDONLY|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Erreur d'ouverture du fichier : ", err)
	}

	defer f.Close()
	
	n, err := f.WriteString(write)
	if err != nil {
		log.Fatal("Erreur d'écriture du fichier : ", err)
	}

	fmt.Println("Nombre d'octet écrit sur le fichier : ", n)
}

func fileStats(file string) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("Erreur d'ouverture du fichier : ", err)
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nom du fichier : ", info.Name())
	fmt.Println("Taille du fichier : ", info.Size())
	fmt.Println("Mode du fichier : ", info.Mode())
	if info.IsDir() {
		fmt.Println("C'est un dossier")
	} else {
		fmt.Println("Ce n'est pas un dossier")
	}
	fmt.Println("Sys : ", info.Sys())
	year, month, day := info.ModTime().Date()
	fmt.Println("ModTime : ", day,"-", month, "-", year)
}

func deleteFile(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fichier supprimée")
}

func renameFile(file string, newName string) {
	err := os.Rename(file, newName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fichier renommée")
}

func copyFile(file string, newFile string) {
	read, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Le fichier n'existe pas ", err)
		} else {
			log.Fatal("Erreur de lecture du fichier ", err)
		}
	}

	error := os.WriteFile(newFile, read, 0600)
	if error != nil {
		if os.IsNotExist(error) {
			log.Fatal("Le nouveau fichier n'existe pas ", err)
		} else {
			log.Fatal("Erreur d'écriture dans le nouveau fichier ", err)
		}
	}

	fmt.Println("Le contenu à été copié")
}