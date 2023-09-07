package test_test

import (
	"testing"

	"github.com/cjwagner/httpcache"
	"github.com/cjwagner/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
