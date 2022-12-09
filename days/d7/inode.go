package d7

type inode interface {
	getName() string
	getSize() int
	getParent() *dir
}
