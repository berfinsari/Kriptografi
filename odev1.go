//md5'i kullanarak rastgele sayılardan ayırt edilemeyecek bir dizi döndüren golang kodu.
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arrayLength int
	var x []string
	var y []int64

	fmt.Println("Dizinin uzunluğunu giriniz: ")
	fmt.Scanln(&arrayLength)

	file, err := os.Open("kripton")
	if err != nil {
		fmt.Println("dosya okunamıyor")
		log.Fatal(err)
	}
	defer file.Close()

	hashFile := md5.New()
	_, err = io.Copy(hashFile, file)
	if err != nil {
		log.Fatal(err)
	}
	hashByte := hashFile.Sum(nil)
	var hashmd5String string
	hashmd5String = hex.EncodeToString(hashByte)

	x = append(x, hashmd5String)
	y = append(y, returncharacter(hashmd5String))

	hsum := hashmd5String
	for i := 1; i < arrayLength; i++ {
		x = append(x, returnmd5(hsum))
		y = append(y, returncharacter(hsum))
		hsum = returnmd5(hsum)
	}
	for j := 0; j < arrayLength; j++ {
		print(y[j], " ")
	}
}
func returnmd5(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	hSum := h.Sum(nil)[:16]
	hmd5String := hex.EncodeToString(hSum)
	return hmd5String
}

func returncharacter(textmd5 string) int64 {
	var array_temp []string
	array_temp = strings.Split(textmd5, "")
	h_character, _ := strconv.ParseInt(array_temp[7], 16, 0)
	return h_character

}
