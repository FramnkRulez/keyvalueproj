import React, { Component } from "react";
import Button from 'react-bootstrap/Button';

import "./Add.css";

class Add extends Component {
    constructor(props) {
        super(props);
        this.state = {key: '', value: ''};
    
        this.handleKeyChange = this.handleKeyChange.bind(this);
        this.handleValueChange = this.handleValueChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
      }

    handleKeyChange(event) {
        this.setState({key: event.target.value});
    }

    handleValueChange(event) {
        this.setState({value: event.target.value});
    }
    
    handleSubmit(event) {
        console.log(this.state);

        const requestOptions = {
            method: 'POST',
            body: JSON.stringify(this.state)
        };

        fetch('http://localhost:8080/kvp', requestOptions)
            .then(response => response.json())
            .then(data =>  {
                this.setState({ key: '', value: '' });
                this.props.addedFunc();
            })
            .catch(console.log);

        event.preventDefault();
    }

    render() {
        return (
            <div>
                <h3>Add new key/value pair</h3>

                Key: <input name="key" value={this.state.key} onChange={this.handleKeyChange}/>
                Value: <input name="value" value={this.state.value} onChange={this.handleValueChange}/>

                <Button onClick={this.handleSubmit}>Add</Button>
          </div>
        );
      }
    }

export default Add;