// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Question is an object representing the database table.
type Question struct {
	ID          string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description string      `boil:"description" json:"description" toml:"description" yaml:"description"`
	CreatedBy   string      `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	CreatedAt   null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	Status      null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	Rating      null.Int64  `boil:"rating" json:"rating,omitempty" toml:"rating" yaml:"rating,omitempty"`

	R *questionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L questionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuestionColumns = struct {
	ID          string
	Title       string
	Description string
	CreatedBy   string
	CreatedAt   string
	Status      string
	Rating      string
}{
	ID:          "id",
	Title:       "title",
	Description: "description",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	Status:      "status",
	Rating:      "rating",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var QuestionWhere = struct {
	ID          whereHelperstring
	Title       whereHelperstring
	Description whereHelperstring
	CreatedBy   whereHelperstring
	CreatedAt   whereHelpernull_Time
	Status      whereHelpernull_String
	Rating      whereHelpernull_Int64
}{
	ID:          whereHelperstring{field: "\"questions\".\"id\""},
	Title:       whereHelperstring{field: "\"questions\".\"title\""},
	Description: whereHelperstring{field: "\"questions\".\"description\""},
	CreatedBy:   whereHelperstring{field: "\"questions\".\"created_by\""},
	CreatedAt:   whereHelpernull_Time{field: "\"questions\".\"created_at\""},
	Status:      whereHelpernull_String{field: "\"questions\".\"status\""},
	Rating:      whereHelpernull_Int64{field: "\"questions\".\"rating\""},
}

// QuestionRels is where relationship names are stored.
var QuestionRels = struct {
	CreatedByUser string
	Answers       string
}{
	CreatedByUser: "CreatedByUser",
	Answers:       "Answers",
}

// questionR is where relationships are stored.
type questionR struct {
	CreatedByUser *User       `boil:"CreatedByUser" json:"CreatedByUser" toml:"CreatedByUser" yaml:"CreatedByUser"`
	Answers       AnswerSlice `boil:"Answers" json:"Answers" toml:"Answers" yaml:"Answers"`
}

// NewStruct creates a new relationship struct
func (*questionR) NewStruct() *questionR {
	return &questionR{}
}

// questionL is where Load methods for each relationship are stored.
type questionL struct{}

var (
	questionAllColumns            = []string{"id", "title", "description", "created_by", "created_at", "status", "rating"}
	questionColumnsWithoutDefault = []string{"id", "title", "description", "created_by"}
	questionColumnsWithDefault    = []string{"created_at", "status", "rating"}
	questionPrimaryKeyColumns     = []string{"id"}
)

type (
	// QuestionSlice is an alias for a slice of pointers to Question.
	// This should generally be used opposed to []Question.
	QuestionSlice []*Question
	// QuestionHook is the signature for custom Question hook methods
	QuestionHook func(context.Context, boil.ContextExecutor, *Question) error

	questionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	questionType                 = reflect.TypeOf(&Question{})
	questionMapping              = queries.MakeStructMapping(questionType)
	questionPrimaryKeyMapping, _ = queries.BindMapping(questionType, questionMapping, questionPrimaryKeyColumns)
	questionInsertCacheMut       sync.RWMutex
	questionInsertCache          = make(map[string]insertCache)
	questionUpdateCacheMut       sync.RWMutex
	questionUpdateCache          = make(map[string]updateCache)
	questionUpsertCacheMut       sync.RWMutex
	questionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var questionBeforeInsertHooks []QuestionHook
var questionBeforeUpdateHooks []QuestionHook
var questionBeforeDeleteHooks []QuestionHook
var questionBeforeUpsertHooks []QuestionHook

var questionAfterInsertHooks []QuestionHook
var questionAfterSelectHooks []QuestionHook
var questionAfterUpdateHooks []QuestionHook
var questionAfterDeleteHooks []QuestionHook
var questionAfterUpsertHooks []QuestionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Question) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Question) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Question) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Question) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Question) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Question) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Question) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Question) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Question) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddQuestionHook registers your hook function for all future operations.
func AddQuestionHook(hookPoint boil.HookPoint, questionHook QuestionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		questionBeforeInsertHooks = append(questionBeforeInsertHooks, questionHook)
	case boil.BeforeUpdateHook:
		questionBeforeUpdateHooks = append(questionBeforeUpdateHooks, questionHook)
	case boil.BeforeDeleteHook:
		questionBeforeDeleteHooks = append(questionBeforeDeleteHooks, questionHook)
	case boil.BeforeUpsertHook:
		questionBeforeUpsertHooks = append(questionBeforeUpsertHooks, questionHook)
	case boil.AfterInsertHook:
		questionAfterInsertHooks = append(questionAfterInsertHooks, questionHook)
	case boil.AfterSelectHook:
		questionAfterSelectHooks = append(questionAfterSelectHooks, questionHook)
	case boil.AfterUpdateHook:
		questionAfterUpdateHooks = append(questionAfterUpdateHooks, questionHook)
	case boil.AfterDeleteHook:
		questionAfterDeleteHooks = append(questionAfterDeleteHooks, questionHook)
	case boil.AfterUpsertHook:
		questionAfterUpsertHooks = append(questionAfterUpsertHooks, questionHook)
	}
}

// One returns a single question record from the query.
func (q questionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Question, error) {
	o := &Question{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for questions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Question records from the query.
func (q questionQuery) All(ctx context.Context, exec boil.ContextExecutor) (QuestionSlice, error) {
	var o []*Question

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Question slice")
	}

	if len(questionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Question records in the query.
func (q questionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count questions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q questionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if questions exists")
	}

	return count > 0, nil
}

// CreatedByUser pointed to by the foreign key.
func (o *Question) CreatedByUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CreatedBy),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Answers retrieves all the answer's Answers with an executor.
func (o *Question) Answers(mods ...qm.QueryMod) answerQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"answers\".\"question_id\"=?", o.ID),
	)

	query := Answers(queryMods...)
	queries.SetFrom(query.Query, "\"answers\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"answers\".*"})
	}

	return query
}

// LoadCreatedByUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (questionL) LoadCreatedByUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuestion interface{}, mods queries.Applicator) error {
	var slice []*Question
	var object *Question

	if singular {
		object = maybeQuestion.(*Question)
	} else {
		slice = *maybeQuestion.(*[]*Question)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &questionR{}
		}
		args = append(args, object.CreatedBy)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionR{}
			}

			for _, a := range args {
				if a == obj.CreatedBy {
					continue Outer
				}
			}

			args = append(args, obj.CreatedBy)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(questionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CreatedByUser = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.CreatedByQuestions = append(foreign.R.CreatedByQuestions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CreatedBy == foreign.ID {
				local.R.CreatedByUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.CreatedByQuestions = append(foreign.R.CreatedByQuestions, local)
				break
			}
		}
	}

	return nil
}

// LoadAnswers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (questionL) LoadAnswers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuestion interface{}, mods queries.Applicator) error {
	var slice []*Question
	var object *Question

	if singular {
		object = maybeQuestion.(*Question)
	} else {
		slice = *maybeQuestion.(*[]*Question)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &questionR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`answers`),
		qm.WhereIn(`answers.question_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load answers")
	}

	var resultSlice []*Answer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice answers")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on answers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for answers")
	}

	if len(answerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Answers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &answerR{}
			}
			foreign.R.Question = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.QuestionID {
				local.R.Answers = append(local.R.Answers, foreign)
				if foreign.R == nil {
					foreign.R = &answerR{}
				}
				foreign.R.Question = local
				break
			}
		}
	}

	return nil
}

