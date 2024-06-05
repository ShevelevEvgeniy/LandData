package custom_errors

import "github.com/pkg/errors"

var ErrKptAlreadyExist = errors.New("kpt already exist")
