import React from 'react';
import {render} from 'react-dom';
import {Provider} from 'react-redux';
import {
    BrowserRouter as Router,
    Route
} from 'react-router-dom';
import configureStore from './store/configureStore';
import './index.less';
import App from './containers/App';
import registerServiceWorker from './registerServiceWorker';
import 'antd/dist/antd.less';

const store = configureStore();

render(
    <Provider store={store}>
        <Router>
            <div>
                <Route component={App}/>
            </div>
        </Router>
    </Provider>,
    document.getElementById('root')
)
registerServiceWorker();