// SetCreatedByUser of the question to the related item.
// Sets o.R.CreatedByUser to related.
// Adds o to related.R.CreatedByQuestions.
func (o *Question) SetCreatedByUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"questions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"created_by"}),
		strmangle.WhereClause("\"", "\"", 2, questionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CreatedBy = related.ID
	if o.R == nil {
		o.R = &questionR{
			CreatedByUser: related,
		}
	} else {
		o.R.CreatedByUser = related
	}

	if related.R == nil {
		related.R = &userR{
			CreatedByQuestions: QuestionSlice{o},
		}
	} else {
		related.R.CreatedByQuestions = append(related.R.CreatedByQuestions, o)
	}

	return nil
}

// AddAnswers adds the given related objects to the existing relationships
// of the question, optionally inserting them as new records.
// Appends related to o.R.Answers.
// Sets related.R.Question appropriately.
func (o *Question) AddAnswers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Answer) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.QuestionID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"answers\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"question_id"}),
				strmangle.WhereClause("\"", "\"", 2, answerPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.QuestionID = o.ID
		}
	}

	if o.R == nil {
		o.R = &questionR{
			Answers: related,
		}
	} else {
		o.R.Answers = append(o.R.Answers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &answerR{
				Question: o,
			}
		} else {
			rel.R.Question = o
		}
	}
	return nil
}

// Questions retrieves all the records using an executor.
func Questions(mods ...qm.QueryMod) questionQuery {
	mods = append(mods, qm.From("\"questions\""))
	return questionQuery{NewQuery(mods...)}
}

// FindQuestion retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuestion(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Question, error) {
	questionObj := &Question{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"questions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, questionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from questions")
	}

	return questionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Question) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no questions provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	questionInsertCacheMut.RLock()
	cache, cached := questionInsertCache[key]
	questionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			questionAllColumns,
			questionColumnsWithDefault,
			questionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(questionType, questionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(questionType, questionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"questions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"questions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into questions")
	}

	if !cached {
		questionInsertCacheMut.Lock()
		questionInsertCache[key] = cache
		questionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Question.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Question) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	questionUpdateCacheMut.RLock()
	cache, cached := questionUpdateCache[key]
	questionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			questionAllColumns,
			questionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update questions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"questions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, questionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(questionType, questionMapping, append(wl, questionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update questions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for questions")
	}

	if !cached {
		questionUpdateCacheMut.Lock()
		questionUpdateCache[key] = cache
		questionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q questionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for questions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuestionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"questions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, questionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in question slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all question")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Question) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no questions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	questionUpsertCacheMut.RLock()
	cache, cached := questionUpsertCache[key]
	questionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			questionAllColumns,
			questionColumnsWithDefault,
			questionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			questionAllColumns,
			questionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert questions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(questionPrimaryKeyColumns))
			copy(conflict, questionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"questions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(questionType, questionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(questionType, questionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert questions")
	}

	if !cached {
		questionUpsertCacheMut.Lock()
		questionUpsertCache[key] = cache
		questionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Question record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Question) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Question provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), questionPrimaryKeyMapping)
	sql := "DELETE FROM \"questions\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for questions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q questionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no questionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for questions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuestionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(questionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"questions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from question slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for questions")
	}

	if len(questionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Question) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQuestion(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuestionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuestionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"questions\".* FROM \"questions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in QuestionSlice")
	}

	*o = slice

	return nil
}

// QuestionExists checks if the Question row exists.
func QuestionExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"questions\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if questions exists")
	}

	return exists, nil
}