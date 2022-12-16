package model

type Storage interface {
	Add(data int64) (index int64)
	Delete(index int64) (ok bool)
	Print()
	SortInc()
	SortDec()
}

func Add(storage Storage, data int64) {
	storage.Add(data)
}

func Delete(storage Storage, index int64) {
	storage.Delete(index)
}

func Print(storage Storage) {
	storage.Print()
}

func SortInc(storage Storage) {
	storage.SortInc()
}

func SortDec(storage Storage) {
	storage.SortDec()
}
