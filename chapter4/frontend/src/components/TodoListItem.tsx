import {FC} from "react";
import {TodoResponse} from "../utils/TodoResponse.ts";

export const TodoListItem: FC<{todos: TodoResponse[]}> = ({todos}) => <ul>
  {todos.map(todo => <li key={todo.id}>{todo.title}</li>)}
</ul>