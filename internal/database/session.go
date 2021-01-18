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
	HasHighDebt  bool              `bson:"has_high_debt"`
	StartedAt    time.Time         `bson:"started_at"`
	SubmittedAt  time.Time         `bson:"submitted_at"`
	Instructions string            `bson:"instructions"`
	Starting     map[string]string `bson:"starting"`
	Submitted    map[string]string `bson:"submitted"`
	Name         string            `bson:"name"`
}

type SurveySubmission struct {
	EducationLevel        string `bson:"education_level"`
	EducationField        string `bson:"education_field"`
	ProgrammingExperience string `bson:"programming_experience"`
	JavaExperience        string `bson:"java_experience"`
	WorkDomain            string `bson:"work_domain"`

	CompanyUsesCodeReviews     bool `bson:"company_uses_code_reviews"`
	CompanyUsesPairProgramming bool `bson:"company_uses_pair_programming"`
	CompanyTracksTechnicalDebt bool `bson:"company_tracks_technical_debt"`

	ScenarioShapesQuality       int `bson:"shapes_quality"`
	ScenarioShapesQualityChange int `bson:"shapes_change_quality"`

	ScenarioBookingQuality       int `bson:"booking_quality"`
	ScenarioBookingQualityChange int `bson:"booking_change_quality"`

	ScenarioShoppingQuality       int `bson:"shopping_quality"`
	ScenarioShoppingQualityChange int `bson:"shopping_change_quality"`
}

type Session struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	RegisterCode      string             `bson:"register_code"`
	IgnoreCount       bool               `bson:"ignore_count"`
	AgreementAccepted bool               `bson:"agreement_accepted"`
	SurveyAnswers     *SurveySubmission  `bson:"survey_answers"`
	Scenarios         []Scenario         `bson:"scenarios"`
	StartedAt         time.Time          `bson:"started_at"`
	EndedAt           time.Time          `bson:"ended_at"`
}

func (d *Database) CountSessions() (int, error) {
	n, err := d.client.Database("test").Collection("sessions").CountDocuments(context.TODO(), bson.M{"ignore_count": false})
	return int(n), err
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

func (d *Database) CreateSession(code string, scenarios []Scenario, ignoreCount bool) (primitive.ObjectID, error) {
	session := Session{
		RegisterCode: code,
		IgnoreCount:  ignoreCount,
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
			{"$set", bson.D{{"agreement_accepted", true}}}, //nolint
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
			{"$set", bson.D{{"ended_at", time.Now()}}}, //nolint
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
			{"$set", bson.D{{"scenarios." + strconv.Itoa(n) + ".started_at", time.Now()}}}, //nolint
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
			{"$set", bson.D{{"scenarios." + strconv.Itoa(n) + ".submitted_at", time.Now()}}}, //nolint
			{"$set", bson.D{{"scenarios." + strconv.Itoa(n) + ".submitted", submission}}}, //nolint
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
			{"$set", bson.D{{"survey_answers", submission}}}, //nolint
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
