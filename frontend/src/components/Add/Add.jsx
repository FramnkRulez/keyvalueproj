import React, { Component } from "react";
import Button from 'react-bootstrap/Button';

import "./Add.css";

class Add extends Component {
    constructor(props) {
        super(props);
        this.state = {key: '', value: '', errorMessage: ''};
    
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

        if(this.state.key.includes('/')) {
            this.setState({errorMessage: "Invalid character '/' found in key"});
        } else {
            const requestOptions = {
                method: 'POST',
                body: JSON.stringify({key: this.state.key, value: this.state.value})
            };

            fetch('http://localhost:8080/kvp', requestOptions)
                .then(response => response.json())
                .then(data =>  {
                    this.setState({ key: '', value: '' , errorMessage: ''});
                    this.props.addedFunc();
                })
                .catch(reason => {
                    console.log(reason);
                    this.setState({errorMessage: 'Unable to add key/value, verify server is running'});
                });
            }

        event.preventDefault();
    }

    render() {
        return (
            <div>
                <h3>Set key/value pair</h3>

                Key: <input name="key" value={this.state.key} onChange={this.handleKeyChange}/>
                Value: <input name="value" value={this.state.value} onChange={this.handleValueChange}/>

                <Button onClick={this.handleSubmit}>Add</Button>
                {this.state.errorMessage && <div className="error"> {this.state.errorMessage}</div>}
          </div>
        );
      }
    }

export default Add;