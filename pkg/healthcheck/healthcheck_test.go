package healthcheck

import (
				"testing"
				"net/http"
			)

func TestFailed(t *testing.T) {
	t.Run("Test state failed", func(t *testing.T) {
		s := State{}
		s.Failed()
		if s.Status != "failed" {
			t.Errorf("Invalid status value %s, should be %s", s.Status, "failed")
		}
	})
	t.Run("Test state ok", func(t *testing.T) {
		s := State{}
		s.Ok()
		if s.Status != "ok" {
			t.Errorf("Invalid status value %s, should be %s", s.Status, "ok")
		}
	})
}

func TestRoute(t *testing.T) {
	t.Run("Test status endpoint", func(t *testing.T) {
		req, err := http.NewRequest("GET", "localhost:8080/status", nil)
		if err != nil {
			t.Fatalf("Couldn't created request: %v", err)
		}
		rec := httptest.NewRecorder()

		Route(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK; got %v", res.Status)
		}
	})
}

func TestInfo(t *testing.T) {
	t.Run("Test info endpoint", func(t *testing.T) {
		req, err := http.NewRequest("GET", "localhost:8080/info", nil)
		if err != nil {
			t.Fatalf("Couldn't created request: %v", err)
		}
		rec := httptest.NewRecorder()

		Info(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK; got %v", res.Status)
		}
	})
	
}
