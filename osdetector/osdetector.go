package osdetector

import (
	"fmt"
	"runtime"
)

// GetLogPath détecte le système d'exploitation et renvoie le chemin du fichier de log approprié.
func GetLogPath() string {
	var logPath string

	// Détection du système d'exploitation
	if runtime.GOOS == "windows" {
		// Sur Windows, utiliser ce chemin
		logPath = "D:\\03.TechMania\\04.Dev\\ali.go\\"
	} else if runtime.GOOS == "linux" {
		// Sur Linux, utiliser ce chemin
		logPath = "/home/support/axmsd/go"
	} else {
		// Si le système d'exploitation n'est pas supporté
		fmt.Println("OS non supporté")
		logPath = ""
	}

	return logPath
}
