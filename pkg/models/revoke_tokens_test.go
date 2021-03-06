// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRevokeTokens(t *testing.T) {
	t.Parallel()

	query := RevokeTokens()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRevokeTokensDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevokeTokensQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RevokeTokens().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevokeTokensSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevokeTokenSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevokeTokensExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RevokeTokenExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RevokeToken exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RevokeTokenExists to return true, but got false.")
	}
}

func testRevokeTokensFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	revokeTokenFound, err := FindRevokeToken(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if revokeTokenFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRevokeTokensBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RevokeTokens().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRevokeTokensOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RevokeTokens().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRevokeTokensAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	revokeTokenOne := &RevokeToken{}
	revokeTokenTwo := &RevokeToken{}
	if err = randomize.Struct(seed, revokeTokenOne, revokeTokenDBTypes, false, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}
	if err = randomize.Struct(seed, revokeTokenTwo, revokeTokenDBTypes, false, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revokeTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revokeTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevokeTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRevokeTokensCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	revokeTokenOne := &RevokeToken{}
	revokeTokenTwo := &RevokeToken{}
	if err = randomize.Struct(seed, revokeTokenOne, revokeTokenDBTypes, false, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}
	if err = randomize.Struct(seed, revokeTokenTwo, revokeTokenDBTypes, false, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revokeTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revokeTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func revokeTokenBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func revokeTokenAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func revokeTokenAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func revokeTokenBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func revokeTokenAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func revokeTokenBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func revokeTokenAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func revokeTokenBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func revokeTokenAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevokeToken) error {
	*o = RevokeToken{}
	return nil
}

func testRevokeTokensHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RevokeToken{}
	o := &RevokeToken{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevokeToken object: %s", err)
	}

	AddRevokeTokenHook(boil.BeforeInsertHook, revokeTokenBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	revokeTokenBeforeInsertHooks = []RevokeTokenHook{}

	AddRevokeTokenHook(boil.AfterInsertHook, revokeTokenAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	revokeTokenAfterInsertHooks = []RevokeTokenHook{}

	AddRevokeTokenHook(boil.AfterSelectHook, revokeTokenAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	revokeTokenAfterSelectHooks = []RevokeTokenHook{}

	AddRevokeTokenHook(boil.BeforeUpdateHook, revokeTokenBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	revokeTokenBeforeUpdateHooks = []RevokeTokenHook{}

	AddRevokeTokenHook(boil.AfterUpdateHook, revokeTokenAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	revokeTokenAfterUpdateHooks = []RevokeTokenHook{}

	AddRevokeTokenHook(boil.BeforeDeleteHook, revokeTokenBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	revokeTokenBeforeDeleteHooks = []RevokeTokenHook{}

	AddRevokeTokenHook(boil.AfterDeleteHook, revokeTokenAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	revokeTokenAfterDeleteHooks = []RevokeTokenHook{}

	AddRevokeTokenHook(boil.BeforeUpsertHook, revokeTokenBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	revokeTokenBeforeUpsertHooks = []RevokeTokenHook{}

	AddRevokeTokenHook(boil.AfterUpsertHook, revokeTokenAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	revokeTokenAfterUpsertHooks = []RevokeTokenHook{}
}

func testRevokeTokensInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevokeTokensInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(revokeTokenColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevokeTokenToOneUserUsingOwner(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RevokeToken
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, revokeTokenDBTypes, false, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OwnerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Owner().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RevokeTokenSlice{&local}
	if err = local.L.LoadOwner(ctx, tx, false, (*[]*RevokeToken)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Owner == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Owner = nil
	if err = local.L.LoadOwner(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Owner == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRevokeTokenToOneSetOpUserUsingOwner(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevokeToken
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revokeTokenDBTypes, false, strmangle.SetComplement(revokeTokenPrimaryKeyColumns, revokeTokenColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetOwner(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Owner != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OwnerRevokeTokens[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OwnerID != x.ID {
			t.Error("foreign key was wrong value", a.OwnerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OwnerID))
		reflect.Indirect(reflect.ValueOf(&a.OwnerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OwnerID != x.ID {
			t.Error("foreign key was wrong value", a.OwnerID, x.ID)
		}
	}
}

func testRevokeTokensReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRevokeTokensReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevokeTokenSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRevokeTokensSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevokeTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	revokeTokenDBTypes = map[string]string{`ID`: `varchar`, `OwnerID`: `varchar`, `Token`: `varchar`}
	_                  = bytes.MinRead
)

func testRevokeTokensUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(revokeTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(revokeTokenAllColumns) == len(revokeTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRevokeTokensSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(revokeTokenAllColumns) == len(revokeTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevokeToken{}
	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revokeTokenDBTypes, true, revokeTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(revokeTokenAllColumns, revokeTokenPrimaryKeyColumns) {
		fields = revokeTokenAllColumns
	} else {
		fields = strmangle.SetComplement(
			revokeTokenAllColumns,
			revokeTokenPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RevokeTokenSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRevokeTokensUpsert(t *testing.T) {
	t.Parallel()

	if len(revokeTokenAllColumns) == len(revokeTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RevokeToken{}
	if err = randomize.Struct(seed, &o, revokeTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevokeToken: %s", err)
	}

	count, err := RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, revokeTokenDBTypes, false, revokeTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevokeToken struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevokeToken: %s", err)
	}

	count, err = RevokeTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
