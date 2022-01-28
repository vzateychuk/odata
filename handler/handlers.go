package handler

import (
	"context"
	"encoding/json"
	"github.com/vez/odata/meta"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var ctx = context.TODO()

type Handler struct {
	Metas *mongo.Collection
}

func NewHandlerInstance(collection *mongo.Collection) *Handler {
	handlers := &Handler{
		Metas: collection,
	}
	return handlers
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	data, err := h.filterMetas(filter) // []*Metadata{}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) filterMetas(filter interface{}) ([]*meta.Metadata, error) {

	// A slice of metas for storing the decoded documents
	var metas []*meta.Metadata

	cur, err := h.Metas.Find(ctx, filter)
	if err != nil {
		return metas, err
	}

	for cur.Next(ctx) {
		var t meta.Metadata
		err := cur.Decode(&t)
		if err != nil {
			return metas, err
		}

		metas = append(metas, &t)
	}

	if err := cur.Err(); err != nil {
		return metas, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(metas) == 0 {
		return metas, mongo.ErrNoDocuments
	}

	return metas, nil
}
