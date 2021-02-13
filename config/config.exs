# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
use Mix.Config

config :gitodo,
  ecto_repos: [Gitodo.Repo]

# Configures the endpoint
config :gitodo, GitodoWeb.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "+c2z69SYv+EKZxPFoTDOafvO1tOEyJKXQH/xFzJmq3aZ2vOD4V6zFdwU+nAQ6qrO",
  render_errors: [view: GitodoWeb.ErrorView, accepts: ~w(html json), layout: false],
  pubsub_server: Gitodo.PubSub,
  live_view: [signing_salt: "P47yFMK9"]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env()}.exs"
