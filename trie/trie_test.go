//  Copyright 2019 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package trie

import (
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func randomPath() string {
	rand.Seed(time.Now().UnixNano())

	nodeNum := rand.Intn(10)
	nodes := make([]string, 0)
	for i := 0; i < nodeNum; i++ {
		nodes = append(nodes, randomStr(10))
	}
	return "/" + strings.Join(nodes, "/")
}

func GenTree(pathNum int) (tree *Tree, dstPath string) {
	tree = New()

	for index := 0; index < pathNum; index++ {
		path := randomPath()
		tree.Put(path, 123)

		dstPath = path
		if index == rand.Intn(pathNum) {
			dstPath = path
		}
	}
	return
}

func benchmarkMatch(tree *Tree, dstPath string, b *testing.B) {
	for i := 0; i < 1000000; i++ {
		if _, ok := tree.Match(dstPath); !ok {
			b.Fatalf("match failed!")
		}
	}
}

func BenchmarkHundred_Match(b *testing.B) {
	hundredTree, dstPath := GenTree(100)
	b.ReportAllocs()
	b.ResetTimer()
	benchmarkMatch(hundredTree, dstPath, b)
}

func BenchmarkThousand_Match(b *testing.B) {
	thousandTree, dstPath := GenTree(1000)
	b.ReportAllocs()
	b.ResetTimer()
	benchmarkMatch(thousandTree, dstPath, b)
}

func BenchmarkFiveThousand_Match(b *testing.B) {
	fiveThousandTree, dstPath := GenTree(10000)
	b.ReportAllocs()
	b.ResetTimer()
	benchmarkMatch(fiveThousandTree, dstPath, b)
}

func BenchmarkTenThousand_Match(b *testing.B) {
	tenThousandTree, dstPath := GenTree(100000)
	b.ReportAllocs()
	b.ResetTimer()
	benchmarkMatch(tenThousandTree, dstPath, b)
}

func TestTrie_Put(t *testing.T) {
	routes := []string{
		"/abc",
		"/eft",
		"/test/test",
		"/test/abc/aaa",
		"/test/abc/bcd",
		"/test/abc",
		"/test/bcd",
		"/test/test2/test",
		"/test/test2/test/abc/aaa",
	}
	tree := New()
	for _, route := range routes {
		tree.Put(route, 123)
	}
	tree.Dump()
}

func TestTrie_GoroutinesMatch(t *testing.T) {
	tree, dstPath := GenTree(10000)
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			tree.Match(dstPath)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestTree_Exist(t *testing.T) {
	routes := []string{
		"/test",
		"/test/test",
		"/test/abc/aaa",
		"/test/abc/bcd",
		"/test/abc",
	}

	tree := New()
	for _, route := range routes {
		tree.Put(route, 123)
	}

	errorRoutes := []string{
		"/",
		"/abc/abc",
		"/test/abf",
		"/test/test2",
		"/test2/abc/abc",
	}
	for _, route := range errorRoutes {
		_, ok := tree.Exist(route)
		assert.True(t, !ok, "%s", route)
	}

	for _, route := range routes {
		_, ok := tree.Exist(route)
		assert.True(t, ok, "%s", route)
	}
}

func TestTrie_Match(t *testing.T) {
	tree := New()

	routes := []string{
		"/test",
		"/test/test",
		"/test/abc/aaa",
		"/test/abc/bcd",
		"/test/abc",
	}
	for _, route := range routes {
		tree.Put(route, 123)
	}
	tree.Dump()

	for _, route := range routes {
		_, ok := tree.Match(route)
		assert.True(t, ok, "match route %s failed", route)
	}

	errorRoutes := []string{
		"/",
		"/abc/abc",
		"/test/abf",
		"/test/test2",
		"/test2/abc/abc",
	}
	for _, route := range errorRoutes {
		_, ok := tree.Match(route)
		assert.True(t, !ok, "match a not exit route %s succ", route)
	}
}

func TestTree_MatchWildcard(t *testing.T) {
	tree := New()

	routes := []string{
		"/onepiece/.+",
		"/onepiece/test",
		"/onepiece/v2/test",
		"/onepiece/v1/guest/register",
		"/test.+",
		"/.+",
	}
	for _, route := range routes {
		tree.Put(route, route)
	}
	tree.Dump()

	type R struct {
		Path    string
		Pattern string
	}

	correctRoutes := []R{
		{"/onepiece/test", "/onepiece/test"},
		{"/onepiece/test2", "/onepiece/.+"},
		{"/onepiece/v2/test", "/onepiece/v2/test"},
		{"/onepiece/v1/guest/register", "/onepiece/v1/guest/register"},
		{"/onepiece/v1/user/profile", "/onepiece/.+"},
		{"/onepiece/v1", "/onepiece/.+"},
		{"/test/abc", "/test.+"},
		{"/test", "/test.+"},
	}
	for _, route := range correctRoutes {
		entry, ok := tree.Match(route.Path)
		assert.True(t, ok, "match route %s failed", route)
		assert.Equal(t, route.Pattern, entry)
	}
}

func TestTree_Empty(t *testing.T) {
	tree := New()
	assert.True(t, tree.Empty(), "[BUG: error in method Empty.]")

	var err error
	tree.Put("/test/abc", 233)
	tree.Dump()

	err = tree.Del("/test/abc")
	assert.NoError(t, err)
	tree.Dump()

	assert.True(t, tree.Empty(), "[BUG: error in method Empty.]")
}

func TestTrie_Del(t *testing.T) {
	routes := []string{
		"/",
		"/abc",
		"/eft",
		"/test/test",
		"/test/abc/aaa",
		"/test/abc/bcd",
		"/test/abc",
		"/test/bcd",
		"/test/test2/test",
		"/test/test2/test/abc/aaa",
		"/onepiece/.+",
		"/onepiece/v2/.+",
		"/onepiece/v1/user/profile",
	}
	tree := New()
	for _, route := range routes {
		tree.Put(route, 123)
	}
	tree.Dump()
	for _, route := range routes {
		_, ok := tree.Match(route)
		assert.True(t, ok, "match route %s failed", route)
	}
	deleteRoutes := []string{
		"/test/test",
		"/test/abc/aaa",
		"/test/abc/bcd",
		"/test/abc",
		"/test/bcd",
		"/test/test2/test/abc/aaa",
	}
	for _, route := range deleteRoutes {
		err := tree.Del(route)
		assert.NoError(t, err, "delete route %s failed", route)
	}
	tree.Dump()
	for _, route := range deleteRoutes {
		_, ok := tree.Match(route)
		assert.True(t, !ok, "match a not exist route succ")
	}

	notExistRoutes := []string{
		"/bcd/test",
	}
	for _, route := range notExistRoutes {
		err := tree.Del(route)
		assert.Error(t, err, "delete a not exist route %s succ", route)
	}
}

//   /
//   |-test
//      |-.+
//      |-test2
//      |-test3
//
//
//
//
