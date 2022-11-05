package note_v1

import (
	"context"
	"fmt"

	desc "github.com/almira-galeeva/Note-Service-API/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("Note Was Created.")
	fmt.Println("Title:", req.GetTitle())
	fmt.Println("Text:", req.GetText())
	fmt.Printf("Author: %s\n\n", req.GetAuthor())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}
