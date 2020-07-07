package bolt

import (
	"context"
	"testing"

	"github.com/agreyfox/gisvs"
	"github.com/stretchr/testify/assert"
)

func TestTransformStorage_Store(t *testing.T) {
	completeTransformation := &gisvs.NewTransformation{ApplicationID: testApplication.ID, FileID: testFile.ID, Actions: gisvs.TransformationOption{Width: 5}}
	transformationMustBeUnique := &gisvs.NewTransformation{ApplicationID: testApplication.ID, FileID: testFile.ID, Actions: gisvs.TransformationOption{Width: 5}}
	NoApplicationID := &gisvs.NewTransformation{FileID: testFile.ID, Actions: gisvs.TransformationOption{Width: 10}}

	tests := []struct {
		newTran     *gisvs.NewTransformation
		expectedErr error
		expected    bool
		description string
	}{
		{completeTransformation, nil, true, "insert complete transformation is successful"},
		{transformationMustBeUnique, gisvs.ErrTransformationNotUnique, false, "insert of duplipacte transformation should fail"},
		{NoApplicationID, nil, false, "application_id shouldn be required"},
	}

	ctx := context.Background()

	for _, test := range tests {
		a, err := testTransformationStorage.Store(ctx, test.newTran)
		result := err == nil

		if !assert.Equal(t, test.expected, result) {
			t.Errorf("test: %s. error: %v", test.description, err)
			return
		}

		if test.expectedErr != nil {
			assert.Equal(t, test.expectedErr, err, "error should match")
		}

		if err == nil {
			assert.NotNil(t, a.ID, test.description)
			assert.NotNil(t, a.CreatedAt, test.description)
			assert.NotNil(t, a.UpdatedAt, test.description)
			assert.Equal(t, test.newTran.ApplicationID, a.ApplicationID, "ApplicationID should be equal", test.description)
			assert.Equal(t, test.newTran.FileID, a.FileID, "FileID should be equal", test.description)
			assert.Equal(t, test.newTran.Actions, a.Actions, "Actions should be equal", test.description)
		}
	}
}
