import React, { useEffect, useState } from 'react';
import { NavLink, Route, Switch, useParams } from 'react-router-dom';
import { getTask } from "../../services/task.service";
import './TaskPage.css'
import SendTaskSolution from "./SolutionFied/SolutionFied";
import CheckResult from "./GetResult/GetResult";
import GetHistory from "./TaskHistory/TaskHistory";

function TaskPage() {
    const { id } = useParams();
    const [task, setTask] = useState({});

    useEffect(() => {
        getTask(id)
            .then(data => {
                setTask(data)
            })
    }, [])

    return(
        <div className="Task-row">
            <SendTaskSolution />
            <div className="Task-left">
                <nav>
                    <ul className="Task-menu">
                        <li><NavLink exact to={"/task/"+id} activeClassName="active">Description</NavLink></li>
                        <li><NavLink to={"/task/"+id+"/output"} activeClassName="active">Output</NavLink></li>
                        <li><NavLink to={"/task/"+id+"/history"} activeClassName="active">History</NavLink></li>
                    </ul>
                </nav>
                <div className="Task-left-content">
                    <Switch>
                        <Route exact path="/task/:id">
                            <h3>{task.name}</h3>
                            <p>{task.description}</p>
                        </Route>
                        <Route exact path="/task/:id/output">
                            <CheckResult />
                        </Route>
                        <Route exact path="/task/:id/history">
                            <GetHistory />
                        </Route>
                    </Switch>
                </div>
                <div className="Task-left-footer"> </div>
            </div>
        </div>
    )
}

export default TaskPage;