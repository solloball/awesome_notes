import './App.css';
import React from 'react';
import AddNote from './addNote';
import Note from './Note';

// TODO: add to config
//const Url = "http://localhost:8082"

class App extends React.Component {
 render() {
  let alias = window.location.href.toString().split(window.location.host)[1];

  // TODO: refactor this 
  if (alias.length > 1) {
    return (
        <div className="App">
            <header className="App-header">
            <div>Pretty note</div>
            </header>
            <aside>
            <Note />
            </aside>
            </div>
        );
  }

  console.log(alias.length);
  return (
    <div className="App">
      <header className="App-header">
      <div>Pretty note</div>
      </header>
      <aside>
        <AddNote />
      </aside>
    </div>
  );
 }
}

export default App;
