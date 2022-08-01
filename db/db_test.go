package db_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/hertg/gopci/db"
	"github.com/hertg/gopci/internal/utils"
	"github.com/jaypipes/pcidb"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	f, err := os.Open("/home/michael/repos/gopci/pci.ids")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	database := db.DB{}
	db.Parse(scanner, &database)

	fmt.Println(database.Vendors[0x1002].Label)
	for _, d := range database.Vendors[0x1002].Devices {
		fmt.Println("\t", d.Label)
	}
	fmt.Printf("%+v\n", database.Devices[0x73bf])

	//fmt.Printf("%+v", database)

	//panic("err")
}

func TestParseByteNum(t *testing.T) {
	num := utils.ParseByteNum([]byte("2f"))
	assert.Equal(t, uint(47), num)

	num = utils.ParseByteNum([]byte("f4ab"))
	assert.Equal(t, uint(62635), num)

	num = utils.ParseByteNum([]byte("c48118"))
	assert.Equal(t, uint(12878104), num)

	num = utils.ParseByteNum([]byte("df29fe6e"))
	assert.Equal(t, uint(3744071278), num)
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.ParseByteNum([]byte("df29fe6e"))
	}
}

func BenchmarkDBInit(b *testing.B) {
	f, err := os.Open("/home/michael/repos/gopci/pci.ids")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		scanner := bufio.NewScanner(f)
		d := db.DB{}
		db.Parse(scanner, &d)
	}
}

func BenchmarkGhwInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pcidb.New()
	}
}

/*
func parseByteNumOld(str []byte) uint {
	// len: 2->8, 4->16, 8->32, 16->64
	shift := (len(str) - 1) * 4
	var num uint
	for _, s := range str {
		if s >= 97 && s <= 102 {
			s = (s+1)&0b0111 | 0b1000
		} else if s >= 48 && s <= 57 {
			s = s & 0b1111
		}
		num = num | (uint(s) << shift)
		shift -= 4
	}
	return num
}

func parseByteNumDirty(str []byte) uint {
	s := string(str)
	n, _ := strconv.ParseInt(s, 16, 32)
	return uint(n)
}

func BenchmarkOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseByteNumOld([]byte("df29fe6e"))
	}
}

func BenchmarkDirty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseByteNumDirty([]byte("df29fe6e"))
	}
}
*/
