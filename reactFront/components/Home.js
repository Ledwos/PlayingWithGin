import React from 'react';


class Home extends React.Component {
    constructor(props) {
        super(props);
        this.authenticate = this.authenticate.bind(this);
      }
    authenticate() {
    this.WebAuth = new auth0.WebAuth({
        domain: process.env.AUTH0_DOMAIN,
        clientID: process.env.AUTH0_CLIENT_ID,
        scope: "openid profile",
        audience: process.env.AUTH0_API_AUDIENCE,
        responseType: "token id_token",
        redirectUri: process.env.AUTH0_CALLBACK_URL
    });
    this.WebAuth.authorize();
    }

    render() {
        return (
        <div className="container">
            <div className="row">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Jokeish</h1>
                    <p>A load of Dad jokes XD</p>
                    <p>Sign in to get access </p>
                    <a onClick={this.authenticate} 
                    className="btn btn-primary btn-lg btn-login btn-block"
                    >
                    Sign In
                    </a>
              </div>
            </div>
          </div>
        )
    }
}

export default Home;

