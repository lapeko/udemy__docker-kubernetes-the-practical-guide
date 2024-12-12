import {FormEvent, useState, FC} from 'react';

export const InputForm: FC<{taskCreatedHandler: (task: string) => void}> = ({taskCreatedHandler}) => {
    const [pending, setPending] = useState(false);
    const [task, setTask] = useState('');
    const submit = (event: FormEvent<HTMLFormElement>) => {
        setPending(true);
        event.preventDefault();
        fetch("/api/tasks", {method: 'POST', body: JSON.stringify({task}), headers: {
            'Authorization': 'Bearer abc',
            'Content-Type': 'application/json'
        }})
            .then<{ createdTask: string, message: string }>((res) => res.json())
            .then((data) => {
                taskCreatedHandler(data.createdTask);
                setTask('');
            })
            .catch((err) => console.log(err))
            .finally(() => setPending(false));
    };
    return (
        <div className="container pt-4 mb-4">
            <form onSubmit={submit}>
                <label htmlFor="task-input" className="form-label">Task name</label>
                <input
                    className="form-control mb-2"
                    id="task-input"
                    placeholder="Provide a task"
                    value={task}
                    onChange={(e) => setTask(e.target.value)}
                />
                <button type="submit" className="btn btn-primary" disabled={pending}>Submit</button>
            </form>
        </div>
    )
}