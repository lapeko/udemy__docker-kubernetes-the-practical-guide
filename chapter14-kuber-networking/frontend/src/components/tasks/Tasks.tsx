import {FC} from "react";
import {Task} from "../task/Task.tsx";

export const Tasks: FC<{tasks: string[]}> = ({tasks}) => (
    <ul className="container">
        {tasks.map((task, idx) => <Task key={idx} task={task} />)}
    </ul>
);