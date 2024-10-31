import {FC} from "react";
import {TodoResponse} from "../utils/TodoResponse.ts";

export const TodoListItem: FC<{todo: TodoResponse, onClick: (id: string) => void}> = ({todo, onClick}) => (
  <li key={todo.id} onClick={() => onClick(todo.id)}>{todo.title}</li>
);
