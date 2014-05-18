package lastfm

import (
	"reflect"
	"testing"
)

func TestClientGetter(t *testing.T) {
	lfm := New("apikey", "apisecret")
	lfmTesting := New("api_key_for_testing", "api_secret_for_testing")

	if reflect.TypeOf(lfm.getter).String() != "*http.Client" {
		t.Errorf("Expected lfm.getter to be *http.Client, but it is %s",
			reflect.TypeOf(lfm.getter).String())
	}

	if reflect.TypeOf(lfmTesting.getter).String() != "*lastfm.dummyClient" {
		t.Errorf("Expected lfmTesting.getter to be *lastfm.dummyClient, but it is %s",
			reflect.TypeOf(lfmTesting.getter).String())
	}
}

func TestClientTokenSetterAndGetter(t *testing.T) {
	token := "test_token"
	lfm := New("api_key_for_testing", "api_secret_for_testing")

	if lfm.HasToken() {
		t.Errorf("Expected token to be empty at the beginning, but it is %v.", lfm.GetToken())
	}

	lfm.SetToken(token)

	if !lfm.HasToken() {
		t.Error("Expected token to be set, but it is empty.")
	}

	if lfm.GetToken() != lfm.GetToken() {
		t.Error("Expected GetToken to be deterministic, but it isn't.")
	}

	if lfm.GetToken() != token {
		t.Errorf("Expected %v but got %v", token, lfm.GetToken())
	}
}

func TestClientUserAgentSetterAndGetter(t *testing.T) {
	userAgent := "test_user_agent"
	lfm := New("api_key_for_testing", "api_secret_for_testing")

	if lfm.HasUserAgent() {
		t.Errorf("Expected user agent to be empty at the beginning, but it is %v.", lfm.GetUserAgent())
	}

	lfm.SetUserAgent(userAgent)

	if !lfm.HasUserAgent() {
		t.Error("Expected user agent to be set, but it is empty.")
	}

	if lfm.GetUserAgent() != lfm.GetUserAgent() {
		t.Error("Expected GetUserAgent to be deterministic, but it isn't.")
	}

	if lfm.GetUserAgent() != userAgent {
		t.Errorf("Expected %v but got %v", userAgent, lfm.GetUserAgent())
	}
}

func TestClientSessionKeySetterAndGetter(t *testing.T) {
	sessionKey := "test_session_key"
	lfm := New("api_key_for_testing", "api_secret_for_testing")

	if lfm.HasSessionKey() {
		t.Errorf("Expected SessionKey to be empty at the beginning, but it is %v.", lfm.GetToken())
	}

	lfm.SetSessionKey(sessionKey)

	if !lfm.HasSessionKey() {
		t.Error("Expected sessionKey to be set, but it is empty.")
	}

	if lfm.GetSessionKey() != lfm.GetSessionKey() {
		t.Error("Expected GetSessionKey to be deterministic, but it isn't.")
	}

	if lfm.GetSessionKey() != sessionKey {
		t.Errorf("Expected %v but got %v", sessionKey, lfm.GetSessionKey())
	}
}
