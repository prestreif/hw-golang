package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		// Согласно интерфейса тут должен вернуться ListItem,
		// так-как значение Hesh, насколько я понял
		// если всё корректно то val нужно сравнивать так:
		require.Equal(t, 100, val.(*ListItem).Value.(cacheItem).value)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val.(*ListItem).Value.(cacheItem).value)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val.(*ListItem).Value.(cacheItem).value)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(5)

		c.Set("test1", 1)
		c.Set("test2", 2)
		c.Set("test3", 3)

		_, ok := c.Get("test1")
		require.True(t, ok)

		_, ok = c.Get("test2")
		require.True(t, ok)

		_, ok = c.Get("test3")
		require.True(t, ok)

		c.Clear()

		_, ok = c.Get("test1")
		require.False(t, ok)

		_, ok = c.Get("test2")
		require.False(t, ok)

		_, ok = c.Get("test3")
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
