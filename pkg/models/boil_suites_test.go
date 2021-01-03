// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Answers", testAnswers)
	t.Run("Questions", testQuestions)
	t.Run("RevokeTokens", testRevokeTokens)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Answers", testAnswersDelete)
	t.Run("Questions", testQuestionsDelete)
	t.Run("RevokeTokens", testRevokeTokensDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Answers", testAnswersQueryDeleteAll)
	t.Run("Questions", testQuestionsQueryDeleteAll)
	t.Run("RevokeTokens", testRevokeTokensQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Answers", testAnswersSliceDeleteAll)
	t.Run("Questions", testQuestionsSliceDeleteAll)
	t.Run("RevokeTokens", testRevokeTokensSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Answers", testAnswersExists)
	t.Run("Questions", testQuestionsExists)
	t.Run("RevokeTokens", testRevokeTokensExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Answers", testAnswersFind)
	t.Run("Questions", testQuestionsFind)
	t.Run("RevokeTokens", testRevokeTokensFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Answers", testAnswersBind)
	t.Run("Questions", testQuestionsBind)
	t.Run("RevokeTokens", testRevokeTokensBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Answers", testAnswersOne)
	t.Run("Questions", testQuestionsOne)
	t.Run("RevokeTokens", testRevokeTokensOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Answers", testAnswersAll)
	t.Run("Questions", testQuestionsAll)
	t.Run("RevokeTokens", testRevokeTokensAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Answers", testAnswersCount)
	t.Run("Questions", testQuestionsCount)
	t.Run("RevokeTokens", testRevokeTokensCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Answers", testAnswersHooks)
	t.Run("Questions", testQuestionsHooks)
	t.Run("RevokeTokens", testRevokeTokensHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Answers", testAnswersInsert)
	t.Run("Answers", testAnswersInsertWhitelist)
	t.Run("Questions", testQuestionsInsert)
	t.Run("Questions", testQuestionsInsertWhitelist)
	t.Run("RevokeTokens", testRevokeTokensInsert)
	t.Run("RevokeTokens", testRevokeTokensInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AnswerToUserUsingCreatedByUser", testAnswerToOneUserUsingCreatedByUser)
	t.Run("AnswerToQuestionUsingQuestion", testAnswerToOneQuestionUsingQuestion)
	t.Run("QuestionToUserUsingCreatedByUser", testQuestionToOneUserUsingCreatedByUser)
	t.Run("RevokeTokenToUserUsingOwner", testRevokeTokenToOneUserUsingOwner)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("QuestionToAnswers", testQuestionToManyAnswers)
	t.Run("UserToCreatedByAnswers", testUserToManyCreatedByAnswers)
	t.Run("UserToCreatedByQuestions", testUserToManyCreatedByQuestions)
	t.Run("UserToOwnerRevokeTokens", testUserToManyOwnerRevokeTokens)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AnswerToUserUsingCreatedByAnswers", testAnswerToOneSetOpUserUsingCreatedByUser)
	t.Run("AnswerToQuestionUsingAnswers", testAnswerToOneSetOpQuestionUsingQuestion)
	t.Run("QuestionToUserUsingCreatedByQuestions", testQuestionToOneSetOpUserUsingCreatedByUser)
	t.Run("RevokeTokenToUserUsingOwnerRevokeTokens", testRevokeTokenToOneSetOpUserUsingOwner)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("QuestionToAnswers", testQuestionToManyAddOpAnswers)
	t.Run("UserToCreatedByAnswers", testUserToManyAddOpCreatedByAnswers)
	t.Run("UserToCreatedByQuestions", testUserToManyAddOpCreatedByQuestions)
	t.Run("UserToOwnerRevokeTokens", testUserToManyAddOpOwnerRevokeTokens)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Answers", testAnswersReload)
	t.Run("Questions", testQuestionsReload)
	t.Run("RevokeTokens", testRevokeTokensReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Answers", testAnswersReloadAll)
	t.Run("Questions", testQuestionsReloadAll)
	t.Run("RevokeTokens", testRevokeTokensReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Answers", testAnswersSelect)
	t.Run("Questions", testQuestionsSelect)
	t.Run("RevokeTokens", testRevokeTokensSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Answers", testAnswersUpdate)
	t.Run("Questions", testQuestionsUpdate)
	t.Run("RevokeTokens", testRevokeTokensUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Answers", testAnswersSliceUpdateAll)
	t.Run("Questions", testQuestionsSliceUpdateAll)
	t.Run("RevokeTokens", testRevokeTokensSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
