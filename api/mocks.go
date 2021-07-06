// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package api

import (
	"context"
	"sync"
)

var (
	lockElasticSearcherMockGetStatus   sync.RWMutex
	lockElasticSearcherMockMultiSearch sync.RWMutex
	lockElasticSearcherMockSearch      sync.RWMutex
)

// Ensure, that ElasticSearcherMock does implement ElasticSearcher.
// If this is not the case, regenerate this file with moq.
var _ ElasticSearcher = &ElasticSearcherMock{}

// ElasticSearcherMock is a mock implementation of ElasticSearcher.
//
//     func TestSomethingThatUsesElasticSearcher(t *testing.T) {
//
//         // make and configure a mocked ElasticSearcher
//         mockedElasticSearcher := &ElasticSearcherMock{
//             GetStatusFunc: func(ctx context.Context) ([]byte, error) {
// 	               panic("mock out the GetStatus method")
//             },
//             MultiSearchFunc: func(ctx context.Context, index string, docType string, request []byte) ([]byte, error) {
// 	               panic("mock out the MultiSearch method")
//             },
//             SearchFunc: func(ctx context.Context, index string, docType string, request []byte) ([]byte, error) {
// 	               panic("mock out the Search method")
//             },
//         }
//
//         // use mockedElasticSearcher in code that requires ElasticSearcher
//         // and then make assertions.
//
//     }
type ElasticSearcherMock struct {
	// GetStatusFunc mocks the GetStatus method.
	GetStatusFunc func(ctx context.Context) ([]byte, error)

	// MultiSearchFunc mocks the MultiSearch method.
	MultiSearchFunc func(ctx context.Context, index string, docType string, request []byte) ([]byte, error)

	// SearchFunc mocks the Search method.
	SearchFunc func(ctx context.Context, index string, docType string, request []byte) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetStatus holds details about calls to the GetStatus method.
		GetStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// MultiSearch holds details about calls to the MultiSearch method.
		MultiSearch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Index is the index argument value.
			Index string
			// DocType is the docType argument value.
			DocType string
			// Request is the request argument value.
			Request []byte
		}
		// Search holds details about calls to the Search method.
		Search []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Index is the index argument value.
			Index string
			// DocType is the docType argument value.
			DocType string
			// Request is the request argument value.
			Request []byte
		}
	}
}

