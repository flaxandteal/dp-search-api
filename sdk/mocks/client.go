// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/v2/health"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-search-api/models"
	"github.com/ONSdigital/dp-search-api/sdk"
	"github.com/ONSdigital/dp-search-api/sdk/errors"
	"github.com/ONSdigital/dp-search-api/transformer"
	"sync"
)

var (
	lockClienterMockChecker                   sync.RWMutex
	lockClienterMockCreateIndex               sync.RWMutex
	lockClienterMockGetReleaseCalendarEntries sync.RWMutex
	lockClienterMockGetSearch                 sync.RWMutex
	lockClienterMockHealth                    sync.RWMutex
	lockClienterMockURL                       sync.RWMutex
)

// Ensure, that ClienterMock does implement sdk.Clienter.
// If this is not the case, regenerate this file with moq.
var _ sdk.Clienter = &ClienterMock{}

// ClienterMock is a mock implementation of sdk.Clienter.
//
//     func TestSomethingThatUsesClienter(t *testing.T) {
//
//         // make and configure a mocked sdk.Clienter
//         mockedClienter := &ClienterMock{
//             CheckerFunc: func(ctx context.Context, check *healthcheck.CheckState) error {
// 	               panic("mock out the Checker method")
//             },
//             CreateIndexFunc: func(ctx context.Context, options sdk.Options) (*models.CreateIndexResponse, errors.Error) {
// 	               panic("mock out the CreateIndex method")
//             },
//             GetReleaseCalendarEntriesFunc: func(ctx context.Context, options sdk.Options) (*transformer.ReleaseTransformer, errors.Error) {
// 	               panic("mock out the GetReleaseCalendarEntries method")
//             },
//             GetSearchFunc: func(ctx context.Context, options sdk.Options) (*models.SearchResponse, errors.Error) {
// 	               panic("mock out the GetSearch method")
//             },
//             HealthFunc: func() *health.Client {
// 	               panic("mock out the Health method")
//             },
//             URLFunc: func() string {
// 	               panic("mock out the URL method")
//             },
//         }
//
//         // use mockedClienter in code that requires sdk.Clienter
//         // and then make assertions.
//
//     }
type ClienterMock struct {
	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, check *healthcheck.CheckState) error

	// CreateIndexFunc mocks the CreateIndex method.
	CreateIndexFunc func(ctx context.Context, options sdk.Options) (*models.CreateIndexResponse, errors.Error)

	// GetReleaseCalendarEntriesFunc mocks the GetReleaseCalendarEntries method.
	GetReleaseCalendarEntriesFunc func(ctx context.Context, options sdk.Options) (*transformer.ReleaseTransformer, errors.Error)

	// GetSearchFunc mocks the GetSearch method.
	GetSearchFunc func(ctx context.Context, options sdk.Options) (*models.SearchResponse, errors.Error)

	// HealthFunc mocks the Health method.
	HealthFunc func() *health.Client

	// URLFunc mocks the URL method.
	URLFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Check is the check argument value.
			Check *healthcheck.CheckState
		}
		// CreateIndex holds details about calls to the CreateIndex method.
		CreateIndex []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Options is the options argument value.
			Options sdk.Options
		}
		// GetReleaseCalendarEntries holds details about calls to the GetReleaseCalendarEntries method.
		GetReleaseCalendarEntries []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Options is the options argument value.
			Options sdk.Options
		}
		// GetSearch holds details about calls to the GetSearch method.
		GetSearch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Options is the options argument value.
			Options sdk.Options
		}
		// Health holds details about calls to the Health method.
		Health []struct {
		}
		// URL holds details about calls to the URL method.
		URL []struct {
		}
	}
}

// Checker calls CheckerFunc.
func (mock *ClienterMock) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("ClienterMock.CheckerFunc: method is nil but Clienter.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Check *healthcheck.CheckState
	}{
		Ctx:   ctx,
		Check: check,
	}
	lockClienterMockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	lockClienterMockChecker.Unlock()
	return mock.CheckerFunc(ctx, check)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedClienter.CheckerCalls())
func (mock *ClienterMock) CheckerCalls() []struct {
	Ctx   context.Context
	Check *healthcheck.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		Check *healthcheck.CheckState
	}
	lockClienterMockChecker.RLock()
	calls = mock.calls.Checker
	lockClienterMockChecker.RUnlock()
	return calls
}

