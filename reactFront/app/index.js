import React from 'React';
import ReactDOM from 'react-dom';
import './index.css';
import Home from '../components/Home';

class App extends React.Component {
    render(){
        if (this.loggedIn) {
            return (<LoggedIn />);
          } else {
            return (<Home />);
          }
    }
}

ReactDOM.render(<App />, document.getElementById('app'));