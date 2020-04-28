import React from 'react';
import { Provider } from 'react-redux';
import './App.css';

import Home from './Home';
import store from './lib/store';

function App() {
  return (
    <Provider store={store}>
      <Home />
    </Provider>
  );
}

export default App;