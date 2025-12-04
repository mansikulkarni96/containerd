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

package images

import (
	"context"

<<<<<<< HEAD:services/images/local.go
=======
	"github.com/containerd/errdefs/pkg/errgrpc"
	"github.com/containerd/log"
>>>>>>> v2.0.7:plugins/services/images/local.go
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

<<<<<<< HEAD:services/images/local.go
	"github.com/containerd/log"

	eventstypes "github.com/containerd/containerd/api/events"
	imagesapi "github.com/containerd/containerd/api/services/images/v1"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/events"
	"github.com/containerd/containerd/gc"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/metadata"
	"github.com/containerd/containerd/pkg/deprecation"
	"github.com/containerd/containerd/pkg/epoch"
	"github.com/containerd/containerd/plugin"
	ptypes "github.com/containerd/containerd/protobuf/types"
	"github.com/containerd/containerd/services"
	"github.com/containerd/containerd/services/warning"
=======
	imagesapi "github.com/containerd/containerd/api/services/images/v1"
	"github.com/containerd/plugin"
	"github.com/containerd/plugin/registry"

	"github.com/containerd/containerd/v2/core/images"
	"github.com/containerd/containerd/v2/core/metadata"
	"github.com/containerd/containerd/v2/pkg/epoch"
	"github.com/containerd/containerd/v2/pkg/gc"
	"github.com/containerd/containerd/v2/pkg/oci"
	ptypes "github.com/containerd/containerd/v2/pkg/protobuf/types"
	"github.com/containerd/containerd/v2/plugins"
	"github.com/containerd/containerd/v2/plugins/services"
	"github.com/containerd/containerd/v2/plugins/services/warning"
>>>>>>> v2.0.7:plugins/services/images/local.go
)

var empty = &ptypes.Empty{}

func init() {
	registry.Register(&plugin.Registration{
		Type: plugins.ServicePlugin,
		ID:   services.ImagesService,
		Requires: []plugin.Type{
<<<<<<< HEAD:services/images/local.go
			plugin.MetadataPlugin,
			plugin.GCPlugin,
			plugin.WarningPlugin,
=======
			plugins.MetadataPlugin,
			plugins.GCPlugin,
			plugins.WarningPlugin,
>>>>>>> v2.0.7:plugins/services/images/local.go
		},
		InitFn: func(ic *plugin.InitContext) (interface{}, error) {
			m, err := ic.GetSingle(plugins.MetadataPlugin)
			if err != nil {
				return nil, err
			}
			g, err := ic.GetSingle(plugins.GCPlugin)
			if err != nil {
				return nil, err
			}
			w, err := ic.GetSingle(plugins.WarningPlugin)
			if err != nil {
				return nil, err
			}
			w, err := ic.Get(plugin.WarningPlugin)
			if err != nil {
				return nil, err
			}

			return &local{
<<<<<<< HEAD:services/images/local.go
				store:     metadata.NewImageStore(m.(*metadata.DB)),
				publisher: ic.Events,
				gc:        g.(gcScheduler),
				warnings:  w.(warning.Service),
=======
				store:    metadata.NewImageStore(m.(*metadata.DB)),
				gc:       g.(gcScheduler),
				warnings: w.(warning.Service),
>>>>>>> v2.0.7:plugins/services/images/local.go
			}, nil
		},
	})
}

type gcScheduler interface {
	ScheduleAndWait(context.Context) (gc.Stats, error)
}

type local struct {
<<<<<<< HEAD:services/images/local.go
	store     images.Store
	gc        gcScheduler
	publisher events.Publisher
	warnings  warning.Service
=======
	store    images.Store
	gc       gcScheduler
	warnings warning.Service
>>>>>>> v2.0.7:plugins/services/images/local.go
}

var _ imagesapi.ImagesClient = &local{}

func (l *local) Get(ctx context.Context, req *imagesapi.GetImageRequest, _ ...grpc.CallOption) (*imagesapi.GetImageResponse, error) {
	image, err := l.store.Get(ctx, req.Name)
	if err != nil {
		return nil, errgrpc.ToGRPC(err)
	}

	imagepb := imageToProto(&image)
	return &imagesapi.GetImageResponse{
		Image: imagepb,
	}, nil
}

func (l *local) List(ctx context.Context, req *imagesapi.ListImagesRequest, _ ...grpc.CallOption) (*imagesapi.ListImagesResponse, error) {
	images, err := l.store.List(ctx, req.Filters...)
	if err != nil {
		return nil, errgrpc.ToGRPC(err)
	}

	return &imagesapi.ListImagesResponse{
		Images: imagesToProto(images),
	}, nil
}

