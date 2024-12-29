package clienthttp

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// ReqHTTP effectue une requête HTTP GET et affiche le corps de la réponse.
func ReqHTTP(url string, logger *log.Logger) {
	// Définir un timeout pour éviter que la requête ne bloque indéfiniment
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Printf("Démarrage de la requête HTTP vers : %s", url)

	// Créer une requête avec un contexte
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.Fatalf("Erreur lors de la création de la requête HTTP : %v", err)
	}

	// Effectuer la requête avec un client HTTP
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			logger.Fatalf("Timeout : le serveur n'a pas répondu dans le délai imparti (%s)", timeout)
		}
		logger.Fatalf("Erreur lors de l'exécution de la requête HTTP : %v", err)
	}
	defer response.Body.Close()

	// Vérifier le code de statut HTTP
	if response.StatusCode != http.StatusOK {
		logger.Fatalf("Le serveur a renvoyé un code de statut inattendu : %d %s", response.StatusCode, response.Status)
	}

	// Lire et afficher le corps de la réponse
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Fatalf("Erreur lors de la lecture du corps de la réponse : %v", err)
	}

	logger.Printf("Réponse reçue du serveur (%d %s)", response.StatusCode, response.Status)
	logger.Printf("Réponse du serveur :\n%s\n", body)
	fmt.Print("*******************************\n")
	fmt.Printf("\n\tIP public :\n\n\t%s\n", body)
	fmt.Print("*******************************\n")
}
