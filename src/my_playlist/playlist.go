// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	PLS_TYPE   = "pls"
	M3U_TYPE   = "m3u"
	M3U_SUFFIX = "." + M3U_TYPE
	PLS_SUFFIX = "." + PLS_TYPE
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

func main() {
	filename := os.Args[1]
	if len(os.Args) == 1 ||
		(!strings.HasSuffix(filename, M3U_SUFFIX) &&
			!strings.HasSuffix(filename, PLS_SUFFIX)) {
		fmt.Printf("usage: %s <file.[pls|m3u]>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawBytes, err := ioutil.ReadFile(filename); err != nil {
		log.Fatal(err)
	} else {
		songs := readPlaylist(ingressPlaylistType(filename), string(rawBytes))
		writePlaylist(ingressPlaylistType(filename), songs)
	}
}

func readPlaylist(playlistType string, data string) (songs []Song) {
	if playlistType == M3U_TYPE {
		songs = readM3uPlaylist(data)
	} else {
		songs = readPlsPlaylist(data)
	}
	return
}

func writePlaylist(playlistType string, songs []Song) {
	if playlistType == M3U_TYPE {
		writePlsPlaylist(songs)
	} else {
		fmt.Println(songs)
		// M3U writing
	}
}

func ingressPlaylistType(filename string) string {
	if strings.HasSuffix(filename, PLS_SUFFIX) {
		return PLS_TYPE
	} else {
		return M3U_TYPE
	}
}

func readM3uPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlatformDirSeparator, line)
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func readPlsPlaylist(data string) (songs []Song) {
	var song Song
	lines := strings.Split(data, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i%3 != 1 || line == "" || line == "[playlist]" ||
			strings.HasPrefix(line, "NumberOfEntries") ||
			strings.HasPrefix(line, "Version") {
			continue
		}
		unmappedFilename := strings.Split(line, "=")[1]
		song.Filename = strings.Map(mapPlatformDirSeparator, unmappedFilename)
		song.Title = strings.Split(lines[i+1], "=")[1]
		song.Seconds, _ = strconv.Atoi(strings.Split(lines[i+2], "=")[1])
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func parseExtinfLine(line string) (title string, seconds int) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]
			var err error
			if seconds, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("failed to read the duration for '%s': %v\n",
					title, err)
				seconds = -1
			}
		}
	}
	return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func writePlsPlaylist(songs []Song) {
	fmt.Println("[playlist]")
	for i, song := range songs {
		i++
		fmt.Printf("File%d=%s\n", i, song.Filename)
		fmt.Printf("Title%d=%s\n", i, song.Title)
		fmt.Printf("Length%d=%d\n", i, song.Seconds)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}
