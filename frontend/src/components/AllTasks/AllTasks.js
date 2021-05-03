import './AllTasks.css'
import React, { useState, useEffect } from "react";
import AuthService from "../../services/auth.service";
import {Link} from "react-router-dom";

const API_URL = process.env.REACT_APP_SERVER_URL;

const AllTasks = () => {
    const [tasks, setTasks] = useState([]);
    const [currentUser, setCurrentUser] = useState(undefined);

    useEffect(() => {
        const user = AuthService.getCurrentUser();

        if (user) {
            setCurrentUser(user);
        }
    }, []);

    useEffect(() => {
        const getTasks = async () => {
            try {
                const response = await fetch(API_URL+'/tasks');

                const responseData = await response.json();

                setTasks(responseData);
            } catch (error) {
                console.error(error);
            }
        };

        getTasks().then();
    }, []);

    return (
        <div className="container">
            <h1 className="Welcome">Ready to challenge?</h1>
            <div className="Tasks">
                {tasks.map(function (task, index) {
                    return (
                        <div className="Single-task" key={index}>
                            <div className="Single-task-header">{task.name}</div>
                            <hr/>
                            <div className="Single-task-body">
                                <p>{task.description}</p>
                            </div>
                            <hr/>
                            <div className="Single-task-footer">
                                {currentUser ? (
                                    <button className="Train"><Link to={"task/"+task.id}>Train</Link></button>
                                ) : (
                                    <></>
                                )}
                            </div>
                        </div>
                    )
                })}
            </div>
        </div>
    )
};

export default AllTasks;