// GetStatus calls GetStatusFunc.
func (mock *ElasticSearcherMock) GetStatus(ctx context.Context) ([]byte, error) {
	if mock.GetStatusFunc == nil {
		panic("ElasticSearcherMock.GetStatusFunc: method is nil but ElasticSearcher.GetStatus was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockElasticSearcherMockGetStatus.Lock()
	mock.calls.GetStatus = append(mock.calls.GetStatus, callInfo)
	lockElasticSearcherMockGetStatus.Unlock()
	return mock.GetStatusFunc(ctx)
}

// GetStatusCalls gets all the calls that were made to GetStatus.
// Check the length with:
//     len(mockedElasticSearcher.GetStatusCalls())
func (mock *ElasticSearcherMock) GetStatusCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockElasticSearcherMockGetStatus.RLock()
	calls = mock.calls.GetStatus
	lockElasticSearcherMockGetStatus.RUnlock()
	return calls
}

// MultiSearch calls MultiSearchFunc.
func (mock *ElasticSearcherMock) MultiSearch(ctx context.Context, index string, docType string, request []byte) ([]byte, error) {
	if mock.MultiSearchFunc == nil {
		panic("ElasticSearcherMock.MultiSearchFunc: method is nil but ElasticSearcher.MultiSearch was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Index   string
		DocType string
		Request []byte
	}{
		Ctx:     ctx,
		Index:   index,
		DocType: docType,
		Request: request,
	}
	lockElasticSearcherMockMultiSearch.Lock()
	mock.calls.MultiSearch = append(mock.calls.MultiSearch, callInfo)
	lockElasticSearcherMockMultiSearch.Unlock()
	return mock.MultiSearchFunc(ctx, index, docType, request)
}

// MultiSearchCalls gets all the calls that were made to MultiSearch.
// Check the length with:
//     len(mockedElasticSearcher.MultiSearchCalls())
func (mock *ElasticSearcherMock) MultiSearchCalls() []struct {
	Ctx     context.Context
	Index   string
	DocType string
	Request []byte
} {
	var calls []struct {
		Ctx     context.Context
		Index   string
		DocType string
		Request []byte
	}
	lockElasticSearcherMockMultiSearch.RLock()
	calls = mock.calls.MultiSearch
	lockElasticSearcherMockMultiSearch.RUnlock()
	return calls
}

// Search calls SearchFunc.
func (mock *ElasticSearcherMock) Search(ctx context.Context, index string, docType string, request []byte) ([]byte, error) {
	if mock.SearchFunc == nil {
		panic("ElasticSearcherMock.SearchFunc: method is nil but ElasticSearcher.Search was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Index   string
		DocType string
		Request []byte
	}{
		Ctx:     ctx,
		Index:   index,
		DocType: docType,
		Request: request,
	}
	lockElasticSearcherMockSearch.Lock()
	mock.calls.Search = append(mock.calls.Search, callInfo)
	lockElasticSearcherMockSearch.Unlock()
	return mock.SearchFunc(ctx, index, docType, request)
}

// SearchCalls gets all the calls that were made to Search.
// Check the length with:
//     len(mockedElasticSearcher.SearchCalls())
func (mock *ElasticSearcherMock) SearchCalls() []struct {
	Ctx     context.Context
	Index   string
	DocType string
	Request []byte
} {
	var calls []struct {
		Ctx     context.Context
		Index   string
		DocType string
		Request []byte
	}
	lockElasticSearcherMockSearch.RLock()
	calls = mock.calls.Search
	lockElasticSearcherMockSearch.RUnlock()
	return calls
}

var (
	lockQueryBuilderMockBuildSearchQuery sync.RWMutex
)

// Ensure, that QueryBuilderMock does implement QueryBuilder.
// If this is not the case, regenerate this file with moq.
var _ QueryBuilder = &QueryBuilderMock{}

// QueryBuilderMock is a mock implementation of QueryBuilder.
//
//     func TestSomethingThatUsesQueryBuilder(t *testing.T) {
//
//         // make and configure a mocked QueryBuilder
//         mockedQueryBuilder := &QueryBuilderMock{
//             BuildSearchQueryFunc: func(ctx context.Context, q string, contentTypes string, sort string, limit int, offset int) ([]byte, error) {
// 	               panic("mock out the BuildSearchQuery method")
//             },
//         }
//
//         // use mockedQueryBuilder in code that requires QueryBuilder
//         // and then make assertions.
//
//     }
type QueryBuilderMock struct {
	// BuildSearchQueryFunc mocks the BuildSearchQuery method.
	BuildSearchQueryFunc func(ctx context.Context, q string, contentTypes string, sort string, limit int, offset int) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// BuildSearchQuery holds details about calls to the BuildSearchQuery method.
		BuildSearchQuery []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Q is the q argument value.
			Q string
			// ContentTypes is the contentTypes argument value.
			ContentTypes string
			// Sort is the sort argument value.
			Sort string
			// Limit is the limit argument value.
			Limit int
			// Offset is the offset argument value.
			Offset int
		}
	}
}

// BuildSearchQuery calls BuildSearchQueryFunc.
func (mock *QueryBuilderMock) BuildSearchQuery(ctx context.Context, q string, contentTypes string, sort string, limit int, offset int) ([]byte, error) {
	if mock.BuildSearchQueryFunc == nil {
		panic("QueryBuilderMock.BuildSearchQueryFunc: method is nil but QueryBuilder.BuildSearchQuery was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Q            string
		ContentTypes string
		Sort         string
		Limit        int
		Offset       int
	}{
		Ctx:          ctx,
		Q:            q,
		ContentTypes: contentTypes,
		Sort:         sort,
		Limit:        limit,
		Offset:       offset,
	}
	lockQueryBuilderMockBuildSearchQuery.Lock()
	mock.calls.BuildSearchQuery = append(mock.calls.BuildSearchQuery, callInfo)
	lockQueryBuilderMockBuildSearchQuery.Unlock()
	return mock.BuildSearchQueryFunc(ctx, q, contentTypes, sort, limit, offset)
}

// BuildSearchQueryCalls gets all the calls that were made to BuildSearchQuery.
// Check the length with:
//     len(mockedQueryBuilder.BuildSearchQueryCalls())
func (mock *QueryBuilderMock) BuildSearchQueryCalls() []struct {
	Ctx          context.Context
	Q            string
	ContentTypes string
	Sort         string
	Limit        int
	Offset       int
} {
	var calls []struct {
		Ctx          context.Context
		Q            string
		ContentTypes string
		Sort         string
		Limit        int
		Offset       int
	}
	lockQueryBuilderMockBuildSearchQuery.RLock()
	calls = mock.calls.BuildSearchQuery
	lockQueryBuilderMockBuildSearchQuery.RUnlock()
	return calls
}

var (
	lockResponseTransformerMockTransformSearchResponse sync.RWMutex
)

// Ensure, that ResponseTransformerMock does implement ResponseTransformer.
// If this is not the case, regenerate this file with moq.
var _ ResponseTransformer = &ResponseTransformerMock{}

// ResponseTransformerMock is a mock implementation of ResponseTransformer.
//
//     func TestSomethingThatUsesResponseTransformer(t *testing.T) {
//
//         // make and configure a mocked ResponseTransformer
//         mockedResponseTransformer := &ResponseTransformerMock{
//             TransformSearchResponseFunc: func(ctx context.Context, responseData []byte, query string) ([]byte, error) {
// 	               panic("mock out the TransformSearchResponse method")
//             },
//         }
//
//         // use mockedResponseTransformer in code that requires ResponseTransformer
//         // and then make assertions.
//
//     }
type ResponseTransformerMock struct {
	// TransformSearchResponseFunc mocks the TransformSearchResponse method.
	TransformSearchResponseFunc func(ctx context.Context, responseData []byte, query string) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// TransformSearchResponse holds details about calls to the TransformSearchResponse method.
		TransformSearchResponse []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ResponseData is the responseData argument value.
			ResponseData []byte
			// Query is the query argument value.
			Query string
		}
	}
}

