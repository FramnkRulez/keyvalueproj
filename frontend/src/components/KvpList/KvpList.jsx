import React, { Component } from "react";
import Button from 'react-bootstrap/Button';
import Table from 'react-bootstrap/Table'

import "./KvpList.css";

class KvpList extends Component {
    constructor(props) {
        super(props);

        this.state = {errorMessage: ''};

        this.handleRemove = this.handleRemove.bind(this);
    }

    handleRemove(key) {
        console.log('Remove for ' + key);

        const requestOptions = {
            method: 'DELETE'
        };

        fetch('http://localhost:8080/kvp/' + encodeURIComponent(key), requestOptions)
            .then(data =>  {
                this.props.deletedFunc();
                this.setState({errorMessage: ''});
            })
            .catch(reason => {
                console.log(reason);
                this.setState({errorMessage: 'Unable to delete key, verify server is running'});
            });
    }

  render() {
    const kvp = this.props.keyvaluepairs.map((pair, index) => (
        <tr key={pair[0]}>
            <td className="col-1">{index+1}</td>
            <td className="col-3">{pair[0]}</td>
            <td className="col-3">{JSON.stringify(pair[1])}</td>
            <td className="col-1"><Button onClick={() => this.handleRemove(pair[0])}>Delete</Button></td>
        </tr>
    ));

    return (
      <div className="KvpList">
        <h3>Key/Value Pair List</h3>

        <Table striped bordered hover>
            <thead>
                <tr>
                    <th>#</th>
                    <th>Key</th>
                    <th>Value</th>
                    <th></th>
                </tr>
            </thead>

            <tbody>
                {kvp}
            </tbody>
        </Table>
        {this.state.errorMessage && <div className="error"> {this.state.errorMessage}</div>}
      </div>
    );
  }
}

export default KvpList;