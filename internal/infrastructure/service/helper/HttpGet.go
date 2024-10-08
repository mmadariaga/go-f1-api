package infrastructure_service_helper

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type cacheValue struct {
	validUntil int64
	value      []byte
}

var reqCache = make(map[string]cacheValue)

func HttpGet(targetUrl string, useCache bool) []byte {

	cachedValue, isCached := reqCache[targetUrl]

	now := time.Now().Unix()
	useCachedValue := useCache && isCached && now <= cachedValue.validUntil

	if useCachedValue {
		return cachedValue.value
	}

	// Fetch data
	resp, err := http.Get(targetUrl)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Panic().Msg(
			fmt.Sprintf("Error al hacer la petición: %v en %v", err, targetUrl),
		)
	}

	defer resp.Body.Close()

	// Extract body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic().Msg(
			fmt.Sprintf("Error al leer el cuerpo de la respuesta: %v", err),
		)
	}

	if useCache {
		tenMinuteAfter := time.Now().Add(10 * time.Minute).Unix()
		reqCache[targetUrl] = cacheValue{
			validUntil: tenMinuteAfter,
			value:      body,
		}
	}

	return body
}
