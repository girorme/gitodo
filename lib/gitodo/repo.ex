defmodule Gitodo.Repo do
  use Ecto.Repo,
    otp_app: :gitodo,
    adapter: Ecto.Adapters.Postgres
end
