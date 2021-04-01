//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package answer

import (
	"context"
	"testing"

	"github.com/semi-technologies/weaviate/entities/search"
	qnamodels "github.com/semi-technologies/weaviate/modules/qna-transformers/additional/models"
	"github.com/semi-technologies/weaviate/modules/qna-transformers/ent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdditionalAnswerProvider(t *testing.T) {
	t.Run("should fail with empty content", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.NotNil(t, err)
		require.NotEmpty(t, out)
		assert.Error(t, err, "empty content")
	})

	t.Run("should fail with empty question", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
				Schema: map[string]interface{}{
					"content": "content",
				},
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.NotNil(t, err)
		require.NotEmpty(t, out)
		assert.Error(t, err, "empty content")
	})

	t.Run("should answer", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
				Schema: map[string]interface{}{
					"content": "content",
				},
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{
			"answer": "fake param",
		}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.Nil(t, err)
		require.NotEmpty(t, out)
		assert.Equal(t, 1, len(in))
		answer, answerOK := in[0].AdditionalProperties["answer"]
		assert.True(t, answerOK)
		assert.NotNil(t, answer)
		answerAdditional, answerAdditionalOK := answer.(*qnamodels.Answer)
		assert.True(t, answerAdditionalOK)
		assert.Equal(t, "answer", answerAdditional.Result)
	})
}

type fakeQnAClient struct{}

func (c *fakeQnAClient) Answer(ctx context.Context,
	text, question string) (*ent.AnswerResult, error) {
	answer := &ent.AnswerResult{
		Text:     question,
		Question: question,
		Answer:   "answer",
	}
	return answer, nil
}

type fakeParamsHelper struct{}

func (h *fakeParamsHelper) GetQuestion(params interface{}) string {
	if params == nil {
		return ""
	}
	return "question"
}
