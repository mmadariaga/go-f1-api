package infrastructure_service_helper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type cacheValue struct {
	validUntil int64
	value      []byte
}

var reqCache = make(map[string]cacheValue)
var rwMutex sync.RWMutex

func HttpGet(targetUrl string, useCache bool) ([]byte, error) {

	rwMutex.RLock()
	cachedValue, isCached := reqCache[targetUrl]
	rwMutex.RUnlock()

	now := time.Now().Unix()
	useCachedValue := useCache && isCached && now <= cachedValue.validUntil

	if useCachedValue {
		return cachedValue.value, nil
	}

	// Fetch data
	resp, err := http.Get(targetUrl)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errorMsg := fmt.Sprintf("Request error: %v at %v", err, targetUrl)
		log.Error().Msg(errorMsg)

		return nil, errors.New(errorMsg)
	}

	defer resp.Body.Close()

	// Extract body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorMsg := fmt.Sprintf("Error reading response body: %v", err)
		log.Error().Msg(errorMsg)

		return nil, errors.New(errorMsg)
	}

	if useCache {
		tenMinuteAfter := time.Now().Add(10 * time.Minute).Unix()

		rwMutex.Lock()
		reqCache[targetUrl] = cacheValue{
			validUntil: tenMinuteAfter,
			value:      body,
		}
		rwMutex.Unlock()

	}

	return body, nil
}
