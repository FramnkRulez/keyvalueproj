import React, { Component } from "react";
import "./App.css";
import 'bootstrap/dist/css/bootstrap.min.css';

//import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import KvpList from "./components/KvpList";
import Add from "./components/Add";

class App extends Component {
  constructor(props) {
    super(props);

    this.update = this.update.bind(this);

    this.state = {
      keyvaluepairs: []
    }
  }

  componentDidMount() {
    this.update();
  }

  update() {
    console.log("updating kvps");

    fetch('http://localhost:8080')
      .then(res => res.json())
      .then((data) => {
        this.setState({ keyvaluepairs: Object.entries(data) })
      })
      .catch(console.log)
  }

  render() {
    return (
      <div className="App">        
        <Header />
        <KvpList keyvaluepairs={this.state.keyvaluepairs} deletedFunc={this.update}/>
        <Add addedFunc={this.update}/>
      </div>
    );
  }
}

export default App;
