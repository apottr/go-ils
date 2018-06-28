import React, { PropTypes } from 'react'
import { Input, Card } from 'semantic-ui-react'
import {
	BrowserRouter as Router,
	Route,
	Link
} from 'react-router-dom'

const ConRouter = () => (
	<Router>
		<div>
		<Route path="/" component={Container} />
		<Route path="/:marcid" component={Viewer} />
		</div>
	</Router>
)

const Viewer = ({match}) => (
	<h1>Hello {match.params.marcid}!</h1>
)

class Container extends React.Component {
	render(){
		return (<div><List /><UPCInput /></div>)
	}
}

class List extends React.Component {
	constructor(props){
		super(props)
		this.state = {records: []}
	}
	componentDidMount() {
		this.interval = setInterval(() => {
			fetch("/marc")
				.then(r => r.json())
				.then(d => this.setState({records: d}))
		}, 5000);
	}
	componentWillUnmount() {
		clearInterval(this.interval);
	}
	render(){
		let a = this.state.records.map(e => (<li>
			<Card>
				<Card.Content>
					<Card.Header><Link to={`/${e}`}>{e}</Link></Card.Header>
				</Card.Content>
			</Card>
			</li>))
		return (
			<ul>
				{a}
			</ul>
		)
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
		this.setState({value: ""})
		event.preventDefault()
	}
	render(){
		return (
			<form onSubmit={this.handleSubmit}>
				<Input defaultValue={this.state.value} onChange={this.handleChange}/>
			</form>
		)
	}
}


export default ConRouter;
