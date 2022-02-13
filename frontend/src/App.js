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
      keyvaluepairs: [],
      errorMessage: ''
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
        this.setState({ keyvaluepairs: Object.entries(data), errorMessage: '' })
      })
      .catch((reason) => {
        console.log(reason);
        this.setState({errorMessage: 'Unable to connect to backend server at http://localhost:8080, reload to retry.'});
      });
  }

  render() {
    return (
      <div className="App">        
        <Header />
        {this.state.errorMessage && <div className="error"> {this.state.errorMessage}</div>}
        {!this.state.errorMessage && <KvpList keyvaluepairs={this.state.keyvaluepairs} deletedFunc={this.update}/>}
        {!this.state.errorMessage && <Add addedFunc={this.update}/>}
      </div>
    );
  }
}

export default App;
