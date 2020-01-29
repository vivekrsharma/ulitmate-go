package chapter3

const (
	rows int = 2 * 1024
	cols int = 2 * 1024
)

var matrix [rows][cols]byte

// LinkedList to compare traversal speed.
type data struct {
	v byte
	n *data
}
var list *data

func rowWiseTraversal() int {
	var count int
	for i:= 0; i< rows;i++ {
		for j := 0; j< cols; j++ {
			if matrix[i][j] == 0xff {
				count++
			}
		}
	}
	return count
}

func columnWiseTraversal() int {
	var count int

	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			if matrix[i][j] == 0xff {
				count++
			}
		}
	}
	return count
}

func linkedListTraversal() int {
	var count int
	for itr:= list; itr!=nil; itr = itr.n{
		if itr.v == 0xff {
		count++
		}
	}
	return count
}

func init() {
	var last *data
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var d data
			if i%2 == 0 {
				matrix[i][j] = 0xff
				d.v = 0xff
			} else {
				d.v = 0x00
			}

			if list == nil {
				list = &d
			} else {
				last.n = &d
			}

			last = &d
		}
	}
}

func main() {

}

