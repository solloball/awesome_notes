import React from "react";
import axios from "axios";
import Url from "./Constants.js"

class Note extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            author: "user",
            note: "note",
            title: "Test title"
        };

        let alias = window.location.href.toString().split(window.location.host)[1];
        axios.get(Url.Url + alias).then((res) => {
            if (res.data.status === "Error") {
                this.setState({
                    author: "Author not found",
                    note: "Note not found",
                    title: "Title not found"
                });
                return;
            }
            this.setState({
                author: res.data.Author,
                note: res.data.Note,
                title: res.data.Title
            });
        })

    }
    render() {
        return (
            <div>
            <div>{this.state.title}</div>
            <h2>Note:</h2>
            <div>{this.state.note}</div>
            <h4>Author:</h4>
            <div>{this.state.author}</div>
            </div>
        );
    }
}

export default Note;
