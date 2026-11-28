## growthbook data-sources post

Create a data source

### Synopsis

Tests the connection and validates each assignment query against the warehouse before saving. The Managed Warehouse, Google Analytics and Event Forwarders are set up in the app.

```
growthbook data-sources post [flags]
```

### Examples

```
  growthbook data-sources post --name <value> --connection '{"type":"databricks","authType":"oauth-m2m","host":"muffled-squid.name","port":694435,"path":"/lib","catalog":"<value>"}'
```

### Options

```
  -a, --assignment-queries id                                                            Replaces the list. An entry with an existing id keeps `description`, `identifierTypes`, `includesNameColumns`, `dimensionColumns`, `projects` and `columnMapping` if left out; clear the others with `""`, `false` or `[]`. Each query is validated against the warehouse.
      --body string                                                                      Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --connection string                                                                JSON value (variants: postgres: { host: string, port: integer, database: string, user: string, ... }, redshift: { host: string, port: integer, database: string, user: string, ... }, vertica: { host: string, port: integer, database: string, user: string, ... }, mysql: { host: string, port: integer, database: string, user: string, ... }, mssql: { server: string, port: integer, database: string, user: string, ... }, snowflake: { account: string, accessUrl: string, username: string, password: string, ... }, bigquery: { authType: string, projectId: string, clientEmail: string, privateKey: string, ... }, athena: { authType: string, accessKeyId: string, secretAccessKey: string, assumeRoleARN: string, ... }, presto: { engine: string, authType: string, host: string, port: integer, ... }, databricks: { authType: string, token: string, oauthClientId: string, oauthClientSecret: string, ... }, clickhouse: { url: string, port: integer, username: string, password: string, ... }, mixpanel: { username: string, secret: string, projectId: string, server: string }, adobe_experience_platform_query_service: { host: string, port: integer, database: string, username: string, ... })
      --connection.adobe-experience-platform-query-service string                        AdobeExperiencePlatformQueryServiceConnection variant as JSON
      --connection.adobe-experience-platform-query-service.credential string             [required]
      --connection.adobe-experience-platform-query-service.database string               [required]
      --connection.adobe-experience-platform-query-service.host string                   [required]
      --connection.adobe-experience-platform-query-service.port int                      [required]
      --connection.adobe-experience-platform-query-service.technical-account-id string   [required]
      --connection.adobe-experience-platform-query-service.username string               [required]
      --connection.athena string                                                         AthenaConnection variant as JSON
      --connection.athena.access-key-id string                                           string value
      --connection.athena.assume-role-arn string                                         string value
      --connection.athena.auth-type auto                                                 auto uses the server's own credentials: self-hosted only (options: auto, accessKey, assumeRole)
      --connection.athena.bucket-uri string                                              [required]
      --connection.athena.catalog string                                                 string value
      --connection.athena.database string                                                string value
      --connection.athena.duration-seconds int                                           integer value
      --connection.athena.external-id string                                             string value
      --connection.athena.region string                                                  [required]
      --connection.athena.result-reuse-max-age-in-minutes string                         string value
      --connection.athena.role-session-name string                                       string value
      --connection.athena.secret-access-key string                                       string value
      --connection.athena.work-group string                                              string value
      --connection.bigquery string                                                       BigQueryConnection variant as JSON
      --connection.bigquery.api-endpoint string                                          string value
      --connection.bigquery.auth-type auto                                               auto uses the server's own credentials: self-hosted only (options: auto, json)
      --connection.bigquery.client-email string                                          string value
      --connection.bigquery.default-dataset string                                       [required]
      --connection.bigquery.default-project string                                       [required]
      --connection.bigquery.private-key string                                           string value
      --connection.bigquery.project-id string                                            string value
      --connection.bigquery.reservation string                                           string value
      --connection.clickhouse string                                                     ClickHouseConnection variant as JSON
      --connection.clickhouse.ca-cert string                                             string value
      --connection.clickhouse.ca-cert-path string                                        string value
      --connection.clickhouse.client-cert string                                         string value
      --connection.clickhouse.client-cert-chain-path *Path                               The *Path fields read PEM files on the GrowthBook server: self-hosted with a single organization only
      --connection.clickhouse.client-cert-path string                                    string value
      --connection.clickhouse.client-key string                                          string value
      --connection.clickhouse.client-key-path string                                     string value
      --connection.clickhouse.database string                                            [required]
      --connection.clickhouse.max-execution-time int                                     integer value
      --connection.clickhouse.password string                                            [required]
      --connection.clickhouse.port int                                                   [required]
      --connection.clickhouse.ssl                                                        boolean flag
      --connection.clickhouse.url string                                                 [required]
      --connection.clickhouse.username string                                            string value
      --connection.databricks string                                                     DatabricksConnection variant as JSON
      --connection.databricks.auth-type string                                           options: pat, oauth-m2m [required]
      --connection.databricks.catalog string                                             [required]
      --connection.databricks.host string                                                [required]
      --connection.databricks.oauth-client-id string                                     For oauth-m2m
      --connection.databricks.oauth-client-secret string                                 For oauth-m2m
      --connection.databricks.path string                                                [required]
      --connection.databricks.port int                                                   [required]
      --connection.databricks.token string                                               For pat
      --connection.mixpanel string                                                       MixpanelConnection variant as JSON
      --connection.mixpanel.project-id string                                            [required]
      --connection.mixpanel.secret string                                                [required]
      --connection.mixpanel.server string                                                options: standard, eu
      --connection.mixpanel.username string                                              [required]
      --connection.mssql string                                                          MssqlConnection variant as JSON
      --connection.mssql.database string                                                 [required]
      --connection.mssql.default-schema string                                           string value
      --connection.mssql.options string                                                  JSON object
      --connection.mssql.password string                                                 [required]
      --connection.mssql.port int                                                        [required]
      --connection.mssql.request-timeout int                                             integer value
      --connection.mssql.server string                                                   [required]
      --connection.mssql.user string                                                     [required]
      --connection.mysql string                                                          MysqlConnection variant as JSON
      --connection.mysql.ca-cert string                                                  string value
      --connection.mysql.client-cert string                                              string value
      --connection.mysql.client-key string                                               string value
      --connection.mysql.database string                                                 [required]
      --connection.mysql.host string                                                     [required]
      --connection.mysql.password string                                                 [required]
      --connection.mysql.port int                                                        [required]
      --connection.mysql.ssl                                                             boolean flag
      --connection.mysql.user string                                                     [required]
      --connection.postgres string                                                       PostgresConnection variant as JSON
      --connection.postgres.ca-cert string                                               string value
      --connection.postgres.client-cert string                                           string value
      --connection.postgres.client-key string                                            string value
      --connection.postgres.database string                                              [required]
      --connection.postgres.default-schema string                                        string value
      --connection.postgres.host string                                                  [required]
      --connection.postgres.password string                                              [required]
      --connection.postgres.port int                                                     [required]
      --connection.postgres.ssl                                                          boolean flag
      --connection.postgres.user string                                                  [required]
      --connection.presto string                                                         PrestoConnection variant as JSON
      --connection.presto.auth-type string                                               options: basicAuth, none, customAuth, kerberos
      --connection.presto.ca-cert string                                                 string value
      --connection.presto.catalog string                                                 [required]
      --connection.presto.client-cert string                                             string value
      --connection.presto.client-key string                                              string value
      --connection.presto.custom-auth string                                             string value
      --connection.presto.engine string                                                  options: presto, trino [required]
      --connection.presto.host string                                                    [required]
      --connection.presto.kerberos-client-principal string                               string value
      --connection.presto.kerberos-service-principal string                              string value
      --connection.presto.kerberos-user string                                           string value
      --connection.presto.password string                                                For basicAuth
      --connection.presto.port int                                                       [required]
      --connection.presto.request-timeout int                                            integer value
      --connection.presto.schema string                                                  [required]
      --connection.presto.source string                                                  string value
      --connection.presto.ssl                                                            boolean flag
      --connection.presto.trino-user string                                              Trino only
      --connection.presto.user string                                                    Presto only
      --connection.presto.username string                                                For basicAuth
      --connection.redshift string                                                       RedshiftConnection variant as JSON
      --connection.redshift.ca-cert string                                               string value
      --connection.redshift.client-cert string                                           string value
      --connection.redshift.client-key string                                            string value
      --connection.redshift.database string                                              [required]
      --connection.redshift.default-schema string                                        string value
      --connection.redshift.host string                                                  [required]
      --connection.redshift.password string                                              [required]
      --connection.redshift.port int                                                     [required]
      --connection.redshift.ssl                                                          boolean flag
      --connection.redshift.user string                                                  [required]
      --connection.snowflake string                                                      SnowflakeConnection variant as JSON
      --connection.snowflake.access-url string                                           string value
      --connection.snowflake.account string                                              [required]
      --connection.snowflake.auth-method string                                          options: password, key-pair, workload-identity
      --connection.snowflake.database string                                             [required]
      --connection.snowflake.password password                                           For password auth
      --connection.snowflake.private-key string                                          string value
      --connection.snowflake.private-key-password string                                 string value
      --connection.snowflake.role string                                                 string value
      --connection.snowflake.schema string                                               [required]
      --connection.snowflake.username string                                             [required]
      --connection.snowflake.warehouse string                                            string value
      --connection.snowflake.workload-identity-provider string                           options: AWS, AZURE, GCP
      --connection.vertica string                                                        VerticaConnection variant as JSON
      --connection.vertica.ca-cert string                                                string value
      --connection.vertica.client-cert string                                            string value
      --connection.vertica.client-key string                                             string value
      --connection.vertica.database string                                               [required]
      --connection.vertica.default-schema string                                         string value
      --connection.vertica.host string                                                   [required]
      --connection.vertica.password string                                               [required]
      --connection.vertica.port int                                                      [required]
      --connection.vertica.ssl                                                           boolean flag
      --connection.vertica.user string                                                   [required]
      --description string                                                               string value
  -e, --event-tracker eventForwarder                                                     eventForwarder is only accepted unchanged (options: segment, snowplow, jitsu, freshpaint, ga4, fullstory, matomo, heap, rudderstack, amplitude, mparticle, mixpanel, firebase, keen, clevertap, langfuse, phoenix, eventForwarder, custom)
  -h, --help                                                                             help for post
      --identifier-join-queries string                                                   Replaces the list
      --identifier-types ""                                                              Replaces the list. An entry matching an existing id keeps its description if left out; send "" to clear it.
      --max-concurrent-queries int                                                       integer value
      --mixpanel-settings string                                                         JSON object
  -n, --name string                                                                      [required]
      --pipeline-settings allowWriting                                                   Merged into the stored settings: fields left out keep their values, but any change must leave allowWriting, `mode`, `writeDataset` and `unitsTableRetentionHours` set, with a non-empty `writeDataset` when `allowWriting` is true. Send `[]` to clear `excludedExperimentIds` or `incrementalOptInExperimentIds`, and `includedExperimentIds: null` to apply the mode to every experiment again (`[]` applies it to none).
      --project-ids stringArray                                                          list of values
      --query-cache-ttl-mins int                                                         integer value
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string            Bearer auth token: your Secret Key or Personal Access Token, sent as an Authorization Bearer header.
      --color string                  Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                         Log request and response diagnostics to stderr
      --domain string                 Server template variable: domain
      --dry-run                       Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray            Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers               Include HTTP response headers in the output
  -q, --jq string                     Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive                Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
      --no-update-check               Disable the once-a-day check for a newer CLI version
  -o, --output-format string          Specify the output format. Options: pretty, json, yaml, toon. (default "pretty")
      --password string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username
```

### SEE ALSO

* [growthbook data-sources](growthbook_data-sources.md)	 - How GrowthBook connects and queries your data, including cached database schema metadata (information schemas) for tables and columns
