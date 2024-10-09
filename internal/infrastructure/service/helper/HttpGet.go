package infrastructure_service_helper

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
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

type HttpGetExtraArgs struct {
	UseCache   bool
	Retry      bool
	RetryCount int
}

const MAX_RETRIES = 2

func HttpGet(targetUrl string, args *HttpGetExtraArgs) ([]byte, error) {

	if args.RetryCount > 0 {
		delay := time.Duration(args.RetryCount*2500 + rand.Intn(1000))
		time.Sleep(delay * time.Millisecond)
	}

	rwMutex.RLock()
	cachedValue, isCached := reqCache[targetUrl]
	rwMutex.RUnlock()

	now := time.Now().Unix()
	useCachedValue := args.UseCache && isCached && now <= cachedValue.validUntil

	if useCachedValue {
		return cachedValue.value, nil
	}

	// Fetch data
	resp, err := http.Get(targetUrl)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {

		if args.Retry && args.RetryCount < MAX_RETRIES {
			log.Debug().Msg("Request error. Retrying request to " + targetUrl)
			args.RetryCount++
			return HttpGet(targetUrl, args)
		}

		errorMsg := fmt.Sprintf("Request error: %v at %v", err, targetUrl)
		log.Error().Msg(errorMsg)

		return nil, errors.New(errorMsg)
	}

	defer resp.Body.Close()

	// Extract body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if args.Retry && args.RetryCount < MAX_RETRIES {
			log.Debug().Msg("Error reading response body. Retrying request to " + targetUrl)
			args.RetryCount++
			return HttpGet(targetUrl, args)
		}

		errorMsg := fmt.Sprintf("Error reading response body: %v", err)
		log.Error().Msg(errorMsg)

		return nil, errors.New(errorMsg)
	}

	if args.UseCache {
		oneHour := time.Now().Add(60 * time.Minute).Unix()

		rwMutex.Lock()
		reqCache[targetUrl] = cacheValue{
			validUntil: oneHour,
			value:      body,
		}
		rwMutex.Unlock()
	}

	return body, nil
}
