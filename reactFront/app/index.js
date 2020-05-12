import React from 'React'
import ReactDOM from 'react-dom'
import './index.css'

class App extends React.Component {
    render(){
        return(
            <div>
                <h2>Hey, it's working :D</h2>
                <p>yup</p>
            </div>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'))