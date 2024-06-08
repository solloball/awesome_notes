import React from "react";
import axios from "axios";
//import {Url} from "./App.js"
const Url = "http://localhost:8082"

class Note extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            author: "user",
            note: "note",
            title: "Test title"
        };

        let alias = window.location.href.toString().split(window.location.host)[1];
        axios.get(Url + alias).then((res) => {
            this.setState({
                author: res.data.Author,
                note: res.data.Note,
                title: res.data.Title
            });
        }).catch(function (error) {
            console.log(error)
        });

    }
    render() {
        return (
            <div>
            <h1>Title:</h1>
            <h2>{this.state.title}</h2>
            <h1>Note:</h1>
            <h2>{this.state.note}</h2>
            <h1>Author:</h1>
            <h2>{this.state.author}</h2>
            </div>
        );
    }
}

export default Note;
