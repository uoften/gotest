package main

import (
	"fmt"
	"github.com/allegro/bigcache"
	"log"
	"time"
)

func main() {

	config := bigcache.Config {
		Shards: 1024,					// 存储的条目数量，值必须是2的幂
		LifeWindow: 1*time.Second,	// 超时后条目被删除
		MaxEntriesInWindow: 0,		// 在 Life Window 中的最大数量，
		MaxEntrySize: 0,			// 条目最大尺寸，以字节为单位
		HardMaxCacheSize: 0,		// 设置缓存最大值，以MB为单位，超过了不在分配内存。0表示无限制分配
	}

	cache, initErr := bigcache.NewBigCache(config)
	if initErr != nil {
		log.Fatal(initErr)
	}

	cache.Set("key12345678", []byte("1zxcvbnmkjhgfdsaqwertyuiopkjhgferettyyyt"))
	cache.Set("key12345679", []byte("2jhgfdsaqwertyuiopkjhgfdxcvnmjuytfghjkng"))
	cache.Set("key12345680", []byte("3jhgfdsaqwertyuiopkjhgfdxcvnmjuytfghjkop"))
	time.Sleep(3 * time.Second)
	cache.Set("key12345681", []byte("4jhgfdsaqwertyuiopkjhgfdxcvnmjuytfghjkop"))


	entry, err := cache.Get("key12345678")
	fmt.Println(string(entry),err)

	entry, err = cache.Get("key12345679")
	fmt.Println(string(entry),err)

	entry, err = cache.Get("key12345680")
	fmt.Println(string(entry),err)

	entry, err = cache.Get("key12345681")
	fmt.Println(string(entry),err)


}