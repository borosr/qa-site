package healthcheck

import "testing"

func TestState_Failed(t *testing.T) {
	t.Run("failed", func(t *testing.T) {
		s := State{}
		s.Failed()
		if s.Status != "failed" {
			t.Errorf("Invalid status value %s, should be %s", s.Status, "failed")
		}
	})
	t.Run("ok", func(t *testing.T) {
		s := State{}
		s.Ok()
		if s.Status != "ok" {
			t.Errorf("Invalid status value %s, should be %s", s.Status, "ok")
		}
	})
}
