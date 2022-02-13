import React, { Component } from "react";
import "./KvpList.css";

class KvpList extends Component {
    constructor(props) {
        super(props);

        this.handleRemove = this.handleRemove.bind(this);
    }

    handleRemove(key) {
        console.log('Remove for ' + key);

        const requestOptions = {
            method: 'DELETE'
        };

        fetch('http://localhost:8080/kvp/' + key, requestOptions)
            .then(data =>  {
                this.props.deletedFunc();
            })
            .catch(console.log);
    }

  render() {
    const kvp = this.props.keyvaluepairs.map((pair, index) => (
      <li key={pair[0]}>
          <span>{pair[0]}={JSON.stringify(pair[1])}</span>
          <button onClick={() => this.handleRemove(pair[0])}>Delete</button>
      </li>
    ));

    return (
      <div className="KvpList">
        <h3>Key/Value Pair List</h3>
        <ul>
        {kvp}
        </ul>
      </div>
    );
  }
}

export default KvpList;