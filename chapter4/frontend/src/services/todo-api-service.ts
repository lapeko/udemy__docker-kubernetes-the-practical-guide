import {TodoResponse} from "../utils/TodoResponse.ts";
import {TodoRequest} from "../utils/TodoRequest.ts";
import {API_URL} from "../environments/env.ts"

type Response<T> = {
  OK: boolean,
  payload: T | null,
  error: string
}

export const TodoApiService = {
  fetchTodos: () => fetch(API_URL)
    .then(res => res.json())
    .then((json: Response<{ todos: TodoResponse[] }>) => json.payload?.todos),

  createTodo: (todo: TodoRequest) => fetch(API_URL, {body: JSON.stringify(todo), method: "POST"})
    .then(res => res.json())
    .then((json: Response<{ InsertedID: string }>) => json.payload?.InsertedID),

  deleteTodo: (id: string) => fetch(`${API_URL}/${id}`, {method: "DELETE"})
    .then(res => res.json())
    .then((json: Response<{ DeletedCount: number }>) => json.payload?.DeletedCount),
}