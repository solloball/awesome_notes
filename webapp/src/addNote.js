import React from "react";
import Url from "./App.js"

class AddNote extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            title: "",
            note: "",
            author: "",
            alias: ""
        }
    }

    render() {
        return (
            <form>
                <input placeholder="title" onChange={
                    (e) => this.setState({title: e.target.value})
                }/>
                <input placeholder="note" onChange={
                    (e) => this.setState({note: e.target.value})
                }/>
                <input placeholder="author" onChange={
                    (e) => this.setState({author: e.target.value})
                }/>
                <input placeholder="alias" onChange={
                    (e) => this.setState({alias: e.target.value})
                }/>
                <button type="button" onClick={() => {
                    if (this.state.title === ""
                        || this.state.note === ""
                        || this.state.author === ""
                    ) {
                        alert("Titile, note and author shouldn't be empty");
                        return;
                    }

                    let json = JSON.stringify({
                        record: {
                            Title: this.state.title,
                            Note: this.state.note,
                            Author: this.state.author
                        },
                        alias: this.state.alias,
                    });
                    console.log(json);

                    // send data to server
                    fetch(Url + "/record", {
                    method: "POST",
                    body: json,
                    headers: {
                        "Content-type": "application/json; charset=UTF-8"
                    }
                    })
                    .then((response) => response.json())
                    .then((json) => {
                        try {
                            alert("your alias is " + json.alias);
                        } catch {
                            alert("Failed to get json");
                        }
                    });

                    
                }}>Create Note</button>
            </form>
        )
    }
}

export default AddNote
