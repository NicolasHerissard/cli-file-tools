package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	create := flag.String("create", "", "Créer un fichier")
	file := flag.String("file", "", "Saisir le nom du fichier")
	write := flag.String("write", "", "Écrire dans un fichier")
	read := flag.Bool("read", false, "Lire le fichier")
	stat := flag.Bool("stat", false, "Voir les stats du fichier")
	delete := flag.Bool("delete", false, "Supprimer le fichier")
	rename := flag.String("rename", "", "Renommer un fichier")
	copy := flag.String("copy", "", "Copier le contenu du fichier")

	flag.Parse()

	if *create != "" {
		createFile(*create)
		return
	}

	if *file == "" {
		log.Fatal("Veuillez spécifier le chemin du fichier avec -file")
	}

	switch {
	case *read:
		readFile(*file)
	case *write != "":
		writeFile(*file, *write)
	case *stat:
		fileStats(*file)
	case *delete:
		deleteFile(*file)
	case *rename != "":
		renameFile(*file, *rename)
	case *copy != "":
		copyFile(*file, *copy)
	default:
		fmt.Println("Opération inconnue")
	}
}