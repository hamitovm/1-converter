package bins

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool, createdAt time.Time, name string) Bin {
	return Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}

type BinList []Bin

func NewBinList(bins ...Bin) BinList {
	return BinList(bins)
}
