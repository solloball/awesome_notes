import React from "react";
import axios from "axios";

class GetNotes extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            author: "empty",
            note: "empty",
            title: "empty"
        };

        let alias = window.location.href.toString().split(window.location.host)[1];
        axios.get("api" + alias).then((res) => {
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

export default GetNotes;