func (l *local) Create(ctx context.Context, req *imagesapi.CreateImageRequest, _ ...grpc.CallOption) (*imagesapi.CreateImageResponse, error) {
	log.G(ctx).WithField("name", req.Image.Name).WithField("target", req.Image.Target.Digest).Debugf("create image")
	if req.Image.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Image.Name required")
	}

	var (
		image = imageFromProto(req.Image)
		resp  imagesapi.CreateImageResponse
	)
	if req.SourceDateEpoch != nil {
		tm := req.SourceDateEpoch.AsTime()
		ctx = epoch.WithSourceDateEpoch(ctx, &tm)
	}
	created, err := l.store.Create(ctx, image)
	if err != nil {
		return nil, errgrpc.ToGRPC(err)
	}

	resp.Image = imageToProto(&created)

<<<<<<< HEAD
<<<<<<< HEAD:services/images/local.go
	if err := l.publisher.Publish(ctx, "/images/create", &eventstypes.ImageCreate{
		Name:   resp.Image.Name,
		Labels: resp.Image.Labels,
	}); err != nil {
		return nil, err
	}

=======
>>>>>>> v2.0.7:plugins/services/images/local.go
	l.emitSchema1DeprecationWarning(ctx, &image)
=======
>>>>>>> v2.1.0
	return &resp, nil
}

func (l *local) Update(ctx context.Context, req *imagesapi.UpdateImageRequest, _ ...grpc.CallOption) (*imagesapi.UpdateImageResponse, error) {
	if req.Image.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Image.Name required")
	}

	var (
		image      = imageFromProto(req.Image)
		resp       imagesapi.UpdateImageResponse
		fieldpaths []string
	)

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		fieldpaths = append(fieldpaths, req.UpdateMask.Paths...)
	}

	if req.SourceDateEpoch != nil {
		tm := req.SourceDateEpoch.AsTime()
		ctx = epoch.WithSourceDateEpoch(ctx, &tm)
	}

	updated, err := l.store.Update(ctx, image, fieldpaths...)
	if err != nil {
		return nil, errgrpc.ToGRPC(err)
	}

	resp.Image = imageToProto(&updated)

<<<<<<< HEAD
<<<<<<< HEAD:services/images/local.go
	if err := l.publisher.Publish(ctx, "/images/update", &eventstypes.ImageUpdate{
		Name:   resp.Image.Name,
		Labels: resp.Image.Labels,
	}); err != nil {
		return nil, err
	}

=======
>>>>>>> v2.0.7:plugins/services/images/local.go
	l.emitSchema1DeprecationWarning(ctx, &image)
=======
>>>>>>> v2.1.0
	return &resp, nil
}

func (l *local) Delete(ctx context.Context, req *imagesapi.DeleteImageRequest, _ ...grpc.CallOption) (*ptypes.Empty, error) {
	log.G(ctx).WithField("name", req.Name).Debugf("delete image")

	var opts []images.DeleteOpt
	if req.Target != nil {
		desc := oci.DescriptorFromProto(req.Target)
		opts = append(opts, images.DeleteTarget(&desc))
	}

	// Sync option handled here after event is published
	if err := l.store.Delete(ctx, req.Name, opts...); err != nil {
		return nil, errgrpc.ToGRPC(err)
	}

	if req.Sync {
		if _, err := l.gc.ScheduleAndWait(ctx); err != nil {
			return nil, err
		}
	}

	return empty, nil
}
<<<<<<< HEAD

func (l *local) emitSchema1DeprecationWarning(ctx context.Context, image *images.Image) {
	if image == nil {
		return
	}
	dgst, ok := image.Labels[images.ConvertedDockerSchema1LabelKey]
	if !ok {
		return
	}
	log.G(ctx).WithField("name", image.Name).WithField("schema1digest", dgst).Warn("conversion from schema 1 images is deprecated")
	l.warnings.Emit(ctx, deprecation.PullSchema1Image)
}

func (l *local) emitSchema1DeprecationWarning(ctx context.Context, image *images.Image) {
	if image == nil {
		return
	}
	dgst, ok := image.Labels[images.ConvertedDockerSchema1LabelKey]
	if !ok {
		return
	}
	log.G(ctx).WithField("name", image.Name).WithField("schema1digest", dgst).Warn("conversion from schema 1 images is deprecated")
	l.warnings.Emit(ctx, deprecation.PullSchema1Image)
}
=======
>>>>>>> v2.1.0
