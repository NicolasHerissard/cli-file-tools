package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	create := flag.String("create", "", "Saisir le nom du fichier à créer")
	file := flag.String("file", "", "Saisir le nom du fichier")
	write := flag.String("write", "", "Saisir le contenu à écrire")
	read := flag.Bool("read", false, "Lire le fichier")
	stat := flag.Bool("stat", false, "Voir les stats du fichier")
	delete := flag.Bool("delete", false, "Supprimer le fichier")
	rename := flag.String("rename", "", "Renommer un fichier")

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
	default:
		fmt.Println("Opération inconnue")
	}
}