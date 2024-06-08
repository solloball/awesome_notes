import './App.css';
import React from 'react';
import AddNote from './addNote';

export class App extends React.Component {
 constructor(props) {
     super(props);
 }
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
