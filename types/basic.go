package types

// Account identifier. Provides access to user's state.
type AccountID = string

// Gas is a type for storing amounts of gas.
type Gas = uint64

// Nonce for transactions.
type Nonce = uint64

// Time nanoseconds fit into uint128. Using existing Balance type which
// implements JSON marshal/unmarshal
type TimeNanos = Balance

// Some more aliases for blocks...
type BlockHeight = uint64
type ShardID = uint64
