import React, { Component } from "react";
import { connect } from 'react-redux';
import "./App.css";

class App extends Component {
  render() {
    console.log(this.props.drones)
    return (
      <div className="container">
        {this.props.drones.length > 0 && <h1 className="header">Drones</h1>}
        <div className="drones">
        {this.props.drones.map(d => 
          <div className="drone" key={d.id}>
            <label>ID: {d.id}</label>
            <label>X: {d.x}</label>
            <label>Y: {d.y}</label>
            <label>Speed: {d.s}</label>
            <label>Active: {d.active ? "true" : "false"}</label>
          </div>
        )}
        </div>
      </div>
    );
  }
}

export default connect(state => ({
  drones: state.drones
}), {})(App);