const API_URL = process.env.REACT_APP_SERVER_URL;

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function repeat(id) {
    await sleep(1000)
    return await getResultInformation(id)
}

export async function getResultInformation(taskID) {
    const url = new URL(taskID, API_URL+'/solution/').toString()
    const token = JSON.parse(localStorage.getItem("token")).Authorization
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
        )
        const st = response.status
        if(st === 200 || st === 400 || st === 408 ) {
            const json = await response.json()
            if(json?.message==="pending") {
                return await repeat(taskID)
            } else {
                return json
            }
        }
    } catch (error) {
        sessionStorage.removeItem("lastSolution")
        return error
    }
}
