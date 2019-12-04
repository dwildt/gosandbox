package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestActive(t *testing.T) {
	activateFeature("test.google")
	assert.Equal(t, isActive("test.google"), true, "deveria estar ativa")
	
	shutdownFeature("test.google")
	assert.Equal(t, isActive("test.google"), false, "deveria estar desligada")
}

func TestWhenFeatureNotAdded(t *testing.T) {
	assert.Equal(t, isActive("test.facebook"), false, "feature não existente, deveria estar desligada")
}
