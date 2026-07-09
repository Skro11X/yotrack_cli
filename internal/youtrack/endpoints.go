package youtrack

const (
	EndpointIssues        = "/api/issues"
	EndpointIssue         = "/api/issues/%s"
	EndpointIssueComments = "/api/issues/%s/comments"
	EndpointIssueLinks    = "/api/issues/%s/links"
	EndpointIssueTags     = "/api/issues/%s/tags"

	EndpointUsers         = "/api/users"
	EndpointCurrentUser   = "/api/users/me"
	EndpointGroups        = "/api/groups"
	EndpointProjects      = "/api/admin/projects"
	EndpointProject       = "/api/admin/projects/%s"
	EndpointProjectIssues = "/api/admin/projects/%s/issues"

	EndpointCommands     = "/api/commands"
	EndpointTags         = "/api/tags"
	EndpointAgileBoards  = "/api/agiles"
	EndpointSavedQueries = "/api/savedQueries"
)

const (
	QueryFields = "fields"
	QuerySearch = "query"
	QueryTop    = "$top"
	QuerySkip   = "$skip"
)
