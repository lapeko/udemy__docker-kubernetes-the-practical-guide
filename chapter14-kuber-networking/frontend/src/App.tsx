import {useEffect, useState} from "react";

import {InputForm} from "./components/input-form/InputForm.tsx";
import {Tasks} from "./components/tasks/Tasks.tsx";

export const App = () => {
    const [tasks, setTasks] = useState([] as string[]);
    useEffect(() => {
        console.log(tasks);
    }, [tasks]);
    useEffect(() => {
        fetch("/api/tasks", {headers: {'Authorization': 'Bearer abc', 'Content-Type': 'application/json'}})
            .then<{ tasks: string[] }>((res) => res.json())
            .then((data) => setTasks(data.tasks))
            .catch((err) => console.log(err));
    }, []);

    const taskCreatedHandler = (task: string) => setTasks(tasks => [...tasks, task]);

    return (
        <>
            <InputForm taskCreatedHandler={taskCreatedHandler} />
            <Tasks tasks={tasks} />
        </>
    );
};