// CreateIndex calls CreateIndexFunc.
func (mock *ClienterMock) CreateIndex(ctx context.Context, options sdk.Options) (*models.CreateIndexResponse, errors.Error) {
	if mock.CreateIndexFunc == nil {
		panic("ClienterMock.CreateIndexFunc: method is nil but Clienter.CreateIndex was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Options sdk.Options
	}{
		Ctx:     ctx,
		Options: options,
	}
	lockClienterMockCreateIndex.Lock()
	mock.calls.CreateIndex = append(mock.calls.CreateIndex, callInfo)
	lockClienterMockCreateIndex.Unlock()
	return mock.CreateIndexFunc(ctx, options)
}

// CreateIndexCalls gets all the calls that were made to CreateIndex.
// Check the length with:
//     len(mockedClienter.CreateIndexCalls())
func (mock *ClienterMock) CreateIndexCalls() []struct {
	Ctx     context.Context
	Options sdk.Options
} {
	var calls []struct {
		Ctx     context.Context
		Options sdk.Options
	}
	lockClienterMockCreateIndex.RLock()
	calls = mock.calls.CreateIndex
	lockClienterMockCreateIndex.RUnlock()
	return calls
}

// GetReleaseCalendarEntries calls GetReleaseCalendarEntriesFunc.
func (mock *ClienterMock) GetReleaseCalendarEntries(ctx context.Context, options sdk.Options) (*transformer.ReleaseTransformer, errors.Error) {
	if mock.GetReleaseCalendarEntriesFunc == nil {
		panic("ClienterMock.GetReleaseCalendarEntriesFunc: method is nil but Clienter.GetReleaseCalendarEntries was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Options sdk.Options
	}{
		Ctx:     ctx,
		Options: options,
	}
	lockClienterMockGetReleaseCalendarEntries.Lock()
	mock.calls.GetReleaseCalendarEntries = append(mock.calls.GetReleaseCalendarEntries, callInfo)
	lockClienterMockGetReleaseCalendarEntries.Unlock()
	return mock.GetReleaseCalendarEntriesFunc(ctx, options)
}

// GetReleaseCalendarEntriesCalls gets all the calls that were made to GetReleaseCalendarEntries.
// Check the length with:
//     len(mockedClienter.GetReleaseCalendarEntriesCalls())
func (mock *ClienterMock) GetReleaseCalendarEntriesCalls() []struct {
	Ctx     context.Context
	Options sdk.Options
} {
	var calls []struct {
		Ctx     context.Context
		Options sdk.Options
	}
	lockClienterMockGetReleaseCalendarEntries.RLock()
	calls = mock.calls.GetReleaseCalendarEntries
	lockClienterMockGetReleaseCalendarEntries.RUnlock()
	return calls
}

// GetSearch calls GetSearchFunc.
func (mock *ClienterMock) GetSearch(ctx context.Context, options sdk.Options) (*models.SearchResponse, errors.Error) {
	if mock.GetSearchFunc == nil {
		panic("ClienterMock.GetSearchFunc: method is nil but Clienter.GetSearch was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Options sdk.Options
	}{
		Ctx:     ctx,
		Options: options,
	}
	lockClienterMockGetSearch.Lock()
	mock.calls.GetSearch = append(mock.calls.GetSearch, callInfo)
	lockClienterMockGetSearch.Unlock()
	return mock.GetSearchFunc(ctx, options)
}

// GetSearchCalls gets all the calls that were made to GetSearch.
// Check the length with:
//     len(mockedClienter.GetSearchCalls())
func (mock *ClienterMock) GetSearchCalls() []struct {
	Ctx     context.Context
	Options sdk.Options
} {
	var calls []struct {
		Ctx     context.Context
		Options sdk.Options
	}
	lockClienterMockGetSearch.RLock()
	calls = mock.calls.GetSearch
	lockClienterMockGetSearch.RUnlock()
	return calls
}

// Health calls HealthFunc.
func (mock *ClienterMock) Health() *health.Client {
	if mock.HealthFunc == nil {
		panic("ClienterMock.HealthFunc: method is nil but Clienter.Health was just called")
	}
	callInfo := struct {
	}{}
	lockClienterMockHealth.Lock()
	mock.calls.Health = append(mock.calls.Health, callInfo)
	lockClienterMockHealth.Unlock()
	return mock.HealthFunc()
}

// HealthCalls gets all the calls that were made to Health.
// Check the length with:
//     len(mockedClienter.HealthCalls())
func (mock *ClienterMock) HealthCalls() []struct {
} {
	var calls []struct {
	}
	lockClienterMockHealth.RLock()
	calls = mock.calls.Health
	lockClienterMockHealth.RUnlock()
	return calls
}

// URL calls URLFunc.
func (mock *ClienterMock) URL() string {
	if mock.URLFunc == nil {
		panic("ClienterMock.URLFunc: method is nil but Clienter.URL was just called")
	}
	callInfo := struct {
	}{}
	lockClienterMockURL.Lock()
	mock.calls.URL = append(mock.calls.URL, callInfo)
	lockClienterMockURL.Unlock()
	return mock.URLFunc()
}

// URLCalls gets all the calls that were made to URL.
// Check the length with:
//     len(mockedClienter.URLCalls())
func (mock *ClienterMock) URLCalls() []struct {
} {
	var calls []struct {
	}
	lockClienterMockURL.RLock()
	calls = mock.calls.URL
	lockClienterMockURL.RUnlock()
	return calls
}
