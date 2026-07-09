package cli

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"try_parse_youtrack/internal/config"
	"try_parse_youtrack/internal/youtrack"
)

func Run(ctx context.Context, args []string, stdout io.Writer, stderr io.Writer) int {
	opts, err := parseOptions(args, stderr)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		fmt.Fprintln(stderr, err)
		return 2
	}

	client, err := youtrack.NewClient(youtrack.Config{
		Token:      opts.Config.Token,
		BaseURL:    opts.Config.BaseURL,
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
	})
	if err != nil {
		fmt.Fprintf(stderr, "create client: %v\n", err)
		return 1
	}

	if err := runRequests(ctx, client, opts, stdout); err != nil {
		fmt.Fprintf(stderr, "request: %v\n", err)
		return 1
	}

	return 0
}

type options struct {
	Config config.Config
	Query  string
	Top    string
}

func parseOptions(args []string, stderr io.Writer) (options, error) {
	defaultConfigPath, err := config.DefaultPath()
	if err != nil {
		return options{}, err
	}

	flags := flag.NewFlagSet("try_parse_youtrack", flag.ContinueOnError)
	flags.SetOutput(stderr)

	var (
		configPath = flags.String("config", defaultConfigPath, "path to JSON config file")
		baseURL    = flags.String("base-url", "", "YouTrack base URL")
		token      = flags.String("token", "", "YouTrack permanent token")
		query      = flags.String("query", "#Unresolved", "YouTrack issues search query")
		top        = flags.String("top", "10", "maximum number of issues to fetch")
	)
	flags.Usage = func() {
		fmt.Fprintf(flags.Output(), "usage: %s [flags]\n\n", flags.Name())
		fmt.Fprintln(flags.Output(), "Configuration priority: flags, env, config file.")
		fmt.Fprintf(flags.Output(), "Config file: %s\n", defaultConfigPath)
		fmt.Fprintf(flags.Output(), "Environment: %s, %s\n\n", config.EnvBaseURL, config.EnvToken)
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		return options{}, err
	}
	if flags.NArg() != 0 {
		return options{}, fmt.Errorf("unexpected arguments: %v", flags.Args())
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		return options{}, err
	}
	cfg = config.ApplyEnv(cfg)
	if *baseURL != "" {
		cfg.BaseURL = *baseURL
	}
	if *token != "" {
		cfg.Token = *token
	}

	if cfg.BaseURL == "" {
		return options{}, fmt.Errorf("missing base URL: set -base-url, %s, or config base_url", config.EnvBaseURL)
	}
	if cfg.Token == "" {
		return options{}, fmt.Errorf("missing token: set -token, %s, or config token", config.EnvToken)
	}

	return options{
		Config: cfg,
		Query:  *query,
		Top:    *top,
	}, nil
}

func runRequests(ctx context.Context, client *youtrack.Client, opts options, stdout io.Writer) error {
	query := url.Values{}
	query.Set(youtrack.QueryFields, "id,idReadable,summary,updated")
	query.Set(youtrack.QuerySearch, opts.Query)
	query.Set(youtrack.QueryTop, opts.Top)

	var issues []map[string]any
	if err := client.Get(ctx, youtrack.EndpointIssues, query, &issues); err != nil {
		return err
	}

	return json.NewEncoder(stdout).Encode(issues)
}
