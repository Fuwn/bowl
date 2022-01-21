package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"time"
)

func readAll() [][]string {
	f, err := os.Open("bowl.csv")
	defer f.Close()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, _ = os.Create("bowl.csv")
			_, err := f.Write([]byte("#type,notes,datetime,ip\n"))
			if err != nil {
				log.Fatalf("error writing to file: %s\n", err)
			}
		} else {
			log.Fatalf("failed to open file: %s\n", err)
		}
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("failed to read file: %s\n", err)
	}

	return records
}

func write(formType string, formNotes string, ip string) {
	records := readAll()

	f, err := os.Create("bowl.csv")
	defer f.Close()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, _ = os.Create("bowl.csv")
			_, err := f.Write([]byte("#type,notes,datetime,ip\n"))
			if err != nil {
				log.Fatalf("error writing to file: %s\n", err)
			}
		} else {
			log.Fatalf("failed to open file: %s\n", err)
		}
	}

	records = append(records, []string{
		formType, formNotes, time.Now().String(), ip,
	})

	w := csv.NewWriter(f)
	defer w.Flush()

	err = w.WriteAll(records)
	if err != nil {
		log.Fatalf("error writing entry to file: %s\n", err)
	}
}
