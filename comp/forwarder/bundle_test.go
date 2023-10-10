// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package forwarder

import (
	"testing"

	"github.com/DataDog/datadog-agent/comp/core"
	"github.com/DataDog/datadog-agent/comp/forwarder/defaultforwarder"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
	"go.uber.org/fx"
)

func TestBundleDependencies(t *testing.T) {
	fxutil.TestBundle(t, Bundle,
		core.MockBundle,
		fx.Supply(defaultforwarder.Params{}),
	)
}
