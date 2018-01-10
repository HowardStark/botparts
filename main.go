package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"errors"
)

const (
	TypeUser = "user"
	TypeComment = "comment"
	TypeMedia = "media"
	TypePart = "part"

	ActionCreate = "create"
	ActionRequest = "request"
	ActionUpdate = "update"
	ActionSearch = "search"
)

var (
	errPieceTypeInvalid = errors.New("error piece type invalid")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleRequest)
	http.Handle("/api/", r)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

}

type Request struct {
	Auth string `json:"auth"`
	Pieces []*Piece `json:"pieces"`
}

type Piece struct {
	Type   string      `json:"type"`
	Action string      `json:"action"`
	Expand bool        `json:"expand"`
	Object interface{} `json:"object"`
}

func (p *Piece) UnmarshalJSON(b []byte) error {
	tp := struct {
		Type   string                 `json:"type"`
		Action string                 `json:"action"`
		Expand bool                   `json:"expand"`
		Misc   map[string]interface{} `json:"-"`
	}{}
	if err := json.Unmarshal(b, &tp); err != nil {
		return err
	}
	p.Type = tp.Type
	p.Action = tp.Action
	p.Expand = tp.Expand
	_, ok := tp.Misc["object"]
	if !ok {
		return nil
	}
	switch tp.Type {
	case TypeComment:
		ot := struct {
			Object Comment `json:"object"`
		}{}
		if err := json.Unmarshal(b, &ot); err != nil {
			return err
		}
		p.Object = ot.Object
		break
	case TypePart:
		ot := struct {
			Object Part `json:"object"`
		}{}
		if err := json.Unmarshal(b, &ot); err != nil {
			return err
		}
		p.Object = ot.Object
		break
	case TypeMedia:
		ot := struct {
			Object Media `json:"object"`
		}{}
		if err := json.Unmarshal(b, &ot); err != nil {
			return err
		}
		p.Object = ot.Object
		break
	case TypeUser:
		ot := struct {
			Object User `json:"object"`
		}{}
		if err := json.Unmarshal(b, &ot); err != nil {
			return err
		}
		p.Object = ot.Object
	default:
		return errPieceTypeInvalid
	}
	return nil
}