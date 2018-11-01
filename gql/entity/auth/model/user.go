package model

import (
	"encoding/json"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

// AggregateID is the global AggregateID for UserAuth Aggregate.
const AggregateID int8 = 1

// User defines the User Aggregate.
type User struct {
	ID        objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID    uuuid.UUID        `bson:"userID,omitempty" json:"userID,omitempty"`
	Email     string            `bson:"email,omitempty" json:"email,omitempty"`
	FirstName string            `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string            `bson:"last_name,omitempty" json:"last_name,omitempty"`
	UserName  string            `bson:"userName,omitempty" json:"userName,omitempty"`
	Password  string            `bson:"password,omitempty" json:"password,omitempty"`
	Role      string            `bson:"role,omitempty" json:"role,omitempty"`
}

// MarshalJSON returns bytes of JSON-type.
func (u *User) MarshalJSON() ([]byte, error) {
	mu := map[string]interface{}{
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"email":     u.Email,
		"userName":  u.UserName,
		"password":  u.Password,
		"role":      u.Role,
		"userID":    u.UserID.String(),
	}

	if u.ID != objectid.NilObjectID {
		mu["_id"] = u.ID.Hex()
	}

	m, err := json.Marshal(mu)
	if err != nil {
		err = errors.Wrap(err, "MarshalJSON Error")
	}
	return m, err
}

// UnmarshalJSON returns JSON-type from bytes.
func (u *User) UnmarshalJSON(in []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	err = u.unmarshalFromMap(m)
	if err != nil {
		err = errors.Wrap(err, "UnmarshalJSON Error")
	}
	return err
}

// unmarshalFromMap unmarshals Map into User.
func (u *User) unmarshalFromMap(m map[string]interface{}) error {
	var err error
	var assertOK bool

	// Hoping to discover a better way to do this someday
	if m["_id"] != nil {
		u.ID, assertOK = m["_id"].(objectid.ObjectID)
		if !assertOK {
			u.ID, err = objectid.FromHex(m["_id"].(string))
			if err != nil {
				err = errors.Wrap(err, "Error while asserting ObjectID")
				return err
			}
		}
	}

	if m["userID"] != nil {
		u.UserID, err = uuuid.FromString(m["userID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting UserID")
			return err
		}
	}

	if m["email"] != nil {
		u.Email, assertOK = m["email"].(string)
		if !assertOK {
			return errors.New("Error while asserting Email")
		}
	}
	if m["firstName"] != nil {
		u.FirstName, assertOK = m["firstName"].(string)
		if !assertOK {
			return errors.New("Error while asserting FirstName")
		}
	}
	if m["lastName"] != nil {
		u.LastName, assertOK = m["lastName"].(string)
		if !assertOK {
			return errors.New("Error while asserting LastName")
		}
	}
	if m["userName"] != nil {
		u.UserName, assertOK = m["userName"].(string)
		if !assertOK {
			return errors.New("Error while asserting Username")
		}
	}
	if m["password"] != nil {
		u.Password, assertOK = m["password"].(string)
		if !assertOK {
			return errors.New("Error while asserting Password")
		}
	}
	if m["role"] != nil {
		u.Role, assertOK = m["role"].(string)
		if !assertOK {
			return errors.New("Error while asserting Role")
		}
	}

	return nil
}
