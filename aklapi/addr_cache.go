package aklapi

type addrResponseCache map[string]AddrResponse

var addrCache = make(addrResponseCache)

func (c addrResponseCache) Lookup(searchText string) (resp AddrResponse, ok bool) {
	resp, ok = c[searchText]
	return
}

func (c addrResponseCache) Add(searchText string, ar AddrResponse) {
	c[searchText] = ar
}
