package gombackoff

import "time"

type retryPolicy struct {
	maxRetries   int
	currentRetry int
	exponential  bool
	factor       int
}

func NewRetryPolicy() *retryPolicy {
	return &retryPolicy{
		maxRetries:   0,
		currentRetry: 1,
	}
}

func (rp *retryPolicy) Times(retries int) *retryPolicy {
	rp.maxRetries = retries
	return rp
}

func (rp *retryPolicy) ExponentialBackoff(factor int) *retryPolicy {
	rp.exponential = true
	rp.factor = factor
	return rp
}

func (rp *retryPolicy) LinearBackoff(factor int) *retryPolicy {
	rp.exponential = false
	rp.factor = factor
	return rp
}

func (rp *retryPolicy) Wait() bool {
	if rp.currentRetry > rp.maxRetries {
		return false
	}
	rp.sleep()
	rp.currentRetry++
	return true
}

func (rp *retryPolicy) sleep() {
	if rp.exponential {
		time.Sleep((time.Duration(rp.factor*rp.currentRetry) * time.Second))
		return
	}
	time.Sleep(time.Duration(rp.factor) * time.Second)
}
