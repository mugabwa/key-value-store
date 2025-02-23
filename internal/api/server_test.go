package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetGetDelete(t *testing.T) {
	server := New()

	req := httptest.NewRequest(http.MethodPut, "/kv/testKey", bytes.NewBufferString("testValue"))
	w := httptest.NewRecorder()
	server.set(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204, got %d", resp.StatusCode)
	}

	req = httptest.NewRequest(http.MethodGet, "/kv/testKey", nil)
	w = httptest.NewRecorder()
	server.get(w, req)

	resp = w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}

	expectedBody := "testValue\n"
	gotBody := w.Body.String()
	if gotBody != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, gotBody)
	}

	req = httptest.NewRequest(http.MethodDelete, "/kv/testKey", nil)
	w = httptest.NewRecorder()
	server.delete(w, req)

	resp = w.Result()
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204, got %d", resp.StatusCode)
	}

	req = httptest.NewRequest(http.MethodGet, "/kv/testKey", nil)
	w = httptest.NewRecorder()
	server.get(w, req)

	resp = w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", resp.StatusCode)
	}
}

func TestSetEmptyValue(t *testing.T) {
	server := New()

	req := httptest.NewRequest(http.MethodPut, "/kv/testKey", bytes.NewBufferString(""))
	w := httptest.NewRecorder()
	server.set(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", resp.StatusCode)
	}
}

func TestDeleteMissingKey(t *testing.T) {
	server := New()

	req := httptest.NewRequest(http.MethodDelete, "/kv/testKey", nil)
	w := httptest.NewRecorder()
	server.delete(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", resp.StatusCode)
	}
}