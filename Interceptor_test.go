/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package propagate

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/grpc_testing"
	"testing"
)

func TestUnary(t *testing.T) {
	// Create a mock context.
	ctx := context.Background()

	// set the metadata.
	ctx = metadata.NewIncomingContext(ctx, metadata.MD{
		"key": []string{"value"},
	})

	// Create test data.
	method := "example.method"
	req := &grpc_testing.Empty{}
	reply := &grpc_testing.Empty{}
	cc := &grpc.ClientConn{}

	// Create a mock invoker function.
	invokerMock := func(c context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		// Verify that the context is correct.
		md, ok := metadata.FromOutgoingContext(c)
		if !ok {
			t.Errorf("metadata not found in context")
		}
		//		check md
		if len(md["key"]) != 1 {
			t.Errorf("metadata key not found in context")
		}
		if md["key"][0] != "value" {
			t.Errorf("metadata key not found in context")
		}
		return nil
	}

	// Call the function being tested.
	err := Unary(ctx, method, req, reply, cc, invokerMock)

	// Verify that the error is nil.
	if err != nil {
		t.Errorf("Unary failed with error: %v", err)
	}
}
