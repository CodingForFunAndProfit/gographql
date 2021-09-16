package gopher

import (
	"errors"
	"fmt"

	"github.com/CodingForFunAndProfit/gographql/job"
	"github.com/graphql-go/graphql"
)

type GopherService struct {
	gophers Repository
	jobs    job.Repository
}

// NewService is a factory that creates a new GopherService
func NewService(repo Repository, jobrepo job.Repository) GopherService {
	return GopherService{
		gophers: repo,
		jobs:    jobrepo,
	}
}

// ResolveGophers will be used to retrieve all available Gophers
func (gs GopherService) ResolveGophers(p graphql.ResolveParams) (interface{}, error) {
	// Fetch gophers from the Repository
	gophers, err := gs.gophers.GetGophers()
	if err != nil {
		return nil, err
	}
	return gophers, nil
}

func (gs *GopherService) ResolveJobs(p graphql.ResolveParams) (interface{}, error) {

	g, ok := p.Source.(Gopher)

	if !ok {
		return nil, errors.New("source was not a Gopher")
	}
	// Here we extract the Argument Company
	company := ""
	if value, ok := p.Args["company"]; ok {
		company, ok = value.(string)
		if !ok {
			return nil, errors.New("id has to be a string")
		}
	}

	// Find Jobs Based on the Gophers ID
	jobs, err := gs.jobs.GetJobs(g.ID, company)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (gs *GopherService) MutateJobs(p graphql.ResolveParams) (interface{}, error) {
	employee, err := grabStringArgument("employeeid", p.Args, true)
	if err != nil {
		return nil, err
	}
	jobid, err := grabStringArgument("jobid", p.Args, true)
	if err != nil {
		return nil, err
	}
	start, err := grabStringArgument("start", p.Args, false)
	if err != nil {
		return nil, err
	}
	end, err := grabStringArgument("end", p.Args, false)
	if err != nil {
		return nil, err
	}

	// Get the job
	job, err := gs.jobs.GetJob(employee, jobid)
	if err != nil {
		return nil, err
	}
	// Modify start and end date if they are set
	if start != "" {
		job.Start = start
	}

	if end != "" {
		job.End = end
	}
	// Update with new values
	return gs.jobs.Update(job)
}

// grabStringArgument is used to grab a string argument
func grabStringArgument(k string, args map[string]interface{}, required bool) (string, error) {
	// first check presense of arg
	if value, ok := args[k]; ok {
		// check string datatype
		v, o := value.(string)
		if !o {
			return "", fmt.Errorf("%s is not a string value", k)
		}
		return v, nil
	}
	if required {
		return "", fmt.Errorf("missing argument %s", k)
	}
	return "", nil
}
