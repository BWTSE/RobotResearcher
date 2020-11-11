package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Scenario struct {
	HasDebt      bool              `bson:"has_debt"`
	StartedAt    time.Time         `bson:"started_at"`
	SubmittedAt  time.Time         `bson:"submitted_at"`
	Instructions string            `bson:"instructions"`
	Starting     map[string]string `bson:"starting"`
	Submitted    map[string]string `bson:"submitted"`
}

type SurveySubmission struct {
	LikesChocolate bool `bson:"likes_chocolate"`
}

type Session struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	RegisterCode      string             `bson:"register_code"`
	AgreementAccepted bool               `bson:"agreement_accepted"`
	SurveyAnswers     *SurveySubmission  `bson:"survey_answers"`
	Scenarios         []Scenario         `bson:"scenarios"`
	StartedAt         time.Time          `bson:"started_at"`
	EndedAt           time.Time          `bson:"ended_at"`
}

func (d *Database) GetSession(id primitive.ObjectID) (Session, error) {
	cursor := d.client.Database("test").Collection("sessions").FindOne(context.TODO(), bson.M{"_id": id})

	var session Session

	err := cursor.Decode(&session)
	if err != nil {
		return session, err
	}

	return session, nil
}

func (d *Database) CreateSession(code string, scenarios []Scenario) (primitive.ObjectID, error) {
	session := Session{
		RegisterCode: code,
		Scenarios:    scenarios,
		StartedAt:    time.Now(),
	}

	cursor, err := d.client.Database("test").Collection("sessions").InsertOne(context.TODO(), session)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return cursor.InsertedID.(primitive.ObjectID), nil
}

func (d *Database) AcceptAgreement(id primitive.ObjectID) error {
	res, err := d.client.Database("test").Collection("sessions").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"agreement_accepted", true}}},
		},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("nothing modified")
	}

	return nil
}

func (d *Database) EndSession(id primitive.ObjectID) error {
	res, err := d.client.Database("test").Collection("sessions").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"ended_at", time.Now()}}},
		},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("nothing modified")
	}

	return nil
}

func (d *Database) OpenScenario(id primitive.ObjectID, n int) error {
	res, err := d.client.Database("test").Collection("sessions").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"scenarios." + strconv.Itoa(n) + ".started_at", time.Now()}}},
		},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("nothing modified")
	}

	return nil
}

func (d *Database) CommitScenario(id primitive.ObjectID, n int, submission map[string]string) error {
	res, err := d.client.Database("test").Collection("sessions").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"scenarios." + strconv.Itoa(n) + ".submitted_at", time.Now()}}},
			{"$set", bson.D{{"scenarios." + strconv.Itoa(n) + ".submitted", submission}}},
		},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("nothing modified")
	}

	return nil
}

func (d *Database) SaveSurveyAnswers(id primitive.ObjectID, submission SurveySubmission) error {
	res, err := d.client.Database("test").Collection("sessions").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"survey_answers", submission}}},
		},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("nothing modified")
	}

	return nil
}
