
const HOST = location.protocol + "//" + window.location.hostname
+ (window.location.port ? ":" + window.location.port : "");

const API_LIST_QUOTES = HOST + "/api/listquotes";

function isEmptyObj(obj){
    return (obj == null) || (Object.keys(obj).length === 0 && obj.constructor === Object)
}

async function getJSONData(url = '', qryParams = {}) {
    if (!url) {
        throw new Error("URl cannot be empty!!")
    }
    const urlObj = new URL(url)
    if (!isEmptyObj(qryParams)){
        urlObj.search = new URLSearchParams(qryParams).toString();
    }
    const response = await fetch(urlObj)
    return await response.json()
}

function loadQuotes(){
    const quotes = getJSONData(API_LIST_QUOTES, null);
    quotes.then(respData => {
        if (respData.error){
            console.log(respData.msg)
        } else {
            // populate table with response data
            const table = document.getElementById("quotes");
            emptyTable(table);
            populateTable(table, respData);
        }
    }).catch(err => {
            console.log(err);
    });
}

function emptyTable(table){
    const rowCount = table.rows.length;
    for (let i=rowCount-1; i > 0; i--){
        table.deleteRow(i);
    }
}

function populateTable(table, data){
    let rc = table.rows.length;
    for(let rec of data){
        const row = table.insertRow(rc);
        rc+=1;
        const cell1 = row.insertCell(0);
        const cell2 = row.insertCell(1);
        const cell3 = row.insertCell(2);
        cell1.className = "col_1";
        cell2.className = "col_2";
        cell1.innerHTML = `"${rec.Quote}"`;
        cell2.innerHTML = rec.Author;
    }
}

// things to do on load
window.onload = () => {
    loadQuotes();
}