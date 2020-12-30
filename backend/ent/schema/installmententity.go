package schema

import "github.com/facebookincubator/ent"

// InstallmentEntity holds the schema definition for the InstallmentEntity entity.
type InstallmentEntity struct {
	ent.Schema
}

// Fields of the InstallmentEntity.
func (InstallmentEntity) Fields() []ent.Field {
	return nil
}

// Edges of the InstallmentEntity.
func (InstallmentEntity) Edges() []ent.Edge {
	return nil
}
