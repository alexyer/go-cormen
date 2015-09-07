package trees

type Tree interface {
	Delete()
	Find(val int) Tree
	Insert(val int) Tree
}
