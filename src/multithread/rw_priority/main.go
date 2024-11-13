package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	NUM_READERS = 5
	NUM_WRITERS = 2
	filename    = "shared_file.txt"
)

var (
	rwlock        sync.RWMutex // Read-Write lock
	counterReader int
	counterWriter int
)

// Reader function
func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		rwlock.RLock() // Acquire read lock
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Erreur lors de l'ouverture du fichier en lecture: %v\n", err)
			rwlock.RUnlock()
			continue
		}

		fmt.Printf("\033[1;32mLecteur %d: ------------------------lecture du fichier----------------------------\033[0m\n", id)
		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if err != nil {
				break
			}
			fmt.Print(string(buf[:n]))
		}
		counterReader++
		fmt.Printf("\nFIN LECTURE, lecteur %d\n", id)
		fmt.Printf("Total acces pour les lecteurs: %d\n", counterReader)

		file.Close()
		rwlock.RUnlock() // Release read lock
		time.Sleep(1 * time.Second)
	}
}

// Writer function
func writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		rwlock.Lock() // Acquire write lock
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Erreur lors de l'ouverture du fichier en ecriture: %v\n", err)
			rwlock.Unlock()
			continue
		}

		fmt.Printf("======================DEBUT ECRITURE, ecrivain %d\n", id)
		_, err = file.WriteString(fmt.Sprintf("Ecrivain %d ecrit dans le fichier.\n", id))
		if err != nil {
			fmt.Printf("Erreur lors de l'ecriture: %v\n", err)
		}
		counterWriter++
		fmt.Printf("======================FIN ECRITURE, ecrivain %d\n", id)
		fmt.Printf("Total acces pour les ecrivains: %d\n", counterWriter)
		file.Sync() // Ensure data is written to disk

		file.Close()
		rwlock.Unlock() // Release write lock
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// Create or clear the shared file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Erreur lors de la creation du fichier: %v\n", err)
		return
	}
	file.Close()

	var wg sync.WaitGroup

	// Create reader goroutines
	for i := 0; i < NUM_READERS; i++ {
		wg.Add(1)
		go reader(i+1, &wg)
	}

	// Create writer goroutines
	for i := 0; i < NUM_WRITERS; i++ {
		wg.Add(1)
		go writer(i+1, &wg)
	}

	// Wait for all goroutines (This will block forever as readers/writers run indefinitely)
	wg.Wait()
}
