package btree


type Btree struct{
	Data int
	Lchild *Btree
	Rchild *Btree
}

func CreateBtree(val int) *Btree {
	node := new(Btree)
	node.Data = val
	node.Lchild = nil
	node.Rchild = nil
	return node
}

func (t *Btree) GetData() int{
	return t.Data
}

func (t *Btree) SetData(val int){
	t.Data = val
}

func (t *Btree) FindNode(node *Btree, x int) *Btree {
	if node == nil{
		return nil
	}
	if node.Data == x {
		return node
	} else if node.Data > x {
		return t.FindNode(node.Lchild, x)
	} else {
		return t.FindNode(node.Rchild, x)
	}
}

func (t *Btree) GetHeight(node *Btree) int {
	if node == nil {
		return 0
	} else{
		lheight := t.GetHeight(node.Lchild)
		rheight := t.GetHeight(node.Rchild)
		if lheight >= rheight {
			return lheight + 1
		} else {
			return rheight + 1
		}
	}
}

// func (t )