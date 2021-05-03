import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import './TaskHistory.css'
import autosize from 'autosize';

const API_URL = process.env.REACT_APP_SERVER_URL;

const GetHistory = () => {
    const [items, setItems] = useState([]);
    const { id } = useParams();
    const url = new URL(id, API_URL+'/solution/task/').toString() + '/history'
    const token = JSON.parse(localStorage.getItem("token")).Authorization

    const getHistory = async () => {
        try {
            const response = await fetch(
                url,
                {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer ' + token,
                        'Content-Type': 'application/json',
                    },
                }
            );

            const responseData = await response.json();
            if(response.status!==200) {
                console.log(response)
            } else {
                setItems(responseData);
            }
        } catch (error) {
            console.error(error);
        }
    };

    useEffect(() => {
        getHistory().then();
    }, []);

    const printLang = function (lang) {
        switch (lang) {
            case 'golang': return 'Go';
            case 'c++': return 'C++';
            default: return ' ';
        }
    }

    return (
        <div className="History">
            {items && items.map(function(item, index) {
                return(
                    <div className="History-item" key={index}>
                        { index > 0 ? (<hr/>) : (<></>) }
                        <div className="History-item-header">
                            <p>{'Language: '+printLang(item?.solution.language)}</p>
                            <p>{(new Date(item?.created_at)).toLocaleString()}</p>
                        </div>
                        <p>Result: {item?.result.results && item?.result.results[0].status}</p>
                        <p className="Task-id">ID: {item?.id}</p>
                        <textarea ref={c=>autosize(c)} disabled={true} value={atob(item?.solution.solution)} className={
                            item?.result.results ? ((item.result.results[0].status==='OK') ? ('Border-correct') : ('Border-incorrect')) : ("")
                        }> </textarea>
                    </div>
                )
            })}
        </div>
    )
}

export default GetHistory;