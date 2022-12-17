package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var goalsBucket = []byte("GoalsBucket")
var db *bolt.DB

type Goal struct {
	Id     int
	Text   string
	Repeat int
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(goalsBucket)
		return err
	})
}

func CreateGoal(goalText string, repeat int) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(goalsBucket)
		id, _ := b.NextSequence()
		goal := &Goal{
			Id:     int(id),
			Text:   goalText,
			Repeat: repeat,
		}
		// Marshal user data into bytes.
		buf, err := json.Marshal(goal)
		if err != nil {
			return err
		}
		return b.Put(itob(goal.Id), buf)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetAllGoals() ([]Goal, error) {
	var goals []Goal
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(goalsBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var goal Goal
			err := json.Unmarshal(v, &goal)
			if err != nil {
				fmt.Println("error:", err)
			}
			goals = append(goals, goal)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return goals, nil
}

func DeleteGoal(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(goalsBucket)
		return b.Delete(itob(id))
	})
}

// edit features of a goal
func UpdateGoal(id int, goalText string, repeat int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(goalsBucket)
		goal := &Goal{
			Id:     id,
			Text:   goalText,
			Repeat: repeat,
		}
		// Marshal user data into bytes.
		buf, err := json.Marshal(goal)
		if err != nil {
			return err
		}
		return b.Put(itob(goal.Id), buf)
	})

	return err
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
