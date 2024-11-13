package processing

import (
    "strings"
    "sync"
)

type KeywordProcessor struct {
    keywordCounts map[string]int
    mu            sync.Mutex
}


func NewKeywordProcessor(keywords []string) *KeywordProcessor {
    counts := make(map[string]int)
    for _, keyword := range keywords {
        counts[strings.ToLower(keyword)] = 0
    }
    return &KeywordProcessor{
        keywordCounts: counts,
    }
}

func (kp *KeywordProcessor) ProcessChunk(chunk string) map[string]int {
    kp.mu.Lock()
    defer kp.mu.Unlock()

    words := strings.Fields(strings.ToLower(chunk))
    for _, word := range words {
        if _, exists := kp.keywordCounts[word]; exists {
            kp.keywordCounts[word]++
        }
    }
    countsCopy := make(map[string]int)
    for k, v := range kp.keywordCounts {
        countsCopy[k] = v
    }
    return countsCopy
}