// TransformSearchResponse calls TransformSearchResponseFunc.
func (mock *ResponseTransformerMock) TransformSearchResponse(ctx context.Context, responseData []byte, query string) ([]byte, error) {
	if mock.TransformSearchResponseFunc == nil {
		panic("ResponseTransformerMock.TransformSearchResponseFunc: method is nil but ResponseTransformer.TransformSearchResponse was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		ResponseData []byte
		Query        string
	}{
		Ctx:          ctx,
		ResponseData: responseData,
		Query:        query,
	}
	lockResponseTransformerMockTransformSearchResponse.Lock()
	mock.calls.TransformSearchResponse = append(mock.calls.TransformSearchResponse, callInfo)
	lockResponseTransformerMockTransformSearchResponse.Unlock()
	return mock.TransformSearchResponseFunc(ctx, responseData, query)
}

// TransformSearchResponseCalls gets all the calls that were made to TransformSearchResponse.
// Check the length with:
//     len(mockedResponseTransformer.TransformSearchResponseCalls())
func (mock *ResponseTransformerMock) TransformSearchResponseCalls() []struct {
	Ctx          context.Context
	ResponseData []byte
	Query        string
} {
	var calls []struct {
		Ctx          context.Context
		ResponseData []byte
		Query        string
	}
	lockResponseTransformerMockTransformSearchResponse.RLock()
	calls = mock.calls.TransformSearchResponse
	lockResponseTransformerMockTransformSearchResponse.RUnlock()
	return calls
}
