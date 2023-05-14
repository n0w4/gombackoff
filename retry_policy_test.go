package gombackoff_test

import (
	"testing"
	"time"

	"github.com/n0w4/gombackoff"
)

func TestExponentialBackoff(t *testing.T) {
	rp := gombackoff.NewRetryPolicy().ExponentialBackoff(2).Times(2)
	currentTime := time.Now()
	rp.Wait()
	rp.Wait()
	elapsedTime := time.Since(currentTime)
	if elapsedTime < 6*time.Second {
		t.Errorf("Expected to wait at least 6 seconds, waited %v", elapsedTime)
	}
}

func TestLinearBackoff(t *testing.T) {
	rp := gombackoff.NewRetryPolicy().LinearBackoff(2).Times(2)
	currentTime := time.Now()
	rp.Wait()
	rp.Wait()
	elapsedTime := time.Since(currentTime)
	if elapsedTime < 4*time.Second {
		t.Errorf("Expected to wait at least 4 seconds, waited %v", elapsedTime)
	}
}
