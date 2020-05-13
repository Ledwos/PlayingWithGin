import React from 'react';
import Joke from './Joke';

class LoggedIn extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      jokes: []
    };

    this.serverRequest = this.serverRequest.bind(this);
    this.logout = this.logout.bind(this);
  }

  logout() {
    localStorage.removeItem("id_token");
    localStorage.removeItem("access_token");
    localStorage.removeItem("profile");
    location.reload();
  }
  
  serverRequest() {
    $.get("http://localhost:3000/api/jokes", res => {
      this.setState({
        jokes: res
      });
    });
  }

  componentDidMount() {
    this.serverRequest();
  }

    render() {
      return (
        <div className="container">
          <br />
          <span className="pull-right">
            <a onClick={this.logout}>Log out</a>
          </span>
          <h2>Jokeish</h2>
          <p>Let's feed you with some funny Jokes!!!</p>
          <div className="row">
            <div className="container">
              {this.state.jokes.map(function(joke, i) {
                return <Joke key={i} joke={joke} />;
              })}
            </div>
          </div>
      </div>
      )
    }
}

export default LoggedIn;