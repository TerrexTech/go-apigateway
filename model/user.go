package model

import (
	"encoding/json"
	"log"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

// User represents a system-registered user.
type User struct {
	ID        objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UUID      uuuid.UUID        `bson:"uuid,omitempty" json:"uuid,omitempty"`
	Email     string            `bson:"email,omitempty" json:"email,omitempty"`
	FirstName string            `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string            `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Username  string            `bson:"username,omitempty" json:"username,omitempty"`
	Password  string            `bson:"password,omitempty" json:"password,omitempty"`
	Role      string            `bson:"role,omitempty" json:"role,omitempty"`
}

// marshalUser is alternative format for User for convenient
// Marshalling/Unmarshalling operations.
type marshalUser struct {
	ID        objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UUID      string            `bson:"uuid,omitempty" json:"uuid,omitempty"`
	Email     string            `bson:"email,omitempty" json:"email,omitempty"`
	FirstName string            `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string            `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Username  string            `bson:"username,omitempty" json:"username,omitempty"`
	Password  string            `bson:"password,omitempty" json:"password,omitempty"`
	Role      string            `bson:"role,omitempty" json:"role,omitempty"`
}

// MarshalBSON converts the User to its BSON representation.
func (u User) MarshalBSON() ([]byte, error) {
	mu := &marshalUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		Role:      u.Role,
	}

	if u.UUID.String() != (uuuid.UUID{}).String() {
		mu.UUID = u.UUID.String()
	}
	return bson.Marshalv2(mu)
}

// MarshalJSON converts the current-user representation into its
// JSON representation.
func (u *User) MarshalJSON() ([]byte, error) {
	mu := &marshalUser{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		Role:      u.Role,
	}

	mu.ID = u.ID
	mu.UUID = u.UUID.String()

	return json.Marshal(mu)
}

// UnmarshalJSON converts the JSON representation of a User into the
// local User-struct.
func (u *User) UnmarshalJSON(in []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	log.Println(string(in))

	u.ID, err = objectid.FromHex(m["_id"].(string))
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error: Error parsing user _id")
		return err
	}

	u.UUID, err = uuuid.FromString(m["uuid"].(string))
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error: Error parsing user uuid")
		return err
	}
	u.Email = m["email"].(string)
	u.FirstName = m["first_name"].(string)
	u.LastName = m["last_name"].(string)
	u.Username = m["username"].(string)

	if m["password"] != nil {
		u.Password = m["password"].(string)
	}
	u.Role = m["role"].(string)

	return nil
}
