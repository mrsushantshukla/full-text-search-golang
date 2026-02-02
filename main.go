package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/mrsushantshukla/full-text-search-golang/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-pages-logging10.xml.gz", "Dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	log.Println("Full text search in progress")
	start := time.Now()
	docs, err := utils.loadDocuments(dumpPath)
	if err != nil {
		log.Fatal("Could not load document: ", err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	start = time.Now()
	for _, id := range matchedIDs {
		log.Printf("%d\t%s\n", id, docs[id].Text)
	}
}
