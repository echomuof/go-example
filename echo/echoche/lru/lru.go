/**
 *
 * @author: echomuof
 * @created: 2020/11/20
 */
package lru

import "container/list"

type Cache struct {
	maxBytes  int64                         //允许使用的最大内存
	nbytes    int64                         //当前已经使用的内存
	ll        *list.List                    //LRU的链表
	cache     map[string]*list.Element      //保存连接中节点的键值对
	OnEvicted func(key string, value Value) //节点被删除后调用的方法
}

// 节点的值
type Value interface {
	Len() int
}

//节点
type entry struct {
	key   string
	value Value
}

/*
 * 初始化缓存
 */
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

/*
 * 从缓存中查找，找到后移动到队首
 */
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

/*
 * 删除队尾元素，最近最少使用
 */
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		//删除后回调函数
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

/*
 * 新增或修改
 */
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	//如果空间不够，清理缓存
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

/*
 * 查询缓存中的数据个数
 */
func (c *Cache) Len() int {
	return c.ll.Len()
}
