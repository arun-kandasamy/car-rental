package validate

import (
	"context"
	"fmt"

	policy "github.com/example/car-rental-service/proto/generated/policy"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func ValidateInboundRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	msg, ok := req.(proto.Message)
	if !ok {
		return handler(ctx, req)
	}

	if err := validateMessage(msg); err != nil {
		log.WithFields(log.Fields{
			"wrapped-error": err.Error(),
		}).Info("request validation failed")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("request validation failed:%v", err))
	}

	// invoke chained interceptor or underlying gRPC handler
	return handler(ctx, req)
}

// Validate validates the given proto message
func validateMessage(m proto.Message) error {
	return validate(m.ProtoReflect())
}

func validate(m protoreflect.Message) error {
	md := m.Descriptor()
	fds := md.Fields()

	// iterate all the field descriptors
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)

		// List, Map, Message kind of fields need to be handled
		if fd.Kind() != protoreflect.StringKind {
			return nil
		}

		opts := fd.Options().(*descriptorpb.FieldOptions)
		vo, ok := proto.GetExtension(opts, policy.E_Validate).(*policy.IdValidationOpts)
		if !ok || vo == nil {
			continue
		}

		field := newField(string(fd.Name()), m.Get(fd).Interface())
		if err := field.runValidator(vo); err != nil {
			return err
		}
	}

	return nil
}
