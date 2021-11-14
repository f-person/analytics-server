package infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

// Creates `primitive.ObjectID` from the given string id.
// `id` is assumed to be valid.
func objectIDFromString(id string) primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return objectID
}
