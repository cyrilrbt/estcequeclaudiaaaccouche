package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func sendRequest(action string, statusManager *StatusManager, fn func(http.ResponseWriter, *http.Request, *StatusManager)) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(action, "", nil)
	w := httptest.NewRecorder()

	handler := NewHandler(fn, statusManager)

	handler.ServeHTTP(w, req)

	return w
}

func TestRenderStatus(t *testing.T) {
	manager := buildStatusManager()

	w := sendRequest("GET", manager, RenderStatus)

	if w.Code != http.StatusOK {
		t.Errorf("RenderStatus didn't return %v", http.StatusOK)
	}
}

func TestAPIStatus(t *testing.T) {
	manager := buildStatusManager()

	w := sendRequest("GET", manager, APIStatus)

	if w.Code != http.StatusOK {
		t.Errorf("APIStatus didn't return %v", http.StatusOK)
	}
}

func TestToggleStatus(t *testing.T) {
	manager := buildStatusManager()

	w := sendRequest("POST", manager, ToggleStatus)

	if w.Code != http.StatusOK {
		t.Errorf("ToggleStatus didn't return %v", http.StatusOK)
	}

	if manager.Value() != Yes {
		t.Errorf("TogglingStatus did not turned status to yes, got %s", manager.Value())
	}

	w = sendRequest("POST", manager, ToggleStatus)

	if w.Code != http.StatusOK {
		t.Errorf("ToggleStatus didn't return %v", http.StatusOK)
	}

	if manager.Value() != No {
		t.Errorf("TogglingStatus did not turned status to no, got %s", manager.Value())
	}
}

func TestToggleStatusWithTwilio(t *testing.T) {
	manager := buildStatusManager()

	w := sendRequest("POST", manager, ToggleStatusWithTwilio)

	if w.Code != http.StatusOK {
		t.Errorf("ToggleStatus didn't return %v", http.StatusOK)
	}

	if manager.Value() != Yes {
		t.Errorf("ToggleStatusWithTwilio did not turned status to yes, got %s", manager.Value())
	}

	w = sendRequest("POST", manager, ToggleStatusWithTwilio)

	if w.Code != http.StatusOK {
		t.Errorf("ToggleStatus didn't return %v", http.StatusOK)
	}

	if manager.Value() != No {
		t.Errorf("ToggleStatusWithTwilio didn not turned status to no, got %s", manager.Value())
	}
}

func TestRenderStatusWithTwilio(t *testing.T) {
	manager := buildStatusManager()

	w := sendRequest("GET", manager, RenderStatusWithTwilio)

	if w.Code != http.StatusOK {
		t.Errorf("RenderStatusWithTwilio didn't return %v", http.StatusOK)
	}
}
