package gout

type Error struct {
	Err  error
	Type uint64
	Meta interface{}
}

type errorMsgs []*Error
