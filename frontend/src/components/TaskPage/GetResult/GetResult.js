import React, { useState, useEffect} from "react";
import {getResultInformation} from '../../../services/result.service'
import './GetResult.css'
import { CircularProgress } from "@material-ui/core";

const CheckResult = () => {
    const [result, setResult] = useState();

    const getResult = async () => {
        setResult(undefined)
        const id = sessionStorage.getItem("lastSolution")
        if (id===null) {
            return
        }

        getResultInformation(id).then(
            (result) => {
                setResult(result)
            },
            (error) => {
                const resMessage =
                    (error.response &&
                        error.response.data &&
                        error.response.data.message) ||
                    error.message ||
                    error.toString();
                console.log(resMessage)
            }
        );
    };

    useEffect(() => {
        getResult().then();
    }, []);

    return (
        <>
            {result ? (
                <div className="Results">
                    <h3>Tests passed: <span className={
                        result.passed_tests_count===result?.tests_count ? ("Correct") : ("Incorrect")
                    }>{result?.passed_tests_count}/{result?.tests_count}</span></h3>
                    <>
                        {result?.results ? (
                            result.results.map((res, id) => {
                                return <div className="Results-item" key={id}>{
                                    <>
                                        { result.results.length > 1 ? (<p className="Results-item-bold">{id+1+'. Basic test:'}</p>) : (<></>) }
                                        <p>
                                            Status: <span className={ res.status==="OK" ? ("Correct") : ("Incorrect") }>{res.status}</span>
                                        </p>
                                        <p>Time: {res.time+"ms"}</p>
                                    </>
                                }</div>;
                            })) : (
                            <></>
                        )}
                    </>
                </div>
            ) : (
                <CircularProgress color="inherit" className="Progress" />
            )}
        </>
    );
};

export default CheckResult;
