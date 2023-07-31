package parser

type tokenListNode struct {
	token *Token
	next  *tokenListNode
	prev  *tokenListNode
}

type tokenList struct {
	head *tokenListNode
	tail *tokenListNode
}

func (self *tokenList) Append(t *Token) {
	newElement := &tokenListNode{token: t, next: nil, prev: nil}
	if self.head == nil {
		self.head = newElement
		self.tail = newElement
	} else {
		self.tail.next = newElement
		newElement.prev = self.tail
		self.tail = newElement
	}
}

func (self *tokenList) GetIterator() *tokenIterator {
	return &tokenIterator{
		pos: &tokenListNode{
			token: nil,
			next:  self.head,
			prev:  nil,
		},
	}
}

type tokenIterator struct {
	pos *tokenListNode
}

func (self *tokenIterator) Get() *Token {
	return self.pos.token
}

func (self *tokenIterator) Next() bool {
	if self.pos != nil && self.pos.next != nil {
		self.pos = self.pos.next
		return true
	} else {
		return false
	}
}
