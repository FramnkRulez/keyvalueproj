import React, { Component } from "react";
import "./KvpList.css";

class KvpList extends Component {
  render() {
    const kvp = this.props.keyvaluepairs.map((pair, index) => (
      <p key={pair[0]}>{pair[0]}={JSON.stringify(pair[1])}</p>
    ));

    return (
      <div className="KvpList">
        <h2>Key/Value Pairs</h2>
        {kvp}
      </div>
    );
  }
}

export default KvpList;