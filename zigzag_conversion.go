package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	result := ""
	if numRows == 1 {
		return s
	}

	pattern := 2*numRows - 2
	tmpPattern := pattern
	lenght := 0

	/*
		Input: s = "PAYPALISHIRING", numRows = 4
		Output: "PINALSIGYAHRPI"

		P
		A   L
		Y A
		P

		I
		S  I
		H R
		I

		N
		G

		olmak üzere bölümlere ayırdım
		bu bölümlere nasıl ayırdığımı soracak olursanız örüntüsünü bulmam gerekti formülü "2*(numRows-currentRow-1) - 2" şeklinde buldum

		stringin uzunluğunu patterne böldüğümde bana integer bir sayı verir
		unutmayın 5/2 = 2.5 değil 2'dir (integer kesirli olmadığından dolayı)

		burdaki kesir NG'yi temsil etmektedir (pattern'im 6 (2*(4-0) - 2) ama NG 2 karakterli)
		eğer kesir varsa uzunluğu bir arttırmalıyım
	*/

	if (len(s)/pattern)*pattern == len(s) {
		lenght = len(s) / pattern
	} else {
		lenght = (len(s) / pattern) + 1
	}

	var stringList []string
	/*
		********************************
		tmpPattern o anlık pattern
		bunu kullanmamın sebebi tmpPattern ve pattern eşit olmadığında 2 tane eşit olduğunda 1 tane char eklenmeli
		mesela 2. satıra geldiğimizde patternimiz 4 (2*(4-1) - 2) olacaktır

		*A   L
		Y A
		P

		*S  I
		H R
		I

		*G
		********************************
		*tmpPattern ve bir sayıyı toplayıp pattern sayısını bulmalıyız
		mesela 2. satırda tmpPattern'imiz 4 ya
		1. bölümün 2. ve 4. karakteri
		P
		*A   L
		Y A
		P

		2. bölümün 2. ve 4. karakteri
		I
		*S  I
		H R
		I

		3. bölümün ise yalnızca 2. karakteri (4. karakteri yok (bknz 95. satır))
		N
		*G

		98. satırdaki sorgulama tmpPattern 0 olduğunda 96. satırdaeklediğini tekrar eklemesin diye var
	*/
	for i := 0; i < numRows; i++ {
		for j := 0; j < lenght; j++ {
			tmpPattern = 2*(numRows-i) - 2

			if tmpPattern == pattern {
				stringList = append(stringList, string(s[(j*tmpPattern)+i]))

			} else {
				if (j*pattern)+i < len(s) {
					stringList = append(stringList, string(s[(j*pattern)+i]))
				}
				if (j*pattern)+i+(tmpPattern) < len(s) && tmpPattern != 0 {
					stringList = append(stringList, string(s[(j*pattern)+i+(tmpPattern)]))
				}

			}

		}
	}
	result = strings.Join(stringList, "")

	return result
}
