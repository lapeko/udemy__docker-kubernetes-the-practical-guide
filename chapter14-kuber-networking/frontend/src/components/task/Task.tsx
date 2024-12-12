import {FC} from "react";

export const Task: FC<{task: string}> = ({task}) => <li>{task}</li>;
