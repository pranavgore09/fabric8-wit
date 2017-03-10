package controller

import (
	"github.com/almighty/almighty-core/app"
	"github.com/almighty/almighty-core/application"
	"github.com/almighty/almighty-core/jsonapi"
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
)

// SpaceIterationsController implements the space-iterations resource.
type SpaceIterationsController struct {
	*goa.Controller
	db application.DB
}

// NewSpaceIterationsController creates a space-iterations controller.
func NewSpaceIterationsController(service *goa.Service, db application.DB) *SpaceIterationsController {
	return &SpaceIterationsController{Controller: service.NewController("SpaceIterationsController"), db: db}
}

// // Create runs the create action.
// func (c *SpaceIterationsController) Create(ctx *app.CreateSpaceIterationsContext) error {
// 	_, err := login.ContextIdentity(ctx)
// 	if err != nil {
// 		return jsonapi.JSONErrorResponse(ctx, goa.ErrUnauthorized(err.Error()))
// 	}
// 	spaceID, err := uuid.FromString(ctx.ID)
// 	if err != nil {
// 		return jsonapi.JSONErrorResponse(ctx, goa.ErrNotFound(err.Error()))
// 	}

// 	// Validate Request
// 	if ctx.Payload.Data == nil {
// 		return jsonapi.JSONErrorResponse(ctx, errors.NewBadParameterError("data", nil).Expected("not nil"))
// 	}
// 	reqIter := ctx.Payload.Data
// 	if reqIter.Attributes.Name == nil {
// 		return jsonapi.JSONErrorResponse(ctx, errors.NewBadParameterError("data.attributes.name", nil).Expected("not nil"))
// 	}

// 	return application.Transactional(c.db, func(appl application.Application) error {
// 		_, err = appl.Spaces().Load(ctx, spaceID)
// 		if err != nil {
// 			return jsonapi.JSONErrorResponse(ctx, goa.ErrNotFound(err.Error()))
// 		}

// 		newItr := iteration.Iteration{
// 			SpaceID: spaceID,
// 			Name:    *reqIter.Attributes.Name,
// 			StartAt: reqIter.Attributes.StartAt,
// 			EndAt:   reqIter.Attributes.EndAt,
// 		}
// 		if reqIter.Attributes.Description != nil {
// 			newItr.Description = reqIter.Attributes.Description
// 		}
// 		err = appl.Iterations().Create(ctx, &newItr)
// 		if err != nil {
// 			return jsonapi.JSONErrorResponse(ctx, err)
// 		}
// 		// For create, count will always be zero hence no need to query
// 		// by passing empty map, updateIterationsWithCounts will be able to put zero values
// 		wiCounts := make(map[string]workitem.WICountsPerIteration)
// 		var responseData *app.Iteration
// 		if newItr.Path.IsEmpty() == false {
// 			allParentsUUIDs := newItr.Path
// 			iterations, error := appl.Iterations().LoadMultiple(ctx, allParentsUUIDs)
// 			if error != nil {
// 				return jsonapi.JSONErrorResponse(ctx, err)
// 			}
// 			itrMap := make(iterationIDMap)
// 			for _, itr := range iterations {
// 				itrMap[itr.ID] = itr
// 			}
// 			responseData = ConvertIteration(ctx.RequestData, &newItr, parentPathResolver(itrMap), updateIterationsWithCounts(wiCounts))
// 		} else {
// 			responseData = ConvertIteration(ctx.RequestData, &newItr, updateIterationsWithCounts(wiCounts))
// 		}
// 		res := &app.IterationSingle{
// 			Data: responseData,
// 		}
// 		ctx.ResponseData.Header().Set("Location", rest.AbsoluteURL(ctx.RequestData, app.IterationHref(res.Data.ID)))
// 		return ctx.Created(res)
// 	})
// }

// List runs the list action.
func (c *SpaceIterationsController) List(ctx *app.ListSpaceIterationsContext) error {
	spaceID, err := uuid.FromString(ctx.ID)
	if err != nil {
		return jsonapi.JSONErrorResponse(ctx, goa.ErrNotFound(err.Error()))
	}

	return application.Transactional(c.db, func(appl application.Application) error {

		_, err = appl.Spaces().Load(ctx, spaceID)
		if err != nil {
			return jsonapi.JSONErrorResponse(ctx, goa.ErrNotFound(err.Error()))
		}
		iterations, err := appl.Iterations().List(ctx, spaceID)
		if err != nil {
			return jsonapi.JSONErrorResponse(ctx, err)
		}
		itrMap := make(iterationIDMap)
		for _, itr := range iterations {
			itrMap[itr.ID] = itr
		}
		// fetch extra information(counts of WI in each iteration of the space) to be added in response
		wiCounts, err := appl.WorkItems().GetCountsPerIteration(ctx, spaceID)
		if err != nil {
			return jsonapi.JSONErrorResponse(ctx, err)
		}
		res := &app.IterationList{}
		res.Data = ConvertIterations(ctx.RequestData, iterations, updateIterationsWithCounts(wiCounts), parentPathResolver(itrMap))
		return ctx.OK(res)
	})
}
