/*
tsuserver, an Attorney Online server
Copyright (C) 2016 tsukasa84 <tsukasadev84@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/hex"
	"errors"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var crypt_const1 uint16 = 53761
var crypt_const2 uint16 = 32618
var crypt_key uint16

var rng *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func calcKeys() {
	var tmp int
	tmp, _ = strconv.Atoi(decryptMessage([]byte("4"), 322))
	crypt_key = uint16(tmp)
}

func hexToBytes(hexstr string) []byte {
	out, _ := hex.DecodeString(hexstr)
	return out
}

func decryptMessage(enc_msg []byte, key uint16) string {
	// "crypt"
	var out string
	for _, val := range enc_msg {
		out += string(uint16(val) ^ (key >> 8))
		key = ((uint16(val) + key) * crypt_const1) + crypt_const2
	}
	return out
}

func encryptMessage(pt string, key uint16) string {
	// "crypt"
	var out string
	for _, chr := range pt {
		val := uint16(chr) ^ (key >> 8)
		out += strconv.FormatUint(uint64(val), 16)
		key = ((val + key) * crypt_const1) + crypt_const2
	}
	return out
}

func loadCharPages(perpage int) []string {
	var ret []string
	var str string = "CI#"

	for i, v := range config.Charlist {
		str += strconv.Itoa(i) + "#" + v + "&None&0&&&0&#"
		if math.Mod(float64(i), float64(perpage)) == float64(perpage-1) {
			str += "#%"
			ret = append(ret, str)
			str = "CI#"
		}
	}

	if len(str) > 3 {
		str += "#%"
		ret = append(ret, str)
	}

	return ret
}

func loadEvidence() []string {
	var ret []string
	var str string = "EI#"
	var cnt int
	for i, v := range config.Evidencelist {
		str += strconv.Itoa(i+1) + "#" + v.Name + "&" + v.Desc + "&" + v.Type + "&" + v.Image + "&##%"
		ret = append(ret, str)
		str = "EI#"
		cnt = i + 1
	}
	cnt++
	for _, v := range cust.Evidencelist {
		str += strconv.Itoa(cnt) + "#" + v.Name + "&" + v.Desc + "&" + v.Type + "&" + v.Image + "&##%"
		ret = append(ret, str)
		str = "EI#"
		cnt++
	}

	return ret
}

func loadMusicPages(perpage int) []string {
	var ret []string
	var str string = "EM#"

	for i, v := range config.Musiclist {
		str += strconv.Itoa(i) + "#" + v.Name + "#"
		if math.Mod(float64(i), float64(perpage)) == float64(perpage-1) {
			str += "#%"
			ret = append(ret, str)
			str = "EM#"
		}
	}

	if len(str) > 3 {
		str += "#%"
		ret = append(ret, str)
	}

	return ret
}

func isValidCharID(id int) bool {
	return id >= 0 && id < len(config.Charlist)
}

func isPosValid(pos string) bool {
	validpos := map[string]bool{"def": true, "pro": true, "hld": true,
		"hlp": true, "wit": true, "jud": true}
	if _, ok := validpos[pos]; ok {
		return true
	}
	return false
}

func stringInSlice(a string, list []string, case_sensitive bool) (string, error) {
	for _, b := range list {
		if case_sensitive {
			if b == a {
				return b, nil
			}
		} else {
			if strings.ToLower(b) == strings.ToLower(a) {
				return b, nil
			}
		}
	}
	return "", errors.New("String not found.")
}

func getCIDfromName(charname string) (int, error) {
	if len(charname) == 0 {
		return -1, errors.New("Empty character name.")
	}

	for i, c := range config.Charlist {
		if strings.ToLower(c) == strings.ToLower(charname) {
			return i, nil
		}
	}
	return -1, errors.New("Character could not be found")
}

func getAreaPtr(areaid int) *Area {
	for i := range config.Arealist {
		if config.Arealist[i].Areaid == areaid {
			return &config.Arealist[i]
		}
	}
	return nil
}

// finds whether a string starts with a character name
// if yes, returns the character name and the rest of the message
// if not, returns an error
func msgStartsWithChar(str string) (string, string, error) {
	if len(str) == 0 {
		return "", "", errors.New("Empty string.")
	}

	for _, v := range config.Charlist {
		if strings.HasPrefix(strings.ToLower(str), strings.ToLower(v)) {
			wordcount := len(strings.Split(v, " "))
			split_str := strings.Split(str, " ")

			return v, strings.Join(split_str[wordcount:], " "), nil
		}
	}

	return "", "", errors.New("Character name not found.")
}

// checks if OOC name is reserved
func isOOCNameReserved(name string) bool {
	if strings.HasPrefix(name, config.Reservedname) {
		return true
	}

	if strings.HasPrefix(name, "<dollar>GLOBAL") {
		return true
	}

	if strings.HasPrefix(name, "<dollar>MOD") {
		return true
	}

	return false
}

// checks if file exists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// creates given file
func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fo.Close()
	return nil
}

// Generates a number between min and max
func randomInt(min int, max int) int {
	result := min + rand.Intn(max)
	return result
}

func (evi *Customevidence) AddEvidence(entry Evidence) []Evidence {
	evi.Evidencelist = append(evi.Evidencelist, entry)
	return evi.Evidencelist
}
