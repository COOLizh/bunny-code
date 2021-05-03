import React, {useEffect, useState} from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom';
import './App.css';
import AllTasks from './components/AllTasks/AllTasks.js';
import TaskPage from "./components/TaskPage/TaskPage";
import AuthService from "./services/auth.service";
import Login from "./components/Auth/Login/Login";
import Register from "./components/Auth/Register/Register";
import Home from "./components/Home";
import Profile from "./components/Profile";

function App() {
    const [currentUser, setCurrentUser] = useState(undefined);

    useEffect(() => {
        const user = AuthService.getCurrentUser();

        if (user) {
            setCurrentUser(user);
        }
    }, []);

    const logOut = () => {
        AuthService.logout();
        setCurrentUser(undefined)
    };

  return (
    <div className="App">
        <BrowserRouter>
            <header className="App-header">
                <nav>
                    <ul className="Menu">
                        {currentUser ? (
                            <div>
                                <li><Link to="/">Home</Link></li>
                                <li><Link to="/tasks">All tasks</Link></li>
                            </div>
                        ) : (
                            <div>
                                <li><Link to="/">Home</Link></li>
                                <li><Link to="/tasks">All tasks</Link></li>
                            </div>
                            )}
                    </ul>
                </nav>
                <ul className="Login-header">
                    {currentUser ? (
                        <div>
                            <li><Link to="/profile">{currentUser}</Link></li>
                            <li><Link to="/login" onClick={logOut}>Logout</Link></li>
                        </div>
                    ) : (
                        <div>
                            <li><Link to="/login">Sign in</Link></li>
                            <li><Link to="/registration">Sign up</Link></li>
                        </div>
                    )}
                </ul>
            </header>
            <Switch>
                <Route path="/tasks">
                    <AllTasks />
                </Route>
                <Route exact path="/" component={Home} />
                <Route exact path="/login" component={Login} />
                <Route exact path="/registration" component={Register} />
                <Route exact path="/profile" component={Profile} />
                <Route path="/task/:id" component={TaskPage} />
            </Switch>
        </BrowserRouter>
    </div>
  );
}

export default App;
