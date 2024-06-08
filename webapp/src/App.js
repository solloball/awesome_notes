import './App.css';
import React from 'react';
import AddNote from './addNote';


export const Url = "http://localhost:8082/"

export class App extends React.Component {
 render() {
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
