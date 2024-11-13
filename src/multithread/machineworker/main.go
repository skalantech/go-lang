package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	N_MACHINES = 5
	N_STOCK    = 100
)

var (
	stock         [N_STOCK]int
	articleTraite [N_STOCK]bool
	articleIndex  int
	mutex         sync.Mutex
	wg            sync.WaitGroup
)

func initialiserStock() {
	for i := 0; i < N_STOCK; i++ {
		stock[i] = i + 1
		articleTraite[i] = false
	}
}

func simulerTache(machineID, article int) {
	resultat := machineID*100 + article
	fmt.Printf("Machine: %d traite l'article %d, assigne SKU: %d\n", machineID, article, resultat)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
}

func machineWorker(machineID int) {
	defer wg.Done()
	machineCompteur := 0

	for {
		mutex.Lock()
		if articleIndex >= N_STOCK {
			mutex.Unlock()
			break
		}

		index := articleIndex
		articleIndex++
		mutex.Unlock()

		if !articleTraite[index] {
			simulerTache(machineID+1, stock[index])
			articleTraite[index] = true
			machineCompteur++
		}
	}
	fmt.Printf("Machine %d a fini son travail avec %d articles traites\n", machineID, machineCompteur)
}

func main() {
	initialiserStock()
	wg.Add(N_MACHINES)

	for i := 0; i < N_MACHINES; i++ {
		go machineWorker(i)
	}

	wg.Wait()

	// Vérification pour tous les articles traités
	flag := true
	for i := 0; i < N_STOCK; i++ {
		if !articleTraite[i] {
			fmt.Printf("Erreur: l'article %d n'a pas ete traite.\n", stock[i])
			flag = false
		}
	}

	if flag {
		fmt.Println("Tous les articles ont ete traites avec succes.")
	}
}
