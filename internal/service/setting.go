package service

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
	innErr "github.com/airunny/timer/errors"
	"github.com/airunny/timer/internal/models"
	"github.com/go-kratos/kratos/v2/log"
)

func (s *Service) Bucket(ctx context.Context, in *v1.BucketRequest) (*v1.BucketReply, error) {
	l := log.Context(ctx)
	if in.NumberOfBuckets <= 0 {
		return nil, innErr.ErrBadRequest
	}

	release, err := s.locker.TryLock(ctx, "create_bucket", time.Second*5)
	if err != nil {
		return nil, innErr.WithMessage(innErr.ErrBadRequest, "frequency limited")
	}
	defer release()

	value, err := s.global.GetByName(ctx, models.NamespaceMetadata, "bucket")
	if err != nil {
		l.Error("GetByName Err:%v", err)
		return nil, err
	}

	var metadata models.MetadataBucket
	err = metadata.FromString(value.Value)
	if err != nil {
		l.Errorf("FromString Err:%v", err)
		return nil, err
	}

	if in.NumberOfBuckets <= metadata.NumberOfBuckets {
		return nil, innErr.WithMessage(innErr.ErrBadRequest, "分桶数量不能少于现有数量")
	}

	metadata.NumberOfBuckets = in.NumberOfBuckets
	for index := metadata.NumberOfBuckets; index <= in.NumberOfBuckets; index++ {
		metadata.Buckets = append(metadata.Buckets, fmt.Sprintf("bucket%d", index))
	}

	value.Value = metadata.String()
	err = s.global.UpdateByKey(ctx, value)
	if err != nil {
		l.Errorf("UpdateByKey Err:%v", err)
		return nil, err
	}
	s.publisher.Publish(ctx, bucketChangeTopic, value.Value)
	return &v1.BucketReply{
		Buckets: metadata.Buckets,
	}, nil
}

func (s *Service) ListBucket(ctx context.Context, in *v1.ListBucketRequest) (*v1.ListBucketReply, error) {
	l := log.Context(ctx)
	value, err := s.global.GetByName(ctx, models.NamespaceMetadata, "bucket")
	if err != nil {
		l.Error("GetByName Err:%v", err)
		return nil, err
	}

	var metadata models.MetadataBucket
	err = metadata.FromString(value.Value)
	if err != nil {
		l.Errorf("FromString Err:%v", err)
		return nil, err
	}

	buckets := make([]*v1.Bucket, 0, len(metadata.Buckets))
	for _, bucket := range metadata.Buckets {
		buckets = append(buckets, &v1.Bucket{
			Bucket:    bucket,
			ServiceIp: "",
		})
	}
	return &v1.ListBucketReply{
		NumberOfBuckets: metadata.NumberOfBuckets,
		Buckets:         buckets,
	}, nil
}
