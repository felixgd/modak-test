package ratelimiterrules

import "time"

// RateLimitRule represents the rate limit rule for a notification type.
type RateLimitRule struct {
	MaxCount int
	Duration time.Duration
}
