import React, { PropTypes } from 'react'

class Container extends React.Component {
	render(){
		return (<div><UPCInput /></div>)
	}
}

class UPCInput extends React.Component {
	constructor(props){
		super(props)
		this.state = {value: ''}
		this.handleChange = this.handleChange.bind(this)
		this.handleSubmit = this.handleSubmit.bind(this)
	}
	handleChange(event) {
		this.setState({value: event.target.value})
	}
	handleSubmit(event){
		fetch(`/marc?upc=${this.state.value}`)
		event.preventDefault()
	}
	render(){
		return (
			<form onSubmit={this.handleSubmit}>
				<input type="text" id="upc" value={this.state.value} onChange={this.handleChange}/>
			</form>
		)
	}
}


export default Container;