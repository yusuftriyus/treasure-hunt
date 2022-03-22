package treasure

type Arr [][]string

func MarkUp(arr Arr, i, j int) {
	if arr[i][j] == "#" {
		return
	}
	if arr[i][j] == "." {
		arr[i][j] = "$"
	}
	MarkUp(arr, i-1, j)
	MarkRight(arr, i, j+1)
}

func MarkRight(arr Arr, i, j int) {
	if arr[i][j] == "#" {
		return
	}
	if arr[i][j] == "." {
		arr[i][j] = "$"
	}
	MarkRight(arr, i, j+1)
	MarkDown(arr, i+1, j)
}

func MarkDown(arr Arr, i, j int) {
	if arr[i][j] == "#" {
		return
	}
	if arr[i][j] == "." {
		arr[i][j] = "$"
	}
	MarkDown(arr, i+1, j)
}
