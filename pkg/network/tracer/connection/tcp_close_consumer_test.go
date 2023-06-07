// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux_bpf

package connection

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DataDog/datadog-agent/pkg/ebpf"
)

func TestTcpCloseConsumerStopRace(t *testing.T) {
	pf := ebpf.NewPerfHandler(10)
	require.NotNil(t, pf)

	c, err := newTCPCloseConsumer(pf, nil)
	require.NoError(t, err)
	require.NotNil(t, c)

	c.Stop()
	c.FlushPending()
}