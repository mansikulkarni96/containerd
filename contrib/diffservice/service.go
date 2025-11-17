/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package diffservice

import (
	"context"

	diffapi "github.com/containerd/containerd/api/services/diff/v1"
<<<<<<< HEAD
	"github.com/containerd/containerd/api/types"
	"github.com/containerd/containerd/diff"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/mount"
	"github.com/containerd/typeurl/v2"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
=======
	"github.com/containerd/errdefs"
	"github.com/containerd/errdefs/pkg/errgrpc"
	"github.com/containerd/typeurl/v2"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/containerd/containerd/v2/core/diff"
	"github.com/containerd/containerd/v2/core/mount"
	"github.com/containerd/containerd/v2/pkg/oci"
>>>>>>> v2.0.7
)

type service struct {
	applier  diff.Applier
	comparer diff.Comparer
	diffapi.UnimplementedDiffServer
}

func FromApplierAndComparer(a diff.Applier, c diff.Comparer) diffapi.DiffServer {
	return &service{
		applier:  a,
		comparer: c,
	}
}
func (s *service) Apply(ctx context.Context, er *diffapi.ApplyRequest) (*diffapi.ApplyResponse, error) {
	if s.applier == nil {
<<<<<<< HEAD
		return nil, errdefs.ToGRPC(errdefs.ErrNotImplemented)
=======
		return nil, errgrpc.ToGRPC(errdefs.ErrNotImplemented)
>>>>>>> v2.0.7
	}

	var (
		ocidesc ocispec.Descriptor
		err     error
<<<<<<< HEAD
		desc    = toDescriptor(er.Diff)
		mounts  = toMounts(er.Mounts)
=======
		desc    = oci.DescriptorFromProto(er.Diff)
		mounts  = mount.FromProto(er.Mounts)
>>>>>>> v2.0.7
	)

	var opts []diff.ApplyOpt
	if er.Payloads != nil {
		payloads := make(map[string]typeurl.Any)
		for k, v := range er.Payloads {
			payloads[k] = v
		}
		opts = append(opts, diff.WithPayloads(payloads))
	}
	opts = append(opts, diff.WithSyncFs(er.SyncFs))

	ocidesc, err = s.applier.Apply(ctx, desc, mounts, opts...)
	if err != nil {
<<<<<<< HEAD
		return nil, errdefs.ToGRPC(err)
	}

	return &diffapi.ApplyResponse{
		Applied: fromDescriptor(ocidesc),
=======
		return nil, errgrpc.ToGRPC(err)
	}

	return &diffapi.ApplyResponse{
		Applied: oci.DescriptorToProto(ocidesc),
>>>>>>> v2.0.7
	}, nil
}

func (s *service) Diff(ctx context.Context, dr *diffapi.DiffRequest) (*diffapi.DiffResponse, error) {
	if s.comparer == nil {
<<<<<<< HEAD
		return nil, errdefs.ToGRPC(errdefs.ErrNotImplemented)
=======
		return nil, errgrpc.ToGRPC(errdefs.ErrNotImplemented)
>>>>>>> v2.0.7
	}
	var (
		ocidesc ocispec.Descriptor
		err     error
<<<<<<< HEAD
		aMounts = toMounts(dr.Left)
		bMounts = toMounts(dr.Right)
=======
		aMounts = mount.FromProto(dr.Left)
		bMounts = mount.FromProto(dr.Right)
>>>>>>> v2.0.7
	)

	var opts []diff.Opt
	if dr.MediaType != "" {
		opts = append(opts, diff.WithMediaType(dr.MediaType))
	}
	if dr.Ref != "" {
		opts = append(opts, diff.WithReference(dr.Ref))
	}
	if dr.Labels != nil {
		opts = append(opts, diff.WithLabels(dr.Labels))
	}
	if dr.SourceDateEpoch != nil {
		tm := dr.SourceDateEpoch.AsTime()
		opts = append(opts, diff.WithSourceDateEpoch(&tm))
	}

	ocidesc, err = s.comparer.Compare(ctx, aMounts, bMounts, opts...)
	if err != nil {
<<<<<<< HEAD
		return nil, errdefs.ToGRPC(err)
	}

	return &diffapi.DiffResponse{
		Diff: fromDescriptor(ocidesc),
	}, nil
}

func toMounts(apim []*types.Mount) []mount.Mount {
	mounts := make([]mount.Mount, len(apim))
	for i, m := range apim {
		mounts[i] = mount.Mount{
			Type:    m.Type,
			Source:  m.Source,
			Target:  m.Target,
			Options: m.Options,
		}
	}
	return mounts
}

func toDescriptor(d *types.Descriptor) ocispec.Descriptor {
	return ocispec.Descriptor{
		MediaType:   d.MediaType,
		Digest:      digest.Digest(d.Digest),
		Size:        d.Size,
		Annotations: d.Annotations,
	}
}

func fromDescriptor(d ocispec.Descriptor) *types.Descriptor {
	return &types.Descriptor{
		MediaType:   d.MediaType,
		Digest:      d.Digest.String(),
		Size:        d.Size,
		Annotations: d.Annotations,
	}
}
=======
		return nil, errgrpc.ToGRPC(err)
	}

	return &diffapi.DiffResponse{
		Diff: oci.DescriptorToProto(ocidesc),
	}, nil
}
>>>>>>> v2.0.7
