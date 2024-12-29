package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"ali.go/go-tools/osdetector"
)

// InitLogger initialise un logger qui écrit dans un fichier de log.
func InitLogger(logFile string, path string) *log.Logger {
	// Obtenir le chemin dynamique du fichier de log basé sur l'OS
	logPath := osdetector.GetLogPath()

	// Vérifier si le chemin a été trouvé
	if logPath == "" {
		log.Fatalf("Erreur : Système d'exploitation non supporté ou chemin de log introuvable")
	}

	// Créer le chemin complet du fichier
	fullPath := filepath.Join(path, logFile)
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture du fichier de log : %v", err)
	}

	// Créer un logger personnalisé
	logger := log.New(file, "", 0) // Désactiver les drapeaux par défaut

	// Définir un writer personnalisé
	logger.SetOutput(&customLogWriter{
		file: file,
	})

	return logger
}

// Structure pour personnaliser l'output des logs
type customLogWriter struct {
	file *os.File
}

// Write implémente un writer qui formate le log avec la date, l'heure et Lshortfile sans extension .go
func (w *customLogWriter) Write(p []byte) (n int, err error) {
	// Récupérer l'appelant pour Lshortfile
	_, file, _, ok := runtime.Caller(3)
	if ok {
		// Retirer l'extension .go
		fileName := strings.TrimSuffix(filepath.Base(file), ".go")
		// Ajouter la date et l'heure
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		// Construire le message formaté
		p = []byte(fmt.Sprintf("%s [%s] %s", timestamp, fileName, string(p)))
	}

	// Écrire dans le fichier
	return w.file.Write(p)
}

// WriteLog écrit un texte dans un fichier de log donné.
func WriteLog(logFile string, path string, text string) {
	// Créer le chemin complet du fichier
	fullPath := filepath.Join(path, logFile)

	// Ouvrir ou créer le fichier avec les permissions adéquates
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture ou de la création du fichier de log : %v", err)
	}
	defer file.Close()

	// Créer un logger pour écrire dans le fichier
	logger := log.New(file, "", log.Ldate|log.Ltime)

	// Écrire le texte dans le fichier de log
	logger.Println(text)

	// Message pour indiquer que l'écriture s'est terminée avec succès
	//fmt.Printf("Texte écrit avec succès dans le fichier de log : %s\n", fullPath)
}

func BeginWriteLogDetails() {
	WriteLog("goGetMyPubIP.log", osdetector.GetLogPath(), "*********************")
	WriteLog("goGetMyPubIP.log", osdetector.GetLogPath(), "* Programme démarré *")
	WriteLog("goGetMyPubIP.log", osdetector.GetLogPath(), "*********************")
}

func EndWriteLogDetails() {
	WriteLog("goGetMyPubIP.log", osdetector.GetLogPath(), "*********************")
	WriteLog("goGetMyPubIP.log", osdetector.GetLogPath(), "* Programme terminé *")
	WriteLog("goGetMyPubIP.log", osdetector.GetLogPath(), "*********************")
}
