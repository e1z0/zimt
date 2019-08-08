package api

import "testing"

func TestBrokerFormatUptimeWhenItIsParseable(t *testing.T) {
	b := Broker{Uptime: "42 seconds"}
	expected := "0d 0h 0m 42s (42 seconds)"
	if b.FormatUptime() != expected {
		t.Errorf("Expected %v, got %v", expected, b.FormatUptime())
	}

	b = Broker{Uptime: "1932719 seconds"}
	expected = "22d 8h 51m 59s (1932719 seconds)"
	if b.FormatUptime() != expected {
		t.Errorf("Expected %v, got %v", expected, b.FormatUptime())
	}
}

func TestBrokerFormatUptimeWhenItIsNotParseable(t *testing.T) {
	b := Broker{Uptime: "foo bar"}
	expected := "foo bar"
	if b.FormatUptime() != expected {
		t.Errorf("Expected %v, got %v", expected, b.FormatUptime())
	}
}
