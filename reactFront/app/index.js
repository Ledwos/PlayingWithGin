import React from 'React';
import ReactDOM from 'react-dom';
require('dotenv').config();
import './index.css';
import Home from '../components/Home';
import LoggedIn from './components/LoggedIn';

class App extends React.Component {
  
  parseHash() {
    this.auth0 = new auth0.WebAuth({
      domain: process.env.AUTH0_DOMAIN,
      clientID: process.env.AUTH0_CLIENT_ID
    });
    this.auth0.parseHash(window.location.hash, (err, authResult) => {
      if (err) {
        return console.log(err);
      }
      if (
        authResult !== null &&
        authResult.accessToken !== null &&
        authResult.idToken !== null
      ) {
        localStorage.setItem("access_token", authResult.accessToken);
        localStorage.setItem("id_token", authResult.idToken);
        localStorage.setItem("profile", JSON.stringify(authResult.idTokenPayload));
        window.location = window.location.href.substr(0, window.location.href.indexOf("#"));
      }
    });  
  }

  setup() {
    $.ajaxSetup({
      beforeSend: (r) => {
        if (localStorage.getItem("access_token")) {
          r.setRequestHeader(
            "Authorization",
            "Bearer " + localStorage.getItem("access_token")
          );
        }
      }
    });
  }

  setState() {
    let idToken = localStorage.getItem("id_token");
    if (idToken) {
      this.loggedIn = true;
    } else {
      this.loggedIn = false;
    }
  }

  componentWillMount() {
    this.setup();
    this.parseHash();
    this.setState();
  }


  render(){
    if (this.loggedIn) {
        return (<LoggedIn />);
      } else {
        return (<Home />);
      }
  }
}

ReactDOM.render(<App />, document.getElementById('app'));