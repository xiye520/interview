// skipList project skipList.go
package kit

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SkipListObj interface {
	Compare(obj SkipListObj) bool
	PrintObj()
}

type Node struct {
	O        SkipListObj
	forward  []*Node
	curLevel int
}

type SkipList struct {
	head     *Node
	length   int
	maxLevel int
	lockType int
	lock     sync.Locker
}

//lockType = 0, no lock
//lockType = 1, Mutex
//lockType = 2, RWMutex
// mode = 1 : exclusive lock
// mode = 2 : if lockType = 1, it is exclusive lock; if lockType = 2, it is shared lock
func (s *SkipList) lockList(mode int) {
	if s.lockType == 0 {
		return
	}
	switch mode {
	case 1:
		s.lock.Lock()
	case 2:
		if s.lockType == 1 {
			s.lock.Lock()
		} else if s.lockType == 2 {
			s.lock.(*sync.RWMutex).RLock()
		}
	default:
		return
	}
}

func (s *SkipList) unLockList(mode int) {
	if s.lockType == 0 {
		return
	}
	switch mode {
	case 1:
		s.lock.Unlock()
	case 2:
		if s.lockType == 1 {
			s.lock.Unlock()
		} else if s.lockType == 2 {
			s.lock.(*sync.RWMutex).RUnlock()
		}
	default:
		return
	}
}

//try to find the first node which not match the user defined Compare() condition
func (s *SkipList) searchInternal(o SkipListObj) (*Node, error) {
	p := s.head

	for i := s.maxLevel - 1; i >= 0; i-- {
		for {
			if p.forward[i] != nil && p.forward[i].O.Compare(o) {
				p = p.forward[i]
			} else {
				break
			}
		}
	}
	p = p.forward[0]
	if p == nil {
		return nil, errors.New("No matched object in skip list")
	} else {
		return p, nil
	}
}

func (s *SkipList) Search(o SkipListObj) (SkipListObj, error) {
	v, err := checkSkipListValid(s)
	if v == false {
		return o, err
	}

	s.lockList(2)
	defer s.unLockList(2)
	res, err := s.searchInternal(o)
	if err == nil {
		if !res.O.Compare(o) && !o.Compare(res.O) {
			return res.O, nil
		} else {
			return o, errors.New("cannot find object in skip list")
		}
	} else {
		return o, err
	}
}

func (s *SkipList) SearchRange(minObj, maxObj SkipListObj) ([]SkipListObj, error) {
	res := make([]SkipListObj, 0)
	v, err := checkSkipListValid(s)
	if v == false {
		return res, err
	}

	s.lockList(2)
	defer s.unLockList(2)

	p, err := s.searchInternal(minObj)
	if err != nil {
		return res, err
	}

	for {
		if p != nil && p.O.Compare(maxObj) {
			res = append(res, p.O)
			p = p.forward[0]
		} else {
			break
		}
	}

	return res, nil
}

func (s *SkipList) Traverse() {
	v, _ := checkSkipListValid(s)
	if v == false {
		return
	}

	var p *Node = s.head

	s.lockList(2)
	defer s.unLockList(2)

	for i := s.maxLevel - 1; i >= 0; i-- {
		for {
			if p != nil {
				p.O.PrintObj()
				if p.forward[i] != nil {
					fmt.Print("-->")
				}
				p = p.forward[i]
			} else {
				break
			}
		}
		fmt.Println()
		p = s.head
	}
}

func (s *SkipList) Insert(obj SkipListObj) (bool, error) {
	v, err := checkSkipListValid(s)
	if v == false {
		return false, err
	}

	var p *Node = s.head
	newNode := new(Node)
	newNode.O = obj
	newNode.forward = make([]*Node, s.maxLevel)
	level := s.createNewNodeLevel()

	s.lockList(1)
	defer s.unLockList(1)

	for i := s.maxLevel - 1; i >= 0; i-- {
		for {
			if p.forward[i] != nil && p.forward[i].O.Compare(obj) {
				p = p.forward[i]
			} else {
				break
			}
		}
		//find the last Node which match user defined Compare() condition in i level
		//insert new Node after the node
		if i <= level {
			newNode.forward[i] = p.forward[i]
			p.forward[i] = newNode
		}
	}
	newNode.curLevel = level
	s.length++

	return true, nil
}

func (s *SkipList) RemoveNode(obj SkipListObj) (bool, error) {
	v, err := checkSkipListValid(s)
	if v == false {
		return false, err
	}

	var update []*Node = make([]*Node, s.maxLevel)
	p := s.head

	s.lockList(1)
	defer s.unLockList(1)

	for i := s.maxLevel - 1; i >= 0; i-- {
		for {
			if p.forward[i] != nil && p.forward[i].O.Compare(obj) {
				p = p.forward[i]
			} else {
				break
			}
		}
		update[i] = p
	}
	p = p.forward[0]

	if p == nil || p.O.Compare(obj) || obj.Compare(p.O) {
		return false, errors.New("cannot find object")
	}

	for i := p.curLevel; i >= 0; i-- {
		update[i].forward[i] = p.forward[i]
	}
	s.length--

	return true, nil
}

func (s *SkipList) ClearSkipList() error {
	v, err := checkSkipListValid(s)
	if v == false {
		return err
	}

	s.lockList(1)
	defer s.unLockList(1)

	for i := s.maxLevel; i >= 0; i-- {
		s.head.forward[i] = nil
	}
	s.length = 0

	return nil
}

func (s *SkipList) LenOfSkipList() (int, error) {
	v, err := checkSkipListValid(s)
	if v == false {
		return -1, err
	}

	s.lockList(2)
	defer s.unLockList(2)

	return s.length, nil
}

func CreateSkipList(minObj SkipListObj, args ...int) (*SkipList, error) {
	if minObj == nil {
		return nil, errors.New("minObj paramter is null")
	}
	var maxlevel, mode int
	for i, v := range args {
		switch i {
		case 0:
			maxlevel = v
		case 1:
			mode = v
		default:
			return nil, errors.New("Too many arguments, expect 2 or 3 arguments")
		}
	}
	if maxlevel <= 0 {
		return nil, errors.New("Max level of skip list must greater than 0")
	}

	s := new(SkipList)
	s.head = new(Node)
	s.maxLevel = maxlevel
	s.head.curLevel = maxlevel - 1
	s.head.forward = make([]*Node, maxlevel)
	s.head.O = minObj
	//The length of skip list didn't include the head node
	s.length = 0
	if mode == 1 {
		s.lockType = 1
		s.lock = new(sync.Mutex)
	} else if mode == 2 {
		s.lockType = 2
		s.lock = new(sync.RWMutex)
	}

	return s, nil
}

func (s *SkipList) createNewNodeLevel() int {
	var level int = 0

	rand.Seed(time.Now().UnixNano())
	for {
		if rand.Intn(2) == 1 || level >= s.maxLevel-1 {
			break
		}
		level++
	}
	return level
}

func checkSkipListValid(s *SkipList) (bool, error) {
	if s == nil {
		return false, errors.New("skip list not exist")
	}
	if s.head == nil {
		return false, errors.New("skip list head is null, must use CreateSkipList() to create skip list")
	}

	return true, nil
}
