### Technical task
#### Description

**Create cache server of web pages with REST API.**

__Functional requirements:__
Operations
1) "put" operation - association of key (page url) with page bytes + increases page hit rating.
2) "get" operation - fetch of page bytes by key (page url) + increases page hit rating.
3) "delete" operation - removes key from cache. 
4) "top" operation - returns ordered list of top 100 keys sorted by hit rating desc.
5) items that was not used during some configurable amount of time should be evicted from cache (TTL).

Configs
1) Config param that limits max count for stored items.
2) Config param that limits max size of single item in bytes.
3) Config param that limits max size of whole cache in bytes.

__Non-functional requirements:__
1) cache should support concurrent access.
2) all operation should not have linear or worse complexity dependence from number of stored items.
3) if cache is full in terms of max count or max overall size limits, it should evict the oldest item (among the equally old elements it should evict one with lower hit rating).
