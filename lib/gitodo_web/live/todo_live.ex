defmodule GitodoWeb.TodoLive do
  use GitodoWeb, :live_view

  alias Gitodo.Todos

  def handle_info({Todos, [:todo | _], _}, socket) do
    {:noreply, fetch(socket)}
  end

  def mount(_params, _session, socket) do
    Todos.subscribe()

    {:ok, fetch(socket)}
  end

  def handle_event("add", %{"todo" => todo}, socket) do
    Todos.create_todo(todo)

    {:noreply, fetch(socket)}
  end

  def handle_event("toggle_done", %{"id" => id}, socket) do
    todo = Todos.get_todo!(id)
    Todos.update_todo(todo, %{done: !todo.done})

    {:noreply, socket}
  end

  defp fetch(socket) do
    assign(socket, todos: Todos.list_todos())
  end
end
