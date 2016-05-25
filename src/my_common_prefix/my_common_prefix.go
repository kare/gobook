package main

func CommonPrefix(paths []string) string {
	numPaths := len(paths)
	for i := 0; i < minLength(paths); i++ {
		for j := 1; j < numPaths; j++ {
			if paths[j][i] != paths[0][i] {
				if i == 0 {
					return ""
				} else {
					return paths[0][:i]
				}
			}
		}
	}
	return paths[0]
}

func minLength(paths []string) int {
	var shortest int
	for i, path := range paths {
		if i == 0 || len(path) < shortest {
			shortest = len(path)
		}
	}
	return shortest
}
