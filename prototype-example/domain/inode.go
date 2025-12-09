package domain

type Inode interface {
	Print(indent string)
	Clone() Inode
}
