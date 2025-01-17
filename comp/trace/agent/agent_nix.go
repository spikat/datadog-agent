// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !windows

package agent

import (
	"context"

	"go.uber.org/fx"
)

func setupShutdown(_ context.Context, shutdowner fx.Shutdowner) {
	// Handle stops properly
	go handleSignal(shutdowner)
}
