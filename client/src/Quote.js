import React from "react";
import axios from "axios";
import {
    Button,
    Input,
    Form,
    TextArea
} from "semantic-ui-react";

let endpoint = "http://localhost:8080";

function GetQuote(props) {
    return (
        <Button 
            primary
            className="get-quote" 
            onClick={props.onClick}>
                Fetch!
        </Button>
    );
}

class QuoteBuilder extends React.Component {
    constructor() {
        super()

        this.state = {
            quote: "",
            author: ""
        }

        this.handleInputChange = this.handleInputChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleInputChange(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;

        this.setState({
            [name]: value
        });
    }

    handleSubmit(event) {
        console.log(this.state);
    }

    render() {
        return(
            <div>
                <Form onSubmit={this.handleSubmit}>
                    <div>
                        <TextArea 
                            placeholder="Enter a new quote"
                            name="quote"
                            value={this.state.quote}
                            onChange={this.handleInputChange} />
                    </div>
                    <Input
                        type="text"
                        name="author"
                        value={this.state.author}
                        onChange={this.handleInputChange} />
                    <Input type="submit" value="Submit" />
                </Form>
            </div>
        );
    }
}

class Quote extends React.Component {
    constructor() {
        super()

        this.state = {
            quote: "",
            author: ""
        }
    }

    componentDidMount() {
        this.fetchQuote();
    }

    fetchQuote = () => {
        axios.get(endpoint + "/quotes").then(response => {
            console.log(response);
            if (response.status === 200) {
                this.setState({
                    quote: response.data.quote,
                    author: response.data.author
                })
            } else {
                console.log("Something went wrong with the response");
            }
        });
    }

    handleClick() {
        this.fetchQuote();
    }

    render() {
        return (
            <section className="quote">
                <div className="quote-content">
                    <h2>"{this.state.quote}"</h2>
                </div>
                <div className="quote-author">
                    <h3>-{this.state.author}</h3>
                </div>
                <div className="fetch-button">
                    <GetQuote
                        onClick={() => this.handleClick()}
                    />
                </div>
            </section>
        )
    }
}

export {
    Quote,
    QuoteBuilder,
}