import './App.css';
import React from 'react';
import AddNotes from './notes/AddNotes';
import GetNotes from './notes/GetNotes';


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
            <GetNotes />
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
        <AddNotes />
      </aside>
    </div>
  );
 }
}

export default App;
