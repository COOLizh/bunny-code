import React, { useState} from 'react';
import { useParams, useHistory} from 'react-router-dom';
import Select from '@material-ui/core/Select';
import AceEditor from "react-ace";
import "ace-builds/src-noconflict/mode-golang";
import "ace-builds/src-noconflict/theme-twilight";
import './SolutionFied.css';
import { makeStyles, MenuItem } from "@material-ui/core";

const API_URL = process.env.REACT_APP_SERVER_URL;

async function sendSolution(solution, id) {
    const url = new URL(id,API_URL+'/solution/task/').toString()
    const token = JSON.parse(localStorage.getItem("token")).Authorization
    try {
        await fetch(
            url,
            {
                    method: 'POST',
                    headers: {
                        'Authorization': 'Bearer ' + token,
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(solution)
            }
        ).then(function (response) {
            return response.json();
        }).then(function (data) {
            sessionStorage.setItem("lastSolution", data.id)
        });
    } catch (error) {
        console.log(error)
    }
}

const useStyles = makeStyles({
    root: {
        'margin': '0 10px',
        'padding': '0',
        'color': 'white',
        'font-size': 'inherit',
        'font-family': 'inherit',
        '& svg': {
            color: 'white',
        }
    },
});

export default function SendTaskSolution() {
    const [solution, setSolution] = useState();
    const [language, setLanguage] = useState();
    const { id } = useParams();
    let history = useHistory();
    const classes = useStyles();

    const handleChange = (event) => {
        setLanguage(event.target.value);
    };

    const handleSubmit = async e => {
        if(language !== undefined) {
            e.preventDefault(true);
            sessionStorage.removeItem("lastSolution",)
            await sendSolution({
                solution,
                language
            }, id);
            history.push("/task/"+id+"/")
            history.push("/task/"+id+"/output")
        } else {
            alert("Please, select language")
        }
    }

    return(
        <div className="SendSol-wrapper">
            <div className="SendSol-header">
                <p>Solution:</p>
                <Select
                    className={classes.root}
                    disableUnderline={true}
                    defaultValue={' '}
                    value={language ? language : " "}
                    onChange={handleChange}
                    displayEmpty
                    inputProps={{ 'aria-label': 'Without label' }}
                >
                    <MenuItem value={' '} disabled>Select a language</MenuItem>
                    <MenuItem value={"golang"}>Go</MenuItem>
                    <MenuItem value={"c++"}>C++</MenuItem>
                </Select>
            </div>
            <form onSubmit={handleSubmit}>
                <AceEditor
                    width="100%"
                    height="calc(80vh - 80px)"
                    className="Solution-textarea"
                    mode="golang"
                    theme="twilight"
                    onChange={code => setSolution(btoa(code))}
                    name="Ace-input"
                    editorProps={{ $blockScrolling: true }}
                    setOptions={{
                        fontSize: '0.9rem',
                        showPrintMargin: false,
                        minLines: 20
                    }}
                />
                <div className="SendSol-footer">
                    <button type="submit" className="SendSol-button">Submit</button>
                </div>
            </form>
        </div>
    )
}