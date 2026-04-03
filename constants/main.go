package main

type cache [100](map[string]string)

func (d *cache) storeCd(Id int, name string) {
	(*d)[string(Id)] = name
}
