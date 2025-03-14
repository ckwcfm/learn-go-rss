package constants

type contextKey string

const UserIDKey contextKey = "userID"
const UserKey contextKey = "user"

// func GetUserID(r *http.Request) (primitive.ObjectID, error) {
// 	userID := r.Context().Value(UserIDKey)
// 	if userID == nil {
// 		return primitive.NilObjectID, errors.New("userID not found")
// 	}
// 	id, err := primitive.ObjectIDFromHex(userID.(string))
// 	if err != nil {
// 		return primitive.NilObjectID, err
// 	}
// 	return id, nil
// }
