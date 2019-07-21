package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/pbkdf2"
	"sync"
)

var key string
var salt string
var log = logrus.New()
var wg sync.WaitGroup

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	flag.StringVar(&key, "k", "", "RestrictionsPasswordKey in base64")
	flag.StringVar(&salt, "s", "", "RestrictionsPasswordSalt in base64")
	flag.Parse()
}

func main() {
	bSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil || salt == "" {
		log.Fatalf("bad salt; could not be decoded")
	}
	bKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil || key == "" {
		log.Fatalf("bad key; could not be decoded")
	}
	result := make(chan string)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go force(fmt.Sprintf("%04d", i), result, bSalt, bKey)
	}

	go func() {
		defer close(result)
		wg.Wait()
	}()

	pass, ok := <-result
	if !ok {
		log.Fatalf("could not find password. Did you copy the values properly?")
	}
	log.Printf("pass is %s", pass)
}

func force(pass string, result chan string, bSalt []byte, bKey []byte) {
	defer wg.Done()
	k := pbkdf2.Key([]byte(pass), bSalt, 1000, 20, sha1.New)
	log.Debugf("pass %s gives %s", pass, base64.StdEncoding.EncodeToString(k))
	if bytes.Equal(k, bKey) {
		result <- pass
	}
}
