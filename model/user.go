package model

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

type User struct {
	ID        objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UUID      uuid.UUID         `bson:"uuid,omitempty" json:"uuid,omitempty"`
	Email     string            `bson:"email,omitempty" json:"email,omitempty"`
	FirstName string            `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string            `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Username  string            `bson:"username,omitempty" json:"username,omitempty"`
	Password  string            `bson:"password,omitempty" json:"password,omitempty"`
	Role      string            `bson:"role,omitempty" json:"role,omitempty"`
}

// marshalUser is a simple
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
	if u.UUID.String() != (uuid.UUID{}).String() {
		mu.UUID = u.UUID.String()
	}

	return json.Marshal(mu)
}

func (u *User) UnmarshalJSON(in []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	u.ID, err = objectid.FromHex(m["_id"].(string))
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error: Error parsing user _id")
		return err
	}

	u.UUID, err = uuid.FromString(m["uuid"].(string))
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
